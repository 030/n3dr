package main

import (
	"flag"
	"n3dr/cli"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	n3drURL := flag.String("n3drURL", "http://localhost:8081", "The Nexus3URL")
	n3drUser := flag.String("n3drUser", "admin", "The Nexus user")
	n3drPass := flag.String("n3drPass", "admin123", "The Nexus password")
	n3drRepo := flag.String("n3drRepo", "maven-releases", "The Nexus3 repository")

	flag.Parse()

	n := cli.Nexus3{URL: *n3drURL, User: *n3drUser, Pass: *n3drPass, Repository: *n3drRepo}
	log.Info("n3drURL: " + n.URL)
	log.Info("n3drUser: " + n.User)
	log.Info("n3drPass: ****")
	log.Info("n3drRepo: " + n.Repository)

	err := n.StoreArtifactsOnDisk()
	if err != nil {
		log.Fatal(err)
	}
}
