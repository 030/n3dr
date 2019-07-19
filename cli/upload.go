package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	mp "github.com/030/go-curl/utils"
	log "github.com/sirupsen/logrus"
)

var foldersWithPOM strings.Builder

func (n Nexus3) detectFoldersWithPOM(d string) error {
	err := filepath.Walk(d, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() && filepath.Ext(path) == ".pom" {
			fmt.Println(path)
			foldersWithPOM.WriteString(filepath.Dir(path) + ",")
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// Upload posts an artifact as a multipart to a specific nexus3 repository
func (n Nexus3) Upload(c bool) error {
	err3 := n.detectFoldersWithPOM(n.Repository)
	if err3 != nil {
		return err3
	}

	foldersWithPOMString := strings.TrimSuffix(foldersWithPOM.String(), ",")
	foldersWithPOMStringSlice := strings.Split(foldersWithPOMString, ",")

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

				sourcesJAR, err := regexp.MatchString(`sources`, path)
				if err != nil {
					return err
				}
				if filepath.Ext(path) == ".jar" && sourcesJAR {
					log.Debug("Sources JAR found " + path)
					s.WriteString("maven2.asset4=@" + path + ",")
					s.WriteString("maven2.asset4.classifier=sources,")
					s.WriteString("maven2.asset4.extension=jar,")
				}

				bundeldPdfsJAR, err := regexp.MatchString(`bundledPdfs`, path)
				if err != nil {
					return err
				}
				if filepath.Ext(path) == ".jar" && bundeldPdfsJAR {
					log.Debug("bundledPdfs JAR found " + path)
					s.WriteString("maven2.asset5=@" + path + ",")
					s.WriteString("maven2.asset5.classifier=bundledPdfs,")
					s.WriteString("maven2.asset5.extension=jar,")
				}

				testexpJAR, err := regexp.MatchString(`testexp`, path)
				if err != nil {
					return err
				}
				if filepath.Ext(path) == ".jar" && testexpJAR {
					log.Debug("testexp JAR found " + path)
					s.WriteString("maven2.asset6=@" + path + ",")
					s.WriteString("maven2.asset6.classifier=testexp,")
					s.WriteString("maven2.asset6.extension=jar,")
				}

				standaloneJAR, err := regexp.MatchString(`standalone`, path)
				if err != nil {
					return err
				}
				if filepath.Ext(path) == ".jar" && standaloneJAR {
					log.Debug("standalone JAR found " + path)
					s.WriteString("maven2.asset7=@" + path + ",")
					s.WriteString("maven2.asset7.classifier=standalone,")
					s.WriteString("maven2.asset7.extension=jar,")
				}

				testResourcesJAR, err := regexp.MatchString(`test-resources`, path)
				if err != nil {
					return err
				}
				if filepath.Ext(path) == ".jar" && testResourcesJAR {
					log.Debug("testResources JAR found " + path)
					s.WriteString("maven2.asset8=@" + path + ",")
					s.WriteString("maven2.asset8.classifier=test-resources,")
					s.WriteString("maven2.asset8.extension=jar,")
				}

				if filepath.Ext(path) == ".zip" {
					log.Debug("ZIP found " + path)
					s.WriteString("maven2.asset9=@" + path + ",")
					s.WriteString("maven2.asset9.extension=zip,")
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
		err2 := u.MultipartUpload(multipartString)
		if err2 != nil {
			if c {
				log.WithFields(log.Fields{
					"continueOnDoesNotAllowUpdatingArtifacts": c,
					"error": err2,
				}).Info("Ensure that the tool will not stop if an artifact does not exist")
			} else {
				return err2
			}
		}
	}
	return nil
}
