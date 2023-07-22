//go:build unit
// +build unit

package repository

import (
	"os"
	"testing"

	"github.com/030/mij"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/n3drtest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	containers := []mij.DockerImage{n3drtest.Image(10001)}
	if err := n3drtest.Setup(containers); err != nil {
		panic(err)
	}

	code := m.Run()
	if err := n3drtest.Shutdown(containers); err != nil {
		panic(err)
	}

	os.Exit(code)
}

func TestCreateDockerHosted(t *testing.T) {
	n := connection.Nexus3{FQDN: "localhost:10001", Pass: "testi", User: "admin"}
	r := Repository{Nexus3: n}
	err := r.CreateDockerHosted(false, 2345, "some-docker-repo")
	assert.NoError(t, err)
}

func TestCreateDockerHostedFail(t *testing.T) {
	n := connection.Nexus3{FQDN: "localhost:10001", Pass: "testi", User: "admin"}
	r := Repository{Nexus3: n}
	err := r.CreateDockerHosted(false, 2345, "")
	assert.EqualError(t, err, "repo name should not be empty")
}

func TestCreateMavenHosted(t *testing.T) {
	n := connection.Nexus3{FQDN: "localhost:10001", Pass: "testi", User: "admin"}
	r := Repository{Nexus3: n}
	err := r.CreateMavenHosted("some-maven-repo", false)
	assert.NoError(t, err)
}

func TestCreateNpmHosted(t *testing.T) {
	n := connection.Nexus3{FQDN: "localhost:10001", Pass: "testi", User: "admin"}
	r := Repository{Nexus3: n}
	err := r.CreateNpmHosted("some-npm-repo", false)
	assert.NoError(t, err)
}

func TestCreateRawHosted(t *testing.T) {
	n := connection.Nexus3{FQDN: "localhost:10001", Pass: "testi", User: "admin"}
	r := Repository{Nexus3: n}
	err := r.CreateRawHosted("some-raw-repo")
	assert.NoError(t, err)
}

func TestCreateYumHosted(t *testing.T) {
	n := connection.Nexus3{FQDN: "localhost:100021", Pass: "testi", User: "admin"}
	r := Repository{Nexus3: n}
	err := r.CreateYumHosted("some-yum-repo")
	assert.NoError(t, err)
}
