#!/usr/bin/env bash
VERSION=$(< ./VERSION)
GITCOMMIT=$(git rev-parse --short HEAD)

if [ -n "$(git status --porcelain --untracked-files=no)" ]; then
  GITCOMMIT="$GITCOMMIT-dirty"
fi

PLATFORMS="linux/amd64 darwin/amd64"
declare -A PLATS=(["linux"]=Linux ["darwin"]=Darwin)

for PLATFORM in $PLATFORMS; do
  GOARCH=${PLATFORM#*/}
  GOOS=${PLATFORM%/*}

  ARCH=`echo "$GOARCH" |sed 's/amd64/x86_64/ig'`
  OS=${PLATS["$GOOS"]}

  CGO_ENABLED=0 GOOS="$GOOS" GOARCH="$GOARCH" \
    go build -a -tags netgo -installsuffix cgo \
    -ldflags "-w -s -X main.GITCOMMIT $GITCOMMIT -X main.VERSION $VERSION" \
    -o "build/horus-$GOOS-$GOARCH" \
    .

  tar --transform "s|horus-$GOOS-$GOARCH|horus|" \
    -vczf "build/horus-$OS-$ARCH.tar.gz" \
    -C build "horus-$GOOS-$GOARCH"
done

# docker build --rm -t cotapreco/horus .
# rm -f build/horus
