package artifacts

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	mp "github.com/030/go-multipart/utils"
	log "github.com/sirupsen/logrus"
)

var foldersWithPOM strings.Builder
var foldersWithPOMStringSlice []string
var artifactIndex int

// detectFoldersWithPOM checks whether there are folders with .pom files.
// Without them, maven artifacts cannout be published to nexus3.
func (n Nexus3) detectFoldersWithPOM(d string) error {
	err := filepath.Walk(d, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() && filepath.Ext(path) == ".pom" {
			log.Debug(path)
			foldersWithPOM.WriteString(filepath.Dir(path) + ",")
		}
		return nil
	})
	if err != nil {
		return err
	}

	foldersWithPOMString := strings.TrimSuffix(foldersWithPOM.String(), ",")
	if foldersWithPOMString == "" {
		return fmt.Errorf("no folders with .pom files detected. Please check whether the '%s' directory contains .pom files", d)
	}

	foldersWithPOMStringSlice = strings.Split(foldersWithPOMString, ",")
	return nil
}

func sbArtifact(sb *strings.Builder, path, ext, classifier string) error {
	log.Debugf("Path: '%s'. Ext: '%s'. Classifier: '%s'.", path, ext, classifier)

	_, err := fmt.Fprintf(sb, "maven2.asset%d=@%s,maven2.asset%d.extension=%s,", artifactIndex, path, artifactIndex, ext)
	if err != nil {
		return err
	}

	if len(classifier) > 0 {
		_, err := fmt.Fprintf(sb, "maven2.asset%d.classifier=%s,", artifactIndex, classifier)
		if err != nil {
			return err
		}
	}

	artifactIndex++
	return nil
}

func artifactTypeDetector(sb *strings.Builder, path string) error {
	var err error

	re := regexp.MustCompile(`^.*\/([\w\.-]+)-([\d\.]+)-?([\w-]+)?\.(\w+)$`)
	if re.Match([]byte(path)) {
		result := re.FindAllStringSubmatch(path, -1)
		log.Debugf("Artifact: '%v'", result[0][1])
		log.Debugf("Version: '%v'", result[0][2])
		classifier := result[0][3]
		ext := result[0][4]
		err = sbArtifact(sb, path, ext, classifier)
	} else {
		log.Warningf("'%v' not an artifact", path)
		// return nil to continue-on-error to ensure that subsequent artifacts
		// will be uploaded
	}

	return err
}

func (n Nexus3) multipartUpload(sb strings.Builder) error {
	multipartString := strings.TrimSuffix(sb.String(), ",")
	url := n.URL + "/service/rest/" + n.APIVersion + "/components?repository=" + n.Repository
	log.WithFields(log.Fields{
		"multipart": multipartString,
		"url":       url,
	}).Debug("URL")

	u := mp.Upload{URL: url, Username: n.User, Password: n.Pass}
	if err := u.MultipartUpload(multipartString); err != nil {
		return err
	}
	return nil
}

func pomDirs(p string) (strings.Builder, error) {
	var sb strings.Builder
	artifactIndex = 1

	err := filepath.Walk(p, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			if err = artifactTypeDetector(&sb, path); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return sb, err
	}
	return sb, nil
}

func (n *Nexus3) readFiles() ([]string, error) {
	file, err := os.Open(n.Repository)
	if err != nil {
		return nil, err
	}
	names, err := file.Readdirnames(0)
	if err != nil {
		return nil, err
	}
	return names, err
}

func (n *Nexus3) uploadFile(file string, req *http.Request) error {
	req.SetBasicAuth(n.User, os.ExpandEnv(n.Pass))
	req.Header.Set("Content-Type", "multipart/form-data")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		log.Infof("Upload of %v Ok. StatusCode: %v.", file, resp.StatusCode)
	} else {
		log.Error(resp)
		return fmt.Errorf("Upload of %v to %v failed. StatusCode: '%v'", file, n.URL, resp.StatusCode)
	}
	return nil
}

