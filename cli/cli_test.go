package cli

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/030/go-utils"
)

var n = Nexus3{
	URL:  "http://localhost:8081",
	User: "admin",
	Pass: "admin123",
}

// See https://stackoverflow.com/a/34102842/2777965
func TestMain(m *testing.M) {
	setup()
	m.Run()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func (n Nexus3) TestDownloadedFiles(t *testing.T) {
	n.downloadTestArtifact("maven-releases", "file20", "1.0.0", "pom")
	n.downloadTestArtifact("maven-releases", "file20", "1.0.0", "jar")

	files := []string{"downloaded-file20-1.0.0.pom", "downloaded-file20-1.0.0.jar"}
	for _, f := range files {
		if !utils.FileExists(f) {
			t.Errorf("File %s should exist, but does not.", f)
		}
	}
	defer cleanupFiles("downloaded-*")
}

func TestContinuationToken(t *testing.T) {
	continuationTokenHashMap := map[string]string{
		"null": "35303a6235633862633138616131326331613030356565393061336664653966613733",
		"35303a6235633862633138616131326331613030356565393061336664653966613733":   "3130303a6235633862633138616131326331613030356565393061336664653966613733",
		"3130303a6235633862633138616131326331613030356565393061336664653966613733": "3135303a6235633862633138616131326331613030356565393061336664653966613733",
		"3135303a6235633862633138616131326331613030356565393061336664653966613733": "null",
	}

	for token, expectedContinuationToken := range continuationTokenHashMap {
		actual := n.continuationToken(token)

		if actual != expectedContinuationToken {
			t.Errorf("ContinuationToken incorrect. Expected %s, but was %s.", expectedContinuationToken, actual)
		}
	}
}

func TestContinuationTokenHash(t *testing.T) {
	expected := []string{"null",
		"3135303a6235633862633138616131326331613030356565393061336664653966613733",
		"3130303a6235633862633138616131326331613030356565393061336664653966613733",
		"35303a6235633862633138616131326331613030356565393061336664653966613733"}
	actual := n.continuationTokenRecursion("null")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Maps not equal. Expected %s, but was %s.", expected, actual)
	}

	actualSize := len(actual)
	expectedSize := 4
	if expectedSize != actualSize {
		t.Errorf("Not equal. Expected: %d. Actual: %d.", expectedSize, actualSize)
	}
}

func TestDownloadURLs(t *testing.T) {
	actual := len(n.downloadURLs())
	expected := 960 //160 artifacts * 6 different files, e.g. pom, pom.md5, pom.sha1, jar, jar.md5, jar.sha1
	if expected != actual {
		t.Errorf("Not equal. Expected: %d. Actual: %d.", expected, actual)
	}
}

func TestStoreArtifactsOnDisk(t *testing.T) {
	n.StoreArtifactsOnDisk()

	files, _ := ioutil.ReadDir("download")

	actual := len(files)
	expected := 961 //+1 due to .gitkeep

	if expected != actual {
		t.Errorf("Not equal. Expected: %d. Actual: %d.", expected, actual)
	}
}
