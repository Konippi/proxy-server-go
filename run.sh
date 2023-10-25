DATE=$(date '+%Y-%m-%d %H:%M:%S')

case "$1" in
"")
    # Run docker container
    echo "(run.sh) [$DATE] Docker conatiner is running..."
    docker compose up -d

    # Run server
    echo "(run.sh) [$DATE] Server is running..."
    cd src
    go run main.go
    ;;
"stop")
    # Stop docker container
    echo "(run.sh) [$DATE] Docker conatiner is stopping..."
    docker compose down
    ;;
esac
