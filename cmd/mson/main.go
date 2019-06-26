package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/gobs/simplejson"
	"github.com/raff/mson"
)

func fatal(params ...interface{}) {
	fmt.Fprintln(os.Stderr, params...)
	os.Exit(1)
}

func fatalf(m string, params ...interface{}) {
	fmt.Fprintf(os.Stderr, m, params...)
	os.Exit(1)
}

func process(r io.Reader, fileName string, prologue bool) {
	reader := bufio.NewReader(r)

	// skip the "prologue" that mongdb shell adds when redirecting the output
	// (basically every up to the first '{')
	if prologue {
		for {
			next, err := reader.Peek(1)
			if err != nil {
				fatalf("%s: %v\n", fileName, err)
			}

			if next[0] == '{' {
				break
			}

			_, _, err = reader.ReadLine()
			if err != nil {
				break
			}
		}
	}

	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, reader)
	if err != nil {
		fatalf("%s: %v\n", fileName, err)
	}

	s := string(buf.Bytes())
	v, ok := mson.NewProcessor().Parse(s)

	if !ok {
		fatal("ERROR PARSING", fileName)
	} else {
		fmt.Println(simplejson.MustDumpString(v, simplejson.Indent(" ")))
	}
}

func main() {
	prologue := flag.Bool("prologue", false, "skip `prologue` added by MongoDB shell")

	flag.Parse()

	if flag.NArg() == 0 {
		process(os.Stdin, "<stdin>", *prologue)
		return
	}

	for _, fileName := range flag.Args() {
		f, err := os.Open(fileName)
		if err != nil {
			fatal(err)
		}

		process(f, fileName, *prologue)
		f.Close()
	}
}
