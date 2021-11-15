package main

import (
	"fmt"
	"os"

	"github.com/030/n3dr/internal/config/repository"
	"github.com/030/n3dr/internal/pkg/http"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	configRepoDelete               bool
	configRepoName, configRepoType string
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

		n := http.Nexus3{FQDN: n3drURL, Pass: n3drPass, User: n3drUser}
		r := repository.Repository{Nexus3: n}

		if configRepoDelete {
			if err := r.Delete(configRepoName); err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}

		switch configRepoType {
		case "raw":
			if err := r.CreateRawHosted(configRepoName); err != nil {
				log.Fatal(err)
			}
		default:
			log.Fatalf("configRepoType should not be empty, but: 'raw' and not: '%s'. Did you populate the --configRepoType parameter?", configRepoType)
		}
	},
}

func init() {
	rootCmd.AddCommand(configRepositoryCmd)

	configRepositoryCmd.Flags().StringVarP(&configRepoName, "configRepoName", "", "", "The repository name")
	if err := configRepositoryCmd.MarkFlagRequired("configRepoName"); err != nil {
		log.Fatal(err)
	}

	configRepositoryCmd.Flags().BoolVar(&configRepoDelete, "configRepoDelete", false, "Delete a repository")
	configRepositoryCmd.Flags().StringVarP(&configRepoType, "configRepoType", "", "", "The repository type")
}
