package main

import (
	"github.com/030/n3dr/internal/app/n3dr/config/security"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configUserAnonymous bool

// configCmd represents the config command.
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config",
	Long:  `config`,
	Run: func(cmd *cobra.Command, args []string) {
		n := connection.Nexus3{
			FQDN:  n3drURL,
			HTTPS: &https,
			Pass:  n3drPass,
			User:  n3drUser,
		}
		s := security.Security{Nexus3: n}
		if err := s.Anonymous(configUserAnonymous); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().BoolVar(&configUserAnonymous, "configUserAnonymous", false, "Whether anonymous mode should be enabled or disabled")
}