func (n *Nexus3) uploadMultipartFile(file *os.File, writer *multipart.Writer, req *http.Request, statusCreated int) error {
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(n.User, os.ExpandEnv(n.Pass))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// For some reason a 200 instead of 201 is returned if an NPM has been uploaded
	if resp.StatusCode == statusCreated {
		log.Infof("Upload of %v Ok. StatusCode: %v.", file, resp.StatusCode)
	} else {
		log.Error(resp)
		return fmt.Errorf("Upload of %v to %v failed. StatusCode: '%v'", file, n.URL, resp.StatusCode)
	}
	return nil
}

func (n *Nexus3) openFileAndUpload(file string) error {
	log.Infof("Uploading file %v to %v", file, n.URL)

	f, err := os.Open("./" + n.Repository + "/" + file)
	if err != nil {
		return err
	}
	defer f.Close()
	req, err := http.NewRequest("POST", n.URL+"/repository/"+n.Repository+"/", f)
	if err != nil {
		return err
	}
	if err := n.uploadFile(file, req); err != nil {
		return err
	}
	return nil
}

func (n *Nexus3) openMultipartFileAndUpload(f, httpMethod, uri string, statusCreated int) error {
	log.Infof("Uploading file %v to %v", f, n.URL)
	fileDir, _ := os.Getwd()
	fileName := f
	filePath := path.Join(fileDir, n.Repository, fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return err
	}

	if _, err := io.Copy(part, file); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	log.Infof("Upload Method: '%v', URL: '%v'", httpMethod, n.URL+uri)
	req, err := http.NewRequest(httpMethod, n.URL+uri, body)
	if err != nil {
		return err
	}

	if err := n.uploadMultipartFile(file, writer, req, statusCreated); err != nil {
		return err
	}
	return nil
}

func (n *Nexus3) readMultipartFilesAndUpload(httpMethod, uri string, statusCreated int) error {
	files, err := n.readFiles()
	if err != nil {
		return err
	}
	for _, file := range files {
		if err := n.openMultipartFileAndUpload(file, httpMethod, uri, statusCreated); err != nil {
			return err
		}
	}
	return nil
}

func (n *Nexus3) readFilesAndUpload() error {
	files, err := n.readFiles()
	if err != nil {
		return err
	}
	for _, file := range files {
		if err := n.openFileAndUpload(file); err != nil {
			return err
		}
	}
	return nil
}

func (n *Nexus3) readMavenFilesAndUpload() error {
	if err := n.detectFoldersWithPOM(n.Repository); err != nil {
		return err
	}
	for i, path := range foldersWithPOMStringSlice {
		log.Debug(strconv.Itoa(i) + " Detecting artifacts in folder '" + path + "'")
		sb, err := pomDirs(path)
		if err != nil {
			return err
		}
		log.Info(strconv.Itoa(i) + " Upload '" + sb.String() + "'")
		if err := n.multipartUpload(sb); err != nil {
			return err
		}
	}
	return nil
}

// Upload posts an artifact as a multipart to a specific nexus3 repository
func (n Nexus3) Upload() error {
	log.Infof("Uploading '%s'", n.ArtifactType)
	switch n.ArtifactType {
	case "apt":
		if err := n.readFilesAndUpload(); err != nil {
			return err
		}
	case "maven2":
		if err := n.readMavenFilesAndUpload(); err != nil {
			return err
		}
	case "npm":
		if err := n.readMultipartFilesAndUpload("POST", "/service/rest/internal/ui/upload/"+n.Repository, http.StatusOK); err != nil {
			return err
		}
	case "nuget":
		if err := n.readMultipartFilesAndUpload("PUT", "/repository/"+n.Repository+"/", http.StatusCreated); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Upload of '%s' is not supported", n.ArtifactType)
	}
	return nil
}
