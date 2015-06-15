#!/bin/bash

export PATH=$PATH:$GOPATH/src/github.com/orian/protoc-gen-gojsgrpc

go build
#protoc --jsgrpc=testdata/ testdata/multi/multi{1,2,3}.proto --proto_path=testdata/
protoc --proto_path=testdata/ --jsgrpc_out=plugins=jsgrpc:testdata/ testdata/multi/multi{1,2,3}.proto

protoc --proto_path=testdata/ --jsgrpc_out=plugins=jsgrpc:testdata/jsgrpc/testing testdata/grpc.proto

