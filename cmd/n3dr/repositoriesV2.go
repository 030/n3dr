package main

import (
	"github.com/030/n3dr/internal/artifactsv2"
	"github.com/030/n3dr/internal/artifactsv2/upload"
	"github.com/030/n3dr/internal/pkg/connection"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	countV2, namesV2, backupV2, uploadV2 bool
)

// repositoriesCmd represents the repositories command
var repositoriesV2Cmd = &cobra.Command{
	Use:   "repositoriesV2",
	Short: "Count the number of repositories or return their names",
	Long: `Count the number of repositories, count the total or
download artifacts from all repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		if !(namesV2 || countV2 || backupV2 || uploadV2) {
			if err := cmd.Help(); err != nil {
				log.Fatal(err)
			}
			log.Fatal("One of the required flags \"names\", \"count\" or \"backup\" not set")
		}
		n := connection.Nexus3{BasePathPrefix: basePathPrefix, FQDN: n3drURL, Pass: n3drPass, User: n3drUser, DownloadDirName: downloadDirName, HTTPS: https}
		a := artifactsv2.Nexus3{Nexus3: &n}
		if namesV2 {
			if err := a.RepositoryNamesV2(); err != nil {
				log.Fatal(err)
			}
		}
		if countV2 {
			if err := a.CountRepositoriesV2(); err != nil {
				log.Fatal(err)
			}
		}
		if backupV2 {
			if err := a.Backup(); err != nil {
				log.Fatal(err)
			}
		}
		if uploadV2 {
			u := upload.Nexus3{Nexus3: &n}
			if err := u.Upload(); err != nil {
				log.Fatal(err)
			}
		}
	},
	Version: rootCmd.Version,
}

func init() {
	repositoriesV2Cmd.Flags().BoolVarP(&namesV2, "names", "", false, "print all repository names")
	repositoriesV2Cmd.Flags().BoolVarP(&countV2, "count", "", false, "count the number of repositories")
	repositoriesV2Cmd.Flags().BoolVarP(&backupV2, "backup", "", false, "backup artifacts from all repositories")
	repositoriesV2Cmd.Flags().BoolVarP(&uploadV2, "upload", "", false, "upload artifacts from all repositories")
	repositoriesV2Cmd.Flags().StringVarP(&regex, "regex", "x", ".*", "only download artifacts that match a regular expression, e.g. 'some/group42'")

	rootCmd.AddCommand(repositoriesV2Cmd)
}
