package main

import (
	"github.com/030/n3dr/internal/app/n3dr/config/user"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var downloadRole, uploadRole bool

// configUserCmd represents the configUser command.
var configRoleCmd = &cobra.Command{
	Use:   "configRole",
	Short: "Configure roles.",
	Long: `Create roles.

Examples:
  # Create a download role:
  n3dr configRole --downloadRole

  # Create an upload role:
  n3dr configRole --uploadRole --https=false --n3drPass X --n3drUser admin --n3drURL localhost:9999
`,
	Run: func(cmd *cobra.Command, args []string) {
		if !downloadRole && !uploadRole {
			log.Fatal("either the downloadRole or uploadRole is required")
		}

		acu := models.APICreateUser{
			EmailAddress: email,
			FirstName:    firstName,
			LastName:     lastName,
			Password:     pass,
			UserID:       id,
		}
		n := connection.Nexus3{
			FQDN:  n3drURL,
			HTTPS: &https,
			Pass:  n3drPass,
			User:  n3drUser,
		}
		u := user.User{APICreateUser: acu, Nexus3: n}

		if downloadRole {
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
		}

		if uploadRole {
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
		}
	},
}

func init() {
	rootCmd.AddCommand(configRoleCmd)

	configRoleCmd.Flags().BoolVar(&downloadRole, "downloadRole", false, "Whether a download role should be created")
	configRoleCmd.Flags().BoolVar(&uploadRole, "uploadRole", false, "Whether an upload role should be created")
}
