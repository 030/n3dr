package cli

import (
	"io/ioutil"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestUploads(t *testing.T) {
	d := "maven-releases"

	// if upload repository does not exist
	err := n.detectFoldersWithPOM(d)
	want := "lstat " + d + ": no such file or directory"
	if err.Error() != want {
		t.Errorf("Error expected. Got '%v'. Want '%v'", err, want)
	}

	// if upload repository exists, without .pom files
	err = os.Mkdir(d, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	err = n.detectFoldersWithPOM(d)
	want = "no folders with .pom files detected. Please check whether the '" + d + "' directory contains .pom files"
	if err.Error() != want {
		t.Errorf("Error expected. Got '%v'. Want '%v'", err, want)
	}

	// happy
	ioutil.WriteFile(d+"/hello.pom", nil, os.ModePerm)
	err = n.detectFoldersWithPOM(d)
	if err != nil {
		t.Errorf("No error expected. Got '%v'. Want 'nil", err)
	}

	createPOM(d, "hello", "1")

	err = n.Upload()
	want = "HTTPStatusCode: '400'; ResponseMessage: 'Repository does not allow updating assets: maven-releases'; ErrorMessage: '<nil>'"
	if err.Error() != want {
		t.Errorf("Error expected. Got '%v'. Want '%v'", err, want)
	}

	// cleanup
	err = os.RemoveAll(d)
	if err != nil {
		log.Fatal(err)
	}
}
