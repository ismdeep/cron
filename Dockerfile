FROM debian:13 AS builder
WORKDIR /src/
COPY ./build .
RUN \
    case $(uname -m) in \
    amd64|x86_64)  cp cron_linux_amd64 cron ;; \
    arm64|aarch64) cp cron_linux_arm64 cron ;; \
    *) echo "[ERROR] unsupported platform: $(uname -m)" && false ;; \
    esac

FROM debian:13
RUN \
    apt-get update && \
    apt-get upgrade -y curl ca-certificates
COPY --from=builder /src/cron /usr/bin/go-cron
ENTRYPOINT ["/usr/bin/go-cron"]
