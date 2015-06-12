package chirp

import (
	"bytes"
	. "fmt"
	"go/ast"
	"log"
	"os"
	"path"
	R "reflect"
	"runtime"
	"strings"
	"sync"
)

var Debug [256]bool

type Shower interface {
	Show() string
}

type Hash map[string]T

type Command func(fr *Frame, argv []T) T

type Scope map[string]Loc

type CmdScope map[string]*CmdNode

// CmdNode makes a singly-linked-list of commands
// at different mixin levels, highest level first.
// A non-mixin command has level 0 and only one CmdNode.
type CmdNode struct {
	Fn         Command
	MixinLevel int
	MixinName  string
	Next       *CmdNode
}

// Frame is a local variable frame.
// There is one for the global variables in the Global struct,
// and a new one is created for each proc or yproc invocation
// (but not for every Command; non-proc commands do not make Frames).
type Frame struct {
	Mu   sync.Mutex
	Vars Scope // local variables
	Cred Hash  // credentials

	Prev *Frame
	G    *Global

	WriterChan chan<- Either // for yproc & yield
	MixinLevel int
	MixinName  string
}

// Global holds the global state of an interpreter,
// mainly the Commands and global variables.
// It also knows if a Mixin is being defined.
// Mixins should be defined by main thread,
// after all overridable procs are defined,
// but before other goroutines start.
type Global struct {
	Cmds CmdScope
	Fr   Frame // global scope

	MixinSerial         int    // Increment before defining Mixin.
	MixinNumberDefining int    // Set nonzero while defining Mixin.
	MixinNameDefining   string // Set nonempty while defining Mixin.
	IsSafe              bool   // Set true for safe subinterpreter.
	Logger              *log.Logger
	Verbosity           int    // Log if message level <= verbosity.
	LogName             string // for logging

	Mu sync.Mutex
}

// Clone produces a copy of the receiving interpreter.
func (g *Global) Clone() *Global {
	z := &Global{
		Cmds: make(CmdScope),
		Fr: Frame{
			Vars: make(Scope),
		},
	}

	z.Fr.G = z

	// Lock Frame before Global.
	// That is, start low at leaves in tree, lock upward towards root.
	// TODO:  Why a g.Mu, instead of it's g.Fr.Mu?
	g.Fr.Mu.Lock()
	defer g.Fr.Mu.Unlock()

	g.Mu.Lock()
	defer g.Mu.Unlock()

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
	USAGE // New in chirp; not in Tcl.
)

// Jump structs are panicked for return, break, and continue.
type Jump struct {
	Status StatusCode
	Result T
}

// Either Bad or Good value.
type Either struct {
	Bad  interface{}
	Good T
}

