package cmd

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	apiVersion, cfgFile, n3drRepo, n3drURL, n3drUser, Version, zipName string
	debug, zip, insecureSkipVerify                                     bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "n3dr",
	Short: "nexus3 Disaster Recovery (N3DR)",
	Long: `n3dr is a tool that is able to download all artifacts from
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

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug logging")
	rootCmd.PersistentFlags().BoolVarP(&zip, "zip", "z", false, "add downloaded artifacts to a ZIP archive")
	rootCmd.PersistentFlags().StringVarP(&zipName, "zipName", "i", "", "the name of the zip file")
	rootCmd.PersistentFlags().BoolVar(&insecureSkipVerify, "insecureSkipVerify", false, "Skip repository certificate check")

	rootCmd.PersistentFlags().StringP("n3drPass", "p", "", "nexus3 password")
	rootCmd.PersistentFlags().StringVarP(&n3drURL, "n3drURL", "n", "", "nexus3 URL")
	rootCmd.PersistentFlags().StringVarP(&n3drUser, "n3drUser", "u", "", "nexus3 user")
	rootCmd.PersistentFlags().StringVarP(&apiVersion, "apiVersion", "v", "v1", "nexus3 APIVersion, e.g. v1 or beta")

	if err := rootCmd.MarkPersistentFlagRequired("n3drURL"); err != nil {
		log.Fatal(err)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile) // Use config file from the flag.
	} else {
		home, err := homedir.Dir() // Find home directory.
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		log.Infof("HomeDir: '%v'", home)

		viper.AddConfigPath(home) // Search config in home directory with name ".n3dr" (without extension).
		viper.SetConfigName(".n3dr")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err != nil {
		log.Warn("~/.n3dr.yaml does not exist or yaml is invalid")
	}

	if insecureSkipVerify {
		log.Warn("Certificate validity check is disabled!")
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

}

func enableDebug() {
	if debug {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	}
}
