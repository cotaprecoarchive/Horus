#!/usr/bin/env bash
VERSION=$(< ./VERSION)
GITCOMMIT=$(git rev-parse --short HEAD)

if [ -n "$(git status --porcelain --untracked-files=no)" ]; then
  GITCOMMIT="$GITCOMMIT-dirty"
fi

PLATFORMS="linux/386 linux/amd64 darwin/386 darwin/amd64"

for PLATFORM in $PLATFORMS; do
  GOOS=${PLATFORM%/*}
  GOARCH=${PLATFORM#*/}

  CGO_ENABLED=0 GOOS="$GOOS" GOARCH="$GOARCH" \
    go build -a \
    -tags netgo \
    -ldflags "-w -s -X main.GITCOMMIT $GITCOMMIT -X main.VERSION $VERSION" \
    -o "build/horus-$GOOS-$GOARCH" \
    .
  tar -vczf "build/horus-$GOOS-$GOARCH.tar.gz" "build/horus-$GOOS-$GOARCH"
done

# docker build --rm -t cotapreco/horus .
# rm -f build/horus
