package php

import (
	"fmt"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type MessageClass struct {
	*ClassBase
}

func NewMessageClassWithPB(messagePB *descriptor.DescriptorProto, fPB *descriptor.FileDescriptorProto) *MessageClass {
	pbName := fmt.Sprintf(".%s.%s", fPB.GetPackage(), messagePB.GetName())

	n, p := GetPHPOptions(fPB)

	cb := NewClassBase(pbName, n, p, messagePB.GetName())

	return &MessageClass{
		ClassBase: cb,
	}
}
