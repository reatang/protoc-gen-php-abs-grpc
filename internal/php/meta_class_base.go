package php

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type ClassBase struct {
	pbName string

	namespace string
	prefix    string
	name      string
}

func NewClassBase(pbName, namespace, prefix, name string) *ClassBase {
	// 如果名称包含php的关键字，则报出警告
	if err := isPHPKeyword(name); err != nil {
		panic(err)
	}

	cb := &ClassBase{pbName: pbName, namespace: namespace, prefix: prefix, name: name}
	namesManager.Add(cb)

	return cb
}

func (cb *ClassBase) PbName() string {
	return cb.pbName
}

func (cb *ClassBase) Namespace() string {
	return cb.namespace
}

func (cb *ClassBase) Prefix() string {
	return cb.prefix
}

func (cb *ClassBase) Name() string {
	return cb.name
}

func (cb *ClassBase) FullName() string {
	return fmt.Sprintf("\\%s\\%s%s", cb.namespace, cb.prefix, cb.name)
}

func GetPHPOptions(f *descriptor.FileDescriptorProto) (namespace, prefix string) {

	if f.Options == nil || f.Options.PhpNamespace == nil {
		namespaceList := strings.Split(f.GetPackage(), ".")
		for i, s := range namespaceList {
			namespaceList[i] = FirstUpper(s)
		}
		namespace = strings.Join(namespaceList, "\\")
	} else {
		namespace = *f.Options.PhpNamespace
	}

	if f.Options != nil && f.Options.PhpClassPrefix != nil {
		prefix = *f.Options.PhpClassPrefix
	}

	return
}

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
