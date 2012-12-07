package terp

import (
	. "fmt"
	"sync"
)

type Any	interface{}

type List	[]Any
type Dict	map[string]Any

type Command func (*Terp, ...Any) Any

type Scope map[string]Any

type CmdScope map[string]Command

type Frame struct {
	Vars	Scope
	Prev	*Frame
}

type Terp struct {
	Top	*Frame		// current frame
	Mu	sync.Mutex
	Cmds	CmdScope
	Fr	Frame		// global scope
}

func Repr(a Any) string { return Sprintf("%#v", a) }

func Must(a, b Any, extra ...Any) {
	if a != b {
		panic(Repr(a) + " vs. " + Repr(b) + " :: " + Repr(extra))
	}
}
