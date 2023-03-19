FROM golang:1.20.2-alpine3.17 as builder
ARG VERSION
ENV USERNAME n3dr
RUN adduser -D -g '' $USERNAME
COPY . /go/${USERNAME}/
WORKDIR /go/${USERNAME}/cmd/${USERNAME}
RUN apk add --no-cache \
  git=~2 && \
  CGO_ENABLED=0 go build -ldflags "-X main.Version=${VERSION}" -buildvcs=false && \
  cp n3dr /n3dr

FROM alpine:3.17.2
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /n3dr /usr/local/bin/n3dr
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER n3dr
ENTRYPOINT ["/usr/local/bin/n3dr"]
