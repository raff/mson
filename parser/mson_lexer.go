// Code generated from Mson.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 171,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 7,
	16, 112, 10, 16, 12, 16, 14, 16, 115, 11, 16, 3, 16, 3, 16, 3, 17, 3, 17,
	3, 17, 5, 17, 122, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 21, 5, 21, 135, 10, 21, 3, 21, 3, 21, 3, 21,
	6, 21, 140, 10, 21, 13, 21, 14, 21, 141, 5, 21, 144, 10, 21, 3, 21, 5,
	21, 147, 10, 21, 3, 22, 3, 22, 3, 22, 7, 22, 152, 10, 22, 12, 22, 14, 22,
	155, 11, 22, 5, 22, 157, 10, 22, 3, 23, 3, 23, 5, 23, 161, 10, 23, 3, 23,
	3, 23, 3, 24, 6, 24, 166, 10, 24, 13, 24, 14, 24, 167, 3, 24, 3, 24, 2,
	2, 25, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 2, 35, 2, 37, 2, 39, 2,
	41, 18, 43, 2, 45, 2, 47, 19, 3, 2, 10, 10, 2, 36, 36, 49, 49, 94, 94,
	100, 100, 104, 104, 112, 112, 116, 116, 118, 118, 5, 2, 50, 59, 67, 72,
	99, 104, 5, 2, 2, 33, 36, 36, 94, 94, 3, 2, 50, 59, 3, 2, 51, 59, 4, 2,
	71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 5, 2, 11, 12, 15, 15, 34, 34, 2,
	175, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2,
	41, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 3, 49, 3, 2, 2, 2, 5, 51, 3, 2, 2, 2,
	7, 53, 3, 2, 2, 2, 9, 55, 3, 2, 2, 2, 11, 57, 3, 2, 2, 2, 13, 59, 3, 2,
	2, 2, 15, 61, 3, 2, 2, 2, 17, 66, 3, 2, 2, 2, 19, 72, 3, 2, 2, 2, 21, 77,
	3, 2, 2, 2, 23, 85, 3, 2, 2, 2, 25, 87, 3, 2, 2, 2, 27, 89, 3, 2, 2, 2,
	29, 100, 3, 2, 2, 2, 31, 108, 3, 2, 2, 2, 33, 118, 3, 2, 2, 2, 35, 123,
	3, 2, 2, 2, 37, 129, 3, 2, 2, 2, 39, 131, 3, 2, 2, 2, 41, 134, 3, 2, 2,
	2, 43, 156, 3, 2, 2, 2, 45, 158, 3, 2, 2, 2, 47, 165, 3, 2, 2, 2, 49, 50,
	7, 125, 2, 2, 50, 4, 3, 2, 2, 2, 51, 52, 7, 46, 2, 2, 52, 6, 3, 2, 2, 2,
	53, 54, 7, 127, 2, 2, 54, 8, 3, 2, 2, 2, 55, 56, 7, 60, 2, 2, 56, 10, 3,
	2, 2, 2, 57, 58, 7, 93, 2, 2, 58, 12, 3, 2, 2, 2, 59, 60, 7, 95, 2, 2,
	60, 14, 3, 2, 2, 2, 61, 62, 7, 118, 2, 2, 62, 63, 7, 116, 2, 2, 63, 64,
	7, 119, 2, 2, 64, 65, 7, 103, 2, 2, 65, 16, 3, 2, 2, 2, 66, 67, 7, 104,
	2, 2, 67, 68, 7, 99, 2, 2, 68, 69, 7, 110, 2, 2, 69, 70, 7, 117, 2, 2,
	70, 71, 7, 103, 2, 2, 71, 18, 3, 2, 2, 2, 72, 73, 7, 112, 2, 2, 73, 74,
	7, 119, 2, 2, 74, 75, 7, 110, 2, 2, 75, 76, 7, 110, 2, 2, 76, 20, 3, 2,
	2, 2, 77, 78, 7, 68, 2, 2, 78, 79, 7, 107, 2, 2, 79, 80, 7, 112, 2, 2,
	80, 81, 7, 70, 2, 2, 81, 82, 7, 99, 2, 2, 82, 83, 7, 118, 2, 2, 83, 84,
	7, 99, 2, 2, 84, 22, 3, 2, 2, 2, 85, 86, 7, 42, 2, 2, 86, 24, 3, 2, 2,
	2, 87, 88, 7, 43, 2, 2, 88, 26, 3, 2, 2, 2, 89, 90, 7, 80, 2, 2, 90, 91,
	7, 119, 2, 2, 91, 92, 7, 111, 2, 2, 92, 93, 7, 100, 2, 2, 93, 94, 7, 103,
	2, 2, 94, 95, 7, 116, 2, 2, 95, 96, 7, 78, 2, 2, 96, 97, 7, 113, 2, 2,
	97, 98, 7, 112, 2, 2, 98, 99, 7, 105, 2, 2, 99, 28, 3, 2, 2, 2, 100, 101,
	7, 75, 2, 2, 101, 102, 7, 85, 2, 2, 102, 103, 7, 81, 2, 2, 103, 104, 7,
	70, 2, 2, 104, 105, 7, 99, 2, 2, 105, 106, 7, 118, 2, 2, 106, 107, 7, 103,
	2, 2, 107, 30, 3, 2, 2, 2, 108, 113, 7, 36, 2, 2, 109, 112, 5, 33, 17,
	2, 110, 112, 5, 39, 20, 2, 111, 109, 3, 2, 2, 2, 111, 110, 3, 2, 2, 2,
	112, 115, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114,
	116, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 116, 117, 7, 36, 2, 2, 117, 32,
	3, 2, 2, 2, 118, 121, 7, 94, 2, 2, 119, 122, 9, 2, 2, 2, 120, 122, 5, 35,
	18, 2, 121, 119, 3, 2, 2, 2, 121, 120, 3, 2, 2, 2, 122, 34, 3, 2, 2, 2,
	123, 124, 7, 119, 2, 2, 124, 125, 5, 37, 19, 2, 125, 126, 5, 37, 19, 2,
	126, 127, 5, 37, 19, 2, 127, 128, 5, 37, 19, 2, 128, 36, 3, 2, 2, 2, 129,
	130, 9, 3, 2, 2, 130, 38, 3, 2, 2, 2, 131, 132, 10, 4, 2, 2, 132, 40, 3,
	2, 2, 2, 133, 135, 7, 47, 2, 2, 134, 133, 3, 2, 2, 2, 134, 135, 3, 2, 2,
	2, 135, 136, 3, 2, 2, 2, 136, 143, 5, 43, 22, 2, 137, 139, 7, 48, 2, 2,
	138, 140, 9, 5, 2, 2, 139, 138, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141,
	139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 144, 3, 2, 2, 2, 143, 137,
	3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 146, 3, 2, 2, 2, 145, 147, 5, 45,
	23, 2, 146, 145, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 42, 3, 2, 2, 2,
	148, 157, 7, 50, 2, 2, 149, 153, 9, 6, 2, 2, 150, 152, 9, 5, 2, 2, 151,
	150, 3, 2, 2, 2, 152, 155, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 153, 154,
	3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 156, 148, 3, 2,
	2, 2, 156, 149, 3, 2, 2, 2, 157, 44, 3, 2, 2, 2, 158, 160, 9, 7, 2, 2,
	159, 161, 9, 8, 2, 2, 160, 159, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161,
	162, 3, 2, 2, 2, 162, 163, 5, 43, 22, 2, 163, 46, 3, 2, 2, 2, 164, 166,
	9, 9, 2, 2, 165, 164, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 165, 3, 2,
	2, 2, 167, 168, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 170, 8, 24, 2, 2,
	170, 48, 3, 2, 2, 2, 14, 2, 111, 113, 121, 134, 141, 143, 146, 153, 156,
	160, 167, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'{'", "','", "'}'", "':'", "'['", "']'", "'true'", "'false'", "'null'",
	"'BinData'", "'('", "')'", "'NumberLong'", "'ISODate'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "STRING", "NUMBER",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "STRING", "ESC", "UNICODE",
	"HEX", "SAFECODEPOINT", "NUMBER", "INT", "EXP", "WS",
}

type MsonLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewMsonLexer(input antlr.CharStream) *MsonLexer {

	l := new(MsonLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
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
