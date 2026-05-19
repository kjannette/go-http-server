# Go HTTP server

A simple Go HTTP server boilerplate with Echo, structured logging, graceful shutdown, and Docker support.

## Prerequisites

- Go 1.25+
- Docker and Docker Compose (for containerized runs)

## Configuration

Copy the example environment file and adjust values if needed:

```bash
cp .env.dist .env
```

The application reads configuration from process environment variables (via [envconfig](https://github.com/kelseyhightower/envconfig)). Docker Compose loads `.env` automatically. For a local `go run`, export variables or source the file:

```bash
set -a && source .env && set +a
```

| Variable | Default | Description |
|----------|---------|-------------|
| `LOG_LEVEL` | `info` | Zap log level (`debug`, `info`, `warn`, `error`) |
| `API_PORT` | `3000` | HTTP listen port |
| `API_READ_TIMEOUT` | `7` | Server read timeout (seconds) |
| `API_WRITE_TIMEOUT` | `5` | Server write timeout (seconds) |
| `API_IDLE_TIMEOUT` | `5` | Server idle timeout (seconds) |
| `API_TIMEOUT` | `5` | Reserved for future use |

Inside Docker, the process always listens on port `3000`; `API_PORT` in `.env` controls the host port mapping.

## Run with Docker (recommended)

```bash
make setup-docker   # copy .env, download modules, build and start containers
```

Or step by step:

```bash
make setup-local    # copy .env (if missing) and go mod download
make run            # docker compose up --build
```

Verify:

```bash
curl http://localhost:3000/health
```

Stop:

```bash
make down
```

## Run locally (without Docker)

```bash
make setup-local
make run-local
```

## Build binary

```bash
make build
./bin/server
```

## Make targets

```bash
make help
```

Common targets: `run`, `run-local`, `down`, `setup-local`, `setup-docker`, `build`, `tests`, `lint`.
