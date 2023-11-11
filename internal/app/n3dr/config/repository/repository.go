package repository

import (
	"fmt"
	"regexp"
	"time"

	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/client/repository_management"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
	log "github.com/sirupsen/logrus"
)

type Repository struct {
	connection.Nexus3
	ProxyRemoteURL string
}

var writePolicy string = "ALLOW_ONCE"

func created(name string, err error) error {
	repositoryCreated, errRegex := regexp.MatchString("status 400", err.Error())
	if errRegex != nil {
		return err
	}
	if repositoryCreated {
		log.Infof("repository: '%s' has already been created", name)
		return nil
	}

	return fmt.Errorf("could not create repository: '%v', err: '%w'", name, err)
}

func (r *Repository) CreateMavenGroup(memberNames []string, name string) error {
	log.Infof("creating maven group: '%s'...", name)
	client, err := r.Nexus3.Client()
	if err != nil {
		return err
	}
	if name == "" {
		return fmt.Errorf("repo name should not be empty")
	}
	if len(memberNames) == 0 {
		return fmt.Errorf("memberNames should not be empty")
	}

	online := true
	mhsa := models.StorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &r.StrictContentTypeValidation}
	group := models.GroupAttributes{MemberNames: memberNames}
	body := models.MavenGroupRepositoryAPIRequest{
		Group:   &group,
		Name:    &name,
		Online:  &online,
		Storage: &mhsa,
	}
	createMavenGroup := repository_management.CreateRepositoryParams{Body: &body}
	createMavenGroup.WithTimeout(time.Second * 30)
	if createRepositoryCreated, err := client.RepositoryManagement.CreateRepository(&createMavenGroup); err != nil {
		log.Debugf("createRepositoryCreated: '%v'", createRepositoryCreated)
		log.Tracef("createRepositoryCreatedError '%v'", err)
		if err := created(name, err); err != nil {
			return err
		}
	}
	log.Infof("created the following maven group: '%v'", name)

	return nil
}

func (r *Repository) CreateAptProxied(name string) error {
	log.Infof("Creating proxied apt repository: '%s'...", name)
	client, err := r.Nexus3.Client()
	if err != nil {
		return err
	}
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
	flat := true
	apt := models.AptProxyRepositoriesAttributes{Distribution: "bionic", Flat: &flat}
	mhsa := models.StorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &r.StrictContentTypeValidation}
	ma := models.AptProxyRepositoryAPIRequest{Apt: &apt, Name: &name, Online: &online, Storage: &mhsa, Proxy: &proxy, NegativeCache: &negativeCache, HTTPClient: &httpClient}
	createAptProxy := repository_management.CreateRepository4Params{Body: &ma}
	createAptProxy.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.CreateRepository4(&createAptProxy); err != nil {
		if err := created(name, err); err != nil {
			return err
		}
	}
	log.Infof("created the following repository: '%v'", name)

	return nil
}

func (r *Repository) CreateNpmProxied(name string) error {
	log.Infof("Creating npm proxy: '%s'...", name)
	client, err := r.Nexus3.Client()
	if err != nil {
		return err
	}
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
	npm := models.NpmAttributes{}
	mhsa := models.StorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &r.StrictContentTypeValidation}
	ma := models.NpmProxyRepositoryAPIRequest{Npm: &npm, Name: &name, Online: &online, Storage: &mhsa, Proxy: &proxy, NegativeCache: &negativeCache, HTTPClient: &httpClient}
	createNpmProxy := repository_management.CreateRepository10Params{Body: &ma}
	createNpmProxy.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.CreateRepository10(&createNpmProxy); err != nil {
		if err := created(name, err); err != nil {
			return err
		}
	}
	log.Infof("created the following repository: '%v'", name)

	return nil
}

func (r *Repository) CreateYumProxied(name string) error {
	log.Infof("Creating proxied yum repository: '%s'...", name)
	client, err := r.Nexus3.Client()
	if err != nil {
		return err
	}
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
	mhsa := models.StorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &r.StrictContentTypeValidation}
	body := models.YumProxyRepositoryAPIRequest{Name: &name, Online: &online, Storage: &mhsa, Proxy: &proxy, NegativeCache: &negativeCache, HTTPClient: &httpClient}
	createYumProxy := repository_management.CreateRepository22Params{Body: &body}
	createYumProxy.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.CreateRepository22(&createYumProxy); err != nil {
		if err := created(name, err); err != nil {
			return err
		}
	}
	log.Infof("created the following repository: '%v'", name)

	return nil
}

