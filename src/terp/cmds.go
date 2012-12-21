package terp

import (
	. "fmt"
	"log"
	"net/http"
	"strconv"
)

var _ = log.Printf
var _ = strconv.ParseFloat

var TBuiltins map[string]TCommand = make(map[string]TCommand, 0)

func (fr *Frame) initTBuiltins() {
	TBuiltins["+"] = MkChainingBinaryFlopTCmd(fr, 0.0, func(a, b float64) float64 { return a + b })
	TBuiltins["*"] = MkChainingBinaryFlopTCmd(fr, 1.0, func(a, b float64) float64 { return a * b })
	TBuiltins["-"] = MkBinaryFlopTCmd(fr, func(a, b float64) float64 { return a - b })
	TBuiltins["/"] = MkBinaryFlopTCmd(fr, func(a, b float64) float64 { return a / b })

	TBuiltins["=="] = MkBinaryFlopBoolTCmd(fr, func(a, b float64) bool { return (a == b) })
	TBuiltins["!="] = MkBinaryFlopBoolTCmd(fr, func(a, b float64) bool { return (a != b) })
	TBuiltins["<"] = MkBinaryFlopBoolTCmd(fr, func(a, b float64) bool { return (a < b) })
	TBuiltins["<="] = MkBinaryFlopBoolTCmd(fr, func(a, b float64) bool { return (a <= b) })
	TBuiltins[">"] = MkBinaryFlopBoolTCmd(fr, func(a, b float64) bool { return (a > b) })
	TBuiltins[">="] = MkBinaryFlopBoolTCmd(fr, func(a, b float64) bool { return (a >= b) })
	TBuiltins["must"] = tcmdMust

	TBuiltins["if"] = tcmdIf
	TBuiltins["get"] = tcmdGet
	TBuiltins["getq"] = tcmdGetQ
	TBuiltins["set"] = tcmdSet
	TBuiltins["puts"] = tcmdPuts
	TBuiltins["proc"] = tcmdProc
	TBuiltins["ls"] = tcmdLs
	TBuiltins["slen"] = tcmdSLen
	TBuiltins["llen"] = tcmdLLen
	TBuiltins["list"] = tcmdList
	TBuiltins["sat"] = tcmdSAt  // a.k.a. string index
	TBuiltins["lat"] = tcmdLAt  // a.k.a. lindex
	TBuiltins["http_handler"] = newcmd(cmdHttpHandler)
	TBuiltins["foreach"] = tcmdForEach
}
type BinaryFlop func(a, b float64) float64
type BinaryFlopBool func(a, b float64) bool

func MkBinaryFlopTCmd(fr *Frame, flop BinaryFlop) TCommand {
	return func(fr *Frame, argv []T) T {
		a, b := TArgv2(argv)
		return MkTf(flop(a.Float(), b.Float()))
	}
}

func MkBinaryFlopBoolTCmd(fr *Frame, flop BinaryFlopBool) TCommand {
	return func(fr *Frame, argv []T) T {
		a, b := TArgv2(argv)
		return MkTb(flop(a.Float(), b.Float()))
	}
}

func MkChainingBinaryFlopTCmd(fr *Frame, starter float64, flop BinaryFlop) TCommand {
	return func(fr *Frame, argv []T) T {
		z := starter // Be sure not to modify starter!  It is captured.
		for _, a := range argv[1:] {
			z = flop(z, a.Float())
		}
		return MkTf(z)
	}
}

