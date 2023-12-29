#!/bin/bash -eu

CURRENT_DATETIME="$(date '+%Y-%m-%d %H:%M:%S')"
export CURRENT_DATETIME

case "$1" in
"")
    source script/docker.sh
    run_container
    exit 0
    ;;
"stop")
    source script/docker.sh
    stop_container
    exit 0
    ;;
"psql")
    source script/docker.sh
    connect_psql
    exit 0
    ;;
"lint")
    source script/golang.sh
    lint
    exit 0
    ;;
"fmt")
    source script/golang.sh
    fmt
    exit 0
    ;;
"fmt-lint")
    source script/golang.sh
    fmt
    wait
    lint
    exit 0
    ;;
"gen-cert")
    source script/mkcert.sh
    generate_key_and_certificate
    exit 0
    ;;
esac
