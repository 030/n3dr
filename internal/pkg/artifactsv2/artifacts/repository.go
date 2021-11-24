package artifacts

import (
	"fmt"
	"time"

	"github.com/030/n3dr/internal/goswagger/client/repository_management"
	"github.com/030/n3dr/internal/goswagger/models"
	"github.com/030/n3dr/internal/pkg/connection"
)

type Nexus3 struct {
	*connection.Nexus3
}

func (n *Nexus3) Repos() ([]*models.AbstractAPIRepository, error) {
	client := n.Nexus3.Client()
	r := repository_management.GetRepositoriesParams{}
	r.WithTimeout(time.Second * 30)
	resp, err := client.RepositoryManagement.GetRepositories(&r)
	if err != nil {
		return nil, fmt.Errorf("cannot get repository names: '%v'", err)
	}
	return resp.Payload, nil
}
