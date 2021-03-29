#!/bin/bash

protoc \
    -I. ./proto/*.proto \
    --go_out=./internal/rpc/services \
    --go-grpc_out=./internal/rpc/services