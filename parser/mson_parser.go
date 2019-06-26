// Code generated from Mson.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Mson
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 93, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 7, 2, 20, 10, 2, 12, 2, 14, 2, 23, 11, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 7, 3, 29, 10, 3, 12, 3, 14, 3, 32, 11, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 5, 3, 38, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 7, 5, 48, 10, 5, 12, 5, 14, 5, 51, 11, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 5, 5, 57, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 5, 6, 69, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 86, 10, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 2,
	2, 99, 2, 21, 3, 2, 2, 2, 4, 37, 3, 2, 2, 2, 6, 39, 3, 2, 2, 2, 8, 56,
	3, 2, 2, 2, 10, 68, 3, 2, 2, 2, 12, 70, 3, 2, 2, 2, 14, 85, 3, 2, 2, 2,
	16, 87, 3, 2, 2, 2, 18, 20, 5, 10, 6, 2, 19, 18, 3, 2, 2, 2, 20, 23, 3,
	2, 2, 2, 21, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 3, 3, 2, 2, 2, 23,
	21, 3, 2, 2, 2, 24, 25, 7, 3, 2, 2, 25, 30, 5, 6, 4, 2, 26, 27, 7, 4, 2,
	2, 27, 29, 5, 6, 4, 2, 28, 26, 3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28,
	3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 33, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2,
	33, 34, 7, 5, 2, 2, 34, 38, 3, 2, 2, 2, 35, 36, 7, 3, 2, 2, 36, 38, 7,
	5, 2, 2, 37, 24, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 38, 5, 3, 2, 2, 2, 39,
	40, 7, 17, 2, 2, 40, 41, 7, 6, 2, 2, 41, 42, 5, 10, 6, 2, 42, 7, 3, 2,
	2, 2, 43, 44, 7, 7, 2, 2, 44, 49, 5, 10, 6, 2, 45, 46, 7, 4, 2, 2, 46,
	48, 5, 10, 6, 2, 47, 45, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2,
	2, 2, 49, 50, 3, 2, 2, 2, 50, 52, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 53,
	7, 8, 2, 2, 53, 57, 3, 2, 2, 2, 54, 55, 7, 7, 2, 2, 55, 57, 7, 8, 2, 2,
	56, 43, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 57, 9, 3, 2, 2, 2, 58, 69, 7, 17,
	2, 2, 59, 69, 7, 18, 2, 2, 60, 69, 5, 12, 7, 2, 61, 69, 5, 14, 8, 2, 62,
	69, 5, 16, 9, 2, 63, 69, 5, 4, 3, 2, 64, 69, 5, 8, 5, 2, 65, 69, 7, 9,
	2, 2, 66, 69, 7, 10, 2, 2, 67, 69, 7, 11, 2, 2, 68, 58, 3, 2, 2, 2, 68,
	59, 3, 2, 2, 2, 68, 60, 3, 2, 2, 2, 68, 61, 3, 2, 2, 2, 68, 62, 3, 2, 2,
	2, 68, 63, 3, 2, 2, 2, 68, 64, 3, 2, 2, 2, 68, 65, 3, 2, 2, 2, 68, 66,
	3, 2, 2, 2, 68, 67, 3, 2, 2, 2, 69, 11, 3, 2, 2, 2, 70, 71, 7, 12, 2, 2,
	71, 72, 7, 13, 2, 2, 72, 73, 7, 18, 2, 2, 73, 74, 7, 4, 2, 2, 74, 75, 7,
	17, 2, 2, 75, 76, 7, 14, 2, 2, 76, 13, 3, 2, 2, 2, 77, 78, 7, 15, 2, 2,
	78, 79, 7, 13, 2, 2, 79, 80, 7, 18, 2, 2, 80, 86, 7, 14, 2, 2, 81, 82,
	7, 15, 2, 2, 82, 83, 7, 13, 2, 2, 83, 84, 7, 17, 2, 2, 84, 86, 7, 14, 2,
	2, 85, 77, 3, 2, 2, 2, 85, 81, 3, 2, 2, 2, 86, 15, 3, 2, 2, 2, 87, 88,
	7, 16, 2, 2, 88, 89, 7, 13, 2, 2, 89, 90, 7, 17, 2, 2, 90, 91, 7, 14, 2,
	2, 91, 17, 3, 2, 2, 2, 9, 21, 30, 37, 49, 56, 68, 85,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "','", "'}'", "':'", "'['", "']'", "'true'", "'false'", "'null'",
	"'BinData'", "'('", "')'", "'NumberLong'", "'ISODate'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "STRING", "NUMBER",
	"WS",
}

