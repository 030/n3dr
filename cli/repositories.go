package cli

import (
	"fmt"

	validate "github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/thedevsaddam/gojsonq"
)

type repositoriesMap map[string]string

func bladibla(reposAndFormatJSON interface{}) {
	var m repositoriesMap
	fmt.Println(reposAndFormatJSON)
	for _, a := range reposAndFormatJSON.([]interface{}) {
		fmt.Println(a)
		if rec, ok := a.(map[string]interface{}); ok {
			for key, value := range rec {
				log.Printf(" [========>] %s = %s", key, value)
				var repo string
				var format string
				if key == "name" {
					fmt.Println(value.(string))
					repo = value.(string)
					fmt.Println("dfffffffffff")
				}
				if key == "format" {
					format = value.(string)
					fmt.Println("dffffffffffsdsfsfdsf")
				}

				m[repo] = format
				fmt.Println(m)
			}
		} else {
			fmt.Printf("record not a map[string]interface{}: %v\n", a)
		}
	}
}

func (n Nexus3) repositoriesSlice() (repositoriesMap, error) {
	jsonResp := n.requestJSON(n.URL + "/service/rest/" + n.APIVersion + "/repositories")
	if jsonResp.err != nil {
		return nil, jsonResp.err
	}

	err, reposAndFormatJSON := repositoryNamesAndFormatsMap(jsonResp.strings)
	if err != nil {
		return nil, err
	}

	bladibla(reposAndFormatJSON)

	return nil, nil
}

func repositoryNamesAndFormatsMap(json string) (error, map[string]string) {
	if err := validate.New().Var(json, "required,json"); err != nil {
		return err, nil
	}
	log.Debugf("JSON: '%v'", json)
	jq := gojsonq.New().JSONString(json).WhereNotEqual("type", "group")
	log.Debugf("JQ Output: '%v'", jq)
	jq.SortBy("name", "asc")
	nameAndFormat := jq.Only("name", "format")
	log.Debugf("NameAndFormat: '%v'", nameAndFormat)

	m := make(map[string]string)
	for _, v := range nameAndFormat.([]interface{}) {
		if rec, ok := v.(map[string]interface{}); ok {
			m[rec["name"].(string)] = rec["format"].(string)
		} else {
			fmt.Println("FAIL")
		}
	}
	return nil, m
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

func (n Nexus3) repositoriesChannel(repos repositoriesMap, dir, regex string) error {
	log.Debugf("Repos: '%v'", repos)
	errs := make(chan error)

	for format, name := range repos {
		log.Debugf("Name: '%v'. Format: '%v'", name, format)

		go func(name string) {
			n.Repository = name
			log.Debugf("Repository: '%v'", n.Repository)
			errs <- n.StoreArtifactsOnDiskChannel(dir, regex)
		}(name)
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
