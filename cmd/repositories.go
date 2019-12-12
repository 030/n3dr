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
		if err := viper.BindPFlag("n3drPass", rootCmd.Flags().Lookup("n3drPass")); err != nil {
			log.Fatal(err)
		}
		enableDebug()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if !(names || count || backup) {
			if err := cmd.Help(); err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
		pw := viper.GetString("n3drPass")
		n := cli.Nexus3{URL: n3drURL, User: n3drUser, Pass: pw, APIVersion: apiVersion, ZIP: zip}
		if names {
			if err := n.RepositoryNames(); err != nil {
				log.Fatal(err)
			}
		}
		if count {
			if err := n.CountRepositories(); err != nil {
				log.Fatal(err)
			}
		}
		if backup {
			err := n.Downloads()
			if err != nil {
				log.Fatal(err)
			}
		}
	},
	Version: rootCmd.Version,
}

func init() {
	rootCmd.AddCommand(repositoriesCmd)
	repositoriesCmd.Flags().BoolVarP(&names, "names", "a", false, "print all repository names")
	repositoriesCmd.Flags().BoolVarP(&count, "count", "c", false, "count the number of repositories")
	repositoriesCmd.Flags().BoolVarP(&backup, "backup", "b", false, "backup artifacts from all repositories")
}
