package rcl

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ RCLVisitor = (*EchoVisitor)(nil)

type EchoVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (e *EchoVisitor) VisitRcl(ctx *RclContext) interface{} {
	pairs := ctx.AllPair()
	for _, pair := range pairs {
		e.VisitPair(pair.(*PairContext))
	}
	return nil
}

func (e *EchoVisitor) VisitPair(ctx *PairContext) interface{} {
	// TODO: use a printer with indent
	fmt.Print(ctx.k.GetText())
	fmt.Print(" ")
	fmt.Print(ctx.Value().Accept(e))
	fmt.Println()
	return nil
}

func (e *EchoVisitor) VisitObj(ctx *ObjContext) interface{} {
	//fmt.Println("VisitObj")
	pairs := ctx.AllPair()
	if len(pairs) == 0 {
		// TODO: how to describe empty object?
		return nil
	}
	for _, pair := range pairs {
		e.VisitPair(pair.(*PairContext))
	}
	return nil
}

func (e *EchoVisitor) VisitArray(ctx *ArrayContext) interface{} {
	items := ctx.AllValue()
	if len(items) == 0 {
		// TODO: how to describe empty array
		return nil
	}
	for _, item := range items {
		item.Accept(e)
	}
	return nil
}

func (e *EchoVisitor) VisitValBool(ctx *ValBoolContext) interface{} {
	// FIXME: handle error
	b, _ := strconv.ParseBool(ctx.GetText())
	return b
}

func (e *EchoVisitor) VisitValInt(ctx *ValIntContext) interface{} {
	// FIXME: handle error
	i, _ := strconv.Atoi(ctx.GetText())
	return i
}

func (e *EchoVisitor) VisitValDouble(ctx *ValDoubleContext) interface{} {
	// FIXME: handle error
	d, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return d
}

func (e *EchoVisitor) VisitValString(ctx *ValStringContext) interface{} {
	return ctx.GetText()
}

func (e *EchoVisitor) VisitValObject(ctx *ValObjectContext) interface{} {
	return e.VisitObj(ctx.Obj().(*ObjContext))
}

func (e *EchoVisitor) VisitValArray(ctx *ValArrayContext) interface{} {
	return e.VisitArray(ctx.Array().(*ArrayContext))
}
