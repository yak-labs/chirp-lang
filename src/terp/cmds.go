package terp

import (
	. "fmt"
	"strconv"
)

var Builtins map[string]Command = make(map[string]Command, 0)

func (t *Terp) initBuiltins() {
	Builtins["+"] = cmdPlus
	Builtins["must"] = cmdMust
	Builtins["-"] = MkBinaryFlopCmd(t, func(a, b float64) float64 { return a - b })
}

type BinaryFlop func(a, b float64) float64

func ToFloat(a Any) float64 {
	switch x := a.(type) {
	case float64:
		return x
	case float32:
		return float64(x)
	case uint64:
		return float64(x)
	case int64:
		return float64(x)
	case int:
		return float64(x)
	case uint:
		return float64(x)
	}
	f, err := strconv.ParseFloat(Str(a), 64)
	if err != nil {
		panic(Sprintf("Cannot Parse Float: %q", a))
	}
	return f
}

func CheckListLen(v List, n int) List {
	if len(v) != n {
		panic(Sprintf("CheckListLen: requires 2 elements (or command argv), but got %#v", v))
	}
	return v
}

func CheckArgs2(v List) (Any, Any) {
	if len(v) != 2 {
		panic(Sprintf("Expected 2 argv, but got %#v", v))
	}
	return v[0], v[1]
}

func MkBinaryFlopCmd(t *Terp, flop BinaryFlop) Command {
	return func(t *Terp, argv List) Any {
		a, b := CheckArgs2(argv)
		return flop(ToFloat(a), ToFloat(b))
	}
}

func MkChainingBinaryFlopCmd(t *Terp, z float64, flop BinaryFlop) Command {
	return func(t *Terp, argv List) Any {
		for _, a := range argv {
			z += ToFloat(a)
		}
		return z //Str(z)
	}
}

func cmdPlus(t *Terp, argv List) Any {
	var z float64 = 0
	for _, a := range argv {
		z += ToFloat(a)
	}
	return z //Str(z)
}

func cmdMust(t *Terp, argv List) Any {
	x := Str(argv[0])
	y := Str(argv[1])

	if x == y {
		return argv[1]
	}

	panic("FAILED: must: " + Repr(argv))
}
