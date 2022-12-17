# docs

- [examples](./../examples)

## Instructions

- [snap](docs/instructions/snap.md)

## Installation

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

- `nx-repository-view-*-*-browse`
- `nx-repository-view-*-*-read`

### N3DR upload user

In order to upload artifacts, additional privileges are required:

- `nx-repository-view-*-*-add`
- `nx-repository-view-*-*-edit`

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

#### Create an uploadUser

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
docker build -t utrecht/n3dr:7.0.0 .
```

[![dockeri.co](https://dockeri.co/image/utrecht/n3dr)](https://hub.docker.com/r/utrecht/n3dr)

### Download

```bash
docker run -it \
  -v /home/${USER}/.n3dr:/root/.n3dr \
  -v /tmp/n3dr:/tmp/n3dr utrecht/n3dr:7.0.0
```

### Upload

```bash
docker run -it \
  --entrypoint=/bin/ash \
  -v /home/${USER}/.n3dr:/root/.n3dr \
  -v /tmp/n3dr:/tmp/n3dr utrecht/n3dr:7.0.0
```

navigate to the repository folder, e.g. `/tmp/n3dr/download*/` and upload:

```bash
n3dr upload -r releases -n <url>
```

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

- download/maven-public
- download/maven-releases

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

- Configure OCI environment and secrets locally: <https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm/>
- Add following options to `~/.n3dr/config.yml`:

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

- <https://github.com/RiotGamesMinions/nexus_cli/>
- <https://github.com/packagemgmt/repositorytools/>
- <https://github.com/thiagofigueiro/nexus3-cli/>

None of them seems to be able to backup all repositories by running
a single command.

[godoc]: https://godoc.org/github.com/030/n3dr
[godoc widget]: https://godoc.org/github.com/030/n3dr?status.svg

## Supported

| type   | backup | upload | label |
| ------ | ------ | ------ | ----- |
| apt    |        | x      |       |
| maven2 | x      | x      | `+`   |
| npm    | x      | x      | `*`   |
| nuget  |        | x      |       |

### repositoriesV2

| type      | backup | upload | label |
| --------- | ------ | ------ | ----- |
| apt       | x      | x      | `^`   |
| bower     |        |        |       |
| cocoapods |        |        |       |
| composer  |        |        |       |
| conan     |        |        |       |
| conda     |        |        |       |
| cpan      |        |        |       |
| docker    | x      | x      | ``    |
| elpa      |        |        |       |
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
| yum       | x      | x      | `#`   |
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

## Development

- [docs](./docs/development/README.md)

### Integration testing on Windows

#### Packer

```bash
packer init build/packer/windows2016.json.pkr.hcl
PACKER_LOG=1 packer build build/packer/windows2016.json.pkr.hcl
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
.\cmd\n3dr\n3dr.exe backup -r maven-releases -n http://192.168.0.42:9999 \
  -u admin -p some-password
.\cmd\n3dr\n3dr.exe upload -r maven-releases -n http://192.168.0.42:9999 \
  -u admin -p some-password
```

To check whether it is possible to upload artifacts to Nexus3 from Windows
using N3DR.
