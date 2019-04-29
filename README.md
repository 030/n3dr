# nexus3-cli

[![Build Status](https://travis-ci.org/030/nexus3-cli.svg?branch=master)](https://travis-ci.org/030/nexus3-cli)

## Sonar

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

The aims of the nexus3-cli tool are:
 * to backup all artifacts from a certain Nexus maven repository.
 * to migrate all artifacts from NexusA to NexusB.

## How do the tests look like?

The tests start a nexus docker container. The tests will be started once
the docker container is running. A springboot app will be deployed,
the Spring dependencies will be downloaded and uploaded to nexus. Finally,
all submitted artifacts will be downloaded.

## How to use this tool?

### Help

In order to read the help menu, one has to run:

```
docker run utrecht/nexus3-cli:1.0.0 -h
```

The output will look as follows:

```
Welcome to nexus3-cli
```

### Download

The download command will download all artifacts that reside in a Nexus maven
repository.

### Upload

The upload command will upload all artifacts to a Nexus maven repository.