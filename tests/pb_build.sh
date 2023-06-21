#!/bin/bash

proto_file="./testa.proto"

GEN_DIR=./gen
GRPC_PHP_PLUGIN=${HOME}/protoc/grpc-1.52.1/grpc_php_plugin
PROTO_PATH="-I ${HOME}/protoc/protoc-23.2-osx-aarch_64/include -I ${HOME}/protoc/googleapis -I ./"

mkdir $GEN_DIR

protoc ${PROTO_PATH} \
      --php_out=$GEN_DIR \
      --grpc_out=$GEN_DIR \
      --php-abs-grpc_out=$GEN_DIR \
      --php-abs-grpc_opt=logtostderr=true,loglevel=debug,genType=grpc \
      --plugin=protoc-gen-grpc=$GRPC_PHP_PLUGIN \
      --plugin=protoc-gen-php-abs-grpc=../cmd/protoc-gen-php-abs-grpc/protoc-gen-php-abs-grpc \
      $proto_file