#!/bin/bash

env GOOS=linux GOARCH=arm GOARM=6 go build -o cmd/main.go -o release/pi-temp &&
    tar -czvf release/pi-temp.tar.gz release/pi-temp