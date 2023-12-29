#!/bin/bash -eu

# arg1: caller file name
# arg2: log message
echo_log() {
    echo "($1) [${CURRENT_DATETIME}] $2"
}
