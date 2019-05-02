package cli

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/030/go-utils"
	log "github.com/sirupsen/logrus"
)

func setup() {
	log.SetReportCaller(true)

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

	defer cleanupFiles("testFiles/file*")
}

func shutdown() {
	cmd := exec.Command("bash", "-c", "docker stop nexus && docker rm nexus")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}
	cleanupFiles("download/file*/file*/*/file*")
	cleanupFiles("download/file*/file*/maven-metadata*")
}

func (n Nexus3) pong() bool {
	pongAvailable := false

	req, err := http.NewRequest("GET", n.URL+"/service/metrics/ping", nil)
	if err != nil {
		// handle err
	}
	req.SetBasicAuth(n.User, n.Pass)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
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
	cmd := exec.Command("bash", "-c", "curl -u "+n.User+":"+n.Pass+" -X POST \""+n.URL+"/service/rest/v1/components?repository=maven-releases\" -H  \"accept: application/json\" -H  \"Content-Type: multipart/form-data\" -F \"maven2.asset1=@"+path+".pom\" -F \"maven2.asset1.extension=pom\" -F \"maven2.asset2=@"+path+".jar\" -F \"maven2.asset2.extension=jar\"")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}
}

func createPOM(d string, f string, number string) {
	createArtifact(d, f+".pom", "<project>\n<modelVersion>4.0.0</modelVersion>\n<groupId>file"+number+"</groupId>\n<artifactId>file"+number+"</artifactId>\n<version>1.0.0</version>\n</project>")
}

func createJAR(d string, f string, number string) {
	createArtifact(d, f+".jar", "some-content")
}

func (n Nexus3) createArtifactsAndSubmit(i int) {
	number := strconv.Itoa(i)
	f := "file" + number
	createPOM("testFiles", f, number)
	createJAR("testFiles", f, number)
	n.submitArtifact("testFiles", f)
}

func cleanupFiles(re string) {
	files, err := filepath.Glob(re)
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0 {
		log.Fatal("No files to be removed were found")
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			log.Fatal(err)
		}
	}
}

func allFiles(dir string) ([]string, error) {
	fileList := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f.Mode().IsRegular() {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList, err
}
