package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/svenfuchs/jq"
	"github.com/thedevsaddam/gojsonq"
)

const (
	pingURI  = "/service/metrics/ping"
	assetURI = "/service/rest/v1/assets?repository="
)

// Nexus3 contains the attributes that are used by several functions
type Nexus3 struct {
	URL        string
	User       string
	Pass       string
	Repository string
}

func (n Nexus3) downloadURL(token string) ([]byte, error) {
	assetURL := n.URL + assetURI + n.Repository
	url := assetURL
	if !(token == "null") {
		url = assetURL + "&continuationToken=" + token
	}
	log.Info("DownloadURL: ", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(n.User, n.Pass)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Info(resp.StatusCode)
		return nil, errors.New("HTTP response not 200. Does the URL exist?")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func (n Nexus3) continuationToken(token string) string {
	bodyBytes, err := n.downloadURL(token)
	if err != nil {
		log.Fatal(err)
	}

	op, err := jq.Parse(".continuationToken")
	if err != nil {
		//
	}

	value, err := op.Apply(bodyBytes)
	if err != nil {
		//
	}
	var tokenWithoutQuotes string
	tokenWithoutQuotes = strings.Trim(string(value), "\"")

	return tokenWithoutQuotes
}

func (n Nexus3) continuationTokenRecursion(s string) []string {
	token := n.continuationToken(s)
	if token == "null" {
		return []string{token}
	}
	return append(n.continuationTokenRecursion(token), token)
}

func createArtifact(d string, f string, content string) {
	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	file, err2 := os.Create(filepath.Join(d, f))
	if err2 != nil {
		log.Fatal(err)
	}

	file.WriteString(content)
	defer file.Close()
}

func (n Nexus3) artifactName(url string) (string, string) {
	re := regexp.MustCompile("^.*/" + n.Repository + "/(.*)/(.+)$")
	match := re.FindStringSubmatch(url)
	d := match[1]
	f := match[2]
	log.Info(d)
	log.Info(f)
	return d, f
}

func (n Nexus3) downloadArtifact(url string) {
	log.Info("CP3")
	d, f := n.artifactName(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(n.User, n.Pass)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	createArtifact(filepath.Join("download", d), f, string(body))
}

func (n Nexus3) downloadURLs() []interface{} {
	log.Info("CP2")
	var downloadURLsInterfaceArrayAll []interface{}
	continuationTokenMap := n.continuationTokenRecursion("null")

	for tokenNumber, token := range continuationTokenMap {
		tokenNumberString := strconv.Itoa(tokenNumber)
		log.Info("ContinuationToken: " + token + "; ContinuationTokenNumber: " + tokenNumberString)
		bytes, err := n.downloadURL(token)
		if err != nil {
			log.Fatal(err)
		}
		json := string(bytes)

		jq := gojsonq.New().JSONString(json)
		downloadURLsInterface := jq.From("items").Pluck("downloadUrl")

		downloadURLsInterfaceArray := downloadURLsInterface.([]interface{})
		downloadURLsInterfaceArrayAll = append(downloadURLsInterfaceArrayAll, downloadURLsInterfaceArray...)
	}

	return downloadURLsInterfaceArrayAll
}

// StoreArtifactsOnDisk downloads all artifacts from nexus and saves them on disk
func (n Nexus3) StoreArtifactsOnDisk() {
	log.Info("CP1")
	for _, downloadURL := range n.downloadURLs() {
		n.downloadArtifact(fmt.Sprint(downloadURL))
	}
}
