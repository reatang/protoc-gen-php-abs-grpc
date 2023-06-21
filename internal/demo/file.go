package demo

import (
	"bytes"
	"path"
	"strings"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	log "github.com/sirupsen/logrus"
)

type FileGen struct {
	File           *FileInfo
	SourceRelative bool
}

func (f *FileGen) GetFile() ([]*plugin.CodeGeneratorResponse_File, error) {
	files := make([]*plugin.CodeGeneratorResponse_File, 0)

	// 生成模板
	tmpl := Template()

	// 开始生成
	w := bytes.NewBufferString("")
	err := tmpl.Execute(w, f.File)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// 文件预处理
	filename := f.Filename(f.File.Name, f.File.Package)
	content := strings.TrimSpace(w.String())

	files = append(files, &plugin.CodeGeneratorResponse_File{
		Name:    &filename,
		Content: &content,
	})

	return files, nil
}

func (f *FileGen) Filename(name, pkg string) string {
	p := strings.ReplaceAll(pkg, ".", "/")

	genFileName := name + ".md"

	if f.SourceRelative {
		return genFileName
	} else {
		return path.Join(p, name+".md")
	}
}
