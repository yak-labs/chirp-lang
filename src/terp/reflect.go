package terp

import (
	//. "fmt"
	"log"
	R "reflect"
	//"strconv"

	G "generated"
)

func Extern(name string) Any {
	if len(name) == 0 || name[0] != '/' {
		return nil
	}
	if f, ok := G.Funcs[name]; ok {
		return f
	}
	if t, ok := G.Types[name]; ok {
		return R.TypeOf(t).Elem()
	}
	if m, ok := G.Members[name]; ok {
		return m
	}
	return nil
}

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
	f := G.Funcs[a]
	return R.ValueOf(f)
}

func cmdTypeX(fr *Frame, argv List) Any {
	a := Str(Argv1(argv))
	x := G.Types[a]
	return R.TypeOf(x).Elem()
}

func cmdCall(fr *Frame, argv List) Any {
	a := argv[1]

	if s, ok := a.(string); ok {
		ext := Extern(s)
		if ext != nil {
			a = ext
		}
	}

	pp := make([]R.Value, 0, 4)
	for _, p := range argv[2:] {
		pp = append(pp, R.ValueOf(p))
	}
	if R.ValueOf(a).Kind() == R.Func {
		xx := R.ValueOf(a).Call(pp)
		zz := make(List, 0, len(xx))
		for _, x := range xx {
			zz = append(zz, x.Interface())
		}
		return zz
	}
	panic("argv1 not Func: " + Repr(a))
}

func cmdLsPkg(fr *Frame, argv List) Any {
	switch len(argv) {
	case 1 + 1:
		key := Str(Argv1(argv))
		return G.Members[key]
	case 0 + 1:
		for k, v := range G.Members {
			log.Printf("KEY: %#v\n", k)
			log.Printf("VALUE: %#v\n", v)
		}
		return len(G.Members)
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
