case "$1" in
"")
    # Run docker container
    echo "[run.sh] Docker conatiner is running..."
    docker compose up -d

    # Run server
    echo "[run.sh] Server is running..."
    cd src
    go run main.go
    ;;
"stop")
    # Stop docker container
    echo "[run.sh] Docker conatiner is stopping..."
    docker compose down
    ;;
esac
