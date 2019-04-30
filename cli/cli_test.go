package cli

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

var n = Nexus3{
	URL:        "http://localhost:9999",
	User:       "admin",
	Pass:       "admin123",
	Repository: "maven-releases",
}

// See https://stackoverflow.com/a/34102842/2777965
func TestMain(m *testing.M) {
	setup()
	m.Run()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestContinuationTokenHash(t *testing.T) {
	actual := n.continuationTokenRecursion("null")

	actualSize := len(actual)
	expectedSize := 3

	if expectedSize != actualSize {
		t.Errorf("Not equal. Expected: %d. Actual: %d.", expectedSize, actualSize)
	}
}

func TestDownloadURLs(t *testing.T) {
	log.Info(n.downloadURLs())
	actual := len(n.downloadURLs())
	expected := 27 // 3files*9
	if expected != actual {
		t.Errorf("Not equal. Expected: %d. Actual: %d.", expected, actual)
	}
}

func TestStoreArtifactsOnDisk(t *testing.T) {
	n.StoreArtifactsOnDisk()

	files, _ := allFiles("download")

	actual := len(files)
	expected := 28 // +1 due to .gitkeep

	if expected != actual {
		t.Errorf("Not equal. Expected: %d. Actual: %d.", expected, actual)
	}
}
