package terp

import (
	"bytes"
	. "fmt"
	"log"
	"net/http"
)

var _ = log.Printf

var Builtins map[string]Command = make(map[string]Command, 0)

func (fr *Frame) initTBuiltins() {
	Builtins["+"] = MkChainingBinaryFlopTCmd(fr, 0.0, func(a, b float64) float64 { return a + b })
	Builtins["*"] = MkChainingBinaryFlopTCmd(fr, 1.0, func(a, b float64) float64 { return a * b })
	Builtins["-"] = MkBinaryFlopTCmd(fr, func(a, b float64) float64 { return a - b })
	Builtins["/"] = MkBinaryFlopTCmd(fr, func(a, b float64) float64 { return a / b })

	Builtins["=="] = MkBinaryFlopBoolTCmd(fr, func(a, b float64) bool { return (a == b) })
	Builtins["!="] = MkBinaryFlopBoolTCmd(fr, func(a, b float64) bool { return (a != b) })
	Builtins["<"] = MkBinaryFlopBoolTCmd(fr, func(a, b float64) bool { return (a < b) })
	Builtins["<="] = MkBinaryFlopBoolTCmd(fr, func(a, b float64) bool { return (a <= b) })
	Builtins[">"] = MkBinaryFlopBoolTCmd(fr, func(a, b float64) bool { return (a > b) })
	Builtins[">="] = MkBinaryFlopBoolTCmd(fr, func(a, b float64) bool { return (a >= b) })
	Builtins["must"] = tcmdMust

	Builtins["if"] = tcmdIf
	Builtins["puts"] = tcmdPuts
	Builtins["proc"] = tcmdProc
	Builtins["yproc"] = tcmdYProc
	Builtins["yield"] = tcmdYield
	Builtins["ls"] = tcmdLs
	Builtins["slen"] = tcmdSLen
	Builtins["llen"] = tcmdLLen
	Builtins["list"] = tcmdList
	Builtins["sat"] = tcmdSAt // a.k.a. string index
	Builtins["lat"] = tcmdLAt // a.k.a. lindex
	Builtins["http_handler"] = tcmdHttpHandler
	Builtins["foreach"] = tcmdForEach
	Builtins["while"] = tcmdWhile
	Builtins["catch"] = tcmdCatch
	Builtins["eval"] = tcmdEval
	Builtins["uplevel"] = tcmdUplevel
	Builtins["concat"] = tcmdConcat
	Builtins["set"] = tcmdSet
	Builtins["upvar"] = tcmdUpVar
	Builtins["return"] = tcmdReturn
	Builtins["break"] = tcmdBreak
	Builtins["continue"] = tcmdContinue
	Builtins["hash"] = tcmdHash
	Builtins["hget"] = tcmdHGet   // FIXME: temporary: Use getf
	Builtins["hset"] = tcmdHSet   // FIXME: temporary: Use setf
	Builtins["hdel"] = tcmdHDel   // FIXME: temporary: Use delf
	Builtins["hkeys"] = tcmdHKeys // FIXME: temporary: use keys
}

type BinaryFlop func(a, b float64) float64
type BinaryFlopBool func(a, b float64) bool

func MkBinaryFlopTCmd(fr *Frame, flop BinaryFlop) Command {
	return func(fr *Frame, argv []T) T {
		a, b := TArgv2(argv)
		return MkTf(flop(a.Float(), b.Float()))
	}
}

func MkBinaryFlopBoolTCmd(fr *Frame, flop BinaryFlopBool) Command {
	return func(fr *Frame, argv []T) T {
		a, b := TArgv2(argv)
		return MkTb(flop(a.Float(), b.Float()))
	}
}

func MkChainingBinaryFlopTCmd(fr *Frame, starter float64, flop BinaryFlop) Command {
	return func(fr *Frame, argv []T) T {
		z := starter // Be sure not to modify starter!  It is captured.
		for _, a := range argv[1:] {
			z = flop(z, a.Float())
		}
		return MkTf(z)
	}
}

func Arg1(argv []T) T {
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

	if fr.TEvalExpr(cond).Truth() {
		return fr.Eval(yes)
	}

	if no != nil {
		return fr.Eval(no)
	}

	return Empty
}

func tcmdPuts(fr *Frame, argv []T) T {
	// TODO:  accept a Writer as first arg.
	out := Arg1(argv)
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

	tcmd := func(fr2 *Frame, argv2 []T) (result T) {
		defer func() {
			if r := recover(); r != nil {
				if j, ok := r.(Jump); ok {
					switch j.Status {
					case RETURN:
						result = j.Result
						return
					case BREAK:
						panic("break command was not inside a loop")
					case CONTINUE:
						panic("continue command was not inside a loop")
					}
				}
				panic(r) // Rethrow errors and unknown Status.
			}
		}()

		if argv2 == nil {
			// Debug Data, if invoked with nil argv2.
			return MkTl(argv)
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

	fr.G.Cmds[name.String()] = tcmd
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
			fr3.SetVar(arg, argv2[i+1])
		}

		// Begin difference from Proc.
		ch := make(chan T, 0)
		fr3.Chan = ch

		go func() {
			defer close(ch)
			defer func() {
				if r := recover(); r != nil {
					if j, ok := r.(Jump); ok {
						switch j.Status {
						case RETURN:
							if !j.Result.IsEmpty() {
								panic("cannot return a value inside a yproc command")
							}
							return
						case BREAK:
							panic("break command was not inside a loop")
						case CONTINUE:
							panic("continue command was not inside a loop")
						}
					}
					panic(r) // Rethrow errors and unknown Status.
				}
			}()
			fr3.Eval(body)
		}()

		return MkTy(ch)
		// End difference from Proc.
	}

	fr.G.Cmds[name.String()] = tcmd
	return Empty
}

