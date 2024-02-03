package main

import (
	"github.com/030/n3dr/internal/app/n3dr/config/user"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	admin, changePass, custom, downloadUser, uploadUser bool
	email, firstName, id, lastName, pass                string
	roles                                               []string
)

// configUserCmd represents the configUser command.
var configUserCmd = &cobra.Command{
	Use:   "configUser",
	Short: "Configure users.",
	Long: `Create users or change their passwords.

Examples:
  # Create an admin user:
  n3dr configUser --pass some-pass --email build@example.org --firstName build --id build --lastName build --admin

  # Create a download user:
  n3dr configUser --pass some-pass --email build@example.org --firstName build --id build --lastName build --downloadUser

  # Create an upload user:
  n3dr configUser --pass some-pass --email build@example.org --firstName build --id build --lastName build --uploadUser

  # Create a custom user and assign certain roles:
  n3dr configUser --pass some-pass --email build@example.org --firstName build --id build --lastName build --roles nx-download,nx-upload --custom

  # Change the admin password:
  n3dr configUser --changePass --https=false --n3drUser admin --n3drURL nexus3:8081 --n3drPass initial-pass --pass some-pass --email admin@example.org --firstName admin --id admin --lastName admin
`,
	Run: func(cmd *cobra.Command, args []string) {
		if !admin && !custom && !downloadUser && !uploadUser && !changePass {
			log.Fatal("either the admin, custom, changePass, create, downloadUser or uploadUser is required")
		}

		acu := models.APICreateUser{
			EmailAddress: email,
			FirstName:    firstName,
			LastName:     lastName,
			Password:     pass,
			Roles:        roles,
			UserID:       id,
		}
		n := connection.Nexus3{
			FQDN:  n3drURL,
			HTTPS: &https,
			Pass:  n3drPass,
			User:  n3drUser,
		}
		u := user.User{APICreateUser: acu, Nexus3: n}

		if admin {
			u.Roles = []string{"nx-admin"}
			if err := u.Create(); err != nil {
				log.Fatal(err)
			}
		}

		if custom {
			u.Roles = roles
			log.Info("roles:", u)

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
	configUserCmd.Flags().BoolVar(&custom, "custom", false, "Create a user and assign certain roles")
	configUserCmd.Flags().BoolVar(&downloadUser, "downloadUser", false, "Whether a user should be able to download")
	configUserCmd.Flags().BoolVar(&uploadUser, "uploadUser", false, "Whether a user should be able to upload")
	configUserCmd.Flags().BoolVar(&changePass, "changePass", false, "Whether a pass should be changed")

	configUserCmd.Flags().StringSliceVar(&roles, "roles", nil, "Which roles have to be assigned to the custom user")
}
