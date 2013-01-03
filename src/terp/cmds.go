package terp

import (
	"bytes"
	. "fmt"
	"log"
	"net/http"
)

var _ = log.Printf

var Builtins map[string]Command = make(map[string]Command, 0)

func (fr *Frame) initBuiltins() {
	Builtins["+"] = MkChainingBinaryFlopCmd(fr, 0.0, func(a, b float64) float64 { return a + b })
	Builtins["*"] = MkChainingBinaryFlopCmd(fr, 1.0, func(a, b float64) float64 { return a * b })
	Builtins["-"] = MkBinaryFlopCmd(fr, func(a, b float64) float64 { return a - b })
	Builtins["/"] = MkBinaryFlopCmd(fr, func(a, b float64) float64 { return a / b })

	Builtins["=="] = MkBinaryFlopBoolCmd(fr, func(a, b float64) bool { return (a == b) })
	Builtins["!="] = MkBinaryFlopBoolCmd(fr, func(a, b float64) bool { return (a != b) })
	Builtins["<"] = MkBinaryFlopBoolCmd(fr, func(a, b float64) bool { return (a < b) })
	Builtins["<="] = MkBinaryFlopBoolCmd(fr, func(a, b float64) bool { return (a <= b) })
	Builtins[">"] = MkBinaryFlopBoolCmd(fr, func(a, b float64) bool { return (a > b) })
	Builtins[">="] = MkBinaryFlopBoolCmd(fr, func(a, b float64) bool { return (a >= b) })
	Builtins["must"] = cmdMust

	Builtins["if"] = cmdIf
	Builtins["puts"] = cmdPuts
	Builtins["proc"] = cmdProc
	Builtins["yproc"] = cmdYProc
	Builtins["mixin"] = cmdMixin
	Builtins["super"] = cmdSuper
	Builtins["yield"] = cmdYield
	Builtins["ls"] = cmdLs
	Builtins["slen"] = cmdSLen
	Builtins["llen"] = cmdLLen
	Builtins["list"] = cmdList
	Builtins["sat"] = cmdSAt // a.k.a. string index
	Builtins["lat"] = cmdLAt // a.k.a. lindex
	Builtins["http_handler"] = cmdHttpHandler
	Builtins["foreach"] = cmdForEach
	Builtins["while"] = cmdWhile
	Builtins["catch"] = cmdCatch
	Builtins["eval"] = cmdEval
	Builtins["uplevel"] = cmdUplevel
	Builtins["concat"] = cmdConcat
	Builtins["set"] = cmdSet
	Builtins["upvar"] = cmdUpVar
	Builtins["return"] = cmdReturn
	Builtins["break"] = cmdBreak
	Builtins["continue"] = cmdContinue
	Builtins["hash"] = cmdHash
	Builtins["hget"] = cmdHGet   // FIXME: temporary: Use getf
	Builtins["hset"] = cmdHSet   // FIXME: temporary: Use setf
	Builtins["hdel"] = cmdHDel   // FIXME: temporary: Use delf
	Builtins["hkeys"] = cmdHKeys // FIXME: temporary: use keys
	Builtins["interp-new"] = cmdInterpNew
	Builtins["interp-alias"] = cmdInterpAlias
	Builtins["interp-eval"] = cmdInterpEval
}

type BinaryFlop func(a, b float64) float64
type BinaryFlopBool func(a, b float64) bool

func MkBinaryFlopCmd(fr *Frame, flop BinaryFlop) Command {
	return func(fr *Frame, argv []T) T {
		a, b := Arg2(argv)
		return MkFloat(flop(a.Float(), b.Float()))
	}
}

func MkBinaryFlopBoolCmd(fr *Frame, flop BinaryFlopBool) Command {
	return func(fr *Frame, argv []T) T {
		a, b := Arg2(argv)
		return MkBool(flop(a.Float(), b.Float()))
	}
}

func MkChainingBinaryFlopCmd(fr *Frame, starter float64, flop BinaryFlop) Command {
	return func(fr *Frame, argv []T) T {
		z := starter // Be sure not to modify starter!  It is captured.
		for _, a := range argv[1:] {
			z = flop(z, a.Float())
		}
		return MkFloat(z)
	}
}

func Arg1(argv []T) T {
	if len(argv) != 1+1 {
		panic(Sprintf("Expected 1 arguments, but got argv=%s", Showv(argv)))
	}
	return argv[1]
}

func Arg1v(argv []T) (T, []T) {
	if len(argv) < 1+1 {
		panic(Sprintf("Expected at least 1 argument, but got argv=%s", Showv(argv)))
	}
	return argv[1], argv[2:]
}

