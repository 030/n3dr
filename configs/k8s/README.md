# kind

## create secret

```bash
echo -n 'helloworld' | tr -d "\n\r" | base64 -w 0
```

## create cluster

```bash
mkdir -p /tmp/k8s/kind/nexus-data && \
sudo chown 8484 -R /tmp/k8s/kind/nexus-data && \
~/go/bin/kind create cluster --config=configs/k8s/kind.yaml
```

## add ingress

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml && \
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
```

## deploy nexus

```bash
kubectl apply -f configs/k8s/nexus3.yaml
```

## access nexus

open `/etc/hosts`, add:

```bash
127.0.0.1  nexus3.some-domain
```

open a web browser, navigate to: <http://nexus3.some-domain/> and login as
`admin` with password: `helloworld`.

## n3dr

```bash
kubectl logs nexus3-0 -n nexus3
```

## cleanup

```bash
~/go/bin/kind delete cluster
```
