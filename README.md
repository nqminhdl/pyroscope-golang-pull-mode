# Demo Profiling for Golang Pull Mode with Grafana-Agent

## Build docker image

```bash
cd src/
docker build -t <tag_name>:0.0.1 .
```

## Update docker compose file

```yaml
services:
  demo:
    image: '<update_image_build_here>'
```

## Start docker compose

```bash
docker-compose up
```

## How to explore profiling data

Access Grafana portal <http://localhost:3000/explore>

Input service name `{service_name="demo"}` and explore profile.
