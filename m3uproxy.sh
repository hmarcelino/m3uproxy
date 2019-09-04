#!/usr/bin/env bash

## m3uproxy.sh is a shell script that will help you achieve several tasks:
##
## Possible actions:
## * build
## * docker-build
##
## Examples:
## * ./m3uproxy.sh build - build the proxy (it will require go locally)
## * ./m3uproxy.sh docker-build - build a docker container with m3u proxy
##

set +x
set +e

CURRENT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

mkdir -p bin/
export GOBIN="$CURRENT_DIR/bin"

action=$1

if [[ $action == "" || ($action != 'build' && $action != 'docker-build') ]]; then
  # print usage information
  echo "No action or invalid action specifed"
  cat m3uproxy.sh | grep '##' | grep -v 'cat' | sed 's/^##//' | sed 's/^ //'
  exit 1
fi

rm -rf bin

if [[ $action == 'build' ]]; then
  go install m3uproxy/main.go
  mv bin/main bin/m3uproxy

elif [[ $action == 'docker-build' ]]; then
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install m3uproxy/main.go
  mv bin/main bin/m3uproxy
  docker build -t m3uproxy:latest .
fi
