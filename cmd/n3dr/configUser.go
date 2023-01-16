package main

import (
	"fmt"

	"github.com/030/n3dr/internal/app/n3dr/config/user"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	admin, changePass, downloadUser, uploadUser bool
	email, firstName, id, lastName, pass        string
)

// configUserCmd represents the configUser command.
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

		if !admin && !downloadUser && !uploadUser && !changePass {
			log.Fatal("either the admin, changePass, downloadUser or uploadUser is required")
		}

		acu := models.APICreateUser{
			EmailAddress: email,
			FirstName:    firstName,
			LastName:     lastName,
			Password:     pass,
			UserID:       id,
		}
		n := connection.Nexus3{FQDN: n3drURL, Pass: n3drPass, User: n3drUser}
		u := user.User{APICreateUser: acu, Nexus3: n}

		if admin {
			u.Roles = []string{"nx-admin"}
			if err := u.Create(); err != nil {
				log.Fatal(err)
			}
		}

		if downloadUser {
			u.Roles = []string{"nx-download"}
			rr := models.RoleXORequest{
				ID:   "nx-download",
				Name: "nx-download",
				Privileges: []string{
					"nx-repository-view-*-*-browse",
					"nx-repository-view-*-*-read",
				},
			}
			r := user.Role{RoleXORequest: rr, Nexus3: n}
			if err := r.CreateRole(); err != nil {
				log.Fatal(err)
			}
			if err := u.Create(); err != nil {
				log.Fatal(err)
			}
		}

		if uploadUser {
			u.Roles = []string{"nx-upload"}
			rr := models.RoleXORequest{
				ID:   "nx-upload",
				Name: "nx-upload",
				Privileges: []string{
					"nx-repository-view-*-*-add",
					"nx-repository-view-*-*-edit",
				},
			}
			r := user.Role{RoleXORequest: rr, Nexus3: n}
			if err := r.CreateRole(); err != nil {
				log.Fatal(err)
			}
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
	configUserCmd.Flags().BoolVar(&downloadUser, "downloadUser", false, "Whether a user should be able to download")
	configUserCmd.Flags().BoolVar(&uploadUser, "uploadUser", false, "Whether a user should be able to upload")
	configUserCmd.Flags().BoolVar(&changePass, "changePass", false, "Whether a pass should be changed")
}
