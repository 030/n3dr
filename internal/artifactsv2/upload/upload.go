package upload

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
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

type mavenParts struct {
	classifier, ext string
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
	cn := connection.Nexus3{BasePathPrefix: n.BasePathPrefix, FQDN: n.FQDN, DownloadDirName: n.DownloadDirName, Pass: n.Pass, User: n.User, HTTPS: n.HTTPS}
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
	exists, err := checkIfLocalArtifactResidesInNexus(continuationToken, localDiskRepo, path, client)
	if err != nil {
		return false, err
	}
	if exists {
		return true, nil
	}
	return false, err
}

func maven(path string, skipErrors bool) (mp mavenParts, err error) {
	regexBase := `^.*\/([\w\-\.]+)\/`

	if runtime.GOOS == "windows" {
		log.Info("N3DR is running on Windows. Correcting the regexBase...")
		regexBase = `^.*\\([\w\-\.]+)\\`
	}

	regexVersion := `(([a-z\d\-]+)|(([a-z\d\.]+)))`
	if rv := os.Getenv("N3DR_MAVEN_UPLOAD_REGEX_VERSION"); rv != "" {
		regexVersion = rv
	}

	regexClassifier := `(-(.*?(\-([\w.]+))?)?)?`
	if rc := os.Getenv("N3DR_MAVEN_UPLOAD_REGEX_CLASSIFIER"); rc != "" {
		regexClassifier = rc
	}

	re := regexp.MustCompile(regexBase + regexVersion + regexClassifier + `\.([a-z]+)$`)

	classifier := ""
	ext := ""
	if re.Match([]byte(path)) {
		result := re.FindAllStringSubmatch(path, -1)
		artifactElements := result[0]
		artifactElementsLength := len(result[0])
		log.Debugf("ArtifactElements: '%s'. ArtifactElementLength: '%d'", artifactElements, artifactElementsLength)
		if artifactElementsLength != 11 {
			err := fmt.Errorf("check whether the regex retrieves ten elements from the artifact. Current: '%s'. Note that element 3 is the artifact itself", artifactElements)
			if skipErrors {
				log.Errorf("skipErrors: '%v'. Error: '%v'", skipErrors, err)
			} else {
				return mp, err
			}
		}

		artifact := artifactElements[3]
		version := artifactElements[1]
		ext = artifactElements[10]

		// Check if the 'version' reported in the artifact name is different from the 'real' version
		if artifactElements[7] != artifactElements[1] {
			classifier = artifactElements[9]
		}

		log.Debugf("Artifact: '%v', Version: '%v', Classifier: '%v', Extension: '%v'.", artifact, version, classifier, ext)
	} else {
		err := fmt.Errorf("check whether regexVersion: '%s' and regexClassifier: '%s' match the artifact: '%s'", regexVersion, regexClassifier, path)
		if skipErrors {
			log.Errorf("skipErrors: '%v'. Error: '%v'", skipErrors, err)
		} else {
			return mp, err
		}
	}
	return mavenParts{classifier: classifier, ext: ext}, nil
}

func UploadSingleArtifact(client *client.Nexus3, path, localDiskRepo, localDiskRepoHome, repoFormat string) error {
	dir := strings.Replace(filepath.Dir(path), localDiskRepoHome+"/", "", -1)
	filename := filepath.Base(path)

	var f, f2, f3 *os.File
	c := components.UploadComponentParams{}
	artifacts.PrintType(repoFormat)
	switch rf := repoFormat; rf {
	case "apt":
		c.Repository = localDiskRepo
		f, err := os.Open(filepath.Clean(path))
		if err != nil {
			return err
		}
		c.AptAsset = f
	case "maven2":
		if filepath.Ext(path) == ".pom" {
			c.Repository = localDiskRepo
			f, err := os.Open(filepath.Clean(path))
			if err != nil {
				return err
			}
			c.Maven2Asset1 = f
			ext1 := "pom"
			c.Maven2Asset1Extension = &ext1

			dir = filepath.Dir(path)
			if err := filepath.WalkDir(dir,
				func(path string, info fs.DirEntry, err error) error {
					if err != nil {
						return err
					}
					if filepath.Ext(path) == ".jar" {
						f2, err = os.Open(filepath.Clean(path))
						if err != nil {
							return err
						}
						mp, err := maven(path, false)
						if err != nil {
							return err
						}
						c.Maven2Asset2 = f2
						c.Maven2Asset2Extension = &mp.ext
						c.Maven2Asset2Classifier = &mp.classifier
					}
					if filepath.Ext(path) == ".zip" {
						f3, err = os.Open(filepath.Clean(path))
						if err != nil {
							return err
						}
						c.Maven2Asset3 = f3
						ext3 := "zip"
						c.Maven2Asset3Extension = &ext3
					}
					return nil
				}); err != nil {
				return err
			}
		}
	case "npm":
		c.Repository = localDiskRepo
		f, err := os.Open(filepath.Clean(path))
		if err != nil {
			return err
		}
		c.NpmAsset = f
	case "nuget":
		c.Repository = localDiskRepo
		f, err := os.Open(filepath.Clean(path))
		if err != nil {
			return err
		}
		c.NugetAsset = f
	case "raw":
		c.Repository = localDiskRepo
		f, err := os.Open(filepath.Clean(path))
		if err != nil {
			return err
		}
		c.RawAsset1 = f
		c.RawDirectory = &dir
		c.RawAsset1Filename = &filename
	case "yum":
		c.Repository = localDiskRepo
		f, err := os.Open(filepath.Clean(path))
		if err != nil {
			return err
		}
		c.YumAsset = f
		c.YumAssetFilename = &filename
	default:
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
			log.Debugf("artifact: '%v' has been uploaded", path)
			return nil
		}

		return fmt.Errorf("cannot upload component: '%s', error: '%v'", path, err)
	}

	files := []*os.File{f, f2, f3}
	for _, file := range files {
		if err := closeFileObjectIfNeeded(file); err != nil {
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

			filesToBeSkipped, err := artifacts.FilesToBeSkipped(filepath.Ext(path))
			if err != nil {
				return err
			}
			if !info.IsDir() && !filesToBeSkipped {
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

func closeFileObjectIfNeeded(f *os.File) error {
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
