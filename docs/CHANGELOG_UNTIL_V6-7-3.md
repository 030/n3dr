# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [6.7.3] - 2022-08-20

### Changed

- Alpine version to 3.16.2.

### Updated

- various Golang libraries.
- Nexus3 to 3.41.1.

## [6.7.2] - 2022-08-06

### Added

- deprecation notice that `repositoriesV2` will replace the `backup`,
  `repositories` and `upload` commands.

### Changed

- Golang version to 1.19.0-alpine3.16.
- GolangCI version to v1.48.0-alpine.
- Nexus3 and API version to 3.41.0.
- `io/ioutil` by either `os` or `io` as it is deprecated since Golang 1.16.

### Updated

- the versions of multiple Golang modules.

## [6.7.1] - 2022-07-24

### Changed

- Alpine version to 3.16.1.
- Golang version to 1.18.4-alpine3.16.
- GolangCI version to v1.47.2-alpine.
- Nexus3 and API version to 3.40.1.

### Updated

- the versions of multiple Golang modules.

## [6.7.0] - 2022-05-27

### Added

- support for uploading backup archives to AWS S3.
- example how to upload a backup archive to AWS S3.
- option to disable logo using the config file.

### Changed

- Alpine version to 3.16.0.
- Golang version to 1.18.2-alpine3.16.
- GolangCI version to v1.46.2-alpine.
- Nexus3 and API version to 3.39.0.

## [6.6.2] - 2022-05-07

### Changed

- API to 3.38.1.

## [6.6.1] - 2022-04-23

### Fixed

