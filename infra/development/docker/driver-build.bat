@echo off
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
if not exist build mkdir build
go build -buildvcs=false -o build/driver-service ./services/driver-service/
