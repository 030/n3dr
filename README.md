# N3DR

[![CI](https://github.com/030/n3dr/workflows/Go/badge.svg?event=push)](https://github.com/030/n3dr/actions?query=workflow%3AGo)
[![GoDoc Widget]][GoDoc]
[![Go Report Card](https://goreportcard.com/badge/github.com/030/n3dr)](https://goreportcard.com/report/github.com/030/n3dr)
[![StackOverflow SE Questions](https://img.shields.io/stackexchange/stackoverflow/t/n3dr.svg?logo=stackoverflow)](https://stackoverflow.com/tags/n3dr)
[![DevOps SE Questions](https://img.shields.io/stackexchange/devops/t/n3dr.svg?logo=stackexchange)](https://devops.stackexchange.com/tags/n3dr)
[![ServerFault SE Questions](https://img.shields.io/stackexchange/serverfault/t/n3dr.svg?logo=serverfault)](https://serverfault.com/tags/n3dr)
![Docker Pulls](https://img.shields.io/docker/pulls/utrecht/n3dr.svg)
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
[![Chocolatey](https://img.shields.io/chocolatey/dt/n3dr)](https://chocolatey.org/packages/n3dr)
[![n3dr](https://snapcraft.io//n3dr/badge.svg)](https://snapcraft.io/n3dr)
[![codebeat badge](https://codebeat.co/badges/6c33543d-d05b-44e4-8924-140382148de9)](https://codebeat.co/projects/github-com-030-n3dr-master)

<a href="https://n3dr.releasesoftwaremoreoften.com">\
<img src="https://github.com/030/n3dr/raw/master/cmd/n3dr/assets/logo/logo.png" width="100"></a>

## Backup or Migrate an entire Nexus Artifact Server

Download all artifacts at once or migrate automatically from Nexus to Nexus.

Although the [Nexus backup and restore documentation](https://help.sonatype.com/repomanager3/backup-and-restore)
indicates that one could backup and restore Nexus, the data seems not to be
restored completely as 500 errors occur when an artifact is downloaded from the
UI after restore. It could also be possible that some steps were not issued as
they have should been. Apart from that, the restore is capable of restoring the
Nexus configuration.

N3DR excludes the backup of group repositories and is able to backup all Maven2
and NPM repositories and migrate and/or restore Maven2 artifacts to another
Nexus server.

Note: uploads to proxy and snapshot repositories are not supported by Nexus
itself. As a workaround one could create a hosted repository in Nexus and
upload the backed up proxy content to it.

The aims of the n3dr tool are:

* to backup all artifacts from a certain Nexus maven repository.
* to migrate all artifacts from NexusA to NexusB.
* configuration-as-code.

## Installation

### Linux

```bash
snap install n3dr
```

Check the downloaded artifacts:

```bash
sudo ls /tmp/snap.n3dr/tmp/n3dr/download<some number, e.g.: 082028764>
sudo cp -r /tmp/snap.n3dr/tmp/n3dr/download<some number, e.g.: 082028764> /home/${USER}/n3dr-backup
sudo chown $USER:$USER -R /home/${USER}/n3dr-backup
```

### MacOSX

Get the darwin artifact from the releases tab.

### Windows

```bash
choco install n3dr
```

## Configuration

### N3DR download user

Create a user, e.g. n3dr-download in Nexus3, create a role, e.g. n3dr-download
and assign the following roles:

* `nx-repository-view-*-*-browse`
* `nx-repository-view-*-*-read`

### N3DR upload user

In order to upload artifacts, additional privileges are required:

* `nx-repository-view-*-*-add`
* `nx-repository-view-*-*-edit`

## Usage

<a href="https://asciinema.org/a/Oqwg69HJV0hFnnxxLZR6vbBeH?autoplay=1">\
<img src="https://asciinema.org/a/Oqwg69HJV0hFnnxxLZR6vbBeH.svg" /></a>

### Check the help menu

```bash
user@computer:~/dev$ n3dr -h
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
  -v, --apiVersion string        The Nexus3 APIVersion, e.g. v1 or beta
                                 (default "v1")
  -d, --debug                    Enable debug logging
  -h, --help                     help for n3dr
      --insecureSkipVerify       Skip repository certificate check
  -p, --n3drPass string          The Nexus3 password
  -n, --n3drURL string           The Nexus3 URL
  -u, --n3drUser string          The Nexus3 user
  -z, --zip                      Add downloaded artifacts to a ZIP archive
      --directory-prefix string  The directory prefix is the directory where
                                 artifacts will be saved

Use "n3dr [command] --help" for more information about a command.
```

### insecureSkipVerify

It is possible to load a custom CA to connect to Nexus3 if one created
self-signed certificates, by using:

```bash
--insecureSkipVerify
```

Note: store the `ca.crt` in the `~/.n3dr` directory.

### Anonymous

In order to download as a anonymous user, one has to use the `--anonymous`
option.

### Configuration-as-code

#### LDAP

```bash
n3dr configLDAP \
  --configLDAPAuthPassword=a \
  --configLDAPAuthUsername=b \
  --configLDAPHost=c \
  --configLDAPName=d \
  -p <admin-pass> \
  -u <admin-user> \
  -n=<FQDN-without-http://-or-https>:<port-if-applicable>
```

#### Create an admin user

```bash
n3dr configUser \
  --email=some-admin-user@some-admin-user.some-admin-user \
  --firstName=some-admin-user \
  --lastName=some-admin-user \
  --pass=some-admin-user \
  --id=some-admin-user \
  --admin \
  -p <admin-pass> \
  -u <admin-user> \
  -n=<FQDN-without-http://-or-https>:<port-if-applicable>
```

#### Create a downloadUser

```bash
n3dr configUser \
  --email=some-admin-user@some-admin-user.some-admin-user \
  --firstName=some-admin-user \
  --lastName=some-admin-user \
  --pass=some-admin-user \
  --id=some-admin-user \
  --downloadUser \
  -p <admin-pass> \
  -u <admin-user> \
  -n=<FQDN-without-http://-or-https>:<port-if-applicable>
```

#### Create a uploadUser

```bash
n3dr configUser \
  --email=some-admin-user@some-admin-user.some-admin-user \
  --firstName=some-admin-user \
  --lastName=some-admin-user \
  --pass=some-admin-user \
  --id=some-admin-user \
  --uploadUser \
  -p <admin-pass> \
  -u <admin-user> \
  -n=<FQDN-without-http://-or-https>:<port-if-applicable>
```

#### Change a user pass

```bash
n3dr configUser \
  --email=admin@example.org \
  --firstName=admin \
  --lastName=admin \
  --pass=some-other-admin-pass \
  --id=admin \
  --changePass \
  -p <admin-pass> \
  -u <admin-user> \
  -n=<FQDN-without-http://-or-https>:<port-if-applicable>
```

#### Delete a repository

```bash
n3dr configRepository \
  --configRepoName some-repo-name \
  --configRepoDelete \
  -p <admin-pass> \
  -u <admin-user> \
  -n=<FQDN-without-http://-or-https>:<port-if-applicable>
```

#### Create a repository

##### Hosted Raw

```bash
n3dr configRepository \
  --configRepoName some-repo \
  --configRepoType raw \
  -p <admin-pass> \
  -u <admin-user> \
  -n=<FQDN-without-http://-or-https>:<port-if-applicable>
```

##### Proxied Apt

```bash
n3dr configRepository \
  --configRepoName some-apt-proxy-repo \
  --configRepoType apt \
  --configRepoRecipe proxy \
  --configRepoProxyURL "http://nl.archive.ubuntu.com/ubuntu/" \
  -p <admin-pass> \
  -u <admin-user> \
  -n=<FQDN-without-http://-or-https>:<port-if-applicable>
```

##### Proxied Yum

```bash
n3dr configRepository \
  --configRepoName some-yum-proxy-repo \
  --configRepoType yum \
  --configRepoRecipe proxy \
  --configRepoProxyURL "http://mirror.centos.org/centos/" \
  -p <admin-pass> \
  -u <admin-user> \
  -n=<FQDN-without-http://-or-https>:<port-if-applicable>
```

#### Anonymous access

##### Disable

```bash
n3dr config \
  -p <admin-pass> \
  -u <admin-user> \
  -n=<FQDN-without-http://-or-https>:<port-if-applicable>
```

##### Enable

```bash
n3dr config \
  -p <admin-pass> \
  -u <admin-user> \
  -n=<FQDN-without-http://-or-https>:<port-if-applicable> \
  --configUserAnonymous
```

## Docker

### Build

```bash
docker build -t utrecht/n3dr:6.4.1 .
```

[![dockeri.co](https://dockeri.co/image/utrecht/n3dr)](https://hub.docker.com/r/utrecht/n3dr)

### Download

```bash
docker run -it \
  -v /home/${USER}/.n3dr:/root/.n3dr \
  -v /tmp/n3dr:/tmp/n3dr utrecht/n3dr:6.4.1
```

### Upload

```bash
docker run -it \
  --entrypoint=/bin/ash \
  -v /home/${USER}/.n3dr:/root/.n3dr \
  -v /tmp/n3dr:/tmp/n3dr utrecht/n3dr:6.4.1
```

navigate to the repository folder, e.g. `/tmp/n3dr/download*/` and upload:

```bash
n3dr upload -r releases -n <url>
```

Note: if the snap package is used to upload artifacts, one has to ensure that
the folder resides in the /home/$USER folder. Otherwise a:
`lstat <repository-name>: no such file or directory` issue could occur.

#### skipErrors

One could use `--skipErrors` or `-s` to continue-on-error:

```bash
N3DR_MAVEN_UPLOAD_REGEX_VERSION=boo \
N3DR_MAVEN_UPLOAD_REGEX_CLASSIFIER=foo \
n3dr upload -n some-nexus-url \
            -r some-repository \
            -s
```

## Store the password in a read-only file

Define the password in `~/.n3dr/config.yml`:

```bash
---
n3drPass: admin123
```

and set the permissions to read-write by issuing:

```bash
chmod 0600 ~/.n3dr/config.yml
```

Note: other variables like `n3drURL` and `n3drUser` could also be added to the
config file and one could use `--config` to overwrite the default config path.

### Backup artifacts from a certain repository

All artifacts from a repository will be stored in a download folder when
the following command is run:

```bash
n3dr backup -u admin -n http://localhost:8081 -r maven-releases
```

### Backup artifacts from a repositories list

All artifacts from a repositories list will be stored in a download folder when
the following command is run:

```bash
n3dr backup -u admin -n http://localhost:8081 -r maven-releases,maven-private
```

### Backup all repositories

All artifacts from various repositories will be stored in a download
folder when the following command is issued:

```bash
n3dr repositories -u admin -n http://localhost:8081 -b
```

Note: a new folder will be created for every repository:

* download/maven-public
* download/maven-releases

### Backup only certain artifacts

It is possible to only download artifacts that match a regular expression. If
one would like to download all artifacts from 'some/group42' then one could do
that as follows:

```bash
n3dr backup -u admin -n http://localhost:8081 -r maven-releases -x 'some/group42'
```

If one would like to deploy is while download from all repositories then use
the `-x` option as well:

```bash
n3dr repositories -u admin -n http://localhost:8081 -b -x 'some/group42'
```

### Add all downloaded archives to a ZIP archive

In order to add all archives to a zip archive, one has to use the --zip or -z flag.

If one would like to overwrite the default zip file name, then one has to use
the `-i` option. Note: the extension '.zip' is obliged.

### Upload all artifacts to a certain repository

It is possible to upload all JARs that reside in a folder by
running the following command:

```bash
n3dr upload -u admin -n http://localhost:8081 -r maven-public
```

#### Upload non maven files

It is possible to upload non maven files like deb files as well by setting the
artifactType option to the repository type, e.g. `-t=apt`. Note that the folder
name that contains the files should match the repository name.

### "Clone" a Nexus3 repository

Suppose that one has created a new Nexus3 repository, e.g. NexusNEW and that
one would like to copy the content of the old repository, e.g. NexusOLD, then
these basic steps could be issued to "clone" NexusOLD:

```bash
n3dr backup -u <old-nexus3-user> -n <old-nexus3-server-url> \
-r <old-repo-source-name>
cd download
mv <old-repo-source-name> <new-repo-target-name>
n3dr upload -u <new-target-nexus3-user> -n <new-target-nexus3-server-url> \
-r <new-repo-target-name>
```

### Backup to OCI Object Storage

`n3dr` supports backing up to [OCI Object Storage](https://www.oracle.com/cloud/storage/object-storage.html).
To enable this option you need to

* Configure OCI environment and secrets locally: <https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm/>
* Add following options to `~/.n3dr/config.yml`:

```bash
ociBucket: nexus_dev_archives
```

If you want to remove local copies (after the object has been uploaded), add
following to `~/.n3dr/config.yml`:

```bash
removeLocalFile: true
```

### Backup NPM repositories

Use the `--npm` parameter to backup NPM artifacts.

```bash
go run main.go backup -npm -n some-url -r some-npm-repo -d --npm
```

## logo

If you want to change the logo, just edit
`cmd/n3dr/assets/logo/text-image-com-n3dr.txt` and rebuild with

```bash
go build
```

## Rationale for N3DR

Although there is a number of equivalent tools:

* <https://github.com/RiotGamesMinions/nexus_cli/>
* <https://github.com/packagemgmt/repositorytools/>
* <https://github.com/thiagofigueiro/nexus3-cli/>

None of them seems to be able to backup all repositories by running
a single command.

[GoDoc]: https://godoc.org/github.com/030/n3dr
[GoDoc Widget]: https://godoc.org/github.com/030/n3dr?status.svg

## Supported

| type   | backup | upload | label |
|--------|--------|--------|-------|
| apt    |        | x      |       |
| maven2 | x      | x      | `+`   |
| npm    | x      | x      | `*`   |
| nuget  |        | x      |       |

### repositoriesV2

| type      | backup | upload | label |
|-----------|--------|--------|-------|
| apt       | x      | x      | `^`   |
| bower     |        |        |       |
| cocoapods |        |        |       |
| conan     |        |        |       |
| conda     |        |        |       |
| docker    |        |        |       |
| gitlfs    |        |        |       |
| go        |        |        |       |
| helm      |        |        |       |
| maven2    | x      | x      | `+`   |
| npm       | x      | x      | `*`   |
| nuget     | x      | x      | `$`   |
| p2        |        |        |       |
| pypi      |        |        |       |
| r         |        |        |       |
| raw       | x      | x      | `%`   |
| rubygems  |        |        |       |
| yum       |        |        |       |
| unknown   | x      | x      | `?`   |

#### backup

`repositoriesV2` command in conjunction with the `--backup` subcommand ensures
that all artifacts are downloaded from Nexus3 and stored in a folder that is
defined with the `--directory-prefix` parameter.

```bash
n3dr repositoriesV2 \
  --backup \
  -n some-url \
  -u some-user \
  -p some-pass \
  --directory-prefix /tmp/some-dir
```

#### upload

```bash
n3dr repositoriesV2 \
  --upload \
  -n some-url \
  -u some-user \
  -p some-pass \
  --directory-prefix /tmp/some-dir
```

#### sync

```bash
n3dr sync \
  --otherNexus3Passwords=some-pass-B,some-pass-C \
  --otherNexus3Users=admin,admin \
  --otherNexus3URLs=localhost:9998,localhost:9997 \
  --directory-prefix /some/dir \
  -p some-pass-A \
  -n localhost:9999 \
  -u admin
```

Note: use `--https=false` in order to connect to a <http://nexus-url/>.

## Stargazers over time

[![Stargazers over time](https://starchart.cc/030/n3dr.svg)](https://starchart.cc/030/n3dr)

## Development

### go-swagger

Comment out `trap` in the integration-tests.sh and run it:

```bash
./test/integration-tests.sh
```

Once Nexus had been started, download the go-swagger, swagger.json and
generate internal go-swagger code:

```bash
export GITHUB_URL=https://github.com
export GS_URI=go-swagger/go-swagger/releases/download
export GS_VERSION=v0.28.0
export GS_URL=${GITHUB_URL}/${GS_URI}/${GS_VERSION}/swagger_linux_amd64
curl -L \
  ${GS_URL} \
  -o swagger
chmod +x swagger
mkdir -p internal/goswagger
curl http://localhost:9999/service/rest/swagger.json -o swagger.json
./swagger generate client \
  --name=nexus3 \
  --spec swagger.json \
  --target=internal/goswagger \
  --skip-validation
go mod tidy
```

### Unit Tests

```bash
go test internal/artifacts/common.go internal/artifacts/common_test.go
```

### Integration testing on Windows

#### Packer

```bash
packer init packer/windows2016.json.pkr.hcl
PACKER_LOG=1 packer build packer/windows2016.json.pkr.hcl
```

#### Vagrant

```bash
vagrant box add virtualbox_windows2016.box --name win2016/n3dr
vagrant box list
vagrant plugin install vagrant-reload vagrant-windows-update
export VAGRANT_N3DR_NETWORK_ADAPTER=$(vboxmanage list bridgedifs |\
  grep Name: | head -1 | awk '{ print $2 }')
VAGRANT_NEXUS3_IP=192.168.0.42 VAGRANT_N3DR_IP=192.168.0.43 vagrant up
vagrant provision nexus3
vagrant destroy -f
vagrant ssh nexus3
```

Login as `vagrant` with pass `vagrant` and issue:

```bash
cd C:\vagrant
.\cmd\n3dr\n3dr.exe download -r maven-releases -n http://192.168.0.42:9999 \
  -u admin -p some-password
.\cmd\n3dr\n3dr.exe upload -r maven-releases -n http://192.168.0.42:9999 \
  -u admin -p some-password
```

To check whether it is possible to upload artifacts to Nexus3 from Windows
using N3DR.
