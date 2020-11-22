package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testDirHome = "/tmp/n3drtest"
)

func TestHashFileMD5(t *testing.T) {
	file := testDirHome + testDirDownload + "/file1/file1/1.0.0/file1-1.0.0.jar"
	_, actualError := hashFileMD5(file)
	expectedError := "open " + testDirHome + testDirDownload + "/file1/file1/1.0.0/file1-1.0.0.jar: no such file or directory"

	if actualError.Error() != expectedError {
		t.Errorf(errMsgTxt, expectedError, actualError)
	}

	file = testFileJar100
	expectedResult := "ad60407c083b4ecc372614b8fcd9f305"
	result, _ := hashFileMD5(file)

	if result != expectedResult {
		t.Errorf(errMsgTxt, expectedResult, result)
	}
}
func TestCreateZip(t *testing.T) {
	assert.Nil(t, n.CreateZip(testDirHome))
	n.ZIP = true
	assert.Nil(t, n.CreateZip(testDirHome))
	n.ZipName = "notAZip"
	assert.EqualError(t, n.CreateZip(testDirHome), "format unrecognized by filename: notAZip")
	n.ZIP = false
}

func TestRequest(t *testing.T) {
	n.User = ""
	n.validate()
	n.Pass = ""
	n.validate()
	_, _, err := n.request("incorrectUrl")
	assert.EqualError(t, err, "Get \"incorrectUrl\": GET incorrectUrl giving up after 1 attempt(s): Get \"incorrectUrl\": unsupported protocol scheme \"\"")
	n.User = "admin"
	n.Pass = "admin123"
}