func Arg2(argv []T) (T, T) {
	if len(argv) != 2+1 {
		panic(Sprintf("Expected 2 arguments, but got argv=%s", Showv(argv)))
	}
	return argv[1], argv[2]
}

func Arg2v(argv []T) (T, T, []T) {
	if len(argv) < 2+1 {
		panic(Sprintf("Expected at least 2 arguments, but got argv=%s", Showv(argv)))
	}
	return argv[1], argv[2], argv[3:]
}

func Arg3(argv []T) (T, T, T) {
	if len(argv) != 3+1 {
		panic(Sprintf("Expected 3 arguments, but got argv=%s", Showv(argv)))
	}
	return argv[1], argv[2], argv[3]
}

func Arg3v(argv []T) (T, T, T, []T) {
	if len(argv) < 3+1 {
		panic(Sprintf("Expected at least 3 arguments, but got argv=%s", Showv(argv)))
	}
	return argv[1], argv[2], argv[3], argv[4:]
}

func cmdMust(fr *Frame, argv []T) T {
	xx, yy := Arg2(argv)
	x := xx.String()
	y := yy.String()

	if x != y {
		panic("FAILED: must: " + Repr(argv) + " #### x=<" + x + "> #### y=<" + y + "> ####")
	}
	return Empty
}

func cmdIf(fr *Frame, argv []T) T {
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

	if fr.EvalExpr(cond).Truth() {
		return fr.Eval(yes)
	}

	if no != nil {
		return fr.Eval(no)
	}

	return Empty
}

func cmdPuts(fr *Frame, argv []T) T {
	// TODO:  accept a Writer as first arg.
	out := Arg1(argv)
	Println(out)
	return Empty
}

func cmdProc(fr *Frame, argv []T) T {
	return procOrYProc(fr, argv, false)
}

func cmdYProc(fr *Frame, argv []T) T {
	return procOrYProc(fr, argv, true)
}

func procOrYProc(fr *Frame, argv []T, generating bool) T {
	name, aa, body := Arg3(argv)
	nameStr := name.String()
	alist := aa.List()
	astrs := make([]string, len(alist))
	for i, arg := range alist {
		astr := arg.String()
		if !IsLocal(astr) {
			panic(Sprintf("Cannot use nonlocal name %q for argument in %s", astr, argv[0]))
		}
		astrs[i] = astr
	}
	n := len(alist) + 1 // Add 1 for argv[0] now rather than at proc call.

	// Capture this variable, so it can be used when the cmd is called.
	captureMixinNumberDefining := fr.G.MixinNumberDefining
	captureMixinNameDefining := fr.G.MixinNameDefining
	var longMixinName string
	if captureMixinNumberDefining > 0 {
		longMixinName = nameStr + "~" + captureMixinNameDefining
	}

	cmd := func(fr2 *Frame, argv2 []T) (result T) {
		if captureMixinNumberDefining > 0 {
			argv2[0] = MkString(longMixinName)
		}
		// If generating, not enough happens in this func (as opposed to
		// in the goroutine) to encounter errors.  So this defer/recover is only
		// for the normal, nongenerating case.
		if !generating {
			defer func() {
				if r := recover(); r != nil {
					if j, ok := r.(Jump); ok {
						switch j.Status {
						case RETURN:
							result = j.Result
							return
						case BREAK:
							r =("break command was not inside a loop")
						case CONTINUE:
							r =("continue command was not inside a loop")
						}
					}
					if rs, ok := r.(string); ok {
						r = rs + "\n\tin proc " + argv[0].String()
					}
					panic(r) // Rethrow errors and unknown Status.
				}
			}()
		}

		if argv2 == nil {
			// Debug Data, if invoked with nil argv2.
			return MkList(argv)
		}
		if len(argv2) != n {
			panic(Sprintf("%s %q expects args %#v but got %#v", argv[0], nameStr, aa, argv2))
		}
		fr3 := fr2.NewFrame()
		fr3.MixinLevel = captureMixinNumberDefining
		for i, arg := range astrs {
			fr3.SetVar(arg, argv2[i+1])
		}

		// Case "proc":
		if !generating {
			return fr3.Eval(body)
		}

		// Case "yproc":
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
					if rs, ok := r.(string); ok {
						r = rs + "\n\tin yproc " + argv[0].String()
					}
					panic(r) // Rethrow errors and unknown Status.
				}
			}()
			fr3.Eval(body)
		}()

		return MkGenerator(ch)
	}

	existingNode := fr.G.Cmds[nameStr]
	node := &CmdNode{
		Fn: cmd,
		MixinLevel: fr.G.MixinNumberDefining,
		MixinName: fr.G.MixinNameDefining,
		Next: existingNode,
	}
	log.Printf("%s: NEW NODE %s: make %#v", argv[0], nameStr, node)
	fr.G.Cmds[nameStr] = node

	// Debug Dump
	node = fr.G.Cmds[nameStr]
	for node != nil {
		log.Printf("%s: NODE DUMP %s: %#v", argv[0], nameStr, node)
		node = node.Next
	}

	if captureMixinNumberDefining > 0 {
		// TODO: Check that long name is unique.
		newNode := &CmdNode{
			Fn: cmd,
			MixinLevel: captureMixinNumberDefining,
			MixinName: captureMixinNameDefining,
			Next: nil,
		}
		fr.G.Cmds[longMixinName] = newNode
	}

	return Empty
}

