#!/bin/bash

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap main.go

./bootstrap &

# Get Bootstrap PID
BOOTSTRAP_PID=$!

# # Get Bootstrap PID (Alternative)
# BOOTSTRAP_PID=$(ps -ef | grep '\.\/bootstrap' | grep -v 'grep' | awk '{print $2}')

# Wait until 120 seconds, and then kill the webserver process
sleep 120
kill $BOOTSTRAP_PID

rm -rf bootstrap