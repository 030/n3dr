package artifacts

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryNamesAndFormatsMap(t *testing.T) {
	testRepositoryNamesAndFormatsJSON := `
[ {
	"name" : "3rdparty-maven",
	"format" : "maven2",
	"type" : "proxy",
	"url" : "https://some-url/repository/3rdparty-maven",
	"attributes" : {
		"proxy" : {
		"remoteUrl" : "https://repo.maven.apache.org/maven2/"
		}
	}
	}, {
	"name" : "3rdparty-npm",
	"format" : "npm",
	"type" : "proxy",
	"url" : "https://some-url/repository/3rdparty-npm",
	"attributes" : {
		"proxy" : {
		"remoteUrl" : "https://registry.npmjs.org/"
		}
	}
	}, {
	"name" : "releases",
	"format" : "maven2",
	"type" : "hosted",
	"url" : "https://some-url/repository/releases",
	"attributes" : { }
	} ]
`
	testRepositoryNamesAndFormatsMap := repositoriesNamesAndFormatsMap{
		"3rdparty-maven": "maven2",
		"3rdparty-npm":   "npm",
		"releases":       "maven2",
	}

	act, err := repositoriesNamesAndFormatsJSONToMap(testRepositoryNamesAndFormatsJSON)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(testRepositoryNamesAndFormatsMap, act) {
		t.Errorf("RepositoryNamesAndFormatMaps not identical. Expected %s, but was %s.", testRepositoryNamesAndFormatsMap, act)
	}
}

func TestRepositoryNamesJSON(t *testing.T) {
	expected := repositoriesNamesAndFormatsMap{
		"maven-central":   "maven2",
		"maven-releases":  "maven2",
		"maven-snapshots": "maven2",
		"nuget-hosted":    "nuget",
		"nuget.org-proxy": "nuget",
	}
	actual, _ := n.repositoriesNamesAndFormatsJSONToMapIncludingRequest()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: '%v'. Actual: '%v'", expected, actual)
	}
}

func TestHappyFlow(t *testing.T) {
	assert.Nil(t, n.CountRepositories())
	assert.Nil(t, n.RepositoryNames())
}
func TestUnhappyFlow(t *testing.T) {
	assert.Nil(t, n.Downloads(".*"))
}
