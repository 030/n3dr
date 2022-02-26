package repository

import (
	"fmt"
	"regexp"
	"time"

	"github.com/030/n3dr/internal/goswagger/client/repository_management"
	"github.com/030/n3dr/internal/goswagger/models"
	"github.com/030/n3dr/internal/pkg/connection"
	log "github.com/sirupsen/logrus"
)

type Repository struct {
	connection.Nexus3
	ProxyRemoteURL string
}

func (r *Repository) CreateAptProxied(name string) error {
	log.Infof("Creating proxied apt repository: '%s'...", name)
	client := r.Nexus3.Client()
	if name == "" {
		return fmt.Errorf("repo name should not be empty")
	}

	httpClientBlocked := false
	httpClientAutoBlocked := true
	httpClient := models.HTTPClientAttributes{AutoBlock: &httpClientAutoBlocked, Blocked: &httpClientBlocked}
	negativeCacheEnabled := true
	var negativeCacheTimeToLive int32 = 1440
	negativeCache := models.NegativeCacheAttributes{Enabled: &negativeCacheEnabled, TimeToLive: &negativeCacheTimeToLive}
	var contentMaxAge int32 = 1440
	var metadataMaxAge int32 = 1440
	remoteURL := r.ProxyRemoteURL
	proxy := models.ProxyAttributes{ContentMaxAge: &contentMaxAge, MetadataMaxAge: &metadataMaxAge, RemoteURL: remoteURL}
	online := true
	strictContentTypeValidation := true
	flat := true
	apt := models.AptProxyRepositoriesAttributes{Distribution: "bionic", Flat: &flat}
	mhsa := models.StorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &strictContentTypeValidation}
	ma := models.AptProxyRepositoryAPIRequest{Apt: &apt, Name: &name, Online: &online, Storage: &mhsa, Proxy: &proxy, NegativeCache: &negativeCache, HTTPClient: &httpClient}
	createAptProxy := repository_management.CreateRepository4Params{Body: &ma}
	createAptProxy.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.CreateRepository4(&createAptProxy); err != nil {
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

func (r *Repository) CreateYumProxied(name string) error {
	log.Infof("Creating proxied yum repository: '%s'...", name)
	client := r.Nexus3.Client()
	if name == "" {
		return fmt.Errorf("repo name should not be empty")
	}

	httpClientBlocked := false
	httpClientAutoBlocked := true
	httpClient := models.HTTPClientAttributes{AutoBlock: &httpClientAutoBlocked, Blocked: &httpClientBlocked}
	negativeCacheEnabled := true
	var negativeCacheTimeToLive int32 = 1440
	negativeCache := models.NegativeCacheAttributes{Enabled: &negativeCacheEnabled, TimeToLive: &negativeCacheTimeToLive}
	var contentMaxAge int32 = 1440
	var metadataMaxAge int32 = 1440
	remoteURL := r.ProxyRemoteURL
	proxy := models.ProxyAttributes{ContentMaxAge: &contentMaxAge, MetadataMaxAge: &metadataMaxAge, RemoteURL: remoteURL}
	online := true
	strictContentTypeValidation := true
	mhsa := models.StorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &strictContentTypeValidation}
	body := models.YumProxyRepositoryAPIRequest{Name: &name, Online: &online, Storage: &mhsa, Proxy: &proxy, NegativeCache: &negativeCache, HTTPClient: &httpClient}
	createYumProxy := repository_management.CreateRepository22Params{Body: &body}
	createYumProxy.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.CreateRepository22(&createYumProxy); err != nil {
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

func (r *Repository) CreateDockerHosted(secure bool, port int32, name string) error {
	log.Infof("Creating docker hosted repository: '%s'...", name)
	client := r.Nexus3.Client()
	if name == "" {
		return fmt.Errorf("repo name should not be empty")
	}

	online := true
	strictContentTypeValidation := true
	writePolicy := "allow_once"
	mhsa := models.HostedStorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &strictContentTypeValidation, WritePolicy: &writePolicy}

	var forceBasicAuth = true
	var v1Enabled = false
	docker := models.DockerAttributes{ForceBasicAuth: &forceBasicAuth, V1Enabled: &v1Enabled}
	if secure {
		docker.HTTPSPort = port
	} else {
		docker.HTTPPort = port
	}
	mr := models.DockerHostedRepositoryAPIRequest{Docker: &docker, Name: &name, Online: &online, Storage: &mhsa}
	createRawHosted := repository_management.CreateRepository18Params{Body: &mr}
	createRawHosted.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.CreateRepository18(&createRawHosted); err != nil {
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

func (r *Repository) CreateRawHosted(name string) error {
	log.Infof("Creating raw hosted repository: '%s'...", name)
	client := r.Nexus3.Client()
	if name == "" {
		return fmt.Errorf("repo name should not be empty")
	}

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

func (r *Repository) CreateYumHosted(name string) error {
	log.Infof("Creating yum hosted repository: '%s'...", name)
	client := r.Nexus3.Client()
	if name == "" {
		return fmt.Errorf("repo name should not be empty")
	}

	online := true
	strictContentTypeValidation := true
	writePolicy := "allow_once"
	mhsa := models.HostedStorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &strictContentTypeValidation, WritePolicy: &writePolicy}

	var repoDataDepth int32 = 0
	yum := models.YumAttributes{DeployPolicy: models.YumAttributesDeployPolicySTRICT, RepodataDepth: &repoDataDepth}
	mr := models.YumHostedRepositoryAPIRequest{Name: &name, Online: &online, Storage: &mhsa, Yum: &yum}
	createYumHosted := repository_management.CreateRepository21Params{Body: &mr}
	createYumHosted.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.CreateRepository21(&createYumHosted); err != nil {
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
