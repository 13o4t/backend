#!/bin/bash
set -e

# openapi
go get -u github.com/deepmap/oapi-codegen/cmd/oapi-codegen

# grpc
sudo apt install protobuf-compiler
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

# staticcheck: static check
go get -u honnef.co/go/tools/cmd/staticcheck

# golangci-lint: lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0

# gocyclo: complexity check
go get -u github.com/fzipp/gocyclo/cmd/gocyclo