package project

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
)

const name = "n3dr"

const hiddenN3DR = "." + name

func Home() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	n3drHomeDir := filepath.Join(home, hiddenN3DR)
	log.Tracef("home: '%v'", n3drHomeDir)
	return n3drHomeDir, nil
}
