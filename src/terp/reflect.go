package terp

import (
	"errors"
	. "fmt"
	"log"
	R "reflect"

	G "generated"
)

var errorInterfaceType R.Type = R.TypeOf(errors.New).Out(0)

func findExternalGoFunctionAsValue(name string) R.Value {
	if len(name) < 2 || name[0] != '/' {
		panic("Doesnt start with /: " + Repr(name))
	}
	var f interface{}
	var ok bool
	if f, ok = G.Funcs[name]; !ok {
		panic("External command not found in G.Funcs" + Repr(name))
	}
	z := R.ValueOf(f)
	Must(R.Func, z.Kind())
	return z
}

func (fr *Frame) initReflect() {
	TBuiltins["call"] = tcmdCall
	TBuiltins["send"] = tcmdSend
	TBuiltins["get"] = tcmdGet
	TBuiltins["set"] = tcmdSet

	TBuiltins["elem"] = tcmdElem
	TBuiltins["index"] = tcmdIndex
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
		if a.IsEmpty() {
			log.Printf("AdaptToValue: Nil for Chan (%s), due to IsEmpty.", t)
			return R.Zero(t)
		}
	case R.Func:
		if a.IsEmpty() {
			log.Printf("AdaptToValue: Nil for Func (%s), due to IsEmpty.", t)
			return R.Zero(t)
		}
		v := R.ValueOf(a)
		if v.Kind() == R.Func {
			return v
		}
	case R.Interface:
		if a.IsEmpty() {
			log.Printf("AdaptToValue: Nil for Interface (%s), due to IsEmpty.", t)
			return R.Zero(t)
		}
	case R.Map:
	case R.Ptr:
		if a.IsEmpty() {
			log.Printf("AdaptToValue: Nil for Ptr (%s), due to IsEmpty.", t)
			return R.Zero(t)
		}
	case R.Slice:
	case R.String:
	case R.Struct:
	case R.UnsafePointer:
	}
	// We haven't checked this is correct;
	//  cmdCall will reject it, if it won't work.
	// But maybe we can do better, so log it.
	log.Printf("AdaptToValue: Default: for type <%s>: %s", t, Show(a))
	return R.ValueOf(a.Raw())
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

func tcmdCall(fr *Frame, argv []T) T {
	funcName := argv[1].String()
	log.Printf("Call fn=%s  len(argv)=%d", funcName, len(argv))

	fn := findExternalGoFunctionAsValue(funcName)
	return commonCall(fr, funcName, fn, argv[2:])
}

func tcmdSend(fr *Frame, argv []T) T {
	t, meth, args := TArgv2v(argv)
	log.Printf("----send----receiver = %s", Show(t))
	methName := meth.String()
	log.Printf("----send----method = %q", methName)

	tv, ok := t.(Tv)
	if !ok {
		panic(Sprintf("'send' command expected Tv receiver, but got %s", Show(t)))
	}

	fn := tv.v.MethodByName(methName)
	return commonCall(fr, "Method:"+methName+":"+tv.v.Type().String(), fn, args)
}

