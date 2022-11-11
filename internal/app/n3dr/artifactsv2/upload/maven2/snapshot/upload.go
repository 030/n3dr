package snapshot

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/030/n3dr/internal/app/n3dr/artifactsv2/artifacts"
	"github.com/hashicorp/go-retryablehttp"
	log "github.com/sirupsen/logrus"
)

type Nexus3 struct {
	HTTPS                                                   bool
	DownloadDirName, FQDN, Pass, RepoFormat, RepoName, User string
}

func (n *Nexus3) Upload() error {
	if err := filepath.WalkDir(n.DownloadDirName,
		func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				log.Info(path)

				f, err := os.Open(path)
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

				re := regexp.MustCompile(`/` + n.RepoName + `/(.*)$`)
				match := re.FindStringSubmatch(path)
				fmt.Println(match[1])

				u := protocol + "://" + n.FQDN + "/repository/" + n.RepoName + "/" + match[1]
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

				if resp.StatusCode == http.StatusCreated {
					log.Trace("file has been uploaded")
					artifacts.PrintType(n.RepoFormat)
				} else {
					bodyBytes, err := io.ReadAll(resp.Body)
					if err != nil {
						return err
					}
					bodyString := string(bodyBytes)
					log.Info(bodyString)
					log.Errorf("bad status: %s", resp.Status)
				}
			}

			return nil
		}); err != nil {
		return err
	}

	return nil
}
