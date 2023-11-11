package main

import (
	"fmt"
	"os"

	"github.com/030/n3dr/internal/app/n3dr/config/repository"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	configRepoDockerPortSecure, configRepoDelete, snapshot, strictContentTypeValidation bool
	configRepoDockerPort                                                                int32
	configRepoName, configRepoRecipe, configRepoType, configRepoProxyURL                string
	configRepoGroupMemberNames                                                          []string
)

type repo struct {
	conn               repository.Repository
	kind, name, recipe string
	snapshot           bool
}

var repoRecipeAndKindNotSupported = "repoRecipe: '%s' not supported in conjunction with repoKind: '%s'"

func (r *repo) createByType() error {
	switch configRepoType {
	case "apt":
		return r.Apt()
	case "docker":
		return r.Docker()
	case "gem":
		return r.Gem()
	case "maven2":
		return r.Maven2()
	case "npm":
		return r.Npm()
	case "raw":
		return r.Raw()
	case "yum":
		return r.Yum()
	default:
		return fmt.Errorf("configRepoType should not be empty, but: 'apt', 'docker', 'gem', 'maven2', 'npm' 'raw' or 'yum' and not: '%s'. Did you populate the --configRepoType parameter?", configRepoType)
	}
}

func (r *repo) Apt() error {
	switch r.recipe {
	case "proxy":
		return r.conn.CreateAptProxied(r.name)
	default:
		return fmt.Errorf(repoRecipeAndKindNotSupported, r.recipe, r.kind)
	}
}

func (r *repo) Docker() error {
	switch r.recipe {
	case "hosted":
		return r.conn.CreateDockerHosted(configRepoDockerPortSecure, configRepoDockerPort, r.name)
	default:
		return fmt.Errorf(repoRecipeAndKindNotSupported, r.recipe, r.kind)
	}
}

func (r *repo) Gem() error {
	switch r.recipe {
	case "hosted":
		return r.conn.CreateGemHosted(r.name)
	default:
		return fmt.Errorf(repoRecipeAndKindNotSupported, r.recipe, r.kind)
	}
}

func (r *repo) Maven2() error {
	switch r.recipe {
	case "group":
		return r.conn.CreateMavenGroup(configRepoGroupMemberNames, r.name)
	case "hosted":
		return r.conn.CreateMavenHosted(r.name, snapshot)
	case "proxy":
		return r.conn.CreateMavenProxied(r.name)
	default:
		return fmt.Errorf(repoRecipeAndKindNotSupported, r.recipe, r.kind)
	}
}

func (r *repo) Npm() error {
	switch r.recipe {
	case "hosted":
		return r.conn.CreateNpmHosted(r.name, snapshot)
	case "proxy":
		return r.conn.CreateNpmProxied(r.name)
	default:
		return fmt.Errorf(repoRecipeAndKindNotSupported, r.recipe, r.kind)
	}
}

func (r *repo) Raw() error {
	switch r.recipe {
	case "hosted":
		return r.conn.CreateRawHosted(r.name)
	default:
		return fmt.Errorf(repoRecipeAndKindNotSupported, r.recipe, r.kind)
	}
}

func (r *repo) Yum() error {
	switch r.recipe {
	case "hosted":
		return r.conn.CreateYumHosted(r.name)
	case "proxy":
		return r.conn.CreateYumProxied(r.name)
	default:
		return fmt.Errorf(repoRecipeAndKindNotSupported, r.recipe, r.kind)
	}
}

