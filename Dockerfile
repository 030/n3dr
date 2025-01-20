FROM golang:1.23.5-alpine3.20 as builder
ARG VERSION
ENV USERNAME n3dr
RUN adduser -D -g '' $USERNAME
COPY . /go/${USERNAME}/
WORKDIR /go/${USERNAME}/cmd/${USERNAME}
RUN apk add --no-cache \
        git=~2 && \
        CGO_ENABLED=0 go build -ldflags "-X main.Version=${VERSION}" -buildvcs=false && \
        cp n3dr /n3dr

FROM alpine:3.21.2
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /n3dr /usr/local/bin/n3dr
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
RUN apk add --no-cache \
        libcrypto3=~3 \
        libssl3=~3
USER n3dr
ENTRYPOINT ["n3dr"]
