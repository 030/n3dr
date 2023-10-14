package main

import (
	"github.com/030/n3dr/internal/app/n3dr/artifactsv2/count"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	csv  string
	sort bool
)

// repositoriesCmd represents the repositories command.
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count the number of artifacts.",
	Long: `Count the number of artifacts and write them to a CSV if required.

Examples:
  # Return the number of artifacts without logging:
  n3dr count --logLevel=none

  # Return the number of artifacts and write them to a /tmp/helloworld.csv:
  n3dr count --csv /tmp/helloworld

  # Return the number of artifacts, write them to a /tmp/helloworld.csv and sort it:
  n3dr count --csv /tmp/helloworld --sort
`,
	Run: func(cmd *cobra.Command, args []string) {
		n := connection.Nexus3{
			AwsBucket:                    awsBucket,
			AwsID:                        awsID,
			AwsRegion:                    awsRegion,
			AwsSecret:                    awsSecret,
			BasePathPrefix:               basePathPrefix,
			DockerHost:                   dockerHost,
			DockerPort:                   dockerPort,
			DockerPortSecure:             dockerPortSecure,
			DownloadDirName:              downloadDirName,
			DownloadDirNameZip:           downloadDirNameZip,
			FQDN:                         n3drURL,
			HTTPS:                        &https,
			Pass:                         n3drPass,
			RepoName:                     n3drRepo,
			SkipErrors:                   skipErrors,
			User:                         n3drUser,
			WithoutWaitGroupArtifacts:    withoutWaitGroupArtifacts,
			WithoutWaitGroupRepositories: withoutWaitGroupRepositories,
			WithoutWaitGroups:            withoutWaitGroups,
			ZIP:                          zip,
		}
		c := count.Nexus3{Nexus3: &n, CsvFile: csv, Sort: sort}

		if err := c.Artifacts(); err != nil {
			log.Fatal(err)
		}
	},
	Version: rootCmd.Version,
}

func init() {
	countCmd.Flags().StringVar(&csv, "csv", "", "write to a csvFile")
	countCmd.Flags().BoolVar(&sort, "sort", false, "sort the csvFile")

	rootCmd.AddCommand(countCmd)
}
