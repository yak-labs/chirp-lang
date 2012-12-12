package terp

import (
	//. "fmt"
	"log"
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
	case 0 + 1:
		for k, v := range ExtMembers {
			log.Printf("KEY: %#v\n", k)
			log.Printf("VALUE: %#v\n", v)
		}
		return len(generated.Members)
	}
	panic("Bad argv to cmdLsPkg")
}
