package main

import (
	"fmt"

	"github.com/030/n3dr/internal/config/user"
	"github.com/030/n3dr/internal/goswagger/models"
	"github.com/030/n3dr/internal/pkg/http"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	admin, changePass                    bool
	email, firstName, id, lastName, pass string
)

// configUserCmd represents the configUser command
var configUserCmd = &cobra.Command{
	Use:   "configUser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configUser called")

		if !admin && !changePass {
			log.Fatal("either the admin or changePass is required")
		}

		acu := models.APICreateUser{
			EmailAddress: email,
			FirstName:    firstName,
			LastName:     lastName,
			Password:     pass,
			UserID:       id,
		}
		n := http.Nexus3{FQDN: n3drURL, Pass: n3drPass, User: n3drUser}
		u := user.User{APICreateUser: acu, Nexus3: n}

		if admin {
			u.Roles = []string{"nx-admin"}
			if err := u.Create(); err != nil {
				log.Fatal(err)
			}
		}

		if changePass {
			if err := u.ChangePass(); err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configUserCmd)

	configUserCmd.Flags().StringVarP(&email, "email", "", "", "The email of the user")
	if err := configUserCmd.MarkFlagRequired("email"); err != nil {
		log.Fatal(err)
	}

	configUserCmd.Flags().StringVarP(&firstName, "firstName", "", "", "The firstName of the user")
	if err := configUserCmd.MarkFlagRequired("firstName"); err != nil {
		log.Fatal(err)
	}

	configUserCmd.Flags().StringVarP(&lastName, "lastName", "", "", "The lastName of the user")
	if err := configUserCmd.MarkFlagRequired("lastName"); err != nil {
		log.Fatal(err)
	}

	configUserCmd.Flags().StringVarP(&pass, "pass", "", "", "The pass of the user")
	if err := configUserCmd.MarkFlagRequired("pass"); err != nil {
		log.Fatal(err)
	}

	configUserCmd.Flags().StringVarP(&id, "id", "", "", "The id of the user")
	if err := configUserCmd.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}

	configUserCmd.Flags().BoolVar(&admin, "admin", false, "Whether a user should be admin")
	configUserCmd.Flags().BoolVar(&changePass, "changePass", false, "Whether a pass should be changed")
}
