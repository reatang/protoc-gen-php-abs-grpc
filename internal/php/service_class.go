package php

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
)

// 标准的 service class

// 支持 box_grpc 的 service class

// ServiceMethodArgument stores the type information about method argument
type ServiceMethodArgument struct {
	// Type is the type of the argument
	Type *MessageClass
	// IsExternal indicate if this type is an external dependency
	IsExternal bool
	// IsRepeated indicates whether the field is a repeated field
	IsRepeated bool
}

type GoogleApiHttpOption struct {
	Path string
}

type ServiceMethod struct {
	Name   string
	Input  *ServiceMethodArgument
	Output *ServiceMethodArgument

	GoogleApiHttpOption
}

type ServiceClass struct {
	*ClassBase

	GrpcFullClassName string
	GrpcClassName     string

	Methods []*ServiceMethod
}

func NewServiceClassByPB(servicePB *descriptor.ServiceDescriptorProto, fPB *descriptor.FileDescriptorProto) *ServiceClass {
	pbName := fmt.Sprintf(".%s.%s", fPB.GetPackage(), servicePB.GetName())

	n, p := GetPHPOptions(fPB)
	if strings.Contains(p, "\\") {
		p = ""
	}

	cb := NewClassBase(pbName, n, p, servicePB.GetName()+"AbsRpc")

	grpcClient := servicePB.GetName() + "Client"
	s := &ServiceClass{
		ClassBase: cb,

		GrpcClassName: grpcClient,

		Methods: make([]*ServiceMethod, 0),
	}

	for _, serviceMethod := range servicePB.GetMethod() {
		s.AddMethod(serviceMethod)
	}

	return s
}

func (s *ServiceClass) AddMethod(methodPB *descriptor.MethodDescriptorProto) {
	m := &ServiceMethod{
		Name: methodPB.GetName(),
	}

	if methodPB.InputType != nil {
		m.Input = &ServiceMethodArgument{
			Type: &MessageClass{ClassBase: namesManager.MustGet(*methodPB.InputType)},
		}
	}

	if methodPB.InputType != nil {
		m.Output = &ServiceMethodArgument{
			Type: &MessageClass{ClassBase: namesManager.MustGet(*methodPB.OutputType)},
		}
	}

	if methodPB.GetOptions() != nil {
		if opt := proto.GetExtension(methodPB.GetOptions(), annotations.E_Http); opt != nil {
			if rule, ok := opt.(*annotations.HttpRule); ok {
				m.GoogleApiHttpOption = GoogleApiHttpOption{Path: rule.GetPost()}
			}
		}
	}

	s.Methods = append(s.Methods, m)
}

func (s *ServiceClass) GetUses() []MessageClass {
	classNames := make([]MessageClass, 0)
	once := make(map[string]struct{})

	for _, method := range s.Methods {
		if method.Input != nil {
			if _, ok := once[method.Input.Type.FullName()]; !ok {
				classNames = append(classNames, *method.Input.Type)
				once[method.Input.Type.FullName()] = struct{}{}
			}
		}

		if method.Output != nil {
			if _, ok := once[method.Output.Type.FullName()]; !ok {
				classNames = append(classNames, *method.Output.Type)
				once[method.Output.Type.FullName()] = struct{}{}
			}
		}
	}

	return classNames
}
