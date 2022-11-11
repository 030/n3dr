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

	"github.com/030/multipart/pkg/multipart"
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
		return 0, fmt.Errorf("http status code not found in error message: '%v'", err)
	}
	statusCode := match[1]
	statusCodeInt, err := strconv.Atoi(statusCode)
	if err != nil {
		return 0, err
	}

	return statusCodeInt, nil
}

func (n *Nexus3) reposOnDisk() (localDiskRepos []string, err error) {
	file, err := os.Open(n.DownloadDirName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	localDiskRepos, err = file.Readdirnames(0)
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

func maven(path string, skipErrors bool) (mp mavenParts, err error) {
	regexClassifier := `([a-z]+)?`
	if rc := os.Getenv("N3DR_MAVEN_UPLOAD_REGEX_CLASSIFIER"); rc != "" {
		regexClassifier = rc
	}

	regexExtension := `\.([a-z]+)$`
	regex := regexClassifier + regexExtension
	re := regexp.MustCompile(regex)

	classifier := ""
	ext := ""
	if re.Match([]byte(path)) {
		result := re.FindAllStringSubmatch(path, -1)
		artifactElements := result[0]
		artifactElementsLength := len(result[0])
		log.Debugf("ArtifactElements: '%s'. ArtifactElementLength: '%d'", artifactElements, artifactElementsLength)
		if !(artifactElementsLength == 2 || artifactElementsLength == 3) {
			err := fmt.Errorf("check whether the regex retrieves ten elements from the artifact. Current: '%s'. Note that element 3 is the artifact itself", artifactElements)
			if skipErrors {
				log.Errorf("skipErrors: '%v'. Error: '%v'", skipErrors, err)
			} else {
				return mp, err
			}
		}

		if artifactElementsLength == 3 {
			classifier = artifactElements[1]
			ext = artifactElements[2]
		} else {
			classifier = artifactElements[0]
			ext = artifactElements[1]
		}

		log.Debugf("Classifier: '%v', Extension: '%v'.", classifier, ext)
	} else {
		err := fmt.Errorf("check whether regex: '%s' match the artifact: '%s'", regex, path)
		if skipErrors {
			log.Errorf("skipErrors: '%v'. Error: '%v'", skipErrors, err)
		} else {
			return mp, err
		}
	}
	return mavenParts{classifier: classifier, ext: ext}, nil
}

func (n *Nexus3) UploadSingleArtifact(client *client.Nexus3, path, localDiskRepo, localDiskRepoHome, repoFormat string) error {
	dir := strings.Replace(filepath.Dir(path), localDiskRepoHome+"/", "", -1)
	filename := filepath.Base(path)

	// multipartString := "maven2.asset1=@../../test/testdata/upload/maven-releases/some/group1/File_1/1.0.0-2/File_1-1.0.0-2.pom,maven2.asset1.extension=pom,maven2.asset2=@../../test/testdata/upload/maven-releases/some/group1/File_1/1.0.0-2/File_1-1.0.0-2.jar,maven2.asset2.extension=jar"
	var multipartString string
	i := 1
	var f *os.File
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
			t := strings.Replace(path, ".pom", "", -1)
			log.Warn(t)
			// log.Warn(path)
			multipartString = "maven2.asset1=@" + path + ",maven2.asset1.extension=pom"

			// c.Repository = localDiskRepo
			// var err error
			// f, err = os.Open(filepath.Clean(path))
			// if err != nil {
			// 	return err
			// }
			// c.Maven2Asset1 = f
			// ext1 := "pom"
			// c.Maven2Asset1Extension = &ext1

			dir = filepath.Dir(path)
			if err := filepath.WalkDir(dir,
				func(path string, info fs.DirEntry, err error) error {
					if err != nil {
						return err
					}

					filesToBeSkipped, err := artifacts.FilesToBeSkipped(filepath.Ext(path))
					if err != nil {
						return err
					}

					ext := filepath.Ext(path)
					if !info.IsDir() && !filesToBeSkipped && ext != ".pom" && strings.Contains(path, t) {
						i++
						// maven2.asset5.classifier=sources,

						mp, err := maven(path, false)
						if err != nil {
							return err
						}

						multipartString = multipartString + ",maven2.asset" + strconv.Itoa(i) + "=@" + path + ",maven2.asset" + strconv.Itoa(i) + ".extension=" + mp.ext
						if mp.classifier != "" {
							multipartString = multipartString + ",maven2.asset" + strconv.Itoa(i) + ".classifier=" + mp.classifier
						}

						// 			if c.Maven2Asset2Extension == nil {
						// 				f2, err = os.Open(filepath.Clean(path))
						// 				if err != nil {
						// 					return err
						// 				}
						// 				mp, err := maven(path, false)
						// 				if err != nil {
						// 					return err
						// 				}
						// 				log.Infof("path: '%s', extension: '%s', classifier: '%s'", path, mp.ext, mp.classifier)
						// 				c.Maven2Asset2 = f2
						// 				c.Maven2Asset2Extension = &mp.ext
						// 				c.Maven2Asset2Classifier = &mp.classifier
						// 				return nil
						// 			}

						// 			if c.Maven2Asset3Extension == nil {
						// 				f3, err = os.Open(filepath.Clean(path))
						// 				if err != nil {
						// 					return err
						// 				}
						// 				mp, err := maven(path, false)
						// 				if err != nil {
						// 					return err
						// 				}
						// 				log.Infof("path: '%s',extension: '%s', classifier: '%s'", path, mp.ext, mp.classifier)
						// 				c.Maven2Asset3 = f3
						// 				c.Maven2Asset3Extension = &mp.ext
						// 				c.Maven2Asset3Classifier = &mp.classifier
						// 				return nil
						// 			}

						// log.Errorf("cannot upload more than three components at once. Affected path: '%s'", path)
					}
					return nil
				}); err != nil {
				return err
			}

			// multipartString := strings.TrimSuffix(sb.String(), ",")

			log.Warn(multipartString)
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

	if repoFormat != "maven2" {
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

			return fmt.Errorf("cannot upload component: '%s', error: '%v'", path, err)
		}

		if err := closeFileObjectIfNeeded(f); err != nil {
			return err
		}
	} else {
		protocol := "http"
		if n.HTTPS {
			protocol = "https"
		}
		url := protocol + "://" + n.FQDN + "/service/rest/v1/components?repository=" + localDiskRepo
		log.WithFields(log.Fields{
			"multipart": multipartString,
			"url":       url,
		}).Debug("URL")
		u := multipart.Upload{URL: url, Username: n.User, Password: n.Pass}
		if err := u.Upload(multipartString); err != nil {
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

			if !info.IsDir() && !filesToBeSkipped && repoFormat == "maven2" && filepath.Ext(path) == ".pom" {
				if err := n.UploadSingleArtifact(c, path, localDiskRepo, localDiskRepoHome, repoFormat); err != nil {
					uploaded, errRegex := regexp.MatchString("HTTPStatusCode: '400'; ResponseMessage: 'Repository does not allow updating assets:", err.Error())
					if errRegex != nil {
						panic(err)
					}
					if uploaded {
						log.Debugf("artifact: '%s' has already been uploaded", path)
						return nil
					}

					log.Errorf("could not upload artifact: '%s'. Err: '%w'", path, err)
				} else {
					artifacts.PrintType(repoFormat)
				}
			}

			if !info.IsDir() && !filesToBeSkipped && repoFormat != "maven2" {
				if err := n.UploadSingleArtifact(c, path, localDiskRepo, localDiskRepoHome, repoFormat); err != nil {
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

func (n *Nexus3) Upload() error {
	localDiskRepos, err := n.reposOnDisk()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, localDiskRepo := range localDiskRepos {
		wg.Add(1)
		go func(localDiskRepo string) {
			defer wg.Done()
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
				client := n.Nexus3.Client()
				r := repository_management.GetRepository2Params{RepositoryName: localDiskRepo}
				r.WithTimeout(time.Second * 30)

				resp, err := client.RepositoryManagement.GetRepository2(&r)
				if err != nil {
					log.Errorf("cannot determine version policy, repository: '%s'", localDiskRepo)
					return
				}
				if resp.Payload.Maven.VersionPolicy == "Snapshot" {
					// log.Errorf("upload to snapshot repositories not supported. Affected repository: '%s'", localDiskRepo)

					s := snapshot.Nexus3{DownloadDirName: n.DownloadDirName, FQDN: n.FQDN, Pass: n.Pass, RepoFormat: repoFormat, RepoName: localDiskRepo, User: n.User}
					if err := s.Upload(); err != nil {
						panic(err)
					}
					return
				}
			}

			localDiskRepoHome := filepath.Join(n.DownloadDirName, localDiskRepo)
			if err := n.ReadLocalDirAndUploadArtifacts(localDiskRepoHome, localDiskRepo, repoFormat); err != nil {
				panic(err)
			}
		}(localDiskRepo)
	}
	wg.Wait()

	return nil
}
