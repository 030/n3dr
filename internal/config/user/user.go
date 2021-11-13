package user

import (
	"fmt"
	"time"

	"github.com/030/n3dr/internal/goswagger/client/security_management_users"
	"github.com/030/n3dr/internal/goswagger/models"
	"github.com/030/n3dr/internal/pkg/http"
)

type User struct {
	http.Nexus3
	models.APICreateUser
}

func (u *User) Create() error {
	status := "active"
	u.Status = &status

	n := u.Nexus3
	client := n.Client()

	createUser := security_management_users.CreateUserParams{Body: &u.APICreateUser}
	createUser.WithTimeout(time.Second * 30)
	resp, err := client.SecurityManagementUsers.CreateUser(&createUser)
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", resp.Payload)

	return nil
}
