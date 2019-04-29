package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/thedevsaddam/gojsonq"
	"nexus3-cli/cli"
	"strconv"
)

func downloadURLs() {
	continuationTokenMap := cli.ContinuationTokenRecursion("null")
	for tokenNumber, token := range continuationTokenMap {
		tokenNumberString := strconv.Itoa(tokenNumber)
		log.Info("ContinuationToken: " + token + "; ContinuationTokenNumber: " + tokenNumberString)
		bytes, err := cli.DownloadURL(token)
		if err != nil {
			log.Fatal(err)
		}
		json := string(bytes)

		jq := gojsonq.New().JSONString(json)
		downloadURLsInterface := jq.From("items").Pluck("downloadUrl")

		downloadURLsInterfaceArray := downloadURLsInterface.([]interface{})

		for i, downloadURL := range downloadURLsInterfaceArray {
			log.Printf("OK: message %d => %s\n", i, downloadURL)
			cli.DownloadArtifact(fmt.Sprint(downloadURL))
		}
	}
}

func main() {
	downloadURLs()
}
