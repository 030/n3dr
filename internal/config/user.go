package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type User struct {
	UserID       string   `json:"userId"`
	FirstName    string   `json:"firstName"`
	LastName     string   `json:"lastName"`
	EmailAddress string   `json:"emailAddress"`
	Password     string   `json:"password"`
	Status       string   `json:"status"`
	Roles        []string `json:"roles"`
	Nexus3       Nexus3
}

func (u *User) exists() (bool, error) {
	log.Info("CP")
	req, err := http.NewRequest("GET", u.Nexus3.Endpoint+"/service/rest/v1/security/users?userId="+u.UserID, nil)
	if err != nil {
		log.Info("CP1")
		return false, err
	}
	req.SetBasicAuth(u.Nexus3.Username, u.Nexus3.Password)
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
	log.Info("--------", bodyString)
	if bodyString == "[ ]" && resp.StatusCode == 200 {
		return false, nil
	}
	return true, nil
}

func (u *User) Create() error {
	if exists, _ := u.exists(); exists {
		log.Info("User already exists")
	} else {
		data := User{
			UserID:       u.UserID,
			FirstName:    u.UserID,
			LastName:     u.LastName,
			EmailAddress: u.EmailAddress,
			Password:     u.Password,
			Status:       "active",
			Roles:        []string{"nx-admin"},
		}
		payloadBytes, err := json.Marshal(data)
		if err != nil {
			return err
		}
		body := bytes.NewReader(payloadBytes)

		req, err := http.NewRequest("POST", u.Nexus3.Endpoint+"/service/rest/v1/security/users", body)
		if err != nil {
			return err
		}
		req.SetBasicAuth(u.Nexus3.Username, u.Nexus3.Password)
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
		log.Info(bodyString)
	}
	return nil
}
