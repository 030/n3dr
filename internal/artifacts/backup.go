package artifacts

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/viper"

	"github.com/asaskevich/govalidator"
	log "github.com/sirupsen/logrus"
	"github.com/svenfuchs/jq"
	"github.com/thedevsaddam/gojsonq"
)

const (
	assetURI1   = "/service/rest/"
	assetURI2   = "/assets?repository="
	tokenErrMsg = "Token should be either a hexadecimal or \"null\" and not: "
	tmpDir      = "/tmp/n3dr"
)

func TempDownloadDir(downloadDirName string) (string, error) {
	if downloadDirName != "" {
		if err := os.MkdirAll(downloadDirName, os.ModePerm); err != nil {
			return "", nil
		}
		log.Infof("Download dir name: '%v'", downloadDirName)
		return downloadDirName, nil
	}
	if err := os.MkdirAll(tmpDir, os.ModePerm); err != nil {
		return "", nil
	}
	dname, err := os.MkdirTemp(tmpDir, "download")
	if err != nil {
		return "", err
	}
	log.Infof("Temp download dir name: '%v'", dname)
	return dname, nil
}

func (n Nexus3) downloadURL(token string) ([]byte, error) {
	assetURL := n.URL + assetURI1 + n.APIVersion + assetURI2 + n.Repository
	constructDownloadURL := assetURL
	if token != "null" {
		constructDownloadURL = assetURL + "&continuationToken=" + token
	}
	u, err := url.Parse(constructDownloadURL)
	if err != nil {
		return nil, err
	}
	log.Debug("DownloadURL: ", u)
	urlString := u.String()

	jsonResp, err := n.request(urlString)
	if err != nil {
		return nil, err
	}

	return jsonResp.bytes, nil
}

func (n Nexus3) continuationToken(token string) (string, error) {
	// The continuationToken should consists of 32 characters and should be a hexadecimal or "null"
	if !((govalidator.IsHexadecimal(token) && govalidator.StringLength(token, "32", "32")) || token == "null") {
		return "", errors.New(tokenErrMsg + token)
	}

	bodyBytes, err := n.downloadURL(token)
	if err != nil {
		return "", err
	}

	op, err := jq.Parse(".continuationToken")
	if err != nil {
		return "", err
	}

	value, err := op.Apply(bodyBytes)
	if err != nil {
		return "", err
	}

	tokenWithoutQuotes := strings.Trim(string(value), "\"")

	return tokenWithoutQuotes, nil
}

func createArtifact(d string, f string, content string, md5sum string) (errs []error) {
	ociBucketname := viper.GetString("ociBucket")
	Filename := d + "/" + f

	// Check if object exists
	objectExists := false
	if ociBucketname != "" {
		objectExists, err := findObject(ociBucketname, Filename, md5sum)
		if err != nil {
			errs = append(errs, err)
			return errs
		}

		if objectExists && viper.GetBool("removeLocalFile") {
			log.Debug("Object " + Filename + " already exist")
			return nil
		}
	}

	log.Debug("Create artifact: '" + d + "/" + f + "'")

	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	filename := filepath.Join(d, f)

	checksumDownloadedArtifact := ""
	if fileExists(filename) {
		dat, err := os.ReadFile(filepath.Clean(filename))
		if err != nil {
			errs = append(errs, err)
			return errs
		}

		checksumDownloadedArtifact = fmt.Sprintf("%x", sha512.Sum512(dat))
		if err != nil {
			errs = append(errs, err)
			return errs
		}
	}

	if checksumDownloadedArtifact == md5sum {
		log.Debug("Skipping as file already exists.")
	} else {
		log.Debug("Creating ", filename)
		file, err := os.Create(filepath.Clean(filename))
		if err != nil {
			errs = append(errs, err)
			return errs
		}
		defer func() {
			if err := file.Close(); err != nil {
				errs = append(errs, err)
			}
		}()

		_, err = file.WriteString(content)
		if err != nil {
			errs = append(errs, err)
			return errs
		}
	}

	if ociBucketname != "" && !objectExists {
		errs := ociBackup(ociBucketname, Filename)
		for _, err := range errs {
			if err != nil {
				errs = append(errs, err)
				return errs
			}
		}
	}
	return nil
}

