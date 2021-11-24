package upload

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/030/n3dr/internal/goswagger/client"
	"github.com/030/n3dr/internal/goswagger/client/components"
	"github.com/030/n3dr/internal/pkg/artifactsv2/artifacts"
	"github.com/030/n3dr/internal/pkg/connection"

	log "github.com/sirupsen/logrus"
)

type Nexus3 struct {
	*connection.Nexus3
}

func uploadStatus(err error) (int, error) {
	re := regexp.MustCompile(`status (\d{3})`)
	match := re.FindStringSubmatch(err.Error())
	if len(match) != 2 {
		return 0, fmt.Errorf("http status code not found")
	}
	statusCode := match[1]
	statusCodeInt, err := strconv.Atoi(statusCode)
	if err != nil {
		return 0, err
	}

	return statusCodeInt, nil
}

func (n *Nexus3) reposOnDisk() (localDiskRepos []string, errs []error) {
	file, err := os.Open(n.DownloadDirName)
	if err != nil {
		errs = append(errs, err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			errs = append(errs, err)
			return
		}
	}()
	localDiskRepos, err = file.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}
	return localDiskRepos, nil
}

func (n *Nexus3) repoFormatLocalDiskRepo(localDiskRepo string) (string, error) {
	cn := connection.Nexus3{FQDN: n.FQDN, DownloadDirName: n.DownloadDirName, Pass: n.Pass, User: n.User, HTTPS: n.HTTPS}
	a := artifacts.Nexus3{Nexus3: &cn}
	repos, err := a.Repos()
	if err != nil {
		return "", err
	}
	var repoFormat string
	for _, repo := range repos {
		if repo.Name == localDiskRepo {
			repoFormat = repo.Format
		}
	}
	return repoFormat, nil
}

func checkIfLocalArtifactResidesInNexus(continuationToken, localDiskRepo, path string, client *client.Nexus3) (bool, error) {
	g := components.GetComponentsParams{ContinuationToken: &continuationToken, Repository: localDiskRepo}
	g.WithTimeout(time.Second * 60)
	resp, err := client.Components.GetComponents(&g)
	if err != nil {
		return false, fmt.Errorf("cannot get components from repo: '%s'. Error: '%v'. Does the repo exist in Nexus3?", localDiskRepo, err)
	}
	continuationToken = resp.GetPayload().ContinuationToken
	for _, item := range resp.GetPayload().Items {
		for _, asset := range item.Assets {
			if filepath.Base(asset.Path) == filepath.Base(path) {
				shaType, nexusFileChecksum := artifacts.Checksum(asset)

				localFileChecksum, checksumLocalFileErrs := artifacts.ChecksumLocalFile(path, shaType)
				for _, checksumLocalFileErr := range checksumLocalFileErrs {
					if checksumLocalFileErr != nil {
						return false, checksumLocalFileErr
					}
				}

				if nexusFileChecksum == localFileChecksum {
					log.Debugf("file: '%v' has already been uploaded", path)
					return true, nil
				}
			}
		}
	}
	if continuationToken == "" {
		return false, err
	}
	bestaat, err := checkIfLocalArtifactResidesInNexus(continuationToken, localDiskRepo, path, client)
	if err != nil {
		return false, err
	}
	if bestaat {
		return true, nil
	}
	return false, err
}

func UploadSingleArtifact(client *client.Nexus3, path, localDiskRepo, localDiskRepoHome, repoFormat string) error {
	dir := strings.Replace(filepath.Dir(path), localDiskRepoHome+"/", "", -1)
	filename := filepath.Base(path)

	var f, f2, f3 *os.File
	c := components.UploadComponentParams{}
	switch rf := repoFormat; rf {
	case "apt":
		c.Repository = localDiskRepo
		f, err := os.Open(filepath.Clean(path))
		if err != nil {
			return err
		}
		c.AptAsset = f
	case "npm":
		fmt.Print("*")
		c.Repository = localDiskRepo
		f, err := os.Open(filepath.Clean(path))
		if err != nil {
			return err
		}
		c.NpmAsset = f
	case "nuget":
		fmt.Print("$")
		c.Repository = localDiskRepo
		f, err := os.Open(filepath.Clean(path))
		if err != nil {
			return err
		}
		c.NugetAsset = f
	case "raw":
		fmt.Print("%")
		c.Repository = localDiskRepo
		f, err := os.Open(filepath.Clean(path))
		if err != nil {
			return err
		}
		c.RawAsset1 = f
		c.RawDirectory = &dir
		c.RawAsset1Filename = &filename
	default:
		fmt.Print("?")
		return nil
	}

	if reflect.ValueOf(c).IsZero() {
		log.Debug("no files to be uploaded")
		return nil
	}
	c.WithTimeout(time.Second * 600)
	if err := client.Components.UploadComponent(&c); err != nil {
		statusCode, uploadStatusErr := uploadStatus(err)
		if uploadStatusErr != nil {
			return uploadStatusErr
		}
		if statusCode == 204 {
			log.Infof("artifact: '%v' has been uploaded", path)
			return nil
		}
		if statusCode == 400 {
			log.Debugf("artifact: '%v' has already been uploaded, perhaps 'redeploy' is disabled?", path)
			return nil
		}

		return fmt.Errorf("cannot upload component: '%s', error: '%v'", path, err)
	}

	files := []*os.File{f, f2, f3}
	for _, file := range files {
		if err := boek(file); err != nil {
			return err
		}
	}
	return nil
}

func (n *Nexus3) ReadLocalDirAndUploadArtifacts(localDiskRepoHome, localDiskRepo, repoFormat string) error {
	var errs []error
	var wg sync.WaitGroup
	client := n.Nexus3.Client()
	if err := filepath.WalkDir(localDiskRepoHome,
		func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && !(filepath.Ext(path) == ".sha512") && !(filepath.Ext(path) == ".sha256") && !(filepath.Ext(path) == ".sha1") && !(filepath.Ext(path) == ".md5") {
				wg.Add(1)
				go func() {
					defer wg.Done()
					exists, err := checkIfLocalArtifactResidesInNexus("", localDiskRepo, path, client)
					if err != nil {
						errs = append(errs, err)
						return
					}
					if exists {
						return
					}
					if err := UploadSingleArtifact(client, path, localDiskRepo, localDiskRepoHome, repoFormat); err != nil {
						errs = append(errs, err)
						return
					}
				}()
			}
			return nil
		}); err != nil {
		return err
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func boek(f *os.File) error {
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	if fi.Size() != 0 {
		if err := f.Close(); err != nil {
			return err
		}
	}
	return err
}

func (n *Nexus3) Upload() error {
	localDiskRepos, repoOnDiskErrs := n.reposOnDisk()
	for _, repoOnDiskErr := range repoOnDiskErrs {
		if repoOnDiskErr != nil {
			return repoOnDiskErr
		}
	}

	var errs []error
	var wg sync.WaitGroup
	for _, localDiskRepo := range localDiskRepos {
		wg.Add(1)
		go func(localDiskRepo string) {
			defer wg.Done()
			log.Infof("Uploading files to Nexus: '%s' repository: '%s'...", n.FQDN, localDiskRepo)
			repoFormat, err := n.repoFormatLocalDiskRepo(localDiskRepo)
			if err != nil {
				errs = append(errs, err)
				return
			}
			localDiskRepoHome := filepath.Join(n.DownloadDirName, localDiskRepo)
			if err := n.ReadLocalDirAndUploadArtifacts(localDiskRepoHome, localDiskRepo, repoFormat); err != nil {
				errs = append(errs, err)
				return
			}
		}(localDiskRepo)
	}

	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
