#!/bin/sh

round=20
# while ! nc -z localhost 2376; do
while ! docker info > /dev/null 2>&1 ; do
    if [ $round -eq 0 ]; then
        >&2 echo "dockerd is not ready"
        exit 1
    fi

    >&2 echo " Waiting for dockerd to be ready..."
    sleep 0.5 # wait for 1/2 of the second before check again
    round=$(( round - 1 ))
done