# Pypi

- [Download N3DR](./snippets/n3dr/DOWNLOAD.md).
- [Start a Nexus3 server](./snippets/nexus3/SERVER.md).
- Create a `some-maven2` repository:

```bash
./n3dr configRepository \
  -u admin \
  -p $(docker exec -it nexus3-n3dr-src cat /nexus-data/admin.password) \
  -n localhost:8081 \
  --https=false \
  --configRepoName some-maven2 \
  --configRepoType maven2
```

- Populate it with artifacts:

```bash
downloadDir=/tmp/some-dir/some-maven2 && \
mkdir -p ${downloadDir} && \
for i in $(seq 20); do
  path=${downloadDir}/some/group/1.0.0/1.0.0-${i}
  filename="Some_Package"
  file="${filename}-1.0.0-${i}"
  filePath="${path}/${file}"
  mkdir -p ${path}
  dd if=/dev/urandom of=${filePath} bs=1M count=${i}

  zip ${filePath}.zip ${filePath}
  cp ${filePath} ${filePath}.jar
  cp ${filePath} ${filePath}-javadoc.jar
  cp ${filePath} ${filePath}-sources.jar
  cp ${filePath} ${filePath}.war
  echo hello > ${filePath}.module

  curl -X "POST" \
    localhost:8081/service/rest/v1/components?repository=some-maven2 \
    -v \
    -f \
    -u "admin:$(docker exec -it nexus3-n3dr-src cat /nexus-data/admin.password)" \
    -H "accept: application/json" \
    -H "Content-Type: multipart/form-data" \
    -F "maven2.asset1=@${filePath}.jar;type=application/x-java-archive" \
    -F "maven2.asset1.extension=jar" \
    -F "maven2.asset2=@${filePath}.war;type=application/x-java-archive" \
    -F "maven2.asset2.extension=war" \
    -F "maven2.asset3=@${filePath}.zip" \
    -F "maven2.asset3.extension=zip" \
    -F "maven2.asset4=@${filePath}.module" \
    -F "maven2.asset4.extension=module" \
    -F "maven2.asset5=@${filePath}-javadoc.jar;type=application/x-java-archive" \
    -F "maven2.asset5.classifier=javadoc" \
    -F "maven2.asset5.extension=jar" \
    -F "maven2.asset6=@${filePath}-sources.jar;type=application/x-java-archive" \
    -F "maven2.asset6.classifier=sources" \
    -F "maven2.asset6.extension=jar" \
    -F "maven2.groupId=some.group" \
    -F "maven2.artifactId=${filename}" \
    -F "maven2.version=1.0.0-${i}" \
    -F "maven2.generate-pom=true"
done
```

- [Backup all artifacts](./snippets/n3dr/BACKUP.md).
- Create a local repository without POM files:

```bash
cp -r /tmp/some-dir-backup/some-maven2 /tmp/some-dir-backup/some-maven2-without-pom
rm /tmp/some-dir-backup/some-maven2-without-pom/some/group/Some_Package/1.0.0-*/*.pom
```

- [Start another Nexus3 server](./snippets/nexus3/ANOTHERSERVER.md).
- Create two repositories in the other Nexus3 server:

```bash
for repo in some-maven2 some-maven2-without-pom; do
  ./n3dr configRepository \
    -u admin \
    -p $(docker exec -it nexus3-n3dr-dest cat /nexus-data/admin.password) \
    -n localhost:9000 \
    --https=false \
    --configRepoName "${repo}" \
    --configRepoType maven2
done
```

- [Upload the artifacts to the other Nexus3 server](./snippets/n3dr/UPLOAD.md).
- [Validate](./snippets/n3dr/VALIDATE.md).
- [Cleanup](./snippets/nexus3/CLEANUP.md).
