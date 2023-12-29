#!/bin/bash -eu

fmt() {
    source script/echo.sh
    echo_log "$0" "Running gofmt..."

    gofmt -w　.
}

lint() {
    source script/echo.sh
    echo_log "$0" "Running golangci-lint..."

    golangci-lint run
}
