#!/usr/bin/env sh

set -e

# install operator-sdk bin
GOPATH=$(go env GOPATH)
mkdir -p $GOPATH/src/github.com/operator-framework
cd $GOPATH/src/github.com/operator-framework
git clone https://github.com/operator-framework/operator-sdk
cd operator-sdk
git checkout master
make dep
make install
