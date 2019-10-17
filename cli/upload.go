package cli

import (
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
			foldersWithPOM.WriteString(filepath.Dir(path) + ",")
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func pomDirectories() []string {
	foldersWithPOMString := strings.TrimSuffix(foldersWithPOM.String(), ",")
	return strings.Split(foldersWithPOMString, ",")
}

func jarClassifier(a string) string {
	re := regexp.MustCompile(".*\\d-(.*)\\.jar$")
	match := re.FindStringSubmatch(a)
	if len(match) == 0 {
		return ""
	}
	return match[1]
}

var sb strings.Builder

func sbJAR(p string, c string, artifactNumber string) {
	log.Debug(c + " found " + p)
	sb.WriteString("maven2.asset" + artifactNumber + "=@" + p + ",")
	sb.WriteString("maven2.asset" + artifactNumber + ".extension=" + c + ",")
}

func sbClassifierJAR(p string, c string, artifactNumber string) {
	sbJAR(p, "jar", artifactNumber)
	sb.WriteString("maven2.asset" + artifactNumber + ".classifier=" + c + ",")
}

func multipartContent(path string) (string, error) {
	switch ext := filepath.Ext(path); ext {
	case ".pom":
		sbJAR(path, "pom", "1")
	case ".war":
		sbJAR(path, "war", "2")
	case ".zip":
		sbJAR(path, "zip", "3")
	case ".jar":
		switch c := jarClassifier(path); c {
		case "sources":
			sbClassifierJAR(path, c, "5")
		case "bundledPdfs":
			sbClassifierJAR(path, c, "6")
		case "testexp":
			sbClassifierJAR(path, c, "7")
		case "standalone":
			sbClassifierJAR(path, c, "8")
		case "test-resources":
			sbClassifierJAR(path, c, "9")
		default:
			sbJAR(path, "jar", "4")
		}
	default:
		log.Debug(path + " not an artifact")
	}

	return strings.TrimSuffix(sb.String(), ","), nil
}

// Upload posts an artifact as a multipart to a specific nexus3 repository
func (n Nexus3) Upload() error {
	err := n.detectFoldersWithPOM(n.Repository)
	if err != nil {
		return err
	}

	for _, v := range pomDirectories() {
		var multipartString string

		err := filepath.Walk(v, func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !f.IsDir() {
				multipartString, err = multipartContent(path)
				if err != nil {
					return err
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		url := n.URL + "/service/rest/" + n.APIVersion + "/components?repository=" + n.Repository
		log.WithFields(log.Fields{
			"multipart": multipartString,
			"url":       url,
		}).Debug("URL")

		u := mp.Upload{URL: url, Username: n.User, Password: n.Pass}
		err = u.MultipartUpload(multipartString)
		if err != nil {
			return err
		}
	}
	return nil
}
