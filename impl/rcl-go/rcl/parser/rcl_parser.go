// Code generated from spec/RCL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // RCL
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 16, 64, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 6, 2, 14,
	10, 2, 13, 2, 14, 2, 15, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 24,
	10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 30, 10, 4, 12, 4, 14, 4, 33, 11, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 39, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5,
	45, 10, 5, 12, 5, 14, 5, 48, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 54, 10,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 62, 10, 6, 3, 6, 2, 2, 7,
	2, 4, 6, 8, 10, 2, 2, 2, 69, 2, 13, 3, 2, 2, 2, 4, 23, 3, 2, 2, 2, 6, 38,
	3, 2, 2, 2, 8, 53, 3, 2, 2, 2, 10, 61, 3, 2, 2, 2, 12, 14, 5, 4, 3, 2,
	13, 12, 3, 2, 2, 2, 14, 15, 3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 15, 16, 3,
	2, 2, 2, 16, 3, 3, 2, 2, 2, 17, 18, 7, 13, 2, 2, 18, 19, 7, 3, 2, 2, 19,
	24, 5, 10, 6, 2, 20, 21, 7, 12, 2, 2, 21, 22, 7, 3, 2, 2, 22, 24, 5, 10,
	6, 2, 23, 17, 3, 2, 2, 2, 23, 20, 3, 2, 2, 2, 24, 5, 3, 2, 2, 2, 25, 26,
	7, 4, 2, 2, 26, 31, 5, 4, 3, 2, 27, 28, 7, 5, 2, 2, 28, 30, 5, 4, 3, 2,
	29, 27, 3, 2, 2, 2, 30, 33, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3,
	2, 2, 2, 32, 34, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 34, 35, 7, 6, 2, 2, 35,
	39, 3, 2, 2, 2, 36, 37, 7, 4, 2, 2, 37, 39, 7, 6, 2, 2, 38, 25, 3, 2, 2,
	2, 38, 36, 3, 2, 2, 2, 39, 7, 3, 2, 2, 2, 40, 41, 7, 7, 2, 2, 41, 46, 5,
	10, 6, 2, 42, 43, 7, 5, 2, 2, 43, 45, 5, 10, 6, 2, 44, 42, 3, 2, 2, 2,
	45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 49, 3,
	2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 50, 7, 8, 2, 2, 50, 54, 3, 2, 2, 2, 51,
	52, 7, 7, 2, 2, 52, 54, 7, 8, 2, 2, 53, 40, 3, 2, 2, 2, 53, 51, 3, 2, 2,
	2, 54, 9, 3, 2, 2, 2, 55, 62, 7, 9, 2, 2, 56, 62, 7, 10, 2, 2, 57, 62,
	7, 11, 2, 2, 58, 62, 7, 12, 2, 2, 59, 62, 5, 6, 4, 2, 60, 62, 5, 8, 5,
	2, 61, 55, 3, 2, 2, 2, 61, 56, 3, 2, 2, 2, 61, 57, 3, 2, 2, 2, 61, 58,
	3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 60, 3, 2, 2, 2, 62, 11, 3, 2, 2, 2,
	9, 15, 23, 31, 38, 46, 53, 61,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'{'", "','", "'}'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "BOOL", "INT", "DOUBLE", "STRING", "ID", "WS",
	"BLOCK_COMMENT", "SINGLE_COMMENT",
}

