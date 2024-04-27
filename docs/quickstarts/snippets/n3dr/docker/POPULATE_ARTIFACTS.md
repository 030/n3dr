# populate artifacts

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
