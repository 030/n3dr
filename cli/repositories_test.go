package cli

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryNamesJSON(t *testing.T) {
	var expected interface{} = []interface{}{"maven-central", "maven-releases", "maven-snapshots", "nuget-hosted", "nuget.org-proxy"}
	actual, _ := n.repositoriesSlice()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: '%v'. Actual: '%v'", expected, actual)
	}
}

func TestHappyFlow(t *testing.T) {
	assert.Nil(t, n.CountRepositories())
	assert.Nil(t, n.RepositoryNames())
	assert.Nil(t, n.Downloads(".*"))
}
func TestUnhappyFlow(t *testing.T) {
	n.Pass = "incorrectPass"
	assert.EqualError(t, n.CountRepositories(), testNexusAuthError)
	assert.EqualError(t, n.RepositoryNames(), testNexusAuthError)
	assert.EqualError(t, n.Downloads(".*"), testNexusAuthError)
	n.Pass = "admin123"
}
