# nexus3-cli

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