func TTruth(a T) bool {
	switch x := a.(type) {
	case Tf:
		return x.f != 0
	case Ts:
		return len(x.s) > 0
	case Tl:
		return len(x.l) > 0
	}
	// To Do: Value(nil) Value(false) are false.
	return true
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

func TArgv1(argv []T) T {
	if len(argv) != 1+1 {
		panic(Sprintf("Expected 1 arguments, but got argv=%#v", argv))
	}
	return argv[1]
}

func TArgv2(argv []T) (T, T) {
	if len(argv) != 2+1 {
		panic(Sprintf("Expected 2 arguments, but got argv=%#v", argv))
	}
	return argv[1], argv[2]
}

func TArgv3(argv []T) (T, T, T) {
	if len(argv) != 3+1 {
		panic(Sprintf("Expected 3 arguments, but got argv=%#v", argv))
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
func tcmdMust(fr *Frame, argv []T) T {
	xx, yy := TArgv2(argv)
	x := xx.String()
	y := yy.String()

	if x == y {
		return argv[2]
	}

	panic("FAILED: must: " + Repr(argv) + " -- x: " + x + " -- y: " + y)
}

func tcmdIf(fr *Frame, argv []T) T {
	if len(argv) < 3 {
		panic(Sprintf("Too few arguments for if: %#v", argv))
	}
	var cond, yes, no T

	switch len(argv) {
	case 5:
		if argv[3].String() != "else" {
			panic(Sprintf("Expected 'else' at argv[3]: %#v", argv))
		}
		cond, yes, no = argv[1], argv[2], argv[4]
	case 3:
		cond, yes = argv[1], argv[2]
	default:
		panic(Sprintf("Wrong len(argv) for if: %#v", argv))
	}

	if TTruth(fr.TEval(cond)) {
		return fr.TEval(yes)
	}

	if no != nil {
		return fr.TEval(no)
	}

	return Empty
}

func tcmdGet(fr *Frame, argv []T) T {
	name := TArgv1(argv)
	return fr.TGetVar(name.String())
}

func tcmdGetQ(fr *Frame, argv []T) T {
	req, key := TArgv2(argv)

	v := req.(Tv)
	r := v.v.Interface().(*http.Request)

	return MkTs(r.URL.Query().Get(key.String()))
}

func tcmdSet(fr *Frame, argv []T) T {
	name, x := TArgv2(argv)
	fr.TSetVar(name.String(), x)
	return x
}

func tcmdPuts(fr *Frame, argv []T) T {
	// TODO:  accept a Writer as first arg.  
	out := TArgv1(argv)
	Println(out)
	return Empty
}

func tcmdProc(fr *Frame, argv []T) T {
	name, aa, body := TArgv3(argv)
	alist := newlist(ParseList(aa.String()))
	astrs := make([]string, len(alist))
	for i, arg := range alist {
		astr := arg.String()
		if !IsLocal(astr) {
			panic(Sprintf("Cannot use nonlocal name %q for argument in proc", arg))
		}
		astrs[i] = astr
	}
	n := len(alist) + 1  // Add 1 for argv[0] now rather than at proc call.

	tcmd := func(fr2 *Frame, argv2 []T) T {
		if argv2 == nil {
			// Debug Data, if invoked with nil argv2.
			return MkTl(argv)
		}
		if len(argv2) != n {
			panic(Sprintf("Proc %q expects args %#v but got %#v", name, aa, argv2))
		}
		fr3 := fr2.NewFrame()
		for i, arg := range astrs {
			fr3.TSetVar(arg, argv2[i+1])
		}
		return fr3.TEval(body)
	}

	fr.G.TCmds[name.String()] = tcmd
	return Empty
}

func tcmdLs(fr *Frame, argv []T) T {
	panic("not usefully implemented yet")
}

func tcmdSLen(fr *Frame, argv []T) T {
	a := TArgv1(argv)
	return MkTi(int64(len(a.String())))
}

func tcmdLLen(fr *Frame, argv []T) T {
	a := TArgv1(argv)
	return MkTi(int64(len(ParseList(a.String()))))
}

func tcmdList(fr *Frame, argv []T) T {
	return MkTl(argv[1:])
}

func tcmdLAt(fr *Frame, argv []T) T {
	v, j := TArgv2(argv)
	i := j.Int()
	return new(ParseList(v)[i])
}

func tcmdSAt(fr *Frame, argv []T) T {
	s, j := TArgv2(argv)
	i := j.Int()
	return MkTs(s.String()[i : i+1])
}

func cmdHttpHandler(fr *Frame, argv List) Any {
	return func (w http.ResponseWriter, r *http.Request) {
		v := make(List, len(argv)-1)
		copy(v, argv[1:])
		v = LAppend(v, w)
		v = LAppend(v, r)
		_ = fr.Apply(v)
	}
}

func tcmdForEach(fr *Frame, argv []T) T {
	v, list, body := TArgv3(argv)

	l := list.(Tl).l
	
	for _, e := range l {
		fr.TSetVar(v.String(), e)
		fr.Eval(body.String())
	}

	return Empty
}

