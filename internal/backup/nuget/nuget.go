package nuget

import (
	"crypto/sha512"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/030/n3dr/internal/pkg/backup"
	"github.com/cavaliercoder/grab"
	"github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type Nexus3 struct {
	BaseDir, Endpoint, Password, Repository, Regex, Username string
}

const (
	componentsRepositoryURI = "/service/rest/v1/components?repository="
)

func (n *Nexus3) componentsRepositoryJSON(ct backup.ContinuationToken) (string, error) {
	url := n.Endpoint + componentsRepositoryURI + n.Repository
	if ct != "" {
		url = url + "&continuationToken=" + string(ct)
	}

	resp, err := grequests.Get(url, &grequests.RequestOptions{
		Auth: []string{n.Username, n.Password}})
	if err != nil {
		return "", err
	}

	statusCode := resp.StatusCode
	log.Debugf("URL: '%v'. StatusCode: '%v'. Response: '%s'",
		url, statusCode, resp.String())
	if statusCode != http.StatusOK {
		return "",
			fmt.Errorf("statusCode URL: '%s' not OK, but: '%d'. "+
				"Enable debug mode to get the response",
				url, statusCode)
	}

	responseString := resp.String()
	if responseString == "" {
		return "", fmt.Errorf("response should not be empty. Actual: '%v'",
			responseString)
	}

	return responseString, nil
}

func (n *Nexus3) repositoryJSONAssets(ct backup.ContinuationToken) (string, error) {
	log.Info("CT1: ", ct)
	json, err := n.componentsRepositoryJSON(ct)
	if err != nil {
		return "", err
	}

	assets := gjson.Get(json, "items.#.assets")
	log.Info("CT1a ", assets)
	for _, asset := range assets.Array() {
		var downloadURL string
		if value := gjson.Get(asset.String(), "#.downloadUrl"); !value.Exists() {
			return "", fmt.Errorf("downloadUrl does not exist in json")
		} else {
			log.Info("CT1b ", value.String())
			re := regexp.MustCompile(`^\[\"(.*)\"\]$`)
			match := re.FindStringSubmatch(value.String())
			downloadURL = match[1]
			log.Info("CT1c: ", downloadURL)
		}

		var checksum512 string
		if value := gjson.Get(asset.String(), "#.checksum.sha512"); !value.Exists() {
			return "", fmt.Errorf("512checksum does not exist in json")
		} else {
			log.Info("CT1d ", value.String())
			re := regexp.MustCompile(`^\[\"(.*)\"\]$`)
			match := re.FindStringSubmatch(value.String())
			checksum512 = match[1]
			log.Info("CT1e: ", checksum512)
		}

		filePathDir := filepath.Join(n.BaseDir, n.Repository)
		if err := os.MkdirAll(filePathDir, os.ModePerm); err != nil {
			return "", err
		}

		client := grab.NewClient()
		req, err := grab.NewRequest(filePathDir, downloadURL)
		if err != nil {
			return "", err
		}
		req.HTTPRequest.SetBasicAuth(n.Username, n.Password)
		resp := client.Do(req)
		// log.Info("File ", resp.Filename)

		fmt.Printf("  %v\n", resp.HTTPResponse.Status)

		// start UI loop
		t := time.NewTicker(500 * time.Millisecond)
		defer t.Stop()

	Loop:
		for {
			select {
			case <-t.C:
				fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
					resp.BytesComplete(),
					resp.Size,
					100*resp.Progress())

			case <-resp.Done:
				// download is complete
				break Loop
			}
		}

		// check for errors
		if err := resp.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Download saved to ./%v \n", resp.Filename)

		// dat, err := ioutil.ReadFile(resp.Filename)
		// if err != nil {
		// 	return "", err
		// }
		// checksum512OnDisk := fmt.Sprintf("%x", sha512.Sum512(dat))
		f, err := os.Open(resp.Filename)
		if err != nil {
			return "", err
		}
		defer f.Close()
		h := sha512.New()
		if _, err := io.Copy(h, f); err != nil {
			return "", err
		}
		checksum512OnDisk := fmt.Sprintf("%x", h.Sum(nil))
		if checksum512 != checksum512OnDisk {
			return "", fmt.Errorf("512checksum mismatch on disk: '%v' vs. '%v'. File: '%s'", checksum512, checksum512OnDisk, resp.Filename)
		}
	}

	ct, err = backup.ContinuationTokenInJSON(json)
	if err != nil {
		return "", err
	}
	log.Info("CT2: ", ct)

	if ct == "" {
		log.Info("CT2a: ")
		return "done", nil
	}

	log.Info("CT3: ", ct)
	return n.repositoryJSONAssets(ct)
}
