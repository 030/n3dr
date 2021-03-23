package main

import (
	cli "github.com/030/n3dr/internal/artifacts"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var artifactType string

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload all artifacts to a specific Nexus3 repository",
	Long: `Use this command in order to upload all artifacts to
a specific Nexus3 repository, e.g. maven-releases`,
	Run: func(cmd *cobra.Command, args []string) {
		n := cli.Nexus3{URL: n3drURL, User: n3drUser, Pass: n3drPass, Repository: n3drRepo, APIVersion: apiVersion, ArtifactType: artifactType}
		if err := n.ValidateNexusURL(); err != nil {
			log.Fatal(err)
		}
		if err := n.Upload(skipErrors); err != nil {
			log.Fatal(err)
		}
	},
	Version: rootCmd.Version,
}

func init() {
	uploadCmd.Flags().StringVarP(&artifactType, "artifactType", "t", "maven2", "type of artifacts that have to be uploaded")
	uploadCmd.PersistentFlags().StringVarP(&n3drRepo, "n3drRepo", "r", "", "nexus3 repository")

	if err := uploadCmd.MarkPersistentFlagRequired("n3drRepo"); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(uploadCmd)
}
