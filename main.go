package main

import (
	"flag"
	"nexus3-cli/cli"

	log "github.com/sirupsen/logrus"
)

func main() {
	nexus3URL := flag.String("nexus3URL", "http://localhost:8081", "The Nexus3URL")
	nexus3user := flag.String("nexus3user", "admin", "The Nexus user")
	nexus3pass := flag.String("nexus3pass", "admin123", "The Nexus password")

	flag.Parse()

	n := cli.Nexus3{URL: *nexus3URL, User: *nexus3user, Pass: *nexus3pass}
	log.Info("Nexus3URL: " + n.URL)
	log.Info("Nexus3user: " + n.User)
	log.Info("Nexus3pass: ****")

	n.StoreArtifactsOnDisk()
}
