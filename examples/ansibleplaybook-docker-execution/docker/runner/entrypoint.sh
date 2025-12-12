#!/bin/sh
set -eu

/usr/local/bin/dockerd-entrypoint.sh 2> /dev/null &
/usr/local/bin/wait-for-dockerd.sh

exec "$@"
