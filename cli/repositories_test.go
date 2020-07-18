package cli

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryNamesJSON(t *testing.T) {
	var expected interface{} = []interface{}{"maven-central", "maven-public", "maven-releases", "maven-snapshots", "nuget-group", "nuget-hosted", "nuget.org-proxy"}
	actual, _ := n.repositoriesSlice()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: '%v'. Actual: '%v'", expected, actual)
	}
}

func TestCountRepositories(t *testing.T) {
	assert.Nil(t, n.CountRepositories())
	assert.EqualError(t, nErrAuth.CountRepositories(), testNexusAuthError)
}

func TestRepositoryNames(t *testing.T) {
	assert.Nil(t, n.RepositoryNames())
	assert.EqualError(t, nErrAuth.RepositoryNames(), testNexusAuthError)
}

func TestDownloads(t *testing.T) {
	assert.Nil(t, n.Downloads(".*"))
	assert.EqualError(t, nErrAuth.Downloads(".*"), testNexusAuthError)
}
