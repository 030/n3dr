package cli

import (
	"reflect"
	"testing"
)

func TestRepositoryNamesJSON(t *testing.T) {
	var expected interface{} = []interface{}{"maven-central", "maven-public", "maven-releases", "maven-snapshots", "nuget-group", "nuget-hosted", "nuget.org-proxy"}
	actual, _ := n.repositoriesSlice()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: '%v'. Actual: '%v'", expected, actual)
	}
}
