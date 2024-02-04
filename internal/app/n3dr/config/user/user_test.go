//go:build integration

package user

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const integrationTestAdminPassword = "someIntegrationTestPassword123!"

var hostAndPort, initialAdminPassword string

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "sonatype/nexus3",
		Tag:        "3.64.0",
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort = resource.GetHostPort("8081/tcp")
	log.Info("hostAndPort", hostAndPort)

	client := http.Client{}

	pool.MaxWait = time.Minute * 3
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet.
	if err = pool.Retry(func() error {
		// Check whether Nexus3 is ready and writable
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%s/service/rest/v1/status/writable", resource.GetPort("8081/tcp")), nil)
		if err != nil {
			return err
		}

		_, err = client.Do(req)

		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	var out bytes.Buffer
	resource.Exec([]string{"cat", "/nexus-data/admin.password"}, dockertest.ExecOptions{
		StdOut: &out,
		StdErr: &out,
	})
	initialAdminPassword = out.String()

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestChangePass(t *testing.T) {
	https := false
	c := connection.Nexus3{
		HTTPS: &https,
		User:  "admin",
	}
	macu := models.APICreateUser{
		EmailAddress: "admin@example.org",
		FirstName:    "admin",
		UserID:       "admin",
		LastName:     "admin",
		Password:     initialAdminPassword,
	}

	tests := []struct {
		adminPassword, expectedErrorString, name, passwordOfUserIdThatHasToBeChanged string
		unsetHostAndPort                                                             bool
	}{
		{
			adminPassword:                      initialAdminPassword,
			expectedErrorString:                "",
			name:                               "If admin password is valid then it should be possible to change it with the same password that is valid.",
			passwordOfUserIdThatHasToBeChanged: "admin",
			unsetHostAndPort:                   false,
		},
		{
			adminPassword:                      "incorrectPassword",
			expectedErrorString:                "password of userID: 'admin' did not change. Error: 'response status code does not match any response statuses defined for this endpoint in the swagger spec (status 401): {}'",
			name:                               "Test that change of password will fail if it is not possible to login to Nexus3 due to an incorrect password.",
			passwordOfUserIdThatHasToBeChanged: "admin",
			unsetHostAndPort:                   false,
		},
		{
			adminPassword:                      initialAdminPassword,
			expectedErrorString:                "password of userID: 'someUserIDThatDoesNotExist' did not change. Error: '[PUT /v1/security/users/{userId}/change-password][404] changePasswordNotFound '",
			name:                               "Test that password change will fail if userID does not exist.",
			passwordOfUserIdThatHasToBeChanged: "someUserIDThatDoesNotExist",
			unsetHostAndPort:                   false,
		},
		{
			adminPassword:                      initialAdminPassword,
			expectedErrorString:                "Key: 'Nexus3.FQDN' Error:Field validation for 'FQDN' failed on the 'required' tag",
			name:                               "Test that password change will fail if host and port have not been set.",
			passwordOfUserIdThatHasToBeChanged: "someUserIDThatDoesNotExist",
			unsetHostAndPort:                   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c.Pass = test.adminPassword
			c.FQDN = hostAndPort
			if test.unsetHostAndPort {
				c.FQDN = ""
			}

			macu.UserID = test.passwordOfUserIdThatHasToBeChanged
			u := User{c, macu}
			err := u.ChangePass()

			if test.expectedErrorString == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.expectedErrorString)
			}
		})
	}
}

