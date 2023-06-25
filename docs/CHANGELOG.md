<a name="unreleased"></a>
## [Unreleased]


<a name="7.2.2"></a>
## [7.2.2] - 2023-06-25
### Build
- **deps:** Update versions.

### Fix
- [[GH-371](https://github.com/030/n3dr/issues/371)] Resolve indentation issue in count table formatting.


<a name="7.2.1"></a>
## [7.2.1] - 2023-03-19

<a name="7.2.0"></a>
## [7.2.0] - 2023-02-26
### Feat
- [[#363](https://github.com/030/n3dr/issues/363)] Backup and upload of Ruby gems.

### Fix
- [[#363](https://github.com/030/n3dr/issues/363)] Remove version update check old README.


<a name="7.1.1"></a>
## [7.1.1] - 2023-02-19
### Fix
- Missing &&'s in Docker quickstarts doc.
- [[#342](https://github.com/030/n3dr/issues/342)] Backup integration test.
- [[#238](https://github.com/030/n3dr/issues/238)] port in quickstart should be 9001.
- [[#312](https://github.com/030/n3dr/issues/312)] Errorlint.
- [[#331](https://github.com/030/n3dr/issues/331)] Stylecheck.
- [[#333](https://github.com/030/n3dr/issues/333)][[#336](https://github.com/030/n3dr/issues/336)] Unused and whitespace.
- [[#319](https://github.com/030/n3dr/issues/319)] Godot.
- [[#312](https://github.com/030/n3dr/issues/312)] Gofumpt.
- [[#311](https://github.com/030/n3dr/issues/311)] Durationcheck.
- **docs:** [[#238](https://github.com/030/n3dr/issues/238)] Improve docker.


<a name="7.1.0"></a>
## [7.1.0] - 2023-01-15
### Feat
- [[#99](https://github.com/030/n3dr/issues/99)][[#297](https://github.com/030/n3dr/issues/297)] Count number of artifacts and save info to a CSV file.


<a name="7.0.3"></a>
## [7.0.3] - 2023-01-03
### Build
- Update CI images to latest.

### Fix
- [[#238](https://github.com/030/n3dr/issues/238)] Validate docker params.
- [[#309](https://github.com/030/n3dr/issues/309)] Enable bodyclose check and autoupdate gosec.
- **snap:** [[#238](https://github.com/030/n3dr/issues/238)] Artifact download failed due to renaming.


<a name="7.0.2"></a>
## [7.0.2] - 2022-12-25
### Fix
- [[#304](https://github.com/030/n3dr/issues/304)] Optional waitgroup to mitigate memory issues in large Nexus3 environments.


<a name="7.0.1"></a>
## [7.0.1] - 2022-12-20
### Fix
- [[#294](https://github.com/030/n3dr/issues/294)] Update n3dr version in docs/README.
- [[#294](https://github.com/030/n3dr/issues/294)] Update n3dr version in README.
- [[#294](https://github.com/030/n3dr/issues/294)] Merge conflicts.
- [[#294](https://github.com/030/n3dr/issues/294)] Apply auto updates.
- **snap:** [[#290](https://github.com/030/n3dr/issues/290)] Restrict release of older snaps to main.


<a name="7.0.0"></a>
## [7.0.0] - 2022-12-17
### Build
- **snap:** [[#290](https://github.com/030/n3dr/issues/290)] Stable version 7.

### Fix
- [[#263](https://github.com/030/n3dr/issues/263)] Remove deprecated commands.
- **logging:** [[#270](https://github.com/030/n3dr/issues/270)] Optional write to syslog and/or file and default loglevel set to info.
- **snap:** [[#290](https://github.com/030/n3dr/issues/290)] Snap and release version 7.

### BREAKING CHANGE

The `--debug` and `-d` shorthand have been replaced by `--logLevel debug`.

The `backup`, `upload` and `repositories` commands have been removed.


<a name="6.8.3"></a>
## [6.8.3] - 2022-11-28
### Build
- **auto-updater:** Update schedule.

### Fix
- [[#254](https://github.com/030/n3dr/issues/254)] Broken snapcraft build.
- **maven2:** [[#254](https://github.com/030/n3dr/issues/254)] Upload snapshots.


<a name="6.8.2"></a>
## [6.8.2] - 2022-11-11
### Fix
- Ensure that auto updater updates the nexus version in the integration test.


<a name="6.8.1"></a>
## [6.8.1] - 2022-11-04
### Fix
- [[#278](https://github.com/030/n3dr/issues/278)] Backup a single repository using repositoriesV2.


<a name="6.8.0"></a>
## [6.8.0] - 2022-10-30
### Feat
- **logging:** [[#270](https://github.com/030/n3dr/issues/270)] Improve by adding a trace level.

### Fix
- **windows:** [[#270](https://github.com/030/n3dr/issues/270)] Omit syslog.


<a name="6.7.5"></a>
## [6.7.5] - 2022-10-23
### Fix
- **docs:** [[#264](https://github.com/030/n3dr/issues/264)] Improve snap.
- **repositoriesV2:** [[#265](https://github.com/030/n3dr/issues/265)] Add missing zip functionality.


<a name="6.7.4"></a>
## [6.7.4] - 2022-10-02
### Build
- **deps:** Update versions.
- **deps:** Add auto updater that creates a PR.


<a name="6.7.3"></a>
## [6.7.3] - 2022-08-20

<a name="6.7.2"></a>
## [6.7.2] - 2022-08-06

<a name="6.7.1"></a>
## [6.7.1] - 2022-07-24

<a name="6.7.0"></a>
## [6.7.0] - 2022-05-27

<a name="6.6.2"></a>
## [6.6.2] - 2022-05-07

<a name="6.6.1"></a>
## [6.6.1] - 2022-04-24

<a name="6.6.0"></a>
## [6.6.0] - 2022-04-19

<a name="6.5.1"></a>
## [6.5.1] - 2022-02-26

<a name="6.5.0"></a>
## [6.5.0] - 2022-01-16

<a name="6.4.3"></a>
## [6.4.3] - 2022-01-10

<a name="6.4.2"></a>
## [6.4.2] - 2022-01-03

<a name="6.4.1"></a>
## [6.4.1] - 2022-01-03

<a name="6.4.0"></a>
## [6.4.0] - 2022-01-02

<a name="6.3.0"></a>
## [6.3.0] - 2021-12-31

<a name="6.2.0"></a>
## [6.2.0] - 2021-11-30

<a name="6.1.0"></a>
## [6.1.0] - 2021-11-23

<a name="6.0.13"></a>
## [6.0.13] - 2021-10-11

<a name="6.0.12"></a>
## [6.0.12] - 2021-10-09

<a name="6.0.11"></a>
## [6.0.11] - 2021-05-28

<a name="6.0.10"></a>
## [6.0.10] - 2021-04-12

<a name="6.0.9"></a>
## [6.0.9] - 2021-04-06

<a name="6.0.8"></a>
## [6.0.8] - 2021-04-03

<a name="6.0.7"></a>
## [6.0.7] - 2021-04-03

<a name="6.0.6"></a>
## [6.0.6] - 2021-03-23

<a name="6.0.5"></a>
## [6.0.5] - 2021-03-21

<a name="6.0.4"></a>
## [6.0.4] - 2021-03-11

<a name="6.0.3"></a>
## [6.0.3] - 2021-03-07

<a name="6.0.2"></a>
## [6.0.2] - 2021-03-06

<a name="6.0.1"></a>
## [6.0.1] - 2021-02-18

<a name="6.0.0"></a>
## [6.0.0] - 2020-12-20

<a name="5.2.7"></a>
## [5.2.7] - 2020-12-14

<a name="5.2.6"></a>
## [5.2.6] - 2020-12-13

<a name="5.2.5"></a>
## [5.2.5] - 2020-12-12

<a name="5.2.4"></a>
## [5.2.4] - 2020-12-09

<a name="5.2.3"></a>
## [5.2.3] - 2020-12-07

<a name="5.2.2"></a>
## [5.2.2] - 2020-12-06

<a name="5.2.1"></a>
## [5.2.1] - 2020-09-12

<a name="5.2.0"></a>
## [5.2.0] - 2020-08-09

<a name="5.1.1"></a>
## [5.1.1] - 2020-08-07

<a name="5.1.0"></a>
## [5.1.0] - 2020-08-03

<a name="5.0.2"></a>
## [5.0.2] - 2020-08-03

<a name="5.0.1"></a>
## [5.0.1] - 2020-07-27

<a name="5.0.0"></a>
## [5.0.0] - 2020-07-24

<a name="4.1.4"></a>
## [4.1.4] - 2020-07-22

<a name="4.1.3"></a>
## [4.1.3] - 2020-07-21

<a name="4.1.2"></a>
## [4.1.2] - 2020-07-19

<a name="4.1.1"></a>
## [4.1.1] - 2020-07-19

<a name="4.1.0"></a>
## [4.1.0] - 2020-07-19

<a name="4.0.0"></a>
## [4.0.0] - 2020-07-19

<a name="3.6.3"></a>
## [3.6.3] - 2020-07-13

<a name="3.6.2"></a>
## [3.6.2] - 2020-07-08

<a name="3.6.1"></a>
## [3.6.1] - 2020-06-26

<a name="3.6.0"></a>
## [3.6.0] - 2020-06-25

<a name="3.5.1"></a>
## [3.5.1] - 2020-04-09

<a name="3.5.0"></a>
## [3.5.0] - 2020-03-29
### Reverts
- [[GH-86](https://github.com/030/n3dr/issues/86)] removed superfluous before_script block


<a name="3.4.0"></a>
## [3.4.0] - 2020-03-25

<a name="3.3.5-rc1"></a>
## [3.3.5-rc1] - 2020-03-01

<a name="3.3.4-rc1"></a>
## [3.3.4-rc1] - 2020-01-16

<a name="3.3.3"></a>
## [3.3.3] - 2019-12-12

<a name="3.3.2"></a>
## [3.3.2] - 2019-09-08

<a name="3.3.1"></a>
## [3.3.1] - 2019-09-06
### Reverts
- [[GH-86](https://github.com/030/n3dr/issues/86)] removed superfluous before_script block


<a name="3.3.0"></a>
## [3.3.0] - 2019-09-02

<a name="3.2.0"></a>
## [3.2.0] - 2019-08-17

<a name="3.1.1"></a>
## [3.1.1] - 2019-08-06

<a name="3.1.0"></a>
## [3.1.0] - 2019-06-02

<a name="3.0.0"></a>
## [3.0.0] - 2019-05-21

<a name="2.3.0"></a>
## [2.3.0] - 2019-05-20

<a name="2.2.1"></a>
## [2.2.1] - 2019-05-20

<a name="2.2.0"></a>
## [2.2.0] - 2019-05-19

<a name="2.1.1"></a>
## [2.1.1] - 2019-05-19

<a name="2.1.0"></a>
## [2.1.0] - 2019-05-19

<a name="2.0.0"></a>
## [2.0.0] - 2019-05-15

<a name="1.0.2"></a>
## [1.0.2] - 2019-05-15

<a name="1.0.1"></a>
## [1.0.1] - 2019-05-14

<a name="1.0.0"></a>
## 1.0.0 - 2019-05-12

[Unreleased]: https://github.com/030/n3dr/compare/7.2.2...HEAD
[7.2.2]: https://github.com/030/n3dr/compare/7.2.1...7.2.2
[7.2.1]: https://github.com/030/n3dr/compare/7.2.0...7.2.1
[7.2.0]: https://github.com/030/n3dr/compare/7.1.1...7.2.0
[7.1.1]: https://github.com/030/n3dr/compare/7.1.0...7.1.1
[7.1.0]: https://github.com/030/n3dr/compare/7.0.3...7.1.0
[7.0.3]: https://github.com/030/n3dr/compare/7.0.2...7.0.3
[7.0.2]: https://github.com/030/n3dr/compare/7.0.1...7.0.2
[7.0.1]: https://github.com/030/n3dr/compare/7.0.0...7.0.1
[7.0.0]: https://github.com/030/n3dr/compare/6.8.3...7.0.0
[6.8.3]: https://github.com/030/n3dr/compare/6.8.2...6.8.3
[6.8.2]: https://github.com/030/n3dr/compare/6.8.1...6.8.2
[6.8.1]: https://github.com/030/n3dr/compare/6.8.0...6.8.1
[6.8.0]: https://github.com/030/n3dr/compare/6.7.5...6.8.0
[6.7.5]: https://github.com/030/n3dr/compare/6.7.4...6.7.5
[6.7.4]: https://github.com/030/n3dr/compare/6.7.3...6.7.4
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
[3.4.0]: https://github.com/030/n3dr/compare/3.3.5-rc1...3.4.0
[3.3.5-rc1]: https://github.com/030/n3dr/compare/3.3.4-rc1...3.3.5-rc1
[3.3.4-rc1]: https://github.com/030/n3dr/compare/3.3.3...3.3.4-rc1
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
