// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MsonLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var msonlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func msonlexerLexerInit() {
	staticData := &msonlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'{'", "','", "'}'", "':'", "'['", "']'", "'true'", "'false'", "'null'",
		"'BinData'", "'('", "')'", "'NumberLong'", "'ISODate'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "STRING",
		"NUMBER", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "STRING", "ESC", "UNICODE",
		"HEX", "SAFECODEPOINT", "NUMBER", "INT", "EXP", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 169, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 5,
		14, 110, 8, 14, 10, 14, 12, 14, 113, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 3, 15, 120, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 19, 3, 19, 133, 8, 19, 1, 19, 1, 19, 1, 19,
		4, 19, 138, 8, 19, 11, 19, 12, 19, 139, 3, 19, 142, 8, 19, 1, 19, 3, 19,
		145, 8, 19, 1, 20, 1, 20, 1, 20, 5, 20, 150, 8, 20, 10, 20, 12, 20, 153,
		9, 20, 3, 20, 155, 8, 20, 1, 21, 1, 21, 3, 21, 159, 8, 21, 1, 21, 1, 21,
		1, 22, 4, 22, 164, 8, 22, 11, 22, 12, 22, 165, 1, 22, 1, 22, 0, 0, 23,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 0, 33, 0, 35, 0, 37, 0, 39, 16, 41,
		0, 43, 0, 45, 17, 1, 0, 8, 8, 0, 34, 34, 47, 47, 92, 92, 98, 98, 102, 102,
		110, 110, 114, 114, 116, 116, 3, 0, 48, 57, 65, 70, 97, 102, 3, 0, 0, 31,
		34, 34, 92, 92, 1, 0, 48, 57, 1, 0, 49, 57, 2, 0, 69, 69, 101, 101, 2,
		0, 43, 43, 45, 45, 3, 0, 9, 10, 13, 13, 32, 32, 173, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 1, 47, 1, 0, 0, 0, 3, 49, 1, 0, 0, 0, 5, 51, 1, 0, 0, 0, 7, 53,
		1, 0, 0, 0, 9, 55, 1, 0, 0, 0, 11, 57, 1, 0, 0, 0, 13, 59, 1, 0, 0, 0,
		15, 64, 1, 0, 0, 0, 17, 70, 1, 0, 0, 0, 19, 75, 1, 0, 0, 0, 21, 83, 1,
		0, 0, 0, 23, 85, 1, 0, 0, 0, 25, 87, 1, 0, 0, 0, 27, 98, 1, 0, 0, 0, 29,
		106, 1, 0, 0, 0, 31, 116, 1, 0, 0, 0, 33, 121, 1, 0, 0, 0, 35, 127, 1,
		0, 0, 0, 37, 129, 1, 0, 0, 0, 39, 132, 1, 0, 0, 0, 41, 154, 1, 0, 0, 0,
		43, 156, 1, 0, 0, 0, 45, 163, 1, 0, 0, 0, 47, 48, 5, 123, 0, 0, 48, 2,
		1, 0, 0, 0, 49, 50, 5, 44, 0, 0, 50, 4, 1, 0, 0, 0, 51, 52, 5, 125, 0,
		0, 52, 6, 1, 0, 0, 0, 53, 54, 5, 58, 0, 0, 54, 8, 1, 0, 0, 0, 55, 56, 5,
		91, 0, 0, 56, 10, 1, 0, 0, 0, 57, 58, 5, 93, 0, 0, 58, 12, 1, 0, 0, 0,
		59, 60, 5, 116, 0, 0, 60, 61, 5, 114, 0, 0, 61, 62, 5, 117, 0, 0, 62, 63,
		5, 101, 0, 0, 63, 14, 1, 0, 0, 0, 64, 65, 5, 102, 0, 0, 65, 66, 5, 97,
		0, 0, 66, 67, 5, 108, 0, 0, 67, 68, 5, 115, 0, 0, 68, 69, 5, 101, 0, 0,
		69, 16, 1, 0, 0, 0, 70, 71, 5, 110, 0, 0, 71, 72, 5, 117, 0, 0, 72, 73,
		5, 108, 0, 0, 73, 74, 5, 108, 0, 0, 74, 18, 1, 0, 0, 0, 75, 76, 5, 66,
		0, 0, 76, 77, 5, 105, 0, 0, 77, 78, 5, 110, 0, 0, 78, 79, 5, 68, 0, 0,
		79, 80, 5, 97, 0, 0, 80, 81, 5, 116, 0, 0, 81, 82, 5, 97, 0, 0, 82, 20,
		1, 0, 0, 0, 83, 84, 5, 40, 0, 0, 84, 22, 1, 0, 0, 0, 85, 86, 5, 41, 0,
		0, 86, 24, 1, 0, 0, 0, 87, 88, 5, 78, 0, 0, 88, 89, 5, 117, 0, 0, 89, 90,
		5, 109, 0, 0, 90, 91, 5, 98, 0, 0, 91, 92, 5, 101, 0, 0, 92, 93, 5, 114,
		0, 0, 93, 94, 5, 76, 0, 0, 94, 95, 5, 111, 0, 0, 95, 96, 5, 110, 0, 0,
		96, 97, 5, 103, 0, 0, 97, 26, 1, 0, 0, 0, 98, 99, 5, 73, 0, 0, 99, 100,
		5, 83, 0, 0, 100, 101, 5, 79, 0, 0, 101, 102, 5, 68, 0, 0, 102, 103, 5,
		97, 0, 0, 103, 104, 5, 116, 0, 0, 104, 105, 5, 101, 0, 0, 105, 28, 1, 0,
		0, 0, 106, 111, 5, 34, 0, 0, 107, 110, 3, 31, 15, 0, 108, 110, 3, 37, 18,
		0, 109, 107, 1, 0, 0, 0, 109, 108, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111,
		109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0, 0, 0, 113, 111,
		1, 0, 0, 0, 114, 115, 5, 34, 0, 0, 115, 30, 1, 0, 0, 0, 116, 119, 5, 92,
		0, 0, 117, 120, 7, 0, 0, 0, 118, 120, 3, 33, 16, 0, 119, 117, 1, 0, 0,
		0, 119, 118, 1, 0, 0, 0, 120, 32, 1, 0, 0, 0, 121, 122, 5, 117, 0, 0, 122,
		123, 3, 35, 17, 0, 123, 124, 3, 35, 17, 0, 124, 125, 3, 35, 17, 0, 125,
		126, 3, 35, 17, 0, 126, 34, 1, 0, 0, 0, 127, 128, 7, 1, 0, 0, 128, 36,
		1, 0, 0, 0, 129, 130, 8, 2, 0, 0, 130, 38, 1, 0, 0, 0, 131, 133, 5, 45,
		0, 0, 132, 131, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0,
		134, 141, 3, 41, 20, 0, 135, 137, 5, 46, 0, 0, 136, 138, 7, 3, 0, 0, 137,
		136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140,
		1, 0, 0, 0, 140, 142, 1, 0, 0, 0, 141, 135, 1, 0, 0, 0, 141, 142, 1, 0,
		0, 0, 142, 144, 1, 0, 0, 0, 143, 145, 3, 43, 21, 0, 144, 143, 1, 0, 0,
		0, 144, 145, 1, 0, 0, 0, 145, 40, 1, 0, 0, 0, 146, 155, 5, 48, 0, 0, 147,
		151, 7, 4, 0, 0, 148, 150, 7, 3, 0, 0, 149, 148, 1, 0, 0, 0, 150, 153,
		1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 155, 1, 0,
		0, 0, 153, 151, 1, 0, 0, 0, 154, 146, 1, 0, 0, 0, 154, 147, 1, 0, 0, 0,
		155, 42, 1, 0, 0, 0, 156, 158, 7, 5, 0, 0, 157, 159, 7, 6, 0, 0, 158, 157,
		1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 161, 3, 41,
		20, 0, 161, 44, 1, 0, 0, 0, 162, 164, 7, 7, 0, 0, 163, 162, 1, 0, 0, 0,
		164, 165, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166,
		167, 1, 0, 0, 0, 167, 168, 6, 22, 0, 0, 168, 46, 1, 0, 0, 0, 12, 0, 109,
		111, 119, 132, 139, 141, 144, 151, 154, 158, 165, 1, 6, 0, 0,
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

// MsonLexerInit initializes any static state used to implement MsonLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMsonLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MsonLexerInit() {
	staticData := &msonlexerLexerStaticData
	staticData.once.Do(msonlexerLexerInit)
}

// NewMsonLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMsonLexer(input antlr.CharStream) *MsonLexer {
	MsonLexerInit()
	l := new(MsonLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &msonlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Mson.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MsonLexer tokens.
const (
	MsonLexerT__0   = 1
	MsonLexerT__1   = 2
	MsonLexerT__2   = 3
	MsonLexerT__3   = 4
	MsonLexerT__4   = 5
	MsonLexerT__5   = 6
	MsonLexerT__6   = 7
	MsonLexerT__7   = 8
	MsonLexerT__8   = 9
	MsonLexerT__9   = 10
	MsonLexerT__10  = 11
	MsonLexerT__11  = 12
	MsonLexerT__12  = 13
	MsonLexerT__13  = 14
	MsonLexerSTRING = 15
	MsonLexerNUMBER = 16
	MsonLexerWS     = 17
)
