#!/bin/sh
set -e

if [[ $(ls | grep -m 1 'go.mod') != 'go.mod' ]]; then
    echo "run build.sh in project root"
    echo "bash deployment/build.sh"
    exit 1
fi

docker build \
    -f deployment/server/Dockerfile \
    -t motion-server:latest .

docker build \
    -f deployment/web-client/Dockerfile \
    -t motion-web-client:latest .
