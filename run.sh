#!/bin/sh

ENV_FILE=newrelic.env
IMAGE=suse-demo

if ! [ -f "$ENV_FILE" ]; then
    echo "File $ENV_FILE not found"
    exit
fi

BUILT=`docker images -q $IMAGE`
if [ -z "$BUILT" ]; then
    echo "Docker Image $IMAGE not found, run build.sh first"
    exit
fi

docker run --env-file=$ENV_FILE --name suse-agent-demo -p 8080:8080 $IMAGE
