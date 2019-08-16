# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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
- Fix 'incorrect folder name if artifact path contains repository name' by [@dbevacqua](https://github.com/dbevacqua).

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
- Defined commands that are used by download and repositories function, globally.
- Indicated what subcommand are required.

## [2.1.1] - 2019-05-19
### Added
- Issue templates.
- URL validation.
- Difference with equivalent tools explained.

### Fixed
- URL in repositories subcommand always set to http://localhost:9999.

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
- Coverage report changed by excluding all files that were created by cobra in a cmd folder.

### Removed
- Not implemented upload subcommand removed from README.

## [1.0.2] - 2019-05-15
### Added
- TestDownloadArtifacts.

### Changed
- Restrict testing to linux as docker is omitted on Mac and Windows build in travis.

### Fixed
- Broken Windows build due to formatting solved by enforcing LF using gitattributes.

## [1.0.1] - 2019-05-14
### Fixed
- Publication of artifacts.

## [1.0.0] - 2019-05-12
### Added
- Download all artifacts from a certain Nexus3 repository.

[Unreleased]: https://github.com/030/n3dr/compare/3.2.0...HEAD
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
