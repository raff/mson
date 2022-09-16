// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Mson
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MsonParser struct {
	*antlr.BaseParser
}

var msonParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func msonParserInit() {
	staticData := &msonParserStaticData
	staticData.literalNames = []string{
		"", "'{'", "','", "'}'", "':'", "'['", "']'", "'true'", "'false'", "'null'",
		"'BinData'", "'('", "')'", "'NumberLong'", "'ISODate'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "STRING",
		"NUMBER", "WS",
	}
	staticData.ruleNames = []string{
		"mson", "obj", "pair", "array", "value", "binData", "numberLong", "isoDate",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 17, 91, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 5, 0, 18, 8, 0, 10, 0, 12,
		0, 21, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 27, 8, 1, 10, 1, 12, 1, 30,
		9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 36, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 5, 3, 46, 8, 3, 10, 3, 12, 3, 49, 9, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 55, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 3, 4, 67, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 84, 8, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 0, 97,
		0, 19, 1, 0, 0, 0, 2, 35, 1, 0, 0, 0, 4, 37, 1, 0, 0, 0, 6, 54, 1, 0, 0,
		0, 8, 66, 1, 0, 0, 0, 10, 68, 1, 0, 0, 0, 12, 83, 1, 0, 0, 0, 14, 85, 1,
		0, 0, 0, 16, 18, 3, 8, 4, 0, 17, 16, 1, 0, 0, 0, 18, 21, 1, 0, 0, 0, 19,
		17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 1, 1, 0, 0, 0, 21, 19, 1, 0, 0,
		0, 22, 23, 5, 1, 0, 0, 23, 28, 3, 4, 2, 0, 24, 25, 5, 2, 0, 0, 25, 27,
		3, 4, 2, 0, 26, 24, 1, 0, 0, 0, 27, 30, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0,
		28, 29, 1, 0, 0, 0, 29, 31, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 31, 32, 5,
		3, 0, 0, 32, 36, 1, 0, 0, 0, 33, 34, 5, 1, 0, 0, 34, 36, 5, 3, 0, 0, 35,
		22, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 36, 3, 1, 0, 0, 0, 37, 38, 5, 15, 0,
		0, 38, 39, 5, 4, 0, 0, 39, 40, 3, 8, 4, 0, 40, 5, 1, 0, 0, 0, 41, 42, 5,
		5, 0, 0, 42, 47, 3, 8, 4, 0, 43, 44, 5, 2, 0, 0, 44, 46, 3, 8, 4, 0, 45,
		43, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0,
		0, 48, 50, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 50, 51, 5, 6, 0, 0, 51, 55,
		1, 0, 0, 0, 52, 53, 5, 5, 0, 0, 53, 55, 5, 6, 0, 0, 54, 41, 1, 0, 0, 0,
		54, 52, 1, 0, 0, 0, 55, 7, 1, 0, 0, 0, 56, 67, 5, 15, 0, 0, 57, 67, 5,
		16, 0, 0, 58, 67, 3, 10, 5, 0, 59, 67, 3, 12, 6, 0, 60, 67, 3, 14, 7, 0,
		61, 67, 3, 2, 1, 0, 62, 67, 3, 6, 3, 0, 63, 67, 5, 7, 0, 0, 64, 67, 5,
		8, 0, 0, 65, 67, 5, 9, 0, 0, 66, 56, 1, 0, 0, 0, 66, 57, 1, 0, 0, 0, 66,
		58, 1, 0, 0, 0, 66, 59, 1, 0, 0, 0, 66, 60, 1, 0, 0, 0, 66, 61, 1, 0, 0,
		0, 66, 62, 1, 0, 0, 0, 66, 63, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 65,
		1, 0, 0, 0, 67, 9, 1, 0, 0, 0, 68, 69, 5, 10, 0, 0, 69, 70, 5, 11, 0, 0,
		70, 71, 5, 16, 0, 0, 71, 72, 5, 2, 0, 0, 72, 73, 5, 15, 0, 0, 73, 74, 5,
		12, 0, 0, 74, 11, 1, 0, 0, 0, 75, 76, 5, 13, 0, 0, 76, 77, 5, 11, 0, 0,
		77, 78, 5, 16, 0, 0, 78, 84, 5, 12, 0, 0, 79, 80, 5, 13, 0, 0, 80, 81,
		5, 11, 0, 0, 81, 82, 5, 15, 0, 0, 82, 84, 5, 12, 0, 0, 83, 75, 1, 0, 0,
		0, 83, 79, 1, 0, 0, 0, 84, 13, 1, 0, 0, 0, 85, 86, 5, 14, 0, 0, 86, 87,
		5, 11, 0, 0, 87, 88, 5, 15, 0, 0, 88, 89, 5, 12, 0, 0, 89, 15, 1, 0, 0,
		0, 7, 19, 28, 35, 47, 54, 66, 83,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MsonParserInit initializes any static state used to implement MsonParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMsonParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MsonParserInit() {
	staticData := &msonParserStaticData
	staticData.once.Do(msonParserInit)
}

// NewMsonParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMsonParser(input antlr.TokenStream) *MsonParser {
	MsonParserInit()
	this := new(MsonParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &msonParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *MsonContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	this := p
	_ = this

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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&124834) != 0 {
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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPairContext); ok {
			len++
		}
	}

	tst := make([]IPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPairContext); ok {
			tst[i] = t.(IPairContext)
			i++
		}
	}

	return tst
}

func (s *ObjContext) Pair(i int) IPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	this := p
	_ = this

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	this := p
	_ = this

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinDataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinDataContext)
}

func (s *ValueContext) NumberLong() INumberLongContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberLongContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberLongContext)
}

func (s *ValueContext) IsoDate() IIsoDateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIsoDateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIsoDateContext)
}

func (s *ValueContext) Obj() IObjContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *ValueContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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
	this := p
	_ = this

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
	this := p
	_ = this

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
	this := p
	_ = this

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
