#!/bin/bash
GIT_COMMIT="$(git rev-parse --verify HEAD)"
BUILD_DATE="$(date -u '+%Y-%m-%d_%I:%M:%S%p')"
echo "Using git commit hash: ${GIT_COMMIT}"

echo "Building Linux binary..."
CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-w -s -X main.builddate=${BUILD_DATE} -X main.githash=${GIT_COMMIT}" -o bin/gowww .

echo "Building OS X binary..."
CGO_ENABLED=0 GOOS=darwin go build -a -ldflags "-w -s -X main.builddate=${BUILD_DATE} -X main.githash=${GIT_COMMIT}" -o bin/gowww-darwin .
