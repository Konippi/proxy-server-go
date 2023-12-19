#!/bin/bash -eu

# Run gofmt
function fmt() {
    echo "(run.sh) [$DATE] Running gofmt..."
    gofmt -w　.
}

# Run golangci-lint
function lint() {
    echo "(run.sh) [$DATE] Running golangci-lint..."
    golangci-lint run
}
