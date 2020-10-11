# Go Cross
> Cross-compile Go projects with C dependencies.

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/rolandjitsu/go-cross/Test?label=tests&style=flat-square)](https://github.com/rolandjitsu/go-cross/actions?query=workflow%3ATest)

## Prerequisites
Install the following tools:
* [Docker](https://docs.docker.com/engine) >= `19.03.13`
* [buildx](https://github.com/docker/buildx#installing) >= `v0.4.1`
* [msgpack-c](https://github.com/msgpack/msgpack-c/blob/c_master/QUICKSTART-C.md#install) - if building on the host

Enable the experimental features for Docker CLI by adding the following config to `~/.docker/config.json`:
```json
{
    "experimental": "enabled"
}
```

And enable the experimental features for Docker Daemon by adding the following config to the `/etc/docker/daemon.json` file (for Linux; on macOS it's `~/.docker/daemon.json`):
```json
{
    "experimental": true
}
```

## Build

### Docker
Build the `hello` binary with buildx [build](https://github.com/docker/buildx#buildx-build-options-path--url---):
```bash
docker buildx build -f Dockerfile.hello -o type=local,dest=./bin .
```

Or build it with [bake](https://github.com/docker/buildx#buildx-bake-options-target):
```bash
docker buildx bake
```

### Native
Build the `hello` cmd on the current platform/host:
```bash
CGO_LDFLAGS="-lmsgpackc" go build -race -o ./bin/hello cmd/hello/main.go
```

## Test
To run tests for `pkg` run:
```bash
go test ./pkg/...
```

To avoid caching during tests use:
```bash
go test -count=1 ./pkg/...
```

To get coverage reports use the `-cover` flag:
```bash
go test -coverprofile=coverage.out ./pkg/...
```

And to view the profile run:
```bash
go tool cover -html=coverage.out
```

To run static analysis on a package/cmd run:
```bash
go vet ./cmd/hello/main.go
```