// configRepositoryCmd represents the configRepository command.
var configRepositoryCmd = &cobra.Command{
	Use:   "configRepository",
	Short: "Configure repositories",
	Long: `Configure repositories, e.g.:
* delete a repository

Examples:
  # Create a Docker repository:
  n3dr configRepository -u some-user -p some-pass -n localhost:9000 --https=false --configRepoName some-name --configRepoType docker

  # Create a Maven2 repository if credentials and FQDN have been set in a '~/.n3dr/config.yml' file:
  n3dr configRepository --configRepoName some-name --configRepoType maven2

  # Create a Maven2 repository:
  n3dr configRepository -u some-user -p some-pass -n localhost:9000 --https=false --configRepoName some-name --configRepoType maven2

  # Create a Maven2 repository without strictContentTypeValidation:
  n3dr configRepository -u some-user -p some-pass -n localhost:9000 --https=false --configRepoName some-name --configRepoType maven2 --strictContentTypeValidation=false

  # Create a Maven2 snapshot repository:
  n3dr configRepository -u some-user -p some-pass -n localhost:9000 --https=false --configRepoName some-name --configRepoType maven2 --snapshot

  # Create a NPM repository:
  n3dr configRepository -u admin -p some-pass -n localhost:9000 --https=false --configRepoName 3rdparty-npm --configRepoType npm

  # Create a Rubygems repository:
  n3dr configRepository -u admin -p some-pass -n localhost:9000 --https=false --configRepoName 3rdparty-rubygems --configRepoType gem

  # Create a Maven2 proxy:
  n3dr configRepository --configRepoType maven2 --configRepoName 3rdparty-maven --configRepoRecipe proxy --configRepoProxyURL https://repo.maven.apache.org/maven2/

  # Create a NPM proxy:
  n3dr configRepository --configRepoType npm --configRepoName 3rdparty-npm --configRepoRecipe proxy --configRepoProxyURL https://registry.npmjs.org/

  # Create a group:
  n3dr configRepository --configRepoType maven2 --configRepoRecipe group --configRepoName some-group --configRepoGroupMemberNames releases,snapshots
`,
	Run: func(cmd *cobra.Command, args []string) {
		n := connection.Nexus3{
			FQDN:                        n3drURL,
			HTTPS:                       &https,
			Pass:                        n3drPass,
			StrictContentTypeValidation: strictContentTypeValidation,
			User:                        n3drUser,
		}
		rr := repository.Repository{Nexus3: n}

		if configRepoDelete {
			if err := rr.Delete(configRepoName); err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}

		if configRepoRecipe == "" {
			log.Fatal("configRepoReceipe should not be empty")
		}

		if configRepoRecipe == "proxy" && configRepoProxyURL == "" {
			log.Fatal("configRepoProxyURL should not be empty")

			rr.ProxyRemoteURL = configRepoProxyURL
		}

		r := repo{conn: rr, kind: configRepoType, name: configRepoName, recipe: configRepoRecipe, snapshot: snapshot}
		if err := r.createByType(); err != nil {
			log.Fatalf("repo not created. Error: '%v'", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configRepositoryCmd)

	configRepositoryCmd.Flags().StringVarP(&configRepoName, "configRepoName", "", "", "The repository name")
	if err := configRepositoryCmd.MarkFlagRequired("configRepoName"); err != nil {
		log.Fatal(err)
	}

	configRepositoryCmd.Flags().StringVar(&configRepoRecipe, "configRepoRecipe", "hosted", "The repository recipe, i.e.: group, hosted, or proxy")
	configRepositoryCmd.Flags().BoolVar(&configRepoDelete, "configRepoDelete", false, "Delete a repository")
	configRepositoryCmd.Flags().BoolVar(&snapshot, "snapshot", false, "snapshot repository")
	configRepositoryCmd.Flags().StringVar(&configRepoType, "configRepoType", "", "The repository type, e.g.: 'apt', 'raw'")
	configRepositoryCmd.Flags().StringVar(&configRepoProxyURL, "configRepoProxyURL", "", "The proxy repository URL, e.g.: 'http://nl.archive.ubuntu.com/ubuntu/'")
	configRepositoryCmd.Flags().Int32Var(&configRepoDockerPort, "configRepoDockerPort", 8082, "The docker connector port, e.g. 8082")
	configRepositoryCmd.Flags().BoolVar(&configRepoDockerPortSecure, "configRepoDockerPortSecure", false, "Whether the docker connector port should be secure")
	configRepositoryCmd.Flags().BoolVar(&strictContentTypeValidation, "strictContentTypeValidation", true, "whether strictContentTypeValidation should be enabled")
	configRepositoryCmd.Flags().StringSliceVar(&configRepoGroupMemberNames, "configRepoGroupMemberNames", []string{}, "The repository type, e.g.: 'apt', 'raw'")
}
