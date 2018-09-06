package cli

/*
	For debugging exec pipes, try this:
	go run chirp.go -recover=0 -c='puts [exec ls -l | sed {s/[0-9]/#/g} | tr {a-z} {A-Z} ]' 2>/dev/null | od -c
*/

import (
	chirp "github.com/yak-labs/chirp-lang"

	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime/pprof"
	"strings"

	"github.com/chzyer/readline"
)

var dFlag = flag.String("d", "", "Debugging flags, each a single letter.")
var cFlag = flag.String("c", "", "Immediate command to execute.")
var recoverFlag = flag.Bool("recover", true, "Set to false to disable recover in the REPL.")
var testFlag = flag.Bool("test", false, "Print test summary at end.")

var scriptName string

func saveArgvStarting(fr *chirp.Frame, i int) {
	argv := []chirp.T{}
	for _, a := range flag.Args() {
		argv = append(argv, chirp.MkString(a))
	}
	fr.SetVar("argv", chirp.MkList(argv)) // Deprecated: argv
	fr.SetVar("Argv", chirp.MkList(argv)) // New: Argv
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
		if ch < 256 {
			chirp.Debug[ch] = true
		}
	}

	if cFlag != nil && *cFlag != "" {
		saveArgvStarting(fr, 1)
		fr.Eval(chirp.MkString(*cFlag))
		goto End
	}

	if len(flag.Args()) > 0 {
		// Script mode.
		scriptName = flag.Arg(0)
		contents, err := ioutil.ReadFile(scriptName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot read file %s: %v", scriptName, err)
			os.Exit(2)
			return
		}
		saveArgvStarting(fr, 1)

		fr.Eval(chirp.MkString(string(contents)))
		goto End
	}

	{
		// Interactive mode.
		home := os.Getenv("HOME")
		if home == "" {
			home = "."
		}

		rl, err := readline.NewEx(&readline.Config{
			Prompt:          "% ",
			HistoryFile:     filepath.Join(home, ".chirp.history"),
			InterruptPrompt: "*SIGINT*",
			EOFPrompt:       "*EOF*",
			// AutoComplete:    completer,
			// HistorySearchFold:   true,
			// FuncFilterInputRune: filterInput,
		})
		if err != nil {
			panic(err)
		}
		defer rl.Close()

		for {
			fmt.Fprint(os.Stderr, "chirp% ") // Prompt to stderr.
			line, err := rl.Readline()
			if err != nil {
				if err.Error() == "EOF" { // TODO: better way?
					goto End
				}
				fmt.Fprintf(os.Stderr, "ERROR in Readline: %s\n", err.Error())
				goto End
			}
			result := EvalStringOrPrintError(fr, string(line))
			if result != "" { // Traditionally, if result is empty, tclsh doesn't print.
				fmt.Println(result)
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

	// Print summary for tests.
	if *testFlag {
		chirp.MustMutex.Lock()
		if chirp.MustFails > 0 {
			fmt.Fprintf(os.Stderr, "TEST FAILS: %q succeeds=%d fails=%d\n", scriptName, chirp.MustSucceeds, chirp.MustFails)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "Test Done: %q succeeds=%d\n", scriptName, chirp.MustSucceeds)
		chirp.MustMutex.Unlock()
	}
}

func EvalStringOrPrintError(fr *chirp.Frame, cmd string) (out string) {
	if *recoverFlag {
		defer func() {
			if r := recover(); r != nil {
				fmt.Fprintln(os.Stderr, "ERROR: ", r) // Error to stderr.
				out = ""
				return
			}
		}()
	}

	return fr.Eval(chirp.MkString(cmd)).String()
}
