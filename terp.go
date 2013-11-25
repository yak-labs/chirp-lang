package chirp

import (
	"bytes"
	. "fmt"
	"go/ast"
	"log"
	"path"
	R "reflect"
	"runtime"
	"strings"
	"sync"
)

type Shower interface {
	Show() string
}

type Hash map[string]T // TODO: Mutex

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
	Vars Scope // local variables
	Cred Hash  // credentials

	Prev *Frame
	G    *Global

	WriterChan chan<- Either // for yproc & yield
	MixinLevel int
	MixinName  string

	Self  *Obj
	Super *Obj
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
	isSafe              bool   // Set true for safe subinterpreter.
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

var Empty = MkString("")
var InvalidValue = *new(R.Value)

// Create a new interpreter, and return the global frame pointer.
func New1(isSafe bool) *Frame {
	g := &Global{
		Cmds: make(CmdScope),
		Fr: Frame{
			Vars: make(Scope),
		},
		isSafe: isSafe,
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

// New() makes a new full interpreter.
func New() *Frame {
	return New1(false)
}

// NewSafe() makes a new safe interpreter.
func NewSafe() *Frame {
	return New1(true)
}

// NewFrame makes a frame for calling another proc.
func (fr *Frame) NewFrame() *Frame {
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

	if name[0] == '_' && fr.Self != nil {
		return fr.Self.Slots
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
	if loc, ok = sc[name]; !ok {
		return false
	}
	return loc.Has()
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
	if sc[name] == nil {
		sc[name] = new(Slot)
	}
	sc[name].Set(x)
}

func (p *UpSlot) Has() bool { return p.Fr.HasVar(p.RemoteName) }
func (p *UpSlot) Get() T    { return p.Fr.GetVar(p.RemoteName) }
func (p *UpSlot) Set(t T)   { p.Fr.SetVar(p.RemoteName, t) }

func (fr *Frame) DefineUpVar(name string, remFr *Frame, remName string) {
	sc := fr.GetVarScope(name)
	sc[name] = &UpSlot{Fr: remFr, RemoteName: remName}
}

func (fr *Frame) FindCommand(name T, callSuper bool) Command {
	// Some day we will not require terpString; for now, it helps debug.
	// TODO: Optimize with terpMulti.
	cmdName := name.String()

	var fn Command
	cmdNode, ok := fr.G.Cmds[cmdName]
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
		localNode, ok = fr.G.Cmds[localCmdName]
		if ok {
			fn = localNode.Fn // Should be singleton.
		}
	}

	if !ok {
		panic(Sprintf("FindCommand: command not found: %q", cmdName))
	}
	return fn
}

// Apply a command with its arguments.
func (fr *Frame) Apply(argv []T) T {
	Sayf("@Apply@   %s", argv[0])
	for i, e := range argv[1:] {
		Sayf("      .%d.  %s", i, e)
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

	head := argv[0]
	fn := fr.FindCommand(head, false) // false: Don't call super.
	z := fn(fr, argv)
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

// Quick internal logging function that needs no Frame.
func Say(args ...interface{}) {
	log.Println(Sprintf("Say --->%s --->", Where()))
	prefix := ": "
	for _, a := range args {
		switch t := a.(type) {
		case Shower:
			log.Println(Sprintf("%s%s", prefix, t.Show()))
		default:
			rv := R.ValueOf(a)
			rvt := rv.Type()
			if rvt.Kind() == R.Slice && rvt.Elem().Kind() == R.Ptr {
				log.Println(Sprintf("%s SLICE [%d] %s ........", prefix, rv.Len(), rvt))
				for i := 0; i < rv.Len(); i++ {
					elem := rv.Index(i)
					Say(Sprintf("...... Index[%d]", i), elem.Interface())
				}
			} else {
				log.Println(Sprintf("%s%#v", prefix, a))
			}
		}
		prefix += "  : "
	}
}
func Sayf(format string, args ...interface{}) {
	log.Println(Sprintf(format, args...))
}
