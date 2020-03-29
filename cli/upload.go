package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	mp "github.com/030/go-multipart/utils"
	log "github.com/sirupsen/logrus"
)

var foldersWithPOM strings.Builder
var foldersWithPOMStringSlice []string

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

func sbArtifact(path string, ext string, number string) string {
	log.Debug(ext + " found " + path)
	return "maven2.asset" + number + "=@" + path + ",maven2.asset" + number + ".extension=" + ext + ","
}

func artifactTypeDetector(path string, sb strings.Builder) strings.Builder {
	switch ext := filepath.Ext(path); ext {
	case ".pom":
		sb.WriteString(sbArtifact(path, "pom", "1"))
	case ".jar":
		sb.WriteString(sbArtifact(path, "jar", "2"))
	case ".war":
		sb.WriteString(sbArtifact(path, "war", "3"))
	case ".zip":
		sb.WriteString(sbArtifact(path, "zip", "42"))
	default:
		log.Debug(path + " not an artifact")
	}
	return sb
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
	err := filepath.Walk(p, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			sb = artifactTypeDetector(path, sb)
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
