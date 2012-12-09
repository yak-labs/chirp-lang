package terp

import (
	. "fmt"
	"log"
	"sync"
)

type Any interface{}

type List []Any
type Dict map[string]Any

type Command func(*Terp, List) Any

type Scope map[string]Any

type CmdScope map[string]Command

type Frame struct {
	Vars Scope
	Prev *Frame
}

type Terp struct {
	Top  *Frame // current frame
	Mu   sync.Mutex
	Cmds CmdScope
	Fr   Frame // global scope
}

func NewTerp() *Terp {
	z := &Terp{
		Cmds: make(CmdScope),
		Fr: Frame{
			Vars: make(Scope),
		},
	}

	z.Top = &z.Fr
	z.initBuiltins()
	return z
}

func (me *Terp) Apply(head Any, args []Any) Any {
	log.Printf("< Apply < %#v < %#v\n", head, args)
	cmdName, ok := head.(string)
	if !ok {
		panic(Sprintf("Command must be a string: %#v", head))
	}

	fn, ok := me.Cmds[cmdName]
	if !ok {
		fn, ok = Builtins[cmdName]
	}
	if !ok {
		panic(Sprintf("Command not found: %q", cmdName))
	}
	z := fn(me, args)
	log.Printf("> Apply > %#v\n", z)
	return z
}

func Repr(a Any) string { return Sprintf("%#v", a) }
func Str(a Any) string {
	if s, ok := a.(string); ok {
		return s
	}
	return Sprintf("%v", a)
}

func Must(a, b Any, extra ...Any) {
	if a != b {
		panic(Repr(a) + " vs. " + Repr(b) + " :: " + Repr(extra))
	}
}
