#! /bin/sh
GOOS=linux GOARCH=amd64 go build -o "ports-service" cmd/ports-service/main.go

docker build -t local/ports-service .

rm -f ports-service