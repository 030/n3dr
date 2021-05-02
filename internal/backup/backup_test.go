// package backup

// import (
// 	"io/ioutil"
// 	"os"
// 	"os/exec"
// 	"strconv"
// 	"testing"

// 	mp "github.com/030/go-multipart/utils"
// 	"github.com/030/mij"
// 	"github.com/030/n3dr/internal/config"
// 	"github.com/levigross/grequests"
// 	"github.com/stretchr/testify/assert"

// 	log "github.com/sirupsen/logrus"
// )

// const (
// 	nexus3Version      = "3.30.1"
// 	nexus3ExternalPort = 9998
// 	nexus3BrowserURI   = "/service/rest/repository/browse/"
// 	nexus3TestDir      = "/tmp/test-n3dr"
// 	npmRepositoryName  = "REPO_NAME_HOSTED_NPM"
// )

// var npmFiles = []string{"test-one", "test-two", "test-three"}

// var d = mij.DockerImage{
// 	Name:                     "sonatype/nexus3",
// 	PortExternal:             nexus3ExternalPort,
// 	PortInternal:             8081,
// 	Version:                  nexus3Version,
// 	ContainerName:            "nexus-backup-test",
// 	LogFile:                  "/nexus-data/log/nexus.log",
// 	LogFileStringHealthCheck: "Started Sonatype Nexus OSS",
// }
// var n = Nexus3{Endpoint: "http://localhost:" + strconv.Itoa(d.PortExternal), Password: "admin123", Username: "admin", BaseDir: nexus3TestDir, Regex: ".*"}

// func TestMain(m *testing.M) {
// 	setup(&d)
// 	code := m.Run()
// 	shutdown(&d)
// 	os.Exit(code)
// }

// type mijDockerImage struct {
// 	*mij.DockerImage
// }

// func (m *mijDockerImage) initialNexusAdminPassword() (string, error) {
// 	b, err := exec.Command("bash", "-c", "docker exec -i "+m.ContainerName+" cat /nexus-data/admin.password").CombinedOutput()
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(b), nil
// }

// func setup(m *mij.DockerImage) {
// 	m.Run()

// 	b := &mijDockerImage{m}
// 	pw, err := b.initialNexusAdminPassword()
// 	if err != nil {
// 		log.Error(err)
// 	}

// 	n.Password = pw

// 	d1 := []byte(pw)
// 	if err := ioutil.WriteFile("/tmp/bla", d1, 0644); err != nil {
// 		// log.Fatal(err)
// 	}

// 	r := config.Repository{Endpoint: n.Endpoint, Name: npmRepositoryName, Password: n.Password, Username: n.Username}
// 	if err := r.Create(); err != nil {
// 		log.Fatal(err)
// 	}

// 	// upload NPM testdata
// 	mpu := mp.Upload{URL: "http://localhost:9998/service/rest/v1/components?repository=" + npmRepositoryName, Username: n.Username, Password: pw}
// 	for _, file := range npmFiles {
// 		if err := mpu.MultipartUpload("npm.asset=@../../test/testdata/npm/" + file + "-1.0.0.tgz"); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

// func shutdown(m *mij.DockerImage) {
// 	m.Stop()
// }

// func TestRepositoryRawHTML(t *testing.T) {
// 	s, _ := n.repositoryRawHTML(n.Endpoint + nexus3BrowserURI + npmRepositoryName)
// 	assert.Equal(t, "\n<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <title>Index of /</title>\n    <meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\"/>\n\n\n    <!--[if lt IE 9]>\n    <script>(new Image).src=\"http://localhost:9998/favicon.ico?3.30.1-01\"</script>\n    <![endif]-->\n    <link rel=\"icon\" type=\"image/png\" href=\"http://localhost:9998/favicon-32x32.png?3.30.1-01\" sizes=\"32x32\">\n    <link rel=\"mask-icon\" href=\"http://localhost:9998/safari-pinned-tab.svg?3.30.1-01\" color=\"#5bbad5\">\n    <link rel=\"icon\" type=\"image/png\" href=\"http://localhost:9998/favicon-16x16.png?3.30.1-01\" sizes=\"16x16\">\n    <link rel=\"shortcut icon\" href=\"http://localhost:9998/favicon.ico?3.30.1-01\">\n    <meta name=\"msapplication-TileImage\" content=\"http://localhost:9998/mstile-144x144.png?3.30.1-01\">\n    <meta name=\"msapplication-TileColor\" content=\"#00a300\">\n\n    <link rel=\"stylesheet\" type=\"text/css\" href=\"http://localhost:9998/static/css/nexus-content.css?3.30.1-01\"/>\n</head>\n<body class=\"htmlIndex\">\n<h1>Index of /</h1>\n\n\n<table cellspacing=\"10\">\n    <tr>\n        <th align=\"left\">Name</th>\n        <th>Last Modified</th>\n        <th>Size</th>\n        <th>Description</th>\n    </tr>\n        <tr>\n            <td><a href=\"test-one/\">test-one</a></td>\n            <td>\n                    &nbsp;\n            </td>\n            <td align=\"right\">\n                    &nbsp;\n            </td>\n            <td></td>\n        </tr>\n        <tr>\n            <td><a href=\"test-three/\">test-three</a></td>\n            <td>\n                    &nbsp;\n            </td>\n            <td align=\"right\">\n                    &nbsp;\n            </td>\n            <td></td>\n        </tr>\n        <tr>\n            <td><a href=\"test-two/\">test-two</a></td>\n            <td>\n                    &nbsp;\n            </td>\n            <td align=\"right\">\n                    &nbsp;\n            </td>\n            <td></td>\n        </tr>\n</table>\n</body>\n</html>\n", s.String())
// }

