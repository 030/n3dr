# development

## go-swagger

Run:

```bash
N3DR_CLEAN_IN_CASE_OF_SUCCESS_OR_FAILURE=false ./test/integration-tests.sh
```

Once Nexus had been started, download the go-swagger, swagger.json and
generate internal go-swagger code:

```bash
export GITHUB_URL=https://github.com
export GS_URI=go-swagger/go-swagger/releases/download
export GS_VERSION=v0.29.0
export GS_URL=${GITHUB_URL}/${GS_URI}/${GS_VERSION}/swagger_linux_amd64
curl -L \
  ${GS_URL} \
  -o swagger
chmod +x swagger
mkdir -p internal/goswagger
curl http://localhost:9999/service/rest/swagger.json -o swagger.json
./swagger generate client \
  --name=nexus3 \
  --spec swagger.json \
  --target=internal/app/n3dr/goswagger \
  --skip-validation
go mod tidy
```

## Unit Tests

```bash
go test internal/artifacts/common.go internal/artifacts/common_test.go
```
