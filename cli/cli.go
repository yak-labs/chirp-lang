package cli

/*
	For debugging exec pipes, try this:
	go run chirp.go -recover=0 -c='puts [exec ls -l | sed {s/[0-9]/#/g} | tr {a-z} {A-Z} ]' 2>/dev/null | od -c
*/

import (
	"bufio"
	"flag"
	. "fmt"
	"github.com/yak-labs/chirp-lang"
	"io/ioutil"
	"os"
	"runtime/pprof"
	"strings"
)

var dFlag = flag.String("d", "", "Debugging flags, each a single letter.")
var cFlag = flag.String("c", "", "Immediate command to execute.")
var recoverFlag = flag.Bool("recover", true, "Set to false to disable recover in the REPL.")

func saveArgvStarting(fr *chirp.Frame, i int) {
	argv := []chirp.T{}
	for _, a := range flag.Args() {
		argv = append(argv, chirp.MkString(a))
	}
	fr.SetVar("argv", chirp.MkList(argv))
}

func setEnvironInChirp(fr *chirp.Frame, varName string) {
	h := make(chirp.Hash)
	for _, s := range os.Environ() {
		kv := strings.SplitN(s, "=", 2)
		if len(kv) == 2 {
			h[kv[0]] = chirp.MkString(kv[1])
		}
	}
	fr.SetVar(varName, chirp.MkHash(h))
}

func Main() {
	flag.Parse()
	fr := chirp.NewInterpreter()
	setEnvironInChirp(fr, "Env")

	for _, ch := range *dFlag {
		chirp.Debug[ch] = true
	}

	if cFlag != nil && *cFlag != "" {
		saveArgvStarting(fr, 1)
		fr.Eval(chirp.MkString(*cFlag))
		goto End
	}

	if len(flag.Args()) > 0 {
		fname := flag.Arg(0)
		contents, err := ioutil.ReadFile(fname)
		if err != nil {
			Fprintf(os.Stderr, "Cannot read file %s: %v", fname, err)
			os.Exit(2)
			return
		}
		saveArgvStarting(fr, 1)
		fr.Eval(chirp.MkString(string(contents)))
		goto End
	}

	{
		bio := bufio.NewReader(os.Stdin)
		for {
			Fprint(os.Stderr, "chirp% ") // Prompt to stderr.
			line, isPrefix, err := bio.ReadLine()
			if err != nil {
				if err.Error() == "EOF" { // TODO: better way?
					goto End
				}
				Fprintf(os.Stderr, "ERROR in ReadLine: %s\n", err.Error())
				goto End
			}
			fullLine := line
			for isPrefix {
				line, isPrefix, err = bio.ReadLine()
				if err != nil {
					Fprintf(os.Stderr, "ERROR in ReadLine: %s\n", err.Error())
					goto End
				}
				fullLine = append(fullLine, line...)
			}
			result := EvalStringOrPrintError(fr, string(fullLine))
			if result != "" { // Traditionally, if result is empty, tclsh doesn't print.
				Println(result)
			}
		}
	}

End:
	logAllCounters()
	if chirp.Debug['h'] {
		pprof.Lookup("heap").WriteTo(os.Stderr, 0)
	}
}

func logAllCounters() {
	if chirp.Debug['c'] {
		chirp.LogAllCounters()
	}
}

func EvalStringOrPrintError(fr *chirp.Frame, cmd string) (out string) {
	if *recoverFlag {
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
