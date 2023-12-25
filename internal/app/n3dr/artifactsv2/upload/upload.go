package upload

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
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
	"github.com/hashicorp/go-retryablehttp"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

type artifactFiles struct {
	f2, f3, f4, f5, f6, f7 *os.File
}

type mavenParts struct {
	artifact, classifier, ext, version string
}

type Nexus3 struct {
	*connection.Nexus3
}

type repoFormatAndType struct {
	format, repoType string
}

// prevent race condition while using global variables in conjunction with go
// routines, see: "Unprotected global variable" paragraph in
// https://go.dev/doc/articles/race_detector
var (
	checkedMavenFolders   []string
	checkedMavenFoldersMu sync.Mutex
)

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

func (n *Nexus3) repoFormatLocalDiskRepo(localDiskRepo string) (repoFormatAndType, error) {
	cn := connection.Nexus3{
		BasePathPrefix:  n.BasePathPrefix,
		FQDN:            n.FQDN,
		DownloadDirName: n.DownloadDirName,
		Pass:            n.Pass,
		User:            n.User,
		HTTPS:           n.HTTPS,
	}
	a := artifacts.Nexus3{Nexus3: &cn}
	repos, err := a.Repos()
	if err != nil {
		return repoFormatAndType{}, err
	}

	var repoFormat string
	var repoType string
	for _, repo := range repos {
		repoName := repo.Name
		if repoName == localDiskRepo {
			repoFormat = repo.Format
			repoType = repo.Type
		}
	}
	log.Infof("format of repo: '%s' is: '%s' and repoType: '%s'", localDiskRepo, repoFormat, repoType)

	return repoFormatAndType{repoFormat, repoType}, nil
}

func maven(path string, skipErrors bool) (mp mavenParts, err error) {
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

	artifact := ""
	classifier := ""
	ext := ""
	version := ""

	if re.MatchString(path) {
		result := re.FindAllStringSubmatch(path, -1)
		artifactElements := result[0]
		artifactElementsLength := len(result[0])
		log.Tracef("ArtifactElements: '%s'. ArtifactElementLength: '%d'", artifactElements, artifactElementsLength)
		if artifactElementsLength != 11 {
			err := fmt.Errorf("check whether the regex retrieves ten elements from the artifact. Current: '%s'. Note that element 3 is the artifact itself", artifactElements)
			if skipErrors {
				log.Errorf("skipErrors: '%v'. Error: '%v'", skipErrors, err)
			} else {
				return mp, err
			}
		}

		artifact = artifactElements[3]
		version = artifactElements[1]
		ext = artifactElements[10]

		// Check if the 'version' reported in the artifact name is different from the 'real' version
		if artifactElements[7] != artifactElements[1] {
			classifier = artifactElements[9]
		}

		log.Tracef("Artifact: '%v', Version: '%v', Classifier: '%v', Extension: '%v'.", artifact, version, classifier, ext)
	} else {
		err := fmt.Errorf("check whether regexVersion: '%s' and regexClassifier: '%s' match the artifact: '%s'", regexVersion, regexClassifier, path)
		if skipErrors {
			log.Errorf("skipErrors: '%v'. Error: '%v'", skipErrors, err)
		} else {
			return mp, err
		}
	}
	return mavenParts{artifact: artifact, classifier: classifier, ext: ext, version: version}, nil
}

