package user

import (
	"fmt"
	"regexp"
	"time"

	"github.com/030/n3dr/internal/goswagger/client/security_management_users"
	"github.com/030/n3dr/internal/goswagger/models"
	"github.com/030/n3dr/internal/pkg/http"
	log "github.com/sirupsen/logrus"
)

type User struct {
	http.Nexus3
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
		return fmt.Errorf("could not create user: '%v', perhaps the user already exists?", err)
	}
	log.Infof("created the following user: '%v'", resp.Payload)

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
