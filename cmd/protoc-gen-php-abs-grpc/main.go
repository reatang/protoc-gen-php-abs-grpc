package main

import (
	"os"

	"github.com/reatang/protoc-gen-php-abs-grpc/internal/php"
	"github.com/reatang/protoc-gen-php-abs-grpc/pkg/generator"
	"github.com/reatang/protoc-gen-php-abs-grpc/pkg/log"
	"github.com/reatang/protoc-gen-php-abs-grpc/pkg/protoio"
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
