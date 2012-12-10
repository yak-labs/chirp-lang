package terp

import (
	"go/ast"
	. "fmt"
	"log"
	"sync"
)

type Any interface{}

type List []Any
type Dict map[string]Any

type Command func(fr *Frame, argv List) Any

type Scope map[string]Any

type CmdScope map[string]Command

type Frame struct {
	Vars *Scope
	Prev *Frame
	G *Global
	Slots *Scope
}

type Global struct {
	Top  *Frame // current frame
	Mu   sync.Mutex
	Cmds CmdScope
	Fr   Frame // global scope
}

func New() *Frame {
	scope := make(Scope)
	g := &Global{
		Cmds: make(CmdScope),
		Fr: Frame{
			Vars: &scope,
		},
	}

	g.Fr.G = g
	g.Top = &g.Fr
	g.Top.initBuiltins()
	return g.Top
}

func IsGlobal(name string) bool {
	return ast.IsExported(name)  // Same criteria, First is Uppercase.
}

func (fr *Frame) GetVar(name string) Any {
	if len(name) == 0 {
		panic("Empty variable name")
	}
	if name[0] == '_' {
		if fr.Slots == nil {
			panic("No slots in this frame: " + name)
		}
		return (*fr.Slots)[name]
	}

	if IsGlobal(name) {
		return (*fr.G.Fr.Vars)[name]
	}
	return (*fr.Vars)[name]
}

func (fr *Frame) Apply(argv List) Any {
	log.Printf("< Apply < %#v\n", argv)
	head := argv[0]
	cmdName, ok := head.(string)
	if !ok {
		panic(Sprintf("Command must be a string: %#v", head))
	}

	fn, ok := fr.G.Cmds[cmdName]
	if !ok {
		fn, ok = Builtins[cmdName]
	}
	if !ok {
		panic(Sprintf("Command not found: %q", cmdName))
	}
	z := fn(fr, argv)
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