func cmdMixin(fr *Frame, argv []T) T {
	name, body := Arg2(argv)
	if fr.G.MixinNumberDefining > 0 {
		panic("already defining a mixin: " + fr.G.MixinNameDefining)
	}
	num := fr.G.MintMixinSerial()

	fr.G.MixinNumberDefining = num
	defer func() {
		fr.G.MixinNumberDefining = 0
	}()
	fr.G.MixinNameDefining = name.String()
	defer func() {
		fr.G.MixinNameDefining = ""
	}()

	return fr.Eval(body)
}

func cmdSuper(fr *Frame, argv []T) T {
	name, _ := Arg1v(argv)
    log.Printf("< Super < %s", Showv(argv))
	log.Printf("= Super = From mixin level %d", fr.MixinLevel)
	if fr.MixinLevel < 1 {
		panic("cannot super from non-mixin")
	}
	fn := fr.FindCommand(name, true)  // true: Call Super.
	z := fn(fr, argv[1:])
    log.Printf("> Super > %s", Show(z))
	return z
}

func cmdYield(fr *Frame, argv []T) T {
	if len(argv) == 2 {
		// Write exactly 1 arg on the channel.
		fr.Chan <- argv[1]
		return argv[1]
	}

	// Write more than 1 arg in a list.
	z := MkList(argv[1:])
	fr.Chan <- z
	return z
}

func cmdLs(fr *Frame, argv []T) T {
	panic("not usefully implemented yet")
}

func cmdSLen(fr *Frame, argv []T) T {
	a := Arg1(argv)
	return MkInt(int64(len(a.String())))
}

func cmdLLen(fr *Frame, argv []T) T {
	a := Arg1(argv)
	return MkInt(int64(len(a.List())))
}

func cmdList(fr *Frame, argv []T) T {
	return MkList(argv[1:])
}

func cmdLAt(fr *Frame, argv []T) T {
	tlist, ti := Arg2(argv)
	list := tlist.List()
	i := ti.Int()
	if i < 0 || i > int64(len(list)) {
		panic(Sprintf("lat: bad index: len(list)=%d but i=%d", len(list), i))
	}
	return list[i]
}

func cmdSAt(fr *Frame, argv []T) T {
	s, j := Arg2(argv)
	i := j.Int()
	return MkString(s.String()[i : i+1])
}

func cmdHttpHandler(fr *Frame, argv []T) T {
	fn := func(w http.ResponseWriter, r *http.Request) {
		v := make([]T, len(argv)-1)
		copy(v, argv[1:])
		v = append(v, MkT(w))
		v = append(v, MkT(r))
		_ = fr.Apply(v)
	}
	return MkT(fn)
}

