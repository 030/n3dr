package artifactsv2

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/030/n3dr/internal/app/n3dr/artifactsv2/artifacts"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/client/components"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
	"github.com/030/n3dr/internal/app/n3dr/s3"
	"github.com/030/p2iwd/pkg/p2iwd"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/mholt/archiver/v3"
	log "github.com/sirupsen/logrus"
)

const chmodDir = 0o750

type Nexus3 struct {
	*connection.Nexus3
}

func (n *Nexus3) download(checksum, downloadedFileChecksum string, asset *models.AssetXO) error {
	if checksum != downloadedFileChecksum {
		log.WithFields(log.Fields{
			"actual":   downloadedFileChecksum,
			"url":      asset.DownloadURL,
			"expected": checksum,
		}).Debug("download artifact as checksum deviates")

		req, err := http.NewRequest("GET", asset.DownloadURL, nil)
		if err != nil {
			return err
		}
		req.SetBasicAuth(n.User, n.Pass)

		retryClient := retryablehttp.NewClient()
		retryClient.Logger = nil
		retryClient.RetryMax = 10
		standardClient := retryClient.StandardClient()

		dir := filepath.Dir(asset.Path)
		if err := os.MkdirAll(filepath.Join(n.DownloadDirName, asset.Repository, dir), chmodDir); err != nil {
			return err
		}

		dst, err := os.Create(filepath.Join(n.DownloadDirName, asset.Repository, asset.Path))
		if err != nil {
			return err
		}
		defer func() {
			if err := dst.Close(); err != nil {
				panic(err)
			}
		}()
		resp, err := standardClient.Do(req)
		if err != nil {
			return err
		}
		defer func() {
			if err := resp.Body.Close(); err != nil {
				panic(err)
			}
		}()

		if resp.StatusCode != http.StatusOK {
			err := fmt.Errorf("bad status: %s", resp.Status)
			return err
		}
		_, err = io.Copy(dst, resp.Body)
		if err != nil {
			return err
		}
		if err := dst.Sync(); err != nil {
			return err
		}

		artifacts.PrintType(asset.Format)
	}
	return nil
}

func (n *Nexus3) downloadSingleArtifact(asset *models.AssetXO, repo string) {
	shaType, checksum := artifacts.Checksum(asset)

	log.WithFields(log.Fields{
		"url":      asset.DownloadURL,
		"format":   asset.Format,
		"checksum": checksum,
	}).Trace("Download artifact")
	assetPath := asset.Path
	filesToBeSkipped, err := artifacts.FilesToBeSkipped(assetPath)
	if err != nil {
		panic(err)
	}
	if !filesToBeSkipped {
		file := filepath.Join(n.DownloadDirName, repo, assetPath)
		downloadedFileChecksum, err := artifacts.ChecksumLocalFile(file, shaType)
		if err != nil {
			panic(err)
		}

		if err := n.download(checksum, downloadedFileChecksum, asset); err != nil {
			panic(err)
		}
	}
}

func (n *Nexus3) downloadIfChecksumMismatchLocalFile(continuationToken, repo string) error {
	var wg sync.WaitGroup

	client, err := n.Nexus3.Client()
	if err != nil {
		return err
	}
	c := components.GetComponentsParams{ContinuationToken: &continuationToken, Repository: repo}
	c.WithTimeout(time.Second * 60)
	resp, err := client.Components.GetComponents(&c)
	if err != nil {
		return fmt.Errorf("cannot get components: '%w'", err)
	}
	continuationToken = resp.GetPayload().ContinuationToken
	for _, item := range resp.GetPayload().Items {
		for _, asset := range item.Assets {
			if n.WithoutWaitGroups || n.WithoutWaitGroupArtifacts {
				n.downloadSingleArtifact(asset, repo)
			} else {
				wg.Add(1)
				go func(assetPreventDataRace *models.AssetXO, repoPreventDataRace string) {
					defer wg.Done()

					n.downloadSingleArtifact(assetPreventDataRace, repoPreventDataRace)
				}(asset, repo)
			}
		}
	}
	wg.Wait()

	if continuationToken == "" {
		return nil
	}

	return n.downloadIfChecksumMismatchLocalFile(continuationToken, repo)
}

func (n *Nexus3) zipFile() error {
	zipFilename := "n3dr-backup-" + time.Now().Format("01-02-2006T15-04-05") + ".zip"
	zipFilenamePath := filepath.Join(n.DownloadDirNameZip, zipFilename)
	if n.AwsBucket != "" || n.ZIP {
		log.Infof("Trying to create a zip file: '%s' in '%s'...", zipFilename, zipFilenamePath)
		if err := archiver.Archive([]string{n.DownloadDirName}, zipFilenamePath); err != nil {
			return err
		}
		log.Infof("Zipfile: '%v' created in '%v'", zipFilename, zipFilenamePath)
	}

	if n.AwsBucket != "" {
		nS3 := s3.Nexus3{AwsBucket: n.AwsBucket, AwsID: n.AwsID, AwsRegion: n.AwsRegion, AwsSecret: n.AwsSecret, ZipFilename: zipFilenamePath}
		if err := nS3.Upload(); err != nil {
			return err
		}
	}

	return nil
}

func (n *Nexus3) SingleRepoBackup() error {
	if err := os.MkdirAll(filepath.Join(n.DownloadDirName, n.RepoName), chmodDir); err != nil {
		return err
	}
	if err := n.downloadIfChecksumMismatchLocalFile("", n.RepoName); err != nil {
		return err
	}

	if err := n.zipFile(); err != nil {
		return err
	}

	return nil
}

func (n *Nexus3) repository(repo *models.AbstractAPIRepository) {
	if repo.Type != "group" {
		if err := os.MkdirAll(filepath.Join(n.DownloadDirName, repo.Name), chmodDir); err != nil {
			panic(err)
		}
		if err := n.downloadIfChecksumMismatchLocalFile("", repo.Name); err != nil {
			panic(err)
		}
	}
}

func (n *Nexus3) Backup() error {
	var wg sync.WaitGroup

	cn := connection.Nexus3{BasePathPrefix: n.BasePathPrefix, FQDN: n.FQDN, DownloadDirName: n.DownloadDirName, Pass: n.Pass, User: n.User, HTTPS: n.HTTPS, DockerHost: n.DockerHost, DockerPort: n.DockerPort, DockerPortSecure: n.DockerPortSecure}
	a := artifacts.Nexus3{Nexus3: &cn}
	repos, err := a.Repos()
	if err != nil {
		return err
	}
	for _, repo := range repos {
		log.Infof("backing up '%s', '%s', %s", repo.Name, repo.Type, repo.Format)
		if repo.Format == "docker" {
			if n.DockerHost == "" || n.DockerPort == 0 {
				return fmt.Errorf("please ensure that the dockerPort and host have been specified, e.g.: --dockerPort 9001 --dockerHost http://localhost")
			}

			h := n.DockerHost + ":" + fmt.Sprint(n.DockerPort)
			pdr := p2iwd.DockerRegistry{Dir: filepath.Join(n.DownloadDirName, "p2iwd"), Host: h, Pass: n.Pass, User: n.User}
			if err := pdr.Backup(); err != nil {
				return err
			}
		} else {
			if n.WithoutWaitGroups || n.WithoutWaitGroupRepositories {
				n.repository(repo)
			} else {
				wg.Add(1)
				go func(repoPreventDataRace *models.AbstractAPIRepository) {
					defer wg.Done()

					n.repository(repoPreventDataRace)
				}(repo)
			}
		}
	}
	wg.Wait()

	if err := n.zipFile(); err != nil {
		return err
	}

	return nil
}
