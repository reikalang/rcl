// Package parser is generated from RCL.g4 using ANTLR4.
// DO NOT edit generated files directly, change the grammar is spec folder and make gen-parser
// Visitor implementations are provided, i.e. EchoVisitor pretty print parse tree (not to be confused with AST).
//
// The grammar can be found in https://github.com/reikalang/rcl/blob/master/spec/RCL.g4
package parser // import "github.com/reikalang/rcl/impl/rcl-go/rcl/parser"

import (
	"github.com/reikalang/rcl/impl/rcl-go/rcl/util/logutil"
)

// TODO: might move global variables to rcl package instead of put it in parser
const (
	Extension    = "rcl"
	DotExtension = ".rcl"
)

var log = logutil.NewPackageLogger()
