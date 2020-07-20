package cmd

import (
	"n3dr/cli"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var regex string

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup all artifacts from a Nexus3 repository",
	Long: `Use this command in order to backup all artifacts that
reside in a certain Nexus3 repository`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := viper.BindPFlag("n3drPass", rootCmd.Flags().Lookup("n3drPass")); err != nil {
			log.Fatal(err)
		}
		enableDebug()
	},
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := cli.TempDownloadDir()
		if err != nil {
			log.Fatal(err)
		}

		n := cli.Nexus3{URL: n3drURL, User: n3drUser, Pass: viper.GetString("n3drPass"), Repository: n3drRepo, APIVersion: apiVersion, ZIP: zip, ZipName: zipName}
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

	if err := backupCmd.MarkPersistentFlagRequired("n3drRepo"); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().StringVarP(&regex, "regex", "x", ".*", "only download artifacts that match a regular expression, e.g. 'some/group42'")
}
