package terp

import (
	"bytes"
	. "fmt"
	"go/ast"
	"log"
	R "reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Hash map[string]T // TODO: Mutex

type Command func(fr *Frame, argv []T) T

type Scope map[string]Loc

type CmdScope map[string]*CmdNode

// CmdNode makes a singly-linked-list of commands
// at different mixin levels, highest level first.
// A non-mixin command has level 0 and only one CmdNode.
type CmdNode struct {
	Fn	Command
	MixinLevel	int
	MixinName	string
	Next	*CmdNode
}

// Frame is a local variable frame.
// There is one for the global variables in the Global struct,
// and a new one is created for each proc or yproc invocation
// (but not for every Command; non-proc commands do not make Frames).
type Frame struct {
	Vars  Scope

	Prev *Frame
	G    *Global

	Chan chan<- T // for yproc & yield
	MixinLevel int
}

// Global holds the global state of an interpreter,
// mainly the Commands and global variables.
// It also knows if a Mixin is being defined.
// Mixins should be defined by main thread,
// after all overridable procs are defined,
// but before other goroutines start.
type Global struct {
	Cmds CmdScope
	Fr    Frame // global scope
	SubTerps	map[string]*Global	

	MixinSerial		int  // Increment before defining Mixin.
	MixinNumberDefining	int  // Set nonzero while defining Mixin.
	MixinNameDefining	string  // Set nonzero while defining Mixin.

	Mu sync.Mutex
}

// Clone produces a copy of the receiving interpreter.
func (g *Global) Clone() *Global {
	z := &Global{
		Cmds: make(CmdScope),
		Fr: Frame{
			Vars: make(Scope),
		},
		SubTerps: make(map[string]*Global),
	}

	z.Fr.G = z

	for k, v := range g.Cmds {
		z.Cmds[k] = v
	}
	
	for k, loc := range g.Fr.Vars {
		z.Fr.SetVar(k, loc.Get())
	}

	return z
}

// StatusCode are the same integers as Tcl/C uses for return, break, and continue.
type StatusCode int
const (
	RETURN = StatusCode(iota + 2)
	BREAK
	CONTINUE
)

// Jump structs are panicked for return, break, and continue.
type Jump struct {
	Status StatusCode
	Result T
}

// Loc is protocol for a variable location.
type Loc interface {
	Get() T
	Set(T)
}

// Slot stores a variable value.
type Slot struct {
	Elem T
}
// UpSlot forwards a variable to another variable.
type UpSlot struct {
	Fr         *Frame
	RemoteName string
}

var Empty = MkString("")
var InvalidValue = *new(R.Value)

// Create a new interpreter, and return the global frame pointer.
func New() *Frame {
	g := &Global{
		Cmds: make(CmdScope),
		Fr: Frame{
			Vars: make(Scope),
		},
		SubTerps: make(map[string]*Global),
	}

	g.Fr.G = g
	g.Fr.initBuiltins()
	g.Fr.initReflect()
	g.Fr.initExpr()
	g.Fr.initDbCmds()
	g.Fr.initTemplateCmds()
	return &g.Fr
}

func (fr *Frame) NewFrame() *Frame {
	return &Frame{
		Vars:  make(Scope),
		Prev:   fr,
		G:      fr.G,
	}
}

// Initial capital letter for a variable means Global.
func IsGlobal(name string) bool {
	if len(name) == 0 {
		panic("Empty variable name")
	}
	return ast.IsExported(name) // Same criteria, First is Uppercase.
}

// Initial capital letter for a variable means local.
func IsLocal(name string) bool {
	if len(name) == 0 {
		panic("Empty variable name")
	}
	return !ast.IsExported(name)
}

func (p *Slot) Get() T  { return p.Elem }
func (p *Slot) Set(t T) { p.Elem = t }

func (fr *Frame) GetVarScope(name string) Scope {
	if len(name) == 0 {
		panic("Empty variable name")
	}

	if IsGlobal(name) {
		return fr.G.Fr.Vars
	}
	return fr.Vars
}

func (fr *Frame) GetVar(name string) T {
	sc := fr.GetVarScope(name)
	var loc Loc
	var ok bool
	if loc, ok = sc[name]; !ok {
		panic(Sprintf("Variable %q does not exist; scope contains %v", name, sc))
	}
	return loc.Get()
}

func (fr *Frame) SetVar(name string, x T) {
	sc := fr.GetVarScope(name)
	if sc[name] == nil {
		sc[name] = new(Slot)
	}
	sc[name].Set(x)
}

func (p *UpSlot) Get() T  { return p.Fr.GetVar(p.RemoteName) }
func (p *UpSlot) Set(t T) { p.Fr.SetVar(p.RemoteName, t) }

func (fr *Frame) UpVar(name string, remFr *Frame, remName string) {
	sc := fr.GetVarScope(name)
	sc[name] = &UpSlot{Fr: remFr, RemoteName: remName}
}

func (fr *Frame) FindCommand(name T, callSuper bool) Command {
	// Some day we will not require terpString; for now, it helps debug.
	cmdName, ok := name.(terpString)
	if !ok {
		panic(Sprintf("Restriction: Command must be a string: %#v", name))
	}

	var fn Command
	cmdNode, ok := fr.G.Cmds[cmdName.s]
	if ok {
		if callSuper {
			maxMixinLevel := fr.MixinLevel - 1
			if maxMixinLevel < 0 {
				panic("cannot callSuper from non-mixin")
			}
	        log.Printf("FindCommand: callSuper: maxMixinLevel=%d try=%#v", maxMixinLevel, cmdNode)
			for cmdNode != nil && cmdNode.MixinLevel > maxMixinLevel {
	        	log.Printf("FindCommand: callSuper: maxMixinLevel=%d try=%#v", maxMixinLevel, cmdNode)
				cmdNode = cmdNode.Next
			}
	        log.Printf("FindCommand: callSuper: OK")
		}
		if cmdNode == nil {
			ok = false
		} else {
	        log.Printf("FindCommand: Choosing mixin level %d from %#v", cmdNode.MixinLevel, cmdNode)
			fn = cmdNode.Fn
		}
	} else {
		fn, ok = Builtins[cmdName.s]
	}
	if !ok {
		panic(Sprintf("FindCommand: command not found: %q", cmdName.s))
	}
	return fn
}

// Apply a command with its arguments.
func (fr *Frame) Apply(argv []T) T {
	head := argv[0]
	log.Printf("< Apply < %q", head)
	for ai, av := range argv[1:] {
		log.Printf("< ...... < [%d] (%T) ## %#v ## %q", ai, av, av, av.String())
	}

	// Some day we will not require terpString; for now, it helps debug.
	cmdName, ok := head.(terpString)
	if !ok {
		panic(Sprintf("Command must be a string: %s", Show(head)))
	}

	if len(cmdName.s) > 1 && cmdName.s[0] == '/' {
		call := []T{MkString("call"), head}
		call = append(call, argv[1:]...) // Append all but first of argv.
		return fr.Apply(call)           // Recurse.
	}

	fn := fr.FindCommand(head, false) // false: Don't call super.
	z := fn(fr, argv)
	log.Printf("> Apply > %s", Show(z))
	return z
}

func Repr(a interface{}) string { return Sprintf("REPR<<%#v>>", a) }

// Must takes 2 T values, and compares their Show()s.
func Must(a, b T) {
	if a.String() != b.String() {
		panic(Show(a) + " .vs. " + Show(b))
	}
}

// MustST takes a string and a T
func MustST(a string, b T) {
	Must(MkString(a), b)
}

// MustA takes Any 2 values, and compares their Repr()s.
func MustA(a, b interface{}) {
	if Repr(a) != Repr(b) {
		panic(Repr(a) + " .vs. " + Repr(b))
	}
}

func Show(a T) string {
	return Sprintf("{(%T) ## %#v ## %q}", a, a, a.String())
}

func Showv(a []T) string {
	buf := bytes.NewBufferString(Sprintf("Slice of T with %d elements:", len(a)))
	for i, e := range a {
		buf.WriteString(Sprintf("\n    ... [%d] = %s", i, Show(e)))
	}
	return buf.String()
}

///////////////////////////////////////

// T is an interface to any Tcl value.
// Use them only through these methods, or fix these methods.
type T interface {
	Raw() interface{}
	String() string
	Float() float64
	Int() int64
	Uint() uint64
	ListElementString() string
	Truth() bool   // Like Python, empty values and 0 values are false.
	IsEmpty() bool // Would String() return ""?
	List() []T
	IsPreservedByList() bool
	HeadTail() (hd, tl T)
	Hash() Hash
	QuickReflectValue() R.Value
}

// terpFloat is a Tcl value holding a float64.
type terpFloat struct { // Implements T.
	f float64
}

// terpString is a Tcl value holding a string.
type terpString struct { // Implements T.
	s string
}

// terpList is a Tcl value holding a List.
type terpList struct { // Implements T.
	l []T
}

// terpValue is a Tcl value holding a Go reflect.Value.
// It is a handle to non-Tcl Go objets.
type terpValue struct { // Implements T.
	v R.Value
}

// terpGenerator holds a channel for reading from a generator (yproc command).
type terpGenerator struct { // Implements T.
	ch <-chan T
	hd T
	tl T
}

// terpHash holds a Hash.
type terpHash struct { // Imlements T.
	h Hash
}

func MkHash() terpHash {
	return terpHash{h: make(Hash, 4)}
}
func MkGenerator(ch <-chan T) terpGenerator {
	return terpGenerator{ch: ch}
}
func MkBool(a bool) terpFloat {
	if a {
		return MkInt(1)
	}
	return MkInt(0)
}
func MkFloat(a float64) terpFloat {
	return terpFloat{f: a}
}
func MkInt(a int64) terpFloat {
	return terpFloat{f: float64(a)}
}
func MkUint(a uint64) terpFloat {
	return terpFloat{f: float64(a)}
}
func MkString(a string) terpString {
	if len(a) > 6 && a[:6] == "Value:" {
		panic(666)
	}
	return terpString{s: a}
}
func MkList(a []T) terpList {
	log.Printf("MkList: from <%T> <%s>", a, a)
	return terpList{l: a}
}
func MkValue(a R.Value) terpValue {
	return terpValue{v: a}
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
		return MkBool(v.Bool())

	case R.Int:
		return MkInt(v.Int())
	case R.Int8:
		return MkInt(v.Int())
	case R.Int16:
		return MkInt(v.Int())
	case R.Int32:
		return MkInt(v.Int())
	case R.Int64:
		return MkInt(v.Int())

	case R.Uint:
		return MkUint(v.Uint())
	case R.Uint8:
		return MkUint(v.Uint())
	case R.Uint16:
		return MkUint(v.Uint())
	case R.Uint32:
		return MkUint(v.Uint())
	case R.Uint64:
		return MkUint(v.Uint())
	case R.Uintptr:
		return MkUint(v.Uint())

	case R.Float32:
		return MkFloat(v.Float())
	case R.Float64:
		return MkFloat(v.Float())

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
			return terpValue{v: v}.terpList()
		*/

		var pointerToT *T
		switch v.Type().Elem() {
		case R.TypeOf(pointerToT).Elem(): // i.e. case T
			return MkList(v.Interface().([]T))
		}
		switch v.Type().Elem().Kind() {
		case R.Uint8:
			return MkString(string(v.Interface().([]byte)))
		}

	case R.String:
		return MkString(v.Interface().(string))
	case R.Struct:
	case R.UnsafePointer:
	}

	// Everything else becomes a terpValue
	log.Printf("MkT --> defaulting to terpValue")
	return MkValue(v)
}