func (n Nexus3) artifactName(url string) (string, string, error) {
	log.Debug("Validate whether: '" + url + "' is an URL")
	if !govalidator.IsURL(url) {
		return "", "", errors.New(url + " is not an URL")
	}

	re := regexp.MustCompile("^.*?/" + n.Repository + "/(.*)/(.+)$")
	if !re.MatchString(url) {
		return "", "", fmt.Errorf("URL: '%s' does not seem to contain a Maven artifact", url)
	}

	group := re.FindStringSubmatch(url)
	d := group[1]
	f := group[2]
	log.Debugf("ArtifactName directory: '%s' and file: '%s'"+d, f)

	return d, f, nil
}

func (n Nexus3) downloadArtifact(dir, url, md5 string) error {
	d, f, err := n.artifactName(url)
	if err != nil {
		return err
	}

	jsonResp, err := n.request(url)
	if err != nil {
		return err
	}

	if errs := createArtifact(filepath.Join(dir, n.Repository, d), f, jsonResp.strings, md5); err != nil {
		for _, err := range errs {
			return err
		}
	}
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (n Nexus3) continuationTokenRecursionChannel(cerr chan error, t, dir, regex string) error {
	token, err := n.continuationToken(t)
	if err != nil {
		return err
	}

	bytes, err := n.downloadURL(token)
	if err != nil {
		return err
	}
	json := string(bytes)
	jq := gojsonq.New().JSONString(json)
	log.Debugf("JQ: '%v'", jq)
	downloadURLAndChecksums := jq.From("items").Only("downloadUrl", "checksum.sha512")

	for _, downloadURLAndChecksum := range downloadURLAndChecksums.([]interface{}) {
		log.Debugf("downloadURLAndChecksum: '%v'", downloadURLAndChecksum)
		go func(downloadURLAndChecksum interface{}) {
			downloadURL := fmt.Sprint(downloadURLAndChecksum.(map[string]interface{})["downloadUrl"])
			md5 := fmt.Sprint(downloadURLAndChecksum.(map[string]interface{})["sha512"])

			log.Debugf("Only download artifacts that match the regex: '%s'. URL: '%s'", regex, downloadURL)
			r, err := regexp.Compile(regex)
			if err != nil {
				cerr <- err
				return
			}
			if r.MatchString(downloadURL) {
				// Exclude download of md5 and sha1 files as these are unavailable
				// unless the metadata.xml is opened first
				regexSha, _ := regexp.Compile("^.sha(1|256|512)$")

				// GH-134: archetype-catalog.xml should be skipped to prevent 'archetype-catalog.xml' does not seem to contain a Maven artifact' issues
				regexArchetypeCatalog := regexp.MustCompile("^.*?/" + n.Repository + "/archetype-catalog.xml$")
				if !(filepath.Ext(downloadURL) == ".md5" || regexSha.MatchString(filepath.Ext(downloadURL)) || regexArchetypeCatalog.MatchString(downloadURL)) {
					log.Debugf("DownloadURL: '%v'", downloadURL)
					if err := n.downloadArtifact(dir, downloadURL, md5); err != nil {
						cerr <- err
						return
					}
					fmt.Print("+")
				}
			}
			cerr <- nil
		}(downloadURLAndChecksum)
	}
	for range downloadURLAndChecksums.([]interface{}) {
		if err := <-cerr; err != nil {
			return err
		}
	}

	if token == "null" {
		return nil
	}
	if err := n.continuationTokenRecursionChannel(cerr, token, dir, regex); err != nil {
		return err
	}
	return nil
}

func (n Nexus3) StoreArtifactsOnDiskChannel(dir, regex string) error {
	log.Infof("Backing up: '%v'", n.Repository)
	cerr := make(chan error)
	defer close(cerr)
	if err := n.continuationTokenRecursionChannel(cerr, "null", dir, regex); err != nil {
		return err
	}
	return nil
}
