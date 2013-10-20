package main

import (
	"flag"
	. "fmt"
	"github.com/yak-labs/chirp-lang"
	_ "github.com/yak-labs/chirp-lang/http"
	_ "github.com/yak-labs/chirp-lang/img"
	"io/ioutil"
	"os"
)

var cFlag = flag.String("c", "", "Immediate command to execute.")

func saveArgvStarting(fr *chirp.Frame, i int) {
	argv := []chirp.T{}
	for _, a := range os.Args[i:] {
		argv = append(argv, chirp.MkString(a))
	}
	fr.SetVar("argv", chirp.MkList(argv))
}

func main() {
	flag.Parse()
	fr := chirp.New()

	if cFlag != nil && *cFlag != "" {
		saveArgvStarting(fr, 1)
		fr.Eval(chirp.MkString(*cFlag))
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
		fr.Eval(chirp.MkString(string(contents)))
		return
	}

	// No os.Args --
	panic("REPL not yet")
}
