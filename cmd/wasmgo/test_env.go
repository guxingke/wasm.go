package main

import (
	"github.com/zxh0/wasm.go/binary"
	"github.com/zxh0/wasm.go/instance"
)

func newTestEnv() instance.Instance {
	env := instance.NewNativeInstance()
	env.RegisterFunc("assert_eq_i32", assertEqI32, binary.ValTypeI32, binary.ValTypeI32)
	return env
}

func assertEqI32(args ...interface{}) (interface{}, error) {
	return args[0].(int32) == args[1].(int32), nil
}
