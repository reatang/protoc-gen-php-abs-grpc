# protobuf-plugin-template
基于golang构建 protobuf plugin 的基本脚手架

## 模块划分

- 分析器
- 内容生成器

## 功能描述

`generator`调用`分析器`返回`内容生成器`构造文件内容，并交由protoc创建`文件`  

## 开发流程

1. 先编写`analyse`，将`protoset`信息转换为自定义的`meta`结构体
2. 再编写`file`生成器，将自定义的格式转换为特定的文件内容