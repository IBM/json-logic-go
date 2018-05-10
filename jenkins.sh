#!/bin/sh

# fail if any command below fails
set -e
cd ..
rm -rf gopath
rm -rf go

# download and install Go
echo "***** Installing Go 1.7.5 *****"
wget https://storage.googleapis.com/golang/go1.7.5.linux-amd64.tar.gz -nv
mv go1.7.5.linux-amd64.tar.gz go.tar.gz
tar -xf go.tar.gz

# set up Go env
export GOROOT=$PWD/go
export PATH=$PATH:$GOROOT/bin
mkdir -p gopath
export GOPATH=$PWD/gopath

# set up json-logic-go src in gopath
mkdir -p $GOPATH/src
mv $WORKSPACE $GOPATH/src

# build & run Go tests
cd $GOPATH/src/json-logic-go
echo "***** Building Go code *****"
go build --tags nopkcs11 ./...
echo "***** Running Go tests *****"
go test --tags nopkcs11 $(go list | grep -v vendor)

# cleanup
echo "***** Cleaning up *****"
rm -rf gopath
rm -rf go
