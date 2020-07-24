package cli

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/mholt/archiver"

	log "github.com/sirupsen/logrus"
)

const projectName = "n3dr"

const (
	CfgFileExt     = "yml"
	HiddenN3DR     = "." + projectName
	DefaultCfgFile = "config"
)

const (
	DefaultCfgFileWithExt = DefaultCfgFile + "." + CfgFileExt
)

// Nexus3 contains the attributes that are used by several functions
type Nexus3 struct {
	URL        string
	User       string
	Pass       string
	Repository string
	APIVersion string
	ZIP        bool
	ZipName    string
}

func (n Nexus3) validate() {
	if n.User == "" {
		log.Debug("Empty user. Verify whether the the subcommand is specified or anonymous mode is used")
	}
	if n.Pass == "" {
		log.Debug("Empty password. Verify whether the 'n3drPass' has been defined in ~/.n3dr.yaml, the subcommand is specified or anonymous mode is used")
	}
}

func (n Nexus3) request(url string) ([]byte, string, error) {
	n.validate()

	log.WithFields(log.Fields{"URL": url, "User": n.User}).Debug("URL Request")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}

	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(n.User, n.Pass)

	bodyBytes, bodyString, err := n.response(req)
	if err != nil {
		return nil, "", err
	}
	return bodyBytes, bodyString, err
}

func (n Nexus3) response(req *http.Request) ([]byte, string, error) {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5
	retryClient.Logger = nil
	standardClient := retryClient.StandardClient()
	resp, err := standardClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	bodyBytes, bodyString, err := n.responseBodyString(resp)
	if err != nil {
		return nil, "", err
	}

	return bodyBytes, bodyString, nil
}

func (n Nexus3) responseBodyString(resp *http.Response) ([]byte, string, error) {
	var bodyString string
	var bodyBytes []byte
	var err error
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, "", err
		}
		bodyString = string(bodyBytes)
		if bodyString == "[ ]" {
			return nil, "", fmt.Errorf("Bodystring should not be empty. Did the authentication to '%s' succeed?", n.URL)
		}
	} else {
		return nil, "", fmt.Errorf("ResponseCode: '%s' and Message '%s' for URL: %s", strconv.Itoa(resp.StatusCode), resp.Status, resp.Request.URL)
	}

	return bodyBytes, bodyString, nil
}

// CreateZip adds all artifacts to a ZIP archive
func (n Nexus3) CreateZip(dir string) error {
	if n.ZIP {
		if n.ZipName == "" {
			n.ZipName = "n3dr-backup-" + time.Now().Format("01-02-2006T15-04-05") + ".zip"
		}
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		log.Warnf("Trying to create a zip file in: '%v'. Note that this could result in a 'permission denied' issue if N3DR has been installed using snap and is run in a different directory than your own home folder.", cwd)
		err = archiver.Archive([]string{dir}, n.ZipName)
		if err != nil {
			return err
		}
		log.Infof("Zipfile: '%v' created in '%v'", n.ZipName, cwd)
	}
	return nil
}
