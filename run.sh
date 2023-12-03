#!/bin/bash -eu

DATE=$(date '+%Y-%m-%d %H:%M:%S')

case "$1" in
"")
    # Run docker container
    echo "(run.sh) [$DATE] Starting Docker conatiners..."
    docker compose up -d
    ;;
"stop")
    # Stop docker container
    echo "(run.sh) [$DATE] Stopping docker conatiners..."
    docker compose down
    ;;
"psql")
    # Connect postgresql
    echo "(run.sh) [$DATE] Connecting postgresql..."
    docker exec -it redis-cluster-study-postgres bash
    ;;
"lint")
    # Run golangci-lint
    echo "(run.sh) [$DATE] Running golangci-lint..."
    golangci-lint run
    ;;
"fmt")
    # Run gofmt
    echo "(run.sh) [$DATE] Running gofmt..."
    gofmt -w .
    ;;
esac
