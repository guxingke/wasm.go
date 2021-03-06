package binary

import (
	"encoding/binary"
	"math"
	"reflect"
)

type WasmWriter struct {
	buf []byte
}

func Encode(module Module) []byte {
	writer := &WasmWriter{}
	writer.writeModule(module)
	return writer.buf
}

func (writer *WasmWriter) writeByte(b byte) {
	writer.buf = append(writer.buf, b)
}
func (writer *WasmWriter) writeU32(n uint32) {
	writer.buf = append(writer.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(writer.buf[len(writer.buf)-4:], n)
}
func (writer *WasmWriter) writeF32(f float32) {
	writer.writeU32(math.Float32bits(f))
}
func (writer *WasmWriter) writeF64(f float64) {
	writer.buf = append(writer.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	n := math.Float64bits(f)
	binary.LittleEndian.PutUint64(writer.buf[len(writer.buf)-8:], n)
}

func (writer *WasmWriter) writeLen(n int) {
	writer.writeVarU32(uint32(n))
}
func (writer *WasmWriter) writeVarU32(n uint32) {
	data := encodeVarUint(uint64(n), 32)
	writer.buf = append(writer.buf, data...)
}
func (writer *WasmWriter) writeVarS32(n int32) {
	data := encodeVarInt(int64(n), 32)
	writer.buf = append(writer.buf, data...)
}
func (writer *WasmWriter) writeVarS64(n int64) {
	data := encodeVarInt(n, 64)
	writer.buf = append(writer.buf, data...)
}

func (writer *WasmWriter) writeBytes(data []byte) {
	writer.writeLen(len(data))
	writer.buf = append(writer.buf, data...)
}
func (writer *WasmWriter) writeName(name string) {
	writer.writeBytes([]byte(name))
}

func (writer *WasmWriter) writeModule(module Module) []byte {
	writer.writeU32(MagicNumber)
	writer.writeU32(Version)
	writer.writeSec(1, module.TypeSec)
	writer.writeSec(2, module.ImportSec)
	writer.writeSec(3, module.FuncSec)
	writer.writeSec(4, module.TableSec)
	writer.writeSec(5, module.MemSec)
	writer.writeSec(6, module.GlobalSec)
	writer.writeSec(7, module.ExportSec)
	if module.StartSec != nil {
		writer.writeByte(8)
		writer.writeVarU32(*module.StartSec)
	}
	writer.writeSec(9, module.ElemSec)
	writer.writeSec(10, module.CodeSec)
	writer.writeSec(11, module.DataSec)
	//writer.writeSec(0, module.CustomSecs)
	return writer.buf
}

func (writer *WasmWriter) writeSec(id byte, vec interface{}) {
	val := reflect.ValueOf(vec)
	if val.Len() > 0 {
		writer.writeByte(id)
		secWriter := &WasmWriter{}
		secWriter.writeVec(vec)
		writer.writeBytes(secWriter.buf)
	}
}

func (writer *WasmWriter) writeVec(vec interface{}) {
	val := reflect.ValueOf(vec)
	writer.writeLen(val.Len())
	for i := 0; i < val.Len(); i++ {
		writer.writeScala(val.Index(i).Interface())
	}
}

func (writer *WasmWriter) writeScala(val interface{}) {
	switch x := val.(type) {
	case Import:
		writer.writeImport(x)
	case Global:
		writer.writeGlobal(x)
	case Export:
		writer.writeExport(x)
	case Elem:
		writer.writeElem(x)
	case Code:
		writer.writeCode(x)
	case Data:
		writer.writeData(x)
	case FuncType:
		writer.writeFuncType(x)
	case TableType:
		writer.writeTableType(x)
	case Limits:
		writer.writeLimits(x)
	case byte:
		writer.writeByte(x)
	case uint32:
		writer.writeVarU32(x)
	default:
		panic("unreachable")
	}
}

func (writer *WasmWriter) writeImport(imp Import) {
	writer.writeName(imp.Module)
	writer.writeName(imp.Name)
	writer.writeByte(imp.Desc.Tag)
	switch imp.Desc.Tag {
	case ImportTagFunc:
		writer.writeVarU32(imp.Desc.FuncType)
	case ImportTagTable:
		writer.writeTableType(imp.Desc.Table)
	case ImportTagMem:
		writer.writeLimits(imp.Desc.Mem)
	case ImportTagGlobal:
		writer.writeGlobalType(imp.Desc.Global)
	}
}
func (writer *WasmWriter) writeGlobal(global Global) {
	writer.writeGlobalType(global.Type)
	writer.writeExpr(global.Expr)
}
func (writer *WasmWriter) writeExport(export Export) {
	writer.writeName(export.Name)
	writer.writeByte(export.Desc.Tag)
	writer.writeVarU32(export.Desc.Idx)
}
func (writer *WasmWriter) writeElem(elem Elem) {
	writer.writeVarU32(elem.Table)
	writer.writeExpr(elem.Offset)
	writer.writeVec(elem.Init)
}
func (writer *WasmWriter) writeCode(code Code) {
	codeWriter := &WasmWriter{}
	codeWriter.writeLen(len(code.Locals))
	for _, locals := range code.Locals {
		codeWriter.writeVarU32(locals.N)
		codeWriter.writeByte(locals.Type)
	}
	codeWriter.writeExpr(code.Expr)
	writer.writeBytes(codeWriter.buf)
}
func (writer *WasmWriter) writeData(data Data) {
	writer.writeVarU32(data.Mem)
	writer.writeExpr(data.Offset)
	writer.writeBytes(data.Init)
}

// types
func (writer *WasmWriter) writeBlockType(bt BlockType) {
	if len(bt) == 0 {
		writer.writeByte(NoVal)
	} else {
		writer.writeByte(bt[0])
	}
}
func (writer *WasmWriter) writeFuncType(ft FuncType) {
	writer.writeByte(FtTag)
	writer.writeVec(ft.ParamTypes)
	writer.writeVec(ft.ResultTypes)
}
func (writer *WasmWriter) writeTableType(tt TableType) {
	writer.writeByte(tt.ElemType)
	writer.writeLimits(tt.Limits)
}
func (writer *WasmWriter) writeGlobalType(gt GlobalType) {
	writer.writeByte(gt.ValType)
	writer.writeByte(gt.Mut)
}
func (writer *WasmWriter) writeLimits(limits Limits) {
	writer.writeByte(limits.Tag)
	writer.writeVarU32(limits.Min)
	if limits.Tag == 1 {
		writer.writeVarU32(limits.Max)
	}
}

func (writer *WasmWriter) writeExpr(expr Expr) {
	for _, instr := range expr {
		writer.writeInstr(instr)
	}
	writer.writeByte(_End)
}

func (writer *WasmWriter) writeInstr(instr Instruction) {
	writer.writeByte(instr.Opcode)
	switch instr.Opcode {
	case Block, Loop:
		args := instr.Args.(BlockArgs)
		writer.writeBlockType(args.RT)
		writer.writeExpr(args.Instrs)
	case If:
		args := instr.Args.(IfArgs)
		writer.writeBlockType(args.RT)
		writer.writeExpr(args.Instrs1)
		writer.buf[len(writer.buf)-1] = _Else
		writer.writeExpr(args.Instrs2)
	case Br, BrIf:
		writer.writeVarU32(instr.Args.(uint32))
	case BrTable:
		args := instr.Args.(BrTableArgs)
		writer.writeVec(args.Labels)
		writer.writeVarU32(args.Default)
	case Call:
		writer.writeVarU32(instr.Args.(uint32))
	case CallIndirect:
		writer.writeVarU32(instr.Args.(uint32))
		writer.writeByte(0)
	case LocalGet, LocalSet, LocalTee, GlobalGet, GlobalSet:
		writer.writeVarU32(instr.Args.(uint32))
	case MemorySize, MemoryGrow:
		writer.writeByte(0)
	case I32Const:
		writer.writeVarS32(instr.Args.(int32))
	case I64Const:
		writer.writeVarS64(instr.Args.(int64))
	case F32Const:
		writer.writeF32(instr.Args.(float32))
	case F64Const:
		writer.writeF64(instr.Args.(float64))
	default:
		if memArg, ok := instr.Args.(MemArg); ok {
			writer.writeVarU32(memArg.Align)
			writer.writeVarU32(memArg.Offset)
		}
	}
}
