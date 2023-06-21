package generator

import (
	"errors"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

var analysers = make(map[string]Builder)

type Builder interface {
	Schema() string
	Analyse(params map[string]string, proto *descriptor.FileDescriptorProto) GenerateFile
}

func AddAnalyse(ab Builder) {
	analysers[ab.Schema()] = ab
}

func GetAnalyse(schema string) (Builder, error) {
	if ab, ok := analysers[schema]; ok {
		return ab, nil
	} else {
		return nil, errors.New("不存在的 Analyse: " + schema)
	}
}
