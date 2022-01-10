package artifacts

import (
	"reflect"
	"testing"
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
