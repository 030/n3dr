//go:build integration
// +build integration

package upload

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	pool.MaxWait = 5 * time.Minute

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// mkdir -p /tmp/nexus-data
	// sudo chown 200:200 -R /tmp/nexus-data
	// sudo chmod 0700 /tmp/nexus-data/change-admin-password.sh
	//
	// #!/bin/bash
	// pw=$(cat /nexus-data/admin.password)
	// curl -u admin:${pw}  -X 'PUT' -H 'Content-Type: text/plain' 'http://localhost:8081/service/rest/v1/security/users/admin/change-password' -d '123456789' -v
	options := &dockertest.RunOptions{
		Repository: "sonatype/nexus3",
		Tag:        "3.58.1",
		Mounts:     []string{"/tmp/nexus-data:/nexus-data"},
		// Cmd:        []string{"cat", "/nexus-data/admin.password"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"8081/tcp": {{HostPort: "9999"}},
		},
		// Env: []string{"MINIO_ACCESS_KEY=MYACCESSKEY", "MINIO_SECRET_KEY=MYSECRETKEY"},
	}

	resource, err := pool.RunWithOptions(options)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	endpoint := fmt.Sprintf("localhost:%s", resource.GetPort("8081/tcp"))

	if err := pool.Retry(func() error {
		url := fmt.Sprintf("http://%s/", endpoint)
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("status code not OK")
		}
		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// exit_code, err := resource.Exec([]string{"/nexus-data/change-admin-password.sh"}, dockertest.ExecOptions{})
	// fmt.Println(exit_code)
	// if exit_code != 0 {
	// 	log.Fatal(exit_code)
	// }

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		exit_code, err := resource.Exec([]string{"/nexus-data/change-admin-password.sh"}, dockertest.ExecOptions{})
		if err != nil {
			return err
		}
		if exit_code != 0 {
			return fmt.Errorf("Exit code was not: 0, but: '%d'", exit_code)
		}
		return nil
	}); err != nil {
		log.Fatalf("Could not change admin pass: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestSomething(t *testing.T) {
	// db.Query()
}

// func TestMain(m *testing.M) {
// 	containers := []mij.DockerImage{n3drtest.Image(10002)}
// 	if err := n3drtest.Setup(containers); err != nil {
// 		panic(err)
// 	}

// 	code := m.Run()
// 	if err := n3drtest.Shutdown(containers); err != nil {
// 		panic(err)
// 	}

// 	os.Exit(code)
// }

// func TestUploadSnapshots(t *testing.T) {
// 	n := connection.Nexus3{FQDN: "localhost:10002", Pass: "testi", User: "admin", DownloadDirName: "../../../../../test/testdata/upload/snapshots"}

// 	r := repository.Repository{Nexus3: n}
// 	err := r.CreateMavenHosted("maven-snapshots", true)
// 	assert.NoError(t, err)

// 	u := Nexus3{Nexus3: &n}
// 	err = u.Upload()
// 	assert.NoError(t, err)
// }

// func TestUploadSnapshotsSkipErrors(t *testing.T) {
// 	n := connection.Nexus3{FQDN: "localhost:10002", Pass: "testi", User: "admin", DownloadDirName: "../../../../../test/testdata/upload/snapshots-fail"}
// 	n.SkipErrors = true

// 	r := repository.Repository{Nexus3: n}
// 	err := r.CreateMavenHosted("maven-snapshots-fail", true)
// 	assert.NoError(t, err)

// 	u := Nexus3{Nexus3: &n}
// 	err = u.Upload()
// 	assert.NoError(t, err)
// }
