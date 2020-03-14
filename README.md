# n3dr

[![GoDoc Widget]][GoDoc]
[![Build Status](https://travis-ci.org/030/n3dr.svg?branch=master)](https://travis-ci.org/030/n3dr)
[![Go Report Card](https://goreportcard.com/badge/github.com/030/n3dr)](https://goreportcard.com/report/github.com/030/n3dr)
![DevOps SE Questions](https://img.shields.io/stackexchange/devops/t/n3dr.svg)
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
[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-web.svg)](https://golangci.com/r/github.com/030/n3dr)

<a href="https://n3dr.releasesoftwaremoreoften.com"><img src="https://github.com/030/n3dr/raw/master/logo/logo.png" width="100"></a>

The aims of the n3dr tool are:
 * to backup all artifacts from a certain Nexus maven repository.
 * to migrate all artifacts from NexusA to NexusB.

## Installation

```
curl -L https://github.com/030/n3dr/releases/download/z.y.z/n3dr-linux -o n3dr-linux
curl -L https://github.com/030/n3dr/releases/download/z.y.z/n3dr-linux.sha512.txt -o n3dr-linux.sha512.txt
sha512sum --check n3dr-linux.sha512.txt
chmod +x n3dr-linux
./n3dr-linux
```

### Debian

```
VERSION=z.y.z && \
curl -L https://github.com/030/n3dr/releases/download/${VERSION}/n3dr_${VERSION}-0.deb -o n3dr.deb && \
sudo apt -y install ./n3dr.deb
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
      --insecureSkipVerify  Skip repository certificate check
  -p, --n3drPass string     The Nexus3 password
  -n, --n3drURL string      The Nexus3 URL
  -u, --n3drUser string     The Nexus3 user
  -z, --zip                 Add downloaded artifacts to a ZIP archive

Use "n3dr [command] --help" for more information about a command.
```

## Store the password in a read-only file

Define the password in `~/.n3dr.yaml`:

```
---
n3drPass: admin123
```

and set the permissions to 400 by issuing:

```
chmod 400 ~/.n3dr.yaml
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

## Add all downloaded archives to a ZIP archive

In order to add all archives to a zip archive, one has to use the --zip or -z flag.

## Upload all artifacts to a certain repository

It is possible to upload all JARs that reside in a folder by
running the following command:

```
./n3dr-linux upload -u admin -n http://localhost:8081 -r maven-public
```

## "Clone" the content of a repository in a different Nexus 3 server in a different repository 

These are the basic steps to "clone" and eventually rename the content of a
repository from one nexus3 server to another one:

```
n3dr backup -u <source-nexus3-user> -n <source-nexus3-server-url> -r <repo-source-name>
cd download
mv <repo-source-name> <repo-target-name>
n3dr upload -u <target-nexus3-user> -n <target-nexus3-server-url> -r <repo-target-name>
```

## Rationale for N3DR

Although there is a number of equivalent tools:

* https://github.com/RiotGamesMinions/nexus_cli
* https://github.com/packagemgmt/repositorytools
* https://github.com/thiagofigueiro/nexus3-cli

None of them seems to be able to backup all repositories by running
a single command.

[GoDoc]: https://godoc.org/github.com/030/n3dr
[GoDoc Widget]: https://godoc.org/github.com/030/n3dr?status.svg
