package main

import (
	"flag"
	"nexus3-cli/cli"

	log "github.com/sirupsen/logrus"
)

func main() {
	nexus3URL := flag.String("nexus3URL", "http://localhost:8081", "The Nexus3URL")
	flag.Parse()

	n := cli.Nexus3{URL: *nexus3URL}
	log.Info("Nexus3URL: " + n.URL)

	n.StoreArtifactsOnDisk()
}
