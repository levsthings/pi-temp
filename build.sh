#!/bin/bash

env GOOS=linux GOARCH=arm GOARM=6 go build -o release/pi-temp cmd/main.go