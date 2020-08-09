package cli

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

func repositoryNamesJSON(json string) interface{} {
	jq := gojsonq.New().JSONString(json).WhereNotEqual("type", "group")
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

func (n Nexus3) repositoriesChannel(repos []interface{}, dir, regex string) error {
	cerr := make(chan error)
	defer close(cerr)
	for _, name := range repos {
		go func(name string) {
			n.Repository = name
			cerr <- n.StoreArtifactsOnDiskChannel(dir, regex)
		}(name.(string))
	}
	for range repos {
		if err := <-cerr; err != nil {
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
