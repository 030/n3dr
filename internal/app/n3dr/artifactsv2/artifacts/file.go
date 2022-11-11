package artifacts

import (
	"crypto/sha1" // #nosec
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
	log "github.com/sirupsen/logrus"
)

func ChecksumLocalFile(file, shaType string) (checksum string, err error) {
	f, err := os.Open(filepath.Clean(file))
	if err != nil {
		log.Debugf("file: '%v' not found on local disk", f)
		return "", nil
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	h := sha512.New()
	if shaType == "sha256" {
		h = sha256.New()
	}
	if shaType == "sha1" {
		/* #nosec */
		h = sha1.New()
	}
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	checksum = fmt.Sprintf("%x", h.Sum(nil))

	return checksum, err
}

func PrintType(assetFormat string) {
	switch af := assetFormat; af {
	case "apt":
		fmt.Print("^")
	case "maven2":
		fmt.Print("+")
	case "npm":
		fmt.Print("*")
	case "nuget":
		fmt.Print("$")
	case "raw":
		fmt.Print("%")
	case "yum":
		fmt.Print("#")
	default:
		fmt.Print("?")
		log.Debugf("Unknown type: '%s'", af)
	}
}

func Checksum(asset *models.AssetXO) (string, string) {
	shaType := "sha512"
	checksum := fmt.Sprintf("%s", asset.Checksum[shaType])
	if len(checksum) != 128 {
		shaType = "sha256"
		checksum = fmt.Sprintf("%s", asset.Checksum[shaType])
	}
	if len(checksum) != 64 {
		shaType = "sha1"
		checksum = fmt.Sprintf("%s", asset.Checksum[shaType])
	}
	return shaType, checksum
}

func FilesToBeSkipped(path string) (bool, error) {
	filesToBeSkipped, err := regexp.MatchString(`^\.(sha(1|256|512)|md5|xml)$`, filepath.Ext(path))
	if err != nil {
		return false, err
	}
	if filesToBeSkipped {
		log.Tracef("file: '%s' will be skipped", path)
	}

	return filesToBeSkipped, nil
}
