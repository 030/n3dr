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

func sbJAR(p string, c string, artifactNumber string) string {
	log.Debug(c + " found " + p)
	return "maven2.asset" + artifactNumber + "=@" + p + ",maven2.asset" + artifactNumber + ".extension=" + c + ","
}

func sbClassifierJAR(p string, c string, artifactNumber string) string {
	return sbJAR(p, "jar", artifactNumber) + "maven2.asset" + artifactNumber + ".classifier=" + c + ","
}

func multipartContent(path string) string {
	var sb strings.Builder
	switch ext := filepath.Ext(path); ext {
	case ".pom":
		sb.WriteString(sbJAR(path, "pom", "1"))
	case ".war":
		sb.WriteString(sbJAR(path, "war", "2"))
	case ".zip":
		sb.WriteString(sbJAR(path, "zip", "3"))
	case ".jar":
		switch c := jarClassifier(path); c {
		case "sources":
			sb.WriteString(sbClassifierJAR(path, c, "5"))
		case "bundledPdfs":
			sb.WriteString(sbClassifierJAR(path, c, "6"))
		case "testexp":
			sb.WriteString(sbClassifierJAR(path, c, "7"))
		case "standalone":
			sb.WriteString(sbClassifierJAR(path, c, "8"))
		case "test-resources":
			sb.WriteString(sbClassifierJAR(path, c, "9"))
		default:
			sb.WriteString(sbJAR(path, "jar", "4"))
		}
	default:
		log.Debug(path + " not an artifact")
	}

	return sb.String()
}

// Upload posts an artifact as a multipart to a specific nexus3 repository
func (n Nexus3) Upload() error {
	err := n.detectFoldersWithPOM(n.Repository)
	if err != nil {
		return err
	}

	var sb strings.Builder
	for _, v := range pomDirectories() {
		err := filepath.Walk(v, func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !f.IsDir() {
				sb.WriteString(multipartContent(path))
			}
			return nil
		})

		if err != nil {
			return err
		}

		multipartString := strings.TrimSuffix(sb.String(), ",")
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
