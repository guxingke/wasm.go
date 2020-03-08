package text

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/zxh0/wasm.go/text/parser"
)

var _ parser.WASTVisitor = (*baseVisitor)(nil)

type baseVisitor struct {
}

/* Code generated by IDEA */

func (v baseVisitor) Visit(tree antlr.ParseTree) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitScript(ctx *parser.ScriptContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitCmd(ctx *parser.CmdContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitWastModule(ctx *parser.WastModuleContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitAction_(ctx *parser.Action_Context) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitAssertion(ctx *parser.AssertionContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitExpected(ctx *parser.ExpectedContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitMeta(ctx *parser.MetaContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitModule(ctx *parser.ModuleContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitWatModule(ctx *parser.WatModuleContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitModuleField(ctx *parser.ModuleFieldContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitTypeDef(ctx *parser.TypeDefContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitImport_(ctx *parser.Import_Context) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitImportDesc(ctx *parser.ImportDescContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitFunc_(ctx *parser.Func_Context) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitFuncLocal(ctx *parser.FuncLocalContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitTable(ctx *parser.TableContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitMemory(ctx *parser.MemoryContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitGlobal(ctx *parser.GlobalContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitExport(ctx *parser.ExportContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitExportDesc(ctx *parser.ExportDescContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitElem(ctx *parser.ElemContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitData(ctx *parser.DataContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitEmbeddedIm(ctx *parser.EmbeddedImContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitEmbeddedEx(ctx *parser.EmbeddedExContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitTypeUse(ctx *parser.TypeUseContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitFuncVars(ctx *parser.FuncVarsContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitValType(ctx *parser.ValTypeContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitBlockType(ctx *parser.BlockTypeContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitGlobalType(ctx *parser.GlobalTypeContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitMemoryType(ctx *parser.MemoryTypeContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitTableType(ctx *parser.TableTypeContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitElemType(ctx *parser.ElemTypeContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitLimits(ctx *parser.LimitsContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitFuncType(ctx *parser.FuncTypeContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitParam(ctx *parser.ParamContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitResult(ctx *parser.ResultContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitExpr(ctx *parser.ExprContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitInstr(ctx *parser.InstrContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitFoldedInstr(ctx *parser.FoldedInstrContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitBlockInstr(ctx *parser.BlockInstrContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitPlainInstr(ctx *parser.PlainInstrContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitConstInstr(ctx *parser.ConstInstrContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitMemArg(ctx *parser.MemArgContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitNat(ctx *parser.NatContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitValue(ctx *parser.ValueContext) interface{} {
	panic("implement me")
}

func (v baseVisitor) VisitVariable(ctx *parser.VariableContext) interface{} {
	panic("implement me")
}