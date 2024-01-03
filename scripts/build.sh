#!/bin/sh

# force go modules
export GOPATH=""

# disable cgo
export CGO_ENABLED=0

export GOPROXY=https://goproxy.cn

set -e
set -x

# linux
GOOS=linux GOARCH=amd64 go build -o release/linux/amd64/kaniko-docker ./cmd/kaniko-docker
