FROM golang:1.20-bullseye AS builder
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /src
COPY . .
RUN set -e; \
    go mod tidy; \
    go build -o ./bin/cron .

FROM debian:11
ENV TZ=Asia/Shanghai
RUN set -e; \
    sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list; \
    sed -i 's/security.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list; \
    apt-get update; \
    apt-get upgrade -y
COPY --from=builder /src/bin/cron /usr/bin/go-cron
ENTRYPOINT ["/usr/bin/go-cron"]
