package cmd

import (
	"strings"

	"github.com/030/n3dr/cli"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var regex string
var npm bool

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup all artifacts from a Nexus3 repository",
	Long: `Use this command in order to backup all artifacts that
reside in a certain Nexus3 repository`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := cli.TempDownloadDir(downloadDirName)
		if err != nil {
			log.Fatal(err)
		}

		selectedRepositories := strings.Split(n3drRepo, ",")
		for _, repository := range selectedRepositories {
			log.Info("Processing repository: ", repository)
			n := cli.Nexus3{URL: n3drURL, User: n3drUser, Pass: n3drPass, Repository: repository, APIVersion: apiVersion, ZIP: zip, ZipName: zipName}
			if err := n.ValidateNexusURL(); err != nil {
				log.Fatal(err)
			}

			if npm {
				log.Info("Backing up an NPM repository...")
				errs := make(chan error, 1)
				if err := n.Bladibla(errs, n.URL+"/service/rest/repository/browse/"+repository); err != nil {
					log.Fatal(err)
				}
				log.Debugf("Go Channel length (outside go routine): '%d'", len(errs))
			} else {
				if err := n.StoreArtifactsOnDiskChannel(dir, regex); err != nil {
					log.Fatal(err)
				}
			}

			if err := n.CreateZip(dir); err != nil {
				log.Fatal(err)
			}
		}
	},
	Version: rootCmd.Version,
}

func init() {
	backupCmd.PersistentFlags().StringVarP(&n3drRepo, "n3drRepo", "r", "", "nexus3 repositories")
	backupCmd.Flags().BoolVarP(&npm, "npm", "", false, "backup an NPM repository")
	backupCmd.Flags().StringVarP(&regex, "regex", "x", ".*", "only download artifacts that match a regular expression, e.g. 'some/group42'")

	if err := backupCmd.MarkPersistentFlagRequired("n3drRepo"); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(backupCmd)
}
