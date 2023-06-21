#!/bin/bash

proto_file="./testa.proto"

GEN_DIR=./gen

mkdir $GEN_DIR

protoc --reatang-demo_out=$GEN_DIR \
      --reatang-demo_opt=logtostderr=true,loglevel=debug \
      --reatang-demo_opt=reqjson=true \
      --reatang-demo_opt=paths=source_relative \
      --plugin=protoc-gen-reatang-demo=../cmd/protoc-gen-reatang-demo/protoc-gen-reatang-demo \
      $proto_file