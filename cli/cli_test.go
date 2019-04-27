package cli

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/030/go-utils"
	log "github.com/sirupsen/logrus"
)

func initializer() {
	cmd := exec.Command("bash", "-c", "docker run -d -p 8081:8081 --name nexus sonatype/nexus3:3.16.1")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}
}

func available() {
	for !utils.URLExists(pingURL) {
		log.Info("Nexus not available.")
	}
}

func submitArtifact(f string) {
	cmd := exec.Command("bash", "-c", "curl -u admin:admin123 -X POST \"http://localhost:8081/service/rest/v1/components?repository=maven-releases\" -H  \"accept: application/json\" -H  \"Content-Type: multipart/form-data\" -F \"maven2.asset1=@file"+f+".pom\" -F \"maven2.asset1.extension=pom\" -F \"maven2.asset2=@file"+f+".jar\" -F \"maven2.asset2.extension=jar\"")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}
}

func createArtifact(f string, content string) {
	file, err := os.Create(f)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(content)
}

func createPOM(f string) {
	createArtifact("file"+f+".pom", "<project>\n<modelVersion>4.0.0</modelVersion>\n<groupId>file"+f+"</groupId>\n<artifactId>file"+f+"</artifactId>\n<version>1.0.0</version>\n</project>")
}

func createJAR(f string) {
	createArtifact("file"+f+".jar", "some-content")
}

func createArtifactsAndSubmit(f string) {
	createPOM(f)
	createJAR(f)
	submitArtifact(f)
}

func postArtifacts() {
	for i := 1; i <= 100; i++ {
		createArtifactsAndSubmit(strconv.Itoa(i))
	}
}

func cleanupFiles() {
	files, err := filepath.Glob("file*")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			log.Fatal(err)
		}
	}
}

func cleanup() {
	cmd := exec.Command("bash", "-c", "docker stop nexus && docker rm nexus")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}
}

func TestSum(t *testing.T) {
	initializer()
	available()
	postArtifacts()
	cleanupFiles()
	cleanup()

	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
