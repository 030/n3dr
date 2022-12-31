package name

import (
	"fmt"

	"github.com/030/n3dr/internal/app/n3dr/artifactsv2/artifacts"
	"github.com/030/n3dr/internal/app/n3dr/connection"
)

func (n *Nexus3) Repositories() error {
	cn := connection.Nexus3{BasePathPrefix: n.BasePathPrefix, FQDN: n.FQDN, DownloadDirName: n.DownloadDirName, Pass: n.Pass, User: n.User, HTTPS: n.HTTPS}
	a := artifacts.Nexus3{Nexus3: &cn}

	repos, err := a.Repos()
	if err != nil {
		return err
	}

	for _, repo := range repos {
		fmt.Println(repo.Name)
	}

	return nil
}
