package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOciBackup(t *testing.T) {
	assert.EqualError(t, ociBackup("", ""), "can not create client, bad configuration: did not find a proper configuration for tenancy")
}

func TestFindObject(t *testing.T) {
	_, err := findObject("", "", "")
	assert.EqualError(t, err, "can not create client, bad configuration: did not find a proper configuration for tenancy")
}
