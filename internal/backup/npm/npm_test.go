package npm

// import (
// 	"crypto/sha512"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"regexp"
// 	"time"

// 	"github.com/cavaliercoder/grab"
// 	"github.com/levigross/grequests"
// 	log "github.com/sirupsen/logrus"
// 	"github.com/tidwall/gjson"
// )

// type Nexus3 struct {
// 	BaseDir, Endpoint, Password, Repository, Regex, Username string
// }

// type continuationToken string

// const (
// 	componentsRepositoryURI = "/service/rest/v1/components?repository="
// )

// func (n *Nexus3) componentsRepositoryJSON(ct continuationToken) (string, error) {
// 	url := n.Endpoint + componentsRepositoryURI + n.Repository
// 	if ct != "" {
// 		url = url + "&continuationToken=" + string(ct)
// 	}

// 	resp, err := grequests.Get(url, &grequests.RequestOptions{
// 		Auth: []string{n.Username, n.Password}})
// 	if err != nil {
// 		return "", err
// 	}

// 	statusCode := resp.StatusCode
// 	log.Debugf("URL: '%v'. StatusCode: '%v'. Response: '%s'",
// 		url, statusCode, resp.String())
// 	if statusCode != http.StatusOK {
// 		return "",
// 			fmt.Errorf("statusCode URL: '%s' not OK, but: '%d'. "+
// 				"Enable debug mode to get the response",
// 				url, statusCode)
// 	}

// 	responseString := resp.String()
// 	if responseString == "" {
// 		return "", fmt.Errorf("response should not be empty. Actual: '%v'",
// 			responseString)
// 	}

// 	return responseString, nil
// }

// func (n *Nexus3) repositoryJSONAssets(ct continuationToken) (string, error) {
// 	log.Info("CT1: ", ct)
// 	json, err := n.componentsRepositoryJSON(ct)
// 	if err != nil {
// 		return "", err
// 	}

// 	// var checksum512 string
// 	assets := gjson.Get(json, "items.#.assets")
// 	log.Info("CT1a ", assets)
// 	for _, asset := range assets.Array() {
// 		var downloadURL string
// 		if value := gjson.Get(asset.String(), "#.downloadUrl"); !value.Exists() {
// 			return "", fmt.Errorf("downloadUrl does not exist in json")
// 		} else {
// 			log.Info("CT1b ", value.String())
// 			re := regexp.MustCompile(`^\[\"(.*)\"\]$`)
// 			match := re.FindStringSubmatch(value.String())
// 			downloadURL = match[1]
// 			log.Info("CT1c: ", downloadURL)
// 		}

// 		var checksum512 string
// 		if value := gjson.Get(asset.String(), "#.checksum.sha512"); !value.Exists() {
// 			return "", fmt.Errorf("512checksum does not exist in json")
// 		} else {
// 			log.Info("CT1d ", value.String())
// 			re := regexp.MustCompile(`^\[\"(.*)\"\]$`)
// 			match := re.FindStringSubmatch(value.String())
// 			checksum512 = match[1]
// 			log.Info("CT1e: ", checksum512)
// 		}

// 		filePathDir := filepath.Join(n.BaseDir, n.Repository)
// 		if err := os.MkdirAll(filePathDir, os.ModePerm); err != nil {
// 			return "", err
// 		}

// 		client := grab.NewClient()
// 		req, err := grab.NewRequest(filePathDir, downloadURL)
// 		if err != nil {
// 			return "", err
// 		}
// 		req.HTTPRequest.SetBasicAuth(n.Username, n.Password)
// 		resp := client.Do(req)
// 		// log.Info("File ", resp.Filename)

// 		fmt.Printf("  %v\n", resp.HTTPResponse.Status)

// 		// start UI loop
// 		t := time.NewTicker(500 * time.Millisecond)
// 		defer t.Stop()

// 	Loop:
// 		for {
// 			select {
// 			case <-t.C:
// 				fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
// 					resp.BytesComplete(),
// 					resp.Size,
// 					100*resp.Progress())

// 			case <-resp.Done:
// 				// download is complete
// 				break Loop
// 			}
// 		}

// 		// check for errors
// 		if err := resp.Err(); err != nil {
// 			fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
// 			os.Exit(1)
// 		}

// 		fmt.Printf("Download saved to ./%v \n", resp.Filename)

// 		// dat, err := ioutil.ReadFile(resp.Filename)
// 		// if err != nil {
// 		// 	return "", err
// 		// }
// 		// checksum512OnDisk := fmt.Sprintf("%x", sha512.Sum512(dat))
// 		f, err := os.Open(resp.Filename)
// 		if err != nil {
// 			return "", err
// 		}
// 		defer f.Close()
// 		h := sha512.New()
// 		if _, err := io.Copy(h, f); err != nil {
// 			return "", err
// 		}
// 		checksum512OnDisk := fmt.Sprintf("%x", h.Sum(nil))
// 		if checksum512 != checksum512OnDisk {
// 			return "", fmt.Errorf("512checksum mismatch on disk: '%v' vs. '%v'. File: '%s'", checksum512, checksum512OnDisk, resp.Filename)
// 		}
// 	}

// 	ct, err = continuationTokenInJSON(json)
// 	if err != nil {
// 		return "", err
// 	}
// 	log.Info("CT2: ", ct)

// 	if ct == "" {
// 		log.Info("CT2a: ")
// 		return "done", nil
// 	}

// 	log.Info("CT3: ", ct)
// 	return n.repositoryJSONAssets(ct)
// }

