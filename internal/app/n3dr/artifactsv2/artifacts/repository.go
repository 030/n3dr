package artifacts

import (
	"fmt"
	"time"

	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/client/repository_management"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
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
		return nil, fmt.Errorf("cannot get repository names: '%w'", err)
	}

	return resp.Payload, nil
}
