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
  cat m3u-proxy.sh | grep '##' | grep -v 'cat' | sed 's/^##//' | sed 's/^ //'
  exit 1
fi

rm -rf  bin
go install m3u-proxy/main.go
mv bin/main bin/m3uproxy

if [[ $action == 'docker-build' ]]; then
   docker build -t m3u-proxy:latest .
fi

