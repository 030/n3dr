package cmd

import (
	"n3dr/cli"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var regex string

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup all artifacts from a Nexus3 repository",
	Long: `Use this command in order to backup all artifacts that
reside in a certain Nexus3 repository`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := cli.TempDownloadDir()
		if err != nil {
			log.Fatal(err)
		}

		n := cli.Nexus3{URL: n3drURL, User: n3drUser, Pass: n3drPass, Repository: n3drRepo, APIVersion: apiVersion, ZIP: zip, ZipName: zipName}
		if err := n.StoreArtifactsOnDisk(dir, regex); err != nil {
			log.Fatal(err)
		}

		if err := n.CreateZip(dir); err != nil {
			log.Fatal(err)
		}
	},
	Version: rootCmd.Version,
}

func init() {
	backupCmd.PersistentFlags().StringVarP(&n3drRepo, "n3drRepo", "r", "", "nexus3 repository")
	backupCmd.Flags().StringVarP(&regex, "regex", "x", ".*", "only download artifacts that match a regular expression, e.g. 'some/group42'")

	if err := backupCmd.MarkPersistentFlagRequired("n3drRepo"); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(backupCmd)
}
