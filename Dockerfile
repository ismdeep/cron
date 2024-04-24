FROM golang:1.20-bullseye AS builder
WORKDIR /src
COPY . .
RUN set -e; \
    go mod tidy; \
    go build -o ./bin/cron -trimpath -ldflags '-s -w' .

FROM debian:11
RUN set -e; \
    apt-get update; \
    apt-get upgrade -y curl ca-certificates
COPY --from=builder /src/bin/cron /usr/bin/go-cron
ENTRYPOINT ["/usr/bin/go-cron"]
