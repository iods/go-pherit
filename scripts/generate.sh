#!/bin/bash

protoc \
    -I. ./proto/*.proto \
    --go_out=./internal/rpc \
    --go-grpc_out=./internal/rpc