package php

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	log "github.com/sirupsen/logrus"
)

func PtrString(s string) *string {
	return &s
}

func PtrBool(b bool) *bool {
	return &b
}

func TestGrpcClassFile_GetFile(t *testing.T) {
	req := &descriptor.DescriptorProto{
		Name: PtrString("Req"),
	}
	resp := &descriptor.DescriptorProto{
		Name: PtrString("Resp"),
	}

	f := &descriptor.FileDescriptorProto{
		Package: PtrString("reatang.grpc.phptest"),
		Options: &descriptor.FileOptions{
			PhpNamespace:         PtrString("Reatang\\Grpc\\Phptest"),
			PhpMetadataNamespace: PtrString("Reatang\\Grpc\\Phptest"),
			PhpClassPrefix:       PtrString("PB\\"),
			PhpGenericServices:   PtrBool(true),
		},
		MessageType: []*descriptor.DescriptorProto{
			req,
			resp,
		},
		Service: []*descriptor.ServiceDescriptorProto{
			{
				Name: PtrString("PHPTest"),
				Method: []*descriptor.MethodDescriptorProto{
					{
						Name:       PtrString("Ping"),
						InputType:  PtrString(".reatang.grpc.phptest.Req"),
						OutputType: PtrString(".reatang.grpc.phptest.Resp"),
					},
				},
			},
		},
	}

	file := AnalyseFilePB("", f)

	getFile, err := file.GetFile()
	if err != nil {
		log.Fatal(err)
	}

	for _, responseFile := range getFile {
		fmt.Println(responseFile.GetName())
		fmt.Println("------------")
		fmt.Println(responseFile.GetContent())
	}
}
