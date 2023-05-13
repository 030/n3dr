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
	goOpenApiRuntime "github.com/go-openapi/runtime"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

type Nexus3 struct {
	*connection.Nexus3
}

type mavenParts struct {
	artifact, classifier, ext, version string
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

func processMaven2Asset(filePath string, asset *goOpenApiRuntime.NamedReadCloser, ext, classifier **string, logFields log.Fields, skipErrors bool, files []*os.File) ([]*os.File, error) {
	if _, err := os.Stat(filePath); err == nil {
		f, err := os.Open(filepath.Clean(filePath))
		if err != nil {
			return nil, err
		}
		files = append(files, f)

		mp, err := maven(filePath, skipErrors)
		if err != nil {
			return nil, err
		}
		*asset = f
		*ext = &mp.ext
		*classifier = &mp.classifier

		log.WithFields(logFields).Trace("Maven2 asset")
	}

	return files, nil
}

func mavenJarAndOtherExtensions(c *components.UploadComponentParams, fileNameWithoutExtIncludingDir string, files []*os.File, skipErrors bool) ([]*os.File, error) {
	var err error

	assets := []struct {
		filePath   string
		mavenAsset *goOpenApiRuntime.NamedReadCloser
		extension  **string
		classifier **string
		logFields  log.Fields
	}{
		{fileNameWithoutExtIncludingDir + "-sources.jar", &c.Maven2Asset2, &c.Maven2Asset2Extension, &c.Maven2Asset2Classifier, log.Fields{"file": c.Maven2Asset2.Name(), "extension": *c.Maven2Asset2Extension, "classifier": *c.Maven2Asset2Classifier}},
		{fileNameWithoutExtIncludingDir + "-javadoc.jar", &c.Maven2Asset7, &c.Maven2Asset7Extension, &c.Maven2Asset7Classifier, log.Fields{"file": c.Maven2Asset7.Name(), "extension": *c.Maven2Asset7Extension, "classifier": *c.Maven2Asset7Classifier}},
		{fileNameWithoutExtIncludingDir + ".jar", &c.Maven2Asset3, &c.Maven2Asset3Extension, &c.Maven2Asset3Classifier, log.Fields{"file": c.Maven2Asset3.Name(), "extension": *c.Maven2Asset3Extension, "classifier": *c.Maven2Asset3Classifier}},
		{fileNameWithoutExtIncludingDir + ".war", &c.Maven2Asset4, &c.Maven2Asset4Extension, &c.Maven2Asset4Classifier, log.Fields{"file": c.Maven2Asset4.Name(), "extension": *c.Maven2Asset4Extension, "classifier": *c.Maven2Asset4Classifier}},
		{fileNameWithoutExtIncludingDir + ".zip", &c.Maven2Asset5, &c.Maven2Asset5Extension, &c.Maven2Asset5Classifier, log.Fields{"file": c.Maven2Asset5.Name(), "extension": *c.Maven2Asset5Extension, "classifier": *c.Maven2Asset5Classifier}},
		{fileNameWithoutExtIncludingDir + ".module", &c.Maven2Asset6, &c.Maven2Asset6Extension, &c.Maven2Asset6Classifier, log.Fields{"file": c.Maven2Asset6.Name(), "extension": *c.Maven2Asset6Extension, "classifier": *c.Maven2Asset6Classifier}},
	}

	for _, asset := range assets {
		files, err = processMaven2Asset(asset.filePath, asset.mavenAsset, asset.extension, asset.classifier, asset.logFields, skipErrors, files)
		if err != nil {
			return nil, err
		}
	}

	return files, nil
}

func removeExtension(fileName string) string {
	ext := filepath.Ext(fileName)
	re := regexp.MustCompile(`(-[a-zA-Z]+)?` + ext + `$`)
	result := re.ReplaceAllString(fileName, "")
	filename := filepath.Base(result)
	return filename
}

var checkedMavenFolders = []string{""}

func (n *Nexus3) compareChecksumLocalArtifactWithRemote(f, localDiskRepo, dir, filename string) (bool, error) {
	downloadedFileChecksum, err := artifacts.ChecksumLocalFile(f, "")
	if err != nil {
		return false, err
	}
	log.Info(f, downloadedFileChecksum)

	retryClient := retryablehttp.NewClient()
	retryClient.Logger = nil
	retryClient.RetryMax = 3
	standardClient := retryClient.StandardClient()

	u := fmt.Sprintf("%s://%s/repository/%s/%s/%s.sha512", func() string {
		if n.HTTPS {
			return "https"
		}
		return "http"
	}(), n.FQDN, localDiskRepo, dir, filename)
	log.Info("URL:             ", u)

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
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Info(bodyString)

	identical := false

	if bodyString == downloadedFileChecksum {
		identical = true
	}

	return identical, nil
}

func asset(c *components.UploadComponentParams, dir, filename, localDiskRepo, path, repoFormat string) (*os.File, error) {
	var err error
	var f *os.File
	c.Repository = localDiskRepo

	if !(repoFormat == "rubygems" && strings.Contains(path, "/quick/")) {
		f, err = os.Open(filepath.Clean(path))
		if err != nil {
			return nil, err
		}
	}

	switch repoFormat {
	case "apt":
		c.AptAsset = f
	case "npm":
		c.NpmAsset = f
	case "nuget":
		c.NugetAsset = f
	case "raw":
		c.RawAsset1 = f
		c.RawDirectory = &dir
		c.RawAsset1Filename = &filename
	case "rubygems":
		c.RubygemsAsset = f
	case "yum":
		c.YumAsset = f
		c.YumAssetFilename = &filename
	}

	return f, nil
}

func aaa(c *components.UploadComponentParams, dirPath, fileNameWithoutExtIncludingDir, filePathPom string, skipErrors bool, files []*os.File) (*components.UploadComponentParams, []*os.File, error) {
	var err error

	log.Debugf("folder: '%s' contains a pom file: '%s'", dirPath, filePathPom)

	files, err = mavenJarAndOtherExtensions(c, fileNameWithoutExtIncludingDir, files, skipErrors)
	if err != nil {
		return nil, nil, err
	}

	f, err := os.Open(filepath.Clean(filePathPom))
	if err != nil {
		return nil, nil, err
	}
	files = append(files, f)

	c.Maven2Asset1 = f
	ext1 := "pom"
	c.Maven2Asset1Extension = &ext1

	return c, files, nil
}

func bbb(c *components.UploadComponentParams, dirPath, fileNameWithoutExtIncludingDir, filePathPom, localDiskRepoHome, path string, skipErrors bool, files []*os.File) (*components.UploadComponentParams, []*os.File, error) {
	var err error

	log.Debugf("folder: '%s' does not contain a pom file: '%s'", dirPath, filePathPom)

	files, err = mavenJarAndOtherExtensions(c, fileNameWithoutExtIncludingDir, files, skipErrors)
	if err != nil {
		return nil, nil, err
	}

	//
	mp, err := maven(path, skipErrors)
	if err != nil {
		return nil, nil, err
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
		return nil, nil, fmt.Errorf("groupID should not be empty, path: '%s' and regex: '%s'", path, regex)
	}
	c.Maven2GroupID = &groupID
	generatePOM := true
	c.Maven2GeneratePom = &generatePOM

	return c, files, nil
}

func logging(c *components.UploadComponentParams) {
	maven2Assets := make([]string, 7)
	for i, asset := range []goOpenApiRuntime.NamedReadCloser{c.Maven2Asset1, c.Maven2Asset2, c.Maven2Asset3, c.Maven2Asset4, c.Maven2Asset5, c.Maven2Asset6, c.Maven2Asset7} {
		if asset != nil {
			maven2Assets[i] = asset.Name()
		} else {
			maven2Assets[i] = "empty"
		}
	}
	log.WithFields(log.Fields{
		"1": maven2Assets[0],
		"2": maven2Assets[1],
		"3": maven2Assets[2],
		"4": maven2Assets[3],
		"5": maven2Assets[4],
		"6": maven2Assets[5],
		"7": maven2Assets[6],
	}).Debug("Maven2 asset upload")
}

func (n *Nexus3) UploadSingleArtifact(client *client.Nexus3, path, localDiskRepo, localDiskRepoHome, repoFormat string, skipErrors bool) error {
	dir := strings.Replace(filepath.Dir(path), localDiskRepoHome+"/", "", -1)
	filename := filepath.Base(path)

	if identical, _ := n.compareChecksumLocalArtifactWithRemote(filepath.Clean(path), localDiskRepo, dir, filename); identical {
		log.Info("Already uploaded")
		return nil
	}

	var err error
	var f *os.File
	var files []*os.File
	c := components.UploadComponentParams{}
	switch repoFormat {
	case "apt":
		f, err = asset(&c, "", "", localDiskRepo, path, repoFormat)
		if err != nil {
			return err
		}
	case "maven2":
		dirPath := filepath.Dir(path)
		fileNameWithoutExt := removeExtension(path)
		fileNameWithoutExtIncludingDir := filepath.Join(dirPath, fileNameWithoutExt)
		filePathPom := fileNameWithoutExtIncludingDir + ".pom"

		if slices.Contains(checkedMavenFolders, dirPath) {
			return nil
		}

		c.Repository = localDiskRepo

		if _, err := os.Stat(filePathPom); err == nil {
			_, files, err = aaa(&c, dirPath, fileNameWithoutExtIncludingDir, filePathPom, skipErrors, files)
			if err != nil {
				return err
			}
		} else {
			_, files, err = bbb(&c, dirPath, fileNameWithoutExtIncludingDir, filePathPom, localDiskRepoHome, path, skipErrors, files)
			if err != nil {
				return err
			}
		}

		logging(&c)

		checkedMavenFolders = append(checkedMavenFolders, dirPath)
		checkedMavenFolders = lo.Uniq(checkedMavenFolders)
	case "npm":
		f, err = asset(&c, "", "", localDiskRepo, path, repoFormat)
		if err != nil {
			return err
		}
	case "nuget":
		f, err = asset(&c, "", "", localDiskRepo, path, repoFormat)
		if err != nil {
			return err
		}
	case "raw":
		f, err = asset(&c, dir, filename, localDiskRepo, path, repoFormat)
		if err != nil {
			return err
		}

	case "rubygems":
		f, err = asset(&c, "", "", localDiskRepo, path, repoFormat)
		if err != nil {
			return err
		}
	case "yum":
		f, err = asset(&c, "", filename, localDiskRepo, path, repoFormat)
		if err != nil {
			return err
		}
	default:
		return nil
	}

	files = append(files, f)
	if err := upload(c, client, path, files); err != nil {
		return err
	}

	return nil
}

func upload(c components.UploadComponentParams, client *client.Nexus3, path string, files []*os.File) error {
	if reflect.ValueOf(c).IsZero() {
		log.Debug("no files to be uploaded")
		return nil
	}
	c.WithTimeout(time.Second * 600)
	if err := client.Components.UploadComponent(&c); err != nil {
		log.Tracef("----------------------------------------------------------------------------------err: '%v' while uploading file: '%s'", err, path)
		statusCode, uploadStatusErr := uploadStatus(err)
		if uploadStatusErr != nil {
			log.Error(path)
			return uploadStatusErr
		}
		if statusCode == 204 {
			log.Debugf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>artifact: '%v' has been uploaded", path)
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

func (n *Nexus3) ReadLocalDirAndUploadArtifacts(localDiskRepoHome, localDiskRepo, repoFormat string) error {
	var wg sync.WaitGroup

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
				wg.Add(1)
				go func(path, localDiskRepo, localDiskRepoHome, repoFormat string, skipErrors bool) {
					defer wg.Done()

					if err := n.UploadSingleArtifact(c, path, localDiskRepo, localDiskRepoHome, repoFormat, skipErrors); err != nil {
						uploaded, errRegex := regexp.MatchString("status 400", err.Error())
						if errRegex != nil {
							panic(err)
						}
						if uploaded {
							log.Debugf("artifact: '%s' has already been uploaded", path)
							return
						}

						errString := fmt.Errorf("could not upload artifact: '%s', err: '%w'", path, err)
						if n.SkipErrors {
							log.Error(errString)
						} else {
							panic(errString)
						}
					} else {
						artifacts.PrintType(repoFormat)
					}
				}(path, localDiskRepo, localDiskRepoHome, repoFormat, n.SkipErrors)
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
	client := n.Nexus3.Client()
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
