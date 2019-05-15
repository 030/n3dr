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

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("download called")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	downloadCmd.Flags().StringP("n3drURL", "n", "http://localhost:8081", "The Nexus3 URL")
	downloadCmd.Flags().StringP("n3drPass", "p", "admin123", "The Nexus3 password")
	downloadCmd.Flags().StringP("n3drRepo", "r", "maven-releases", "The Nexus3 repository")
	downloadCmd.Flags().StringP("n3drUser", "u", "admin", "The Nexus3 user")
	downloadCmd.MarkFlagRequired("n3drURL")
	downloadCmd.MarkFlagRequired("n3drPass")
	downloadCmd.MarkFlagRequired("n3drRepo")
	downloadCmd.MarkFlagRequired("n3drUser")
}

func lookupCobraFlag(cobraFlag string) string {
	v, err := downloadCmd.Flags().GetString(cobraFlag)
	log.Info(cobraFlag + ": " + v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func downloadArtifacts() {
	n := cli.Nexus3{URL: lookupCobraFlag("n3drURL"), User: lookupCobraFlag("n3drUser"), Pass: lookupCobraFlag("n3drPass"), Repository: lookupCobraFlag("n3drRepo")}
	err := n.StoreArtifactsOnDisk()
	if err != nil {
		log.Fatal(err)
	}
}
