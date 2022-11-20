#!/bin/sh

GOOS=linux GOARCH=amd64 go build -o limebackendtester-linux
GOOS=darwin GOARCH=amd64 go build -o limebackendtester-osx