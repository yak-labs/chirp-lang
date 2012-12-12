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
	Builtins["get"] = cmdGet
	Builtins["set"] = cmdSet
	Builtins["proc"] = cmdProc
	Builtins["ls"] = cmdLs
	Builtins["slen"] = cmdSLen
	Builtins["llen"] = cmdLLen
	Builtins["list"] = cmdList
	Builtins["sat"] = cmdSAt
	Builtins["lat"] = cmdLAt
}

type BinaryFlop func(a, b float64) float64

func MkBinaryFlopCmd(fr *Frame, flop BinaryFlop) Command {
	return func(fr *Frame, argv List) Any {
		a, b := Argv2(argv)
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

func Argv1(argv List) Any {
	if len(argv) != 1+1 {
		panic(Sprintf("Expected 1 arguments, but got %#v", argv))
	}
	return argv[1]
}

func Argv2(argv List) (Any, Any) {
	if len(argv) != 2+1 {
		panic(Sprintf("Expected 2 arguments, but got %#v", argv))
	}
	return argv[1], argv[2]
}

func Argv3(argv List) (Any, Any, Any) {
	if len(argv) != 3+1 {
		panic(Sprintf("Expected 3 arguments, but got %#v", argv))
	}
	return argv[1], argv[2], argv[3]
}

func cmdMust(fr *Frame, argv List) Any {
	x := Str(argv[1])
	y := Str(argv[2])

	if x == y {
		return argv[2]
	}

	panic("FAILED: must: " + Repr(argv) + " -- x: " + x + " -- y: " + y)
}

func cmdIf(fr *Frame, argv List) Any {
	if len(argv) < 3+1 {
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

func cmdGet(fr *Frame, argv List) Any {
	name := Argv1(argv)
	return fr.GetVar(Str(name))
}

func cmdSet(fr *Frame, argv List) Any {
	name, x := Argv2(argv)
	fr.SetVar(Str(name), x)
	return x
}

func cmdProc(fr *Frame, argv List) Any {
	name, aa, body := Argv3(argv)
	alist := fr.ParseList(aa)
	astrs := make([]string, len(alist))
	for i, arg := range alist {
		astr := Str(arg)
		if !IsLocal(astr) {
			panic(Sprintf("Cannot use nonlocal name %q for argument in proc", arg))
		}
		astrs[i] = astr
	}
	n := len(alist) + 1
	cmd := func(fr2 *Frame, argv2 List) Any {
		if argv2 == nil {
			// Debug Data, if invoked with nil argv2.
			return argv
		}
		if len(argv2) != n {
			panic(Sprintf("Proc %q expects args %#v but got %#v", name, aa, argv2))
		}
		fr3 := fr2.NewFrame()
		for i, arg := range astrs {
			fr3.SetVar(arg, argv2[i+1])
		}
		return fr3.Eval(body)
	}
	fr.G.Cmds[Str(name)] = cmd
	return nil
}

func cmdLs(fr *Frame, argv List) Any {
	name := Argv1(argv)
	fn := fr.G.Cmds[Str(name)]
	return fn(fr, nil)
}

func cmdSLen(fr *Frame, argv List) Any {
	a := Argv1(argv)
	return float64(len(Str(a)))
}

func cmdLLen(fr *Frame, argv List) Any {
	a := Argv1(argv)
	return float64(len(fr.ParseList(a)))
}

func cmdList(fr *Frame, argv List) Any {
	return argv[1:]
}

func cmdLAt(fr *Frame, argv List) Any {
	v, j := Argv2(argv)
	f := ToFloat(j)
	i := int(f)
	return fr.ParseList(v)[i]
}

func cmdSAt(fr *Frame, argv List) Any {
	s, j := Argv2(argv)
	f := ToFloat(j)
	i := int(f)
	return Str(s)[i : i+1]
}