// commonCall is common to both "call" and "send".
// args already has function name or receiver and message name stripped; it's just the args.
func commonCall(fr *Frame, funcName string, fn R.Value, args []T) T {
	ty := fn.Type()

	log.Printf("... fn <%s> type: <%s> %s", funcName, ty.Kind(), ty)

	nin := ty.NumIn()
	nout := ty.NumOut()
	vari := ty.IsVariadic() // Last arg has "..." and presents as slice type

	// Log the function's expected arguments.
	for i := 0; i < nin; i++ {
		log.Printf("... fn expect in[%d] : <%s> %s", i, ty.In(i).Kind(), ty.In(i))
	}
	for i := 0; i < nout; i++ {
		log.Printf("... fn returns out[%d] : <%s>  %s", i, ty.Out(i).Kind(), ty.Out(i))
	}

	// Check num args
	if vari {
		if len(args) < nin-1 {
			panic(Sprintf("command <%s> expected at least %d args but got %d.", funcName, nin-1, len(args)))
		}
	} else {
		if len(args) != nin {
			panic(Sprintf("command <%s> expected %d args but got %d.", funcName, nin, len(args)))
		}
	}

	// Convert actual args to Values, into pp. 
	pp := make([]R.Value, len(args))
	for i, p := range args {
		var target R.Type
		if vari && i >= nin-1 {
			target = ty.In(nin - 1).Elem() // element type of variadic args...t
		} else {
			target = ty.In(i)
		}
		pp[i] = AdaptToValue(p, target)
		log.Printf("........ passing [%d] %s as (%s) %s", i, Show(p), pp[i].Kind().String(), pp[i].Type().String())

	}

	// Call it.
	log.Printf("...(calling)...  %#v  (  %#v  )", funcName, pp)
	xx := fn.Call(pp)
	log.Printf("...(called)...")

	// Convert results from Values into zz.
	zz := make([]T, len(xx))
	for i, x := range xx {
		z := x.Interface()
		zz[i] = MkT(z)
		log.Printf("........ result [%d] %s from (%s) %s", i, Show(zz[i]), x.Kind().String(), x.Type().String())
	}

	// If last arg is type error, remove it if nil, or panic it.
	if nout > 0 && ty.Out(nout-1).Implements(errorInterfaceType) {
		log.Printf("Last result implements error; checking it: %#v", zz[nout-1])
		if !xx[nout-1].IsNil() {
			panic(Sprintf("call %q returns error: %q", funcName, xx[nout-1].Interface().(error).Error()))
		}
		log.Printf("Last result implements error; was nil; dropping it.")
		zz = zz[:nout-1]
	}

	// Decide how to return, based on number of results.
	switch len(zz) {
	case 0:
		return Empty // If no result, return the Tcl Empty
	case 1:
		return zz[0] // If single result, return it simply.
	}
	return MkTl(zz) // If multiple results, return a list of them.
}

func derefChain(fr *Frame, chain []T) R.Value {
	Must(true, len(chain) > 0)
	// First name in chain is starting variable name.
	start := chain[0].String()

	var av R.Value
	if len(start) > 0 && start[0] == '/' {
		// Case Go Var
		ptr, ok := G.Vars[start]
		if !ok {
			panic(Sprintf("Cannot find std lib Go var %q", start))
		}
		av = R.ValueOf(ptr).Elem()
	} else {
		// Case Tcl Var
		a := fr.TGetVar(start)
		av = R.ValueOf(a)
	}

	// For additional names, use Fields (or other navigation) to deref.
	for _, e := range chain[1:] {
		log.Printf("------DEREF <<< type(av)=<%s>%s", av.Kind(), av.Type())
		if av.Kind() == R.Ptr || av.Kind() == R.Interface {
			log.Printf("------DEREF %s", av.Kind())
			av = av.Elem()
		}
		log.Printf("------DEREF BY %q", e.String())
		av2 := av.FieldByName(e.String())
		log.Printf("------DEREF >>> type(av)=<%s>%s", av2.Kind(), av2.Type())
		if ! av2.IsValid() {
			// Better: ShowValue(av)
			panic(Sprintf("invalid field %s of %s", e.String(), Show(MkT(av.Interface()))))
		}
		av = av2
	}
	return av
}

func tcmdGet(fr *Frame, argv []T) T {
	z := derefChain(fr, argv[1:]).Interface()

	// We might have a T, or we might have some other Go value.
	if zt, ok := z.(T); ok {
		return zt
	}
	return MkT(z)
}

func tcmdSet(fr *Frame, argv []T) T {
	n := len(argv)
	switch n {
	case 0, 1, 2:
		panic(Sprintf("set command expects at least two args, got %q", Showv(argv)))
	case 3:
		name, x := TArgv2(argv)
		start := name.String()
		log.Printf(".... start=%q", start)
		if len(start) == 0 || start[0] != '/' {
			log.Printf(".... TSetVar %q %s", start, Show(x))
			fr.TSetVar(name.String(), x)
			return x
		}
	}

	// Case 4 or more:
	loc := derefChain(fr, argv[1:n-1])  // Location to assign to.
	if !loc.CanSet() {
		panic(Sprintf("set command deref'ed to an Unsetable Location: %q", Showv(argv)))
	}

	x := argv[n-1] // The value to be assigned.
	zv := AdaptToValue(x, loc.Type())
	log.Printf(".... Reflect Set loc <%s> %s = %s", zv.Kind(), zv.Type(), Show(x))
	loc.Set(zv)
	return x
}
