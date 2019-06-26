// Code generated from Mson.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Mson
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMsonListener is a complete listener for a parse tree produced by MsonParser.
type BaseMsonListener struct{}

var _ MsonListener = &BaseMsonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMsonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMsonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMsonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMsonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMson is called when production mson is entered.
func (s *BaseMsonListener) EnterMson(ctx *MsonContext) {}

// ExitMson is called when production mson is exited.
func (s *BaseMsonListener) ExitMson(ctx *MsonContext) {}

// EnterObj is called when production obj is entered.
func (s *BaseMsonListener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BaseMsonListener) ExitObj(ctx *ObjContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseMsonListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseMsonListener) ExitPair(ctx *PairContext) {}

// EnterArray is called when production array is entered.
func (s *BaseMsonListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseMsonListener) ExitArray(ctx *ArrayContext) {}

// EnterValue is called when production value is entered.
func (s *BaseMsonListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseMsonListener) ExitValue(ctx *ValueContext) {}

// EnterBinData is called when production binData is entered.
func (s *BaseMsonListener) EnterBinData(ctx *BinDataContext) {}

// ExitBinData is called when production binData is exited.
func (s *BaseMsonListener) ExitBinData(ctx *BinDataContext) {}

// EnterNumberLong is called when production numberLong is entered.
func (s *BaseMsonListener) EnterNumberLong(ctx *NumberLongContext) {}

// ExitNumberLong is called when production numberLong is exited.
func (s *BaseMsonListener) ExitNumberLong(ctx *NumberLongContext) {}

// EnterIsoDate is called when production isoDate is entered.
func (s *BaseMsonListener) EnterIsoDate(ctx *IsoDateContext) {}

// ExitIsoDate is called when production isoDate is exited.
func (s *BaseMsonListener) ExitIsoDate(ctx *IsoDateContext) {}
