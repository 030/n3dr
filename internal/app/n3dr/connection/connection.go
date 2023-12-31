package connection

import (
	apiclient "github.com/030/n3dr/internal/app/n3dr/goswagger/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

type Nexus3 struct {
	AwsBucket, AwsID, AwsRegion, AwsSecret, BasePathPrefix, DockerHost, DownloadDirName, DownloadDirNameZip, Pass, Regex, RepoName, User       string
	DockerPort                                                                                                                                 int32
	DockerPortSecure, SkipErrors, StrictContentTypeValidation, WithoutWaitGroups, WithoutWaitGroupArtifacts, WithoutWaitGroupRepositories, ZIP bool
	HTTPS                                                                                                                                      *bool  `validate:"required"`
	FQDN                                                                                                                                       string `validate:"required"`
}

func (n *Nexus3) Client() (*apiclient.Nexus3, error) {
	if err := validator.New().Struct(n); err != nil {
		return nil, err
	}

	schemes := apiclient.DefaultSchemes
	if *n.HTTPS {
		schemes = []string{"http", "https"}
	}
	basePath := apiclient.DefaultBasePath
	if n.BasePathPrefix != "" {
		log.Tracef("adding '%s' as a prefix to the basePath", n.BasePathPrefix)
		basePath = n.BasePathPrefix + "/" + apiclient.DefaultBasePath
	}
	r := httptransport.New(n.FQDN, basePath, schemes)
	r.DefaultAuthentication = httptransport.BasicAuth(n.User, n.Pass)

	return apiclient.New(r, strfmt.Default), nil
}
