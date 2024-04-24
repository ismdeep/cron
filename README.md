# cron

# Quick Start

docker-compose.yaml

```
services:
  cron-service:
    image: ismdeep/cron:latest
    environment:
      TZ: Asia/Shanghai
      CRON_SPEC: '@every 1m'
      CRON_COMMAND: 'date'
      CRON_RUN_AT_START: true
    restart: always
```

