package artifacts

import (
	"testing"
)

const (
	errMsgTxt       = "Incorrect. Expected: %v. Actual: %v"
	testDirDownload = "/download"
	testDirUpload   = "/testFiles"
)

var n = Nexus3{
	URL:        "http://localhost:9999",
	User:       "admin",
	Pass:       "admin123",
	Repository: "maven-releases",
	APIVersion: "v1",
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
	actualFileErrors := createArtifact(testDirHome+testDirUpload, "file100/file100", "some-content", "ba1f2511fc30423bdbb183fe33f3dd0f")
	expectedErrorFile := "open " + testDirHome + testDirUpload + "/file100/file100: no such file or directory"

	for _, actualFileError := range actualFileErrors {
		if actualFileError.Error() != expectedErrorFile {
			t.Errorf(errMsgTxt, expectedErrorFile, actualFileError)
		}
	}
}

func TestDownloadArtifact(t *testing.T) {
	actualError := n.downloadArtifact(testDirDownload, "http://releasesoftwaremoreoften.com", "")
	expectedError := "URL: 'http://releasesoftwaremoreoften.com' does not seem to contain a Maven artifact"

	if actualError.Error() != expectedError {
		t.Errorf(errMsgTxt, expectedError, actualError)
	}
}

func TestFileExists(t *testing.T) {
	file := "file1/file1/1.0.0/file1-1.0.0.jar"
	result := fileExists(file)
	expectedResult := false

	if result != expectedResult {
		t.Errorf(errMsgTxt, expectedResult, result)
	}
}
