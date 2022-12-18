package main

import (
	"github.com/030/n3dr/internal/app/n3dr/artifactsv2"
	"github.com/030/n3dr/internal/app/n3dr/artifactsv2/upload"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var countV2, namesV2, backupV2, uploadV2 bool

// repositoriesCmd represents the repositories command
var repositoriesV2Cmd = &cobra.Command{
	Use:   "repositoriesV2",
	Short: "Count the number of repositories or return their names.",
	Long: `Count the number of repositories, count the total or
download artifacts from all repositories.

Examples:
  # Return the number of repositories without logging:
  n3dr repositoriesV2 --count --logLevel=none

  # Return the repository names:
  n3dr repositoriesV2 --names

  # Backup a single repository:
  n3dr repositoriesV2 --backup --n3drRepo some-repo --directory-prefix /tmp/some-dir

  # Backup all artifacts:
  n3dr repositoriesV2 --backup --directory-prefix /tmp/some-dir

  # Backup all artifacts, set log level to trace and write it to a file and syslog:
  n3dr repositoriesV2 --backup --directory-prefix /tmp/some-dir --logFile some-file.log --logLevel trace --syslog

  # Backup all artifacts that reside in a Nexus3 server in a certain dir and store these in a zip file:
  n3dr repositoriesV2 --backup --directory-prefix /tmp/some-dir --directory-prefix-zip /tmp/some-dir/some-zip --zip

  # Backup all artifacts including docker images:
  n3dr repositoriesV2 --backup -u some-user -p some-pass -n localhost:9000 --https=false --directory-prefix /tmp/some-dir --dockerPort 9001 --dockerHost http://localhost

  # Upload artifacts, print errors on stderr and write them to syslog:
  n3dr repositoriesV2 --upload -u some-user -p some-pass -n localhost:9000 --https=false --directory-prefix /tmp/some-dir --logLevel=none --logFile
`,
	Run: func(cmd *cobra.Command, args []string) {
		if !(namesV2 || countV2 || backupV2 || uploadV2) {
			if err := cmd.Help(); err != nil {
				log.Fatal(err)
			}
			log.Fatal("One of the required flags \"names\", \"count\" or \"backup\" not set")
		}
		n := connection.Nexus3{AwsBucket: awsBucket, AwsId: awsId, AwsRegion: awsRegion, AwsSecret: awsSecret, BasePathPrefix: basePathPrefix, FQDN: n3drURL, Pass: n3drPass, User: n3drUser, DownloadDirName: downloadDirName, DownloadDirNameZip: downloadDirNameZip, HTTPS: https, DockerHost: dockerHost, DockerPort: dockerPort, DockerPortSecure: dockerPortSecure, ZIP: zip, RepoName: n3drRepo, SkipErrors: skipErrors}
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
			if n.RepoName != "" {
				if err := a.SingleRepoBackup(); err != nil {
					log.Fatal(err)
				}
			} else {
				if err := a.Backup(); err != nil {
					log.Fatal(err)
				}
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
	repositoriesV2Cmd.Flags().StringVar(&n3drRepo, "n3drRepo", "", "backup a single nexus3 repository")

	rootCmd.AddCommand(repositoriesV2Cmd)
}