var ruleNames = []string{
	"mson", "obj", "pair", "array", "value", "binData", "numberLong", "isoDate",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type MsonParser struct {
	*antlr.BaseParser
}

func NewMsonParser(input antlr.TokenStream) *MsonParser {
	this := new(MsonParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Mson.g4"

	return this
}

// MsonParser tokens.
const (
	MsonParserEOF    = antlr.TokenEOF
	MsonParserT__0   = 1
	MsonParserT__1   = 2
	MsonParserT__2   = 3
	MsonParserT__3   = 4
	MsonParserT__4   = 5
	MsonParserT__5   = 6
	MsonParserT__6   = 7
	MsonParserT__7   = 8
	MsonParserT__8   = 9
	MsonParserT__9   = 10
	MsonParserT__10  = 11
	MsonParserT__11  = 12
	MsonParserT__12  = 13
	MsonParserT__13  = 14
	MsonParserSTRING = 15
	MsonParserNUMBER = 16
	MsonParserWS     = 17
)

// MsonParser rules.
const (
	MsonParserRULE_mson       = 0
	MsonParserRULE_obj        = 1
	MsonParserRULE_pair       = 2
	MsonParserRULE_array      = 3
	MsonParserRULE_value      = 4
	MsonParserRULE_binData    = 5
	MsonParserRULE_numberLong = 6
	MsonParserRULE_isoDate    = 7
)

// IMsonContext is an interface to support dynamic dispatch.
type IMsonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMsonContext differentiates from other interfaces.
	IsMsonContext()
}

type MsonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsonContext() *MsonContext {
	var p = new(MsonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MsonParserRULE_mson
	return p
}

func (*MsonContext) IsMsonContext() {}

func NewMsonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MsonContext {
	var p = new(MsonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsonParserRULE_mson

	return p
}

func (s *MsonContext) GetParser() antlr.Parser { return s.parser }

func (s *MsonContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *MsonContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *MsonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MsonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.EnterMson(s)
	}
}

func (s *MsonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.ExitMson(s)
	}
}

func (p *MsonParser) Mson() (localctx IMsonContext) {
	localctx = NewMsonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MsonParserRULE_mson)
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MsonParserT__0)|(1<<MsonParserT__4)|(1<<MsonParserT__6)|(1<<MsonParserT__7)|(1<<MsonParserT__8)|(1<<MsonParserT__9)|(1<<MsonParserT__12)|(1<<MsonParserT__13)|(1<<MsonParserSTRING)|(1<<MsonParserNUMBER))) != 0 {
		{
			p.SetState(16)
			p.Value()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = MsonParserRULE_obj
	return p
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsonParserRULE_obj

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

func (s *ObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.EnterObj(s)
	}
}

func (s *ObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.ExitObj(s)
	}
}

func (p *MsonParser) Obj() (localctx IObjContext) {
	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MsonParserRULE_obj)
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

	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Match(MsonParserT__0)
		}
		{
			p.SetState(23)
			p.Pair()
		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MsonParserT__1 {
			{
				p.SetState(24)
				p.Match(MsonParserT__1)
			}
			{
				p.SetState(25)
				p.Pair()
			}

			p.SetState(30)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(31)
			p.Match(MsonParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.Match(MsonParserT__0)
		}
		{
			p.SetState(34)
			p.Match(MsonParserT__2)
		}

	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MsonParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsonParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) GetKey() antlr.Token { return s.key }

func (s *PairContext) SetKey(v antlr.Token) { s.key = v }

func (s *PairContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(MsonParserSTRING, 0)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *MsonParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MsonParserRULE_pair)

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
	{
		p.SetState(37)

		var _m = p.Match(MsonParserSTRING)

		localctx.(*PairContext).key = _m
	}
	{
		p.SetState(38)
		p.Match(MsonParserT__3)
	}
	{
		p.SetState(39)
		p.Value()
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
	p.RuleIndex = MsonParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsonParserRULE_array

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

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.ExitArray(s)
	}
}

