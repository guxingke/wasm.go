package main

import (
	"github.com/zxh0/wasm.go/binary"
	"github.com/zxh0/wasm.go/instance"
)

var _ instance.Instance = (*nativeInstance)(nil)

type nativeInstance struct {
	exported map[string]interface{}
}

func newNativeInstance() *nativeInstance {
	return &nativeInstance{
		exported: map[string]interface{}{},
	}
}

func (n *nativeInstance) RegisterFunc(name string,
	ft binary.FuncType, f instance.GoFunc) {

	n.exported[name] = nativeFunction{t: ft, f: f}
}
func (n *nativeInstance) RegisterNoResultsFunc(name string,
	f instance.GoFunc, params ...binary.ValType) {

	ft := binary.FuncType{
		ParamTypes:  params,
		ResultTypes: []binary.ValType{},
	}
	n.RegisterFunc(name, ft, f)
}

func (n *nativeInstance) Register(name string, x interface{}) {
	n.exported[name] = x
}

func (n *nativeInstance) Get(name string) interface{} {
	return n.exported[name]
}

func (n *nativeInstance) CallFunc(name string, args ...interface{}) (interface{}, error) {
	panic("implement me")
}

func (n *nativeInstance) GetGlobalValue(name string) (interface{}, error) {
	panic("implement me")
}
