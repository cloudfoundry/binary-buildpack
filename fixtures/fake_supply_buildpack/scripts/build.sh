#!/usr/bin/env bash
set -exuo pipefail

GOOS=linux   go build -ldflags="-s -w" -o ./my_buildpack_assets/main     ./my_buildpack_assets/main.go
GOOS=windows go build -ldflags="-s -w" -o ./my_buildpack_assets/main.exe ./my_buildpack_assets/main.go
