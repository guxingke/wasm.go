package aot

import (
	"fmt"
	"github.com/zxh0/wasm.go/binary"
)

type moduleInfo struct {
	module           binary.Module
	importedFuncs    []binary.Import
	importedTables   []binary.Import
	importedMemories []binary.Import
	importedGlobals  []binary.Import
	globalTypes      []binary.GlobalType
	maxOperandStacks []int
}

func newModuleInfo(module binary.Module) moduleInfo {
	info := moduleInfo{module: module}
	for _, imp := range module.ImportSec {
		switch imp.Desc.Tag {
		case binary.ImportTagFunc:
			info.importedFuncs = append(info.importedFuncs, imp)
		case binary.ImportTagTable:
			info.importedTables = append(info.importedTables, imp)
		case binary.ImportTagMem:
			info.importedMemories = append(info.importedMemories, imp)
		case binary.ImportTagGlobal:
			info.importedGlobals = append(info.importedGlobals, imp)
		}
	}
	return info
}

func (mi moduleInfo) getFuncNameAndType(funcIdx int) (string, binary.FuncType) {
	if funcIdx < len(mi.importedFuncs) {
		imp := mi.importedFuncs[funcIdx]
		ftIdx := imp.Desc.FuncType
		return imp.Name, mi.module.TypeSec[ftIdx]
	}

	ftIdx := mi.module.FuncSec[funcIdx-len(mi.importedFuncs)]
	name := fmt.Sprintf("func#%d", funcIdx)
	return name, mi.module.TypeSec[ftIdx]
}

func getMemPageMin(m binary.Module) int {
	if len(m.MemSec) > 0 {
		return int(m.MemSec[0].Min)
	}
	return 0
}
