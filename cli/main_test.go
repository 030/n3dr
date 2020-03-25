package cli

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/030/go-utils"
	log "github.com/sirupsen/logrus"

	mp "github.com/030/go-multipart/utils"
)

// See https://stackoverflow.com/a/34102842/2777965
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

const (
	testFilesDir = "testFiles"
)

var n = Nexus3{
	URL:        "http://localhost:9999",
	User:       "admin",
	Pass:       "admin123",
	Repository: "maven-releases",
	APIVersion: "v1",
}

func setup() {
	// Start docker nexus
	cmd := exec.Command("bash", "-c", "docker run -d -p 9999:8081 --name nexus sonatype/nexus3:3.16.1")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}

	pingURL := n.URL + pingURI

	// Waitfor ping URL to become available
	for !utils.URLExists(pingURL) {
		log.Info("Nexus not available.")
		time.Sleep(30 * time.Second)
	}

	// Waitfor ping endpoint to return pong
	for !n.pong() {
		log.Info("Nexus Pong not returned yet.")
		time.Sleep(3 * time.Second)
	}

	// Send test artifacts to docker nexus
	for i := 1; i <= 3; i++ {
		n.createArtifactsAndSubmit(i)
	}
}

func shutdown() {
	cmd := exec.Command("bash", "-c", "docker stop nexus && docker rm nexus")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}

	testFiles := filepath.Join(testFilesDir, "/file*")
	testDownloads := filepath.Join("download", n.Repository, "file*", "file*", "*", "file*")
	testDownloadsMetadata := filepath.Join("download", n.Repository, "file*", "file*", "maven-metadata*")
	cleanupFilesSlice := []string{testFiles, testDownloads, testDownloadsMetadata}
	for _, f := range cleanupFilesSlice {
		err := cleanupFiles(f)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (n Nexus3) pong() bool {
	pongAvailable := false

	req, err := http.NewRequest("GET", n.URL+"/service/metrics/ping", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(n.User, n.Pass)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)

		if bodyString == "pong\n" {
			pongAvailable = true
		}
	}
	return pongAvailable
}

func (n Nexus3) submitArtifact(d string, f string) {
	path := filepath.Join(d, f)
	url := n.URL + "/service/rest/v1/components?repository=" + n.Repository
	u := mp.Upload{URL: url, Username: n.User, Password: n.Pass}
	err := u.MultipartUpload("maven2.asset1=@" + path + ".pom,maven2.asset1.extension=pom,maven2.asset2=@" + path + ".jar,maven2.asset2.extension=jar")
	if err != nil {
		log.Fatal(err)
	}
}

func createPOM(d string, f string, number string) {
	if err := createArtifact(d, f+".pom", "<project>\n<modelVersion>4.0.0</modelVersion>\n<groupId>file"+number+"</groupId>\n<artifactId>file"+number+"</artifactId>\n<version>1.0.0</version>\n</project>"); err != nil {
		log.Fatal(err)
	}
}

func createJAR(d string, f string) {
	if err := createArtifact(d, f+".jar", "some-content"); err != nil {
		log.Fatal(err)
	}
}

func (n Nexus3) createArtifactsAndSubmit(i int) {
	number := strconv.Itoa(i)
	f := "file" + number
	createPOM(testFilesDir, f, number)
	createJAR(testFilesDir, f)
	n.submitArtifact(testFilesDir, f)
}

func cleanupFiles(re string) error {
	files, err := filepath.Glob(re)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return errors.New("No files to be removed were found in: '" + re + "'")
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return err
		}
	}
	return nil
}

func allFiles(dir string) ([]string, error) {
	fileList := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if f.Mode().IsRegular() {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList, err
}
