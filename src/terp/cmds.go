package terp

import (
	"bytes"
	. "fmt"
	"log"
	"net/http"
)

var _ = log.Printf

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
	TBuiltins["puts"] = tcmdPuts
	TBuiltins["proc"] = tcmdProc
	TBuiltins["yproc"] = tcmdYProc
	TBuiltins["yield"] = tcmdYield
	TBuiltins["ls"] = tcmdLs
	TBuiltins["slen"] = tcmdSLen
	TBuiltins["llen"] = tcmdLLen
	TBuiltins["list"] = tcmdList
	TBuiltins["sat"] = tcmdSAt // a.k.a. string index
	TBuiltins["lat"] = tcmdLAt // a.k.a. lindex
	TBuiltins["http_handler"] = tcmdHttpHandler
	TBuiltins["foreach"] = tcmdForEach
	TBuiltins["while"] = tcmdWhile
	TBuiltins["catch"] = tcmdCatch
	TBuiltins["eval"] = tcmdEval
	TBuiltins["uplevel"] = tcmdUplevel
	TBuiltins["concat"] = tcmdConcat
	TBuiltins["set"] = tcmdSet
	TBuiltins["get"] = tcmdGet
	TBuiltins["return"] = tcmdReturn
	TBuiltins["break"] = tcmdBreak
	TBuiltins["continue"] = tcmdContinue
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

func TArgv1(argv []T) T {
	if len(argv) != 1+1 {
		panic(Sprintf("Expected 1 arguments, but got argv=%s", Showv(argv)))
	}
	return argv[1]
}

func TArgv1v(argv []T) (T, []T) {
	if len(argv) < 1+1 {
		panic(Sprintf("Expected at least 1 argument, but got argv=%s", Showv(argv)))
	}
	return argv[1], argv[2:]
}

func TArgv2(argv []T) (T, T) {
	if len(argv) != 2+1 {
		panic(Sprintf("Expected 2 arguments, but got argv=%s", Showv(argv)))
	}
	return argv[1], argv[2]
}

func TArgv2v(argv []T) (T, T, []T) {
	if len(argv) < 2+1 {
		panic(Sprintf("Expected at least 2 arguments, but got argv=%s", Showv(argv)))
	}
	return argv[1], argv[2], argv[3:]
}

func TArgv3(argv []T) (T, T, T) {
	if len(argv) != 3+1 {
		panic(Sprintf("Expected 3 arguments, but got argv=%s", Showv(argv)))
	}
	return argv[1], argv[2], argv[3]
}

func TArgv3v(argv []T) (T, T, T, []T) {
	if len(argv) < 3+1 {
		panic(Sprintf("Expected at least 3 arguments, but got argv=%s", Showv(argv)))
	}
	return argv[1], argv[2], argv[3], argv[4:]
}

func tcmdMust(fr *Frame, argv []T) T {
	xx, yy := TArgv2(argv)
	x := xx.String()
	y := yy.String()

	if x != y {
		panic("FAILED: must: " + Repr(argv) + " #### x=<" + x + "> #### y=<" + y + "> ####")
	}
	return Empty
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

	if TTruth(fr.TEvalExpr(cond)) {
		return fr.TEval(yes)
	}

	if no != nil {
		return fr.TEval(no)
	}

	return Empty
}

func tcmdPuts(fr *Frame, argv []T) T {
	// TODO:  accept a Writer as first arg.
	out := TArgv1(argv)
	Println(out)
	return Empty
}

