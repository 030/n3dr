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

	"github.com/030/n3dr/internal/app/n3dr/artifactsv2/artifacts"
	"github.com/030/n3dr/internal/app/n3dr/artifactsv2/upload/maven2/snapshot"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/client"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/client/components"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/client/repository_management"
	"github.com/030/p2iwd/pkg/p2iwd"
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
		return 0, fmt.Errorf("http status code not found in error message: '%w'", err)
	}
	statusCode := match[1]
	statusCodeInt, err := strconv.Atoi(statusCode)
	if err != nil {
		return 0, err
	}

	return statusCodeInt, nil
}

func (n *Nexus3) reposOnDisk() ([]string, error) {
	file, err := os.Open(n.DownloadDirName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	localDiskRepos, err := file.Readdirnames(0)
	if err != nil {
		return nil, err
	}
	log.Infof("found the following localDiskRepos: '%v'", localDiskRepos)
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
		repoName := repo.Name
		if repoName == localDiskRepo {
			repoFormat = repo.Format
		}
	}
	log.Infof("format of repo: '%s' is: '%s'", localDiskRepo, repoFormat)

	return repoFormat, nil
}

func maven(path string, skipErrors bool) (mavenParts, error) {
	regexBase := `^.*\/([\w\-\.]+)\/`

	if runtime.GOOS == "windows" {
		log.Info("N3DR is running on Windows. Correcting the regexBase...")
		regexBase = `^.*\\([\w\-\.]+)\\`
	}

	regexVersion := `(([A-Za-z\d\-_]+)|(([a-z\d\.]+)))`
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
				return mavenParts{}, err
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
			return mavenParts{}, err
		}
	}
	return mavenParts{classifier: classifier, ext: ext}, nil
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
	case "maven2":
		if filepath.Ext(path) == ".pom" {
			c.Repository = localDiskRepo
			var err error
			f, err = os.Open(filepath.Clean(path))
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
		log.Tracef("err: '%v' while uploading file: '%s'", err, path)
		statusCode, uploadStatusErr := uploadStatus(err)
		if uploadStatusErr != nil {
			log.Error(path)
			return uploadStatusErr
		}
		if statusCode == 204 {
			log.Debugf("artifact: '%v' has been uploaded", path)
			return nil
		}

		return fmt.Errorf("cannot upload component: '%s', error: '%w'", path, err)
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
	c := n.Nexus3.Client()
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
				if err := UploadSingleArtifact(c, path, localDiskRepo, localDiskRepoHome, repoFormat); err != nil {
					uploaded, errRegex := regexp.MatchString("status 400", err.Error())
					if errRegex != nil {
						panic(err)
					}
					if uploaded {
						log.Debugf("artifact: '%s' has already been uploaded", path)
						return nil
					}

					log.Errorf("could not upload artifact: '%s', err: '%v'", path, err)
				} else {
					artifacts.PrintType(repoFormat)
				}
			}
			return nil
		}); err != nil {
		return err
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

func (n *Nexus3) maven2SnapshotsUpload(localDiskRepo string) {
	client := n.Nexus3.Client()
	r := repository_management.GetRepository2Params{RepositoryName: localDiskRepo}
	r.WithTimeout(time.Second * 30)
	resp, err := client.RepositoryManagement.GetRepository2(&r)
	if err != nil {
		log.Errorf("cannot determine version policy, repository: '%s'", localDiskRepo)
		return
	}
	vp := resp.Payload.Maven.VersionPolicy
	log.Tracef("VersionPolicy: ''%s", vp)

	if strings.EqualFold(vp, "snapshot") {
		s := snapshot.Nexus3{DownloadDirName: n.DownloadDirName, FQDN: n.FQDN, Pass: n.Pass, RepoFormat: "maven2", RepoName: localDiskRepo, SkipErrors: n.SkipErrors, User: n.User}
		if err := s.Upload(); err != nil {
			if !n.SkipErrors {
				panic(err)
			}
			log.Error(err)
		}
	}
}

func (n *Nexus3) uploadArtifactsSingleDir(localDiskRepo string) {
	log.Infof("Uploading files to Nexus: '%s' repository: '%s'...", n.FQDN, localDiskRepo)

	if localDiskRepo == "p2iwd" {
		h := n.DockerHost + ":" + fmt.Sprint(n.DockerPort)
		pdr := p2iwd.DockerRegistry{Dir: filepath.Join(n.DownloadDirName, "p2iwd"), Host: h, Pass: n.Pass, User: n.User}
		if err := pdr.Upload(); err != nil {
			panic(err)
		}
		return
	}

	repoFormat, err := n.repoFormatLocalDiskRepo(localDiskRepo)
	if err != nil {
		panic(err)
	}
	if repoFormat == "" {
		log.Errorf("repoFormat not detected. Verify whether repo: '%s' resides in Nexus", localDiskRepo)
		return
	}

	if repoFormat == "maven2" {
		n.maven2SnapshotsUpload(localDiskRepo)
	}

	localDiskRepoHome := filepath.Join(n.DownloadDirName, localDiskRepo)
	if err := n.ReadLocalDirAndUploadArtifacts(localDiskRepoHome, localDiskRepo, repoFormat); err != nil {
		panic(err)
	}
}

func (n *Nexus3) Upload() error {
	localDiskRepos, err := n.reposOnDisk()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, localDiskRepo := range localDiskRepos {
		if n.WithoutWaitGroups {
			n.uploadArtifactsSingleDir(localDiskRepo)
		} else {
			wg.Add(1)
			go func(localDiskRepo string) {
				defer wg.Done()

				n.uploadArtifactsSingleDir(localDiskRepo)
			}(localDiskRepo)
		}
	}
	wg.Wait()

	return nil
}
