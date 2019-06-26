// Code generated from Mson.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Mson
import "github.com/antlr/antlr4/runtime/Go/antlr"

// MsonListener is a complete listener for a parse tree produced by MsonParser.
type MsonListener interface {
	antlr.ParseTreeListener

	// EnterMson is called when entering the mson production.
	EnterMson(c *MsonContext)

	// EnterObj is called when entering the obj production.
	EnterObj(c *ObjContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterBinData is called when entering the binData production.
	EnterBinData(c *BinDataContext)

	// EnterNumberLong is called when entering the numberLong production.
	EnterNumberLong(c *NumberLongContext)

	// EnterIsoDate is called when entering the isoDate production.
	EnterIsoDate(c *IsoDateContext)

	// ExitMson is called when exiting the mson production.
	ExitMson(c *MsonContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitBinData is called when exiting the binData production.
	ExitBinData(c *BinDataContext)

	// ExitNumberLong is called when exiting the numberLong production.
	ExitNumberLong(c *NumberLongContext)

	// ExitIsoDate is called when exiting the isoDate production.
	ExitIsoDate(c *IsoDateContext)
}