func (af artifactFiles) mavenJarAndOtherExtensions(c *components.UploadComponentParams, fileNameWithoutExtIncludingDir string, skipErrors bool) error {
	filePathJar := fileNameWithoutExtIncludingDir + ".jar"
	filePathSourcesJar := fileNameWithoutExtIncludingDir + "-sources.jar"
	filePathJavadocJar := fileNameWithoutExtIncludingDir + "-javadoc.jar"

	if _, err := os.Stat(filePathSourcesJar); err == nil {
		af.f2, err = os.Open(filepath.Clean(filePathSourcesJar))
		if err != nil {
			return err
		}
		mp, err := maven(filePathSourcesJar, skipErrors)
		if err != nil {
			return err
		}
		c.Maven2Asset2 = af.f2
		c.Maven2Asset2Extension = &mp.ext
		c.Maven2Asset2Classifier = &mp.classifier

		log.WithFields(log.Fields{
			"file":       c.Maven2Asset2.Name(),
			"extension":  *c.Maven2Asset2Extension,
			"classifier": *c.Maven2Asset2Classifier,
		}).Trace("Maven2 asset2")
	}

	if _, err := os.Stat(filePathJavadocJar); err == nil {
		af.f7, err = os.Open(filepath.Clean(filePathJavadocJar))
		if err != nil {
			return err
		}
		mp, err := maven(filePathJavadocJar, skipErrors)
		if err != nil {
			return err
		}
		c.Maven2Asset7 = af.f7
		c.Maven2Asset7Extension = &mp.ext
		c.Maven2Asset7Classifier = &mp.classifier

		log.WithFields(log.Fields{
			"file":       c.Maven2Asset7.Name(),
			"extension":  *c.Maven2Asset7Extension,
			"classifier": *c.Maven2Asset7Classifier,
		}).Trace("Maven2 asset7")
	}

	if _, err := os.Stat(filePathJar); err == nil {
		af.f3, err = os.Open(filepath.Clean(filePathJar))
		if err != nil {
			return err
		}
		mp, err := maven(filePathJar, skipErrors)
		if err != nil {
			return err
		}
		c.Maven2Asset3 = af.f3
		c.Maven2Asset3Extension = &mp.ext
		c.Maven2Asset3Classifier = &mp.classifier

		log.WithFields(log.Fields{
			"file":       c.Maven2Asset3.Name(),
			"extension":  *c.Maven2Asset3Extension,
			"classifier": *c.Maven2Asset3Classifier,
		}).Trace("Maven2 asset3")
	}

	filePathWar := fileNameWithoutExtIncludingDir + ".war"
	if _, err := os.Stat(filePathWar); err == nil {
		af.f4, err = os.Open(filepath.Clean(filePathWar))
		if err != nil {
			return err
		}
		mp, err := maven(filePathWar, skipErrors)
		if err != nil {
			return err
		}
		c.Maven2Asset4 = af.f4
		c.Maven2Asset4Extension = &mp.ext
		c.Maven2Asset4Classifier = &mp.classifier

		log.WithFields(log.Fields{
			"file":       c.Maven2Asset4.Name(),
			"extension":  *c.Maven2Asset4Extension,
			"classifier": *c.Maven2Asset4Classifier,
		}).Trace("Maven2 asset4")
	}

	filePathZip := fileNameWithoutExtIncludingDir + ".zip"
	if _, err := os.Stat(filePathZip); err == nil {
		af.f5, err = os.Open(filepath.Clean(filePathZip))
		if err != nil {
			return err
		}
		c.Maven2Asset5 = af.f5
		ext5 := "zip"
		c.Maven2Asset5Extension = &ext5

		log.WithFields(log.Fields{
			"file":      c.Maven2Asset5.Name(),
			"extension": *c.Maven2Asset5Extension,
		}).Trace("Maven2 asset5")
	}

	filePathModule := fileNameWithoutExtIncludingDir + ".module"
	if _, err := os.Stat(filePathModule); err == nil {
		af.f6, err = os.Open(filepath.Clean(filePathModule))
		if err != nil {
			return err
		}
		c.Maven2Asset6 = af.f6
		ext6 := "module"
		c.Maven2Asset6Extension = &ext6

		log.WithFields(log.Fields{
			"file":      c.Maven2Asset6.Name(),
			"extension": *c.Maven2Asset6Extension,
		}).Trace("Maven2 asset6")
	}

	return nil
}

func removeExtension(fileName string) string {
	ext := filepath.Ext(fileName)
	re := regexp.MustCompile(`(-[a-zA-Z]+)?` + ext + `$`)
	result := re.ReplaceAllString(fileName, "")
	filename := filepath.Base(result)
	return filename
}

