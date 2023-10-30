# export env variables
# export $(cat .env | grep -v "#" | xargs)

DATE=$(date '+%Y-%m-%d %H:%M:%S')

case "$1" in
"")
    # Run docker container
    echo "(run.sh) [$DATE] Docker conatiner is running..."
    docker compose up -d

    # Run server
    echo "(run.sh) [$DATE] Backend server is running..."
    cd src
    go run main.go
    ;;
"stop")
    # Stop docker container
    echo "(run.sh) [$DATE] Docker conatiner is stopping..."
    docker compose down
    ;;
"psql")
    # Connect postgresql
    echo "(run.sh) [$DATE] Connecting postgresql..."
    docker exec -it redis-cluster-study-postgres bash
    ;;
esac
