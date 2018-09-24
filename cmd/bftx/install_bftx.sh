#!/bin/sh

go get github.com/blockfreight/go-bftx
cd $GOPATH/src/github.com/blockfreight/go-bftx
dep ensure

cd $GOPATH/src/github.com/blockfreight/go-bftx/cmd/bftx
go install -v

bftx node start