package terp

import (
	//. "fmt"
	"log"
	"reflect"
	//"strconv"

	"generated"
)

var ExtTypes = generated.Types
var ExtMembers = generated.Members

func (fr *Frame) initExterns() {
	Builtins["lspkg"] = cmdLsPkg

	Builtins["peek"] = cmdPeek
	Builtins["type"] = cmdType
	Builtins["kindT"] = cmdKindT
	Builtins["kind"] = cmdKind
	Builtins["value"] = cmdValue
	Builtins["zeroT"] = cmdZeroT
	Builtins["anyV"] = cmdAnyV
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

func cmdPeek(fr *Frame, argv List) Any {
	a := Argv1(argv)
	return Repr(a)
}

func cmdType(fr *Frame, argv List) Any {
	a := Argv1(argv)
	return reflect.TypeOf(a)
}

func cmdKindT(fr *Frame, argv List) Any {
	a := Argv1(argv)
	t := a.(reflect.Type)
	return t.Kind().String()
}

func cmdKind(fr *Frame, argv List) Any {
	a := Argv1(argv)
	return reflect.ValueOf(a).Kind().String()
}

func cmdValue(fr *Frame, argv List) Any {
	a := Argv1(argv)
	return reflect.ValueOf(a)
}

func cmdZeroT(fr *Frame, argv List) Any {
	a := Argv1(argv)
	t := a.(reflect.Type)
	return reflect.Zero(t)
}

func cmdAnyV(fr *Frame, argv List) Any {
	a := Argv1(argv)
	v := a.(reflect.Value)
	return v.Interface()
}
