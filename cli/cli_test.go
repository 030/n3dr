package cli

import (
	"os/exec"
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

func submitArtifacts() {
	cmd := exec.Command("bash", "-c", "curl -u admin:admin123 -X POST \"http://localhost:8081/service/rest/v1/components?repository=maven-releases\" -H  \"accept: application/json\" -H  \"Content-Type: multipart/form-data\" -F \"maven2.asset1=@5.1.6.RELEASE.pom\" -F \"maven2.asset1.extension=pom\" -F \"maven2.asset2=@5.1.6.RELEASE.jar\" -F \"maven2.asset2.extension=jar\"")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
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
	submitArtifacts()
	cleanup()

	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
