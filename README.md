# n3dr

[![GoDoc Widget]][GoDoc]
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
[![codecov](https://codecov.io/gh/030/n3dr/branch/master/graph/badge.svg)](https://codecov.io/gh/030/n3dr)
[![BCH compliance](https://bettercodehub.com/edge/badge/030/n3dr?branch=master)](https://bettercodehub.com/results/030/n3dr)

The aims of the n3dr tool are:
 * to backup all artifacts from a certain Nexus maven repository.
 * to migrate all artifacts from NexusA to NexusB.

## Download and verify n3dr

```
curl -L https://github.com/030/n3dr/releases/download/3.1.1/n3dr-linux -o n3dr-linux
curl -L https://github.com/030/n3dr/releases/download/3.1.1/n3dr-linux.sha512.txt -o n3dr-linux.sha512.txt
sha512sum --check n3dr-linux.sha512.txt
chmod +x n3dr-linux
./n3dr-linux
```

## Check the help menu

```
user@computer:~/dev$ ./n3dr-linux -h
N3DR is a tool that is able to download all artifacts from
a certain Nexus3 repository.

Usage:
  n3dr [command]

Available Commands:
  backup       Backup all artifacts from a Nexus3 repository
  help         Help about any command
  repositories Count the number of repositories or return their names
  upload       Upload all artifacts to a specific Nexus3 repository

Flags:
  -v, --apiVersion string   The Nexus3 APIVersion, e.g. v1 or beta (default "v1")
  -d, --debug               Enable debug logging
  -h, --help                help for n3dr
  -p, --n3drPass string     The Nexus3 password
  -n, --n3drURL string      The Nexus3 URL
  -u, --n3drUser string     The Nexus3 user

Use "n3dr [command] --help" for more information about a command.
```

## Backup artifacts from a certain repository

All artifacts from a repository will be stored in a download folder when
the following command is run:

```
./n3dr-linux backup -u admin -n http://localhost:8081 -r maven-releases
```

## Backup all repositories

All artifacts from various repositories will be stored in a download
folder when the following command is issued:

```
./n3dr-linux repositories -u admin -n http://localhost:8081 -b
```

Note: a new folder will be created for every repository:

* download/maven-public
* download/maven-releases

## Upload all artifacts to a certain repository

It is possible to upload all JARs that reside in a folder by
running the following command:

```
./n3dr-linux upload -u admin -n http://localhost:8081 -r maven-public
```

------------------------------------------------------------------------------------

## How do the tests look like?

The tests start a nexus docker container. The tests will be started once
the docker container is running and fake artifacts have been uploaded. Finally,
all submitted artifacts will be downloaded.

## How to use this tool?

### Docker-compose

One could use [this docker-compose.yml](docker-compose.yml) and start a backup
of all repositories, after modifying it, and issue:

```
docker-compose up
```

### Password

Define the password in `~/.n3dr.yaml`:

```
---
n3drPass: admin123
```

and set the permissions to 400 by issuing:

```
chmod 400 ~/.n3dr.yaml
```

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
  backup       Backup all artifacts from a Nexus3 repository
  help         Help about any command
  repositories Count the number of repositories or return their names

Flags:
  -h, --help   help for n3dr

Use "n3dr [command] --help" for more information about a command.
```

### Backup

The backup command will backup all artifacts that reside in a Nexus maven
repository.

```
[user@localhost n3dr]$ ./n3dr backup -h
Use this command in order to backup all artifacts that
reside in a certain Nexus3 repository

Usage:
  n3dr backup [flags]

Flags:
  -h, --help              help for backup
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

#### Download artifacts from all repositories

```
go run main.go repositories --download --n3drURL http://localhost:9999
```

[![dockeri.co](https://dockeri.co/image/utrecht/n3dr)](https://hub.docker.com/r/utrecht/n3dr)

```
docker run --rm -e HTTPS_PROXY=some-proxy \
           -v ~/.n3dr.yaml:/home/n3dr/.n3dr.yaml \
           -v ${PWD}/nexus3backup:/download utrecht/n3dr:3.0.0 \
           backup -p pass -r maven-releases \
           -n https://some-nexus-repo -u admin
```

### Difference with equivalent tools

There is a number of equivalent tools:

* https://github.com/RiotGamesMinions/nexus_cli
* https://github.com/packagemgmt/repositorytools
* https://github.com/thiagofigueiro/nexus3-cli

The difference is that n3dr is able to download artifacts from all Nexus3
repositories.

[GoDoc]: https://godoc.org/github.com/030/n3dr
[GoDoc Widget]: https://godoc.org/github.com/030/n3dr?status.svg