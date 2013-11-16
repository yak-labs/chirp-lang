package chirp

import (
	"bytes"
	. "fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

// Safes are builtin commands that safe subinterps can call.
// Conventionally these contain no hyphen.
var Safes map[string]Command

// Unsafes are commands that only the trusted, toplevel terp can call.
// Conventionally these contain a hyphen.
var Unsafes map[string]Command

type binaryFlop func(a, b float64) float64
type binaryFlopBool func(a, b float64) bool
type binaryStringBool func(a, b string) bool

func MkBinaryFlopCmd(flop binaryFlop) Command {
	return func(fr *Frame, argv []T) T {
		a, b := Arg2(argv)
		return MkFloat(flop(a.Float(), b.Float()))
	}
}

func MkBinaryFlopBoolCmd(op binaryFlopBool) Command {
	return func(fr *Frame, argv []T) T {
		a, b := Arg2(argv)
		return MkBool(op(a.Float(), b.Float()))
	}
}

func MkBinaryStringBoolCmd(op binaryStringBool) Command {
	return func(fr *Frame, argv []T) T {
		a, b := Arg2(argv)
		return MkBool(op(a.String(), b.String()))
	}
}

func MkChainingBinaryFlopCmd(starter float64, flop binaryFlop) Command {
	return func(fr *Frame, argv []T) T {
		z := starter // Be sure not to modify starter!  It is captured.
		for _, a := range argv[1:] {
			z = flop(z, a.Float())
		}
		return MkFloat(z)
	}
}

func Arg0(argv []T) {
	if len(argv) != 1 {
		panic(Sprintf("Expected 0 arguments, but got argv=%s", Showv(argv)))
	}
}

func Arg0v(argv []T) []T {
	if len(argv) < 1 {
		panic(Sprintf("Expected at least 0 arguments, but got argv=%s", Showv(argv)))
	}
	return argv[1:]
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

func Arg4(argv []T) (T, T, T, T) {
	if len(argv) != 4+1 {
		panic(Sprintf("Expected 4 arguments, but got argv=%s", Showv(argv)))
	}
	return argv[1], argv[2], argv[3], argv[4]
}

func Arg5(argv []T) (T, T, T, T, T) {
	if len(argv) != 5+1 {
		panic(Sprintf("Expected 5 arguments, but got argv=%s", Showv(argv)))
	}
	return argv[1], argv[2], argv[3], argv[4], argv[5]
}

func Arg6(argv []T) (T, T, T, T, T, T) {
	if len(argv) != 6+1 {
		panic(Sprintf("Expected 6 arguments, but got argv=%s", Showv(argv)))
	}
	return argv[1], argv[2], argv[3], argv[4], argv[5], argv[6]
}

func Arg7(argv []T) (T, T, T, T, T, T, T) {
	if len(argv) != 7+1 {
		panic(Sprintf("Expected 7 arguments, but got argv=%s", Showv(argv)))
	}
	return argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7]
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

	if fr.EvalExpr(cond).Bool() {
		return fr.Eval(yes)
	}

	if no != nil {
		return fr.Eval(no)
	}

	return Empty
}

func cmdCase(fr *Frame, argv []T) T {
	// Two possible syntaxes for Tcl 6.7 case command:
	//   (1) case string ?in? patList body ?patList body ...?
	//   (2) case string ?in? {patList body ?patList body ...?}
	topicL, rest := Arg1v(argv)
	topic := topicL.String()

	if len(rest) < 1 {
		panic(Sprintf("Too few arguments for 'case': %#v", argv))
	}
	if rest[0].String() == "in" {
		// ?in? exists; delete it.
		rest = rest[1:]
	}

	if len(rest) == 1 {
		// Case (2).  Expand the one arg into its parts.
		rest = rest[0].List()
	}

	if (len(rest) & 1) == 1 {
		panic(Sprintf("Odd number of items in {patList body} list of stride two: %v", argv))
	}

	var dflt T
	for i := 0; i < len(rest); i += 2 {
		pats := rest[i].List()
		if len(pats) == 1 && pats[0].String() == "default" {
			dflt = rest[i+1]
			continue
		}
		for _, pat := range pats {
			if StringMatch(pat.String(), topic) {
				return fr.Eval(rest[i+1])
			}
		}
	}

	if dflt == nil {
		return Empty
	}
	return fr.Eval(dflt)
}

func cmdPuts(fr *Frame, argv []T) T {
	// TODO:  accept a Writer as first arg.
	out := Arg1(argv)
	Println(out)
	return Empty
}

func cmdProc(fr *Frame, argv []T) T {
	return procOrYProc(fr, argv, false, nil)
}

func cmdYProc(fr *Frame, argv []T) T {
	return procOrYProc(fr, argv, true, nil)
}

func procOrYProc(fr *Frame, argv []T, generating bool, super *Obj) T {
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
	n := len(alist)

	// Capture this variable, so it can be used when the cmd is called.
	captureMixinNumberDefining := fr.G.MixinNumberDefining
	captureMixinNameDefining := fr.G.MixinNameDefining

	// If the proc is being defined by a mixin, put it in the same mixin.
	if fr.MixinLevel > 0 {
		captureMixinNumberDefining = fr.MixinLevel
		captureMixinNameDefining = fr.MixinName
	}
	var longMixinName string
	if captureMixinNumberDefining > 0 {
		longMixinName = captureMixinNameDefining + "~" + nameStr
	}

	var compiled Evaler
	if !body.IsPreservedByList() { // TODO: reconsider this test.
		compiled = CompileSequence(fr, body.String())
	}

	cmd := func(fr2 *Frame, argv2 []T) (result T) {
		var self *Obj
		if super != nil {
			// Argv2 is [ Receiver, Message, args... ].
			// Remove Receiver, defining self.
			// Leave message to be argv[0] for the command.
			self = argv2[0].Raw().(*Obj)
			argv2 = argv2[1:]
		}
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
							r = ("break command was not inside a loop")
						case CONTINUE:
							r = ("continue command was not inside a loop")
						}
					}
					if rs, ok := r.(string); ok {
						rs = rs + "\n\tin proc " + argv2[0].String()
						rs = rs + Sprintf("\n\t\t(caller's MixinLevel=%d)", fr2.MixinLevel)
						rs = rs + Sprintf("\n\t\t(caller's MixinName=%q)", fr2.MixinName)
						// TODO: Require debug level for the args.
						for ai, ae := range argv2[1:] {
							as := ae.String()
							if len(as) > 80 {
								as = as[:80] + "..."
							}
							rs = rs + Sprintf("\n\t\targ:%d = %q", ai, as)
						}
						// TODO: Require debug level for the locals.
						for vk, vv := range fr2.Vars {
							vs := vv.Get().String()
							if len(vs) > 80 {
								vs = vs[:80] + "..."
							}
							rs = rs + Sprintf("\n\t\tlocal:%s = %q", vk, vs)
						}
						r = rs
					}
					panic(r) // Rethrow errors and unknown Status.
				}
			}()
		}

		if argv2 == nil {
			// Debug Data, if invoked with nil argv2.
			return MkList(argv)
		}

		var varargs bool = false
		if len(astrs) > 0 && astrs[len(astrs)-1] == "args" {
			varargs = true
			if len(argv2) < n {
				panic(Sprintf("%s %q expects arguments %#v but got %#v", argv[0], nameStr, aa, argv2))
			}
		} else {
			if len(argv2) != n+1 {
				panic(Sprintf("%s %q expects arguments %#v but got %#v", argv[0], nameStr, aa, argv2))
			}
		}

		fr3 := fr2.NewFrame()
		fr3.MixinLevel = captureMixinNumberDefining
		fr3.MixinName = captureMixinNameDefining
		fr3.Self = self
		fr3.Super = super

		if varargs {
			for i, arg := range astrs[:len(astrs)-1] {
				fr3.SetVar(arg, argv2[i+1])
			}

			fr3.SetVar("args", MkList(argv2[len(astrs):]))
		} else {
			for i, arg := range astrs {
				fr3.SetVar(arg, argv2[i+1])
			}
		}

		// Case "proc":
		if !generating {
			if compiled != nil {
				return compiled.Eval(fr3)
			}
			return fr3.Eval(body)
		}

		// Case "yproc":
		ch := make(chan Either, 0)
		z := MkGenerator(ch) // Save reader half in z.
		fr3.WriterChan = ch  // Save writer half in frame.
		ch = nil

		go func() {
			defer close(fr3.WriterChan)
			defer func() {
				if r := recover(); r != nil {
					var ei Either
					ei.Bad = Sprintf("%v", r)
					if j, ok := r.(Jump); ok {
						switch j.Status {
						case RETURN:
							if !j.Result.IsEmpty() {
								ei.Bad = "yproc command: cannot return a value"
							}
							return
						case BREAK:
							ei.Bad = "yproc command: break command was not inside a loop"
						case CONTINUE:
							ei.Bad = "yproc command: continue command was not inside a loop"
						default:
							ei.Bad = Sprintf("yproc command: unknown Jump Status: %d", int(j.Status))
						}
					}
					if rs, ok := r.(string); ok {
						r = rs + "\n\tin yproc " + argv[0].String()
					}
					fr3.WriterChan <- ei
				}
			}()
			fr3.Eval(body)
		}()

		return z
	}

	builtin := Safes[nameStr]
	if builtin != nil {
		panic(Sprintf("cannot redefine a builtin: %q", nameStr))
	}

	existingNode := fr.G.Cmds[nameStr]

	if IsGlobal(nameStr) {
		// Procs that behave as Mixins have Capital Initial Letter.

		node := &CmdNode{
			Fn:         cmd,
			MixinLevel: fr.G.MixinNumberDefining,
			MixinName:  fr.G.MixinNameDefining,
			Next:       existingNode,
		}
		fr.G.Cmds[nameStr] = node

		// Debug Dump
		node = fr.G.Cmds[nameStr]
		for node != nil {
			node = node.Next
		}
	} else {
		if existingNode != nil {
			panic(Sprintf("Name already defined at base level; cannot redefine: %q", nameStr))
		}
		if captureMixinNumberDefining == 0 {
			// Install base command.
			node := &CmdNode{
				Fn:         cmd,
				MixinLevel: fr.G.MixinNumberDefining,
				MixinName:  fr.G.MixinNameDefining,
				Next:       nil,
			}
			fr.G.Cmds[nameStr] = node
		} else {
			// Install as Long Name below.
		}
	}

	if captureMixinNumberDefining > 0 {
		// TODO: Check that long name is unique.
		newNode := &CmdNode{
			Fn:         cmd,
			MixinLevel: captureMixinNumberDefining,
			MixinName:  captureMixinNameDefining,
			Next:       nil,
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
	if fr.MixinLevel < 1 {
		panic("cannot super from non-mixin")
	}
	fn := fr.FindCommand(name, true) // true: Call Super.
	z := fn(fr, argv[1:])
	return z
}

func cmdYield(fr *Frame, argv []T) T {
	if len(argv) == 2 {
		// Write exactly 1 arg on the channel.
		fr.WriterChan <- Either{Good: argv[1]}
		return argv[1]
	}

	// Write more than 1 arg in a list.
	z := MkList(argv[1:])
	fr.WriterChan <- Either{Good: z}
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

func cmdNull(fr *Frame, argv []T) T {
	a := Arg1(argv)
	return MkBool(a.IsEmpty())
}

func cmdNotNull(fr *Frame, argv []T) T {
	a := Arg1(argv)
	return MkBool(!a.IsEmpty())
}

func cmdList(fr *Frame, argv []T) T {
	return MkList(argv[1:])
}

func cmdLIndex(fr *Frame, argv []T) T {
	tlist, ti := Arg2(argv)
	list := tlist.List()
	i := ti.Int()
	if i < 0 || i > int64(len(list)) {
		panic(Sprintf("lindex: bad index: len(list)=%d but i=%d", len(list), i))
	}
	return list[i]
}

func cmdLSort(fr *Frame, argv []T) T {
	tlist := Arg1(argv)
	list := tlist.List()
	n := len(list)

	// Consider a list with 1 or less elements already sorted.
	if n <= 1 {
		return tlist
	}

	strs := make([]string, n)

	// Convert our list to a slice of strings.
	for i, t := range list {
		strs[i] = t.String()
	}

	// Sort our strings.
	sort.Strings(strs)

	// Put our sorted strings back into the list.
	for i, s := range strs {
		list[i] = MkString(s)
	}

	// Return the sorted list.
	return MkList(list)
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
	varLT, list, body := Arg3(argv)
	varL := varLT.List()

	toBreak := false
	toContinue := false

Outer:
	for {
		var hd T
		var tl T
		for _, varT := range varL {
			hd, tl = list.HeadTail()
			if hd == nil {
				// This does leave vars in a slightly skewed state if the stride of varL
				// doesn't fit the data.  So just don't mismatch the lengths.
				// It's not worth the complexity to fix this code.
				break Outer
			}
			list = tl

			fr.SetVar(varT.String(), hd)
		}

		func() {
			defer func() {
				if r := recover(); r != nil {
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
			fr.Eval(body)
		}()
		if toBreak {
			break
		}
		if toContinue {
			continue
		}
	}

	return Empty
}

func cmdWhile(fr *Frame, argv []T) T {
	cond, body := Arg2(argv)

	toBreak := false
	toContinue := false

	for {
		c := fr.EvalExpr(cond)
		if !c.Bool() {
			break
		}

		func() {
			defer func() {
				if r := recover(); r != nil {
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
			fr.Eval(body)
		}()
		if toBreak {
			break
		}
		if toContinue {
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
			status = True
		}
	}()

	z := fr.Eval(body)
	fr.SetVar(varName, z)
	return False
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
	fr.DefineUpVar(locName, remFr, remName)
	return Empty
}

func cmdSet(fr *Frame, argv []T) T {
	target, _ := Arg1v(argv)
	targ := target.String()
	if len(targ) == 0 {
		panic("command 'set' target is empty")
	}
	n := len(targ)
	if targ[n-1] == ')' {
		// Case Subscript:
		i := strings.Index(targ, "(")
		if i < 0 {
			panic("command 'set' target ends with ')' but has no '('")
		}
		if i < 1 {
			panic("command 'set' target is empty before '('")
		}

		name := targ[:i]
		key := targ[i+1 : n-1]
		if len(argv) == 2 {
			h := fr.GetVar(name)
			return h.GetAt(MkString(key))
		}
		if !fr.HasVar(name) {
			fr.SetVar(name, MkHash(nil))
		}
		_, x := Arg2(argv)
		h := fr.GetVar(name)
		h.PutAt(x, MkString(key))
		return x
	}

	// Case No Subscript:
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
	return MkHash(nil)
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

type SafeSubInterp struct {
	fr *Frame // Private member.
}

func cmdInterp(fr *Frame, argv []T) T {
	Arg0(argv)

	z := &SafeSubInterp{
		fr: NewSafe(),
	}
	return MkT(z)
}

func (ssi *SafeSubInterp) Alias(fr *Frame, newcmdnameStr string, prefix T) {
	cmd := func(fr2 *Frame, argv2 []T) (result T) {
		defer func() {
			if r := recover(); r != nil {
				if j, ok := r.(Jump); ok {
					switch j.Status {
					case RETURN:
						r = "return reached in an interp-alias"
					case BREAK:
						r = "break command was not inside a loop"
					case CONTINUE:
						r = "continue command was not inside a loop"
					}
				}
				if rs, ok := r.(string); ok {
					r = rs + "\n\tin alias " + argv2[0].String()
				}
				panic(r) // Rethrow errors and unknown Status.
			}
		}()

		z := make([]T, 0, 4)
		z = append(z, prefix.List()...)
		z = append(z, argv2[1:]...)

		fr3 := fr2.NewFrame()
		fr3.G = fr.G
		return fr3.Apply(z)
	}

	if _, ok := ssi.fr.G.Cmds[newcmdnameStr]; ok {
		panic("The command already exists within that subinterp.")
	}

	node := &CmdNode{
		Fn: cmd,
	}
	ssi.fr.G.Cmds[newcmdnameStr] = node
}

func (ssi *SafeSubInterp) Eval(script T) T {
	return ssi.fr.Eval(script)
}

func (ssi *SafeSubInterp) Clone() *SafeSubInterp {
	cloned := ssi.fr.G.Clone()
	z := &SafeSubInterp{
		fr: &cloned.Fr,
	}
	return z
}

func (ssi *SafeSubInterp) CopyCredFrom(fr *Frame) {
	if fr.Cred != nil {
		if ssi.fr.Cred == nil {
			ssi.fr.Cred = make(Hash)
		}
		for k, v := range fr.Cred {
			ssi.fr.Cred[k] = v
		}
	}
}

// Tcl requires integers, but our base numeric value is float64.
// TODO: Make the increment argument optional.
func cmdIncr(fr *Frame, argv []T) T {
	varName, incr := Arg2(argv)

	name := varName.String()

	if !fr.HasVar(name) {
		fr.SetVar(name, MkInt(0))
	}
	v := fr.GetVar(name).Float()
	i := incr.Float()
	z := MkFloat(v + i)

	fr.SetVar(name, z)

	return z
}

func cmdAppend(fr *Frame, argv []T) T {
	varName, values := Arg1v(argv)

	name := varName.String()

	if !fr.HasVar(name) {
		fr.SetVar(name, Empty)
	}
	v := fr.GetVar(name)

	i := 0
	n := len(values)

	if n == 0 {
		// We get to return early.
		return v
	}

	buf := bytes.NewBufferString(v.String())

	for i < n {
		buf.WriteString(values[i].String())
		i++
	}

	z := MkString(buf.String())
	fr.SetVar(name, z)
	return z
}

func cmdLAppend(fr *Frame, argv []T) T {
	varName, values := Arg1v(argv)

	name := varName.String()

	if !fr.HasVar(name) {
		fr.SetVar(name, Empty)
	}
	v := fr.GetVar(name).List()
	v = append(v, values...)

	z := MkList(v)
	fr.SetVar(name, z)
	return z
}

func cmdError(fr *Frame, argv []T) T {
	message := Arg1(argv)

	panic(message.String())
}

// Modern Tcl uses "return --code" to throw strange codes,
// but our "return" takes multiple values, so we cannot use it.
// Tcl 6.7 had no way to do it.
// We add a command "throw code result" to do it.
func cmdThrow(fr *Frame, argv []T) T {
	statusT, resultT := Arg2(argv)
	status := statusT.Int()

	panic(Jump{
		Status: StatusCode(status),
		Result: resultT,
	})
}

var stringEnsemble = []EnsembleItem{
	EnsembleItem{Name: "length", Cmd: cmdSLen},
	EnsembleItem{Name: "range", Cmd: cmdStringRange},
	EnsembleItem{Name: "slice", Cmd: cmdStringSlice},
	EnsembleItem{Name: "first", Cmd: cmdStringFirst},
	EnsembleItem{Name: "index", Cmd: cmdStringIndex},
	EnsembleItem{Name: "match", Cmd: cmdStringMatch},
}

// Follows Tcl's string range spec.
func cmdStringRange(fr *Frame, argv []T) T {
	str, first, last := Arg3(argv)

	strS := str.String()
	n := len(strS)
	firstI := int(first.Int()) // The index of the first character to include.

	keep := 1     // Tcl's string range includes the character indexed by last
	var lastI int // The index of the last character to include.
	if last.IsEmpty() || last.String() == "end" {
		lastI = n - keep
	} else {
		lastI = int(last.Int())
	}

	low, high, ok := slicer(n, firstI, lastI, keep)
	if !ok {
		return Empty
	}

	return MkString(strS[low:high])
}

// Follows golang's slice spec.
func cmdStringSlice(fr *Frame, argv []T) T {
	str, first, last := Arg3(argv)

	strS := str.String()
	n := len(strS)
	firstI := int(first.Int()) // The index of the first character to include.

	var lastI int // The number characters to include.
	if last.IsEmpty() || last.String() == "end" {
		lastI = n
	} else {
		lastI = int(last.Int())
	}

	low, high, ok := slicer(n, firstI, lastI, 0)
	if !ok {
		return Empty
	}

	return MkString(strS[low:high])
}

// Slicer will find the low and high values for slicing a golang slice.
// http://golang.org/ref/spec#Slices
//
// Parameters:
// length - The length of the slice.
// first  - The index of the first element to take.
// last   - The high value for the slice.
//          If keep is 0, this will return a low/high value that will satisfy
//          0 <= low <= high <= length, like in go.
// keep   - The number of elements to keep.
//
// Returns:
// low    - The low value for the slice.
// high   - The high value for the slice.
// ok     - false if there is an invalid request.
func slicer(length int, first, last int, keep int) (int, int, bool) {
	// If first is too small, Zero.
	if first < 0 {
		first = 0
	}

	// If first is too large, Empty.
	if first > length {
		return -1, -1, false
	}

	// Last may be negative, like in Python.
	if last < 0 {
		last += length - keep
	}

	// If last is too small, Empty.
	if last < first {
		return -1, -1, false
	}

	// If last is too large, End.
	if last > length-keep {
		last = length - keep
	}

	return first, last + keep, true
}

// TODO: Add optional argument "startIndex"
func cmdStringFirst(fr *Frame, argv []T) T {
	needle, haystack := Arg2(argv)

	i := strings.Index(haystack.String(), needle.String())

	return MkFloat(float64(i))
}

func cmdStringIndex(fr *Frame, argv []T) T {
	str, charIndex := Arg2(argv)

	s := str.String()
	i := int(charIndex.Int())
	n := len(s)

	if i < 0 || i >= n {
		return Empty
	}

	z := string(s[i])
	return MkString(z)
}

func cmdStringMatch(fr *Frame, argv []T) T {
	pattern, str := Arg2(argv)

	return MkBool(StringMatch(pattern.String(), str.String()))
}

func StringMatch(pattern, str string) bool {
	plen, slen := len(pattern), len(str)
	pidx, cidx := 0, 0
	var p, c uint8

Loop:
	for pidx < plen {
		p = pattern[pidx]

		// c is unset.
		if p == '*' {
			// Skip successive *'s in the pattern
			for p == '*' {
				pidx++
				if pidx < plen {
					p = pattern[pidx]
				} else {
					return true
				}
			}

			// Loop through the string until satisfied.
			// p is the pattern after the * we found.
			// pidx != plen
			for cidx < slen {
				// Optimization:
				// If 'p' isn't a special character, look ahead for the next matching
				// character in the string.
				if p != '[' && p != '?' && p != '\\' {

					// c is the next character to try and match.
				StarLookAhead:
					for cidx < slen {
						c = str[cidx]

						if c == p {
							break StarLookAhead
						}

						cidx++

						// We reached the end of str so we can return early.
						if cidx == slen {
							return false
						}
					}
					// c should now be the first character that matches p
					// cidx should be the index of c in str
				}

				if StringMatch(pattern[pidx:], str[cidx:]) {
					return true
				}

				cidx++
			}
			// reached end of str
			// p is unmatched
			return false
		}

		if p == '?' {
			pidx++
			cidx++
			continue Loop
		}

		// Populate c if we can.
		if cidx < slen {
			c = str[cidx]
		} else {
			// We've run out of string.
			return false
		}

		if p == '[' {
			var start, end uint8

			// Skip the pidx to point to the next char
			pidx++

		BracketLoop:
			for {
				if pidx == plen {
					return false
				}

				p = pattern[pidx]
				if p == ']' {
					return false
				}

				start = p

				pidx++
				p = pattern[pidx]

				if p == '-' {
					// Match a range of characters.
					pidx++
					if pidx == plen {
						return false
					}

					p = pattern[pidx]
					end = p

					if (start <= c && c <= end) || (end <= c && c <= start) {
						break BracketLoop
					}
				} else if start == c {
					break BracketLoop
				}
			}

			// Skip to after the ending bracket.
			for p != ']' {
				pidx++
				if pidx < plen {
					p = pattern[pidx]
				} else {
					p--
					break
				}
			}

			// We succeeded in matching our character.  Continue the loop.
			pidx++
			cidx++
			continue Loop
		}

		// Strip off the '\' so we do an exact match on the following char.
		if p == '\\' {
			pidx++
			if pidx == plen {
				return false
			}

			p = pattern[pidx]
		}

		// The normal case, with no special characters.
		if c != p {
			return false
		}

		pidx++
		cidx++
	}

	// Are we at the end of both the pattern and the string?
	if pidx == plen {
		return cidx == slen
	}

	return false
}

var infoEnsemble = []EnsembleItem{
	EnsembleItem{Name: "commands", Cmd: cmdInfoCommands},
	EnsembleItem{Name: "globals", Cmd: cmdInfoGlobals},
	EnsembleItem{Name: "locals", Cmd: cmdInfoLocals},
}

func cmdInfoCommands(fr *Frame, argv []T) T {
	Arg0(argv) // TODO: optional pattern
	z := make([]T, 0, 100)

	for k, _ := range fr.G.Cmds {
		z = append(z, MkString(k))
	}
	return MkList(z)
}
func cmdInfoGlobals(fr *Frame, argv []T) T {
	Arg0(argv) // TODO: optional pattern
	z := make([]T, 0, 100)

	for k, _ := range fr.G.Fr.Vars {
		z = append(z, MkString(k))
	}
	return MkList(z)
}
func cmdInfoLocals(fr *Frame, argv []T) T {
	Arg0(argv) // TODO: optional pattern
	z := make([]T, 0, 100)

	for k, _ := range fr.Vars {
		z = append(z, MkString(k))
	}
	return MkList(z)
}

func cmdSplit(fr *Frame, argv []T) T {
	str, delimV := Arg1v(argv)
	s := str.String()
	if s == "" {
		return Empty // Special case in Tcl.
	}

	var delim string
	switch len(delimV) {
	case 0:
		delim = ""
	case 1:
		delim = delimV[0].String()
	default:
		panic("Usage: split str ?delims?")
	}
	if delim == "" {
		delim = " \t\n\r" // White Space.
	}

	z := make([]T, 0, 4)
	for {
		i := strings.IndexAny(s, delim)
		if i == -1 {
			z = append(z, MkString(s))
			break
		}
		z = append(z, MkString(s[:i]))
		s = s[i+1:]
	}
	return MkList(z)
}

func cmdJoin(fr *Frame, argv []T) T {
	list, joinV := Arg1v(argv)

	var joiner string
	switch len(joinV) {
	case 0:
		joiner = " "
	case 1:
		joiner = joinV[0].String()
	default:
		panic("Usage: join list ?joinString?")
	}

	buf := bytes.NewBuffer(nil)
	for i, e := range list.List() {
		if i > 0 {
			buf.WriteString(joiner)
		}
		buf.WriteString(e.String())
	}
	return MkString(buf.String())
}

func cmdDropNull(fr *Frame, argv []T) T {
	listT := Arg1(argv)
	list := listT.List()
	z := make([]T, 0, len(list))
	for _, e := range list {
		if !e.IsEmpty() {
			z = append(z, e)
		}
	}
	return MkList(z)
}

func cmdSubst(fr *Frame, argv []T) T {
	args := Arg0v(argv)

	if len(args) == 0 {
		panic("'subst' commmand needs an argument")
	}

	var flags SubstFlags
	for len(args) > 1 {
		a := args[0].String()
		switch true {
		case StringMatch("-nob*", a):
			flags |= NoBackslash
		case StringMatch("-noc*", a):
			flags |= NoSquare
		case StringMatch("-nov*", a):
			flags |= NoDollar
		default:
			panic(Sprintf("Bad flag for 'subst' commmand: %q", a))
		}
		args = args[1:]
	}

	return MkString(fr.SubstString(args[0].String(), flags))
}

// Getting cred is safe.  Setting it is unsafe.  This is the getter.
// Returns Empty if none.
func cmdCred(fr *Frame, argv []T) T {
	name := Arg1(argv)

	key := name.String()
	if _, ok := fr.Cred[key]; !ok {
		return Empty
	}
	z := fr.Cred[key]
	if z == nil {
		return Empty
	}
	return z
}

// Usage: log <level> <messages>...
// Creates a new stderr logger, if Global has no logger yet.
func cmdLog(fr *Frame, argv []T) T {
	levelT, messageT := Arg2(argv)
	Log(fr, levelT.String(), messageT.String())
	return Empty
}

func Log(fr *Frame, levelStr string, message string) {
	var panicky, fatally bool
	if len(levelStr) != 1 {
		panic(Sprintf("Log level should be 'p', 'f', or in '0'..'9' but is %q", levelStr))
	}
	lev := levelStr[0]
	level := -1 // for case 'p' or 'f'

	if lev == 'p' { // "p"anic level
		panicky = true
	} else if lev == 'f' { // "f"atal level
		fatally = true
	} else if '0' <= lev && lev <= '9' {
		level = int(lev) - int('0')
	} else {
		panic(Sprintf("Log level should be 'p', 'f', or in '0'..'9' but is %q", level))
	}

	if level > fr.G.Verbosity {
		return // Not enough verbosity for this message.
	}

	if fr.G.Logger == nil {
		logName := fr.G.LogName
		if logName == "" {
			logName = "chirp" // Default LogName
		}
		fr.G.Logger = log.New(os.Stderr, logName, log.LstdFlags)
	}

	message = SubstStringOrOrig(fr, message)
	fr.G.Logger.Println(message)

	if panicky {
		panic(Sprintf("log p: %s", message))
	}
	if fatally {
		fr.G.Logger.Println("Exiting after fatal log message.")
		os.Exit(13) // Unlucky Exit.
	}
}

func SubstStringOrOrig(fr *Frame, s string) (z string) {
	defer func() {
		if r := recover(); r != nil {
			z = Sprintf("ERROR ignored while substituting log message: %s", s)
			return
		}
	}()
	return fr.SubstString((s), 0) // 0 is all substitutions.
}

func init() {
	if Safes == nil {
		Safes = make(map[string]Command, 333)
	}
	Safes["+"] = MkChainingBinaryFlopCmd(0.0, func(a, b float64) float64 { return a + b })
	Safes["*"] = MkChainingBinaryFlopCmd(1.0, func(a, b float64) float64 { return a * b })
	Safes["-"] = MkBinaryFlopCmd(func(a, b float64) float64 { return a - b })
	Safes["/"] = MkBinaryFlopCmd(func(a, b float64) float64 { return a / b })

	Safes["=="] = MkBinaryFlopBoolCmd(func(a, b float64) bool { return (a == b) })
	Safes["!="] = MkBinaryFlopBoolCmd(func(a, b float64) bool { return (a != b) })
	Safes["<"] = MkBinaryFlopBoolCmd(func(a, b float64) bool { return (a < b) })
	Safes["<="] = MkBinaryFlopBoolCmd(func(a, b float64) bool { return (a <= b) })
	Safes[">"] = MkBinaryFlopBoolCmd(func(a, b float64) bool { return (a > b) })
	Safes[">="] = MkBinaryFlopBoolCmd(func(a, b float64) bool { return (a >= b) })

	Safes["eq"] = MkBinaryStringBoolCmd(func(a, b string) bool { return (a == b) })
	Safes["ne"] = MkBinaryStringBoolCmd(func(a, b string) bool { return (a != b) })
	Safes["lt"] = MkBinaryStringBoolCmd(func(a, b string) bool { return (a < b) })
	Safes["le"] = MkBinaryStringBoolCmd(func(a, b string) bool { return (a <= b) })
	Safes["gt"] = MkBinaryStringBoolCmd(func(a, b string) bool { return (a > b) })
	Safes["ge"] = MkBinaryStringBoolCmd(func(a, b string) bool { return (a >= b) })

	Safes["must"] = cmdMust
	Safes["if"] = cmdIf
	Safes["case"] = cmdCase
	Safes["puts"] = cmdPuts
	Safes["proc"] = cmdProc
	Safes["yproc"] = cmdYProc
	Safes["mixin"] = cmdMixin
	Safes["super"] = cmdSuper
	Safes["yield"] = cmdYield
	Safes["ls"] = cmdLs
	Safes["null"] = cmdNull
	Safes["notnull"] = cmdNotNull
	Safes["list"] = cmdList
	Safes["lindex"] = cmdLIndex
	Safes["lsort"] = cmdLSort
	Safes["llength"] = cmdLLen
	Safes["http_handler"] = cmdHttpHandler
	Safes["foreach"] = cmdForEach
	Safes["while"] = cmdWhile
	Safes["catch"] = cmdCatch
	Safes["eval"] = cmdEval
	Safes["uplevel"] = cmdUplevel
	Safes["concat"] = cmdConcat
	Safes["set"] = cmdSet
	Safes["upvar"] = cmdUpVar
	Safes["return"] = cmdReturn
	Safes["break"] = cmdBreak
	Safes["continue"] = cmdContinue
	Safes["hash"] = cmdHash
	Safes["hget"] = cmdHGet   // FIXME: temporary: Use getf
	Safes["hset"] = cmdHSet   // FIXME: temporary: Use setf
	Safes["hdel"] = cmdHDel   // FIXME: temporary: Use delf
	Safes["hkeys"] = cmdHKeys // FIXME: temporary: use keys
	Safes["interp"] = cmdInterp
	Safes["incr"] = cmdIncr
	Safes["append"] = cmdAppend
	Safes["lappend"] = cmdLAppend
	Safes["error"] = cmdError
	Safes["throw"] = cmdThrow
	Safes["string"] = MkEnsemble(stringEnsemble)
	Safes["info"] = MkEnsemble(infoEnsemble)
	Safes["split"] = cmdSplit
	Safes["join"] = cmdJoin
	Safes["dropnull"] = cmdDropNull
	Safes["subst"] = cmdSubst
	Safes["cred"] = cmdCred
	Safes["log"] = cmdLog
}
