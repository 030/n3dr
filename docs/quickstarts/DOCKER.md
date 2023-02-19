# Docker

- [Download N3DR](./snippets/n3dr/DOWNLOAD.md).
- [Start a Nexus3 server](./snippets/nexus3/SERVER.md).
- Populate it with artifacts:

```bash
docker login localhost:8082 \
  -p $(docker exec -it nexus3-n3dr-src cat /nexus-data/admin.password) \
  -u admin && \
for d in $(seq 5); do
  docker_registry_tag=localhost:8082/repository/docker-images/utrecht/n3dr:6.${d}.0
  docker pull "utrecht/n3dr:6.${d}.0"
  docker tag "utrecht/n3dr:6.${d}.0" localhost:8082/repository/docker-images/utrecht/n3dr:6.${d}.0
  docker push localhost:8082/repository/docker-images/utrecht/n3dr:6.${d}.0
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
  --configRepoName docker-images \
  --configRepoType docker
```

- [Upload the artifacts to the other Nexus3 server](./snippets/n3dr/UPLOAD.md).
- [Validate](./snippets/n3dr/VALIDATE.md).
- [Cleanup](./snippets/nexus3/CLEANUP.md).
