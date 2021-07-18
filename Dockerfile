FROM golang:1.16.6-alpine3.14 as builder
ENV PROJECT n3dr
RUN mkdir $PROJECT && \
    adduser -D -g '' $PROJECT
COPY cmd ./$PROJECT/cmd/
COPY internal ./$PROJECT/internal/
COPY go.mod go.sum ./$PROJECT/
WORKDIR $PROJECT/cmd/n3dr
RUN apk add git && \
    CGO_ENABLED=0 go build && \
    cp n3dr /n3dr

FROM alpine:3.13.5
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /n3dr /usr/local/bin/n3dr
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER n3dr
ENTRYPOINT ["/usr/local/bin/n3dr"]