// terpHash implements T

func (t terpHash) Raw() interface{} {
	panic("not implemented on generator (terpHash)")
}
func (t terpHash) String() string {
	return Repr(t)
}
func (t terpHash) Float() float64 {
	panic("not implemented on generator (terpHash)")
}
func (t terpHash) Int() int64 {
	panic("not implemented on generator (terpHash)")
}
func (t terpHash) Uint() uint64 {
	panic("not implemented on generator (terpHash)")
}
func (t terpHash) ListElementString() string {
	panic("not implemented on generator (terpHash)")
}
func (t terpHash) Truth() bool {
	return len(t.h) > 0
}
func (t terpHash) IsEmpty() bool {
	return len(t.h) == 0
}

type SortListByStringTSlice []T

func (p SortListByStringTSlice) Len() int           { return len(p) }
func (p SortListByStringTSlice) Less(i, j int) bool { return p[i].String() < p[j].String() }
func (p SortListByStringTSlice) Swap(i, j int)      { p[j], p[i] = p[i], p[j] }

func SortListByString(list []T) {
	sort.Sort(SortListByStringTSlice(list))
}

func SortedKeysOfHash(h Hash) []string {
	// TODO: mutex
	keys := make([]string, 0, len(h))

	for k, v := range h {
		if v == nil {
			continue // Omit phantoms and deletions.
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (t terpHash) IsPreservedByList() bool { return true }
func (t terpHash) List() []T {
	keys := SortedKeysOfHash(t.h)
	z := make([]T, 0, 2*len(keys))
	// TODO: mutex
	for _, k := range keys {
		v := t.h[k]
		if v == nil {
			continue // Omit phantoms and deletions.
		}
		z = append(z, MkString(k), v)
	}
	return z
}
func (t terpHash) HeadTail() (hd, tl T) {
	return MkList(t.List()).HeadTail()
}
func (t terpHash) Hash() Hash {
	return t.h
}
func (t terpHash) QuickReflectValue() R.Value { return InvalidValue }

// terpGenerator implements T

func (t terpGenerator) Raw() interface{} {
	panic("not implemented on generator (terpGenerator)")
}
func (t terpGenerator) String() string {
	return Repr(t)
}
func (t terpGenerator) Float() float64 {
	panic("not implemented on generator (terpGenerator)")
}
func (t terpGenerator) Int() int64 {
	panic("not implemented on generator (terpGenerator)")
}
func (t terpGenerator) Uint() uint64 {
	panic("not implemented on generator (terpGenerator)")
}
func (t terpGenerator) ListElementString() string {
	panic("not implemented on generator (terpGenerator)")
}
func (t terpGenerator) Truth() bool {
	panic("not implemented on generator (terpGenerator)")
}
func (t terpGenerator) IsEmpty() bool {
	panic("not implemented on generator (terpGenerator)")
}
func (t terpGenerator) IsPreservedByList() bool { return true }
func (t terpGenerator) List() []T {
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
func (t terpGenerator) HeadTail() (hd, tl T) {
	if t.ch == nil {
		return t.hd, t.tl
	}
	t.hd = <-t.ch
	if t.hd == nil {
		t.ch = nil
		return nil, nil
	}
	t.tl = terpGenerator{ch: t.ch}
	return t.hd, t.tl
}
func (t terpGenerator) Hash() Hash {
	panic("terpGenerator is not a Hash")
}
func (t terpGenerator) QuickReflectValue() R.Value { return InvalidValue }

// terpFloat implements T

func (t terpFloat) Raw() interface{} {
	return t.f
}
func (t terpFloat) String() string {
	return Sprintf("%g", t.f)
}
func (t terpFloat) ListElementString() string {
	return t.String()
}
func (t terpFloat) Truth() bool {
	return t.f != 0
}
func (t terpFloat) IsEmpty() bool {
	return false
}
func (t terpFloat) Float() float64 {
	return t.f
}
func (t terpFloat) Int() int64 {
	return int64(t.f)
}
func (t terpFloat) Uint() uint64 {
	return uint64(t.f)
}
func (t terpFloat) IsPreservedByList() bool { return true }
func (t terpFloat) List() []T {
	return []T{t}
}
func (t terpFloat) HeadTail() (hd, tl T) {
	return MkList(t.List()).HeadTail()
}
func (t terpFloat) Hash() Hash {
	panic(" is not a Hash")
}
func (t terpFloat) QuickReflectValue() R.Value { return InvalidValue }

// terpString implements T

func (t terpString) Raw() interface{} {
	return t.s
}
func (t terpString) String() string {
	return t.s
}
func (t terpString) ListElementString() string {
	return ToListElementString(t.s)
}
func (t terpString) Truth() bool {
	return t.s != "" && t.s != "0" // TODO: Reconsider Truth.
}
func (t terpString) IsEmpty() bool {
	return t.s == ""
}
func (t terpString) Float() float64 {
	z, err := strconv.ParseFloat(t.s, 64)
	if err != nil {
		panic(err)
	}
	return z
}
func (t terpString) Int() int64 {
	return int64(t.Float()) //TODO
}
func (t terpString) Uint() uint64 {
	return uint64(t.Float()) //TODO
}
func (t terpString) IsPreservedByList() bool { return false }
func (t terpString) List() []T {
	return ParseList(t.s)
}
func (t terpString) HeadTail() (hd, tl T) {
	return MkList(t.List()).HeadTail()
}
func (t terpString) Hash() Hash {
	panic("A string is not a Hash")
}
func (t terpString) QuickReflectValue() R.Value { return InvalidValue }

// terpList implements T

func (t terpList) Raw() interface{} {
	z := make([]interface{}, len(t.l))
	for i, e := range t.l {
		z[i] = e.Raw() // Recurse.
	}
	return z
}
func (t terpList) String() string {
	z := ""
	for k, v := range t.l {
		if k > 0 {
			z += " "
		}
		z += v.ListElementString()
	}
	return z
}
func (t terpList) ListElementString() string {
	return ToListElementString(t.String())
}
func (t terpList) Truth() bool {
	return len(t.l) != 0
}
func (t terpList) IsEmpty() bool {
	return len(t.l) == 0
}
func (t terpList) Float() float64 {
	if len(t.l) != 1 {
		panic("cant")
	}
	return t.l[0].Float()
}
func (t terpList) Int() int64 {
	if len(t.l) != 1 {
		panic("cant")
	}
	return t.l[0].Int()
}
func (t terpList) Uint() uint64 {
	if len(t.l) != 1 {
		panic("cant")
	}
	return t.l[0].Uint()
}
func (t terpList) IsPreservedByList() bool { return true }
func (t terpList) List() []T {
	return t.l
}
func (t terpList) HeadTail() (hd, tl T) {
	if len(t.l) == 0 {
		return nil, nil
	}
	return t.l[0], MkList(t.l[1:])
}
func (t terpList) Hash() Hash {
	panic("A List is not a Hash")
}
func (t terpList) QuickReflectValue() R.Value { return InvalidValue }

// terpValue implements T

func (t terpValue) Raw() interface{} {
	return t.v.Interface()
}
func (t terpValue) String() string {
	s := Sprintf("Value:%s:%s", t.v.Kind(), t.v.Type())
	log.Printf("Warning: converting to terpValue to String: %q", s)
	return s
}
func (t terpValue) ListElementString() string {
	return ToListElementString(t.String())
}
func (t terpValue) Truth() bool {
	// TODO
	return !t.IsEmpty()
	// TODO // panic("Restriction: cannot test terpValue for Truth")
}
func (t terpValue) IsEmpty() bool {
	switch t.v.Kind() {
	// IsNil() can only be called on this 6 Kinds:
	case R.Func, R.Interface, R.Ptr, R.Slice, R.Map, R.Chan:
		return t.v.IsNil()
	}
	// Strings, numbers, and bools should not be in terpValue so we don't look for emptiness or zeroness in them.
	return false
}
func (t terpValue) Float() float64 {
	panic("cant yet")
}
func (t terpValue) Int() int64 {
	panic("cant yet")
}
func (t terpValue) Uint() uint64 {
	panic("cant yet")
}
func (t terpValue) IsPreservedByList() bool { return true }
func (t terpValue) List() []T {
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
			panic(Sprintf("Slice of Uint8 should not be in terpValue: %q", string(t.v.Interface().([]byte))))
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
func (t terpValue) HeadTail() (hd, tl T) {
	return MkList(t.List()).HeadTail()
}

func NeedsOctalEscape(b byte) bool {
	switch b {
		case '{' : return true
		case '}' : return true
		case '\\' : return true
	}
	if b < ' ' {return true}
	return false
}

func OctalEscape (s string) string {
	needsEscaping := false
	for _, b := range []byte (s) {
		if NeedsOctalEscape(b) {
			needsEscaping = true
			break
		}
	}
	if !needsEscaping {return s}
	buf := bytes.NewBuffer(nil)
	for _, b := range []byte (s) {
		if NeedsOctalEscape(b) {
			buf.WriteString(Sprintf("\\%03o",b))
		} else {
			buf.WriteByte(b)
		}
	}
	return(buf.String())
}

func ToListElementString(s string) string {
	// TODO: Not perfect, but we are not doing \ yet.
	// TODO: Broken for mismatched {}.
	if s == "" {
		return "{}"
	}

	if strings.ContainsAny(s, " \t\n\r{}\\") {
		return "{" + OctalEscape(s) + "}"
	}
	return s
}
func (t terpValue) Hash() Hash {
	panic("A GoValue is not a Hash")
}
func (t terpValue) QuickReflectValue() R.Value { return t.v }

func (g *Global) MintMixinSerial() int {
     g.Mu.Lock()
     defer g.Mu.Unlock()

	 g.MixinSerial++
	 return g.MixinSerial
}

type EnsembleItem struct {
	Name	string
	Cmd		Command	
	Doc		string
}

func MkEnsemble(items []EnsembleItem) Command {
	cmd := func(fr *Frame, argv []T) T {
		switch len(argv) {
		case 0:
			panic("TODO: doc string")
		case 1:
			panic("Ensemble command needs a subcommand argument: " + Showv(argv))
		}
		subName := argv[1].String()
		for _, e := range items {
			if e.Name == subName {
				return e.Cmd(fr, argv[1:])
			}
		}
		panic("Ensemble subcommand not found: " + Showv(argv))
	}
	return cmd
}

func NonEmpty(v []string) []string {
	z := make([]string, 0, len(v))
	for _, e := range v {
		if len(e) > 1 {
			z = append(z, e)
		}
	}
	return z
}
