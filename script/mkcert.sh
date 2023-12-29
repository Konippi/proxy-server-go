#!/bin/bash -eu

readonly cert_dir="cert"
readonly domains=(
    "localhost"
)

generate_file_with_generated_date_and_expiration_date() {
    if date -d "next year" &>/dev/null; then
        # GNU: https://www.gnu.org/software/coreutils/manual/html_node/date-invocation.html
        expiration_date=$(date -d "+2 years +3 months" "+%Y-%m-%d %H:%M:%S")
    else
        # BSD: https://man.freebsd.org/cgi/man.cgi?date
        expiration_date=$(date -v+2y -v+3m "+%Y-%m-%d %H:%M:%S")
    fi

    echo "${expiration_date}" >"${cert_dir}"/expiration-date.txt
}

generate_key_and_certificate() {
    source script/echo.sh
    echo_log "$0" "Generating key and certificate..."

    mkcert -install

    ca_root=$(mkcert -CAROOT)
    echo_log "$0" "CA ROOT: \"${ca_root}\""

    for domain in "${domains[@]}"; do
        mkdir -p "${cert_dir}"/"${domain}"
        mkcert -key-file "${cert_dir}"/"${domain}"/"${domain}"-key.pem -cert-file "${cert_dir}"/"${domain}"/"${domain}"-cert.pem "${domain}"
        chmod 600 "${cert_dir}"/"${domain}"/"${domain}"-key.pem "${cert_dir}"/"${domain}"/"${domain}"-cert.pem
        generate_file_with_generated_date_and_expiration_date
    done
}
