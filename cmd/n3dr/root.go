package main

import (
	"crypto/tls"
	"crypto/x509"
	_ "embed"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/030/logging/pkg/logging"
	cli "github.com/030/n3dr/internal/app/n3dr/artifacts"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

//go:embed assets/logo/text-image-com-n3dr.txt
var logo string

var (
	apiVersion, awsBucket, awsId, awsRegion, awsSecret, basePathPrefix, cfgFile, n3drRepo, n3drURL, n3drPass, n3drUser, Version, zipName, downloadDirName, downloadDirNameZip, dockerHost, logLevel string
	anonymous, awsS3, debug, dockerPortSecure, logFile, https, insecureSkipVerify, showLogo, skipErrors, zip                                                                                        bool
	dockerPort                                                                                                                                                                                      int32
)

var rootCmd = &cobra.Command{
	Use:   "n3dr",
	Short: "nexus3 Disaster Recovery (N3DR)",
	Long: `N3DR is a tool that is capable of backing up all artifacts from a certain
Nexus3 repository and restoring them.`,
	Version: Version,
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&awsS3, "awsS3", "", false, "whether the backup zip should be uploaded to AWS S3")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/"+cli.DefaultCfgFileWithExt+")")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug logging")
	rootCmd.PersistentFlags().BoolVarP(&zip, "zip", "z", false, "add downloaded artifacts to a ZIP archive")
	rootCmd.PersistentFlags().StringVarP(&zipName, "zipName", "i", "", "the name of the zip file")
	rootCmd.PersistentFlags().BoolVar(&insecureSkipVerify, "insecureSkipVerify", false, "Skip repository certificate check")
	rootCmd.PersistentFlags().StringVarP(&n3drPass, "n3drPass", "p", "", "nexus3 password")
	rootCmd.PersistentFlags().StringVarP(&n3drURL, "n3drURL", "n", "", "nexus3 URL")
	rootCmd.PersistentFlags().StringVarP(&n3drUser, "n3drUser", "u", "", "nexus3 user")
	rootCmd.PersistentFlags().StringVar(&downloadDirName, "directory-prefix", "", "directory to store downloaded artifacts")
	rootCmd.PersistentFlags().StringVar(&downloadDirNameZip, "directory-prefix-zip", "", "directory to store the zipped artifacts")
	rootCmd.PersistentFlags().StringVarP(&apiVersion, "apiVersion", "v", "v1", "nexus3 APIVersion, e.g. v1 or beta")
	rootCmd.PersistentFlags().BoolVar(&anonymous, "anonymous", false, "Skip authentication")
	rootCmd.PersistentFlags().BoolVarP(&skipErrors, "skipErrors", "s", false, "Skip errors")
	rootCmd.PersistentFlags().BoolVarP(&https, "https", "", true, "https true or false")
	rootCmd.PersistentFlags().StringVarP(&basePathPrefix, "basePathPrefix", "", "", "the nexus basePathPrefix. Default: \"\"")
	rootCmd.PersistentFlags().BoolVarP(&showLogo, "showLogo", "", true, "show N3DR logo or not. Default: true")
	rootCmd.PersistentFlags().StringVar(&dockerHost, "dockerHost", "", "The docker host, e.g. localhost")
	rootCmd.PersistentFlags().Int32Var(&dockerPort, "dockerPort", 0, "The docker connector port, e.g. 8082")
	rootCmd.PersistentFlags().BoolVar(&dockerPortSecure, "dockerPortSecure", false, "Whether the docker connector port should be secure")
	rootCmd.PersistentFlags().StringVar(&logLevel, "logLevel", "info", "change the log level")
	rootCmd.PersistentFlags().BoolVar(&logFile, "logFile", false, "write the log to syslog")
}

func n3drHiddenHome() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	n3drHomeDir := filepath.Join(home, cli.HiddenN3DR)
	log.Infof("n3drHomeDir: '%v'", n3drHomeDir)
	return n3drHomeDir, nil
}

