package connection

import (
	apiclient "github.com/030/n3dr/internal/goswagger/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Nexus3 struct {
	DownloadDirName, FQDN, Pass, User string
	HTTPS                             bool
}

func (n *Nexus3) Client() *apiclient.Nexus3 {
	schemes := apiclient.DefaultSchemes
	if n.HTTPS {
		schemes = []string{"http", "https"}
	}
	r := httptransport.New(n.FQDN, apiclient.DefaultBasePath, schemes)
	r.DefaultAuthentication = httptransport.BasicAuth(n.User, n.Pass)
	return apiclient.New(r, strfmt.Default)
}
