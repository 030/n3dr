# Validate

```bash
dir=/tmp/some-dir-validate
mkdir -p "${dir}"
./n3dr count \
  -u admin \
  -p $(docker exec -it nexus3-n3dr-src cat /nexus-data/admin.password) \
  -n localhost:8081 \
  --https=false \
  --csv ${dir}/nexus3-n3dr-src \
  --sort && \
./n3dr count \
  -u admin \
  -p $(docker exec -it nexus3-n3dr-dest cat /nexus-data/admin.password) \
  -n localhost:9000 \
  --https=false \
  --csv ${dir}/nexus3-n3dr-dest \
  --sort && \
echo "Comparing Nexus3 src content:"
cat ${dir}/nexus3-n3dr-src-sorted.csv | sed -e 1d | cut -d "," -f9- | sha256sum && \
echo "with Nexus3 dest:"
cat ${dir}/nexus3-n3dr-dest-sorted.csv | sed -e 1d | cut -d "," -f9- | sha256sum && \
echo "Check whether the sha256sums are identical, e.g.:"
sed -i "s|:8081|:X|g" ${dir}/nexus3-n3dr-src-sorted.csv
sed -i "s|:9000|:X|g" ${dir}/nexus3-n3dr-dest-sorted.csv
cut -d ',' -f1-5,8-11 --output-delimiter=',' ${dir}/nexus3-n3dr-src-sorted.csv > ${dir}/nexus3-n3dr-src-sorted-without-time.csv
cut -d ',' -f1-5,8-11 --output-delimiter=',' ${dir}/nexus3-n3dr-dest-sorted.csv > ${dir}/nexus3-n3dr-dest-sorted-without-time.csv
echo "Number of deviating artifacts between Nexus3 A and B: $(diff ${dir}/nexus3-n3dr-src-sorted-without-time.csv ${dir}/nexus3-n3dr-dest-sorted-without-time.csv | wc -l)"
echo "vimdiff ${dir}/nexus3-n3dr-src-sorted-without-time.csv ${dir}/nexus3-n3dr-dest-sorted-without-time.csv"
```
