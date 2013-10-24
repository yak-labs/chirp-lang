package main

/*
	For debugging exec pipes, try this:
	go run chirp.go -panic -c='puts [exec ls -l | sed {s/[0-9]/#/g} | tr {a-z} {A-Z} ]' 2>/dev/null | od -c
*/

import (
	_ "github.com/yak-labs/chirp-lang/http"
	_ "github.com/yak-labs/chirp-lang/img"
	_ "github.com/yak-labs/chirp-lang/posix"
)

import (
	"bufio"
	"flag"
	. "fmt"
	"github.com/yak-labs/chirp-lang"
	"io/ioutil"
	"os"
)

var cFlag = flag.String("c", "", "Immediate command to execute.")
var panicFlag = flag.Bool("panic", false, "Don't catch panic in REPL.")

func saveArgvStarting(fr *chirp.Frame, i int) {
	argv := []chirp.T{}
	for _, a := range flag.Args() {
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

	if len(flag.Args()) > 1 {
		fname := flag.Arg(1)
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

	bio := bufio.NewReader(os.Stdin)
	for {
		Fprint(os.Stderr, "chirp% ") // Prompt to stderr.
		line, isPrefix, err := bio.ReadLine()
		if err != nil {
			if err.Error() == "EOF" { // TODO: better way?
				return
			}
			Fprintf(os.Stderr, "ERROR in ReadLine: %s\n", err.Error())
			return
		}
		fullLine := line
		for isPrefix {
			line, isPrefix, err = bio.ReadLine()
			if err != nil {
				Fprintf(os.Stderr, "ERROR in ReadLine: %s\n", err.Error())
				return
			}
			fullLine = append(fullLine, line...)
		}
		result := EvalStringOrPrintError(fr, string(fullLine))
		if result != "" { // Traditionally, if result is empty, tclsh doesn't print.
			Println(result)
		}
	}
}

func EvalStringOrPrintError(fr *chirp.Frame, cmd string) (out string) {
	if panicFlag != nil {
		defer func() {
			if r := recover(); r != nil {
				Fprintln(os.Stderr, "ERROR: ", r) // Error to stderr.
				out = ""
				return
			}
		}()
	}

	return fr.Eval(chirp.MkString(cmd)).String()
}
