package generator

import (
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type GenerateFile interface {
	GetFile() ([]*plugin.CodeGeneratorResponse_File, error)
}

type ProtoFileGenerator struct {
	Registry *Registry
}

func NewProtoFileGenerator(params map[string]string) *ProtoFileGenerator {
	return &ProtoFileGenerator{
		Registry: NewRegistry(params),
	}
}

func (p *ProtoFileGenerator) Generate(analyser string, req *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	resp := &plugin.CodeGeneratorResponse{}

	// 分析文件
	genFiles, err := p.Registry.Analyse(analyser, req)
	if err != nil {
		return nil, errors.Wrap(err, "error analysing proto files")
	}

	// 循环生成文件
	for fileName, file := range genFiles {
		// 判断是否要生成
		if !p.Registry.IsFileToGenerate(fileName) {
			continue
		}

		log.Printf("generate %v", fileName)

		// 使用文件和模板生成目标代码
		generated, err := file.GetFile()
		if err != nil {
			log.Debugf("error generating file %v", fileName)
			return nil, errors.Wrap(err, "error generating file")
		}

		for _, responseFile := range generated {
			log.Debugf("output file: %v", responseFile.GetName())
		}

		resp.File = append(resp.File, generated...)
	}

	return resp, nil
}
