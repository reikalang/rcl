package rcl

import (
	dlog "github.com/dyweb/gommon/log"
)

const (
	// Extension is file extension for RCL without dot
	Extension = "rcl"
	// DotExtension is '.' + Extension
	DotExtension = ".rcl"
)

var (
	logReg = dlog.NewRegistry()
	log    = logReg.Logger()
)
