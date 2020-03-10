package instance

import (
	"github.com/zxh0/wasm.go/binary"
)

var _ Instance = (*NativeInstance)(nil)

type NativeInstance struct {
	exported map[string]interface{}
}

func NewNativeInstance() *NativeInstance {
	return &NativeInstance{
		exported: map[string]interface{}{},
	}
}

func (n *NativeInstance) RegisterFunc(name string,
	ft binary.FuncType, f GoFunc) {

	n.exported[name] = nativeFunction{t: ft, f: f}
}
func (n *NativeInstance) RegisterNoResultsFunc(name string,
	f GoFunc, params ...binary.ValType) {

	ft := binary.FuncType{
		ParamTypes:  params,
		ResultTypes: []binary.ValType{},
	}
	n.RegisterFunc(name, ft, f)
}

func (n *NativeInstance) Register(name string, x interface{}) {
	n.exported[name] = x
}

func (n *NativeInstance) Get(name string) interface{} {
	return n.exported[name]
}

func (n *NativeInstance) CallFunc(name string, args ...interface{}) (interface{}, error) {
	panic("implement me")
}

func (n *NativeInstance) GetGlobalValue(name string) (interface{}, error) {
	panic("implement me")
}
