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
)

type Hash map[string]T

type TCommand func(fr *Frame, argv []T) T

type TScope map[string]Loc

type TCmdScope map[string]TCommand

type Frame struct {
	TVars  TScope
	TSlots TScope

	Prev *Frame
	G    *Global
	Mu   sync.Mutex
	Chan	chan<- T  // for yproc & yield
}

type Global struct {
	TCmds TCmdScope
	Fr    Frame // global scope

	Mu sync.Mutex
}

type StatusCode int 
const (
	RETURN = StatusCode(iota + 2)
	BREAK
	CONTINUE
)

// Panic a Jump for return, break, and continue.
type Jump struct {
	Status StatusCode
	Result T
}

type Loc interface {
	Get() T
	Set(T)
}

type Slot struct {
	Elem	T
}
type UpSlot struct {
	Fr		*Frame
	RemoteName string
}

func New() *Frame {
	g := &Global{
		TCmds: make(TCmdScope),
		Fr: Frame{
			TVars: make(TScope),
		},
	}

	g.Fr.G = g
	g.Fr.initTBuiltins()
	g.Fr.initReflect()
	g.Fr.initExpr()
	return &g.Fr
}

func (fr *Frame) NewFrame() *Frame {
	return &Frame{
		TVars:  make(TScope),
		TSlots: nil,
		Prev:   fr,
		G:      fr.G,
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

func (p* Slot) Get() T { return p.Elem }
func (p* Slot) Set(t T) { p.Elem = t }

func (fr *Frame) TGetVarScope(name string) TScope {
	if len(name) == 0 {
		panic("Empty variable name")
	}
	if name[0] == '_' {
		if fr.TSlots == nil {
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
	return fr.TGetVarScope(name)[name].Get()
}

func (fr *Frame) TSetVar(name string, x T) {
	sc := fr.TGetVarScope(name)
	if sc[name] == nil {
		sc[name] = new(Slot)
	}
	sc[name].Set(x)
}

func (p* UpSlot) Get() T { return p.Fr.TGetVar(p.RemoteName) }
func (p* UpSlot) Set(t T) { p.Fr.TSetVar(p.RemoteName, t) }

func (fr *Frame) TUpVar(name string, remFr *Frame, remName string) {
	sc := fr.TGetVarScope(name)
	sc[name] = &UpSlot{Fr: remFr, RemoteName: remName}
}

func (fr *Frame) FindCommand(name T) TCommand {
	// Some day we will not require Ts; for now, it helps debug.
	cmdName, ok := name.(Ts)
	if !ok {
		panic(Sprintf("Restriction: Command must be a string: %#v", name))
	}

	fn, ok := fr.G.TCmds[cmdName.s]
	if !ok {
		fn, ok = TBuiltins[cmdName.s]
	}
	if !ok {
		panic(Sprintf("FindCommand: command not found: %q", cmdName.s))
	}
	return fn
}

func (fr *Frame) TApply(argv []T) T {
	head := argv[0]
	log.Printf("< TApply < %q", head)
	for ai, av := range argv[1:] {
		log.Printf("< ...... < [%d] (%T) ## %#v ## %q", ai, av, av, av.String())
	}

	// Some day we will not require Ts; for now, it helps debug.
	cmdName, ok := head.(Ts)
	if !ok {
		panic(Sprintf("Command must be a string: %s", Show(head)))
	}

	if len(cmdName.s) > 1 && cmdName.s[0] == '/' {
		call := []T{MkTs("call"), head}
		call = append(call, argv[1:]...) // Append all but first of argv.
		return fr.TApply(call)           // Recurse.
	}

	fn := fr.FindCommand(head)
	z := fn(fr, argv)
	log.Printf("> TApply > (%T) ## %#v ## %q", z, z, z.String())
	return z
}

func Repr(a interface{}) string { return Sprintf("REPR<<%#v>>", a) }

func TMust(a, b T) {
	if a.String() != b.String() {
		panic(Show(a) + " .vs. " + Show(b))
	}
}

// MustT takes a string and a T
func MustT(a string, b T) {
	TMust(MkTs(a), b)
}

func Must(a, b interface{}) {
	// Otherwise use Repr equality
	if Repr(a) != Repr(b) {
		panic(Repr(a) + " .vs. " + Repr(b))
	}
}

func Show(a T) string {
	return Sprintf("{(%T) ## %#v ## %q}", a, a, a.String())
}

func Showv(a []T) string {
	buf := bytes.NewBufferString(Sprintf("Slice of T with %d slots:", len(a)))
	for i, e := range a {
		buf.WriteString(Sprintf("\n    ... [%d] = %s", i, Show(e)))
	}
	return buf.String()
}

///////////////////////////////////////

type T interface {
	Raw() interface{}
	String() string
	Float() float64
	Int() int64
	Uint() uint64
	Bool() bool
	ListElement() string
	Truth() bool   // Like Python, empty values and 0 values are false.
	IsEmpty() bool // Would String() return ""?
	List() []T
	HeadTail() (hd, tl T)
}

type Tf struct { // Implements T.
	f float64
}
type Ts struct { // Implements T.
	s string
}
type Tl struct { // Implements T.
	l []T
}
type Tv struct { // Implements T.
	v R.Value
}

type Ty struct { // Implements T.
	ch <-chan T
	hd T
	tl T
}

var Empty = MkTs("")

func MkTy(ch <-chan T) Ty {
	return Ty{ch: ch}
}

func MkTb(a bool) Tf {
	if a {
		return MkTi(1)
	}
	return MkTi(0)
}
func MkTf(a float64) Tf {
	return Tf{f: a}
}
func MkTi(a int64) Tf {
	return Tf{f: float64(a)}
}
func MkTu(a uint64) Tf {
	return Tf{f: float64(a)}
}
func MkTs(a string) Ts {
	if len(a) > 6 && a[:6] == "Value:" {
		panic(666)
	}
	return Ts{s: a}
}
func MkTl(a []T) Tl {
	log.Printf("MkTl: from <%T> <%s>", a, a)
	return Tl{l: a}
}
func MkTv(a R.Value) T {
	return Tv{v: a}
}
func MkT(a interface{}) T {
	log.Printf("MkT <-- (%T)   %v", a, a)

	// Very specific type cases.
	switch x := a.(type) {
	case T:
		panic(Sprintf("Calling MkT() on a T: <%T> <%#v> %s", x, x, x.String()))
	case R.Value:
		// Some day we'll allow this, but for now, flag an error.
		panic(Sprintf("Calling MkT() on a R.Value: <%T> <%#v> %s", x, x, x.String()))
	case nil:
		return Empty

	case string:
		if len(x) > 6 && x[:6] == "Value:" {
			panic(666)
		}

	}

	// Use reflection to figure it out.
	v := R.ValueOf(a)
	switch v.Kind() {

	case R.Bool:
		return MkTb(v.Bool())

	case R.Int:
		return MkTi(v.Int())
	case R.Int8:
		return MkTi(v.Int())
	case R.Int16:
		return MkTi(v.Int())
	case R.Int32:
		return MkTi(v.Int())
	case R.Int64:
		return MkTi(v.Int())

	case R.Uint:
		return MkTu(v.Uint())
	case R.Uint8:
		return MkTu(v.Uint())
	case R.Uint16:
		return MkTu(v.Uint())
	case R.Uint32:
		return MkTu(v.Uint())
	case R.Uint64:
		return MkTu(v.Uint())
	case R.Uintptr:
		return MkTu(v.Uint())

	case R.Float32:
		return MkTf(v.Float())
	case R.Float64:
		return MkTf(v.Float())

	case R.Complex64:
	case R.Complex128:

	case R.Array:
	case R.Chan:
		if v.IsNil() {
			return Empty
		}
	case R.Func:
		if v.IsNil() {
			return Empty
		}
	case R.Interface:
		if v.IsNil() {
			return Empty
		}
	case R.Map:
		if v.IsNil() {
			return Empty
		}
	case R.Ptr:
		if v.IsNil() {
			return Empty
		}
	case R.Slice:
		if v.IsNil() {
			return Empty
		}

		/*
			// This will convert slices to lists.
			// Is this a good idea?
			return Tv{v: v}.Tl()
		*/

		var pointerToT *T
		switch v.Type().Elem() {
		case R.TypeOf(pointerToT).Elem(): // i.e. case T
			return MkTl(v.Interface().([]T))
		}
		switch v.Type().Elem().Kind() {
		case R.Uint8:
			return MkTs(string(v.Interface().([]byte)))
		}

	case R.String:
		return MkTs(v.Interface().(string))
	case R.Struct:
	case R.UnsafePointer:
	}

	// Everything else becomes a Tv
	log.Printf("MkT --> defaulting to Tv")
	return MkTv(v)
}

// Ty implements T

func (t Ty) Raw() interface{} {
	panic("not implemented on generator (Ty)")
}
func (t Ty) String() string {
	return Repr(t)
}
func (t Ty) Float() float64 {
	panic("not implemented on generator (Ty)")
}
func (t Ty) Int() int64 {
	panic("not implemented on generator (Ty)")
}
func (t Ty) Uint() uint64 {
	panic("not implemented on generator (Ty)")
}
func (t Ty) Bool() bool {
	panic("not implemented on generator (Ty)")
}
func (t Ty) ListElement() string {
	panic("not implemented on generator (Ty)")
}
func (t Ty) Truth() bool {
	panic("not implemented on generator (Ty)")
}
func (t Ty) IsEmpty() bool {
	panic("not implemented on generator (Ty)")
}
func (t Ty) List() []T {
	z := make([]T, 0, 4)
	for {
		t := <-t.ch
		if t == nil {
			break
		}
		z = append(z, t)
	}
	return z
}
func (t Ty) HeadTail() (hd, tl T) {
	if t.ch == nil {
		return t.hd, t.tl
	}
	t.hd = <-t.ch
	if t.hd == nil {
		t.ch = nil
		return nil, nil
	}
	t.tl = Ty{ch: t.ch}
	return t.hd, t.tl
}

// Tf implements T

func (t Tf) Raw() interface{} {
	return t.f
}
func (t Tf) String() string {
	return Sprintf("%g", t.f)
}
func (t Tf) ListElement() string {
	return t.String()
}
func (t Tf) Truth() bool {
	return t.f != 0
}
func (t Tf) IsEmpty() bool {
	return false
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
func (t Tf) List() []T {
	return []T{t}
}
func (t Tf) HeadTail() (hd, tl T) {
	return MkTl(t.List()).HeadTail()
}

// Ts implements T

func (t Ts) Raw() interface{} {
	return t.s
}
func (t Ts) String() string {
	return t.s
}
func (t Ts) ListElement() string {
	return ToListElement(t.s)
}
func (t Ts) Truth() bool {
	return t.s != ""
}
func (t Ts) IsEmpty() bool {
	return t.s == ""
}
func (t Ts) Float() float64 {
	z, err := strconv.ParseFloat(t.s, 64)
	if err != nil {
		panic(err)
	}
	return z
}
func (t Ts) Int() int64 {
	return int64(t.Float()) //TODO
}
func (t Ts) Uint() uint64 {
	return uint64(t.Float()) //TODO
}
func (t Ts) Bool() bool {
	if t.s == "" || t.s == "0" {
		return false
	}
	return true
}
func (t Ts) List() []T {
	return ParseList(t.s)
}
func (t Ts) HeadTail() (hd, tl T) {
	return MkTl(t.List()).HeadTail()
}

// Tl implements T

func (t Tl) Raw() interface{} {
	z := make([]interface{}, len(t.l))
	for i, e := range t.l {
		z[i] = e.Raw() // Recurse.
	}
	return z
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
func (t Tl) Truth() bool {
	return len(t.l) != 0
}
func (t Tl) IsEmpty() bool {
	return len(t.l) == 0
}
func (t Tl) Float() float64 {
	if len(t.l) != 1 {
		panic("cant")
	}
	return t.l[0].Float()
}
func (t Tl) Int() int64 {
	if len(t.l) != 1 {
		panic("cant")
	}
	return t.l[0].Int()
}
func (t Tl) Uint() uint64 {
	if len(t.l) != 1 {
		panic("cant")
	}
	return t.l[0].Uint()
}
func (t Tl) Bool() bool {
	if len(t.l) == 0 {
		return false
	}
	return true
}
func (t Tl) List() []T {
	return t.l
}
func (t Tl) HeadTail() (hd, tl T) {
	if len(t.l) == 0 {
		return nil, nil
	}
	return t.l[0], MkTl(t.l[1:])
}

// Tv implements T

func (t Tv) Raw() interface{} {
	return t.v.Interface()
}
func (t Tv) String() string {
	s := Sprintf("Value:%s:%s", t.v.Kind(), t.v.Type())
	log.Printf("Warning: converting to Tv to String: %q", s)
	return s
}
func (t Tv) ListElement() string {
	return ToListElement(t.String())
}
func (t Tv) Truth() bool {
	// TODO
	panic("Restriction: cannot test Tv for Truth")
}
func (t Tv) IsEmpty() bool {
	switch t.v.Kind() {
	// IsNil() can only be called on this 6 Kinds:
	case R.Func, R.Interface, R.Ptr, R.Slice, R.Map, R.Chan:
		return t.v.IsNil()
	}
	// Strings, numbers, and bools should not be in Tv so we don't look for emptiness or zeroness in them.
	return false
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
func (t Tv) List() []T {
/***
	Is this a good idea?

	At times, it is really convenient to have a Raw Slice be a list.

	But other times, we want to edit that Raw Slice in place.

	Maybe this is right -- only when you explicitly ask for a List() do we explode it.

	Is treating Ptr and Iface like a Singleton List a good idea?
***/
	switch t.v.Kind() {

	// Treat Pointer and Interface as a singleton list.
	case R.Ptr, R.Interface:
		x := MkT(t.v.Elem().Interface())
		return []T{x}

	// Slices and Arrays are naturally lists (unless they're bytes)
	case R.Slice, R.Array:
		if t.v.Type().Elem().Kind() == R.Uint8 {
			panic(Sprintf("Slice of Uint8 should not be in Tv: %q", string(t.v.Interface().([]byte))))
		}
		n := t.v.Len()
		z := make([]T, n)
		for i := 0; i < n; i++ {
			z[i] = MkT(t.v.Index(i).Interface())
		}
		return z
	}
/********/
	return []T{t}
}
func (t Tv) HeadTail() (hd, tl T) {
	return MkTl(t.List()).HeadTail()
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
