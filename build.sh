#!/usr/bin/env bash

set +x
set +e

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

mkdir -p bin/
export GOBIN="$CURRENT_DIR/bin"

go install src/main.go