# Docker

Download the [latest N3DR binary](https://github.com/030/n3dr/releases/tag/7.1.0):

```bash
cd /tmp && \
curl -L https://github.com/030/n3dr/releases/download/7.1.0/n3dr-ubuntu-latest \
  -o n3dr-ubuntu-latest && \
curl -L https://github.com/030/n3dr/releases/download/7.1.0/\
n3dr-ubuntu-latest.sha512.txt \
  -o n3dr-ubuntu-latest.sha512.txt && \
sha512sum -c n3dr-ubuntu-latest.sha512.txt && \
chmod +x n3dr-ubuntu-latest && \
mv n3dr-ubuntu-latest n3dr && \
./n3dr --version
```

Start a Nexus3 server:

```bash
nexus_docker_name=nexus3-n3dr-src
docker run \
  --rm \
  -d \
  -p 8081:8081 \
  -p 8082:8082 \
  --name ${nexus_docker_name} \
  sonatype/nexus3:3.47.1 && \
until docker logs ${nexus_docker_name} | grep -q 'Started Sonatype Nexus OSS'; do
  sleep 10
done
```

Create a repository:

```bash
./n3dr configRepository \
  -u admin \
  -p $(docker exec -it nexus3-n3dr-src cat /nexus-data/admin.password) \
  -n localhost:8081 \
  --https=false \
  --configRepoName docker-images \
  --configRepoType docker
```

Populate it with artifacts:

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

Backup all artifacts:

```bash
./n3dr repositoriesV2 \
  --backup \
  --directory-prefix /tmp/some-dir \
  -u admin \
  -p $(docker exec -it nexus3-n3dr-src cat /nexus-data/admin.password) \
  --dockerPort 8082 \
  --dockerHost http://localhost \
  -n localhost:8081 \
  --https=false
```

Start another nexus3 server:

```bash
nexus_docker_name=nexus3-n3dr-dest
docker run \
  --rm \
  -d \
  -p 9000:8081 \
  -p 9001:8082 \
  --name ${nexus_docker_name} \
  sonatype/nexus3:3.47.1 && \
until docker logs ${nexus_docker_name} | grep -q 'Started Sonatype Nexus OSS'; do
  sleep 10
done
```

Create a repository in the other nexus server:

```bash
./n3dr configRepository \
  -u admin \
  -p $(docker exec -it nexus3-n3dr-dest cat /nexus-data/admin.password) \
  -n localhost:9000 \
  --https=false \
  --configRepoName docker-images \
  --configRepoType docker
```

Upload the artifacts to the other nexus server:

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

Validate:

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

Cleanup:

```bash
docker stop nexus3-n3dr-src nexus3-n3dr-dest
```
