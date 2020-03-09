package aot

import "github.com/zxh0/wasm.go/binary"

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
