# Backup

```bash
./n3dr repositoriesV2 \
  --backup \
  --directory-prefix /tmp/some-dir-backup \
  -u admin \
  -p $(docker exec -it nexus3-n3dr-src cat /nexus-data/admin.password) \
  --dockerPort 8082 \
  --dockerHost http://localhost \
  -n localhost:8081 \
  --https=false
```
