package rcl

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ RCLVisitor = (*EchoVisitor)(nil)

type EchoVisitor struct {
	*antlr.BaseParseTreeVisitor
	p *indentPrinter
}

func NewEchoVisitor() *EchoVisitor {
	return &EchoVisitor{
		p: &indentPrinter{w: os.Stdout},
	}
}

func (e *EchoVisitor) VisitRcl(ctx *RclContext) interface{} {
	pairs := ctx.AllPair()
	for _, pair := range pairs {
		e.VisitPair(pair.(*PairContext))
	}
	return nil
}

func (e *EchoVisitor) VisitPair(ctx *PairContext) interface{} {
	e.p.Print(ctx.k.GetText())
	fmt.Print(" ")
	fmt.Print(ctx.Value().Accept(e))
	fmt.Println()
	return nil
}

func (e *EchoVisitor) VisitObj(ctx *ObjContext) interface{} {
	//fmt.Println("VisitObj")
	e.p.Indent()
	//e.p.Println()
	pairs := ctx.AllPair()
	if len(pairs) == 0 {
		// TODO: how to describe empty object?
		return nil
	}
	for _, pair := range pairs {
		e.VisitPair(pair.(*PairContext))
	}
	//e.p.Println()
	e.p.UnIndent()
	return nil
}

func (e *EchoVisitor) VisitArray(ctx *ArrayContext) interface{} {
	e.p.Indent()
	e.p.Println()
	items := ctx.AllValue()
	if len(items) == 0 {
		// TODO: how to describe empty array
		log.Info("empty array")
		return nil
	}
	for _, item := range items {
		item.Accept(e)
	}
	//e.p.Println()
	e.p.UnIndent()
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

const indentUnit = 4

// global instance used for test only
var printer = &indentPrinter{w: os.Stdout}

type indentPrinter struct {
	indentSize int
	w          io.Writer

	indent string
	arg    []interface{}
}

func (p *indentPrinter) Indent() {
	// TODO: might allow config indent
	p.indentSize += indentUnit
	p.updateCache()
}

func (p *indentPrinter) UnIndent() {
	if p.indentSize < indentUnit {
		return
	}
	p.indentSize -= indentUnit
	p.updateCache()
}

func (p *indentPrinter) updateCache() {
	b := make([]byte, p.indentSize)
	for i := 0; i < p.indentSize; i++ {
		b[i] = ' '
	}
	p.indent = string(b)
	t := make([]interface{}, 1)
	t[0] = p.indent
	p.arg = t
}

func (p *indentPrinter) Print(args ...interface{}) {
	p.w.Write([]byte(p.indent + fmt.Sprint(args...)))
}

func (p *indentPrinter) Println(args ...interface{}) {
	p.w.Write([]byte(p.indent + fmt.Sprintln(args...)))
}

func (p *indentPrinter) Printf(format string, args ...interface{}) {
	p.w.Write([]byte(p.indent + fmt.Sprintf(format, args...)))
}
