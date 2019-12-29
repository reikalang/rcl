// Package ast defines AST nodes and provides common visitors.
// It is imported by the parser package so packages operating on AST don't need to import ANTLR runtime.
package ast

import (
	"github.com/reikalang/rcl/impl/rcl-go/rcl/util/logutil"
)

var log = logutil.NewPackageLogger()
