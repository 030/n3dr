# Another server

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
