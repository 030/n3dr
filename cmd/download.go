// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"n3dr/cli"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	n3drURL  string
	n3drPass string
	n3drRepo string
	n3drUser string
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download all artifacts from a Nexus3 repository",
	Long: `Use this command in order to download all artifacts that
reside in a certain Nexus3 repository`,
	Run: func(cmd *cobra.Command, args []string) {
		n := cli.Nexus3{URL: n3drURL, User: n3drUser, Pass: n3drPass, Repository: n3drRepo}
		err := n.StoreArtifactsOnDisk()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringVarP(&n3drURL, "n3drURL", "n", "http://localhost:8081", "The Nexus3 URL")
	downloadCmd.Flags().StringVarP(&n3drPass, "n3drPass", "p", "admin123", "The Nexus3 password")
	downloadCmd.Flags().StringVarP(&n3drRepo, "n3drRepo", "r", "maven-releases", "The Nexus3 repository")
	downloadCmd.Flags().StringVarP(&n3drUser, "n3drUser", "u", "admin", "The Nexus3 user")
}
