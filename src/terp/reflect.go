package terp

import (
	//. "fmt"
	"log"
	R "reflect"
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

	Builtins["funcX"] = cmdFuncX
	Builtins["typeX"] = cmdTypeX
	Builtins["call"] = cmdCall
}

func cmdFuncX(fr *Frame, argv List) Any {
	a := Str(Argv1(argv))
	f := generated.Funcs[a]
	return R.ValueOf(f)
}

func cmdTypeX(fr *Frame, argv List) Any {
	a := Str(Argv1(argv))
	x := generated.Types[a]
	return R.TypeOf(x).Elem()
}

func cmdCall(fr *Frame, argv List) Any {
	a := argv[1]
	pp := make([]R.Value, 0, 4)
	for _, p := range argv[2:] {
		pp = append(pp, R.ValueOf(p))
	}
	if R.ValueOf(a).Kind() == R.Func {
		xx := R.ValueOf(a).Call(pp)
		zz := make([]Any, 0, len(xx))
		for _, x := range xx {
			zz = append(zz, x.Interface())
		}
		return xx
	}
	panic("argv1 not Func: " + Repr(a))
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
	return R.TypeOf(a)
}

func cmdKindT(fr *Frame, argv List) Any {
	a := Argv1(argv)
	t := a.(R.Type)
	return t.Kind().String()
}

func cmdKind(fr *Frame, argv List) Any {
	a := Argv1(argv)
	return R.ValueOf(a).Kind().String()
}

func cmdValue(fr *Frame, argv List) Any {
	a := Argv1(argv)
	return R.ValueOf(a)
}

func cmdZeroT(fr *Frame, argv List) Any {
	a := Argv1(argv)
	t := a.(R.Type)
	return R.Zero(t)
}

func cmdAnyV(fr *Frame, argv List) Any {
	a := Argv1(argv)
	v := a.(R.Value)
	return v.Interface()
}
