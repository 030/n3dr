package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func (n Nexus3) validatePassword() error {
	if n.Pass == "" {
		return errors.New("Empty password. Verify whether the 'n3drPass' has been defined in ~/.n3dr.yaml")
	}
	return nil
}

func (n Nexus3) request(url string) ([]byte, string, error) {
	n.validatePassword()
	log.WithFields(log.Fields{"URL": url, "User": n.User}).Debug("URL Request")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}

	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(n.User, n.Pass)

	bodyBytes, bodyString, err := n.response(req)
	if err != nil {
		return nil, "", err
	}
	return bodyBytes, bodyString, err
}

func (n Nexus3) response(req *http.Request) ([]byte, string, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	bodyBytes, bodyString, err := n.responseBodyString(resp)
	if err != nil {
		return nil, "", err
	}

	return bodyBytes, bodyString, nil
}

func (n Nexus3) responseBodyString(resp *http.Response) ([]byte, string, error) {
	var bodyString string
	var bodyBytes []byte
	var err error
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, "", err
		}
		bodyString = string(bodyBytes)
		if bodyString == "[ ]" {
			return nil, "", fmt.Errorf("Bodystring should not be empty. Did the authentication to '%s' succeed?", n.URL)
		}
	} else {
		return nil, "", fmt.Errorf("ResponseCode: '%s' and Message '%s' for URL: %s", strconv.Itoa(resp.StatusCode), resp.Status, resp.Request.URL)
	}

	return bodyBytes, bodyString, nil
}
