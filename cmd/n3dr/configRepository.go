package main

import (
	"fmt"
	"os"

	"github.com/030/n3dr/internal/config/repository"
	"github.com/030/n3dr/internal/pkg/connection"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	configRepoDockerPortSecure, configRepoDelete                         bool
	configRepoDockerPort                                                 int32
	configRepoName, configRepoRecipe, configRepoType, configRepoProxyURL string
)

// configRepositoryCmd represents the configRepository command
var configRepositoryCmd = &cobra.Command{
	Use:   "configRepository",
	Short: "Configure repositories",
	Long: `Configure repositories, e.g.:
* delete a repository
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configRepository called")

		n := connection.Nexus3{FQDN: n3drURL, Pass: n3drPass, User: n3drUser}
		r := repository.Repository{Nexus3: n}

		if configRepoDelete {
			if err := r.Delete(configRepoName); err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}

		if configRepoRecipe == "" {
			log.Fatal("configRepoReceipe should not be empty")
		}

		if configRepoRecipe == "proxy" {
			if configRepoProxyURL == "" {
				log.Fatal("configRepoProxyURL should not be empty")
			}
			r.ProxyRemoteURL = configRepoProxyURL
		}

		switch configRepoType {
		case "apt":
			if configRepoRecipe == "proxy" {
				if err := r.CreateAptProxied(configRepoName); err != nil {
					log.Fatal(err)
				}
			}
		case "docker":
			if configRepoRecipe == "hosted" {
				if err := r.CreateDockerHosted(configRepoDockerPortSecure, configRepoDockerPort, configRepoName); err != nil {
					log.Fatal(err)
				}
			}
		case "raw":
			if configRepoRecipe == "hosted" {
				if err := r.CreateRawHosted(configRepoName); err != nil {
					log.Fatal(err)
				}
			}
		case "yum":
			if configRepoRecipe == "hosted" {
				if err := r.CreateYumHosted(configRepoName); err != nil {
					log.Fatal(err)
				}
			} else if configRepoRecipe == "proxy" {
				if err := r.CreateYumProxied(configRepoName); err != nil {
					log.Fatal(err)
				}
			} else {
				log.Fatalf("configRepoRecipe: '%s' not supported in conjunction with configRepoType: '%s'", configRepoRecipe, configRepoType)
			}
		default:
			log.Fatalf("configRepoType should not be empty, but: 'apt' or 'raw' and not: '%s'. Did you populate the --configRepoType parameter?", configRepoType)
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
	configRepositoryCmd.Flags().StringVar(&configRepoType, "configRepoType", "", "The repository type, e.g.: 'apt', 'raw'")
	configRepositoryCmd.Flags().StringVar(&configRepoProxyURL, "configRepoProxyURL", "", "The proxy repository URL, e.g.: 'http://nl.archive.ubuntu.com/ubuntu/'")
	configRepositoryCmd.Flags().Int32Var(&configRepoDockerPort, "configRepoDockerPort", 8082, "The docker connector port, e.g. 8082")
	configRepositoryCmd.Flags().BoolVar(&configRepoDockerPortSecure, "configRepoDockerPortSecure", false, "Whether the docker connector port should be secure")
}
