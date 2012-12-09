package terp

import (
	. "fmt"
	"sync"
)

type Any interface{}

type List []Any
type Dict map[string]Any

type Command func(*Terp, []Any) Any

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
	return z
}

func (me *Terp) Eval(s string) string {
	v := ParseCmd(s)
	if len(v) < 1 {
		return ""
	}
	cmdName := Str(v[0])
	cmd, ok := me.Cmds[cmdName]
	if !ok {
	  cmd, ok = Builtins[cmdName]
	}
	if !ok {panic(Sprintf("cmd not found: %q", cmdName))}
	cmdArgs := v[1:]
	z := cmd(me, cmdArgs)
	return Repr(z)
}

func Repr(a Any) string { return Sprintf("%#v", a) }
func Str(a Any) string { return Sprintf("%v", a) }

func Must(a, b Any, extra ...Any) {
	if a != b {
		panic(Repr(a) + " vs. " + Repr(b) + " :: " + Repr(extra))
	}
}
