package generator

import (
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	log "github.com/sirupsen/logrus"
)

type Registry struct {
	// 分析参数
	Params map[string]string

	FilesToGenerate map[string]bool
}

// NewRegistry 初始化一个插件注册工具
func NewRegistry(params map[string]string) *Registry {
	r := &Registry{
		Params: params,
	}

	return r
}

// IsFileToGenerate 只生成输入的文件
func (r *Registry) IsFileToGenerate(name string) bool {
	result, ok := r.FilesToGenerate[name]
	return ok && result
}

// Analyse 将proto数据解析为 预生成 php file 的数据
func (r *Registry) Analyse(schema string, req *plugin.CodeGeneratorRequest) (map[string]GenerateFile, error) {
	r.FilesToGenerate = make(map[string]bool)
	for _, f := range req.GetFileToGenerate() {
		r.FilesToGenerate[f] = true
	}

	files := req.GetProtoFile()
	log.Debugf("about to start anaylyse files, %d in total", len(files))

	analyseBuilder, err := GetAnalyse(schema)
	if err != nil {
		return nil, err
	}

	dataFiles := make(map[string]GenerateFile)
	// analyse all files in the request first
	for _, f := range files {
		dataFiles[f.GetName()] = analyseBuilder.Analyse(r.Params, f)
	}

	return dataFiles, nil
}
