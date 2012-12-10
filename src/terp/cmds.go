package terp

import (
	. "fmt"
	"log"
	"strconv"
)

var Builtins map[string]Command = make(map[string]Command, 0)

func (fr *Frame) initBuiltins() {
	Builtins["+"] = MkChainingBinaryFlopCmd(fr, 0.0, func(a, b float64) float64 { return a + b })
	Builtins["*"] = MkChainingBinaryFlopCmd(fr, 1.0, func(a, b float64) float64 { return a * b })
	Builtins["-"] = MkBinaryFlopCmd(fr, func(a, b float64) float64 { return a - b })
	Builtins["/"] = MkBinaryFlopCmd(fr, func(a, b float64) float64 { return a / b })

	Builtins["=="] = MkBinaryFlopCmd(fr, func(a, b float64) float64 { return ToFloat(a == b) })
	Builtins["!="] = MkBinaryFlopCmd(fr, func(a, b float64) float64 { return ToFloat(a != b) })
	Builtins["<"] = MkBinaryFlopCmd(fr, func(a, b float64) float64 { return ToFloat(a < b) })
	Builtins["<="] = MkBinaryFlopCmd(fr, func(a, b float64) float64 { return ToFloat(a <= b) })
	Builtins[">"] = MkBinaryFlopCmd(fr, func(a, b float64) float64 { return ToFloat(a > b) })
	Builtins[">="] = MkBinaryFlopCmd(fr, func(a, b float64) float64 { return ToFloat(a >= b) })

	Builtins["must"] = cmdMust
	Builtins["if"] = cmdIf
}

type BinaryFlop func(a, b float64) float64

func MkBinaryFlopCmd(fr *Frame, flop BinaryFlop) Command {
	return func(fr *Frame, argv List) Any {
		a, b := CheckArgv2(argv)
		return flop(ToFloat(a), ToFloat(b))
	}
}

func MkChainingBinaryFlopCmd(fr *Frame, starter float64, flop BinaryFlop) Command {
	return func(fr *Frame, argv List) Any {
		z := starter // Be sure not to modify starter!  It is captured.
		for _, a := range argv[1:] {
			z = flop(z, ToFloat(a))
		}
		return z //Str(z)
	}
}

func Truth(a Any) bool {
	switch x := a.(type) {
	case float64:
		return x != 0
	case float32:
		return x != 0
	case uint64:
		return x != 0
	case int64:
		return x != 0
	case int:
		return x != 0
	case uint:
		return x != 0
	case byte:
		return x != 0
	case rune:
		return x != 0
	case bool:
		return x
	case string:
		return len(x) != 0
	case List:
		return len(x) != 0
	case Dict:
		return len(x) != 0
	}
	panic(Sprintf("Case not supported in Truth(): %#v", a))
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
	case byte:
		return float64(x)
	case rune:
		return float64(x)
	case bool:
		log.Printf("BOOL %v", x)
		if x {
			return 1.0
		}
		return 0.0
	}
	f, err := strconv.ParseFloat(Str(a), 64)
	if err != nil {
		panic(Sprintf("Cannot Parse Float: %q", a))
	}
	return f
}

func CheckArgv2(argv List) (Any, Any) {
	if len(argv) != 2 + 1 {
		panic(Sprintf("Expected 2 arguments, but got %#v", argv))
	}
	return argv[1], argv[2]
}

func cmdMust(fr *Frame, argv List) Any {
	x := Str(argv[1])
	y := Str(argv[2])

	if x == y {
		return argv[2]
	}

	panic("FAILED: must: " + Repr(argv))
}

func cmdIf(fr *Frame, argv List) Any {
	if len(argv) < 3 + 1 {
		panic(Sprintf("Too few arguments for if: %#v", argv))
	}
	var cond, yes, no Any

	switch len(argv) {
	case 5:
		if argv[3] != "else" {
			panic(Sprintf("Expected 'else' at argv[3]: %#v", argv))
		}
		cond, yes, no = argv[1], argv[2], argv[4]
	case 3:
		cond, yes = argv[1], argv[2]
	default:
		panic(Sprintf("Wrong len(argv) for if: %#v", argv))
	}

	if Truth(fr.Eval(cond)) {
		return fr.Eval(yes)
	}

	if no != nil {
		return fr.Eval(no)
	}
		
	return nil
}


