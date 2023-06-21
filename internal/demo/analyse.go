package demo

import (
	"github/reatang/protobuf-plugin-template/pkg/generator"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

const DemoAnalyse = "demo"

type DemoAnalyseBuilder struct{}

func NewDemoAnalyseBuilder() *DemoAnalyseBuilder {
	return &DemoAnalyseBuilder{}
}

func (a *DemoAnalyseBuilder) Schema() string {
	return DemoAnalyse
}

func (a *DemoAnalyseBuilder) Analyse(params map[string]string, proto *descriptor.FileDescriptorProto) generator.GenerateFile {

	// 分析 FileDescriptorProto 转换为自己可用的FileInfo
	fileInfo := &FileInfo{
		Package:  proto.GetPackage(),
		Name:     proto.GetName(),
		Messages: make([]*MessageInfo, len(proto.MessageType)),
		Services: make([]*ServiceInfo, len(proto.Service)),
	}

	// 分析 message
	for i, descriptorProto := range proto.MessageType {
		fileInfo.Messages[i] = &MessageInfo{Name: descriptorProto.GetName()}
	}

	// 分析 service
	for i, serviceDescriptorProto := range proto.Service {
		service := &ServiceInfo{
			Name:    serviceDescriptorProto.GetName(),
			Methods: make([]*ServiceMethodInfo, len(serviceDescriptorProto.Method)),
		}

		for i2, methodDescriptorProto := range serviceDescriptorProto.Method {
			service.Methods[i2] = &ServiceMethodInfo{
				Name:   methodDescriptorProto.GetName(),
				Input:  methodDescriptorProto.GetInputType(),
				Output: methodDescriptorProto.GetOutputType(),
			}
		}

		fileInfo.Services[i] = service
	}

	// 转交给生成器
	fileGen := &FileGen{File: fileInfo}
	if paths, ok := params["paths"]; ok && paths == "source_relative" {
		fileGen.SourceRelative = true
	}

	return fileGen
}
