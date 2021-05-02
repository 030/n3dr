package npm

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

var numberOfArtifactsToBeBackedUp int

func (n *Nexus3) AllArtifacts() error {
	errs := make(chan error)
	log.Info(n)
	s, err := n.repositoryJSONAssets(errs, "")
	if err != nil {
		return err
	}
	log.Info("DONE", s)

	time.Sleep(2 * time.Second)
	log.Infof("numberOfArtifactsToBeBackedUp: '%d'", numberOfArtifactsToBeBackedUp)

	for {
		time.Sleep(500 * time.Millisecond)
		log.Info("AAAAA")
		log.Infof("numberOfArtifactsToBeBackedUp: '%d'", numberOfArtifactsToBeBackedUp)

		if err := <-errs; err != nil {
			return err
		}
		if numberOfArtifactsToBeBackedUp == 0 {
			break
		}
		numberOfArtifactsToBeBackedUp--
	}
	return nil
}

func (n *Nexus3) componentsRepositoryJSON(ct backup.ContinuationToken) (string, error) {
	url := n.Endpoint + componentsRepositoryURI + n.Repository
	if ct != "" {
		url = url + "&continuationToken=" + string(ct)
	}
	log.Infof("URL: '%s'", url)

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

func (n *Nexus3) repositoryJSONAssets(errs chan error, ct backup.ContinuationToken) (string, error) {
	log.Info("CT1: ", ct)
	json, err := n.componentsRepositoryJSON(ct)
	if err != nil {
		return "", err
	}
	if !gjson.Valid(json) {
		return "", errors.New("invalid json")
	}

	filePathDir := filepath.Join(n.BaseDir, n.Repository)
	if err := os.MkdirAll(filePathDir, os.ModePerm); err != nil {
		return "", err
	}

	assets := gjson.Get(json, "items.#.assets")
	for _, asset := range assets.Array() {
		log.Info("CCCCC: ")
		for _, artifactElement := range asset.Array() {
			log.Info("DDDDD")
			downloadURL := gjson.Get(artifactElement.String(), "downloadUrl")
			if !downloadURL.Exists() {
				return "", fmt.Errorf("downloadUrl does not exist in json")
			}
			log.Info("downloadUrl: ", downloadURL.Value())
			log.Info("EEEEE")
			checksum512 := gjson.Get(artifactElement.String(), "checksum.sha512")
			if !checksum512.Exists() {
				return "", fmt.Errorf("checksum512 does not exist in json")
			}
			log.Info("checksum512: ", checksum512.Value())
			log.Info("FFFFF")
			go func(checksum512, filePathDir, downloadURL string, n Nexus3) {
				log.Info("GGGGG", n.Username, n.Password)
				s, err := n.downloadAndStoreOnDisk(checksum512, downloadURL, filePathDir)
				errs <- err
				log.Info(s)
				numberOfArtifactsToBeBackedUp++
				log.Info("ARTIFIACTSS ", numberOfArtifactsToBeBackedUp)
			}(checksum512.String(), filePathDir, downloadURL.String(), *n)
		}
	}

	ct, err = backup.ContinuationTokenInJSON(json)
	if err != nil {
		return "", err
	}
	log.Info("CT2: ", ct)

	if ct == "" {
		log.Info("CT2a done: ")
		return "done", nil
	}

	log.Info("CT3: ", ct)
	return n.repositoryJSONAssets(errs, ct)
}

func (n *Nexus3) downloadAndStoreOnDisk(checksum512, downloadURL, filePathDir string) (string, error) {
	log.Info("HHHHH", n.Username, n.Password)
	log.Info(checksum512, filePathDir, downloadURL, n.Username, n.Password)
	client := grab.NewClient()
	req, err := grab.NewRequest(filePathDir, downloadURL)
	if err != nil {
		return "", err
	}

	log.Infof("Downloading %v...\n", req.URL())

	req.HTTPRequest.SetBasicAuth(n.Username, n.Password)
	log.Info("HHHHH2a")
	resp := client.Do(req)
	log.Info("HHHHH3", resp.HTTPResponse.StatusCode)

	fmt.Printf("  %v\n", resp.HTTPResponse.Status)
	log.Info("HHHHH4")

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()
	log.Info("IIIII")
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
		return "", err
	}

	fmt.Printf("Download saved to %v \n", resp.Filename)

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
	log.Info("=======================> FILE STORED")
	return "file stored on disk", nil
}