// func continuationTokenInJSON(json string) (continuationToken, error) {
// 	if value := gjson.Get(json, "continuationToken"); !value.Exists() {
// 		log.Debugf("JSON: '%s'", json)
// 		return "", fmt.Errorf("continuationToken does not exist in json")
// 	} else {
// 		return continuationToken(value.String()), nil
// 	}
// }
import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"

	mp "github.com/030/go-multipart/utils"
	"github.com/030/mij"
	"github.com/030/n3dr/internal/config"
	log "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

type mijDockerImage struct {
	*mij.DockerImage
}

const (
	nexus3Version      = "3.28.1"
	nexus3ExternalPort = 9996
	nexus3BrowserURI   = "/service/rest/repository/browse/"
	nexus3TestDir      = "/tmp/test-n3dr/npm"
	repositoryName     = "REPO_NAME_HOSTED_NPM"
)

var (
	d = mij.DockerImage{
		Name:                     "sonatype/nexus3",
		PortExternal:             nexus3ExternalPort,
		PortInternal:             8081,
		Version:                  nexus3Version,
		ContainerName:            "n3dr-backup-test-npm",
		LogFile:                  "/nexus-data/log/nexus.log",
		LogFileStringHealthCheck: "Started Sonatype Nexus OSS",
	}
	files = []string{"test-one", "test-two", "test-three"}
	n     = Nexus3{Endpoint: "http://localhost:" + strconv.Itoa(d.PortExternal), Username: "admin", BaseDir: nexus3TestDir, Regex: ".*", Repository: repositoryName}
)

func TestMain(m *testing.M) {
	setup(&d)
	code := m.Run()
	shutdown(&d)
	os.Exit(code)
}

func (m *mijDockerImage) initialNexusAdminPassword() (string, error) {
	log.Infof("Container name: '%s'", m.ContainerName)
	b, err := exec.Command("bash", "-c", "docker exec -i "+m.ContainerName+" cat /nexus-data/admin.password").CombinedOutput()
	initialAdminPass := string(b)
	if err != nil {
		return "", fmt.Errorf("cannot get initialAdminPass: '%s'", initialAdminPass)
	}
	return initialAdminPass, nil
}

func setup(m *mij.DockerImage) {
	// m.Run()

	mdi := &mijDockerImage{m}
	pw, err := mdi.initialNexusAdminPassword()
	if err != nil {
		log.Fatalf("cannot get initial nexus3 admin pass. Error: '%v'", err)
	}
	n.Password = pw

	u := config.User{UserID: "test-n3dr", LastName: "test-n3dr", EmailAddress: "test-n3dr@n3dr-test", Password: "some-pass-test-n3dr", Nexus3: config.Nexus3{Endpoint: n.Endpoint, Password: n.Password, Username: n.Username}}
	u.Create()

	r := config.Repository{Name: repositoryName, Nexus3: config.Nexus3{Endpoint: n.Endpoint, Password: n.Password, Username: n.Username}}
	r.Create()

	// if err := n.uploadTestData(); err != nil {
	// 	log.Fatal(err)
	// }
}

func shutdown(m *mij.DockerImage) {
	// m.Stop()
}

func (n *Nexus3) uploadTestData() error {
	mpu := mp.Upload{URL: "http://localhost:" + strconv.Itoa(nexus3ExternalPort) + "/service/rest/v1/components?repository=" + repositoryName, Username: n.Username, Password: n.Password}
	for _, file := range files {
		if err := mpu.MultipartUpload("npm.asset=@../../../test/testdata/npm/" + file + "-1.0.0.tgz"); err != nil {
			return err
		}
	}
	return nil
}

// func TestComponentsRepositoryJSON(t *testing.T) {
// 	n.Repository = repositoryName
// 	s, err := n.componentsRepositoryJSON("")
// 	exp := `test-one\/-\/test-one-1.0.0.tgz`
// 	assert.Nil(t, err)
// 	assert.Regexp(t, exp, s)
// }

func TestRepositoryJSONAssets(t *testing.T) {
	err := n.AllArtifacts()
	assert.Nil(t, err)

	i := 0
	files, err := ioutil.ReadDir(filepath.Join(n.BaseDir, n.Repository))
	assert.Nil(t, err)
	for _, file := range files {
		if !file.IsDir() {
			i++
		}
	}
	assert.Equal(t, 3, i)
}

// func TestContinuationTokenInJSON(t *testing.T) {
// 	json := `{
// 		"items": [ {
// 		} ],
// 		"continuationToken": "boo"
// 	}`
// 	s, err := backup.ContinuationTokenInJSON(json)
// 	assert.Nil(t, err)
// 	assert.Equal(t, backup.ContinuationToken("boo"), s)

// 	json = `{
// 		"items": [ {
// 		} ],
// 		"continuationToken": null
// 	}`
// 	s, err = backup.ContinuationTokenInJSON(json)
// 	assert.Nil(t, err)
// 	assert.Equal(t, backup.ContinuationToken(""), s)
// }

// func TestContinuationTokenErrors(t *testing.T) {
// 	jsons := []string{
// 		`{
// 		"items": [ {
// 		} ],
// 		"continuationTokenDoesNotExist" : "something"
// 		}`,
// 		`{`,
// 		``,
// 		`abc`,
// 	}
// 	for _, json := range jsons {
// 		s, err := backup.ContinuationTokenInJSON(json)
// 		assert.EqualError(t, err, "continuationToken does not exist in json")
// 		assert.Equal(t, backup.ContinuationToken(""), s)
// 	}
// }
