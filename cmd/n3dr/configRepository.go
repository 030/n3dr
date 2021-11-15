package main

import (
	"fmt"

	"github.com/030/n3dr/internal/config/repository"
	"github.com/030/n3dr/internal/pkg/http"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	configRepoDelete bool
	configRepoName   string
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

		if !configRepoDelete {
			log.Fatal("configRepoDelete is required")
		}

		if configRepoDelete {
			n := http.Nexus3{FQDN: n3drURL, Pass: n3drPass, User: n3drUser}
			r := repository.Repository{Nexus3: n}
			if err := r.Delete(configRepoName); err != nil {
				log.Fatal(err)
			}
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
}