func tcmdProc(fr *Frame, argv []T) T {
	name, aa, body := TArgv3(argv)
	alist := aa.List()
	astrs := make([]string, len(alist))
	for i, arg := range alist {
		astr := arg.String()
		if !IsLocal(astr) {
			panic(Sprintf("Cannot use nonlocal name %q for argument in proc", arg))
		}
		astrs[i] = astr
	}
	n := len(alist) + 1 // Add 1 for argv[0] now rather than at proc call.

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

func tcmdYProc(fr *Frame, argv []T) T {
	name, aa, body := TArgv3(argv)
	alist := aa.List()
	astrs := make([]string, len(alist))
	for i, arg := range alist {
		astr := arg.String()
		if !IsLocal(astr) {
			panic(Sprintf("Cannot use nonlocal name %q for argument in yproc", arg))
		}
		astrs[i] = astr
	}
	n := len(alist) + 1 // Add 1 for argv[0] now rather than at proc call.

	tcmd := func(fr2 *Frame, argv2 []T) T {
		if argv2 == nil {
			// Debug Data, if invoked with nil argv2.
			return MkTl(argv)
		}
		if len(argv2) != n {
			panic(Sprintf("yproc %q expects args %#v but got %#v", name, aa, argv2))
		}
		fr3 := fr2.NewFrame()
		for i, arg := range astrs {
			fr3.TSetVar(arg, argv2[i+1])
		}

		// Begin difference from Proc.
		ch := make(chan T, 0)
		fr3.Chan = ch

		go func() {
			fr3.TEval(body)
			close(ch)
		}()

		return MkTy(ch)
		// End difference from Proc.
	}

	fr.G.TCmds[name.String()] = tcmd
	return Empty
}

func tcmdYield(fr *Frame, argv []T) T {
	if len(argv)==2 {
		// Write exactly 1 arg on the channel.
		fr.Chan <- argv[1]
		return argv[1]
	}

	// Write more than 1 arg in a list.
	z := MkTl(argv[1:])
	fr.Chan <- z
	return z
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
	return MkTi(int64(len(a.List())))
}

func tcmdList(fr *Frame, argv []T) T {
	return MkTl(argv[1:])
}

func tcmdLAt(fr *Frame, argv []T) T {
	tlist, ti := TArgv2(argv)
	list := tlist.List()
	i := ti.Int()
	if i < 0 || i > int64(len(list)) {
		panic(Sprintf("lat: bad index: len(list)=%d but i=%d", len(list), i))
	}
	return list[i]
}

func tcmdSAt(fr *Frame, argv []T) T {
	s, j := TArgv2(argv)
	i := j.Int()
	return MkTs(s.String()[i : i+1])
}

func tcmdHttpHandler(fr *Frame, argv []T) T {
	fn := func(w http.ResponseWriter, r *http.Request) {
		v := make([]T, len(argv)-1)
		copy(v, argv[1:])
		v = append(v, MkT(w))
		v = append(v, MkT(r))
		_ = fr.TApply(v)
	}
	return MkT(fn)
}

func tcmdForEach(fr *Frame, argv []T) T {
	v, list, body := TArgv3(argv)

	for {
		hd, tl := list.HeadTail()
		if hd == nil {
			break
		}

		fr.TSetVar(v.String(), hd)
		fr.TEval(body)
		list = tl
	}

	return Empty
}

func tcmdWhile(fr *Frame, argv []T) T {
	cond, body := TArgv2(argv)

	for {
		c := fr.TEvalExpr(cond)
		if !c.Truth() {
			break
		}

		fr.TEval(body)
	}

	return Empty
}

func tcmdCatch(fr *Frame, argv []T) (status T) {
	body, varT := TArgv2(argv)
	varName := varT.String()

	defer func() {
		if r := recover(); r != nil {
			if j, ok := r.(Jump); ok {
				fr.TSetVar(varName, j.Result)
				status = MkTi(int64(j.Status))
				return
			}
			fr.TSetVar(varName, MkT(r))
			status = MkTi(1)
		}
	}()

	z := fr.TEval(body)
	fr.TSetVar(varName, z)
	return MkTi(0)
}

func tcmdEval(fr *Frame, argv []T) T {
	return EvalOrApplyLists(fr, argv[1:])
}

func tcmdUplevel(fr *Frame, argv []T) T {
	specArg, rest := TArgv1v(argv)
	spec := specArg.String()

	// Special case for #0 meaning global.
	if spec == "#0" {
		return EvalOrApplyLists(&fr.G.Fr, rest)
	}

	// Count back number of frames specified.
	level := specArg.Int()
	for i := int64(0); i < level; i++ {
		if fr.Prev != nil {
			fr = fr.Prev
		}
	}
	return EvalOrApplyLists(fr, rest)
}

func EvalOrApplyLists(fr *Frame, lists []T) T {
	// Are they already lists?
	areLists := true
	for _, e := range lists {
		_, ok := e.(Tl)
		if !ok {
			areLists = false
		}
	}

	if areLists {
		return fr.TApply(ConcatLists(lists))
	}

	buf := bytes.NewBuffer(nil)
	for _, e := range lists {
		buf.WriteString(e.String())
		buf.WriteRune(' ')
	}
	return fr.TEval(MkTs(buf.String()))
}

func ConcatLists(lists []T) []T {
	z := make([]T, 0, 4)
	for _, e := range lists {
		z = append(z, e.List() ...)
	}
	return z
}

func tcmdConcat(fr *Frame, argv []T) T {
	return MkTl(ConcatLists(argv[1:]))
}

func tcmdGet(fr *Frame, argv []T) T {
    name := TArgv1(argv)
    return fr.TGetVar(name.String())
}

func tcmdSet(fr *Frame, argv []T) T {
	if len(argv) == 2 {
		// Normal Tcl allows 'set' to 'get'
		return tcmdGet(fr, argv)
	}
    name, x := TArgv2(argv)
    fr.TSetVar(name.String(), x)
    return x
}

func tcmdReturn(fr *Frame, argv []T) T {
	var z T
	if len(argv) == 2 {
		z = argv[1]
	}
	if len(argv) > 2 {
		z = MkTl(argv[1:])
	}
	panic(Jump{Status: RETURN, Result: z})
}

func tcmdBreak(fr *Frame, argv []T) T {
	panic(Jump{Status: BREAK})
}

func tcmdContinue(fr *Frame, argv []T) T {
	panic(Jump{Status: CONTINUE})
}
