package main

import (
	"fmt"

	"github.com/030/n3dr/internal/app/n3dr/config/security"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configLDAPConnectionRetryDelaySeconds, configLDAPConnectionTimeoutSeconds, configLDAPMaxIncidentsCount, configLDAPPort int32
var configLDAPGroupsAsRoles bool
var configLDAPAuthPassword, configLDAPAuthScheme, configLDAPAuthUsername, configLDAPGroupType, configLDAPHost, configLDAPName, configLDAPProtocol, configLDAPSearchBase, configLDAPUserBaseDn, configLDAPUserEmailAddressAttribute, configLDAPUserIDAttribute, configLDAPUserMemberOfAttribute, configLDAPUserObjectClass, configLDAPUserRealNameAttribute string

// configLDAPCmd represents the configLDAP command
var configLDAPCmd = &cobra.Command{
	Use:   "configLDAP",
	Short: "configLDAP",
	Long:  `configLDAP`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configure LDAP")

		n := connection.Nexus3{FQDN: n3drURL, Pass: n3drPass, User: n3drUser}
		s := security.Security{Nexus3: n}
		m := models.CreateLdapServerXo{
			AuthPassword:                &configLDAPAuthPassword,
			AuthScheme:                  &configLDAPAuthScheme,
			AuthUsername:                configLDAPAuthUsername,
			ConnectionRetryDelaySeconds: &configLDAPConnectionRetryDelaySeconds,
			ConnectionTimeoutSeconds:    &configLDAPConnectionTimeoutSeconds,
			GroupType:                   &configLDAPGroupType,
			Host:                        &configLDAPHost,
			LdapGroupsAsRoles:           configLDAPGroupsAsRoles,
			MaxIncidentsCount:           &configLDAPMaxIncidentsCount,
			Name:                        &configLDAPName,
			Port:                        &configLDAPPort,
			Protocol:                    &configLDAPProtocol,
			SearchBase:                  &configLDAPSearchBase,
			UserBaseDn:                  configLDAPUserBaseDn,
			UserEmailAddressAttribute:   configLDAPUserEmailAddressAttribute,
			UserIDAttribute:             configLDAPUserIDAttribute,
			UserMemberOfAttribute:       &configLDAPUserMemberOfAttribute,
			UserRealNameAttribute:       configLDAPUserRealNameAttribute,
			UserObjectClass:             configLDAPUserObjectClass,
		}
		l := security.LDAPParams{Security: &s, CreateLdapServerXo: m}
		if err := l.LDAP(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configLDAPCmd)

	configLDAPCmd.Flags().Int32VarP(&configLDAPConnectionRetryDelaySeconds, "configLDAPConnectionRetryDelaySeconds", "", 300, "The number of seconds before retrying to connect")
	configLDAPCmd.Flags().Int32VarP(&configLDAPConnectionTimeoutSeconds, "configLDAPConnectionTimeoutSeconds", "", 30, "Number of seconds until timeout")
	configLDAPCmd.Flags().StringVarP(&configLDAPAuthScheme, "configLDAPAuthScheme", "", "SIMPLE", "The LDAP login scheme, e.g.: NONE, SIMPLE, DIGEST_MD5 or CRAM_MD5")
	configLDAPCmd.Flags().BoolVarP(&configLDAPGroupsAsRoles, "configLDAPGroupsAsRoles", "", true, "Enable groups as roles")
	configLDAPCmd.Flags().StringVarP(&configLDAPGroupType, "configLDAPGroupType", "", "DYNAMIC", "The group type, e.g. DYNAMIC or STATIC")
	configLDAPCmd.Flags().Int32VarP(&configLDAPMaxIncidentsCount, "configLDAPMaxIncidentsCount", "", 3, "The max number of retries to connect to the LDAP server")
	configLDAPCmd.Flags().Int32VarP(&configLDAPPort, "configLDAPPort", "", 389, "The config LDAP port")
	configLDAPCmd.Flags().StringVarP(&configLDAPProtocol, "configLDAPProtocol", "", "LDAP", "The config LDAP protocol, e.g.: LDAP or LDAPS")
	configLDAPCmd.Flags().StringVarP(&configLDAPSearchBase, "configLDAPSearchBase", "", "dc=example,dc=com", "The search base, e.g.: dc=example,dc=com")
	configLDAPCmd.Flags().StringVarP(&configLDAPUserBaseDn, "configLDAPUserBaseDn", "", "cn=users", "The user base dn, e.g. cn=users")
	configLDAPCmd.Flags().StringVarP(&configLDAPUserEmailAddressAttribute, "configLDAPUserEmailAddressAttribute", "", "mail", "The user email address attribute, e.g. email")
	configLDAPCmd.Flags().StringVarP(&configLDAPUserIDAttribute, "configLDAPUserIDAttribute", "", "sAMAccountName", "The user id attribute, e.g.: sAMAccountName")
	configLDAPCmd.Flags().StringVarP(&configLDAPUserMemberOfAttribute, "configLDAPUserMemberOfAttribute", "", "memberOf", "The user member of attribute, e.g.: memberOf")
	configLDAPCmd.Flags().StringVarP(&configLDAPUserRealNameAttribute, "configLDAPUserRealNameAttribute", "", "cn", "The user real name attribute, e.g. cn")
	configLDAPCmd.Flags().StringVarP(&configLDAPUserObjectClass, "configLDAPUserObjectClass", "", "user", "The user object class, e.g. user")

	configLDAPCmd.Flags().StringVarP(&configLDAPAuthPassword, "configLDAPAuthPassword", "", "", "The LDAP login password")
	if err := configLDAPCmd.MarkFlagRequired("configLDAPAuthPassword"); err != nil {
		log.Fatal(err)
	}
	configLDAPCmd.Flags().StringVarP(&configLDAPAuthUsername, "configLDAPAuthUsername", "", "", "The LDAP login name")
	if err := configLDAPCmd.MarkFlagRequired("configLDAPAuthUsername"); err != nil {
		log.Fatal(err)
	}
	configLDAPCmd.Flags().StringVarP(&configLDAPHost, "configLDAPHost", "", "", "The LDAP host")
	if err := configLDAPCmd.MarkFlagRequired("configLDAPHost"); err != nil {
		log.Fatal(err)
	}
	configLDAPCmd.Flags().StringVarP(&configLDAPName, "configLDAPName", "", "", "The name of the LDAP configuration")
	if err := configLDAPCmd.MarkFlagRequired("configLDAPName"); err != nil {
		log.Fatal(err)
	}
}
