#!/bin/sh
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o dingtalk-liker_Windows_i386.exe dingtalk-liker.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dingtalk-liker_Windows_amd64.exe dingtalk-liker.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dingtalk-liker_Linux_amd64 dingtalk-liker.go
