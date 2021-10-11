/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"time"

	apiclient "github.com/030/n3dr/internal/go-swagger/client"
	"github.com/030/n3dr/internal/go-swagger/client/security_management_users"
	"github.com/030/n3dr/internal/go-swagger/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")

		// n := security_management_users.Client{}
		// n.CreateUser(&u)

		// var myRoundTripper http.RoundTripper = createRoundTripper()

		// transport.Transport = myRoundTripper
		// todoListClient := client.New(transport, nil)

		// r := httptransport.New(client.DefaultHost, client.DefaultBasePath, client.DefaultSchemes)
		// r := httptransport.New("localhost:9999", "/service/rest", []string{"http"})
		// r := httptransport.New("localhost:9999", "/", []string{"http"})
		// r.DefaultAuthentication = httptransport.BasicAuth("admin", "fce8bc4c-cd3a-46ad-8555-6300e8fd67de")
		// client2 := client.New(r, strfmt.Default)

		// a, err := client2.SecurityManagementUsers.CreateUser(&u)
		// fmt.Println(a, err)

		r := httptransport.New("localhost:9999", apiclient.DefaultBasePath, apiclient.DefaultSchemes)
		r.DefaultAuthentication = httptransport.BasicAuth("admin", "fce8bc4c-cd3a-46ad-8555-6300e8fd67de")
		client := apiclient.New(r, strfmt.Default)
		//
		status := "active"
		roles := []string{"nx-admin"}
		b := models.APICreateUser{EmailAddress: "piet.snot@bladibla.bladibla", FirstName: "Piet", LastName: "Snot", Password: "wiewiewie123", UserID: "piet2", Status: &status, Roles: roles}
		u := security_management_users.CreateUserParams{Body: &b}
		u.WithTimeout(time.Second * 30)
		resp, err := client.SecurityManagementUsers.CreateUser(&u)
		//
		//
		// u := security_management_users.GetUsersParams{}
		//

		// resp, err := client.SecurityManagementUsers.GetUsers(&u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v\n", resp.Payload)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