// Loc is protocol for a variable location.
type Loc interface {
	Has() bool
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

type BitsWord uint32 // We cannot fit uint64 into the float -- until we support actual uint64, we must use shorter BitsWords.

var TypeT = R.TypeOf(MkT(""))
var TypeType = R.TypeOf(TypeT)

var Empty = MkString("")
var InvalidValue = *new(R.Value)

// Create a new interpreter, and return the global frame pointer.
func newEitherInterpreter(isSafe bool) *Frame {
	g := &Global{
		Cmds: make(CmdScope),
		Fr: Frame{
			Vars: make(Scope),
		},
		IsSafe: isSafe,
	}

	g.Fr.G = g

	// Copy Safes to commands.
	for k, v := range Safes {
		node := CmdNode{Fn: v}
		g.Cmds[k] = &node
	}

	if !isSafe {
		// In unsafe terp, copy Unsafes to commands.
		for k, v := range Unsafes {
			node := CmdNode{Fn: v}
			g.Cmds[k] = &node
		}
	}

	return &g.Fr
}

// NewInterpreter() makes a new full interpreter.
func NewInterpreter() *Frame {
	return newEitherInterpreter(false)
}

// NewSafeInterpreter() makes a new safe interpreter.
func NewSafeInterpreter() *Frame {
	return newEitherInterpreter(true)
}

// NewFrame makes a frame for calling another proc.
func (fr *Frame) NewFrame() *Frame {
	NewFrameCounter.Incr()
	return &Frame{
		Vars: make(Scope), // new local var scope
		Cred: fr.Cred,     // same credentials as caller
		Prev: fr,          // link back to prev frame
		G:    fr.G,        // the Global struct
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

func (p *Slot) Has() bool { return p.Elem != nil }
func (p *Slot) Get() T    { return p.Elem }
func (p *Slot) Set(t T)   { p.Elem = t }

func (fr *Frame) GetVarScope(name string) Scope {
	if len(name) == 0 {
		panic("Empty variable name")
	}

	if IsGlobal(name) {
		return fr.G.Fr.Vars
	}
	return fr.Vars
}

func (fr *Frame) HasVar(name string) bool {
	sc := fr.GetVarScope(name)
	var loc Loc
	var ok bool
	fr.Mu.Lock()
	loc, ok = sc[name]
	fr.Mu.Unlock()
	if !ok {
		return false
	}
	return loc.Has()
}

func (fr *Frame) GetVar(name string) T {
	sc := fr.GetVarScope(name)
	var loc Loc
	var ok bool
	fr.Mu.Lock()
	loc, ok = sc[name]
	fr.Mu.Unlock()
	if !ok {
		panic(Sprintf("Variable %q does not exist; scope contains %v", name, sc))
	}
	return loc.Get()
}

func (fr *Frame) SetVar(name string, x T) {
	if strings.Contains(name, ",") {
		// Support destructuring list assignment syntax.
		xs := x.List()
		names := strings.Split(name, ",")
		for i, n := range names {
			if len(n) > 0 {
				if i < len(xs) {
					fr.SetVar(n, xs[i])
				} else {
					fr.SetVar(n, Empty) // Missing values become empty.
				}
			}
		}
		return
	}
	sc := fr.GetVarScope(name)
	fr.Mu.Lock()
	ptr := sc[name]
	if ptr == nil {
		ptr = new(Slot)
		sc[name] = ptr
	}
	fr.Mu.Unlock()
	ptr.Set(x)
}

func (p *UpSlot) Has() bool { return p.Fr.HasVar(p.RemoteName) }
func (p *UpSlot) Get() T    { return p.Fr.GetVar(p.RemoteName) }
func (p *UpSlot) Set(t T)   { p.Fr.SetVar(p.RemoteName, t) }

func (fr *Frame) DefineUpVar(name string, remFr *Frame, remName string) {
	sc := fr.GetVarScope(name)
	fr.Mu.Lock()
	sc[name] = &UpSlot{Fr: remFr, RemoteName: remName}
	fr.Mu.Unlock()
}

func (fr *Frame) FindCommand(name T, callSuper bool) Command {
	// TODO: Optimize with terpMulti.
	cmdName := name.String()

	var fn Command
	fr.Mu.Lock()
	cmdNode, ok := fr.G.Cmds[cmdName]
	fr.Mu.Unlock()
	if ok {
		if callSuper {
			maxMixinLevel := fr.MixinLevel - 1
			if maxMixinLevel < 0 {
				panic("cannot callSuper from non-mixin")
			}
			for cmdNode != nil && cmdNode.MixinLevel > maxMixinLevel {
				cmdNode = cmdNode.Next
			}
		}
		if cmdNode == nil {
			ok = false
		} else {
			fn = cmdNode.Fn
		}
	}

	// Mixin Local Commands:
	if !ok && fr.MixinLevel > 0 && !IsGlobal(cmdName) {
		// Use long name for mixin local fn.
		localCmdName := fr.MixinName + "~" + cmdName
		var localNode *CmdNode
		fr.Mu.Lock()
		localNode, ok = fr.G.Cmds[localCmdName]
		fr.Mu.Unlock()
		if ok {
			fn = localNode.Fn // Should be singleton.
		}
	}

	if !ok {
		return nil
	}
	return fn
}

// Apply a command with its arguments.
func (fr *Frame) Apply(argv []T) T {
	if Debug['a'] {
		Sayf("Apply:  <%q>", argv[0])
		for i, e := range argv[1:] {
			Sayf(".....:  arg%d=<%q>", i+1, e)
		}
	}

	defer func() {
		if r := recover(); r != nil {
			if re, ok := r.(error); ok {
				r = re.Error() // Convert error to string.
			}
			if rs, ok := r.(string); ok {
				rs = rs + Sprintf("\n\tin Apply\n\t\t%q", argv[0])

				// TODO: Require debug level for the args.
				for _, ae := range argv[1:] {
					as := ae.String()
					if len(as) > 40 {
						as = as[:40] + "..."
					}
					rs = rs + Sprintf(" %q", as)
				}

				if fr.MixinLevel > 0 {
					rs = rs + Sprintf("\n\t\t(frame's MixinLevel=%d)", fr.MixinLevel)
				}
				if len(fr.MixinName) > 0 {
					rs = rs + Sprintf("\n\t\t(frame's MixinName=%q)", fr.MixinName)
				}
				r = rs
			}
			panic(r)
		}
	}()

	// First try to find the Command function.
	head := argv[0]
	fn := fr.FindCommand(head, false) // false: Don't call super.
	if fn != nil {
		// Found it; use it.
		z := fn(fr, argv)
		if Debug['a'] {
			Sayf("Apply...returns <%q>", z.String())
		}
		return z
	}

	// Next try to find a root.
	rootName := head.String()
	if len(rootName) > 0 && rootName[0] == '/' {
		return LookupRootAndApply(fr, rootName, argv)
	}

	panic(Sprintf("No such command: %q", rootName))
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
		log.Printf("[A]: %s", Repr(a))
		log.Printf("[B]: %s", Repr(b))
		panic(Repr(a) + " .vs. " + Repr(b))
	}
}

// MustB takes two bytes.
func MustB(a, b byte) {
	if a != b {
		panic(Sprintf("MustB Fails: %d %q .vs. %d %q", a, string(a), b, string(b)))
	}
}

// MustSp takes Any 2 values, and compares their Repr()s, without spaces.
func MustNoSp(a, b interface{}) {
	if DropSpaces(Repr(a)) != DropSpaces(Repr(b)) {
		log.Printf("[A]: %s", DropSpaces(Repr(a)))
		log.Printf("[B]: %s", DropSpaces(Repr(b)))
		panic(Repr(a) + " .vs. " + Repr(b))
	}
}
func DropSpaces(s string) string {
	z := ""
	for _, c := range s {
		if c != ' ' {
			z += string(c)
		}
	}
	return z
}

func Show(a T) string {
	if a == nil {
		return "{(T)nil}"
	}
	return Sprintf("{(%T) ## %#v ## %q}", a, a, a.String())
}

func Showv(a []T) string {
	buf := bytes.NewBufferString(Sprintf("Slice of T with %d elements:", len(a)))
	for i, e := range a {
		buf.WriteString(Sprintf("\n    ... [%d] = %s", i, Show(e)))
	}
	return buf.String()
}

func Where() string {
	sb := bytes.NewBuffer(nil)
	for skip := 6; skip > 1; skip-- {
		_, file, line, ok := runtime.Caller(skip)
		if ok {
			base := path.Base(file)
			n := len(base)
			if base[n-3:] == ".go" {
				base = base[:n-3] // Strip trailing .go
			}
			sb.WriteString(Sprintf(" %s:%d", base, line))
		}
	}
	return sb.String()
}

func Logf(fmt string, args ...interface{}) {
	log.Println(Sprintf(fmt, args...))
}

var SayPrefix = "Say" // TODO: Global Prefix WILL BREAK WITH goROUTINES.

// Quick internal logging function that needs no Frame.
func Say(args ...interface{}) {
	savedPrefix := SayPrefix
	defer func() { SayPrefix = savedPrefix }()

	if len(SayPrefix) < 4 {
		log.Println(Sprintf("%s --->%s --->", SayPrefix, Where()))
	}

	prefix := " :::"
	for _, a := range args {
		switch t := a.(type) {
		case Shower:
			log.Println(Sprintf("%s %s", SayPrefix, t.Show()))
		case terpValue:
			view := Sprintf("%v", t.v.Interface())
			if len(view) > 80 {
				view = view[:80] + "..."
			}
			targetCanSet := " "
			targAddress := "t@?"
			switch t.v.Kind() {
			case R.Ptr, R.Interface:
				targetCanSet = Sprintf("targCanSet=%v", t.v.Elem().CanSet())
				if t.v.Elem().CanAddr() {
					targAddress = Sprintf("t@%x@%x", t.v.Elem().Addr(), t.v.Pointer())
				}
			}
			address := "@?"
			if t.v.CanAddr() {
				address = Sprintf("@%x", t.v.Addr())
			}
			log.Println(Sprintf("%s terpVALUE{{{CanSet=%v %s %s %s :::%s:::%s::: %s:::%#v}}}", SayPrefix, t.v.CanSet(), targetCanSet, address, targAddress, t.v.Kind(), t.v.Type(), view, t.v))
		case R.Value:
			view := Sprintf("%v", t.Interface())
			if len(view) > 80 {
				view = view[:80] + "..."
			}
			targetCanSet := " "
			targAddress := "t@?"
			switch t.Kind() {
			case R.Ptr, R.Interface:
				targetCanSet = Sprintf("targCanSet=%v", t.Elem().CanSet())
				if t.Elem().CanAddr() {
					targAddress = Sprintf("t@%x@%x", t.Elem().Addr(), t.Pointer())
				}
			}
			address := "@?"
			if t.CanAddr() {
				address = Sprintf("@%x", t.Addr())
			}
			log.Println(Sprintf("%s VALUE{{{CanSet=%v %s %s %s :::%s:::%s::: %s:::%#v}}}", SayPrefix, t.CanSet(), targetCanSet, address, targAddress, t.Kind(), t.Type(), view, t))
		default:
			rv := R.ValueOf(a)
			rvt := rv.Type()
			if rvt.Kind() == R.Slice {
				log.Println(Sprintf("%s SLICE [%d] %s ........", SayPrefix, rv.Len(), rvt))
				for i := 0; i < rv.Len(); i++ {
					SayPrefix = savedPrefix + Sprintf("SLICE [%d]: ", i)
					elem := rv.Index(i)
					Say(elem.Interface())
				}
			} else {
				log.Println(Sprintf("%s %#v", prefix, a))
			}
		}
	}
}
func Sayf(format string, args ...interface{}) {
	log.Println(Sprintf(format, args...))
}

func SetDebugFromEnv() {
	letters := os.Getenv("CHIRP_DEBUG")
	for _, ch := range letters {
		Debug[ch] = true
	}
}

type Counter struct {
	count int64
	name  string
	next  *Counter
}

var Counters *Counter

func (p *Counter) Incr() {
	p.count++
}

func (p *Counter) Show() string {
	return Sprintf("%d %s", p.count, p.name)
}

func (p *Counter) Register(name string) {
	p.name = name
	p.next = Counters
	Counters = p
}

func LogAllCounters() {
	for p := Counters; p != nil; p = p.next {
		Logf("Counter: %s", p.Show())
	}
}

func ClearAllCounters() {
	for p := Counters; p != nil; p = p.next {
		p.count = 0
	}
}

func ShowAllCounters() string {
	buf := bytes.NewBuffer(nil)
	for p := Counters; p != nil; p = p.next {
		buf.WriteString(p.Show())
		buf.WriteByte('\n')
	}
	return buf.String()
}

var NewFrameCounter Counter

func init() {
	NewFrameCounter.Register("NewFrame")
}
