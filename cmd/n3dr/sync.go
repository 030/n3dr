package main

import (
	"fmt"
	"sync"

	"github.com/030/n3dr/internal/artifactsv2"
	"github.com/030/n3dr/internal/artifactsv2/upload"
	"github.com/030/n3dr/internal/pkg/connection"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var otherNexus3URLs, otherNexus3Users, otherNexus3Passwords []string

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "sync",
	Long:  `sync`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sync called")
		numberOfURLs := len(otherNexus3URLs)
		numberOfPasswords := len(otherNexus3Passwords)
		numberOfUsers := len(otherNexus3Users)
		if numberOfURLs != numberOfPasswords || numberOfURLs != numberOfUsers {
			log.Fatal("incorrect number of elements. Ensure that the number of elements is identical")
		}

		n := connection.Nexus3{FQDN: n3drURL, Pass: n3drPass, User: n3drUser, DownloadDirName: downloadDirName}
		a := artifactsv2.Nexus3{Nexus3: &n}
		if err := a.Backup(); err != nil {
			log.Fatal(err)
		}

		var errs []error
		var wg sync.WaitGroup
		for i, otherNexus3User := range otherNexus3Users {
			n.User = otherNexus3User
			n.Pass = otherNexus3Passwords[i]
			n.FQDN = otherNexus3URLs[i]

			wg.Add(1)
			go func(n connection.Nexus3) {
				defer wg.Done()

				u := upload.Nexus3{Nexus3: &n}
				if err := u.Upload(); err != nil {
					errs = append(errs, err)
					return
				}
			}(n)
		}
		wg.Wait()
		for _, err := range errs {
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.PersistentFlags().StringSliceVarP(&otherNexus3URLs, "otherNexus3URLs", "", nil, "specify the other Nexus3 URLs in a comma separated list")
	if err := syncCmd.MarkPersistentFlagRequired("otherNexus3URLs"); err != nil {
		log.Fatal(err)
	}
	syncCmd.PersistentFlags().StringSliceVarP(&otherNexus3Users, "otherNexus3Users", "", nil, "specify the other Nexus3 users in a comma separated list")
	if err := syncCmd.MarkPersistentFlagRequired("otherNexus3Users"); err != nil {
		log.Fatal(err)
	}
	syncCmd.PersistentFlags().StringSliceVarP(&otherNexus3Passwords, "otherNexus3Passwords", "", nil, "specify the other Nexus3 passwords in a comma separated list")
	if err := syncCmd.MarkPersistentFlagRequired("otherNexus3Passwords"); err != nil {
		log.Fatal(err)
	}
}
