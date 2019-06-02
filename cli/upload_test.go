package cli

import (
	"testing"
)

func TestUploads(t *testing.T) {
	err := n.Upload()
	want := "HTTPStatusCode: '400'; ResponseMessage: 'Repository does not allow updating assets: maven-releases'; ErrorMessage: '<nil>'"
	if err.Error() == want {
		t.Errorf("Error expected. Got '%v'. Want '%v'", err, want)
	}
}
