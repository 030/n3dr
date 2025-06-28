# N3DR

[![CI](https://github.com/030/n3dr/workflows/Go/badge.svg?event=push)](https://github.com/030/n3dr/actions?query=workflow%3AGo)
[![GoDoc Widget]][godoc]
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/030/n3dr?logo=go)
[![Go Report Card](https://goreportcard.com/badge/github.com/030/n3dr)](https://goreportcard.com/report/github.com/030/n3dr)
[![StackOverflow SE Questions](https://img.shields.io/stackexchange/stackoverflow/t/n3dr.svg?logo=stackoverflow)](https://stackoverflow.com/tags/n3dr)
[![DevOps SE Questions](https://img.shields.io/stackexchange/devops/t/n3dr.svg?logo=stackexchange)](https://devops.stackexchange.com/tags/n3dr)
[![ServerFault SE Questions](https://img.shields.io/stackexchange/serverfault/t/n3dr.svg?logo=serverfault)](https://serverfault.com/tags/n3dr)
[![Docker Pulls](https://img.shields.io/docker/pulls/utrecht/n3dr?logo=docker&logoColor=white)](https://hub.docker.com/r/utrecht/n3dr)
[![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/utrecht/n3dr?logo=docker&logoColor=white&sort=semver)](https://hub.docker.com/r/utrecht/n3dr)
![Issues](https://img.shields.io/github/issues-raw/030/n3dr.svg)
![Pull requests](https://img.shields.io/github/issues-pr-raw/030/n3dr.svg)
![Total downloads](https://img.shields.io/github/downloads/030/n3dr/total.svg)
![GitHub forks](https://img.shields.io/github/forks/030/n3dr?label=fork&style=plastic)
![GitHub watchers](https://img.shields.io/github/watchers/030/n3dr?style=plastic)
![GitHub stars](https://img.shields.io/github/stars/030/n3dr?style=plastic)
![License](https://img.shields.io/github/license/030/n3dr.svg)
![Repository Size](https://img.shields.io/github/repo-size/030/n3dr.svg)
![Contributors](https://img.shields.io/github/contributors/030/n3dr.svg)
![Commit activity](https://img.shields.io/github/commit-activity/m/030/n3dr.svg)
![Last commit](https://img.shields.io/github/last-commit/030/n3dr.svg)
![Release date](https://img.shields.io/github/release-date/030/n3dr.svg)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/030/n3dr?logo=github&sort=semver)](https://github.com/030/n3dr/releases/latest)
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
[![codecov](https://codecov.io/gh/030/n3dr/branch/main/graph/badge.svg)](https://codecov.io/gh/030/n3dr)
[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-web.svg)](https://golangci.com/r/github.com/030/n3dr)
[![Chocolatey](https://img.shields.io/chocolatey/dt/n3dr)](https://chocolatey.org/packages/n3dr)
[![n3dr](https://snapcraft.io//n3dr/badge.svg)](https://snapcraft.io/n3dr)
[![codebeat badge](https://codebeat.co/badges/f4aa5086-a4d5-41cd-893a-5da816ee9107)](https://codebeat.co/projects/github-com-030-n3dr-main)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&logoColor=white)](https://conventionalcommits.org)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

[godoc]: https://godoc.org/github.com/030/n3dr
[godoc widget]: https://godoc.org/github.com/030/n3dr?status.svg

[![dockeri.co](https://dockeri.co/image/utrecht/n3dr)](https://hub.docker.com/r/utrecht/n3dr)

Nexus3 Disaster Recovery (N3DR).

## Backup or Migrate an entire Nexus Artifact Server

Download all artifacts at once or migrate automatically from Nexus to Nexus.

Although the [Nexus backup and restore documentation](https://help.sonatype.com/repomanager3/backup-and-restore)
indicates that one could backup and restore Nexus, the data seems not to be
restored completely as 500 errors occur when an artifact is downloaded from the
UI after restore. It could also be possible that some steps were not issued as
they have should been. Apart from that, the restore is capable of restoring the
Nexus configuration.

N3DR excludes the backup of group repositories and is able to backup various
repositories, configure them, restore artifacts or migrate them to another
Nexus3 server. Table 1 indicates what is supported by N3DR.

Note: uploads to proxy repositories are not supported by Nexus3 itself. As a
workaround one could create a hosted repository in Nexus and upload the backed
up proxy content to it.

The aims of the n3dr tool are:

- to backup artifacts from a certain Nexus3 repository.
- to migrate and/or restore them to another Nexus3 server.
- configuration-as-code (cac).

| type      | backup | upload | label | cac |
| --------- | ------ | ------ | ----- | --- |
| apt       | x      | x      | `^`   | x   |
| cargo     |        |        |       |     |
| cocoapods |        |        |       |     |
| composer  |        |        |       |     |
| conan     |        |        |       |     |
| conda     |        |        |       |     |
| docker    | x      | x      | ``    | x   |
| gitlfs    |        |        |       |     |
| go        |        |        |       |     |
| helm      |        |        |       |     |
|huggingface|        |        |       |     |
| maven2    | x      | x      | `+`   | x   |
| npm       | x      | x      | `*`   | x   |
| nuget     | x      | x      | `$`   | x   |
| p2        |        |        |       |     |
| pypi      |        |        |       |     |
| r         |        |        |       |     |
| raw       | x      | x      | `%`   | x   |
| rubygems  | x      | x      | `-`   | x   |
| yum       | x      | x      | `#`   | x   |
| unknown   | n/a    | n/a    | `?`   |     |

**Table 1**: Overview of Nexus3 types that can be downloaded, uploaded and/or
configured by N3DR.

## Quickstarts

- [docker](./docs/quickstarts/DOCKER.md)
- [maven2](./docs/quickstarts/MAVEN2.md)
- [rubygems](./docs/quickstarts/RUBYGEMS.md)

## Instructions

- [snap](./docs/instructions/snap.md)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/030/n3dr.svg)](https://starchart.cc/030/n3dr)
