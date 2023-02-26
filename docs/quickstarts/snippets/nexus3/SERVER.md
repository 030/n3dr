# Server

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
