package user

import (
	"fmt"
	"regexp"
	"time"

	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/client/security_management_roles"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/client/security_management_users"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
	log "github.com/sirupsen/logrus"
)

type Role struct {
	connection.Nexus3
	models.RoleXORequest
}

type User struct {
	connection.Nexus3
	models.APICreateUser
}

func (u *User) Create() error {
	log.Infof("creating user: '%s'...", u.UserID)
	status := "active"
	u.Status = &status

	client := u.Nexus3.Client()

	createUser := security_management_users.CreateUserParams{Body: &u.APICreateUser}
	createUser.WithTimeout(time.Second * 30)
	resp, err := client.SecurityManagementUsers.CreateUser(&createUser)
	if err != nil {
		userCreated, errRegex := regexp.MatchString("status 500", err.Error())
		if errRegex != nil {
			return err
		}
		if userCreated {
			log.Infof("user: '%s' has already been created", u.UserID)
			return nil
		}
		return fmt.Errorf("could not create user: '%v'", err)
	}
	log.Infof("created the following user: '%v'", resp.Payload)

	return nil
}

func (r *Role) CreateRole() error {
	log.Infof("creating role: '%s'...", r.ID)

	client := r.Nexus3.Client()

	createRole := security_management_roles.CreateParams{Body: &r.RoleXORequest}
	createRole.WithTimeout(time.Second * 30)
	resp, err := client.SecurityManagementRoles.Create(&createRole)
	if err != nil {
		roleCreated, errRegex := regexp.MatchString("status 400", err.Error())
		if errRegex != nil {
			return err
		}
		if roleCreated {
			log.Infof("role: '%s' has already been created", r.Name)
			return nil
		}

		return fmt.Errorf("could not create role: '%v', perhaps the role already exists?", err)
	}
	log.Infof("created the following role: '%v'", resp.Payload)

	return nil
}

func (u *User) ChangePass() error {
	log.Infof("changing pass user: '%s'...", u.UserID)
	client := u.Nexus3.Client()

	changePass := security_management_users.ChangePasswordParams{Body: u.Password, UserID: u.UserID}
	changePass.WithTimeout(time.Second * 30)
	if err := client.SecurityManagementUsers.ChangePassword(&changePass); err != nil {
		passwordChanged, errRegex := regexp.MatchString("status 204", err.Error())
		if errRegex != nil {
			return err
		}

		if passwordChanged {
			log.Infof("user: '%s' pass has been changed", u.UserID)
			return nil
		}
		return err
	}

	return nil
}