func (p *MsonParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MsonParserRULE_array)
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

	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Match(MsonParserT__4)
		}
		{
			p.SetState(42)
			p.Value()
		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MsonParserT__1 {
			{
				p.SetState(43)
				p.Match(MsonParserT__1)
			}
			{
				p.SetState(44)
				p.Value()
			}

			p.SetState(49)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(50)
			p.Match(MsonParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.Match(MsonParserT__4)
		}
		{
			p.SetState(53)
			p.Match(MsonParserT__5)
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
	p.RuleIndex = MsonParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsonParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(MsonParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MsonParserNUMBER, 0)
}

func (s *ValueContext) BinData() IBinDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinDataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinDataContext)
}

func (s *ValueContext) NumberLong() INumberLongContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberLongContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberLongContext)
}

func (s *ValueContext) IsoDate() IIsoDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsoDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsoDateContext)
}

func (s *ValueContext) Obj() IObjContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *ValueContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *MsonParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MsonParserRULE_value)

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

	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MsonParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.Match(MsonParserSTRING)
		}

	case MsonParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Match(MsonParserNUMBER)
		}

	case MsonParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(58)
			p.BinData()
		}

	case MsonParserT__12:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(59)
			p.NumberLong()
		}

	case MsonParserT__13:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(60)
			p.IsoDate()
		}

	case MsonParserT__0:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(61)
			p.Obj()
		}

	case MsonParserT__4:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(62)
			p.Array()
		}

	case MsonParserT__6:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(63)
			p.Match(MsonParserT__6)
		}

	case MsonParserT__7:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(64)
			p.Match(MsonParserT__7)
		}

	case MsonParserT__8:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(65)
			p.Match(MsonParserT__8)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBinDataContext is an interface to support dynamic dispatch.
type IBinDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTyp returns the typ token.
	GetTyp() antlr.Token

	// GetData returns the data token.
	GetData() antlr.Token

	// SetTyp sets the typ token.
	SetTyp(antlr.Token)

	// SetData sets the data token.
	SetData(antlr.Token)

	// IsBinDataContext differentiates from other interfaces.
	IsBinDataContext()
}

type BinDataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	typ    antlr.Token
	data   antlr.Token
}

func NewEmptyBinDataContext() *BinDataContext {
	var p = new(BinDataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MsonParserRULE_binData
	return p
}

func (*BinDataContext) IsBinDataContext() {}

func NewBinDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinDataContext {
	var p = new(BinDataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsonParserRULE_binData

	return p
}

func (s *BinDataContext) GetParser() antlr.Parser { return s.parser }

func (s *BinDataContext) GetTyp() antlr.Token { return s.typ }

func (s *BinDataContext) GetData() antlr.Token { return s.data }

func (s *BinDataContext) SetTyp(v antlr.Token) { s.typ = v }

func (s *BinDataContext) SetData(v antlr.Token) { s.data = v }

func (s *BinDataContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MsonParserNUMBER, 0)
}

func (s *BinDataContext) STRING() antlr.TerminalNode {
	return s.GetToken(MsonParserSTRING, 0)
}

func (s *BinDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinDataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.EnterBinData(s)
	}
}

func (s *BinDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.ExitBinData(s)
	}
}

func (p *MsonParser) BinData() (localctx IBinDataContext) {
	localctx = NewBinDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MsonParserRULE_binData)

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
	{
		p.SetState(68)
		p.Match(MsonParserT__9)
	}
	{
		p.SetState(69)
		p.Match(MsonParserT__10)
	}
	{
		p.SetState(70)

		var _m = p.Match(MsonParserNUMBER)

		localctx.(*BinDataContext).typ = _m
	}
	{
		p.SetState(71)
		p.Match(MsonParserT__1)
	}
	{
		p.SetState(72)

		var _m = p.Match(MsonParserSTRING)

		localctx.(*BinDataContext).data = _m
	}
	{
		p.SetState(73)
		p.Match(MsonParserT__11)
	}

	return localctx
}

