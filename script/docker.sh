#!/bin/bash -eu

run_container() {
    source script/echo.sh
    echo_log "$0" "Starting Docker conatiners..."

    docker compose up -d
}

stop_container() {
    source script/echo.sh
    echo_log "$0" "Stopping Docker conatiners..."

    docker compose down
}

connect_psql() {
    source script/echo.sh
    echo_log "$0" "Connecting postgresql..."

    docker exec -it redis-cluster-study-postgres bash
}