var ruleNames = []string{
	"rcl", "pair", "obj", "array", "value",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type RCLParser struct {
	*antlr.BaseParser
}

func NewRCLParser(input antlr.TokenStream) *RCLParser {
	this := new(RCLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "RCL.g4"

	return this
}

// RCLParser tokens.
const (
	RCLParserEOF            = antlr.TokenEOF
	RCLParserT__0           = 1
	RCLParserT__1           = 2
	RCLParserT__2           = 3
	RCLParserT__3           = 4
	RCLParserT__4           = 5
	RCLParserT__5           = 6
	RCLParserBOOL           = 7
	RCLParserINT            = 8
	RCLParserDOUBLE         = 9
	RCLParserSTRING         = 10
	RCLParserID             = 11
	RCLParserWS             = 12
	RCLParserBLOCK_COMMENT  = 13
	RCLParserSINGLE_COMMENT = 14
)

// RCLParser rules.
const (
	RCLParserRULE_rcl   = 0
	RCLParserRULE_pair  = 1
	RCLParserRULE_obj   = 2
	RCLParserRULE_array = 3
	RCLParserRULE_value = 4
)

// IRclContext is an interface to support dynamic dispatch.
type IRclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRclContext differentiates from other interfaces.
	IsRclContext()
}

type RclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRclContext() *RclContext {
	var p = new(RclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCLParserRULE_rcl
	return p
}

func (*RclContext) IsRclContext() {}

func NewRclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RclContext {
	var p = new(RclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCLParserRULE_rcl

	return p
}

func (s *RclContext) GetParser() antlr.Parser { return s.parser }

func (s *RclContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *RclContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *RclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCLVisitor:
		return t.VisitRcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCLParser) Rcl() (localctx IRclContext) {
	localctx = NewRclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RCLParserRULE_rcl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RCLParserSTRING || _la == RCLParserID {
		{
			p.SetState(10)
			p.Pair()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetK returns the k token.
	GetK() antlr.Token

	// SetK sets the k token.
	SetK(antlr.Token)

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	k      antlr.Token
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCLParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCLParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) GetK() antlr.Token { return s.k }

func (s *PairContext) SetK(v antlr.Token) { s.k = v }

func (s *PairContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) ID() antlr.TerminalNode {
	return s.GetToken(RCLParserID, 0)
}

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(RCLParserSTRING, 0)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCLVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCLParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RCLParserRULE_pair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(21)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RCLParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(15)

			var _m = p.Match(RCLParserID)

			localctx.(*PairContext).k = _m
		}
		{
			p.SetState(16)
			p.Match(RCLParserT__0)
		}
		{
			p.SetState(17)
			p.Value()
		}

	case RCLParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(18)

			var _m = p.Match(RCLParserSTRING)

			localctx.(*PairContext).k = _m
		}
		{
			p.SetState(19)
			p.Match(RCLParserT__0)
		}
		{
			p.SetState(20)
			p.Value()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObjContext is an interface to support dynamic dispatch.
type IObjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjContext differentiates from other interfaces.
	IsObjContext()
}

type ObjContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjContext() *ObjContext {
	var p = new(ObjContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCLParserRULE_obj
	return p
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCLParserRULE_obj

	return p
}

func (s *ObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *ObjContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *ObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCLVisitor:
		return t.VisitObj(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCLParser) Obj() (localctx IObjContext) {
	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RCLParserRULE_obj)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.Match(RCLParserT__1)
		}
		{
			p.SetState(24)
			p.Pair()
		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RCLParserT__2 {
			{
				p.SetState(25)
				p.Match(RCLParserT__2)
			}
			{
				p.SetState(26)
				p.Pair()
			}

			p.SetState(31)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(32)
			p.Match(RCLParserT__3)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.Match(RCLParserT__1)
		}
		{
			p.SetState(35)
			p.Match(RCLParserT__3)
		}

	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCLParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCLParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ArrayContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCLVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCLParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RCLParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Match(RCLParserT__4)
		}
		{
			p.SetState(39)
			p.Value()
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RCLParserT__2 {
			{
				p.SetState(40)
				p.Match(RCLParserT__2)
			}
			{
				p.SetState(41)
				p.Value()
			}

			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(47)
			p.Match(RCLParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Match(RCLParserT__4)
		}
		{
			p.SetState(50)
			p.Match(RCLParserT__5)
		}

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCLParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValBoolContext struct {
	*ValueContext
}

func NewValBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValBoolContext {
	var p = new(ValBoolContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ValBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValBoolContext) BOOL() antlr.TerminalNode {
	return s.GetToken(RCLParserBOOL, 0)
}

func (s *ValBoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCLVisitor:
		return t.VisitValBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValObjectContext struct {
	*ValueContext
}

func NewValObjectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValObjectContext {
	var p = new(ValObjectContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ValObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValObjectContext) Obj() IObjContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *ValObjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCLVisitor:
		return t.VisitValObject(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValIntContext struct {
	*ValueContext
}

func NewValIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValIntContext {
	var p = new(ValIntContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ValIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValIntContext) INT() antlr.TerminalNode {
	return s.GetToken(RCLParserINT, 0)
}

func (s *ValIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCLVisitor:
		return t.VisitValInt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValStringContext struct {
	*ValueContext
}

func NewValStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValStringContext {
	var p = new(ValStringContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ValStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(RCLParserSTRING, 0)
}

func (s *ValStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCLVisitor:
		return t.VisitValString(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValDoubleContext struct {
	*ValueContext
}

func NewValDoubleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValDoubleContext {
	var p = new(ValDoubleContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ValDoubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValDoubleContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(RCLParserDOUBLE, 0)
}

func (s *ValDoubleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCLVisitor:
		return t.VisitValDouble(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValArrayContext struct {
	*ValueContext
}

func NewValArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValArrayContext {
	var p = new(ValArrayContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ValArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValArrayContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ValArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCLVisitor:
		return t.VisitValArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RCLParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(59)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RCLParserBOOL:
		localctx = NewValBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Match(RCLParserBOOL)
		}

	case RCLParserINT:
		localctx = NewValIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Match(RCLParserINT)
		}

	case RCLParserDOUBLE:
		localctx = NewValDoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.Match(RCLParserDOUBLE)
		}

	case RCLParserSTRING:
		localctx = NewValStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.Match(RCLParserSTRING)
		}

	case RCLParserT__1:
		localctx = NewValObjectContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(57)
			p.Obj()
		}

	case RCLParserT__4:
		localctx = NewValArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(58)
			p.Array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
