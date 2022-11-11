package n3drtest

import (
	"os/exec"
	"strconv"
	"sync"

	"github.com/030/mij"
	"github.com/030/n3dr/internal/app/n3dr/config/user"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
	log "github.com/sirupsen/logrus"
)

func Setup(containers []mij.DockerImage) error {
	log.SetLevel(log.DebugLevel)
	var wg sync.WaitGroup
	for _, container := range containers {
		wg.Add(1)
		go func(m mij.DockerImage) {
			defer wg.Done()
			if err := m.Run(); err != nil {
				panic(err)
			}

			if err := pass(m); err != nil {
				panic(err)
			}
		}(container)
	}
	wg.Wait()
	return nil
}

func Shutdown(containers []mij.DockerImage) error {
	var wg sync.WaitGroup
	for _, container := range containers {
		wg.Add(1)
		go func(m mij.DockerImage) {
			defer wg.Done()
			if err := m.Stop(); err != nil {
				panic(err)
			}
		}(container)
	}
	wg.Wait()
	return nil
}

func Image(port int) mij.DockerImage {
	return mij.DockerImage{
		Name:                     "sonatype/nexus3",
		PortExternal:             port,
		PortInternal:             8081,
		Version:                  "3.43.0",
		ContainerName:            "nexus" + strconv.Itoa(port),
		LogFile:                  "/nexus-data/log/nexus.log",
		LogFileStringHealthCheck: "Started Sonatype Nexus OSS",
	}
}

func pass(m mij.DockerImage) error {
	b, err := exec.Command("bash", "-c", "docker exec -i nexus"+strconv.Itoa(m.PortExternal)+" cat /nexus-data/admin.password").CombinedOutput() // #nosec G204
	if err != nil {
		return err
	}
	acu := models.APICreateUser{
		EmailAddress: "admin@example.org",
		FirstName:    "admin",
		LastName:     "admin",
		Password:     "testi",
		UserID:       "admin",
	}
	n := connection.Nexus3{FQDN: "localhost:" + strconv.Itoa(m.PortExternal), Pass: string(b), User: "admin"}
	u := user.User{APICreateUser: acu, Nexus3: n}
	if err := u.ChangePass(); err != nil {
		return err
	}
	return nil
}
