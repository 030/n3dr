package cli

import (
	"reflect"
	"testing"
)

// func TestUploads(t *testing.T) {
// 	err := n.Upload()
// 	want := "HTTPStatusCode: '400'; ResponseMessage: 'Repository does not allow updating assets: maven-releases'; ErrorMessage: '<nil>'"
// 	if err.Error() != want {
// 		t.Errorf("Error expected. Got '%v'. Want '%v'", err, want)
// 	}
// }

func TestDetectFoldersWithPOM(t *testing.T) {
	err := n.detectFoldersWithPOM("../maven-releases")
	if err != nil {
		t.Errorf("No error expected, but was %v", err)
	}

	err = n.detectFoldersWithPOM("maven-releases-FOLDER_DOES_NOT_EXIST")
	want := "lstat maven-releases-FOLDER_DOES_NOT_EXIST: no such file or directory"
	if err.Error() != want {
		t.Errorf("Error expected. Got '%v'. Want '%v'", err, want)
	}
}

func TestDirectoriesContainingPOM(t *testing.T) {
	got := pomDirectories()
	want := []string{
		"../maven-releases/com/palantir/docker/compose/docker-compose-rule-core/0.32.0",
		"../maven-releases/com/palantir/docker/compose/docker-compose-rule-core/0.35.0",
		"../maven-releases/file1/file1/1.0.0",
		"../maven-releases/file2/file2/1.0.0",
		"../maven-releases/file22/file22/1.0.0",
		"../maven-releases/file3/file3/1.0.0",
		"../maven-releases/file42/file42/1.0.0",
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Slice not identical. Expected %v, but was %v.", want, got)
	}
}

func TestMultipartContent(t *testing.T) {
	artifacts := map[string]string{
		"../maven-releases/com/palantir/docker/compose/docker-compose-rule-core/0.32.0/docker-compose-rule-core-0.32.0.pom": "maven2.asset1=@../maven-releases/com/palantir/docker/compose/docker-compose-rule-core/0.32.0/docker-compose-rule-core-0.32.0.pom,maven2.asset1.extension=pom",
		"../maven-releases/file22/file22/1.0.0/file22-1.0.0.war":                                                            "maven2.asset2=@../maven-releases/file22/file22/1.0.0/file22-1.0.0.war,maven2.asset2.extension=war",
		"../maven-releases/file22/file22/1.0.0/file22-1.0.0.zip":                                                            "maven2.asset3=@../maven-releases/file22/file22/1.0.0/file22-1.0.0.zip,maven2.asset3.extension=zip",
		"../maven-releases/com/palantir/docker/compose/docker-compose-rule-core/0.35.0/docker-compose-rule-core-0.35.0.jar": "maven2.asset4=@../maven-releases/com/palantir/docker/compose/docker-compose-rule-core/0.35.0/docker-compose-rule-core-0.35.0.jar,maven2.asset4.extension=jar",
		// "../maven-releases/file22/file22/1.0.0/file22-1.0.0-sources.jar":                                                    "maven2.asset2=@../maven-releases/file22/file22/1.0.0/file22-1.0.0-sources.jar,maven2.asset2.extension=jar",
		// "../maven-releases/file22/file22/1.0.0/file22-1.0.0-test-resources.jar":                                             "maven2.asset2=@../maven-releases/file22/file22/1.0.0/file22-1.0.0-test-resources.jar,maven2.asset2.extension=jar",
		"not/an/artifact": "",
	}

	for k, want := range artifacts {
		got, _ := multipartContent(k)
		// if err.Error() != want {
		// 	t.Errorf("Error expected. Got '%v'. Want '%v'", err, want)
		// }
		if got != want {
			t.Errorf("Mismatch. Got '%v', Want '%v'", got, want)
		}
	}
}

func TestOtherJAR(t *testing.T) {
	artifacts := map[string]string{
		"file22-1.0.0-test-resources.jar": "test-resources",
		"file22-1.0.0-sources.jar":        "sources",
		"file22-1.0.0-bundledPdfs.jar":    "bundledPdfs",
		"":                                "",
		"hello":                           "",
	}

	for v, k := range artifacts {
		got := jarClassifier(v)
		want := k

		if got != want {
			t.Errorf("Mismatch. Got '%v', Want '%v'", got, want)
		}
	}
}
