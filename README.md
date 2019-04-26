# nexus3-cli

The aims of the nexus3-cli tool are:
 * to download all artifacts from a certain Nexus maven repository.
 * to upload all artifacts that reside in a local folder.

## How do the tests look like?

The tests start a nexus docker container. The tests will be started once
the docker container is running. A springboot app will be deployed,
the Spring dependencies will be downloaded and uploaded to nexus. Finally,
all submitted artifacts will be downloaded.