// func TestRepositoryRawHTMLError(t *testing.T) {
// 	_, err := n.repositoryRawHTML("bla")
// 	assert.EqualError(t, err, "Get \"bla\": unsupported protocol scheme \"\"")

// 	_, err = n.repositoryRawHTML("http://localhost:" + strconv.Itoa(nexus3ExternalPort) + "/does-not-exist")
// 	assert.EqualError(t, err, "StatusCode URL: 'http://localhost:9998/does-not-exist' not OK, but: '404'. Enable debug mode to get the response")
// }

// func TestRepositoryDirectoriesAndFiles(t *testing.T) {
// 	dirs, _ := repositoryRawHTMLDirectoriesOrArtifacts("\n<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <title>Index of /</title>\n    <meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\"/>\n\n\n    <!--[if lt IE 9]>\n    <script>(new Image).src=\"http://localhost:9998/favicon.ico?3.30.1-01\"</script>\n    <![endif]-->\n    <link rel=\"icon\" type=\"image/png\" href=\"http://localhost:9998/favicon-32x32.png?3.30.1-01\" sizes=\"32x32\">\n    <link rel=\"mask-icon\" href=\"http://localhost:9998/safari-pinned-tab.svg?3.30.1-01\" color=\"#5bbad5\">\n    <link rel=\"icon\" type=\"image/png\" href=\"http://localhost:9998/favicon-16x16.png?3.30.1-01\" sizes=\"16x16\">\n    <link rel=\"shortcut icon\" href=\"http://localhost:9998/favicon.ico?3.30.1-01\">\n    <meta name=\"msapplication-TileImage\" content=\"http://localhost:9998/mstile-144x144.png?3.30.1-01\">\n    <meta name=\"msapplication-TileColor\" content=\"#00a300\">\n\n    <link rel=\"stylesheet\" type=\"text/css\" href=\"http://localhost:9998/static/css/nexus-content.css?3.30.1-01\"/>\n</head>\n<body class=\"htmlIndex\">\n<h1>Index of /</h1>\n\n\n<table cellspacing=\"10\">\n    <tr>\n        <th align=\"left\">Name</th>\n        <th>Last Modified</th>\n        <th>Size</th>\n        <th>Description</th>\n    </tr>\n        <tr>\n            <td><a href=\"%40babel/\">@babel</a></td>\n            <td>\n                    &nbsp;\n            </td>\n            <td align=\"right\">\n                    &nbsp;\n            </td>\n            <td></td>\n        </tr>\n</table>\n</body>\n</html>\n")
// 	assert.Equal(t, 1, len(dirs))

