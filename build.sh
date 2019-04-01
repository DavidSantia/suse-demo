#!/bin/sh

PROJECT=github.com/DavidSantia/suse-demo

if ! [ -d $GOPATH/src/$PROJECT ]; then
    echo "Project $PROJECT not found"
    exit
fi

APP=sles12-nr/demo-server

# Build for Linux, statically linked
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $APP
ls -l $APP

# Build docker image
docker build -t suse-demo sles12-nr

