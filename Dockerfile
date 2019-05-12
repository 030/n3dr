FROM golang:1.12.4-alpine as builder
RUN mkdir n3dr && \
    adduser -D -g '' n3dr
COPY main.go go.mod go.sum ./n3dr/
COPY cli ./n3dr/cli
WORKDIR n3dr
RUN ls && \
    apk add git && \
    ls && \
    CGO_ENABLED=0 go build && \
    ls && \
    cp n3dr /n3dr

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /n3dr /usr/local/n3dr
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER n3dr
ENTRYPOINT ["/usr/local/n3dr"]
