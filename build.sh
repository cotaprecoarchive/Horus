#!/usr/bin/env bash
VERSION=$(< ./VERSION)
GITCOMMIT=$(git rev-parse --short HEAD)

if [ -n "$(git status --porcelain --untracked-files=no)" ]; then
  GITCOMMIT="$GITCOMMIT-dirty"
fi

CGO_ENABLED=0 GOOS=linux \
  go build -a \
  -ldflags "-w -s -X main.GITCOMMIT $GITCOMMIT -X main.VERSION $VERSION" \
  -o build/horus \
  .

docker build --rm -t cotapreco/horus .
rm -f build/horus
