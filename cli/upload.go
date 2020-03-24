package cli

import (
	"fmt"
	"os"
	"path/filepath"
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

// Upload posts an artifact as a multipart to a specific nexus3 repository
func (n Nexus3) Upload() error {
	if err := n.detectFoldersWithPOM(n.Repository); err != nil {
		return err
	}

	for _, v := range foldersWithPOMStringSlice {
		var s strings.Builder
		err := filepath.Walk(v, func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !f.IsDir() {
				if filepath.Ext(path) == ".pom" {
					log.Debug("POM found " + path)
					s.WriteString("maven2.asset1=@" + path + ",")
					s.WriteString("maven2.asset1.extension=pom,")
				}

				if filepath.Ext(path) == ".jar" {
					log.Debug("JAR found " + path)
					s.WriteString("maven2.asset2=@" + path + ",")
					s.WriteString("maven2.asset2.extension=jar,")
				}

				if filepath.Ext(path) == ".war" {
					log.Debug("WAR found " + path)
					s.WriteString("maven2.asset3=@" + path + ",")
					s.WriteString("maven2.asset3.extension=war,")
				}

				if filepath.Ext(path) == ".zip" {
					log.Debug("ZIP found " + path)
					s.WriteString("maven2.asset42=@" + path + ",")
					s.WriteString("maven2.asset42.extension=zip,")
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		multipartString := strings.TrimSuffix(s.String(), ",")
		url := n.URL + "/service/rest/" + n.APIVersion + "/components?repository=" + n.Repository
		log.WithFields(log.Fields{
			"multipart": multipartString,
			"url":       url,
		}).Debug("URL")

		u := mp.Upload{URL: url, Username: n.User, Password: n.Pass}
		if err := u.MultipartUpload(multipartString); err != nil {
			return err
		}
	}
	return nil
}
