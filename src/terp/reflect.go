package terp

import (
	"errors"
	. "fmt"
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
	return nil
}

func (fr *Frame) initReflect() {

	TBuiltins["lspkg"] = newcmd(cmdLsPkg)

	TBuiltins["peek"] = newcmd(cmdPeek)
	TBuiltins["type"] = newcmd(cmdType)
	TBuiltins["kindT"] = newcmd(cmdKindT)
	TBuiltins["kind"] = newcmd(cmdKind)
	TBuiltins["value"] = newcmd(cmdValue)
	TBuiltins["zeroT"] = newcmd(cmdZeroT)
	TBuiltins["anyV"] = newcmd(cmdAnyV)

	TBuiltins["funcX"] = newcmd(cmdFuncX)
	TBuiltins["typeX"] = newcmd(cmdTypeX)
	TBuiltins["call"] = newcmd(cmdCall)
	TBuiltins["send"] = tcmdSend
	TBuiltins["elem"] = tcmdElem
	TBuiltins["index"] = tcmdIndex
	TBuiltins["tolist"] = tcmdToList // Hack

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

func AdaptToValue(a T, t R.Type) R.Value {
	switch t.Kind() {
	case R.Bool:
		if TTruth(a) {
			return R.ValueOf(true)
		}
		return R.ValueOf(false)
    case R.Int:
		return R.ValueOf(int(a.Int()))
    case R.Int8:
		return R.ValueOf(int8(a.Int()))
    case R.Int16:
		return R.ValueOf(int16(a.Int()))
    case R.Int32:
		return R.ValueOf(int32(a.Int()))
    case R.Int64:
		return R.ValueOf(a.Int())
    case R.Uint:
		return R.ValueOf(uint(a.Uint()))
    case R.Uint8:
		return R.ValueOf(uint8(a.Uint()))
    case R.Uint16:
		return R.ValueOf(uint16(a.Uint()))
    case R.Uint32:
		return R.ValueOf(uint32(a.Uint()))
    case R.Uint64:
		return R.ValueOf(a.Uint())
    case R.Uintptr:
		return R.ValueOf(uintptr(a.Uint()))
    case R.Float32:
		return R.ValueOf(uint32(a.Float()))
    case R.Float64:
		return R.ValueOf(a.Float())
    case R.Complex64:
    case R.Complex128:
    case R.Array:
    case R.Chan:
    case R.Func:
		v := R.ValueOf(a)
		if v.Kind() == R.Func {
			return v
		}
    case R.Interface:
    case R.Map:
    case R.Ptr:
    case R.Slice:
    case R.String:
    case R.Struct:
    case R.UnsafePointer:
	}
	panic(Sprintf("cannot AdaptToValue from %T to %s", a, t))
}

func tcmdSend(fr *Frame, argv []T) T {
	t, meth := TArgv2(argv)

	tv := t.(Tv)

	log.Printf("tv.v = %#v", tv.v)
	log.Printf("tv.v.Type() = %s", tv.v.Type())
	log.Printf("tv.v.Kind() = %s", tv.v.Kind())
	log.Printf("meth.String() = %#v", meth.String())
	m := tv.v.MethodByName(meth.String())
	log.Printf("m = %#v", m)

	vals := m.Call([]R.Value {})

	return MkT(vals[0].Interface())
}

func tcmdElem(fr *Frame, argv []T) T {
	p := TArgv1(argv)

	tv := p.(Tv)
	e := tv.v.Elem()

	return MkT(e.Interface())
}

func tcmdIndex(fr *Frame, argv []T) T {
	slice, i := TArgv2(argv)

	tv := slice.(Tv)
	z := tv.v.Index(int(i.Int()))

	return MkT(z.Interface())
}

func cmdCall(fr *Frame, argv List) Any {
	a := argv[1]

	if s, ok := a.(string); ok {
		ext := Extern(s)
		if ext != nil {
			a = ext
		}
	}

	ty := R.TypeOf(a)
	if ty.Kind() != R.Func {
		panic("argv1 not Func: " + Repr(a))
	}
	nin := ty.NumIn()
	nout := ty.NumOut()
	log.Printf("NumIn=%d NumOut=%d", nin, nout)
	for i := 0; i < nin; i++ {
		log.Printf("Type expect in[%d] : <%s> %s", i, ty.In(i).Kind(), ty.In(i))
	}

	pp := make([]R.Value, 0, 4)
	for i, p := range argv[2:] {
		if p == nil || p == "" {
			// pp = append(pp, R.Zero(R.TypeOf(p)))
			// TODO fix for ...
			pp = append(pp, R.Zero(ty.In(i)))
		} else {
			pp = append(pp, R.ValueOf(p))
		}
	}

	for i, ai := range argv[2:] {
		if ai == nil {
			log.Printf("Type actual in[%d] : nil  %T", i, ai)
		} else {
			log.Printf("Type actual in[%d] : <%s> %T = %#v", i, R.TypeOf(ai).Kind(), ai, ai)
		}
		if pp[i].Kind() == R.Interface {
			log.Printf("............. <%s> %s  <%x @ %x>", pp[i].Kind(), pp[i], pp[i].InterfaceData()[0], pp[i].InterfaceData()[1]) 
		} else {
			log.Printf("............. <%s> %s", pp[i].Kind(), pp[i]) 
		}
	}
	for i := 0; i < nout; i++ {
		log.Printf("Type out[%d] : <%s>  %s", i, ty.Out(i).Kind(), ty.Out(i))
	}

	log.Printf("...(calling)...  %#v  (  %#v  )", a, pp)
	xx := R.ValueOf(a).Call(pp)
	log.Printf("...(called)...")

	zz := make(List, 0, len(xx))
	for i, x := range xx {
		z := x.Interface()
		log.Printf("Result out[%d] : %T = %#v", i, z, z)
		zz = append(zz, z)
	}

	errorI := R.TypeOf(errors.New).Out(0)
	
	// if nout > 0 && ty.Out(nout-1).Name() == "error" #
	if nout > 0 && ty.Out(nout-1).Implements(errorI) {
		log.Printf("Last result implements error; checking it: %#v", zz[nout-1])
		if zz[nout-1] != nil {
			panic(zz[nout-1])
		}
		log.Printf("Last result implements error; was nil; dropping it.")
		zz = zz[:nout-1]
	}
	switch len(zz) {
	case 0:
		return nil
	case 1:
		return zz[0]
	}
	return zz
}

func cmdLsPkg(fr *Frame, argv List) Any {
	switch len(argv) {
	case 1 + 1:
		//key := Str(Argv1(argv))
		//return G.Members[key]
	case 0 + 1:
		//for k, v := range G.Members {
		//	log.Printf("KEY: %#v\n", k)
		//	log.Printf("VALUE: %#v\n", v)
		//}
		//return len(G.Members)
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

func tcmdToList(fr *Frame, argv []T) T {
	a := TArgv1(argv)

	v := a.(Tv).v
	n := v.Len()
	z := []T {}

	for i := 0; i < n; i++ {
		z = append(z, MkT(v.Index(i).Interface()))
		log.Printf("tcmdToTList: z <- %s", z)
	}

	log.Printf("tcmdToList: MkTl from <%T> <%s>", z, z)
	return MkTl(z)
}
