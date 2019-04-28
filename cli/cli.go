package cli

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/svenfuchs/jq"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	pingURL  = "http://localhost:8081/service/metrics/ping"
	assetURL = "http://localhost:8081/service/rest/v1/search/assets?repository=maven-releases"
)

func DownloadURL(token string) ([]byte, error) {
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
	bodyBytes, err := DownloadURL(token)
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

func ContinuationTokenRecursion(s string) []string {
	token := continuationToken(s)
	tokenMap = append(tokenMap, token)
	if token == "null" {
		return tokenMap
	}
	return ContinuationTokenRecursion(token)
}
