#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux go build --ldflags '-extldflags "-static"' -o build/horus .
docker build --rm -t cotapreco/horus .
rm -f build/horus