func TestCreateRole(t *testing.T) {
	https := false
	c := connection.Nexus3{
		HTTPS: &https,
		User:  "admin",
	}
	mrxr := models.RoleXORequest{
		Name: "nx-upload",
	}

	tests := []struct {
		adminPassword, expectedErrorString, name, passwordOfUserIdThatHasToBeChanged, roleID string
		rolePrivileges                                                                       []string
		unsetHostAndPort                                                                     bool
	}{
		{
			adminPassword:                      initialAdminPassword,
			expectedErrorString:                "",
			name:                               "Create a role.",
			passwordOfUserIdThatHasToBeChanged: "admin",
			unsetHostAndPort:                   false,
			roleID:                             "nx-upload",
			rolePrivileges: []string{
				"nx-repository-view-*-*-add",
			},
		},
		{
			adminPassword:                      initialAdminPassword,
			expectedErrorString:                "",
			name:                               "Trying to create a role that exists should be handled by the code and return without any error.",
			passwordOfUserIdThatHasToBeChanged: "admin",
			unsetHostAndPort:                   false,
			roleID:                             "nx-upload",
			rolePrivileges: []string{
				"nx-repository-view-*-*-edit",
			},
		},
		{
			adminPassword:                      initialAdminPassword,
			expectedErrorString:                "could not create role. Response: '<nil>'. Error: 'response status code does not match any response statuses defined for this endpoint in the swagger spec (status 400): {}'",
			name:                               "Trying to create a role with a privilege that does not exist.",
			passwordOfUserIdThatHasToBeChanged: "admin",
			unsetHostAndPort:                   false,
			roleID:                             "role-that-does-not-exist",
			rolePrivileges: []string{
				"a-privilege-that-does-not-exist",
			},
		},
		{
			adminPassword:                      "incorrectPassword",
			expectedErrorString:                "could not get role: 'role-that-does-not-exist'. Error: 'response status code does not match any response statuses defined for this endpoint in the swagger spec (status 401): {}'",
			name:                               "Test that role creatiom will fail if it is not possible to login to Nexus3 due to an incorrect password.",
			passwordOfUserIdThatHasToBeChanged: "admin",
			unsetHostAndPort:                   false,
			roleID:                             "role-that-does-not-exist",
			rolePrivileges: []string{
				"a-privilege-that-does-not-exist",
			},
		},
		{
			adminPassword:                      "incorrectPassword",
			expectedErrorString:                "Key: 'Nexus3.FQDN' Error:Field validation for 'FQDN' failed on the 'required' tag",
			name:                               "Test that role creatiom will fail if it is not possible to login to Nexus3 due to an incorrect password.",
			passwordOfUserIdThatHasToBeChanged: "admin",
			unsetHostAndPort:                   true,
			roleID:                             "role-that-does-not-exist",
			rolePrivileges: []string{
				"a-privilege-that-does-not-exist",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c.Pass = test.adminPassword
			c.FQDN = hostAndPort
			if test.unsetHostAndPort {
				c.FQDN = ""
			}

			r := Role{c, mrxr}
			r.RoleXORequest.ID = test.roleID
			r.RoleXORequest.Privileges = test.rolePrivileges
			err := r.CreateRole()

			if test.expectedErrorString == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.expectedErrorString)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	https := false
	c := connection.Nexus3{
		HTTPS: &https,
		User:  "admin",
	}

	tests := []struct {
		adminPassword, expectedErrorString, name, passwordOfUserIdThatHasToBeChanged, roleID string
		rolePrivileges                                                                       []string
		unsetHostAndPort                                                                     bool
		macu                                                                                 models.APICreateUser
	}{
		{
			adminPassword:       initialAdminPassword,
			expectedErrorString: "",
			name:                "Create a user that exists.",
			unsetHostAndPort:    false,
			macu: models.APICreateUser{
				EmailAddress: "admin@example.org",
				FirstName:    "admin",
				LastName:     "admin",
				Password:     "some-password1234!",
				UserID:       "admin",
				Roles:        []string{"nx-anonymous"},
			},
		},
		{
			adminPassword:       initialAdminPassword,
			expectedErrorString: "",
			name:                "Create a user that does not exist.",
			unsetHostAndPort:    false,
			macu: models.APICreateUser{
				EmailAddress: "admin42@example.org",
				FirstName:    "admin42",
				LastName:     "admin42",
				Password:     "some-password1234!",
				UserID:       "admin42",
				Roles:        []string{"nx-anonymous"},
			},
		},
		{
			adminPassword:       initialAdminPassword,
			expectedErrorString: "",
			name:                "Create a user that has just been created for the second time.",
			unsetHostAndPort:    false,
			macu: models.APICreateUser{
				EmailAddress: "admin42@example.org",
				FirstName:    "admin42",
				LastName:     "admin42",
				Password:     "some-password1234!",
				UserID:       "admin42",
				Roles:        []string{"nx-anonymous"},
			},
		},
		{
			adminPassword:       initialAdminPassword,
			expectedErrorString: "could not create user: '[POST /v1/security/users][400] createUserBadRequest '",
			name:                "Create a new user without specifying a role should return a 400 bad request as roles are required.",
			unsetHostAndPort:    false,
			macu: models.APICreateUser{
				EmailAddress: "admin43@example.org",
				FirstName:    "admin43",
				LastName:     "admin43",
				Password:     "some-password1234!",
				UserID:       "admin43",
			},
		},
		{
			adminPassword:       initialAdminPassword,
			expectedErrorString: "Key: 'Nexus3.FQDN' Error:Field validation for 'FQDN' failed on the 'required' tag",
			name:                "Create a new user, but this should fail as the Nexus3 URL is invalid.",
			unsetHostAndPort:    true,
			macu: models.APICreateUser{
				EmailAddress: "admin43@example.org",
				FirstName:    "admin43",
				LastName:     "admin43",
				Password:     "some-password1234!",
				UserID:       "admin43",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c.Pass = test.adminPassword
			c.FQDN = hostAndPort
			if test.unsetHostAndPort {
				c.FQDN = ""
			}

			u := User{c, test.macu}
			err := u.Create()

			if test.expectedErrorString == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.expectedErrorString)
			}
		})
	}
}
