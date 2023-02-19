# Validate

```bash
./n3dr count \
  -u admin \
  -p $(docker exec -it nexus3-n3dr-src cat /nexus-data/admin.password) \
  -n localhost:8081 \
  --https=false \
  --csv /tmp/some-dir/nexus3-n3dr-src && \
./n3dr count \
  -u admin \
  -p $(docker exec -it nexus3-n3dr-dest cat /nexus-data/admin.password) \
  -n localhost:9000 \
  --https=false \
  --csv /tmp/some-dir/nexus3-n3dr-dest && \
echo "Comparing Nexus3 src content:"
cat /tmp/some-dir/nexus3-n3dr-src.csv | sed -e 1d | cut -d "," -f9- | sha256sum && \
echo "with Nexus3 dest:"
cat /tmp/some-dir/nexus3-n3dr-dest.csv | sed -e 1d | cut -d "," -f9- | sha256sum && \
echo "Check whether the sha256sums are identical."
```
