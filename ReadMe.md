# gRPC-Gateway Setup

This project uses [gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway) to generate RESTful APIs and OpenAPI (Swagger) documentation from `.proto` files.

## Environment

- **Go Version**: `go1.24.3 linux/amd64`
- **OS**: Linux
- **Architecture**: amd64

## Prerequisites

Ensure you have the following tools installed:

- [Go](https://go.dev/dl/) (already installed: `1.24.3`)
- [Protocol Buffers Compiler (protoc)](https://github.com/protocolbuffers/protobuf/releases)
- Your Go environment (`GOPATH`, `GOBIN`, etc.) configured

### Set Go Environment (if not already configured)

```bash
export GOPATH=/home/huan/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
