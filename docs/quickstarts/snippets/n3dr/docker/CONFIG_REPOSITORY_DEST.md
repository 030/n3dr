# config repository docker dest

```bash
./n3dr configRepository \
  -u admin \
  -p $(docker exec -it nexus3-n3dr-dest cat /nexus-data/admin.password) \
  -n localhost:9000 \
  --https=false \
  --configRepoName docker-images \
  --configRepoType docker
```
