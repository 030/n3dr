package cli

import (
	"encoding/json"
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
)

var testRepositories = `[ {
  "name" : "maven-snapshots",
  "format" : "maven2",
  "type" : "hosted",
  "url" : "http://localhost:9999/repository/maven-snapshots"
}, {
  "name" : "maven-central",
  "format" : "maven2",
  "type" : "proxy",
  "url" : "http://localhost:9999/repository/maven-central"
}, {
  "name" : "nuget-group",
  "format" : "nuget",
  "type" : "group",
  "url" : "http://localhost:9999/repository/nuget-group"
}, {
  "name" : "nuget.org-proxy",
  "format" : "nuget",
  "type" : "proxy",
  "url" : "http://localhost:9999/repository/nuget.org-proxy"
}, {
  "name" : "maven-releases",
  "format" : "maven2",
  "type" : "hosted",
  "url" : "http://localhost:9999/repository/maven-releases"
}, {
  "name" : "nuget-hosted",
  "format" : "nuget",
  "type" : "hosted",
  "url" : "http://localhost:9999/repository/nuget-hosted"
}, {
  "name" : "maven-public",
  "format" : "maven2",
  "type" : "group",
  "url" : "http://localhost:9999/repository/maven-public"
} ]`

var bla = `{"name":{"first":"Tom","last":"Hanks"},"age":61}`

func marshal(j string) []byte {
	m, err := json.Marshal(j)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func TestRepositories(t *testing.T) {
	expected := marshal(testRepositories)
	actual := marshal(repositories())

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: '%s'. Actual: '%s'", expected, actual)
	}
}

func TestRepositoryNames(t *testing.T) {
	var expected interface{} = []interface{}{"maven-central", "maven-public", "maven-releases", "maven-snapshots", "nuget-group", "nuget-hosted", "nuget.org-proxy"}
	actual := repositoryNames(testRepositories)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: '%v'. Actual: '%v'", expected, actual)
	}
}
