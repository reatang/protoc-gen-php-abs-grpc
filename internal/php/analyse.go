package php

import (
	"bytes"
	"github/reatang/protoc-gen-php-abs-grpc/pkg/generator"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

const Analyse = "php_grpc"
const GenTypeGRPC = "grpc"

// GrpcClassFile PHP box grpc 文件生成器
type GrpcClassFile struct {
	genType string

	Messages []*MessageClass
	Services []*ServiceClass
}

func (c *GrpcClassFile) GetFile() ([]*plugin.CodeGeneratorResponse_File, error) {
	files := make([]*plugin.CodeGeneratorResponse_File, len(c.Services))

	var tmpl *template.Template
	if c.genType == GenTypeGRPC {
		tmpl = GrpcTemplate()
	} else {
		panic("[PHP_ABS_GRPC] genType [" + c.genType + "] don't support")
	}

	// 生成文件
	for i, service := range c.Services {
		w := bytes.NewBufferString("")
		err := tmpl.Execute(w, service)
		if err != nil {
			return nil, err
		}

		fileName := c.GetPHPFileName(service.FullName())
		content := strings.TrimSpace(w.String())

		genFile := &plugin.CodeGeneratorResponse_File{
			Name:           &fileName,
			InsertionPoint: nil,
			Content:        &content,
		}

		files[i] = genFile
	}

	return files, nil
}

func (c *GrpcClassFile) GetPHPFileName(fullClassName string) string {
	baseName := filepath.Base(fullClassName)

	p := path.Join(filepath.Dir(fullClassName), baseName+".php")

	return strings.ReplaceAll(p, "\\", "/")
}

// GrpcPHPAnalyseBuilder PHP分析器注册
type GrpcPHPAnalyseBuilder struct {
}

func (a *GrpcPHPAnalyseBuilder) Schema() string {
	return Analyse
}

func NewGrpcPHPAnalyseBuilder() *GrpcPHPAnalyseBuilder {
	return &GrpcPHPAnalyseBuilder{}
}

func (a *GrpcPHPAnalyseBuilder) Analyse(params map[string]string, proto *descriptor.FileDescriptorProto) generator.GenerateFile {
	if genType, ok := params["genType"]; ok {
		return AnalyseFilePB(genType, proto)
	}

	// 默认生成grpc协议的客户端
	return AnalyseFilePB(GenTypeGRPC, proto)
}

func AnalyseFilePB(genType string, f *descriptor.FileDescriptorProto) *GrpcClassFile {
	cf := &GrpcClassFile{
		genType:  genType,
		Messages: make([]*MessageClass, len(f.MessageType)),
		Services: make([]*ServiceClass, len(f.Service)),
	}

	// 分析 message
	for i, messagePB := range f.MessageType {
		cf.Messages[i] = NewMessageClassWithPB(messagePB, f)
	}

	// 分析 service
	for i, servicePB := range f.Service {
		cf.Services[i] = NewServiceClassByPB(servicePB, f)
	}

	return cf
}
