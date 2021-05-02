package nuget

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"

	mp "github.com/030/go-multipart/utils"
	"github.com/030/mij"
	"github.com/030/n3dr/internal/pkg/backup"
	"github.com/cavaliercoder/grab"
	log "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

type mijDockerImage struct {
	*mij.DockerImage
}

const (
	nexus3Version       = "3.30.1"
	nexus3ExternalPort  = 9997
	nexus3BrowserURI    = "/service/rest/repository/browse/"
	nexus3TestDir       = "/tmp/test-n3dr/nuget"
	nugetRepositoryName = "nuget-hosted"
)

var (
	d = mij.DockerImage{
		Name:                     "sonatype/nexus3",
		PortExternal:             nexus3ExternalPort,
		PortInternal:             8081,
		Version:                  nexus3Version,
		ContainerName:            "n3dr-backup-test-nuget",
		LogFile:                  "/nexus-data/log/nexus.log",
		LogFileStringHealthCheck: "Started Sonatype Nexus OSS",
	}
	files = []string{"5.2.6", "6.0.0", "6.0.1", "6.0.2", "6.0.3", "6.0.4", "6.0.5", "6.0.6", "6.0.7", "6.0.8", "6.0.9", "6.0.10"}
	n     = Nexus3{Endpoint: "http://localhost:" + strconv.Itoa(d.PortExternal), Password: "admin123", Username: "admin", BaseDir: nexus3TestDir, Regex: ".*"}
)

func TestMain(m *testing.M) {
	setup(&d)
	code := m.Run()
	shutdown(&d)
	os.Exit(code)
}

func (m *mijDockerImage) initialNexusAdminPassword() (string, error) {
	b, err := exec.Command("bash", "-c", "docker exec -i "+m.ContainerName+" cat /nexus-data/admin.password").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func setup(m *mij.DockerImage) {
	m.Run()

	mdi := &mijDockerImage{m}
	pw, err := mdi.initialNexusAdminPassword()
	if err != nil {
		log.Error(err)
	}
	n.Password = pw

	if err := uploadTestData(pw); err != nil {
		log.Fatal(err)
	}
}

func shutdown(m *mij.DockerImage) {
	m.Stop()
}

func uploadTestData(password string) error {
	for _, file := range files {
		resp, err := grab.Get("/tmp", "https://chocolatey.org/api/v2/package/n3dr/"+file)
		if err != nil {
			return err
		}
		mpu := mp.Upload{URL: "http://localhost:" + strconv.Itoa(nexus3ExternalPort) + "/service/rest/v1/components?repository=" + nugetRepositoryName, Username: n.Username, Password: password}
		if err := mpu.MultipartUpload("nuget.asset=@" + resp.Filename); err != nil {
			return err
		}
	}
	return nil
}

func TestComponentsRepositoryJSON(t *testing.T) {
	n.Repository = "nuget-hosted"
	s, err := n.componentsRepositoryJSON("")
	exp := `n3dr\/5.2.6`
	assert.Nil(t, err)
	assert.Regexp(t, exp, s)
}

func TestRepositoryJSONAssets(t *testing.T) {
	s, err := n.repositoryJSONAssets("")
	assert.Nil(t, err)
	assert.Equal(t, "done", s)

	i := 0
	files, err := ioutil.ReadDir(filepath.Join(n.BaseDir, n.Repository))
	assert.Nil(t, err)
	for _, file := range files {
		if !file.IsDir() {
			i++
		}
	}
	assert.Equal(t, 12, i)
}

func TestContinuationTokenInJSON(t *testing.T) {
	json := `{
		"items": [ {
		} ],
		"continuationToken": "boo"
	}`
	s, err := backup.ContinuationTokenInJSON(json)
	assert.Nil(t, err)
	assert.Equal(t, backup.ContinuationToken("boo"), s)

	json = `{
		"items": [ {
		} ],
		"continuationToken": null
	}`
	s, err = backup.ContinuationTokenInJSON(json)
	assert.Nil(t, err)
	assert.Equal(t, backup.ContinuationToken(""), s)
}

func TestContinuationTokenErrors(t *testing.T) {
	jsons := []string{
		`{
		"items": [ {
		} ],
		"continuationTokenDoesNotExist" : "something"
		}`,
		`{`,
		``,
		`abc`,
	}
	for _, json := range jsons {
		s, err := backup.ContinuationTokenInJSON(json)
		assert.EqualError(t, err, "continuationToken does not exist in json")
		assert.Equal(t, backup.ContinuationToken(""), s)
	}
}
