package cli

import (
	"reflect"
	"testing"

	"github.com/030/go-utils"
)

const (
	errMsg    = "Not equal. Expected: %d. Actual: %d."
	errMsgTxt = "Error incorrect. Expected: %v. Actual: %v"
)

func TestContinuationTokenHash(t *testing.T) {
	actual, _ := n.continuationTokenRecursion("null")
	actualSize := len(actual)
	expectedSize := 3
	if expectedSize != actualSize {
		t.Errorf(errMsg, expectedSize, actualSize)
	}

	tokenSlice := []string{
		"foo",
		"boo",
		"",
		"----",
		"123",
		"11111111111111111111111111111111111",
	}
	for _, token := range tokenSlice {
		_, actualError := n.continuationTokenRecursion(token)

		expectedError := tokenErrMsg + token
		if actualError.Error() != expectedError {
			t.Errorf(errMsgTxt, expectedError, actualError)
		}
	}
}

func TestDownloadURLs(t *testing.T) {
	url, _ := n.downloadURLs()
	actual := len(url)
	expected := 27 // 3files*9
	if expected != actual {
		t.Errorf(errMsg, expected, actual)
	}
}

func TestStoreArtifactsOnDisk(t *testing.T) {
	n.StoreArtifactsOnDisk()

	actual, _ := allFiles(downloadDir)

	actualFileNumber := len(actual)
	expected := 28 // +1 due to .gitkeep
	if expected != actualFileNumber {
		t.Errorf(errMsg, expected, actualFileNumber)
	}

	expectedDownloads := []string{
		"download/.gitkeep",
		"download/maven-releases/file1/file1/1.0.0/file1-1.0.0.jar",
		"download/maven-releases/file1/file1/1.0.0/file1-1.0.0.jar.md5",
		"download/maven-releases/file1/file1/1.0.0/file1-1.0.0.jar.sha1",
		"download/maven-releases/file1/file1/1.0.0/file1-1.0.0.pom",
		"download/maven-releases/file1/file1/1.0.0/file1-1.0.0.pom.md5",
		"download/maven-releases/file1/file1/1.0.0/file1-1.0.0.pom.sha1",
		"download/maven-releases/file1/file1/maven-metadata.xml",
		"download/maven-releases/file1/file1/maven-metadata.xml.md5",
		"download/maven-releases/file1/file1/maven-metadata.xml.sha1",
		"download/maven-releases/file2/file2/1.0.0/file2-1.0.0.jar",
		"download/maven-releases/file2/file2/1.0.0/file2-1.0.0.jar.md5",
		"download/maven-releases/file2/file2/1.0.0/file2-1.0.0.jar.sha1",
		"download/maven-releases/file2/file2/1.0.0/file2-1.0.0.pom",
		"download/maven-releases/file2/file2/1.0.0/file2-1.0.0.pom.md5",
		"download/maven-releases/file2/file2/1.0.0/file2-1.0.0.pom.sha1",
		"download/maven-releases/file2/file2/maven-metadata.xml",
		"download/maven-releases/file2/file2/maven-metadata.xml.md5",
		"download/maven-releases/file2/file2/maven-metadata.xml.sha1",
		"download/maven-releases/file3/file3/1.0.0/file3-1.0.0.jar",
		"download/maven-releases/file3/file3/1.0.0/file3-1.0.0.jar.md5",
		"download/maven-releases/file3/file3/1.0.0/file3-1.0.0.jar.sha1",
		"download/maven-releases/file3/file3/1.0.0/file3-1.0.0.pom",
		"download/maven-releases/file3/file3/1.0.0/file3-1.0.0.pom.md5",
		"download/maven-releases/file3/file3/1.0.0/file3-1.0.0.pom.sha1",
		"download/maven-releases/file3/file3/maven-metadata.xml",
		"download/maven-releases/file3/file3/maven-metadata.xml.md5",
		"download/maven-releases/file3/file3/maven-metadata.xml.sha1",
	}
	for _, f := range expectedDownloads {
		if !utils.FileExists(f) {
			t.Errorf("File %s should exist, but does not.", f)
		}
	}

	if !reflect.DeepEqual(expectedDownloads, actual) {
		t.Errorf("Slice not identical. Expected %s, but was %s.", expectedDownloads, actual)
	}
}
func TestDownloadURL(t *testing.T) {
	_, actualError := n.downloadURL("some-token")
	expectedError := "HTTP response not 200. Does the URL: http://localhost:9999/service/rest/v1/assets?repository=maven-releases&continuationToken=some-token exist?"

	if actualError.Error() != expectedError {
		t.Errorf(errMsgTxt, expectedError, actualError)
	}
}

func TestArtifactName(t *testing.T) {
	actualDir, actualFile, _ := n.artifactName("http://localhost:9999/repository/maven-releases/file1/file1/1.0.0/file1-1.0.0.jar")
	expectedDir := "file1/file1/1.0.0"
	expectedFile := "file1-1.0.0.jar"

	if expectedDir != actualDir || expectedFile != actualFile {
		t.Errorf("Dir and file incorrect. Expected: %v & %v. Actual: %v & %v", expectedDir, expectedFile, actualDir, actualFile)
	}

	_, _, actualError := n.artifactName("some-url")
	expectedError := "some-url is not an URL"

	if actualError.Error() != expectedError {
		t.Errorf(errMsgTxt, expectedError, actualError)
	}
}

func TestArtifactNameContainingRepositoryName(t *testing.T) {
	actualDir, actualFile, _ := n.artifactName("http://localhost:9999/repository/maven-releases/com/maven-releases/tools/1.0.0/tools-1.0.0.jar")
	expectedDir := "com/maven-releases/tools/1.0.0"
	expectedFile := "tools-1.0.0.jar"

	if expectedDir != actualDir || expectedFile != actualFile {
		t.Errorf("Dir and file incorrect. Expected: %v & %v. Actual: %v & %v", expectedDir, expectedFile, actualDir, actualFile)
	}

	_, _, actualError := n.artifactName("some-url")
	expectedError := "some-url is not an URL"

	if actualError.Error() != expectedError {
		t.Errorf(errMsgTxt, expectedError, actualError)
	}
}

func TestCreateArtifact(t *testing.T) {
	actualErrorFile := createArtifact("testFiles", "file100/file100", "some-content")
	expectedErrorFile := "open testFiles/file100/file100: no such file or directory"

	if actualErrorFile.Error() != expectedErrorFile {
		t.Errorf(errMsgTxt, expectedErrorFile, actualErrorFile)
	}
}

func TestDownloadArtifact(t *testing.T) {
	actualError := n.downloadArtifact("http://releasesoftwaremoreoften.com")
	expectedError := "URL: 'http://releasesoftwaremoreoften.com' does not seem to contain an artifactName"

	if actualError.Error() != expectedError {
		t.Errorf("Error incorrect. Expected: %v. Actual: %v", expectedError, actualError)
	}
}
