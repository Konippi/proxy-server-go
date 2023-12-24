#!/bin/bash -eu

# Run gofmt
function fmt() {
    . echo.sh
    echo_log $0 "Running gofmt..."

    gofmt -w　.
}

# Run golangci-lint
function lint() {
    . echo.sh
    echo_log $0 "Running golangci-lint..."

    golangci-lint run
}
