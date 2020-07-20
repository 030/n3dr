package cli

import (
	"reflect"
	"testing"

	"github.com/030/go-utils"
	log "github.com/sirupsen/logrus"
)

const (
	errMsg         = "Not equal. Expected: %d. Actual: %d."
	errMsgTxt      = "Incorrect. Expected: %v. Actual: %v"
	testFileJar100 = testDirHome + testDirDownload + "/maven-releases/file1/file1/1.0.0/file1-1.0.0.jar"
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
	if err := n.StoreArtifactsOnDisk(testDirHome+testDirDownload, ".*"); err != nil {
		log.Fatal(err)
	}

	actual, _ := allFiles(testDirHome + testDirDownload)

	actualFileNumber := len(actual)
	expected := 9
	if expected != actualFileNumber {
		t.Errorf(errMsg, expected, actualFileNumber)
	}

	expectedDownloads := []string{
		testFileJar100,
		"/tmp/n3drtest/download/maven-releases/file1/file1/1.0.0/file1-1.0.0.pom",
		"/tmp/n3drtest/download/maven-releases/file1/file1/maven-metadata.xml",
		"/tmp/n3drtest/download/maven-releases/file2/file2/1.0.0/file2-1.0.0.jar",
		"/tmp/n3drtest/download/maven-releases/file2/file2/1.0.0/file2-1.0.0.pom",
		"/tmp/n3drtest/download/maven-releases/file2/file2/maven-metadata.xml",
		"/tmp/n3drtest/download/maven-releases/file3/file3/1.0.0/file3-1.0.0.jar",
		"/tmp/n3drtest/download/maven-releases/file3/file3/1.0.0/file3-1.0.0.pom",
		"/tmp/n3drtest/download/maven-releases/file3/file3/maven-metadata.xml",
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
	expectedError := "ResponseCode: '406' and Message '406 Not Acceptable' for URL: http://localhost:9999/service/rest/v1/assets?repository=maven-releases&continuationToken=some-token"

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
	actualErrorFile := createArtifact(testDirHome+testDirUpload, "file100/file100", "some-content", "ba1f2511fc30423bdbb183fe33f3dd0f")
	expectedErrorFile := "open " + testDirHome + testDirUpload + "/file100/file100: no such file or directory"

	if actualErrorFile.Error() != expectedErrorFile {
		t.Errorf(errMsgTxt, expectedErrorFile, actualErrorFile)
	}
}

func TestDownloadArtifact(t *testing.T) {
	url := make(map[string]interface{})
	url["downloadUrl"] = "http://releasesoftwaremoreoften.com"
	actualError := n.downloadArtifact(testDirDownload, url)
	expectedError := "URL: 'http://releasesoftwaremoreoften.com' does not seem to contain an artifactName"

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

	file = testFileJar100
	result = fileExists(file)
	expectedResult = true

	if result != expectedResult {
		t.Errorf(errMsgTxt, expectedResult, result)
	}

}