func cmdForEach(fr *Frame, argv []T) T {
	v, list, body := Arg3(argv)

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

func cmdWhile(fr *Frame, argv []T) T {
	cond, body := Arg2(argv)

	toBreak := false
	toContinue := false

	for {
		c := fr.EvalExpr(cond)
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

func cmdCatch(fr *Frame, argv []T) (status T) {
	body, varT := Arg2(argv)
	varName := varT.String()

	defer func() {
		if r := recover(); r != nil {
			if j, ok := r.(Jump); ok {
				fr.SetVar(varName, j.Result)
				status = MkInt(int64(j.Status))
				return
			}
			fr.SetVar(varName, MkT(r))
			status = MkInt(1)
		}
	}()

	z := fr.Eval(body)
	fr.SetVar(varName, z)
	return MkInt(0)
}

func cmdEval(fr *Frame, argv []T) T {
	return EvalOrApplyLists(fr, argv[1:])
}

func cmdUplevel(fr *Frame, argv []T) T {
	specArg, rest := Arg1v(argv)
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
	return fr.Eval(MkString(buf.String()))
}

func ConcatLists(lists []T) []T {
	z := make([]T, 0, 4)
	for _, e := range lists {
		z = append(z, e.List()...)
	}
	return z
}

func cmdConcat(fr *Frame, argv []T) T {
	return MkList(ConcatLists(argv[1:]))
}

func cmdUpVar(fr *Frame, argv []T) T {
	lev, rem, loc := Arg3(argv)
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

func cmdSet(fr *Frame, argv []T) T {
	if len(argv) == 2 {
		// Retrieve value of variable, if 2nd arg is missing.
		name := Arg1(argv)
		return fr.GetVar(name.String())
	}
	name, x := Arg2(argv)
	fr.SetVar(name.String(), x)
	return x
}

func cmdReturn(fr *Frame, argv []T) T {
	var z T = Empty
	if len(argv) == 2 {
		z = argv[1]
	}
	if len(argv) > 2 {
		z = MkList(argv[1:])
	}
	// Jump with status RETURN.
	panic(Jump{Status: RETURN, Result: z})
}

func cmdBreak(fr *Frame, argv []T) T {
	panic(Jump{Status: BREAK}) // Jump with status BREAK.
}

func cmdContinue(fr *Frame, argv []T) T {
	panic(Jump{Status: CONTINUE}) // Jump with status CONTINUE.
}

func cmdHash(fr *Frame, argv []T) T {
	return MkHash()
}

func cmdHGet(fr *Frame, argv []T) T {
	hash, key := Arg2(argv)
	h := hash.Hash()
	k := key.String()
	value := h[k]
	if value == nil {
		panic(Sprintf("Hash does not contain key: %q", k))
	}
	return value
}

func cmdHSet(fr *Frame, argv []T) T {
	hash, key, value := Arg3(argv)
	h := hash.Hash()
	k := key.String()
	h[k] = value
	return value
}

func cmdHDel(fr *Frame, argv []T) T {
	hash, key := Arg2(argv)
	h := hash.Hash()
	k := key.String()
	h[k] = nil // TODO: how to delete?
	return Empty
}

func cmdHKeys(fr *Frame, argv []T) T {
	hash := Arg1(argv)
	h := hash.Hash()
	z := make([]T, 0, len(h))
	for _, k := range SortedKeysOfHash(h) {
		z = append(z, MkString(k))
	}
	return MkList(z)
}

func cmdInterpNew(fr *Frame, argv []T) T {
	name := Arg1(argv)
	nameStr := name.String()

	if _, ok := fr.G.SubTerps[nameStr]; ok {
		panic("Subterp already exists.")
	}

	// TODO - restrict our subinterp to only safe go functions.
	fr.G.SubTerps[nameStr] = New().G

	return Empty
}

func cmdInterpAlias(fr *Frame, argv []T) T {
	name, newcmdname, prefix := Arg3(argv)
	nameStr := name.String()
	newcmdnameStr := newcmdname.String()

	var sub *Global
	var ok bool
	if sub, ok = fr.G.SubTerps[nameStr]; !ok {
		panic("Subterp does not exist.")
	}

	cmd := func(fr2 *Frame, argv2 []T) (result T) {
		defer func() {
			if r := recover(); r != nil {
				if j, ok := r.(Jump); ok {
					switch j.Status {
					case RETURN:
						r ="return reached in an interp-alias"
					case BREAK:
						r ="break command was not inside a loop"
					case CONTINUE:
						r ="continue command was not inside a loop"
					}
				}
				if rs, ok := r.(string); ok {
					r = rs + "\n\tin alias " + argv2[0].String()
				}
				panic(r) // Rethrow errors and unknown Status.
			}
		}()

		if argv2 == nil {
			// Debug Data, if invoked with nil argv2.
			return MkList(argv)
		}

		z := make([]T, 0, 4)
		z = append(z, prefix.List()...)
		z = append(z, argv2[1:]...)

		return fr.Apply(z)
	}

	if _, ok := sub.Cmds[newcmdnameStr]; ok {
		panic("The command already exists within that subinterp.")
	}

	node := &CmdNode{
		Fn: cmd,
	}
	log.Printf("%s: NEW Interp-Alias NODE %s: make %#v", argv[0], newcmdnameStr, node)
	sub.Cmds[newcmdnameStr] = node

	return Empty
}


func cmdInterpEval(fr *Frame, argv []T) T {
	name, script := Arg2(argv)
	nameStr := name.String()

	var sub *Global
	var ok bool
	if sub, ok = fr.G.SubTerps[nameStr]; !ok {
		panic("Subterp does not exist.")
	}

	return sub.Fr.Eval(script)
}
