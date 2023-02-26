# Rubygems

- [Download N3DR](./snippets/n3dr/DOWNLOAD.md).
- [Start a Nexus3 server](./snippets/nexus3/SERVER.md).
- Create a `some-rubygems` repository:

```bash
./n3dr configRepository \
  -u admin \
  -p $(docker exec -it nexus3-n3dr-src cat /nexus-data/admin.password) \
  -n localhost:8081 \
  --https=false \
  --configRepoName some-rubygems \
  --configRepoType gem
```

- Populate it with artifacts:

```bash
downloadDir=/tmp/some-dir/some-rubygems && \
mkdir -p ${downloadDir} && \
for f in chef-17.4.25 chef-18.1.0 puppet-7.23.0 rack-3.0.4.1; do
curl -L https://rubygems.org/downloads/${f}.gem > ${downloadDir}/${f}.gem
curl -X 'POST' \
    localhost:8081/service/rest/v1/components?repository=some-rubygems \
    -s \
    -f \
    -u "admin:$(docker exec -it nexus3-n3dr-src cat /nexus-data/admin.password)" \
    -H 'accept: application/json' \
    -H 'Content-Type: multipart/form-data' \
    -F "rubygems.asset=@${downloadDir}/${f}.gem;type=application/x-tar"
done
```

- [Backup all artifacts](./snippets/n3dr/BACKUP.md).
- [Start another Nexus3 server](./snippets/nexus3/ANOTHERSERVER.md).
- Create a repository in the other Nexus3 server:

```bash
./n3dr configRepository \
  -u admin \
  -p $(docker exec -it nexus3-n3dr-dest cat /nexus-data/admin.password) \
  -n localhost:9000 \
  --https=false \
  --configRepoName some-rubygems \
  --configRepoType gem
```

- [Upload the artifacts to the other Nexus3 server](./snippets/n3dr/UPLOAD.md).
- [Validate](./snippets/n3dr/VALIDATE.md).
- [Cleanup](./snippets/nexus3/CLEANUP.md).
