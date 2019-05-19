# n3dr

[![Build Status](https://travis-ci.org/030/n3dr.svg?branch=master)](https://travis-ci.org/030/n3dr)
[![Go Report Card](https://goreportcard.com/badge/github.com/030/n3dr)](https://goreportcard.com/report/github.com/030/n3dr)
![DevOps SE Questions](https://img.shields.io/stackexchange/devops/t/n3dr.svg)
![Docker Pulls](https://img.shields.io/docker/pulls/utrecht/n3dr.svg)
![Issues](https://img.shields.io/github/issues-raw/030/n3dr.svg)
![Pull requests](https://img.shields.io/github/issues-pr-raw/030/n3dr.svg)
![Total downloads](https://img.shields.io/github/downloads/030/n3dr/total.svg)
![License](https://img.shields.io/github/license/030/n3dr.svg)
![Repository Size](https://img.shields.io/github/repo-size/030/n3dr.svg)
![Contributors](https://img.shields.io/github/contributors/030/n3dr.svg)
![Commit activity](https://img.shields.io/github/commit-activity/m/030/n3dr.svg)
![Last commit](https://img.shields.io/github/last-commit/030/n3dr.svg)
![Release date](https://img.shields.io/github/release-date/030/n3dr.svg)
![Latest Production Release Version](https://img.shields.io/github/release/030/n3dr.svg)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=030_n3dr&metric=bugs)](https://sonarcloud.io/dashboard?id=030_n3dr)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=030_n3dr&metric=code_smells)](https://sonarcloud.io/dashboard?id=030_n3dr)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=030_n3dr&metric=coverage)](https://sonarcloud.io/dashboard?id=030_n3dr)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=030_n3dr&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=030_n3dr)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=030_n3dr&metric=ncloc)](https://sonarcloud.io/dashboard?id=030_n3dr)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_n3dr&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=030_n3dr)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=030_n3dr&metric=alert_status)](https://sonarcloud.io/dashboard?id=030_n3dr)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_n3dr&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=030_n3dr)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=030_n3dr&metric=security_rating)](https://sonarcloud.io/dashboard?id=030_n3dr)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=030_n3dr&metric=sqale_index)](https://sonarcloud.io/dashboard?id=030_n3dr)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=030_n3dr&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=030_n3dr)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/2810/badge)](https://bestpractices.coreinfrastructure.org/projects/2810)

The aims of the n3dr tool are:
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
./n3dr
```

The output will look as follows:

```
N3DR is a tool that is able to download all artifacts from
a certain Nexus3 repository.

Usage:
  n3dr [command]

Available Commands:
  download     Download all artifacts from a Nexus3 repository
  help         Help about any command
  repositories Count the number of repositories or return their names

Flags:
  -h, --help   help for n3dr

Use "n3dr [command] --help" for more information about a command.
```

### Download

The download command will download all artifacts that reside in a Nexus maven
repository.

```
[user@localhost n3dr]$ ./n3dr download -h
Use this command in order to download all artifacts that
reside in a certain Nexus3 repository

Usage:
  n3dr download [flags]

Flags:
  -h, --help              help for download
  -p, --n3drPass string   The Nexus3 password (default "admin123")
  -r, --n3drRepo string   The Nexus3 repository (default "maven-releases")
  -n, --n3drURL string    The Nexus3 URL (default "http://localhost:8081")
  -u, --n3drUser string   The Nexus3 user (default "admin")
```

### Repositories

In order to get an overview of all repositories that are available in a certain
Nexus3 instance, one could use the following commands:

```
Count the number of repositories or
count the total

Usage:
  n3dr repositories [flags]

Flags:
  -c, --count   Count the number of repositories
  -h, --help    help for repositories
  -n, --names   Print all repository names
```

[![dockeri.co](https://dockeri.co/image/utrecht/n3dr)](https://hub.docker.com/r/utrecht/n3dr)

```
docker run utrecht/n3dr:2.0.0 -h
```
