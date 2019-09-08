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
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	apiVersion, cfgFile, n3drRepo, n3drURL, n3drUser, Version string
	debug, zip                                                bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "n3dr",
	Short: "Nexus3 Disaster Recovery (N3DR)",
	Long: `N3DR is a tool that is able to download all artifacts from
a certain Nexus3 repository.`,
	Version: Version,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")
	rootCmd.PersistentFlags().BoolVarP(&zip, "zip", "z", false, "Add downloaded artifacts to a ZIP archive")

	rootCmd.PersistentFlags().StringP("n3drPass", "p", "", "The Nexus3 password")
	rootCmd.PersistentFlags().StringVarP(&n3drURL, "n3drURL", "n", "", "The Nexus3 URL")
	rootCmd.PersistentFlags().StringVarP(&n3drUser, "n3drUser", "u", "", "The Nexus3 user")
	rootCmd.PersistentFlags().StringVarP(&apiVersion, "apiVersion", "v", "v1", "The Nexus3 APIVersion, e.g. v1 or beta")

	rootCmd.MarkPersistentFlagRequired("n3drURL")
	rootCmd.MarkPersistentFlagRequired("n3drUser")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".n3dr" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".n3dr")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("~/.n3dr.yaml does not exist or yaml is invalid")
	}
}

func enableDebug() {
	if debug {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	}
}
