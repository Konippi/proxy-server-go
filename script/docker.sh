#!/bin/bash -eu

# Run container
function run_container() {
    echo "(run.sh) [$DATE] Starting Docker conatiners..."
    docker compose up -d
}

# Stop container
function stop_container() {
    echo "(run.sh) [$DATE] Stopping docker conatiners..."
    docker compose down
}

# Connect postgresql
function connect_psql() {
    echo "(run.sh) [$DATE] Connecting postgresql..."
    docker exec -it redis-cluster-study-postgres bash
}
