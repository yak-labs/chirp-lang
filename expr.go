package chirp

import (
	// "bytes"
	// . "fmt"
	"runtime"
	// "strconv"
	"strings"
)

var _ = runtime.Breakpoint

var False = MkInt(0)
var True = MkInt(1)

// Concatenate the arguments, adding a space separator, before evaluating the
// expression.
func cmdExpr(fr *Frame, argv []T) T {
	// Optimization: It's already a single arg.
	if len(argv) == 2 {
		return fr.EvalExpr(argv[1])
	}

	strs := make([]string, len(argv)-1)

	for i, t := range argv[1:] {
		strs[i] = t.String()
	}

	ex := strings.Join(strs, " ")

	return fr.EvalExpr(MkString(ex))
}

// Takes a single word that represents an expression and returns the result.
func (fr *Frame) EvalExpr(a T) (result T) {
	return a.EvalExpr(fr)
	/*
		// Optimization: It's already a number.
		if a.IsQuickNumber() {
			return a
		}

		s := a.String()
		// Optimization: "0" and "1"
		if len(s) == 1 {
			if s == "0" {
				return False
			}
			if s == "1" {
				return True
			}
		}

		return fr.ParseExpression(s)
	*/
}

func init() {
	if Safes == nil {
		Safes = make(map[string]Command, 333)
	}

	Safes["expr"] = cmdExpr
}
