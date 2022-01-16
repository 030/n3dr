package artifactsv2

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/030/n3dr/internal/goswagger/client/components"
	"github.com/030/n3dr/internal/goswagger/models"
	"github.com/030/n3dr/internal/pkg/artifactsv2/artifacts"
	"github.com/030/n3dr/internal/pkg/connection"
	"github.com/hashicorp/go-retryablehttp"

	log "github.com/sirupsen/logrus"
)

const chmodDir = 0750

type Nexus3 struct {
	*connection.Nexus3
}

func (n *Nexus3) RepositoryNamesV2() error {
	cn := connection.Nexus3{BasePathPrefix: n.BasePathPrefix, FQDN: n.FQDN, DownloadDirName: n.DownloadDirName, Pass: n.Pass, User: n.User, HTTPS: n.HTTPS}
	a := artifacts.Nexus3{Nexus3: &cn}
	repos, err := a.Repos()
	if err != nil {
		return err
	}
	for _, repo := range repos {
		fmt.Println(repo.Name)
	}
	return nil
}

func (n *Nexus3) CountRepositoriesV2() error {
	cn := connection.Nexus3{BasePathPrefix: n.BasePathPrefix, FQDN: n.FQDN, DownloadDirName: n.DownloadDirName, Pass: n.Pass, User: n.User, HTTPS: n.HTTPS}
	a := artifacts.Nexus3{Nexus3: &cn}
	repos, err := a.Repos()
	if err != nil {
		return err
	}
	fmt.Println(len(repos))
	return nil
}

func (n *Nexus3) download(checksum, downloadedFileChecksum string, asset *models.AssetXO) (errs []error) {
	if checksum != downloadedFileChecksum {
		req, err := http.NewRequest("GET", asset.DownloadURL, nil)
		if err != nil {
			errs = append(errs, err)
		}
		req.SetBasicAuth(n.User, n.Pass)

		retryClient := retryablehttp.NewClient()
		retryClient.Logger = nil
		retryClient.RetryMax = 10
		standardClient := retryClient.StandardClient()

		dir := filepath.Dir(asset.Path)
		if err := os.MkdirAll(filepath.Join(n.DownloadDirName, asset.Repository, dir), chmodDir); err != nil {
			errs = append(errs, err)
		}

		out, err := os.Create(filepath.Join(n.DownloadDirName, asset.Repository, asset.Path))
		if err != nil {
			errs = append(errs, err)
		}
		defer func() {
			if err := out.Close(); err != nil {
				errs = append(errs, err)
			}
		}()
		resp, err := standardClient.Do(req)
		if err != nil {
			errs = append(errs, err)
		}
		defer func() {
			if err := resp.Body.Close(); err != nil {
				errs = append(errs, err)
			}
		}()

		if resp.StatusCode != http.StatusOK {
			err := fmt.Errorf("bad status: %s", resp.Status)
			errs = append(errs, err)
		}
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			errs = append(errs, err)
		}

		artifacts.PrintType(asset.Format)
	}
	return nil
}

func (n *Nexus3) downloadIfChecksumMismatchLocalFile(continuationToken, repo string) error {
	var errs []error
	var wg sync.WaitGroup
	client := n.Nexus3.Client()
	c := components.GetComponentsParams{ContinuationToken: &continuationToken, Repository: repo}
	c.WithTimeout(time.Second * 60)
	resp, err := client.Components.GetComponents(&c)
	if err != nil {
		return fmt.Errorf("cannot get components: '%v'", err)
	}
	continuationToken = resp.GetPayload().ContinuationToken
	for _, item := range resp.GetPayload().Items {
		for _, asset := range item.Assets {
			wg.Add(1)
			go func(asset *models.AssetXO) {
				defer wg.Done()
				shaType, checksum := artifacts.Checksum(asset)

				log.Debugf("downloadURL: '%s', format: '%s', checksum: '%s'", asset.DownloadURL, asset.Format, checksum)
				assetPath := asset.Path
				filesToBeSkipped, err := artifacts.FilesToBeSkipped(assetPath)
				if err != nil {
					errs = append(errs, err)
					return
				}
				if !filesToBeSkipped {
					file := filepath.Join(n.DownloadDirName, repo, assetPath)
					downloadedFileChecksum, checksumLocalFileErrs := artifacts.ChecksumLocalFile(file, shaType)
					for _, checksumLocalFileErr := range checksumLocalFileErrs {
						if checksumLocalFileErr != nil {
							errs = append(errs, checksumLocalFileErr)
							return
						}
					}

					downloadErrs := n.download(checksum, downloadedFileChecksum, asset)
					for _, downloadErr := range downloadErrs {
						if downloadErr != nil {
							errs = append(errs, downloadErr)
							return
						}
					}
				}
			}(asset)
		}
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	if continuationToken == "" {
		return nil
	}
	if err := n.downloadIfChecksumMismatchLocalFile(continuationToken, repo); err != nil {
		return err
	}

	return nil
}

func (n *Nexus3) Backup() error {
	var errs []error
	var wg sync.WaitGroup

	cn := connection.Nexus3{BasePathPrefix: n.BasePathPrefix, FQDN: n.FQDN, DownloadDirName: n.DownloadDirName, Pass: n.Pass, User: n.User, HTTPS: n.HTTPS}
	a := artifacts.Nexus3{Nexus3: &cn}
	repos, err := a.Repos()
	if err != nil {
		return err
	}
	for _, repo := range repos {
		wg.Add(1)
		go func(repo *models.AbstractAPIRepository) {
			defer wg.Done()
			if repo.Type != "group" {
				log.Infof("backing up '%s', '%s', %s", repo.Name, repo.Type, repo.Format)
				if err := os.MkdirAll(filepath.Join(n.DownloadDirName, repo.Name), chmodDir); err != nil {
					errs = append(errs, err)
				}
				if err := n.downloadIfChecksumMismatchLocalFile("", repo.Name); err != nil {
					errs = append(errs, err)
				}
			}
		}(repo)
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
