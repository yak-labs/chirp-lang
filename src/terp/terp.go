package terp

import (
	. "fmt"
	"go/ast"
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
	Vars  Scope
	Slots Scope

	Prev *Frame
	G    *Global
	Mu   sync.Mutex
}

type Global struct {
	Cmds CmdScope
	Fr   Frame // global scope

	Mu sync.Mutex
}

func New() *Frame {
	g := &Global{
		Cmds: make(CmdScope),
		Fr: Frame{
			Vars: make(Scope),
		},
	}

	g.Fr.G = g
	g.Fr.initBuiltins()
	g.Fr.initExterns()
	return &g.Fr
}

func (fr *Frame) NewFrame() *Frame {
	return &Frame{
		Vars:  make(Scope),
		Slots: nil,
		Prev:  fr,
		G:     fr.G,
	}
}

func NewList(a ...Any) List { return List(a) }

func IsGlobal(name string) bool {
	if len(name) == 0 {
		panic("Empty variable name")
	}
	return ast.IsExported(name) // Same criteria, First is Uppercase.
}

func IsLocal(name string) bool {
	if len(name) == 0 {
		panic("Empty variable name")
	}
	return !ast.IsExported(name) && name[0] != '_'
}

func (fr *Frame) GetVarScope(name string) Scope {
	if len(name) == 0 {
		panic("Empty variable name")
	}
	if name[0] == '_' {
		if fr.Slots == nil {
			panic("No slots in this frame: " + name)
		}
		return fr.Slots
	}

	if IsGlobal(name) {
		return fr.G.Fr.Vars
	}
	return fr.Vars
}

func (fr *Frame) GetVar(name string) Any {
	return fr.GetVarScope(name)[name]
}

func (fr *Frame) SetVar(name string, x Any) {
	fr.GetVarScope(name)[name] = x
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
