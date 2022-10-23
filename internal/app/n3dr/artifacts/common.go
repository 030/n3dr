package artifacts

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/mholt/archiver"
	log "github.com/sirupsen/logrus"
	"gopkg.in/validator.v2"
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
	URL                                                                                            string `validate:"nonzero,regexp=^http(s)?://[a-z0-9\\.-]+(:[0-9]+)?(/[a-z0-9\\.-]+)*$"`
	APIVersion, ArtifactType, DownloadDirName, DownloadDirNameZip, Pass, Repository, User, ZipName string
	ZIP                                                                                            bool
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
		return fmt.Errorf("the Nexus3 URL seems to be incorrect. Verify that it complies to the regex that is defined in the 'Nexus3 Struct' and that it does not end with a '/'. Error: '%v'", errs)
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

	bodyBytes, bodyString, errs := n.response(req)
	for _, err := range errs {
		if err != nil {
			return requestJSONResponse{}, err
		}
	}

	return requestJSONResponse{bodyBytes, bodyString}, nil
}

func (n Nexus3) response(req *http.Request) (b []byte, s string, errs []error) {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5
	retryClient.Logger = &RetryLogAdaptor{}
	standardClient := retryClient.StandardClient()
	resp, err := standardClient.Do(req)
	if err != nil {
		errs = append(errs, err)
		return nil, "", errs
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			errs = append(errs, err)
		}
	}()

	bodyBytes, bodyString, err := n.responseBodyString(resp)
	if err != nil {
		errs = append(errs, err)
		return nil, "", errs
	}

	return bodyBytes, bodyString, nil
}

func (n Nexus3) responseBodyString(resp *http.Response) ([]byte, string, error) {
	var bodyString string
	var bodyBytes []byte
	var err error
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, "", err
		}
		bodyString = string(bodyBytes)
		if bodyString == "[ ]" {
			return nil, "", fmt.Errorf("bodystring should not be empty. Did the authentication to '%s' succeed?", n.URL)
		}
	} else {
		return nil, "", fmt.Errorf("ResponseCode: '%s' and Message '%s' for URL: %s", strconv.Itoa(resp.StatusCode), resp.Status, resp.Request.URL)
	}

	return bodyBytes, bodyString, nil
}

// CreateZip adds all artifacts to a ZIP archive
func (n Nexus3) CreateZip(dir string, zipDirDest string) (err error) {
	if n.ZIP {
		if n.ZipName == "" {
			n.ZipName = "n3dr-backup-" + time.Now().Format("01-02-2006T15-04-05") + ".zip"
		}
		if zipDirDest == "" {
			zipDirDest, err = os.Getwd()
			if err != nil {
				return err
			}
		}
		log.Warnf("Trying to create a zip file in: '%v'. Note that this could result in a 'permission denied' issue if N3DR has been installed using snap and is run in a different directory than your own home folder.", zipDirDest)
		err = archiver.Archive([]string{dir}, filepath.Join(zipDirDest, n.ZipName))
		if err != nil {
			return err
		}
		log.Infof("Zipfile: '%v' created in '%v'", n.ZipName, zipDirDest)
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