- cannot upload Maven artifacts containing underscores and uppercases.
  Reported by [fekal0id](https://github.com/fekal0id).

## [6.6.0] - 2022-04-19

### Added

- create a hosted docker repository.
- `showLogo` flag to hide N3DR logo.
- backup, upload and sync docker images.

### Changed

- Alpine to 3.15.4.
- Golang 1.17 to 1.18.
- GoSec 2.8.1 to 2.11.0.
- GolangCILint 1.39 to 1.45.

### Fixed

- cumbersome error handling in repositoriesV2.
- broken codebeat badge.

### Removed

- troublesome existence check while uploading.

## [6.5.1] - 2022-02-26

### Fixed

- version omitted when docker --version is run.

## [6.5.0] - 2022-01-16

### Added

- create hosted yum repositories.
- example how to create a hosted yum repository.
- upload RPMs into yum repository.

### Changed

- master renamed to main branch.

### Fixed

- return error if upload of an artifact fails to ensure that one is aware if
  some file cannot be uploaded and that files are not uploaded if they already
  exist in a Nexus3 repository.

### Removed

- some duplicated code by using `FilesToBeSkipped` and `PrintType` functions.

## [6.4.3] - 2022-01-10

### Added

- docs folder.
- examples.

### Changed

- libraries updated.
- moved CHANGELOG to docs folder.

### Fixed

- unable to use repositoriesV2 in conjunction with different Nexus3 basePath.

### Removed

- integration tests from unit testing.

## [6.4.2] - 2022-01-03

### Fixed

- N3DR does not fail if parameter is empty and not populated in
  `~/.n3dr/config.yml` reported by
  [der-eismann](https://github.com/der-eismann).

## [6.4.1] - 2022-01-03

### Changed

- Use Go embed for logo and change path for logo file by
  [der-eismann](https://github.com/der-eismann).

## [6.4.0] - 2022-01-02

### Added

- Upload of Maven2 artifacts using `repositoriesV2`.

## [6.3.0] - 2021-12-31

### Added

- Configure LDAP

### Fixed

- No backup of `group` repositories anymore.

## [6.2.0] - 2021-11-30

### Added

- Nuget downloads.
- Sync between multiple nexus repositories.

### Changed

- Moved the integration test to a separate github workflow file.

### Fixed

- Inconsistent file zip size while running `./test/integration.sh`.

## [6.1.0] - 2021-11-23

### Added

- Configuration-as-code:
  - Enable and disable anonymous access
  - Create a hosted raw repository
  - Configure a proxied apt repository
  - Delete repositories
  - Change the admin user pass
  - Create an admin user
  - Create a yum proxied repository
  - Generate a download and upload user

## [6.0.13] - 2021-10-11

### Fixed

- Resolve gosec G304, G307, G401 and G501 issues.
- Upload fails from Windows, reported by
  [TheMaddinFromTqg](https://github.com/TheMaddinFromTqg).

## [6.0.12] - 2021-10-09

### Fixed

- --n3drURL does not allow nexus repo URL with a context path, reported by
  [@thjenny](https://github.com/thjenny).

## [6.0.11] - 2021-05-28

### Fixed

- Some repositories were not backed up while using the `repositories` command.

## [6.0.10] - 2021-04-12

### Fixed

- Some Maven artifacts could not be uploaded, fixed by
[@kubealex](https://github.com/kubealex)

## [6.0.9] - 2021-04-06

### Fixed

- G402 by ensuring that self-signed CA should be used, reported by
[@titou10titou10](https://github.com/titou10titou10)

## [6.0.8] - 2021-04-03

### Added

- Release snap by CI

### Changed

- Replace walk by walkDir as it is faster

## [6.0.7] - 2021-04-02

### Added

- Scan waste in docker image using Dive
- Security scanning using Trivy
- Build and push docker image on new release

### Changed

- Alpine: 3.13.4
- Golang: 1.16.3
- Nexus3: 3.30.0

### Fixed

- Nexus3 IP address not allowed, reported by
  [@Just4test](https://github.com/Just4test)

## [6.0.6] - 2021-03-23

### Fixed

- Not possible to skip errors, reported by [@tunix](https://github.com/tunix)

## [6.0.5] - 2021-03-21

### Fixed

- Typo in app name.
- Omission of brew package.

## [6.0.4] - 2021-03-11

### Fixed

- Do not run docker image as root.

## [6.0.3] - 2021-03-07

### Fixed

- Docker not launching, reported by [@YoShiiro](https://github.com/YoShiiro)

## [6.0.2] - 2021-03-06

### Fixed

- Some artifacts could not be uploaded due to regex issue, reported by
  [@tunix](https://github.com/tunix)

## [6.0.1] - 2021-02-18

### Changed

- travis replaced by github actions.

### Fixed

- some maven files were not uploaded if they contained a `.`.

### Improved

- project structure

## [6.0.0] - 2020-12-20

### Added

- npm upload
- nuget upload

### Changed

- `maven` parameter replaced by `--artifactType`

## [5.2.7] - 2020-12-14

### Added

- shellcheck

### Fixed

- archetype-catalog.xml' does not seem to contain a Maven artifact

## [5.2.6] - 2020-12-13

### Added

- markdownlint

### Fixed

- Allow any kind of maven artifacts fix by [@michenux](https://github.com/Michenux)

## [5.2.5] - 2020-12-12

### Added

- Module extensions, aar and classifier support added by [@arcao](https://github.com/arcao)

### Changed

- NPM marked as `*` in logging

## [5.2.4] - 2020-12-08

### Fixed

- Integration test broken due to additional backup of sha256 and sha512 files
- Version subcommand broken due to change of repository path in go.mod
- Logo not included in darwin, linux and windows binaries

## [5.2.3] - 2020-12-07

### Fixed

- Asci art logo was not included in CLI

## [5.2.2] - 2020-12-06

### Fixed

- Backup of NPM repositories

### Removed

- Several superfluous tests

## [5.2.1] - 2020-09-12

### Fixed

- Infinite loop in `continuationTokenRecursionChannel` recursion due to
  omission of error checking reported by
  [@jdonkervliet](https://github.com/jdonkervliet)
- URL was not validated reported by
  [@jdonkervliet](https://github.com/jdonkervliet)

## [5.2.0] - 2020-08-09

### Changed

- Use channels to speed up backups

### Removed

- Progressbar

## [5.1.1] - 2020-08-07

### Fixed

- Upload of debian files did not work

## [5.1.0] - 2020-08-03

### Added

- Specify the directory where the backup should be stored
  added by [@vsoloviov](https://github.com/vsoloviov)

## [5.0.2] - 2020-08-03

### Fixed

- Retry HTTP logger was set to nil. Fixed by
[@vsoloviov](https://github.com/vsoloviov)

## [5.0.1] - 2020-07-27

### Added

- Select multiple repositories for backup added by [@vsoloviov](https://github.com/vsoloviov)

## [5.0.0] - 2020-07-24

### Added

- `--config` option to override default config path
- Lookup of `n3drUser` and `n3drURL` in config file. This means that these
  subcommands could be omitted when running n3dr.

### Changed

- Default config changed from `~/.n3dr.yaml` to `~/.n3dr/config.yml`. Note that
  the extension has been changed as well.
- Use the `--anonymous` subcommand to backup artifacts anonymously

### Deleted

- Superfluous viper lookup var calls on each command level by looking it up
  once on root level

### Fixed

- Unclarity regarding what subcommand should at least be used in conjunction
  with `repositories`

## [4.1.4] - 2020-07-21

### Added

- Warning that explains permission denied issue when running N3DR that was installed using snap

## [4.1.3] - 2020-07-20

### Added

- Overwrite default zipFileName

## [4.1.2] - 2020-07-19

### Fixed

- Duplicated downloads issue resolved by excluding backup of groups

## [4.1.1] - 2020-07-19

### Fixed

- Semi colons in zip file name not interpreted on some machines

## [4.1.0] - 2020-07-19

### Added

- Backup by anonymous users

### Fixed

- Incorrect version in snapcraft

## [4.0.0] - 2020-07-18

### Added

- Log the name of the n3dr backup zip

### Changed

- Added hour, minute and second to backup zip to prevent collision with previous backup zip
- Increased code coverage

### Fixed

- Old n3dr references updated in README
- Download folder should be cleaned to prevent pollution in consecutive backups

## [3.6.3] - 2020-07-13

### Fixed

- Backup of repositories failed due to HTTP timeout of a single artifact

## [3.6.2] - 2020-07-08

### Added

- Asciicast
- Configuration instructions minimal permissions n3dr user

### Changed

- Installation instructions

### Fixed

- Couple of code smells reported by sonar cloud

### Removed

- Superfluous go modules

## [3.6.1] - 2020-06-26

### Fixed

- Solved broken Darwin and Windows publication

## [3.6.0] - 2020-06-25

### Added

- OCI storage backend

## [3.5.1] - 2020-04-09

### Changed

- Golang version updated to 1.14.2

## [3.5.0] - 2020-03-29

### Added

- Download of specific artifacts using a regex
- Instructions added to README how to use this new feature

### Improved

- Section about how to clone an old Nexus3 repository

### Fixed

- Three code smells that were reported by SonarCloud

## [3.4.0] - 2020-03-25

### Added

- Upload of zip artifacts

## Changed

- Create files for testing rather than storing them in git

## [3.3.3] - 2019-12-12

### Added

- GolangCI Badge and scan
- Debian installation package

### Fixed

- Consistent help menu
- Multiple ignored errors
- Broken integration tests
- Allow `-p` in CI by preventing that omission of `~/.n3dr` returns an exit 1
  as suggested by [@jorianvo](https://github.com/jorianvo)

### Changed

- Refactored artifactName and initConfig functions to solve 'write shorter
  units of code' issues

## [3.3.2] - 2019-09-08

### Added

- Display of version by specifying `--version` thanks to explanation by
  [@umarcor](https://github.com/umarcor).

### Removed

- Lambda support as this tool is not suitable for running in serverless as it
  could take over more than 15 minutes to complete.

## [3.3.1] - 2019-09-06

### Fixed

- No error handling if password is omitted or incorrect. Issue reported by
  [@jorianvo](https://github.com/jorianvo).

## [3.3.0] - 2019-09-02

### Added

- Lambda support

## [3.2.0] - 2019-08-16

### Added

- Possibility to create a ZIP of downloaded artifacts
- Described in README how to add artifacts to a ZIP archive

### Changed

- Instructions how to use n3dr updated in README.md

### Removed

- Docker support

## [3.1.1] - 2019-08-06

### Fixed

- Fix 'incorrect folder name if artifact path contains repository name' by
  [@dbevacqua](https://github.com/dbevacqua).

## [3.1.0] - 2019-06-02

### Added

- Upload artifacts to a specific Nexus3 repository.

## [3.0.0] - 2019-05-21

### Added

- Enable debug logging.
- Progress bar as suggested by [@jorianvo](https://github.com/jorianvo).

### Changed

- Download command changed to backup.
- Majority of info logging changed to debug.

## [2.3.0] - 2019-05-20

### Added

- docker-compose example.

## [2.2.1] - 2019-05-20

### Fixed

- Broken repositories subcommands due to omission of authentication request.

## [2.2.0] - 2019-05-20

### Added

- Password lookup using viper.
- Documented how to lookup password from file.

### Fixed

- Help menu was not returned when invoking subcommand.

### Changed

- Defined commands that are used by download and repositories function,
  globally.
- Indicated what subcommand are required.

## [2.1.1] - 2019-05-19

### Added

- Issue templates.
- URL validation.
- Difference with equivalent tools explained.

### Fixed

- URL in repositories subcommand always set to <http://localhost:9999>.

## [2.1.0] - 2019-05-19

### Added

- Count number of repositories in a certain Nexus3 instance.
- Display names of all repositories.
- Download artifacts from all repositories.
- Updated README how to download all artifacts.

### Changed

- Create dedicated download folder for each repository.

## [2.0.0] - 2019-05-15

### Added

- Command Line Interface using [Cobra](https://github.com/spf13/cobra).

### Changed

- Coverage report changed by excluding all files that were created by
  cobra in a cmd folder.

### Removed

- Not implemented upload subcommand removed from README.

## [1.0.2] - 2019-05-15

### Added

- TestDownloadArtifacts.

### Changed

- Restrict testing to linux as docker is omitted on Mac and Windows build in
  travis.

### Fixed

- Broken Windows build due to formatting solved by enforcing LF using
  gitattributes.

## [1.0.1] - 2019-05-14

### Fixed

- Publication of artifacts.

## [1.0.0] - 2019-05-12

### Added

- Download all artifacts from a certain Nexus3 repository.

[Unreleased]: https://github.com/030/n3dr/compare/6.7.3...HEAD
[6.7.3]: https://github.com/030/n3dr/compare/6.7.2...6.7.3
[6.7.2]: https://github.com/030/n3dr/compare/6.7.1...6.7.2
[6.7.1]: https://github.com/030/n3dr/compare/6.7.0...6.7.1
[6.7.0]: https://github.com/030/n3dr/compare/6.6.2...6.7.0
[6.6.2]: https://github.com/030/n3dr/compare/6.6.1...6.6.2
[6.6.1]: https://github.com/030/n3dr/compare/6.6.0...6.6.1
[6.6.0]: https://github.com/030/n3dr/compare/6.5.1...6.6.0
[6.5.1]: https://github.com/030/n3dr/compare/6.5.0...6.5.1
[6.5.0]: https://github.com/030/n3dr/compare/6.4.3...6.5.0
[6.4.3]: https://github.com/030/n3dr/compare/6.4.2...6.4.3
[6.4.2]: https://github.com/030/n3dr/compare/6.4.1...6.4.2
[6.4.1]: https://github.com/030/n3dr/compare/6.4.0...6.4.1
[6.4.0]: https://github.com/030/n3dr/compare/6.3.0...6.4.0
[6.3.0]: https://github.com/030/n3dr/compare/6.2.0...6.3.0
[6.2.0]: https://github.com/030/n3dr/compare/6.1.0...6.2.0
[6.1.0]: https://github.com/030/n3dr/compare/6.0.13...6.1.0
[6.0.13]: https://github.com/030/n3dr/compare/6.0.12...6.0.13
[6.0.12]: https://github.com/030/n3dr/compare/6.0.11...6.0.12
[6.0.11]: https://github.com/030/n3dr/compare/6.0.10...6.0.11
[6.0.10]: https://github.com/030/n3dr/compare/6.0.9...6.0.10
[6.0.9]: https://github.com/030/n3dr/compare/6.0.8...6.0.9
[6.0.8]: https://github.com/030/n3dr/compare/6.0.7...6.0.8
[6.0.7]: https://github.com/030/n3dr/compare/6.0.6...6.0.7
[6.0.6]: https://github.com/030/n3dr/compare/6.0.5...6.0.6
[6.0.5]: https://github.com/030/n3dr/compare/6.0.4...6.0.5
[6.0.4]: https://github.com/030/n3dr/compare/6.0.3...6.0.4
[6.0.3]: https://github.com/030/n3dr/compare/6.0.2...6.0.3
[6.0.2]: https://github.com/030/n3dr/compare/6.0.1...6.0.2
[6.0.1]: https://github.com/030/n3dr/compare/6.0.0...6.0.1
[6.0.0]: https://github.com/030/n3dr/compare/5.2.7...6.0.0
[5.2.7]: https://github.com/030/n3dr/compare/5.2.6...5.2.7
[5.2.6]: https://github.com/030/n3dr/compare/5.2.5...5.2.6
[5.2.5]: https://github.com/030/n3dr/compare/5.2.4...5.2.5
[5.2.4]: https://github.com/030/n3dr/compare/5.2.3...5.2.4
[5.2.3]: https://github.com/030/n3dr/compare/5.2.2...5.2.3
[5.2.2]: https://github.com/030/n3dr/compare/5.2.1...5.2.2
[5.2.1]: https://github.com/030/n3dr/compare/5.2.0...5.2.1
[5.2.0]: https://github.com/030/n3dr/compare/5.1.1...5.2.0
[5.1.1]: https://github.com/030/n3dr/compare/5.1.0...5.1.1
[5.1.0]: https://github.com/030/n3dr/compare/5.0.2...5.1.0
[5.0.2]: https://github.com/030/n3dr/compare/5.0.1...5.0.2
[5.0.1]: https://github.com/030/n3dr/compare/5.0.0...5.0.1
[5.0.0]: https://github.com/030/n3dr/compare/4.1.4...5.0.0
[4.1.4]: https://github.com/030/n3dr/compare/4.1.3...4.1.4
[4.1.3]: https://github.com/030/n3dr/compare/4.1.2...4.1.3
[4.1.2]: https://github.com/030/n3dr/compare/4.1.1...4.1.2
[4.1.1]: https://github.com/030/n3dr/compare/4.1.0...4.1.1
[4.1.0]: https://github.com/030/n3dr/compare/4.0.0...4.1.0
[4.0.0]: https://github.com/030/n3dr/compare/3.6.3...4.0.0
[3.6.3]: https://github.com/030/n3dr/compare/3.6.2...3.6.3
[3.6.2]: https://github.com/030/n3dr/compare/3.6.1...3.6.2
[3.6.1]: https://github.com/030/n3dr/compare/3.6.0...3.6.1
[3.6.0]: https://github.com/030/n3dr/compare/3.5.1...3.6.0
[3.5.1]: https://github.com/030/n3dr/compare/3.5.0...3.5.1
[3.5.0]: https://github.com/030/n3dr/compare/3.4.0...3.5.0
[3.4.0]: https://github.com/030/n3dr/compare/3.3.3...3.4.0
[3.3.3]: https://github.com/030/n3dr/compare/3.3.2...3.3.3
[3.3.2]: https://github.com/030/n3dr/compare/3.3.1...3.3.2
[3.3.1]: https://github.com/030/n3dr/compare/3.3.0...3.3.1
[3.3.0]: https://github.com/030/n3dr/compare/3.2.0...3.3.0
[3.2.0]: https://github.com/030/n3dr/compare/3.1.1...3.2.0
[3.1.1]: https://github.com/030/n3dr/compare/3.1.0...3.1.1
[3.1.0]: https://github.com/030/n3dr/compare/3.0.0...3.1.0
[3.0.0]: https://github.com/030/n3dr/compare/2.3.0...3.0.0
[2.3.0]: https://github.com/030/n3dr/compare/2.2.1...2.3.0
[2.2.1]: https://github.com/030/n3dr/compare/2.2.0...2.2.1
[2.2.0]: https://github.com/030/n3dr/compare/2.1.1...2.2.0
[2.1.1]: https://github.com/030/n3dr/compare/2.1.0...2.1.1
[2.1.0]: https://github.com/030/n3dr/compare/2.0.0...2.1.0
[2.0.0]: https://github.com/030/n3dr/compare/1.0.2...2.0.0
[1.0.2]: https://github.com/030/n3dr/compare/1.0.1...1.0.2
[1.0.1]: https://github.com/030/n3dr/compare/1.0.0...1.0.1
[1.0.0]: https://github.com/030/n3dr/releases/tag/1.0.0
