package main

import (
	"github/reatang/protoc-gen-php-abs-grpc/internal/php"
	"github/reatang/protoc-gen-php-abs-grpc/pkg/generator"
	"github/reatang/protoc-gen-php-abs-grpc/pkg/log"
	"github/reatang/protoc-gen-php-abs-grpc/pkg/protoio"
	"os"
)

func main() {
	// 解码参数
	req := protoio.RequestDecode(os.Stdin)
	params := protoio.ParamsMapDecode(req)
	err := log.ConfigureLogging(params)
	if err != nil {
		panic(err)
	}

	if v, ok := params["reqjson"]; ok && v == "true" {
		protoio.RequestJson(req)
	}

	// 注册分析器
	generator.AddAnalyse(php.NewGrpcPHPAnalyseBuilder())

	// 启动生成器
	g := generator.NewProtoFileGenerator(params)

	// 执行生成器
	resp, err := g.Generate(php.Analyse, req)

	// 返回 protoc
	protoio.ResponseEncode(resp)
}
