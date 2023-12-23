#!/bin/bash -eu

function run_container() {
    . echo.sh
    echo_log $0 "Starting Docker conatiners..."

    docker compose up -d
}

function stop_container() {
    . echo.sh
    echo_log $0 "Stopping Docker conatiners..."

    docker compose down
}

function connect_psql() {
    . echo.sh
    echo_log $0 "Connecting postgresql..."

    docker exec -it redis-cluster-study-postgres bash
}
