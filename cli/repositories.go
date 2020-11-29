package cli

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/thedevsaddam/gojsonq"
)

func repositoryNamesAndFormatJSON(json string) interface{} {
	jq := gojsonq.New().JSONString(json).WhereNotEqual("type", "group")
	log.Debugf("JQ Output: '%v'", jq)
	jq.SortBy("name", "asc")
	nameAndFormat := jq.Only("name", "format")
	log.Debugf("NameAndFormat: '%v'", nameAndFormat)
	return nameAndFormat
}

func (n Nexus3) repositoriesSlice() ([]interface{}, error) {
	_, repos, err := n.request(n.URL + "/service/rest/" + n.APIVersion + "/repositories")
	if err != nil {
		return nil, err
	}
	return repositoryNamesAndFormatJSON(repos).([]interface{}), nil
}

func (n Nexus3) RepositoryNames() error {
	repos, err := n.repositoriesSlice()
	if err != nil {
		return err
	}
	for _, name := range repos {
		fmt.Printf("%s\n", name)
	}
	return nil
}

func (n Nexus3) CountRepositories() error {
	log.Debug("Counting repositories...")
	repos, err := n.repositoriesSlice()
	if err != nil {
		return err
	}
	fmt.Println(len(repos))
	return nil
}

type response2 struct {
	Format string `json:"format"`
	Name   string `json:"name"`
}

func (n Nexus3) repositoriesChannel(repos []interface{}, dir, regex string) error {
	log.Debugf("Repos: '%v'", repos)
	errs := make(chan error)
	fmt.Println("----------------")
	fmt.Println(repos)
	for _, a := range repos {
		fmt.Println(a)
		if rec, ok := a.(map[string]interface{}); ok {
			for key, val := range rec {
				log.Printf(" [========>] %s = %s", key, val)
			}
		} else {
			fmt.Printf("record not a map[string]interface{}: %v\n", a)
		}
	}
	fmt.Println("----------------")

	for format, name := range repos {
		log.Debugf("Name: '%v'. Format: '%v'", name, format)

		go func(name string) {
			n.Repository = name
			log.Debugf("Repository: '%v'", n.Repository)
			errs <- n.StoreArtifactsOnDiskChannel(dir, regex)
		}(name.(string))
	}
	for range repos {
		if err := <-errs; err != nil {
			return err
		}
	}
	return nil
}

func (n Nexus3) downloadAllArtifactsFromRepositories(dir, regex string) error {
	repos, err := n.repositoriesSlice()
	if err != nil {
		return err
	}
	log.Debugf("Repositories: '%v'", repos)
	if err := n.repositoriesChannel(repos, dir, regex); err != nil {
		return err
	}
	return nil
}

// Downloads retrieves artifacts from all repositories
func (n Nexus3) Downloads(regex string) error {
	dir, err := TempDownloadDir(n.DownloadDirName)
	if err != nil {
		return err
	}
	if err := n.downloadAllArtifactsFromRepositories(dir, regex); err != nil {
		return err
	}
	if err := n.CreateZip(dir); err != nil {
		return err
	}
	return nil
}
