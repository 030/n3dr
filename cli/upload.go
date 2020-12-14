package cli

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

func sbArtifact(sb *strings.Builder, path string, ext string, classifier string) error {
	log.Debug(ext + " found " + path)

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

	re := regexp.MustCompile(`^.*\/([\w-]+)-([\d.]+)-?([\w-]+)?\.(\w+)$`)
	if re.Match([]byte(path)) {
		result := re.FindAllStringSubmatch(path, -1)
		//artifact := result[0][1]
		//version := result[0][2]
		classifier := result[0][3]
		ext := result[0][4]
		err = sbArtifact(sb, path, ext, classifier)
	} else {
		log.Debug(path + " not an artifact")
		return nil
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
			err = artifactTypeDetector(&sb, path)
			if err != nil {
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

// Upload posts an artifact as a multipart to a specific nexus3 repository
func (n Nexus3) Upload() error {
	log.Infof("Uploading '%s'", n.ArtifactType)
	switch n.ArtifactType {
	case "apt":
		file, err := os.Open(n.Repository)
		if err != nil {
			return err
		}
		names, err := file.Readdirnames(0)
		if err != nil {
			return err
		}

		for _, name := range names {
			log.Infof("Uploading file %v to %v", name, n.URL)

			f, err := os.Open("./" + n.Repository + "/" + name)
			if err != nil {
				return err
			}
			defer f.Close()
			req, err := http.NewRequest("POST", n.URL+"/repository/"+n.Repository+"/", f)
			if err != nil {
				return err
			}
			req.SetBasicAuth(n.User, os.ExpandEnv(n.Pass))
			req.Header.Set("Content-Type", "multipart/form-data")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusCreated {
				log.Infof("Upload of %v Ok. StatusCode: %v.", name, resp.StatusCode)
			} else {
				log.Error(resp)
				return fmt.Errorf("Upload of %v to %v failed. StatusCode: '%v'", name, n.URL, resp.StatusCode)
			}
		}
	case "maven2":
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
	case "npm":
		file, err := os.Open(n.Repository)
		if err != nil {
			return err
		}
		names, err := file.Readdirnames(0)
		if err != nil {
			return err
		}
		fmt.Println(names)

		for _, name := range names {
			fmt.Println("./" + n.Repository + "/" + name)
			fileDir, _ := os.Getwd()
			fileName := name
			filePath := path.Join(fileDir, n.Repository, fileName)

			file, _ := os.Open(filePath)
			defer file.Close()

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
			io.Copy(part, file)
			writer.Close()

			fmt.Println(n.URL + "/service/rest/internal/ui/upload/" + n.Repository)
			req, err := http.NewRequest("POST", n.URL+"/service/rest/internal/ui/upload/"+n.Repository, body)
			if err != nil {
				return err
			}
			req.Header.Add("Content-Type", writer.FormDataContentType())
			req.SetBasicAuth(n.User, os.ExpandEnv(n.Pass))

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			// For some reason a 200 instead of 201 is returned if an NPM has been uploaded
			if resp.StatusCode == http.StatusOK {
				log.Infof("Upload of %v Ok. StatusCode: %v.", name, resp.StatusCode)
			} else {
				log.Error(resp)
				return fmt.Errorf("Upload of %v to %v failed. StatusCode: '%v'", name, n.URL, resp.StatusCode)
			}
		}
	case "nuget":
		file, err := os.Open(n.Repository)
		if err != nil {
			return err
		}
		names, err := file.Readdirnames(0)
		if err != nil {
			return err
		}
		fmt.Println(names)

		for _, name := range names {
			fmt.Println("./" + n.Repository + "/" + name)
			fileDir, _ := os.Getwd()
			fileName := name
			filePath := path.Join(fileDir, n.Repository, fileName)

			file, _ := os.Open(filePath)
			defer file.Close()

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
			io.Copy(part, file)
			writer.Close()

			fmt.Println(n.URL + "/repository/" + n.Repository + "/")
			req, err := http.NewRequest("PUT", n.URL+"/repository/"+n.Repository+"/", body)
			if err != nil {
				return err
			}
			req.Header.Add("Content-Type", writer.FormDataContentType())
			req.SetBasicAuth(n.User, os.ExpandEnv(n.Pass))

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusCreated {
				log.Infof("Upload of %v Ok. StatusCode: %v.", name, resp.StatusCode)
			} else {
				log.Error(resp)
				return fmt.Errorf("Upload of %v to %v failed. StatusCode: '%v'", name, n.URL, resp.StatusCode)
			}
		}
	default:
		return fmt.Errorf("Upload of '%s' is not supported", n.ArtifactType)
	}

	return nil
}
