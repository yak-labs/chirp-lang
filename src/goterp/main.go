package main

import (
	"flag"
	. "fmt"
	"io/ioutil"
	"os"

	"terp"
)

var cFlag = flag.String("c", "", "Immediate command to execute.")

func saveArgvStarting(fr *terp.Frame, i int) {
	argv := []terp.T{}
	for _, a := range os.Args[i:] {
		argv = append(argv, terp.MkString(a))
	}
	fr.SetVar("argv", terp.MkList(argv))
}

func main() {
	flag.Parse()
	fr := terp.New()

	if cFlag != nil && *cFlag != "" {
		saveArgvStarting(fr, 1)
		Printf("T<<< %#v\n", *cFlag)
		z := fr.Eval(terp.MkString(*cFlag))
		Printf("T>>> %#v\n", z)
		return
	}

	if len(os.Args) > 1 {
		fname := os.Args[1]
		contents, err := ioutil.ReadFile(fname)
		if err != nil {
			Fprintf(os.Stderr, "Cannot read file %s: %v", fname, err)
			os.Exit(2)
			return
		}
		saveArgvStarting(fr, 2)
		Printf("T<<< fname = %#v\n", fname)
		z := fr.Eval(terp.MkString(string(contents)))
		Printf("T>>> %#v\n", z)
		return
	}

	// No os.Args --
	panic("REPL not yet")
}
