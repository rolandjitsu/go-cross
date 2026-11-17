# Go Cross
> Cross-compile Go projects with C dependencies for ARM.

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/rolandjitsu/go-cross/test.yml?branch=master&style=flat-square)](https://github.com/rolandjitsu/go-cross/actions?query=workflow%3ATest)

## Prerequisites
Install the following tools:
* [Docker](https://docs.docker.com/engine) >= `20.10.0`
* [buildx](https://github.com/docker/buildx#installing) >= `v0.8.0`
* [msgpack-c](https://github.com/msgpack/msgpack-c/blob/c_master/QUICKSTART-C.md#install) - if building directly on the host

Modern Docker (v23+) supports BuildKit and multi-platform compilation out of the box without requiring experimental configuration.

## Build

### Docker (Buildx / Bake)
Build for 64-bit ARM (`arm64`):
```bash
docker buildx bake hello-arm64
```

Build for 32-bit ARM (`arm`):
```bash
docker buildx bake hello-arm
```

Or build with standard `buildx build`:
```bash
docker buildx build -f Dockerfile.hello --build-arg TARGETARCH=arm64 -o type=local,dest=./bin/arm64 .
```

### Native
Build the `hello` cmd on your current host architecture:
```bash
CGO_LDFLAGS="-lmsgpackc" go build -race -o ./bin/hello cmd/hello/main.go
```

## Test
To run tests for `pkg`:
```bash
go test ./pkg/...
```

To avoid caching during tests:
```bash
go test -count=1 ./pkg/...
```

To get coverage reports:
```bash
go test -coverprofile=coverage.out ./pkg/...
go tool cover -html=coverage.out
```

To run static analysis:
```bash
go vet ./cmd/hello/main.go
```
