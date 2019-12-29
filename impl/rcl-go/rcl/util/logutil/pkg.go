// Package logutil is the registry of loggers
package logutil

import (
	"github.com/dyweb/gommon/log"
)

// Registry is the root logger of rcl library
var Registry = log.NewLibraryLogger()

// NewPackageLogger create a new logger for the calling package using Registry as its parent
func NewPackageLogger() *log.Logger {
	l := log.NewPackageLoggerWithSkip(1)
	Registry.AddChild(l)
	return l
}
