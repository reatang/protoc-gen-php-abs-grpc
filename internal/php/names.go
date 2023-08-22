package php

import (
	"errors"
	"strings"
)

// pb名映射class名管理器

var ErrNameIsPHPKeyword = errors.New("message or method name is php keyword")

var namesManager *NamesManager

func init() {
	namesManager = &NamesManager{make(map[string]*ClassBase)}
}

type NamesManager struct {
	// map[pb message全名]类名
	names map[string]*ClassBase
}

func (nm *NamesManager) Add(base *ClassBase) {
	nm.names[base.pbName] = base
}

func (nm *NamesManager) Get(pbName string) (*ClassBase, error) {
	if v, ok := nm.names[pbName]; ok {
		return v, nil
	}

	return nil, errors.New("[NamesManager] Not found " + pbName)
}

func (nm *NamesManager) MustGet(pbName string) *ClassBase {
	if v, err := nm.Get(pbName); err == nil {
		return v
	} else {
		panic(err)
	}
}

// 列举了一些常用的可能会被设置成名字的关键字
func isPHPKeyword(name string) error {
	switch strings.ToLower(name) {
	case "isset":
	case "empty":
	case "default":
	case "global":
	case "list":
	case "map":
		return nil
	}

	return ErrNameIsPHPKeyword
}
