#!/bin/sh
cd ../user
go build app.go
cd ../video
go build app.go
cd ../relation
go build app.go
cd ../docker
docker-compose up -d