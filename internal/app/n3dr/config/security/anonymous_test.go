//go:build !integration

package security

import (
	"os"
	"testing"

	"github.com/030/mij"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/n3drtest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	containers := []mij.DockerImage{n3drtest.Image(10000)}
	if err := n3drtest.Setup(containers); err != nil {
		panic(err)
	}

	code := m.Run()
	if err := n3drtest.Shutdown(containers); err != nil {
		panic(err)
	}

	os.Exit(code)
}

func TestAnonymousEnable(t *testing.T) {
	https := false
	n := connection.Nexus3{FQDN: "localhost:10000", Pass: "testi", User: "admin", HTTPS: &https}
	s := Security{Nexus3: n}
	err := s.Anonymous(true)
	assert.NoError(t, err)
}

func TestAnonymousDisable(t *testing.T) {
	https := false
	n := connection.Nexus3{FQDN: "localhost:10000", Pass: "testi", User: "admin", HTTPS: &https}
	s := Security{Nexus3: n}
	err := s.Anonymous(false)
	assert.NoError(t, err)
}

func TestAnonymousFail(t *testing.T) {
	https := false
	n := connection.Nexus3{FQDN: "localhost:10000", Pass: "testi-incorrect", User: "admin", HTTPS: &https}
	s := Security{Nexus3: n}
	err := s.Anonymous(false)
	assert.EqualError(t, err, "could not change anonymous access mode: 'response status code does not match any response statuses defined for this endpoint in the swagger spec (status 401): {}'")
}
