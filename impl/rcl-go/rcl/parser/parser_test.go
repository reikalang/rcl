package rcl

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	asst "github.com/stretchr/testify/assert"
)

func TestRCLParser_Rcl(t *testing.T) {
	assert := asst.New(t)
	input, err := antlr.NewFileStream("testdata/db-xephonk.rcl")
	assert.Nil(err)
	lexer := NewRCLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewRCLParser(stream)
	p.BuildParseTrees = true
	// TODO: error listener
	tree := p.Rcl()
	// TODO: visit the tree
	t.Log(tree)
}
