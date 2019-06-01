package cli

import (
	"fmt"
	"os"
	"path/filepath"
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
func (n Nexus3) Upload() error {
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
			}
			return nil
		})

		if err != nil {
			return err
		}

		multipartString := strings.TrimSuffix(s.String(), ",")
		fmt.Println(multipartString)
		url := n.URL + "/service/rest/" + n.APIVersion + "/components?repository=" + n.Repository
		u := mp.Upload{URL: url, Username: n.User, Password: n.Pass}
		err2 := u.MultipartUpload(multipartString)
		if err2 != nil {
			return err2
		}
	}
	return nil
}