func configFile() (string, error) {
	n3drHomeDir, err := n3drHiddenHome()
	if err != nil {
		return "", err
	}

	viper.AddConfigPath(n3drHomeDir)
	viper.SetConfigName(cli.DefaultCfgFile)
	viper.SetConfigType(cli.CfgFileExt)

	file := filepath.Join(n3drHomeDir, cli.DefaultCfgFileWithExt)
	log.Debugf("configFile: '%v'", file)
	return file, nil
}

func configFilePath() (string, error) {
	if cfgFile != "" {
		log.Infof("Reading configFile: '%v'", cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		file, err := configFile()
		if err != nil {
			return "", err
		}
		cfgFile = file
	}
	return cfgFile, nil
}

func ascii() {
	if showLogo {
		fmt.Println(logo)
	}
}

func initConfig() {
	l := logging.Logging{File: "n3dr.log", Level: logLevel, Syslog: logFile}
	if _, err := l.Setup(); err != nil {
		log.Fatal(err)
	}

	enableDebug()

	cf, err := configFilePath()
	if err != nil {
		log.Fatal(err)
	}
	if err := parseConfig(cf); err != nil {
		log.Fatal(err)
	}

	viper.AutomaticEnv()
	ascii()

	if err := insecureCerts(); err != nil {
		log.Fatal(err)
	}
}

func valueInConfigFile(key string) (string, error) {
	conf := viper.ConfigFileUsed()
	log.Debugf("reading parameter: '%s' from config file", key)
	value := viper.GetString(key)
	if value == "" {
		return "", fmt.Errorf("key: '%s' does not seem to contain a value. Check whether this key is populated in the config file: '%s'", key, conf)
	}
	return value, nil
}

func parseVarsFromConfig() error {
	var err error
	if !anonymous {
		if n3drUser == "" {
			n3drUser, err = valueInConfigFile("n3drUser")
			if err != nil {
				return err
			}
		}

		if n3drPass == "" {
			n3drPass, err = valueInConfigFile("n3drPass")
			if err != nil {
				return err
			}
		}
	}

	if n3drURL == "" {
		n3drURL, err = valueInConfigFile("n3drURL")
		if err != nil {
			return err
		}
	}

	showLogo = viper.GetBool("showLogo")

	if awsS3 {
		awsBucket, err = valueInConfigFile("awsBucket")
		if err != nil {
			return err
		}
		awsId, err = valueInConfigFile("awsId")
		if err != nil {
			return err
		}
		awsRegion, err = valueInConfigFile("awsRegion")
		if err != nil {
			return err
		}
		awsSecret, err = valueInConfigFile("awsSecret")
		if err != nil {
			return err
		}
	}

	return nil
}

func parseConfig(cfgFile string) error {
	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: '%v'", viper.ConfigFileUsed())
		if err := parseVarsFromConfig(); err != nil {
			return err
		}
	} else {
		log.Warnf("Looked for config file: '%v', but found: '%v' including err: '%v'. Check whether it exists, the YAML is correct and the content is valid", cfgFile, viper.ConfigFileUsed(), err)
	}

	return nil
}

func insecureCerts() error {
	if insecureSkipVerify {
		log.Infof("Loading CA in order to connect to Nexus3...")
		n3drHomeDir, err := n3drHiddenHome()
		if err != nil {
			return err
		}
		caCert, err := os.ReadFile(filepath.Clean(filepath.Join(n3drHomeDir, "ca.crt")))
		if err != nil {
			return err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		log.Warn("Certificate validity check is disabled!")
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{MinVersion: tls.VersionTLS12, MaxVersion: tls.VersionTLS13, RootCAs: caCertPool}
	}
	return nil
}

func enableDebug() {
	log.SetReportCaller(true)
	if debug {
		log.SetLevel(log.DebugLevel)

		// Added to be able to debug viper (used to read the config file)
		// Viper is using a different logger
		jww.SetLogThreshold(jww.LevelTrace)
		jww.SetStdoutThreshold(jww.LevelDebug)
	}
}
