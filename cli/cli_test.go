package cli

import (
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/030/go-utils"
	log "github.com/sirupsen/logrus"
)

// See https://stackoverflow.com/a/34102842/2777965
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	cmd := exec.Command("bash", "-c", "docker run -d -p 8081:8081 --name nexus sonatype/nexus3:3.16.1")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}

	// Waitfor ping URL to become available
	for !utils.URLExists(pingURL) {
		log.Info("Nexus not available.")
		time.Sleep(30 * time.Second)
	}

	// Waitfor ping endpoint to return pong
	for !pong() {
		log.Info("Nexus Pong not returned yet.")
		time.Sleep(3 * time.Second)
	}

	// Send test artifacts to docker nexus
	for i := 1; i <= 160; i++ {
		createArtifactsAndSubmit(strconv.Itoa(i))
	}

	defer cleanupFiles("file*")
}

func TestDownloadedFiles(t *testing.T) {
	downloadTestArtifact("maven-releases", "file20", "1.0.0", "pom")
	downloadTestArtifact("maven-releases", "file20", "1.0.0", "jar")

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
		actual := continuationToken(token)

		if actual != expectedContinuationToken {
			t.Errorf("ContinuationToken incorrect. Expected %s, but was %s.", expectedContinuationToken, actual)
		}
	}
}

func TestContinuationTokenHash(t *testing.T) {
	expected := []string{"35303a6235633862633138616131326331613030356565393061336664653966613733",
		"3130303a6235633862633138616131326331613030356565393061336664653966613733",
		"3135303a6235633862633138616131326331613030356565393061336664653966613733",
		"null"}
	actual := ContinuationTokenRecursion("null")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Maps not equal. Expected %s, but was %s.", expected, actual)
	}
}
