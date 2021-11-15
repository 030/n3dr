package repository

import (
	"fmt"
	"regexp"
	"time"

	"github.com/030/n3dr/internal/goswagger/client/repository_management"
	"github.com/030/n3dr/internal/goswagger/models"
	"github.com/030/n3dr/internal/pkg/http"
	log "github.com/sirupsen/logrus"
)

type Repository struct {
	http.Nexus3
}

func (r *Repository) CreateRawHosted(name string) error {
	log.Infof("Creating repository: '%s'...", name)
	client := r.Nexus3.Client()
	if name == "" {
		return fmt.Errorf("repo name should not be empty")
	}

	// "storage":{"blobStoreName":"default","strictContentTypeValidation":true,"writePolicy":"ALLOW_ONCE"},"maven": {"versionPolicy": "RELEASE","layoutPolicy": "STRICT"}}'

	online := true
	strictContentTypeValidation := true
	writePolicy := "allow_once"
	mhsa := models.HostedStorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &strictContentTypeValidation, WritePolicy: &writePolicy}
	mr := models.RawHostedRepositoryAPIRequest{Name: &name, Online: &online, Storage: &mhsa}
	createRawHosted := repository_management.CreateRepository6Params{Body: &mr}
	createRawHosted.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.CreateRepository6(&createRawHosted); err != nil {
		repositoryCreated, errRegex := regexp.MatchString("status 400", err.Error())
		if errRegex != nil {
			return err
		}
		if repositoryCreated {
			log.Infof("repository: '%s' has already been created", name)
			return nil
		}

		return fmt.Errorf("could not create repository: '%v', err: '%v'", name, err)
	}
	log.Infof("created the following repository: '%v'", name)
	return nil
}

func (r *Repository) Delete(name string) error {
	log.Infof("Deleting repository: '%s'...", name)
	client := r.Nexus3.Client()
	if name == "" {
		return fmt.Errorf("repo name should not be empty")
	}

	deleteRepo := repository_management.DeleteRepositoryParams{RepositoryName: name}
	deleteRepo.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.DeleteRepository(&deleteRepo); err != nil {
		deleteRepositoryNotFound, errRegex := regexp.MatchString("deleteRepositoryNotFound", err.Error())
		if errRegex != nil {
			return err
		}
		if deleteRepositoryNotFound {
			log.Infof("repository: '%s' not found. It seems that it already has been deleted", name)
			return nil
		}

		return fmt.Errorf("could not delete repository: '%v', err: '%v'", name, err)
	}
	log.Infof("deleted the following repository: '%v'", name)
	return nil
}
