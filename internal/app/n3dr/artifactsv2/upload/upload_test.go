package upload

import (
	"testing"

	"github.com/030/mij"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/n3drtest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	containers := []mij.DockerImage{n3drtest.Image(10002)}
	if err := n3drtest.Setup(containers); err != nil {
		panic(err)
	}

	// code := m.Run()
	// if err := n3drtest.Shutdown(containers); err != nil {
	// 	panic(err)
	// }

	// os.Exit(code)
}

func TestUpload(t *testing.T) {
	n := connection.Nexus3{FQDN: "localhost:10002", Pass: "testi", User: "admin", DownloadDirName: "../../../../../test/testdata/upload"}
	u := Nexus3{Nexus3: &n}
	err := u.Upload()
	assert.NoError(t, err)
}
