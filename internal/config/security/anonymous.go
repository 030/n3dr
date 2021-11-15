package security

import (
	"fmt"
	"time"

	"github.com/030/n3dr/internal/goswagger/client/security_management_anonymous_access"
	"github.com/030/n3dr/internal/goswagger/models"
	"github.com/030/n3dr/internal/pkg/http"
	log "github.com/sirupsen/logrus"
)

type Security struct {
	http.Nexus3
}

func (s *Security) Anonymous(enabled bool) error {
	aasx := models.AnonymousAccessSettingsXO{Enabled: enabled}

	log.Info("changing anonymous access")

	client := s.Nexus3.Client()

	anonymousAccess := security_management_anonymous_access.UpdateParams{Body: &aasx}
	anonymousAccess.WithTimeout(time.Second * 30)

	resp, err := client.SecurityManagementAnonymousAccess.Update(&anonymousAccess)
	if err != nil {
		return fmt.Errorf("could not change anonymous access mode: '%v'", err)
	}
	if resp.Payload.Enabled {
		log.Info("anonymous access enabled")
	} else {
		log.Info("anonymous access disabled")
	}

	return nil
}
