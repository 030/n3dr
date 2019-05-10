# nexus3-cli

[![Build Status](https://travis-ci.org/030/nexus3-cli.svg?branch=master)](https://travis-ci.org/030/nexus3-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/030/nexus3-cli)](https://goreportcard.com/report/github.com/030/nexus3-cli)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=030_nexus3-cli&metric=bugs)](https://sonarcloud.io/dashboard?id=030_nexus3-cli)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=030_nexus3-cli&metric=code_smells)](https://sonarcloud.io/dashboard?id=030_nexus3-cli)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=030_nexus3-cli&metric=coverage)](https://sonarcloud.io/dashboard?id=030_nexus3-cli)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=030_nexus3-cli&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=030_nexus3-cli)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=030_nexus3-cli&metric=ncloc)](https://sonarcloud.io/dashboard?id=030_nexus3-cli)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_nexus3-cli&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=030_nexus3-cli)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=030_nexus3-cli&metric=alert_status)](https://sonarcloud.io/dashboard?id=030_nexus3-cli)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_nexus3-cli&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=030_nexus3-cli)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=030_nexus3-cli&metric=security_rating)](https://sonarcloud.io/dashboard?id=030_nexus3-cli)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=030_nexus3-cli&metric=sqale_index)](https://sonarcloud.io/dashboard?id=030_nexus3-cli)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=030_nexus3-cli&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=030_nexus3-cli)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/2762/badge)](https://bestpractices.coreinfrastructure.org/projects/2762)

The aims of the nexus3-cli tool are:
 * to backup all artifacts from a certain Nexus maven repository.
 * to migrate all artifacts from NexusA to NexusB.

## How do the tests look like?

The tests start a nexus docker container. The tests will be started once
the docker container is running and fake artifacts have been uploaded. Finally,
all submitted artifacts will be downloaded.

## How to use this tool?

### Help

In order to read the help menu, one has to run:

```
go run main.go -h
```

The output will look as follows:

```
Usage of /tmp/go-build840407725/b001/exe/main:
  -nexus3URL string
        The Nexus3URL (default "http://localhost:8081")
  -nexus3pass string
        The Nexus password (default "admin123")
  -nexus3repo string
        The Nexus3 repository (default "maven-releases")
  -nexus3user string
        The Nexus user (default "admin")
exit status 2
```

### Download

The download command will download all artifacts that reside in a Nexus maven
repository.

### Upload

The upload command will upload all artifacts to a Nexus maven repository.
