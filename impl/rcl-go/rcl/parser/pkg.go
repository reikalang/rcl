package parser // import "github.com/reikalang/rcl/impl/rcl-go/rcl/parser"

import (
	dlog "github.com/dyweb/gommon/log"
)

const (
	Extension    = "rcl"
	DotExtension = ".rcl"
)

var Registry = dlog.NewLibraryLogger()
var log = Registry
