package count

import (
	"fmt"

	"github.com/030/n3dr/internal/app/n3dr/artifactsv2/artifacts"
	"github.com/030/n3dr/internal/app/n3dr/connection"
)

func (n *Nexus3) Repositories() error {
	cn := connection.Nexus3{BasePathPrefix: n.BasePathPrefix, DownloadDirName: n.DownloadDirName, FQDN: n.FQDN, HTTPS: n.HTTPS, Pass: n.Pass, User: n.User}
	a := artifacts.Nexus3{Nexus3: &cn}

	repos, err := a.Repos()
	if err != nil {
		return err
	}

	fmt.Println(len(repos))

	return nil
}
