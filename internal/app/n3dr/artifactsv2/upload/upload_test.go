package upload

import (
	"os"
	"testing"

	"github.com/030/mij"
	"github.com/030/n3dr/internal/app/n3dr/config/repository"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/n3drtest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	containers := []mij.DockerImage{n3drtest.Image(10002)}
	if err := n3drtest.Setup(containers); err != nil {
		panic(err)
	}

	code := m.Run()
	if err := n3drtest.Shutdown(containers); err != nil {
		panic(err)
	}

	os.Exit(code)
}

func TestUploadSnapshots(t *testing.T) {
	n := connection.Nexus3{FQDN: "localhost:10002", Pass: "testi", User: "admin", DownloadDirName: "../../../../../test/testdata/upload/snapshots"}

	r := repository.Repository{Nexus3: n}
	err := r.CreateMavenHosted("maven-snapshots", true)
	assert.NoError(t, err)

	u := Nexus3{Nexus3: &n}
	err = u.Upload()
	assert.NoError(t, err)
}

func TestUploadSnapshotsSkipErrors(t *testing.T) {
	n := connection.Nexus3{FQDN: "localhost:10002", Pass: "testi", User: "admin", DownloadDirName: "../../../../../test/testdata/upload/snapshots-fail"}
	n.SkipErrors = true

	r := repository.Repository{Nexus3: n}
	err := r.CreateMavenHosted("maven-snapshots-fail", true)
	assert.NoError(t, err)

	u := Nexus3{Nexus3: &n}
	err = u.Upload()
	assert.NoError(t, err)
}
