//go:generate antlr -Dlanguage=Go -package parser -o parser Mson.g4
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobs/simplejson"
	"github.com/raff/mson/parser"
)

type obj = map[string]interface{}
type arr = []interface{}

// Processor defines the JsonPath grammar processor
type Processor struct {
	*parser.BaseMsonListener
	value  []interface{}
	errors bool
}

// NewProcessor creates a new grammar Processor
func NewProcessor() *Processor {
	return &Processor{}
}

func (p *Processor) Errors() bool {
	return p.errors
}

func (p *Processor) Value() interface{} {
	return p.value
}

//
// Parse parses a Mson object
//
func (p *Processor) Parse(mson string) (result interface{}, ok bool) {
	p.errors = false
	p.value = nil

	input := antlr.NewInputStream(mson)
	lexer := parser.NewMsonLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewMsonParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	antlr.ParseTreeWalkerDefault.Walk(p, parser.Mson())

	if len(p.value) == 1 {
		result = p.value[0]
	} else {
		result = p.value
	}

	return result, p.errors == false
}

//
// VisitErrorNode is called when there is an error
//
func (p *Processor) VisitErrorNode(node antlr.ErrorNode) {
	p.errors = true
}

// ExitMson is called when exiting the mson production.
func (p *Processor) ExitMson(c *parser.MsonContext) {
	for _, v := range c.AllValue() {
		p.value = append(p.value, parseValue(v))
	}
}

func unquote(s string) string {
	n := len(s) - 1
	return s[1:n]
}

func parseNumber(s string) interface{} {
	if strings.ContainsAny(s, ".Ee") {
		f, _ := strconv.ParseFloat(s, 64)
		return f
	}

	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func parseInteger(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func parseUnsigned(s string) uint64 {
	i, _ := strconv.ParseUint(s, 10, 64)
	return i
}

func parsePair(c parser.IPairContext) (string, interface{}) {
	pc := c.(*parser.PairContext)
	key := unquote(pc.GetKey().GetText())
	value := parseValue(pc.Value())
	return key, value
}

func parseValue(c parser.IValueContext) interface{} {
	vc := c.(*parser.ValueContext)

	switch {
	case vc.Obj() != nil:
		co := vc.Obj().(*parser.ObjContext)
		oo := obj{}
		for _, pair := range co.AllPair() {
			k, v := parsePair(pair)
			oo[k] = v
		}
		return oo

	case vc.Array() != nil:
		ca := vc.Array().(*parser.ArrayContext)
		aa := arr{}
		for _, v := range ca.AllValue() {
			aa = append(aa, parseValue(v))
		}
		return aa

	case vc.STRING() != nil:
		return unquote(vc.STRING().GetText())
	case vc.NUMBER() != nil:
		return parseNumber(vc.NUMBER().GetText())
	case vc.BinData() != nil:
		return obj{
			"bindata": obj{
				"type": parseUnsigned(vc.BinData().GetTyp().GetText()),
				"data": unquote(vc.BinData().GetData().GetText()),
			},
		}
	case vc.IsoDate() != nil:
		return obj{
			"date": unquote(vc.IsoDate().GetDate().GetText()),
		}
	case vc.NumberLong() != nil:
		if vc.NumberLong().GetNum() != nil {
			return parseInteger(vc.NumberLong().GetNum().GetText())
		} else {
			return parseNumber(unquote(vc.NumberLong().GetStr().GetText()))
		}
	case vc.GetText() == "null":
		return nil
	case vc.GetText() == "true":
		return true
	case vc.GetText() == "false":
		return false
	}

	panic("unexpected value for " + vc.GetText())
	return nil
}

func main() {
	fileName := os.Args[1]
	buf := bytes.NewBuffer(nil)

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(f)

	// skip the "prologue" that mongdb shell adds when redirecting the output
	// (basically every up to the first '{')
	for {
		next, err := reader.Peek(1)
		if err != nil {
			panic(err)
		}

		if next[0] == '{' {
			break
		}

		_, _, err = reader.ReadLine()
		if err != nil {
			break
		}
	}

	_, err = io.Copy(buf, reader)
	if err != nil {
		fmt.Println(err)
	}

	f.Close()

	s := string(buf.Bytes())
	v, ok := NewProcessor().Parse(s)

	if !ok {
		fmt.Println("ERROR PARSING", fileName)
	} else {
		fmt.Println(simplejson.MustDumpString(v, simplejson.Indent(" ")))
	}
}
