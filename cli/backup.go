package cli

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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
	dname, err := ioutil.TempDir(tmpDir, "download")
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

	bodyBytes, _, err := n.request(urlString)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
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

func (n Nexus3) continuationTokenRecursion(t string) ([]string, error) {
	token, err := n.continuationToken(t)
	if err != nil {
		return nil, err
	}
	if token == "null" {
		return []string{token}, nil
	}
	tokenSlice, err := n.continuationTokenRecursion(token)
	if err != nil {
		return nil, err
	}
	return append(tokenSlice, token), nil
}

func createArtifact(d string, f string, content string, md5sum string) error {
	ociBucketname := viper.GetString("ociBucket")
	Filename := d + "/" + f

	// Check if object exists
	objectExists := false
	if ociBucketname != "" {
		objectExists, err := findObject(ociBucketname, Filename, md5sum)

		if err != nil {
			return err
		}

		if objectExists && viper.GetBool("removeLocalFile") {
			log.Debug("Object " + Filename + " already exist")
			return nil
		}
	}

	log.Debug("Create artifact: '" + d + "/" + f + "'")

	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		return err
	}

	filename := filepath.Join(d, f)

	md5sumLocal := ""
	if fileExists(filename) {
		md5sumLocal, err = hashFileMD5(filename)
		if err != nil {
			return err
		}
	}

	if md5sumLocal == md5sum {
		log.Debug("Skipping as file already exists.")
	} else {
		log.Debug("Creating ", filename)
		file, err := os.Create(filename)
		if err != nil {
			return err
		}

		_, err = file.WriteString(content)
		defer file.Close()
		if err != nil {
			return err
		}
	}

	if ociBucketname != "" && !objectExists {
		err := ociBackup(ociBucketname, Filename)
		if err != nil {
			return err
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

	_, bodyString, err := n.request(url)
	if err != nil {
		return err
	}

	if err := createArtifact(filepath.Join(dir, n.Repository, d), f, bodyString, md5); err != nil {
		return err
	}
	return nil
}

// HashFileMD5 returns MD5 checksum of a file
func hashFileMD5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
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
	downloadURLAndMd5s := jq.From("items").Only("downloadUrl", "checksum.md5")

	for _, downloadURLAndMd5 := range downloadURLAndMd5s.([]interface{}) {
		log.Debugf("downloadURLAndMd5: '%v'", downloadURLAndMd5)
		go func(downloadURLAndMd5 interface{}) {
			downloadURL := fmt.Sprint(downloadURLAndMd5.(map[string]interface{})["downloadUrl"])
			md5 := fmt.Sprint(downloadURLAndMd5.(map[string]interface{})["md5"])

			log.Debugf("Only download artifacts that match the regex: '%s'. URL: '%s'", regex, downloadURL)
			r, err := regexp.Compile(regex)
			if err != nil {
				cerr <- err
				return
			}
			if r.MatchString(downloadURL) {
				// Exclude download of md5 and sha1 files as these are unavailable
				// unless the metadata.xml is opened first
				if !(filepath.Ext(downloadURL) == ".md5" || filepath.Ext(downloadURL) == ".sha1") {
					log.Debugf("DownloadURL: '%v'", downloadURL)
					if err := n.downloadArtifact(dir, downloadURL, md5); err != nil {
						cerr <- err
						return
					}
					fmt.Print("+")
				}
			}
			cerr <- nil
		}(downloadURLAndMd5)
	}
	for range downloadURLAndMd5s.([]interface{}) {
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
