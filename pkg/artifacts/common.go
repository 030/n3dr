package artifacts

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/mholt/archiver"
	"gopkg.in/validator.v2"

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
	URL             string `validate:"nonzero,regexp=^http(s)?://.*[a-z]+(:[0-9]+)?$"`
	User            string
	Pass            string
	Repository      string
	APIVersion      string
	ZIP             bool
	ZipName         string
	DownloadDirName string
	ArtifactType    string
}

// RetryLogAdaptor adapts the retryablehttp.Logger interface to the logrus logger.
type RetryLogAdaptor struct{}

func (n Nexus3) validate() {
	if n.User == "" {
		log.Debug("Empty user. Verify whether the the subcommand is specified or anonymous mode is used")
	}
	if n.Pass == "" {
		log.Debug("Empty password. Verify whether the 'n3drPass' has been defined in ~/.n3dr.yaml, the subcommand is specified or anonymous mode is used")
	}
}

func (n *Nexus3) ValidateNexusURL() error {
	if errs := validator.Validate(n); errs != nil {
		return fmt.Errorf("The Nexus3 URL seems to be incorrect. Ensure that it does not end with a '/'. Error: '%v'", errs)
	}
	return nil
}

type requestJSONResponse struct {
	bytes   []byte
	strings string
}

func (n Nexus3) request(url string) (requestJSONResponse, error) {
	n.validate()

	log.WithFields(log.Fields{"URL": url, "User": n.User}).Debug("URL Request")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return requestJSONResponse{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(n.User, n.Pass)

	bodyBytes, bodyString, err := n.response(req)
	if err != nil {
		return requestJSONResponse{}, err
	}
	return requestJSONResponse{bodyBytes, bodyString}, nil
}

func (n Nexus3) response(req *http.Request) ([]byte, string, error) {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5
	retryClient.Logger = &RetryLogAdaptor{}
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

// Printf implements the retryablehttp.Logger interface
func (*RetryLogAdaptor) Printf(fmtStr string, vars ...interface{}) {
	switch {
	case strings.HasPrefix(fmtStr, "[DEBUG]"):
		log.Debugf(strings.TrimSpace(fmtStr[7:]), vars...)
	case strings.HasPrefix(fmtStr, "[ERR]"):
		log.Errorf(strings.TrimSpace(fmtStr[5:]), vars...)
	case strings.HasPrefix(fmtStr, "[ERROR]"):
		log.Errorf(strings.TrimSpace(fmtStr[7:]), vars...)
	case strings.HasPrefix(fmtStr, "[WARN]"):
		log.Warnf(strings.TrimSpace(fmtStr[6:]), vars...)
	case strings.HasPrefix(fmtStr, "[INFO]"):
		log.Infof(strings.TrimSpace(fmtStr[6:]), vars...)
	default:
		log.Printf(fmtStr, vars...)
	}
}