func (r *Repository) CreateMavenProxied(name string) error {
	log.Infof("creating the following maven proxy: '%s'...", name)
	client, err := r.Nexus3.Client()
	if err != nil {
		return err
	}
	if name == "" {
		return fmt.Errorf("repo name should not be empty")
	}

	httpClientBlocked := false
	httpClientAutoBlocked := true
	httpClient := models.HTTPClientAttributesWithPreemptiveAuth{AutoBlock: &httpClientAutoBlocked, Blocked: &httpClientBlocked}
	negativeCacheEnabled := true
	var negativeCacheTimeToLive int32 = 1440
	negativeCache := models.NegativeCacheAttributes{Enabled: &negativeCacheEnabled, TimeToLive: &negativeCacheTimeToLive}
	var contentMaxAge int32 = 1440
	var metadataMaxAge int32 = 1440
	remoteURL := r.ProxyRemoteURL
	log.Infof("remoteURL: '%s'", remoteURL)
	proxy := models.ProxyAttributes{ContentMaxAge: &contentMaxAge, MetadataMaxAge: &metadataMaxAge, RemoteURL: remoteURL}
	online := true
	maven := models.MavenAttributes{LayoutPolicy: "STRICT", VersionPolicy: "MIXED"}
	mhsa := models.StorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &r.StrictContentTypeValidation}
	ma := models.MavenProxyRepositoryAPIRequest{Maven: &maven, Name: &name, Online: &online, Storage: &mhsa, Proxy: &proxy, NegativeCache: &negativeCache, HTTPClient: &httpClient}
	createMavenProxy := repository_management.CreateRepository2Params{Body: &ma}
	createMavenProxy.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.CreateRepository2(&createMavenProxy); err != nil {
		if err := created(name, err); err != nil {
			return err
		}
	}
	log.Infof("created the following maven proxy: '%v'", name)

	return nil
}

func (r *Repository) CreateDockerHosted(secure bool, port int32, name string) error {
	log.Infof("Creating docker hosted repository: '%s'...", name)
	client, err := r.Nexus3.Client()
	if err != nil {
		return err
	}
	if name == "" {
		return fmt.Errorf("repo name should not be empty")
	}

	online := true
	dhsa := models.DockerHostedStorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &r.StrictContentTypeValidation, WritePolicy: &writePolicy}

	forceBasicAuth := true
	v1Enabled := false
	docker := models.DockerAttributes{ForceBasicAuth: &forceBasicAuth, V1Enabled: &v1Enabled}
	if secure {
		docker.HTTPSPort = port
	} else {
		docker.HTTPPort = port
	}
	mr := models.DockerHostedRepositoryAPIRequest{Docker: &docker, Name: &name, Online: &online, Storage: &dhsa}
	createRawHosted := repository_management.CreateRepository18Params{Body: &mr}
	createRawHosted.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.CreateRepository18(&createRawHosted); err != nil {
		if err := created(name, err); err != nil {
			return err
		}
	}
	log.Infof("created the following repository: '%v'", name)

	return nil
}

func (r *Repository) CreateGemHosted(name string) error {
	log.Infof("creating gem hosted repository: '%s'...", name)
	client, err := r.Nexus3.Client()
	if err != nil {
		return err
	}
	if name == "" {
		return fmt.Errorf("repo name should not be empty")
	}

	online := true
	rhsa := models.HostedStorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &r.StrictContentTypeValidation, WritePolicy: &writePolicy}
	mr := models.RubyGemsHostedRepositoryAPIRequest{Name: &name, Online: &online, Storage: &rhsa}
	createRubyGemsHosted := repository_management.CreateRepository15Params{Body: &mr}
	createRubyGemsHosted.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.CreateRepository15(&createRubyGemsHosted); err != nil {
		if err := created(name, err); err != nil {
			return err
		}
	}
	log.Infof("created the following repository: '%v'", name)

	return nil
}

