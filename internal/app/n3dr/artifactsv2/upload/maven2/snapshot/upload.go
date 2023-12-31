package snapshot

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/hashicorp/go-retryablehttp"
	log "github.com/sirupsen/logrus"
)

type Nexus3 struct {
	HTTPS, SkipErrors                                              bool
	DownloadDirName, FQDN, Pass, Regex, RepoFormat, RepoName, User string
}

func (n *Nexus3) statusCode(resp *http.Response) error {
	if resp.StatusCode == http.StatusCreated {
		log.Trace("file has been uploaded")
	} else {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		log.Tracef("bodyString: '%s'", bodyString)

		if !n.SkipErrors {
			return fmt.Errorf("bad status: %s", resp.Status)
		}
		log.Errorf("bad status: %s", resp.Status)
	}

	return nil
}

func (n *Nexus3) readRetryAndUpload(path string) error {
	// skip upload of artifact if it does not match the regex
	r, err := regexp.Compile(n.Regex)
	if err != nil {
		return err
	}
	if !r.MatchString(path) {
		log.Debugf("file: '%s' skipped as it does not match regex: '%s'", path, n.Regex)
		return nil
	}

	log.Debugf("reading path: '%s' and uploading it", path)
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	protocol := "http"
	if n.HTTPS {
		protocol = "https"
	}

	log.Debugf("repoName: '%s' and path: '%s'", n.RepoName, path)
	re := regexp.MustCompile(`/` + n.RepoName + `/(.*)$`)
	match := re.FindStringSubmatch(path)
	uri := match[1]
	log.Tracef("uri: '%s'", uri)

	u := protocol + "://" + n.FQDN + "/repository/" + n.RepoName + "/" + uri
	log.Tracef("snapshot upload url: '%s'", u)
	req, err := http.NewRequest("PUT", u, f)
	if err != nil {
		return err
	}
	req.SetBasicAuth(n.User, n.Pass)

	retryClient := retryablehttp.NewClient()
	retryClient.Logger = nil
	retryClient.RetryMax = 10
	standardClient := retryClient.StandardClient()

	resp, err := standardClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	if err := n.statusCode(resp); err != nil {
		return err
	}

	return nil
}

func (n *Nexus3) Upload() error {
	if err := filepath.WalkDir(filepath.Join(n.DownloadDirName, n.RepoName),
		func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				if err := n.readRetryAndUpload(path); err != nil {
					return err
				}
			}

			return nil
		}); err != nil {
		return err
	}

	return nil
}
