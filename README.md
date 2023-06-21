# protoc-gen-php-abs-grpc

支持php grpc包 `https://github.com/reatang/grpc-php-abstract` 的客户端生成器

## 使用

```shell
# 生成目录
GEN_DIR=./grpcgen
# php grpc生成器，如果有新的需更改
GRPC_PHP_PLUGIN=${HOME}/protoc/grpc-1.52.1/grpc_php_plugin
# protobuf官方库 和 google 扩展库，如果有新的需更改
PROTO_PATH="-I ${HOME}/protoc/protoc-23.2-osx-aarch_64/include -I ${HOME}/protoc/googleapis -I ./"

mkdir $GEN_DIR

protoc ${PROTO_PATH} \
      --php_out=$GEN_DIR \ # 生成 php 数据类
      --grpc_out=$GEN_DIR \ # 生成 php grpc 客户端
      --php-abs-grpc_out=$GEN_DIR \ # 本项目的 php 兼容客户端
      --php-abs-grpc_opt=logtostderr=true,loglevel=debug,genType=grpc \ # 可选的生成参数
      --plugin=protoc-gen-grpc=$GRPC_PHP_PLUGIN \ # 导入生成器
      --plugin=protoc-gen-php-abs-grpc=${GOPATH}/bin/protoc-gen-php-abs-grpc \ # 导入生成器
      $protofile

```