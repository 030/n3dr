package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"

	log "github.com/sirupsen/logrus"
)

type Repository struct {
	Name    string  `json:"name"`
	Online  bool    `json:"online"`
	Storage Storage `json:"storage"`
	Nexus3  Nexus3
}

type Storage struct {
	BlobStoreName               string `json:"blobStoreName"`
	StrictContentTypeValidation bool   `json:"strictContentTypeValidation"`
	WritePolicy                 string `json:"writePolicy"`
}

func (r *Repository) exists() (bool, error) {
	log.Info("CP")
	req, err := http.NewRequest("GET", r.Nexus3.Endpoint+"/service/rest/v1/repositories", nil)
	if err != nil {
		log.Info("CP1")
		return false, err
	}
	req.SetBasicAuth(r.Nexus3.Username, r.Nexus3.Password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Info("CP2")
		return false, err
	}
	defer resp.Body.Close()
	log.Info(resp.StatusCode)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	log.Info("CP2a")
	if err != nil {
		log.Info("CP3")
		return false, err
	}
	bodyString := string(bodyBytes)
	log.Info("-------->", bodyString)

	repoExists, err := regexp.MatchString(r.Name, bodyString)
	if err != nil {
		return false, err
	}

	if repoExists && resp.StatusCode == 200 {
		return true, nil
	}
	return false, nil
}

func (r *Repository) Create() error {
	if exists, _ := r.exists(); exists {
		log.Info("Repository already exists")
	} else {
		data := Repository{
			Name:   r.Name,
			Online: true,
			Storage: Storage{
				BlobStoreName:               "default",
				StrictContentTypeValidation: true,
				WritePolicy:                 "ALLOW_ONCE",
			},
		}
		payloadBytes, err := json.Marshal(data)
		if err != nil {
			return err
		}
		body := bytes.NewReader(payloadBytes)

		req, err := http.NewRequest("POST", r.Nexus3.Endpoint+"/service/rest/v1/repositories/npm/hosted", body)
		if err != nil {
			return err
		}
		req.SetBasicAuth(r.Nexus3.Username, r.Nexus3.Password)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		log.Info("---------------------", bodyString)
	}
	return nil
}
