# development

## go-swagger

Run:

```bash
N3DR_CLEAN_IN_CASE_OF_SUCCESS_OR_FAILURE=false ./test/integration-tests.sh
```

Once Nexus had been started, download the swagger.json:

```bash
curl http://localhost:8081/service/rest/swagger.json -o \
  configs/swagger/nexus3.json
```

Lookup the `maven2.asset3` json snippet, septuple it and change it to:
`maven2.asset4`, `maven2.asset5`, `maven2.asset6`, `maven2.asset7`,
`maven2.asset8`, `maven2.asset9`, `maven2.asset10` and `maven2.asset11` respectively. After
adding the seven snippets and renaming them, download go-swagger and generate
internal go-swagger code:

```bash
export GITHUB_URL=https://github.com
export GS_URI=go-swagger/go-swagger/releases/download
export GS_VERSION=v0.30.4
export GS_URL=${GITHUB_URL}/${GS_URI}/${GS_VERSION}/swagger_linux_amd64
export GS_DIR=internal/app/n3dr/goswagger
curl -L \
  ${GS_URL} \
  -o swagger
chmod +x swagger
mkdir -p "${GS_DIR}"
./swagger generate client \
  --name=nexus3 \
  --spec configs/swagger/nexus3.json \
  --target="${GS_DIR}" \
  --skip-validation && \
go mod tidy && \
rm swagger*
```

## Unit Tests

```bash
go test internal/artifacts/common.go internal/artifacts/common_test.go
```
