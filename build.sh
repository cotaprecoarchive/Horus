#!/usr/bin/env bash
# gox -os="linux" -ldflags '-w' --output build/horus
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w' -o build/horus .
docker build --rm -t cotapreco/horus .
rm -f build/horus
