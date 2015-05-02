#!/bin/bash -e
ORG_PATH="github.com/CotaPreco"
REPO_PATH="${ORG_PATH}/Horus"

rm -rf $GOPATH/src/$REPO_PATH
mkdir -p $GOPATH/src/$ORG_PATH
cp -r /app $GOPATH/src/$REPO_PATH

eval $(go env)

# cross
cd $GOROOT/src
GOOS=darwin GOARCH=amd64 ./make.bash

cd $GOPATH/src/$REPO_PATH
go get ./...
bash ./build.sh
rm -rf /app/build/* && mkdir -p /app/build
mv $GOPATH/src/$REPO_PATH/build/* /app/build
rm -rf $GOPATH/src/$REPO_PATH/build/*

exit
