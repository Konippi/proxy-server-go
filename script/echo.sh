DATE=$(date '+%Y-%m-%d %H:%M:%S')

# arg1: caller file name
# arg2: log message
function echo_log() {
    echo "($1) [$DATE] $2"
}
