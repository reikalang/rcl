package rcl_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

// go_test.go tests behavior of go's std package

func TestGoAST(t *testing.T) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "testdata/foo.go", nil, parser.ParseComments)
	require.Nil(t, err)
	dumpFile, err := os.Create("testdata/foo.dump.txt")
	require.Nil(t, err)
	defer dumpFile.Close()
	spew.Fdump(dumpFile, node)
	ast.Inspect(node, func(node ast.Node) bool {
		if node == nil {
			return false
		}
		pos := fset.Position(node.Pos())
		t.Logf("%s %s", reflect.TypeOf(node), pos)
		return true
	})
}