// INumberLongContext is an interface to support dynamic dispatch.
type INumberLongContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num token.
	GetNum() antlr.Token

	// GetStr returns the str token.
	GetStr() antlr.Token

	// SetNum sets the num token.
	SetNum(antlr.Token)

	// SetStr sets the str token.
	SetStr(antlr.Token)

	// IsNumberLongContext differentiates from other interfaces.
	IsNumberLongContext()
}

type NumberLongContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	num    antlr.Token
	str    antlr.Token
}

func NewEmptyNumberLongContext() *NumberLongContext {
	var p = new(NumberLongContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MsonParserRULE_numberLong
	return p
}

func (*NumberLongContext) IsNumberLongContext() {}

func NewNumberLongContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberLongContext {
	var p = new(NumberLongContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsonParserRULE_numberLong

	return p
}

func (s *NumberLongContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberLongContext) GetNum() antlr.Token { return s.num }

func (s *NumberLongContext) GetStr() antlr.Token { return s.str }

func (s *NumberLongContext) SetNum(v antlr.Token) { s.num = v }

func (s *NumberLongContext) SetStr(v antlr.Token) { s.str = v }

func (s *NumberLongContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MsonParserNUMBER, 0)
}

func (s *NumberLongContext) STRING() antlr.TerminalNode {
	return s.GetToken(MsonParserSTRING, 0)
}

func (s *NumberLongContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLongContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberLongContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.EnterNumberLong(s)
	}
}

func (s *NumberLongContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.ExitNumberLong(s)
	}
}

func (p *MsonParser) NumberLong() (localctx INumberLongContext) {
	localctx = NewNumberLongContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MsonParserRULE_numberLong)

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

	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Match(MsonParserT__12)
		}
		{
			p.SetState(76)
			p.Match(MsonParserT__10)
		}
		{
			p.SetState(77)

			var _m = p.Match(MsonParserNUMBER)

			localctx.(*NumberLongContext).num = _m
		}
		{
			p.SetState(78)
			p.Match(MsonParserT__11)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Match(MsonParserT__12)
		}
		{
			p.SetState(80)
			p.Match(MsonParserT__10)
		}
		{
			p.SetState(81)

			var _m = p.Match(MsonParserSTRING)

			localctx.(*NumberLongContext).str = _m
		}
		{
			p.SetState(82)
			p.Match(MsonParserT__11)
		}

	}

	return localctx
}

// IIsoDateContext is an interface to support dynamic dispatch.
type IIsoDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDate returns the date token.
	GetDate() antlr.Token

	// SetDate sets the date token.
	SetDate(antlr.Token)

	// IsIsoDateContext differentiates from other interfaces.
	IsIsoDateContext()
}

type IsoDateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	date   antlr.Token
}

func NewEmptyIsoDateContext() *IsoDateContext {
	var p = new(IsoDateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MsonParserRULE_isoDate
	return p
}

func (*IsoDateContext) IsIsoDateContext() {}

func NewIsoDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsoDateContext {
	var p = new(IsoDateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsonParserRULE_isoDate

	return p
}

func (s *IsoDateContext) GetParser() antlr.Parser { return s.parser }

func (s *IsoDateContext) GetDate() antlr.Token { return s.date }

func (s *IsoDateContext) SetDate(v antlr.Token) { s.date = v }

func (s *IsoDateContext) STRING() antlr.TerminalNode {
	return s.GetToken(MsonParserSTRING, 0)
}

func (s *IsoDateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsoDateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsoDateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.EnterIsoDate(s)
	}
}

func (s *IsoDateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsonListener); ok {
		listenerT.ExitIsoDate(s)
	}
}

func (p *MsonParser) IsoDate() (localctx IIsoDateContext) {
	localctx = NewIsoDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MsonParserRULE_isoDate)

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
	{
		p.SetState(85)
		p.Match(MsonParserT__13)
	}
	{
		p.SetState(86)
		p.Match(MsonParserT__10)
	}
	{
		p.SetState(87)

		var _m = p.Match(MsonParserSTRING)

		localctx.(*IsoDateContext).date = _m
	}
	{
		p.SetState(88)
		p.Match(MsonParserT__11)
	}

	return localctx
}
