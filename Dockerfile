FROM golang:1.12.4-alpine as builder
COPY . ./n3dr/
WORKDIR n3dr
RUN adduser -D -g '' n3dr && \
    apk add git && \
    CGO_ENABLED=0 go build && \
    cp n3dr /n3dr && \
    chmod 100 /n3dr

FROM scratch
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder --chown=n3dr:n3dr /n3dr /usr/local/n3dr
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER n3dr
ENTRYPOINT ["/usr/local/n3dr"]
