package cli

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/svenfuchs/jq"
	"github.com/thedevsaddam/gojsonq"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	log.Info("DownloadURL: ", url)

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

var tokenMap = []string{}

func continuationTokenRecursion(s string) []string {
	token := continuationToken(s)
	tokenMap = append(tokenMap, token)
	if token == "null" {
		return tokenMap
	}
	return continuationTokenRecursion(token)
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

// DownloadURLs is able to find the downloadURLs of all artifacts that reside in nexus
func DownloadURLs() {
	continuationTokenMap := continuationTokenRecursion("null")
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

		downloadURLsInterfaceArray := downloadURLsInterface.([]interface{})

		for i, downloadURL := range downloadURLsInterfaceArray {
			log.Printf("OK: message %d => %s\n", i, downloadURL)
			downloadArtifact(fmt.Sprint(downloadURL))
		}
	}
}
