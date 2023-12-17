#!/bin/bash -eu

DATE=$(date '+%Y-%m-%d %H:%M:%S')

case "$1" in
"")
    # Run docker container
    echo "(run.sh) [$DATE] Starting Docker conatiners..."
    docker compose up -d
    exit 0
    ;;
"stop")
    # Stop docker container
    echo "(run.sh) [$DATE] Stopping docker conatiners..."
    docker compose down
    exit 0
    ;;
"psql")
    # Connect postgresql
    echo "(run.sh) [$DATE] Connecting postgresql..."
    docker exec -it redis-cluster-study-postgres bash
    exit 0
    ;;
"lint")
    # Run golangci-lint
    echo "(run.sh) [$DATE] Running golangci-lint..."
    golangci-lint run
    exit 0
    ;;
"fmt")
    # Run gofmt
    echo "(run.sh) [$DATE] Running gofmt..."
    gofmt -w .
    exit 0
    ;;
"mkcert")
    # Run mkcert
    echo "(run.sh) [$DATE] Creating a pair of certificate and key..."
    mkcert -install
    mkcert -cert-file ./docker/nginx/ssl/$2.pem -key-file ./docker/nginx/ssl/$2-key.pem $2
    exit 0
    ;;
esac
