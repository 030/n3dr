# Upload

```bash
./n3dr repositoriesV2 \
  --upload \
  -u admin \
  -p $(docker exec -it nexus3-n3dr-dest cat /nexus-data/admin.password) \
  -n localhost:9000 \
  --https=false \
  --directory-prefix /tmp/some-dir \
  --dockerPort 9001 \
  --dockerHost http://localhost
```
