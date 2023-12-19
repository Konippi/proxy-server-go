#!/bin/bash -eu

DATE=$(date '+%Y-%m-%d %H:%M:%S')

case "$1" in
"")
    . ./script/docker.sh
    run_container
    exit 0
    ;;
"stop")
    . ./script/docker.sh
    stop_container
    exit 0
    ;;
"psql")
    . ./script/docker.sh
    connect_psql
    exit 0
    ;;
"lint")
    . ./script/golang.sh
    lint
    exit 0
    ;;
"fmt")
    . ./script/golang.sh
    fmt
    exit 0
    ;;
"fmt-lint")
    . ./script/golang.sh
    fmt
    wait
    lint
    exit 0
    ;;
esac
