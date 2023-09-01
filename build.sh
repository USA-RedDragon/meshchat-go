#!/bin/bash

GOOS=linux GOARCH=mips GOMIPS=softfloat go build -o out/meshchat-mips -trimpath -ldflags="-s -w"
GOOS=linux GOARCH=arm go build -o out/meshchat-arm -trimpath -ldflags="-s -w"
GOOS=linux GOARCH=arm64 go build -o out/meshchat-arm64 -trimpath -ldflags="-s -w"
GOOS=linux GOARCH=amd64 go build -o out/meshchat-amd64 -trimpath -ldflags="-s -w"
GOOS=windows GOARCH=amd64 go build -o out/meshchat-amd64.exe -trimpath -ldflags="-s -w"
GOOS=linux GOARCH=386 go build -o out/meshchat-x86 -trimpath -ldflags="-s -w"
GOOS=windows GOARCH=386 go build -o out/meshchat-x86.exe -trimpath -ldflags="-s -w"
