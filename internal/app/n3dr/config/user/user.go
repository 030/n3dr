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
	exists, err := u.exists()
	if err != nil {
		return err
	}
	if !exists {
		status := "active"
		u.Status = &status

		client, err := u.Nexus3.Client()
		if err != nil {
			return err
		}

		createUser := security_management_users.CreateUserParams{Body: &u.APICreateUser}
		createUser.WithTimeout(time.Second * 30)
		resp, err := client.SecurityManagementUsers.CreateUser(&createUser)
		if err != nil {
			return fmt.Errorf("could not create user: '%w'", err)
		}
		log.Infof("created the following user: '%v'", resp.Payload)
		return nil
	}
	log.Infof("user: '%s' exists already", u.APICreateUser.UserID)

	return nil
}

func (u *User) exists() (bool, error) {
	exists := false

	log.Info("checking whether user exists...")

	client, err := u.Nexus3.Client()
	if err != nil {
		return exists, err
	}

	getUsersParams := security_management_users.GetUsersParams{}
	getUsersParams.WithTimeout(time.Second * 30)
	resp, err := client.SecurityManagementUsers.GetUsers(&getUsersParams)
	if err != nil {
		return exists, fmt.Errorf("could not check whether user: '%s' exists. Error: '%w'", u.APICreateUser.UserID, err)
	}

	users := resp.GetPayload()
	for _, user := range users {
		log.Infof("found user: '%+v'", user)
		log.Infof("looking for user: '%+v'", u.APICreateUser)

		if user.UserID == u.APICreateUser.UserID {
			log.Infof("user: '%s' exists", u.APICreateUser.UserID)
			exists = true
		}
	}

	return exists, nil
}

func (r *Role) CreateRole() error {
	log.Infof("creating role: '%s' if it does not exist...", r.ID)

	exists, err := r.checkWhetherRoleExists()
	if err != nil {
		return err
	}
	if !exists {
		client, err := r.Nexus3.Client()
		if err != nil {
			return err
		}

		createRole := security_management_roles.CreateParams{Body: &r.RoleXORequest}
		createRole.WithTimeout(time.Second * 30)
		resp, err := client.SecurityManagementRoles.Create(&createRole)
		if err != nil {
			return fmt.Errorf("could not create role. Error: '%w'", err)
		}

		log.Infof("created the following role: '%+v'", resp.Payload)
		return nil
	}
	log.Infof("role: '%s' exists already", r.RoleXORequest.ID)

	return nil
}

func (r *Role) checkWhetherRoleExists() (bool, error) {
	exists := false

	log.Infof("checking whether role id: '%s' exists...", r.ID)
	client, err := r.Nexus3.Client()
	if err != nil {
		return exists, err
	}

	getRoleParams := security_management_roles.GetRoleParams{ID: r.RoleXORequest.ID}
	getRoleParams.WithTimeout(time.Second * 30)
	resp, err := client.SecurityManagementRoles.GetRole(&getRoleParams)
	if err != nil {
		roleDoesNotExist, errRegex := regexp.MatchString(`\]\[404\] getRoleNotFound`, err.Error())
		if errRegex != nil {
			return exists, err
		}
		if roleDoesNotExist {
			log.Infof("role id: '%s' does not exist...", r.ID)
			return exists, nil
		}

		return exists, fmt.Errorf("could not get role: '%s'. Error: '%w'", getRoleParams.ID, err)
	}

	log.Infof("role: '%s' exists", resp.GetPayload().ID)
	exists = true

	return exists, nil
}

func (u *User) ChangePass() error {
	log.Infof("changing pass user: '%s'...", u.UserID)

	client, err := u.Nexus3.Client()
	if err != nil {
		return err
	}

	changePass := security_management_users.ChangePasswordParams{Body: u.Password, UserID: u.UserID}
	changePass.WithTimeout(time.Second * 30)

	if err := client.SecurityManagementUsers.ChangePassword(&changePass); err != nil {
		passwordChanged, errRegex := regexp.MatchString("status 204", err.Error())
		if errRegex != nil {
			return err
		}

		if !passwordChanged {
			return fmt.Errorf("password of userID: '%s' did not change. Error: '%w'", u.UserID, err)
		}
	}

	log.Infof("user: '%s' pass has been changed", u.UserID)

	return nil
}
