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

	"github.com/spf13/cobra"
)

var (
	names bool
	count bool
)

// repositoriesCmd represents the repositories command
var repositoriesCmd = &cobra.Command{
	Use:   "repositories",
	Short: "Count the number of repositories or return their names",
	Long: `Count the number of repositories or
count the total`,
	Run: func(cmd *cobra.Command, args []string) {
		if names {
			cli.RepositoryNames()
		}
		if count {
			cli.CountRepositories()
		}
	},
}

func init() {
	rootCmd.AddCommand(repositoriesCmd)
	repositoriesCmd.Flags().BoolVarP(&names, "names", "n", false, "Print all repository names")
	repositoriesCmd.Flags().BoolVarP(&count, "count", "c", false, "Count the number of repositories")
}
