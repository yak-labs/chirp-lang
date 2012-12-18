package terp

import (
	"bytes"
	. "fmt"
	"go/ast"
	"log"
	R "reflect"
	"strconv"
	"strings"
	"sync"

	"generated"
)

type Any interface{}

type List []Any
type Dict map[string]Any
type TDict map[string]T

type Command func(fr *Frame, argv List) Any
type TCommand func(fr *Frame, argv []T) T

type Scope map[string]Any
type TScope map[string]T

type CmdScope map[string]Command
type TCmdScope map[string]TCommand

type Frame struct {
	Vars  Scope
	Slots Scope
	TVars  TScope
	TSlots TScope

	Prev *Frame
	G    *Global
	Mu   sync.Mutex
}

type Global struct {
	Cmds CmdScope
	TCmds TCmdScope
	Fr   Frame // global scope

	Mu sync.Mutex
}

func New() *Frame {
	g := &Global{
		Cmds: make(CmdScope),
		TCmds: make(TCmdScope),
		Fr: Frame{
			Vars: make(Scope),
			TVars: make(TScope),
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
		TVars:  make(TScope),
		TSlots: nil,
		Prev:  fr,
		G:     fr.G,
	}
}

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

func (fr *Frame) TGetVarScope(name string) TScope {
	if len(name) == 0 {
		panic("Empty variable name")
	}
	if name[0] == '_' {
		if fr.Slots == nil {
			panic("No slots in this frame: " + name)
		}
		return fr.TSlots
	}

	if IsGlobal(name) {
		return fr.G.Fr.TVars
	}
	return fr.TVars
}

func (fr *Frame) TGetVar(name string) T {
	return fr.TGetVarScope(name)[name]
}

func (fr *Frame) TSetVar(name string, x T) {
	fr.TGetVarScope(name)[name] = x
}

func (fr *Frame) Apply(argv List) Any {
	log.Printf("< Apply < %#v\n", argv)
	head := argv[0]
	cmdName, ok := head.(string)
	if !ok {
		// Some day this may not be true; for now, it helps debug.
		panic(Sprintf("Command must be a string: %#v", head))
	}

	fn, ok := fr.G.Cmds[cmdName]
	log.Printf("Looked in Cmds %v %v %v", fn, ok, cmdName)
	if !ok {
		fn, ok = Builtins[cmdName]
		log.Printf("Looked in Builtins %v %v %v", fn, ok, cmdName)
	}
	if !ok {
		_, ok = generated.Funcs[cmdName]
		log.Printf("Looked in gen.Funcs -- %v %v", ok, cmdName)
		if ok {
			fn = cmdCall
			tmp := List{"call", cmdName}
			for _, a := range argv[1:] {
				tmp = append(tmp, a)
			}
			argv = tmp
		    log.Printf("NEW argv: $#v", argv)
		}
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
	if a == nil {
		return "{}"
	}
	s := Str(a)
	if s == "" {
		return "{}"
	}

	if strings.ContainsAny(s, " \t\n\r{}\\") {
		return "{" + s + "}"
	}
	return s
}
func Repr(a Any) string { return Sprintf("REPR<<%#v>>", a) }
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
	// panic(Sprintf("DEFAULT Str: %#v", a))
	return Sprintf("{%#v}", a)
}

func Must(a, b Any, extra ...Any) {
	if a != b {
		panic(Repr(a) + " vs. " + Repr(b) + " :: " + Repr(extra))
	}
}

func NewList(a ...Any) List { return List(a) }

func ToList(a Any) List {
	switch x := a.(type) {
		case List: return x
		case []Any: return List(x)
		case []interface{}:
			z := make([]Any, 0, 4)
			for _, e := range x {
				z = append(z, e)
			}
			return List(z)
		case string: return ParseList(x)
	}
	return ParseList(Str(a))
}
func LAppend(p Any, a ...Any) List {
	v := ToList(p)
	for _, e := range a {
		v = append(v, e)
	}
	return v
}

///////////////////////////////////////

type T interface {
	String() string
	Float() float64
	Int() int64
	Uint() uint64
	Bool() bool
	ListElement() string

	ToTf() Tf
	ToTs() Ts
	ToTl() Tl
}

type Tf struct {  // Tfloat
	f float64
}
type Ts struct {  // Tstring
	s string
}
type Tl struct {  // Tlist
	l []T
}
type Tv struct {  // Tvalue
	v R.Value
}

func FloatToT(a float64) Tf {
	return Tf{f: a}
}
func IntToT(a int64) Tf {
	return Tf{f: float64(a)}
}
func UintToT(a uint64) Tf {
	return Tf{f: float64(a)}
}
func StringToT(a string) Ts {
	return Ts{s: a}
}
func ListToT(a []T) Tl {
	return Tl{l: a}
}
func ValueToT(a R.Value) T {
	return Tv{v: a}
}
func ToT(a interface{}) T {
	switch z := a.(type) {
	case float64: return FloatToT(z)
	case float32: return FloatToT(float64(z))
	case int: return IntToT(int64(z))
	case int8: return IntToT(int64(z))
	case int16: return IntToT(int64(z))
	case int32: return IntToT(int64(z))
	case int64: return IntToT(z)
	case uint: return UintToT(uint64(z))
	case uint8: return UintToT(uint64(z))
	case uint16: return UintToT(uint64(z))
	case uint32: return UintToT(uint64(z))
	case uint64: return UintToT(z)
	case string: return StringToT(z)
	case T: panic("Already a T")
	}
	return ValueToT(R.ValueOf(a))
}
	


func (t Tf) String() string {
	return Sprintf("%g", t.f)
}
func (t Tf) ListElement() string {
	return t.String()
}
func (t Tf) Float() float64 {
	return t.f
}
func (t Tf) Int() int64 {
	return int64(t.f)
}
func (t Tf) Uint() uint64 {
	return uint64(t.f)
}
func (t Tf) Bool() bool {
	if t.f == 0 {
		return false
	}
	return true
}
func (t Tf) ToTf() Tf {
	return t
}
func (t Tf) ToTs() Ts {
	return Ts{s: t.String()}
}
func (t Tf) ToTl() Tl {
	return Tl{l: []T{t,}}
}



func (t Ts) String() string {
	return t.s
}
func (t Ts) ListElement() string {
	return ToListElement(t.s)
}
func (t Ts) Float() float64 {
	z, err := strconv.ParseFloat(t.s, 64)
	if err != nil {
		panic(err)
	}
	return z
}
func (t Ts) Int() int64 {
	return int64(t.Float())  //TODO
}
func (t Ts) Uint() uint64 {
	return uint64(t.Float())  //TODO
}
func (t Ts) Bool() bool {
	if t.s == "" || t.s == "0" {
		return false
	}
	return true
}
func (t Ts) ToTf() Tf {
	return FloatToT(t.Float())
}
func (t Ts) ToTs() Ts {
	return t
}
func (t Ts) ToTl() Tl {
	v := ParseList(t.s)
	z := make([]T, len(v))
	for i, e := range v {
		z[i] = ToT(e)
	}
	return Tl{l: z}
}



func (t Tl) String() string {
	z := ""
	for k, v := range t.l {
		if k > 0 {
			z += " "
		}
		z += v.ListElement()
	}
	return z
}
func (t Tl) ListElement() string {
	return ToListElement(t.String())
}
func (t Tl) Float() float64 {
	if len(t.l) != 1 {panic("cant")}
	return t.l[0].Float()
}
func (t Tl) Int() int64 {
	if len(t.l) != 1 {panic("cant")}
	return t.l[0].Int()
}
func (t Tl) Uint() uint64 {
	if len(t.l) != 1 {panic("cant")}
	return t.l[0].Uint()
}
func (t Tl) Bool() bool {
	if len(t.l) == 0 {
		return false
	}
	return true
}
func (t Tl) ToTf() Tf {
	return FloatToT(t.Float())
}
func (t Tl) ToTs() Ts {
	return StringToT(t.String())
}
func (t Tl) ToTl() Tl {
	return t
}


func (t Tv) String() string {
	return Sprintf("Value:%s:%s:%d", t.v.Kind(), t.v.Type(), t.v.Addr())
}
func (t Tv) ListElement() string {
	return ToListElement(t.String())
}
func (t Tv) Float() float64 {
	panic("cant yet")
}
func (t Tv) Int() int64 {
	panic("cant yet")
}
func (t Tv) Uint() uint64 {
	panic("cant yet")
}
func (t Tv) Bool() bool {
	panic("cant yet")
}
func (t Tv) ToTf() Tf {
	return Tf{f: t.Float()}
}
func (t Tv) ToTs() Ts {
	return Ts{s: t.String()}
}
func (t Tv) ToTl() Tl {
	return Tl{l: []T{t,}}
}


func ToListElement(s string) string {
	// TODO: Not perfect, but we are not doing \ yet.
	// TODO: Broken for mismatched {}.
	if s == "" {
		return "{}"
	}

	if strings.ContainsAny(s, " \t\n\r{}\\") {
		return "{" + s + "}"
	}
	return s
}




