package cli

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/mholt/archiver"
	log "github.com/sirupsen/logrus"
	"github.com/svenfuchs/jq"
	"github.com/thedevsaddam/gojsonq"
	"gopkg.in/cheggaaa/pb.v1"
)

const (
	downloadDir = "download"
	pingURI     = "/service/metrics/ping"
	assetURI1   = "/service/rest/"
	assetURI2   = "/assets?repository="
	tokenErrMsg = "Token should be either a hexadecimal or \"null\" and not: "
)

// Nexus3 contains the attributes that are used by several functions
type Nexus3 struct {
	URL        string
	User       string
	Pass       string
	Repository string
	APIVersion string
	ZIP        bool
}

func (n Nexus3) downloadURL(token string) ([]byte, error) {
	var err error
	var constructDownloadURL string

	assetURL := n.URL + assetURI1 + n.APIVersion + assetURI2 + n.Repository
	if !(token == "null") {
		constructDownloadURL = assetURL + "&continuationToken=" + token
	}

	var u *url.URL
	if u, err = url.Parse(constructDownloadURL); err != nil {
		return nil, err
	}

	log.Debug("DownloadURL: ", u)

	var bodyBytes []byte
	if bodyBytes, _, err = n.request(u.String()); err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func (n Nexus3) continuationToken(token string) (string, error) {
	var err error
	// The continuationToken should consists of 32 characters and should be a hexadecimal or "null"
	if !((govalidator.IsHexadecimal(token) && govalidator.StringLength(token, "32", "32")) || token == "null") {
		return "", errors.New(tokenErrMsg + token)
	}

	var bodyBytes []byte
	if bodyBytes, err = n.downloadURL(token); err != nil {
		return "", err
	}

	var op jq.Op
	if op, err = jq.Parse(".continuationToken"); err != nil {
		return "", err
	}

	var value []byte
	if value, err = op.Apply(bodyBytes); err != nil {
		return "", err
	}

	tokenWithoutQuotes := strings.Trim(string(value), "\"")

	return tokenWithoutQuotes, nil
}

func (n Nexus3) continuationTokenRecursion(t string) ([]string, error) {
	var err error

	var token string
	if token, err = n.continuationToken(t); err != nil {
		return nil, err
	}

	if token == "null" {
		return []string{token}, nil
	}

	var tokenSlice []string
	if tokenSlice, err = n.continuationTokenRecursion(token); err != nil {
		return nil, err
	}
	return append(tokenSlice, token), nil
}

func createArtifact(d string, f string, content string) error {
	var err error

	if err = os.MkdirAll(d, os.ModePerm); err != nil {
		return err
	}

	var file *os.File
	if file, err = os.Create(filepath.Join(d, f)); err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		return err
	}

	return nil
}

func (n Nexus3) artifactName(url string) (string, string, error) {
	log.Debug("Validate whether: '" + url + "' is an URL")
	if !govalidator.IsURL(url) {
		return "", "", errors.New(url + " is not an URL")
	}

	re := regexp.MustCompile("^.*?/" + n.Repository + "/(.*)/(.+)$")
	match := re.FindStringSubmatch(url)
	if match == nil {
		return "", "", errors.New("URL: '" + url + "' does not seem to contain an artifactName")
	}

	d := match[1]
	f := match[2]
	log.Debug("ArtifactName directory: " + d + " and file: " + f)

	return d, f, nil
}

func (n Nexus3) downloadArtifact(url string) error {
	d, f, err := n.artifactName(url)
	if err != nil {
		return err
	}

	_, bodyString, err := n.request(url)
	if err != nil {
		return err
	}

	if err = createArtifact(filepath.Join(downloadDir, n.Repository, d), f, bodyString); err != nil {
		return err
	}
	return nil
}

func (n Nexus3) downloadURLs() ([]interface{}, error) {
	var downloadURLsInterfaceArrayAll []interface{}
	continuationTokenMap, err := n.continuationTokenRecursion("null")
	if err != nil {
		return nil, err
	}

	count := len(continuationTokenMap)
	if count > 1 {
		log.Info("Assembling downloadURLs '" + n.Repository + "'")
		bar := pb.StartNew(count)
		for tokenNumber, token := range continuationTokenMap {
			tokenNumberString := strconv.Itoa(tokenNumber)
			log.Debug("ContinuationToken: " + token + "; ContinuationTokenNumber: " + tokenNumberString)
			bytes, err := n.downloadURL(token)
			if err != nil {
				return nil, err
			}
			json := string(bytes)

			jq := gojsonq.New().JSONString(json)
			downloadURLsInterface := jq.From("items").Pluck("downloadUrl")

			downloadURLsInterfaceArray := downloadURLsInterface.([]interface{})
			downloadURLsInterfaceArrayAll = append(downloadURLsInterfaceArrayAll, downloadURLsInterfaceArray...)
			bar.Increment()
			time.Sleep(time.Millisecond)
		}
		bar.FinishPrint("Done")
	}
	return downloadURLsInterfaceArrayAll, nil
}

// StoreArtifactsOnDisk downloads all artifacts from nexus and saves them on disk
func (n Nexus3) StoreArtifactsOnDisk() error {
	urls, err := n.downloadURLs()
	if err != nil {
		return err
	}

	countURLs := len(urls)
	if countURLs > 0 {
		log.Info("Backing up artifacts '" + n.Repository + "'")
		bar := pb.StartNew(len(urls))
		for _, downloadURL := range urls {
			_ = n.downloadArtifact(fmt.Sprint(downloadURL))
			bar.Increment()
		}
		bar.FinishPrint("Done")
	} else {
		log.Info("No artifacts found in '" + n.Repository + "'")
	}

	return nil
}

// CreateZip adds all artifacts to a ZIP archive
func (n Nexus3) CreateZip() error {
	if n.ZIP {
		today := time.Now().Format("01-02-2006")
		archiveName := fmt.Sprintf("n3dr-backup-%s.zip", today)
		if err := archiver.Archive([]string{downloadDir}, archiveName); err != nil {
			return err
		}
	}
	return nil
}
