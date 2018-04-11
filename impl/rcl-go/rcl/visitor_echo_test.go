package rcl

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	asst "github.com/stretchr/testify/assert"
)

func TestEchoVisitor_VisitRcl(t *testing.T) {
	assert := asst.New(t)
	input, err := antlr.NewFileStream("testdata/db-xephonk.rcl")
	assert.Nil(err)
	lexer := NewRCLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewRCLParser(stream)
	p.BuildParseTrees = true
	// TODO: error listener
	tree := p.Rcl()
	visitor := NewEchoVisitor()
	res := tree.Accept(visitor)
	t.Log(res)
}

func TestEchoVisitor_Printer(t *testing.T) {
	printer.Println("1")
	printer.Indent()
	printer.Println("2")
	printer.Indent()
	printer.Println("3")
	printer.UnIndent()
	printer.Println("2")
	printer.UnIndent()
	printer.Println("1")
}
