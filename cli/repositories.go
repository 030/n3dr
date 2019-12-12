package cli

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

func repositoryNamesJSON(json string) interface{} {
	jq := gojsonq.New().JSONString(json)
	jq.SortBy("name", "asc")
	name := jq.Pluck("name")
	return name
}

func (n Nexus3) repositoriesSlice() ([]interface{}, error) {
	_, repos, err := n.request(n.URL + "/service/rest/" + n.APIVersion + "/repositories")
	if err != nil {
		return nil, err
	}
	return repositoryNamesJSON(repos).([]interface{}), nil
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
	repos, err := n.repositoriesSlice()
	if err != nil {
		return err
	}

	fmt.Println(len(repos))
	return nil
}

// Downloads retrieves artifacts from all repositories
func (n Nexus3) Downloads() error {
	repos, err := n.repositoriesSlice()
	if err != nil {
		return err
	}

	for _, name := range repos {
		n := Nexus3{URL: n.URL, User: n.User, Pass: n.Pass, Repository: name.(string), APIVersion: n.APIVersion, ZIP: n.ZIP}
		if err := n.StoreArtifactsOnDisk(); err != nil {
			return err
		}
	}

	if err := n.CreateZip(); err != nil {
		return err
	}

	return nil
}