func tcmdYield(fr *Frame, argv []T) T {
	if len(argv) == 2 {
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
	a := Arg1(argv)
	return MkTi(int64(len(a.String())))
}

func tcmdLLen(fr *Frame, argv []T) T {
	a := Arg1(argv)
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
		_ = fr.Apply(v)
	}
	return MkT(fn)
}

func tcmdForEach(fr *Frame, argv []T) T {
	v, list, body := TArgv3(argv)

	toBreak := false
	toContinue := false

	for {
		hd, tl := list.HeadTail()
		if hd == nil {
			break
		}

		fr.SetVar(v.String(), hd)
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("foreach recovered: %#v", r)
					if j, ok := r.(Jump); ok {
						switch j.Status {
						case BREAK:
							toBreak = true
							return
						case CONTINUE:
							toContinue = true
							return
						}
					}
					panic(r) // Rethrow errors and unknown Status.
				}
			}()
			log.Printf("foreach before: %q", body.String())
			fr.Eval(body)
			log.Printf("foreach after: %q", body.String())
		}()
		if toBreak {
			log.Printf("foreach breaks ======================================")
			break
		}
		if toContinue {
			log.Printf("foreach continues =====================================")
			continue
		}
		list = tl
	}

	return Empty
}

func tcmdWhile(fr *Frame, argv []T) T {
	cond, body := TArgv2(argv)

	toBreak := false
	toContinue := false

	for {
		c := fr.TEvalExpr(cond)
		if !c.Truth() {
			break
		}

		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("while recovered: %#v", r)
					if j, ok := r.(Jump); ok {
						switch j.Status {
						case BREAK:
							toBreak = true
							return
						case CONTINUE:
							toContinue = true
							return
						}
					}
					panic(r) // Rethrow errors and unknown Status.
				}
			}()
			log.Printf("while before: %q", body.String())
			fr.Eval(body)
			log.Printf("while after: %q", body.String())
		}()
		if toBreak {
			log.Printf("while breaks ======================================")
			break
		}
		if toContinue {
			log.Printf("while continues =====================================")
			continue
		}
	}

	return Empty
}

func tcmdCatch(fr *Frame, argv []T) (status T) {
	body, varT := TArgv2(argv)
	varName := varT.String()

	defer func() {
		if r := recover(); r != nil {
			if j, ok := r.(Jump); ok {
				fr.SetVar(varName, j.Result)
				status = MkTi(int64(j.Status))
				return
			}
			fr.SetVar(varName, MkT(r))
			status = MkTi(1)
		}
	}()

	z := fr.Eval(body)
	fr.SetVar(varName, z)
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
		if !e.IsPreservedByList() {
			areLists = false
			break
		}
	}

	if areLists {
		return fr.Apply(ConcatLists(lists))
	}

	buf := bytes.NewBuffer(nil)
	for _, e := range lists {
		buf.WriteString(e.String())
		buf.WriteRune(' ')
	}
	return fr.Eval(MkTs(buf.String()))
}

func ConcatLists(lists []T) []T {
	z := make([]T, 0, 4)
	for _, e := range lists {
		z = append(z, e.List()...)
	}
	return z
}

func tcmdConcat(fr *Frame, argv []T) T {
	return MkTl(ConcatLists(argv[1:]))
}

func tcmdUpVar(fr *Frame, argv []T) T {
	lev, rem, loc := TArgv3(argv)
	level := lev.Int()
	remName := rem.String()
	locName := loc.String()
	remFr := fr
	for i := 0; i < int(level); i++ {
		remFr = remFr.Prev
	}
	fr.UpVar(locName, remFr, remName)
	return Empty
}

func tcmdSet(fr *Frame, argv []T) T {
	if len(argv) == 2 {
		// Retrieve value of variable, if 2nd arg is missing.
		name := Arg1(argv)
		return fr.GetVar(name.String())
	}
	name, x := TArgv2(argv)
	fr.SetVar(name.String(), x)
	return x
}

func tcmdReturn(fr *Frame, argv []T) T {
	var z T = Empty
	if len(argv) == 2 {
		z = argv[1]
	}
	if len(argv) > 2 {
		z = MkTl(argv[1:])
	}
	// Jump with status RETURN.
	panic(Jump{Status: RETURN, Result: z})
}

func tcmdBreak(fr *Frame, argv []T) T {
	panic(Jump{Status: BREAK}) // Jump with status BREAK.
}

func tcmdContinue(fr *Frame, argv []T) T {
	panic(Jump{Status: CONTINUE}) // Jump with status CONTINUE.
}

func tcmdHash(fr *Frame, argv []T) T {
	return MkTh()
}

func tcmdHGet(fr *Frame, argv []T) T {
	hash, key := TArgv2(argv)
	h := hash.Hash()
	k := key.String()
	value := h[k]
	if value == nil {
		panic(Sprintf("Hash does not contain key: %q", k))
	}
	return value
}

func tcmdHSet(fr *Frame, argv []T) T {
	hash, key, value := TArgv3(argv)
	h := hash.Hash()
	k := key.String()
	h[k] = value
	return value
}

func tcmdHDel(fr *Frame, argv []T) T {
	hash, key := TArgv2(argv)
	h := hash.Hash()
	k := key.String()
	h[k] = nil // TODO: how to delete?
	return Empty
}

func tcmdHKeys(fr *Frame, argv []T) T {
	hash := Arg1(argv)
	h := hash.Hash()
	z := make([]T, 0, len(h))
	for _, k := range SortedKeysOfHash(h) {
		z = append(z, MkTs(k))
	}
	return MkTl(z)
}