func (n *Nexus3) checkLocalChecksumAndCompareWithOneInRemote(f, localDiskRepo, dir, filename string) (bool, error) {
	identical := false

	downloadedFileChecksum, err := artifacts.ChecksumLocalFile(f, "")
	if err != nil {
		return false, err
	}
	log.Debugf("checksum of file: '%s' is '%s'", f, downloadedFileChecksum)

	retryClient := retryablehttp.NewClient()
	retryClient.Logger = nil
	retryClient.RetryMax = 3
	standardClient := retryClient.StandardClient()

	scheme := "http"
	if *n.HTTPS {
		scheme = "https"
	}

	u := scheme + "://" + n.FQDN + "/repository/" + localDiskRepo + "/" + dir + "/" + filename + ".sha512"
	log.Debugf("upload URL: '%s'", u)

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return false, err
	}
	req.SetBasicAuth(n.User, n.Pass)

	resp, err := standardClient.Do(req)
	if err != nil {
		return false, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	bodyString := string(bodyBytes)
	log.Debugf("checksum of artifact in nexus3: '%s'", bodyString)

	if bodyString == downloadedFileChecksum {
		identical = true
	}

	return identical, nil
}

func (n *Nexus3) UploadSingleArtifact(client *client.Nexus3, path, localDiskRepo, localDiskRepoHome, repoFormat string, skipErrors bool) (bool, error) {
	dir := strings.Replace(filepath.Dir(path), localDiskRepoHome+"/", "", -1)
	filename := filepath.Base(path)

	var f *os.File
	af := artifactFiles{}

	if identical, _ := n.checkLocalChecksumAndCompareWithOneInRemote(filepath.Clean(path), localDiskRepo, dir, filename); identical {
		log.Debugf("artifact: '%s' has already been uploaded", filename)
		return true, nil
	}

	c := components.UploadComponentParams{}
	switch rf := repoFormat; rf {
	case "apt":
		c.Repository = localDiskRepo
		f, err := os.Open(filepath.Clean(path))
		if err != nil {
			return false, err
		}
		c.AptAsset = f
	case "maven2":
		dirPath := filepath.Dir(path)
		fileNameWithoutExt := removeExtension(path)
		fileNameWithoutExtIncludingDir := filepath.Join(dirPath, fileNameWithoutExt)
		filePathPom := fileNameWithoutExtIncludingDir + ".pom"

		if slices.Contains(checkedMavenFolders, dirPath) {
			return false, nil
		}

		if _, err := os.Stat(filePathPom); err == nil {
			log.Debugf("folder: '%s' contains a pom file: '%s'", dirPath, filePathPom)

			c.Repository = localDiskRepo

			if err := af.mavenJarAndOtherExtensions(&c, fileNameWithoutExtIncludingDir, skipErrors); err != nil {
				return false, err
			}

			var err error
			f, err = os.Open(filepath.Clean(filePathPom))
			if err != nil {
				return false, err
			}
			c.Maven2Asset1 = f
			ext1 := "pom"
			c.Maven2Asset1Extension = &ext1

			maven2Asset1 := "empty"
			maven2Asset2 := "empty"
			maven2Asset3 := "empty"
			maven2Asset4 := "empty"
			maven2Asset5 := "empty"
			maven2Asset6 := "empty"
			maven2Asset7 := "empty"

			if c.Maven2Asset1 != nil {
				maven2Asset1 = c.Maven2Asset1.Name()
			}
			if c.Maven2Asset2 != nil {
				maven2Asset2 = c.Maven2Asset2.Name()
			}
			if c.Maven2Asset3 != nil {
				maven2Asset3 = c.Maven2Asset3.Name()
			}
			if c.Maven2Asset4 != nil {
				maven2Asset4 = c.Maven2Asset4.Name()
			}
			if c.Maven2Asset5 != nil {
				maven2Asset5 = c.Maven2Asset5.Name()
			}
			if c.Maven2Asset6 != nil {
				maven2Asset6 = c.Maven2Asset6.Name()
			}
			if c.Maven2Asset7 != nil {
				maven2Asset7 = c.Maven2Asset7.Name()
			}

			log.WithFields(log.Fields{
				"1": maven2Asset1,
				"2": maven2Asset2,
				"3": maven2Asset3,
				"4": maven2Asset4,
				"5": maven2Asset5,
				"6": maven2Asset6,
				"7": maven2Asset7,
			}).Debug("Maven2 asset upload")
		} else {
			log.Debugf("folder: '%s' does not contain a pom file: '%s'", dirPath, filePathPom)

			c.Repository = localDiskRepo

			if err := af.mavenJarAndOtherExtensions(&c, fileNameWithoutExtIncludingDir, skipErrors); err != nil {
				return false, err
			}

			//
			mp, err := maven(path, skipErrors)
			if err != nil {
				return false, err
			}
			c.Maven2ArtifactID = &mp.artifact
			c.Maven2Version = &mp.version

			// Match "/some/group/" and capture the group name.
			regex := `^` + localDiskRepoHome + `/([\w+\/]+)/` + mp.artifact
			re := regexp.MustCompile(regex)
			groupID := ""

			// Extract the group name from the path.
			match := re.FindStringSubmatch(path)
			if len(match) >= 2 {
				groupID = match[1]
				groupID = strings.ReplaceAll(groupID, `/`, `.`)
			} else {
				return false, fmt.Errorf("groupID should not be empty, path: '%s' and regex: '%s'", path, regex)
			}
			c.Maven2GroupID = &groupID
			generatePOM := true
			c.Maven2GeneratePom = &generatePOM

			maven2Asset1 := "empty"
			maven2Asset2 := "empty"
			maven2Asset3 := "empty"
			maven2Asset4 := "empty"
			maven2Asset5 := "empty"
			maven2Asset6 := "empty"
			maven2Asset7 := "empty"

			if c.Maven2Asset1 != nil {
				maven2Asset1 = c.Maven2Asset1.Name()
			}
			if c.Maven2Asset2 != nil {
				maven2Asset2 = c.Maven2Asset2.Name()
			}
			if c.Maven2Asset3 != nil {
				maven2Asset3 = c.Maven2Asset3.Name()
			}
			if c.Maven2Asset4 != nil {
				maven2Asset4 = c.Maven2Asset4.Name()
			}
			if c.Maven2Asset5 != nil {
				maven2Asset5 = c.Maven2Asset5.Name()
			}
			if c.Maven2Asset6 != nil {
				maven2Asset6 = c.Maven2Asset6.Name()
			}
			if c.Maven2Asset7 != nil {
				maven2Asset7 = c.Maven2Asset7.Name()
			}

			log.WithFields(log.Fields{
				"1": maven2Asset1,
				"2": maven2Asset2,
				"3": maven2Asset3,
				"4": maven2Asset4,
				"5": maven2Asset5,
				"6": maven2Asset6,
				"7": maven2Asset7,
			}).Debug("Maven2 asset upload")
		}

		checkedMavenFoldersMu.Lock()
		defer checkedMavenFoldersMu.Unlock()
		checkedMavenFolders = append(checkedMavenFolders, dirPath)
		checkedMavenFolders = lo.Uniq(checkedMavenFolders)
	case "npm":
		c.Repository = localDiskRepo
		f, err := os.Open(filepath.Clean(path))
		if err != nil {
			return false, err
		}
		c.NpmAsset = f
	case "nuget":
		c.Repository = localDiskRepo
		f, err := os.Open(filepath.Clean(path))
		if err != nil {
			return false, err
		}
		c.NugetAsset = f
	case "raw":
		c.Repository = localDiskRepo
		f, err := os.Open(filepath.Clean(path))
		if err != nil {
			return false, err
		}
		c.RawAsset1 = f
		c.RawDirectory = &dir
		c.RawAsset1Filename = &filename
	case "rubygems":
		// Uploading files from the quick folder resulted in 500 issues
		if !strings.Contains(path, "/quick/") {
			c.Repository = localDiskRepo
			f, err := os.Open(filepath.Clean(path))
			if err != nil {
				return false, err
			}
			c.RubygemsAsset = f
		}
	case "yum":
		c.Repository = localDiskRepo
		f, err := os.Open(filepath.Clean(path))
		if err != nil {
			return false, err
		}
		c.YumAsset = f
		c.YumAssetFilename = &filename
	default:
		return false, nil
	}

	files := []*os.File{f, af.f2, af.f3, af.f4, af.f5, af.f6, af.f2, af.f7}
	if err := upload(c, client, path, files); err != nil {
		return false, err
	}

	return false, nil
}

