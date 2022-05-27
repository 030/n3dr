package connection

import (
	apiclient "github.com/030/n3dr/internal/goswagger/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

type Nexus3 struct {
	AwsBucket, AwsId, AwsRegion, AwsSecret, BasePathPrefix, DockerHost, DownloadDirName, DownloadDirNameZip, FQDN, Pass, User string
	DockerPort                                                                                                                int32
	DockerPortSecure, HTTPS                                                                                                   bool
}

func (n *Nexus3) Client() *apiclient.Nexus3 {
	schemes := apiclient.DefaultSchemes
	if n.HTTPS {
		schemes = []string{"http", "https"}
	}
	basePath := apiclient.DefaultBasePath
	if n.BasePathPrefix != "" {
		log.Infof("adding '%s' as a prefix to the basePath", n.BasePathPrefix)
		basePath = n.BasePathPrefix + "/" + apiclient.DefaultBasePath
	}
	r := httptransport.New(n.FQDN, basePath, schemes)
	r.DefaultAuthentication = httptransport.BasicAuth(n.User, n.Pass)
	return apiclient.New(r, strfmt.Default)
}
