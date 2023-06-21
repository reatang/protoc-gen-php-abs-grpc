package protoio

import (
	"encoding/json"
	"io"
	"os"
	"strings"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

// RequestDecode 解码protoc传来的语法数据
func RequestDecode(read io.Reader) *plugin.CodeGeneratorRequest {
	req := &plugin.CodeGeneratorRequest{}
	data, err := io.ReadAll(read)
	if err != nil {
		panic(err)
	}
	err = proto.Unmarshal(data, req)
	if err != nil {
		panic(err)
	}
	return req
}

// ParamsMapDecode 解码opt参数
func ParamsMapDecode(req *plugin.CodeGeneratorRequest) map[string]string {
	paramsMap := make(map[string]string)
	params := req.GetParameter()

	for _, p := range strings.Split(params, ",") {
		if i := strings.Index(p, "="); i < 0 {
			paramsMap[p] = ""
		} else {
			paramsMap[p[0:i]] = p[i+1:]
		}
	}

	return paramsMap
}

// ResponseEncode 返回编码后的内容
func ResponseEncode(resp proto.Message) {
	data, err := proto.Marshal(resp)
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		panic(err)
	}
}

func RequestJson(req *plugin.CodeGeneratorRequest) {
	if log.GetLevel() != log.DebugLevel {
		return
	}

	o, err := os.Create("req.json")
	if err != nil {
		panic(err)
	}
	bytes, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		panic(err)
	}
	_, err = o.Write(bytes)
	if err != nil {
		panic(err)
	}
}