func (r *Repository) CreateMavenHosted(name string, snapshot bool) error {
	log.Infof("creating maven hosted repository: '%s'...", name)
	client, err := r.Nexus3.Client()
	if err != nil {
		return err
	}
	if name == "" {
		return fmt.Errorf("repo name should not be empty")
	}

	online := true
	mm := models.MavenAttributes{VersionPolicy: "RELEASE", LayoutPolicy: "STRICT", ContentDisposition: "INLINE"}
	if snapshot {
		mm.VersionPolicy = "SNAPSHOT"
	}

	mhsa := models.HostedStorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &r.StrictContentTypeValidation, WritePolicy: &writePolicy}
	mr := models.MavenHostedRepositoryAPIRequest{Maven: &mm, Name: &name, Online: &online, Storage: &mhsa}
	createMavenHosted := repository_management.CreateRepository1Params{Body: &mr}
	createMavenHosted.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.CreateRepository1(&createMavenHosted); err != nil {
		if err := created(name, err); err != nil {
			return err
		}
	}
	log.Infof("created the following repository: '%v'", name)

	return nil
}

func (r *Repository) CreateNpmHosted(name string, snapshot bool) error {
	log.Infof("creating npm hosted repository: '%s'...", name)
	client, err := r.Nexus3.Client()
	if err != nil {
		return err
	}
	if name == "" {
		return fmt.Errorf("repo name should not be empty")
	}

	online := true
	mhsa := models.HostedStorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &r.StrictContentTypeValidation, WritePolicy: &writePolicy}
	mr := models.NpmHostedRepositoryAPIRequest{Name: &name, Online: &online, Storage: &mhsa}
	createNpmHosted := repository_management.CreateRepository9Params{Body: &mr}
	createNpmHosted.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.CreateRepository9(&createNpmHosted); err != nil {
		if err := created(name, err); err != nil {
			return err
		}
	}
	log.Infof("created the following repository: '%v'", name)

	return nil
}

func (r *Repository) CreateRawHosted(name string) error {
	log.Infof("Creating raw hosted repository: '%s'...", name)
	client, err := r.Nexus3.Client()
	if err != nil {
		return err
	}
	if name == "" {
		return fmt.Errorf("repo name should not be empty")
	}

	online := true
	mhsa := models.HostedStorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &r.StrictContentTypeValidation, WritePolicy: &writePolicy}
	mr := models.RawHostedRepositoryAPIRequest{Name: &name, Online: &online, Storage: &mhsa}
	createRawHosted := repository_management.CreateRepository6Params{Body: &mr}
	createRawHosted.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.CreateRepository6(&createRawHosted); err != nil {
		if err := created(name, err); err != nil {
			return err
		}
	}
	log.Infof("created the following repository: '%v'", name)

	return nil
}

func (r *Repository) CreateYumHosted(name string) error {
	log.Infof("Creating yum hosted repository: '%s'...", name)
	client, err := r.Nexus3.Client()
	if err != nil {
		return err
	}
	if name == "" {
		return fmt.Errorf("repo name should not be empty")
	}

	online := true
	mhsa := models.HostedStorageAttributes{BlobStoreName: "default", StrictContentTypeValidation: &r.StrictContentTypeValidation, WritePolicy: &writePolicy}

	var repoDataDepth int32 = 0
	yum := models.YumAttributes{DeployPolicy: models.YumAttributesDeployPolicySTRICT, RepodataDepth: &repoDataDepth}
	mr := models.YumHostedRepositoryAPIRequest{Name: &name, Online: &online, Storage: &mhsa, Yum: &yum}
	createYumHosted := repository_management.CreateRepository21Params{Body: &mr}
	createYumHosted.WithTimeout(time.Second * 30)
	if _, err := client.RepositoryManagement.CreateRepository21(&createYumHosted); err != nil {
		if err := created(name, err); err != nil {
			return err
		}
	}
	log.Infof("created the following repository: '%v'", name)

	return nil
}

func (r *Repository) Delete(name string) error {
	log.Infof("Deleting repository: '%s'...", name)
	client, err := r.Nexus3.Client()
	if err != nil {
		return err
	}
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

		return fmt.Errorf("could not delete repository: '%v', err: '%w'", name, err)
	}
	log.Infof("deleted the following repository: '%v'", name)

	return nil
}
