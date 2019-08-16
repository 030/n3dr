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
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	names  bool
	count  bool
	backup bool
)

// repositoriesCmd represents the repositories command
var repositoriesCmd = &cobra.Command{
	Use:   "repositories",
	Short: "Count the number of repositories or return their names",
	Long: `Count the number of repositories, count the total or
download artifacts from all repositories`,
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("n3drPass", rootCmd.Flags().Lookup("n3drPass"))
		enableDebug()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if !(names || count || backup) {
			cmd.Help()
			os.Exit(0)
		}
		pw := viper.GetString("n3drPass")
		n := cli.Nexus3{URL: n3drURL, User: n3drUser, Pass: pw, APIVersion: apiVersion, ZIP: zip}
		if names {
			n.RepositoryNames()
		}
		if count {
			n.CountRepositories()
		}
		if backup {
			err := n.Downloads()
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(repositoriesCmd)
	repositoriesCmd.Flags().BoolVarP(&names, "names", "a", false, "Print all repository names")
	repositoriesCmd.Flags().BoolVarP(&count, "count", "c", false, "Count the number of repositories")
	repositoriesCmd.Flags().BoolVarP(&backup, "backup", "b", false, "Backup artifacts from all repositories")
}
