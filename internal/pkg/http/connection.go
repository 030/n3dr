package http

import (
	apiclient "github.com/030/n3dr/internal/goswagger/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Nexus3 struct {
	FQDN, Pass, User string
}

func (n *Nexus3) Client() *apiclient.Nexus3 {
	r := httptransport.New(n.FQDN, apiclient.DefaultBasePath, apiclient.DefaultSchemes)
	r.DefaultAuthentication = httptransport.BasicAuth(n.User, n.Pass)
	return apiclient.New(r, strfmt.Default)
}
