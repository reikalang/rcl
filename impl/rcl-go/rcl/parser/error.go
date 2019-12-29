package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TODO: error struct and error listener
// https://github.com/antlr/antlr4/blob/master/runtime/Go/antlr/error_listener.go
// I think need to set it on antlr.BaseLexer and antlr.BaseParser
// this logic can actually be generalized and used in other ANTLR based implementation

var _ antlr.ErrorListener = (*ErrorListener)(nil)

type ErrorListener struct {
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {

}

func (l *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {

}

func (l *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {

}

func (l *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {

}