// 	dirs, _ = repositoryRawHTMLDirectoriesOrArtifacts("\n<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <title>Index of /</title>\n    <meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\"/>\n\n\n    <!--[if lt IE 9]>\n    <script>(new Image).src=\"http://localhost:9998/favicon.ico?3.30.1-01\"</script>\n    <![endif]-->\n    <link rel=\"icon\" type=\"image/png\" href=\"http://localhost:9998/favicon-32x32.png?3.30.1-01\" sizes=\"32x32\">\n    <link rel=\"mask-icon\" href=\"http://localhost:9998/safari-pinned-tab.svg?3.30.1-01\" color=\"#5bbad5\">\n    <link rel=\"icon\" type=\"image/png\" href=\"http://localhost:9998/favicon-16x16.png?3.30.1-01\" sizes=\"16x16\">\n    <link rel=\"shortcut icon\" href=\"http://localhost:9998/favicon.ico?3.30.1-01\">\n    <meta name=\"msapplication-TileImage\" content=\"http://localhost:9998/mstile-144x144.png?3.30.1-01\">\n    <meta name=\"msapplication-TileColor\" content=\"#00a300\">\n\n    <link rel=\"stylesheet\" type=\"text/css\" href=\"http://localhost:9998/static/css/nexus-content.css?3.30.1-01\"/>\n</head>\n<body class=\"htmlIndex\">\n<h1>Index of /</h1>\n\n\n<table cellspacing=\"10\">\n    <tr>\n        <th align=\"left\">Name</th>\n        <th>Last Modified</th>\n        <th>Size</th>\n        <th>Description</th>\n    </tr>\n        <tr>\n            <td><a href=\"test-one/\">test-one</a></td>\n            <td>\n                    &nbsp;\n            </td>\n            <td align=\"right\">\n                    &nbsp;\n            </td>\n            <td></td>\n        </tr>\n        <tr>\n            <td><a href=\"test-three/\">test-three</a></td>\n            <td>\n                    &nbsp;\n            </td>\n            <td align=\"right\">\n                    &nbsp;\n            </td>\n            <td></td>\n        </tr>\n        <tr>\n            <td><a href=\"test-two/\">test-two</a></td>\n            <td>\n                    &nbsp;\n            </td>\n            <td align=\"right\">\n                    &nbsp;\n            </td>\n            <td></td>\n        </tr>\n</table>\n</body>\n</html>\n")
// 	assert.Equal(t, 3, len(dirs))
// }
// func TestRepositoryDirectoriesAndFilesError(t *testing.T) {
// 	_, err := repositoryRawHTMLDirectoriesOrArtifacts("html>")
// 	assert.EqualError(t, err, "did not find any directories or artifacts; directoriesAndArtifactsSize: '0'")

// 	_, err = repositoryRawHTMLDirectoriesOrArtifacts("<html></html>")
// 	assert.EqualError(t, err, "did not find any directories or artifacts; directoriesAndArtifactsSize: '0'")
// }

// func TestDownload(t *testing.T) {
// 	var expectedLabel nexusArtifactLabel = "*"
// 	label, err := n.download("http://localhost:9998/repository/REPO_NAME_HOSTED_NPM/test-one/-/test-one-1.0.0.tgz")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	assert.Equal(t, expectedLabel, label)
// }
// func TestDownloadError(t *testing.T) {
// 	_, err := n.download("http://localhost:9998/repository/REPO_NAME_HOSTED_NPM/test-one/-/does-not-exist-1.0.0.tgz")
// 	assert.EqualError(t, err, "StatusCode URL: 'http://localhost:9998/repository/REPO_NAME_HOSTED_NPM/test-one/-/does-not-exist-1.0.0.tgz' not OK, but: '404'. Enable debug mode to get the response")

// 	n.BaseDir = "/tmp-does-not-exist/n3dr"
// 	_, err = n.download("http://localhost:9998/repository/REPO_NAME_HOSTED_NPM/test-one/-/test-one-1.0.0.tgz")
// 	assert.EqualError(t, err, "mkdir /tmp-does-not-exist: permission denied")
// }

// func TestDownloadAndPrintLabel(t *testing.T) {
// 	n.Regex = "does-not-match"
// 	var expectedLabel nexusArtifactLabel = ""
// 	label, _ := n.downloadAndPrintLabel("/tmp/bladibla/boo", "http://localhost:9998/repository/REPO_NAME_HOSTED_NPM/test-one/-/does-not-exist-1.0.0.tgz", &grequests.Response{Ok: false})
// 	assert.Equal(t, expectedLabel, label)
// }
// func TestDownloadAndPrintLabelError(t *testing.T) {
// 	n.Regex = ".*"
// 	_, err := n.downloadAndPrintLabel("/tmp/bladibla/boo", "http://localhost:9998/repository/REPO_NAME_HOSTED_NPM/test-one/-/does-not-exist-1.0.0.tgz", &grequests.Response{Ok: false})
// 	assert.EqualError(t, err, "open /tmp/bladibla/boo: no such file or directory")
// }

// func TestDownloadURL(t *testing.T) {
// 	n.BaseDir = nexus3TestDir
// 	n.Repository = npmRepositoryName
// 	err := n.AllArtifacts()
// 	assert.Nil(t, err)
// }
