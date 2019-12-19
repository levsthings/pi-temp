#!/bin/bash

env GOOS=linux GOARCH=arm GOARM=6 go build -o cmd/main.go release/pi-temp