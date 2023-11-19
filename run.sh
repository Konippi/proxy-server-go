# export env variables
# export $(cat .env | grep -v "#" | xargs)

DATE=$(date '+%Y-%m-%d %H:%M:%S')

case "$1" in
"")
    # Run docker container
    echo "(run.sh) [$DATE] Running Docker conatiners..."
    docker compose up -d

    # Run server
    echo "(run.sh) [$DATE] Running backend server..."
    cd cmd/app
    go run main.go
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
"golint")
    # Run golangci-lint
    echo "(run.sh) [$DATE] Running golangci-lint..."
    golangci-lint run
    ;;
esac
