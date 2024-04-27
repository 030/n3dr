# config repository docker

```bash
./n3dr configRepository \
  -u admin \
  -p $(docker exec -it nexus3-n3dr-src cat /nexus-data/admin.password) \
  -n localhost:8081 \
  --https=false \
  --configRepoName docker-images \
  --configRepoType docker
```
