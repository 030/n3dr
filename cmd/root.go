package cmd

import (
	"crypto/tls"
	"fmt"
	"n3dr/cli"
	"net/http"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

var (
	apiVersion, cfgFile, n3drRepo, n3drURL, n3drPass, n3drUser, Version, zipName string
	anonymous, debug, zip, insecureSkipVerify                                    bool
)

var rootCmd = &cobra.Command{
	Use:   "n3dr",
	Short: "nexus3 Disaster Recovery (N3DR)",
	Long: `n3dr is a tool that is able to download all artifacts from
a certain Nexus3 repository.`,
	Version: Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/"+cli.DefaultCfgFileWithExt+")")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug logging")
	rootCmd.PersistentFlags().BoolVarP(&zip, "zip", "z", false, "add downloaded artifacts to a ZIP archive")
	rootCmd.PersistentFlags().StringVarP(&zipName, "zipName", "i", "", "the name of the zip file")
	rootCmd.PersistentFlags().BoolVar(&insecureSkipVerify, "insecureSkipVerify", false, "Skip repository certificate check")
	rootCmd.PersistentFlags().StringVarP(&n3drPass, "n3drPass", "p", "", "nexus3 password")
	rootCmd.PersistentFlags().StringVarP(&n3drURL, "n3drURL", "n", "", "nexus3 URL")
	rootCmd.PersistentFlags().StringVarP(&n3drUser, "n3drUser", "u", "", "nexus3 user")
	rootCmd.PersistentFlags().StringVarP(&apiVersion, "apiVersion", "v", "v1", "nexus3 APIVersion, e.g. v1 or beta")
	rootCmd.PersistentFlags().BoolVar(&anonymous, "anonymous", false, "Skip authentication")
}

func configFile() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	n3drHomeDir := filepath.Join(home, cli.HiddenN3DR)
	log.Infof("n3drHomeDir: '%v'", n3drHomeDir)

	viper.AddConfigPath(n3drHomeDir)
	viper.SetConfigName(cli.DefaultCfgFile)
	viper.SetConfigType(cli.CfgFileExt)

	file := filepath.Join(n3drHomeDir, cli.DefaultCfgFileWithExt)
	log.Debugf("configFile: '%v'", file)
	return file, nil
}

func initConfig() {
	enableDebug()
	insecureCerts()

	if cfgFile != "" {
		log.Infof("Reading configFile: '%v'", cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		file, err := configFile()
		if err != nil {
			log.Fatal(err)
		}
		cfgFile = file
	}
	parseConfig(cfgFile)
	viper.AutomaticEnv()
}

func parseVarsFromConfig() {
	if !anonymous {
		if n3drUser == "" {
			log.Infof("n3drUser empty. Reading if from '%v'", viper.ConfigFileUsed())
			n3drUser = viper.GetString("n3drUser")
		}

		if n3drPass == "" {
			log.Infof("n3drPass empty. Reading if from '%v'", viper.ConfigFileUsed())
			n3drPass = viper.GetString("n3drPass")
		}
	}

	if n3drURL == "" {
		log.Infof("n3drURL empty. Reading if from '%v'", viper.ConfigFileUsed())
		n3drURL = viper.GetString("n3drURL")
	}
}

func parseConfig(cfgFile string) {
	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: '%v'", viper.ConfigFileUsed())
		parseVarsFromConfig()
	} else {
		log.Warnf("Looked for config file: '%v', but found: '%v' including err: '%v'. Check whether it exists, the YAML is correct and the content is valid", cfgFile, viper.ConfigFileUsed(), err)
	}
}

func insecureCerts() {
	if insecureSkipVerify {
		log.Warn("Certificate validity check is disabled!")
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
}

func enableDebug() {
	if debug {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)

		// Added to be able to debug viper (used to read the config file)
		// Viper is using a different logger
		jww.SetLogThreshold(jww.LevelTrace)
		jww.SetStdoutThreshold(jww.LevelDebug)
	}
}
