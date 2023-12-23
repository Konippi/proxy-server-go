#!/bin/bash -eu

# Generate CA
GENARATED_DIR="./docker/proxy-server-go/ssl/"
DATE=$(date '+%Y-%m-%d %H:%M:%S')

function generate_file_with_expiration_date() {
    GENERATED_FILE="generated_date.txt"

    . echo.sh
    echo_log $0 "Generating file with expiration date..."

    cd $GENARATED_DIR
    touch $GENERATED_FILE
    echo $DATE >>$GENERATED_FILE

    chmod 444 $GENERATED_FILE
}

function generate_private_key() {
    . echo.sh
    echo_log $0 "Generating private key ..."

    openssl genpkey -alogrithm ed25519 -out ca-private.key
}

function generate_csr() {
    . echo.sh
    echo_log $0 "Generating CSR..."

    openssl req -new -key ca-private.key -out ca.csr -subj "/C=JP/ST=Tokyo/L=Tokyo/O=Example Inc./OU=Web/CN=example.com"
}
