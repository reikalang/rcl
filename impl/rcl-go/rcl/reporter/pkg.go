// Package reporter defines interface for error collector
// TODO: its subpackage provides sinks in common format, text, json TODO: language server
// ref https://github.com/at15/reika/tree/master/impl/reika-j/src/main/java/me/at15/reika/compiler/reporter
package reporter

// Position is used in both ast and parser
type Position struct {
	Line   int
	Column int
	Offset int
	File   string
}

// ref https://github.com/Microsoft/language-server-protocol/blob/gh-pages/specification.md#range
type Range struct {
	Start Position
	End Position
}

type ErrorType string

const (
	SyntaxError ErrorType = "syntax"
	// TODO: semantic error?
)

type Error interface {
	Type() ErrorType
	Pos() Position
}

type ErrorCollector interface {
	Reset()
	HasError() bool
	ErrorsCount() int
	Errors() []Error
}
