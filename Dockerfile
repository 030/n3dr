ARG APPLICATION=n3dr \
    VERSION=0.1.0-rc.1

FROM golang:1.24.5-alpine AS builder
ARG APPLICATION \
    VERSION
RUN adduser -D -g '' ${APPLICATION}
COPY . /go/${APPLICATION}/
WORKDIR /go/${APPLICATION}/cmd/${APPLICATION}
RUN apk add --no-cache \
        git=~2 && \
        CGO_ENABLED=0 go build -ldflags "-X main.Version=${VERSION}" -buildvcs=false && \
        cp ${APPLICATION} /${APPLICATION}

FROM alpine:3.22.0
ARG APPLICATION
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /${APPLICATION} /usr/local/bin/app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
RUN apk add --no-cache \
        libcrypto3=~3 \
        libssl3=~3
USER ${APPLICATION}
ENTRYPOINT ["app"]
