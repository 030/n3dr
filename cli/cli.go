package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/svenfuchs/jq"
	"github.com/thedevsaddam/gojsonq"
)

const (
	pingURL  = "http://localhost:8081/service/metrics/ping"
	assetURL = "http://localhost:8081/service/rest/v1/search/assets?repository=maven-releases"
)

func downloadURL(token string) ([]byte, error) {
	url := assetURL
	if !(token == "null") {
		url = assetURL + "&continuationToken=" + token
	}
	// log.Info("DownloadURL: ", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Info(resp.StatusCode)
		return nil, errors.New("HTTP response not 200")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func continuationToken(token string) string {
	bodyBytes, err := downloadURL(token)
	if err != nil {
		log.Fatal(err)
	}

	op, err := jq.Parse(".continuationToken")
	if err != nil {
		//
	}

	// data := []byte(`{"items":[{"hello":"world"},{"hello","bye"}],"hi":"bye"}`) // sample input
	value, err := op.Apply(bodyBytes) // value == '"world"'
	if err != nil {
		//
	}
	var tokenWithoutQuotes string
	tokenWithoutQuotes = strings.Trim(string(value), "\"")

	return tokenWithoutQuotes
}

func continuationTokenRecursion(s string) []string {
	token := continuationToken(s)
	if token == "null" {
		return []string{token}
	}
	return append(continuationTokenRecursion(token), token)
}

func createArtifact(f string, content string) {
	file, err := os.Create(f)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(content)
	defer file.Close()
}

func artifactName(url string) string {
	re := regexp.MustCompile("^.*/(.+)$")
	match := re.FindStringSubmatch(url)
	f := match[1]
	log.Info(f)
	return f
}

func downloadArtifact(url string) {
	f := artifactName(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "admin123")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	createArtifact("download/"+f, string(body))
}

func downloadURLs() []interface{} {
	var downloadURLsInterfaceArray []interface{}
	continuationTokenMap := continuationTokenRecursion("null")
	log.Info("hi")
	log.Info(continuationTokenMap)
	log.Info(continuationTokenMap)
	log.Info("bye")
	for tokenNumber, token := range continuationTokenMap {
		tokenNumberString := strconv.Itoa(tokenNumber)
		log.Info("ContinuationToken: " + token + "; ContinuationTokenNumber: " + tokenNumberString)
		bytes, err := downloadURL(token)
		if err != nil {
			log.Fatal(err)
		}
		json := string(bytes)

		jq := gojsonq.New().JSONString(json)
		downloadURLsInterface := jq.From("items").Pluck("downloadUrl")

		downloadURLsInterfaceArray = downloadURLsInterface.([]interface{})

		// downloadURLsInterfaceArrayAll = append(downloadURLsInterfaceArrayAll, downloadURLsInterfaceArray)
	}

	log.Info("CP5")
	return downloadURLsInterfaceArray
}

// StoreArtifactsOnDisk download all artifacts from nexus and saves them on disk
func StoreArtifactsOnDisk() {
	log.Info("CP10")
	for i, downloadURL := range downloadURLs() {
		log.Printf("OK: message %d => %s\n", i, downloadURL)
		downloadArtifact(fmt.Sprint(downloadURL))
	}
}
