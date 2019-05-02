package main

import (
	"flag"
	"nexus3-cli/cli"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	nexus3URL := flag.String("nexus3URL", "http://localhost:8081", "The Nexus3URL")
	nexus3user := flag.String("nexus3user", "admin", "The Nexus user")
	nexus3pass := flag.String("nexus3pass", "admin123", "The Nexus password")
	nexus3repo := flag.String("nexus3repo", "maven-releases", "The Nexus3 repository")

	flag.Parse()

	n := cli.Nexus3{URL: *nexus3URL, User: *nexus3user, Pass: *nexus3pass, Repository: *nexus3repo}
	log.Info("Nexus3URL: " + n.URL)
	log.Info("Nexus3user: " + n.User)
	log.Info("Nexus3pass: ****")
	log.Info("Nexus3repo: " + n.Repository)

	err := n.StoreArtifactsOnDisk()
	if err != nil {
		log.Fatal(err)
	}
}
