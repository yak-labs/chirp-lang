package terp

import (
	//. "fmt"
	//"log"
	//"strconv"

	"generated"
)

var ExtTypes = generated.Types
var ExtMembers = generated.Members

func (fr *Frame) initExterns() {
	Builtins["lspkg"] = cmdLsPkg
}

func cmdLsPkg(fr *Frame, argv List) Any {
	switch len(argv) {
	case 1 + 1:
		key := Str(Argv1(argv))
		return ExtMembers[key]
	}
	panic("Bad argv to cmdLsPkg")
}