func upload(c components.UploadComponentParams, client *client.Nexus3, path string, files []*os.File) error {
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

	for _, file := range files {
		if err := closeFileObjectIfNeeded(file); err != nil {
			return err
		}
	}

	return nil
}

func (n *Nexus3) uploadAndPrintRepoFormat(c *client.Nexus3, path, localDiskRepo, localDiskRepoHome, repoFormat string, skipErrors bool) error {
	identical, err := n.UploadSingleArtifact(c, path, localDiskRepo, localDiskRepoHome, repoFormat, skipErrors)
	if err != nil {
		uploaded, errRegex := regexp.MatchString("status 400", err.Error())
		if errRegex != nil {
			return err
		}
		if uploaded {
			log.Debugf("artifact: '%s' has already been uploaded", path)
			return nil
		}

		errString := fmt.Errorf("could not upload artifact: '%s', err: '%w'", path, err)
		if n.SkipErrors {
			log.Error(errString)
		} else {
			return errString
		}
	} else {
		artifacts.PrintType(repoFormat)
	}
	if identical {
		log.Debugf("checksum file: '%s' locally is identical compared to one in nexus", path)
		return nil
	}

	return nil
}

func (n *Nexus3) ReadLocalDirAndUploadArtifacts(localDiskRepoHome, localDiskRepo, repoFormat string) error {
	var wg sync.WaitGroup

	c, err := n.Nexus3.Client()
	if err != nil {
		return err
	}
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
				go func(cPreventDataRace *client.Nexus3, pathPreventDataRace, localDiskRepoPreventDataRace, localDiskRepoHomePreventDataRace, repoFormatPreventDataRace string, skipErrors bool) {
					defer wg.Done()

					if err := n.uploadAndPrintRepoFormat(cPreventDataRace, pathPreventDataRace, localDiskRepoPreventDataRace, localDiskRepoHomePreventDataRace, repoFormatPreventDataRace, skipErrors); err != nil {
						panic(err)
					}
				}(c, path, localDiskRepo, localDiskRepoHome, repoFormat, n.SkipErrors)
			}
			return nil
		}); err != nil {
		return err
	}
	wg.Wait()

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
	client, err := n.Nexus3.Client()
	if err != nil {
		panic(err)
	}
	r := repository_management.GetRepository2Params{RepositoryName: localDiskRepo}
	r.WithTimeout(time.Second * 30)
	resp, err := client.RepositoryManagement.GetRepository2(&r)
	if err != nil {
		log.Errorf("cannot determine version policy, repository: '%s'", localDiskRepo)
		return
	}
	vp := resp.Payload.Maven.VersionPolicy
	log.Tracef("VersionPolicy: '%s'", vp)

	if strings.EqualFold(vp, "snapshot") {
		s := snapshot.Nexus3{DownloadDirName: n.DownloadDirName, FQDN: n.FQDN, HTTPS: *n.HTTPS, Pass: n.Pass, RepoFormat: "maven2", RepoName: localDiskRepo, SkipErrors: n.SkipErrors, User: n.User}

		if err := s.Upload(); err != nil {
			uploaded, errRegex := regexp.MatchString("bad status: 400 Repository does not allow updating assets", err.Error())
			if errRegex != nil {
				panic(err)
			}
			if uploaded {
				log.Debugf("artifact from localDiskRepo: '%s' has been uploaded already", localDiskRepo)
				return
			}
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

	repoFormatAndType, err := n.repoFormatLocalDiskRepo(localDiskRepo)
	if err != nil {
		panic(err)
	}
	if repoFormatAndType.format == "" {
		log.Errorf("repoFormat not detected. Verify whether repo: '%s' resides in Nexus", localDiskRepo)
		return
	}

	log.Warnf("only uploads to 'hosted' repositories are supported. Current: '%v'", repoFormatAndType)
	if repoFormatAndType.repoType == "hosted" {
		if repoFormatAndType.format == "maven2" {
			log.Info("upload to snapshot repo")
			n.maven2SnapshotsUpload(localDiskRepo)
		}

		localDiskRepoHome := filepath.Join(n.DownloadDirName, localDiskRepo)
		if err := n.ReadLocalDirAndUploadArtifacts(localDiskRepoHome, localDiskRepo, repoFormatAndType.format); err != nil {
			panic(err)
		}
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
			go func(localDiskRepoPreventDataRace string) {
				defer wg.Done()

				n.uploadArtifactsSingleDir(localDiskRepoPreventDataRace)
			}(localDiskRepo)
		}
	}
	wg.Wait()

	return nil
}
