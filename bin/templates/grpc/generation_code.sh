#!/bin/bash 

ADAPT_PATH=./

clear
protoc --go_out=${ADAPT_PATH} --nrpc_out=./ --go_opt=paths=import --go-grpc_out=${ADAPT_PATH} --go-grpc_opt=paths=import ${ADAPT_PATH}Sync_service.proto
