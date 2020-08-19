package cli

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	log "github.com/sirupsen/logrus"

	mp "github.com/030/go-multipart/utils"
	"github.com/030/mij"
)

const (
	testDirDownload    = "/download"
	testDirUpload      = "/testFiles"
	testNexusAuthError = "ResponseCode: '401' and Message '401 Unauthorized' for URL: http://localhost:9999/service/rest/v1/repositories"
)

var n = Nexus3{
	URL:        "http://localhost:9999",
	User:       "admin",
	Pass:       "admin123",
	Repository: "maven-releases",
	APIVersion: "v1",
}

func TestMain(m *testing.M) {
	d := mij.DockerImage{
		Name:                     "sonatype/nexus3",
		PortExternal:             9999,
		PortInternal:             8081,
		Version:                  "3.16.1",
		ContainerName:            "nexus",
		LogFile:                  "/nexus-data/log/nexus.log",
		LogFileStringHealthCheck: "Started Sonatype Nexus OSS",
	}

	setup(&d)
	code := m.Run()
	shutdown(&d)
	os.Exit(code)
}

func setup(m *mij.DockerImage) {
	m.Run()

	// Send test artifacts to docker nexus once it is healthy
	for i := 1; i <= 3; i++ {
		n.createArtifactsAndSubmit(i)
	}
}

func shutdown(m *mij.DockerImage) {
	m.Stop()

	testFiles := filepath.Join(testDirHome, testDirUpload, "/file*")
	testDownloads := filepath.Join(testDirHome, testDirDownload, n.Repository, "file*", "file*", "*", "file*")
	testDownloadsMetadata := filepath.Join(testDirHome, testDirDownload, n.Repository, "file*", "file*", "maven-metadata*")
	cleanupFilesSlice := []string{testFiles, testDownloads, testDownloadsMetadata}
	for _, f := range cleanupFilesSlice {
		err := cleanupFiles(f)
		if err != nil {
			log.Fatal(err)
		}
	}

	cleanupDirsSlice := []string{tmpDir, testDirHome}
	for _, d := range cleanupDirsSlice {
		if err := os.RemoveAll(d); err != nil {
			log.Fatal(err)
		}
	}
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
	if err := createArtifact(d, f+".pom", "<project>\n<modelVersion>4.0.0</modelVersion>\n<groupId>file"+number+"</groupId>\n<artifactId>file"+number+"</artifactId>\n<version>1.0.0</version>\n</project>", "ba1f2511fc30423bdbb183fe33f3dd0f"); err != nil {
		log.Fatal(err)
	}
}

func createJAR(d string, f string) {
	if err := createArtifact(d, f+".jar", "some-content", "eae7286512c52715673c3878c10d2d55"); err != nil {
		log.Fatal(err)
	}
}

func (n Nexus3) createArtifactsAndSubmit(i int) {
	number := strconv.Itoa(i)
	f := "file" + number
	createPOM(testDirHome+"/"+testDirUpload, f, number)
	createJAR(testDirHome+"/"+testDirUpload, f)
	n.submitArtifact(testDirHome+"/"+testDirUpload, f)
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
