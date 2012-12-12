package terp

import (
	"bytes"
	. "fmt"
	"go/ast"
	"log"
	"strings"
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

func Bool2Int(b bool) int {
	if b {return 1}
	return 0
}
func Bool2Str(b bool) string {
	if b {return "1"}
	return "0"
}
func Float32Str(x float32) string {
	if float32(int64(x)) == x {
		return Sprintf("%d", int64(x))
	}
	return Sprintf("%g", float32(x))
}
func Float64Str(x float64) string {
	if float64(int64(x)) == x {
		return Sprintf("%d", int64(x))
	}
	return Sprintf("%g", float64(x))
}
func List2Str(v List) string {
	buf := bytes.NewBuffer(nil)
	sep := ""
	for _, e := range v {
		buf.WriteString(sep)
		sep = " "
		estr := Any2ListElement(e)
		buf.WriteString(estr)
	}
	return buf.String()
}
func Any2ListElement(a Any) string {
	// TODO: Not perfect, but we are not doing \ yet.
	// TODO: Broken for mismatched {}.
	s := Str(a)
	if strings.ContainsAny(s, " \t\n\r{}\\") {
		return "{" + s + "}"
	}
	return s
}
func Repr(a Any) string { return Sprintf("<* %#v *>", a) }
func Str(a Any) string {
	switch x := a.(type) {
	case nil: return ""
	case string: return x
	case uint: return Sprintf("%d", x)
	case uint8: return Sprintf("%d", x)
	case uint16: return Sprintf("%d", x)
	case uint32: return Sprintf("%d", x)
	case uint64: return Sprintf("%d", x)
	case uintptr: return Sprintf("%d", x)
	case int: return Sprintf("%d", x)
	case int8: return Sprintf("%d", x)
	case int16: return Sprintf("%d", x)
	case int32: return Sprintf("%d", x)
	case int64: return Sprintf("%d", x)
	case float32: return Float32Str(x)
	case float64: return Float64Str(x)
	case bool: return Bool2Str(x)
	case error: return Sprintf("%#v", x)
	case List: return List2Str(x)
	}
	panic(Sprintf("DEFAULT Str: %#v", a))
	return Sprintf("%#v", a)
}

func Must(a, b Any, extra ...Any) {
	if a != b {
		panic(Repr(a) + " vs. " + Repr(b) + " :: " + Repr(extra))
	}
}
