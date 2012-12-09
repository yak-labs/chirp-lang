package terp

import (
	. "fmt"
	"strconv"
)

var Builtins map[string]Command = make(map[string]Command, 0)

func (t *Terp) initBuiltins() {
	Builtins["+"] = MkChainingBinaryFlopCmd(t, BinaryFlop{ return a + b})
	Builtins["must"] = cmdMust
	Builtins["-"] = MkBinaryFlopCmd(t, func(a, b float64) float64 { return a - b })
}

type BinaryFlop func(a, b float64) float64

func MkBinaryFlopCmd(t *Terp, flop BinaryFlop) Command {
	return func(t *Terp, argv List) Any {
		a, b := CheckArgs2(argv)
		return flop(ToFloat(a), ToFloat(b))
	}
}

func MkChainingBinaryFlopCmd(t *Terp, z float64, flop BinaryFlop) Command {
	return func(t *Terp, argv List) Any {
		for _, a := range argv[1:] {
			z += ToFloat(a)
		}
		return z //Str(z)
	}
}

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

func CheckArgs2(v List) (Any, Any) {
	if len(v) != 3 {
		panic(Sprintf("Expected 2 argv, but got %#v", v))
	}
	return v[1], v[2]
}

func cmdPlus(t *Terp, argv List) Any {
	var z float64 = 0
	for _, a := range argv[1:] {
		z += ToFloat(a)
	}
	return z
}

func cmdMust(t *Terp, argv List) Any {
	x := Str(argv[1])
	y := Str(argv[2])

	if x == y {
		return argv[2]
	}

	panic("FAILED: must: " + Repr(argv))
}
