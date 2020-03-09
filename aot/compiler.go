package aot

import (
	"fmt"
	"strings"

	"github.com/zxh0/wasm.go/binary"
)

func Compile(module binary.Module) {
	c := &moduleCompiler{
		printer:    printer{sb: &strings.Builder{}},
		moduleInfo: newModuleInfo(module),
	}
	c.compile()
	fmt.Println(c.sb.String())
}
