// Package parser is generated from RCL.g4 using ANTLR4.
// DO NOT edit generated files directly, change the grammar is spec folder and make gen-parser
// Visitor implementations are provided, i.e. EchoVisitor pretty print parse tree (not to be confused with AST).
//
// The grammar can be found in https://github.com/reikalang/rcl/blob/master/spec/RCL.g4
package parser // import "github.com/reikalang/rcl/impl/rcl-go/rcl/parser"

import (
	"github.com/reikalang/rcl/impl/rcl-go/rcl/util/logutil"
)

// TODO: error listener
// TODO: collect comments in parser using multiple channels
// TODO: visitor to build ast

var log = logutil.NewPackageLogger()

// Position is same in ast and parser, but we copy because I couldn't figure out a good name for the package
type Position struct {
	Line   int
	Column int
	Offset int
	File   string
}
