package chirp

import (
	"errors"
	. "fmt"
	"log"
	R "reflect"
	"unsafe"
)

var (
	// Maps to hold bindings to GO API for Reflection.
	RFuncs map[string]interface{}	    // Holds functions.
	RVars map[string]interface{}		// Holds pointers to vars.
	RTypes map[string]interface{}		// Holds a pointer to a new() instance of the type.
	RConsts map[string]interface{}		// Holds a reified constant value.
)

var errorInterfaceType R.Type = R.TypeOf(errors.New).Out(0)

func findExternalGoFunctionAsValue(name string) R.Value {
	if len(name) < 2 || name[0] != '/' {
		panic("Doesnt start with /: " + Repr(name))
	}
	var f interface{}
	var ok bool
	if f, ok = RFuncs[name]; !ok {
		panic("External command not found in RFuncs" + Repr(name))
	}
	z := R.ValueOf(f)
	if z.Kind() != R.Func {
		panic("Element of RFuncs should have been kind Func: " + Repr(name))
	}
	return z
}

func  AdaptToValue(fr *Frame, a T, ty R.Type) R.Value {
	switch ty.Kind() {
	case R.Bool:
		var tmp bool = a.Bool()
		return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()

	case R.Int:
		var tmp int = int(a.Int())
		return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()

	case R.Int8:
		var tmp int8 = int8(a.Int())
		return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()

	case R.Int16:
		var tmp int16 = int16(a.Int())
		return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()

	case R.Int32:
		var tmp int32 = int32(a.Int())
		return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()

	case R.Int64:
		var tmp int64 = a.Int()
		return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()

	case R.Uint:
		var tmp uint = uint(a.Uint())
		return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()

	case R.Uint8:
		var tmp uint8 = uint8(a.Uint())
		return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()

	case R.Uint16:
		var tmp uint16 = uint16(a.Uint())
		return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()

	case R.Uint32:
		var tmp uint32 = uint32(a.Uint())
		return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()

	case R.Uint64:
		var tmp uint64 = a.Uint()
		return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()

	case R.Uintptr:
		var tmp uintptr = uintptr(a.Uint())
		return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()

	case R.Float32:
		var tmp float32 = float32(a.Float())
		return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()

	case R.Float64:
		var tmp float64 = a.Float()
		return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()

	case R.Complex64:
	case R.Complex128:
	case R.Array:
	case R.Chan:
		if a.IsEmpty() {
			log.Printf("AdaptToValue: Nil for Chan (%s), due to IsEmpty.", ty)
			return R.Zero(ty)
		}
	case R.Func:
		if a.IsEmpty() {
			log.Printf("AdaptToValue: Nil for Func (%s), due to IsEmpty.", ty)
			return R.Zero(ty)
		}
		v := R.ValueOf(a)
		if v.Kind() == R.Func {
			return v
		}
	case R.Interface:
		if a.IsEmpty() {
			log.Printf("AdaptToValue: Nil for Interface (%s), due to IsEmpty.", ty)
			return R.Zero(ty)
		}
		// Very special case of T: Return arg as a Value.
		if ty.String() == "terp.T" {
			return R.ValueOf(a)
		}
	case R.Map:
	case R.Ptr:
		if a.IsEmpty() {
			log.Printf("AdaptToValue: Nil for Ptr (%s), due to IsEmpty.", ty)
			return R.Zero(ty)
		}
		// Very special case of *Frame:
		framePtrValue := R.ValueOf(fr)
		framePtrType := framePtrValue.Type()
		if ty == framePtrType {
			return framePtrValue
		}
	case R.Slice:
		raw := a.Raw()
		val := R.ValueOf(raw)

		switch ty.Elem().Kind() {
		case R.Uint8:
			var tmp []byte = make([]byte, val.Len())
			copy(tmp, val.String())
			return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()
		}

	case R.String:
		raw := a.Raw()
		val := R.ValueOf(raw)

		switch val.Kind() {
		case R.Slice:
			switch val.Elem().Kind() {
			case R.Uint8:
				var tmp string = string(val.Bytes())
				return R.NewAt(ty, unsafe.Pointer(&tmp)).Elem()
			}
		}

	case R.Struct:
	case R.UnsafePointer:
	}
	// We haven't checked this is correct;
	//  cmdCall will reject it, if it won't work.
	// But maybe we can do better, so log it.
	log.Printf("AdaptToValue: Default: for type <%s>: %s", ty, Show(a))
	return R.ValueOf(a.Raw())
}

func cmdElem(fr *Frame, argv []T) T {
	p := Arg1(argv)

	rv := p.QuickReflectValue()
	if !rv.IsValid() {
		panic("cannot use 'elem' on non-reflect-value")
	}
	e := rv.Elem()

	return MkT(e.Interface())
}

func cmdIndex(fr *Frame, argv []T) T {
	slice, i := Arg2(argv)

	rv := slice.QuickReflectValue()
	if !rv.IsValid() {
		panic("cannot 'index' on non-reflect-value")
	}
	z := rv.Index(int(i.Int()))

	return MkT(z.Interface())
}

func cmdCall(fr *Frame, argv []T) T {
	funcName := argv[1].String()
	log.Printf("Call fn=%s  len(argv)=%d", funcName, len(argv))

	fn := findExternalGoFunctionAsValue(funcName)
	return commonCall(fr, funcName, fn, argv[2:])
}

func cmdSend(fr *Frame, argv []T) T {
	r, meth, args := Arg2v(argv)
	log.Printf("----send----receiver = %s", Show(r))
	methName := meth.String()
	log.Printf("----send----method = %q", methName)

	rv := r.QuickReflectValue()
	if !rv.IsValid() {
		panic("cannot 'go-send' to non-reflect-value")
	}

	fn := rv.MethodByName(methName)
	return commonCall(fr, "Method:"+methName+":"+rv.Type().String(), fn, args)
}

// commonCall is common to both "go-call" and "go-send".
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
		pp[i] = AdaptToValue(fr, p, target)
		log.Printf("........ passing [%d] %s as (%s) %s", i, Show(p), pp[i].Kind().String(), pp[i].Type().String())

	}

	// Call it.
	log.Printf("...(calling)...  %v  (  %v  )", funcName, pp)
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
	return MkList(zz) // If multiple results, return a list of them.
}

func derefChain(fr *Frame, argv []T) R.Value {
	start, rest := Arg1v(argv)
	av := R.ValueOf(start.Raw())

	// For additional names, use Fields (or other navigation) to deref.
	for _, e := range rest {
		log.Printf("------DEREF <<< type(av)=<%s>%s", av.Kind(), av.Type())

		// TODO: IS IT OKAY TO DEREF Ptr or Interface?
		if av.Kind() == R.Ptr || av.Kind() == R.Interface {
			log.Printf("------EXTRA DEREF %s", av.Kind())
			av = av.Elem()
		}
		log.Printf("------DEREF BY %q", e.String())
		av2 := av.FieldByName(e.String())
		log.Printf("------DEREF >>> type(av)=<%s>%s", av2.Kind(), av2.Type())
		if !av2.IsValid() {
			// Better: ShowValue(av)
			panic(Sprintf("invalid field %s of %s", e.String(), Show(MkT(av.Interface()))))
		}
		av = av2
	}
	return av
}

func cmdGetf(fr *Frame, argv []T) T {
	z := derefChain(fr, argv).Interface()
	return MkT(z)
}

func cmdSetf(fr *Frame, argv []T) T {
	n := len(argv)
	loc := derefChain(fr, argv[:n-1]) // Location to assign to.
	if !loc.CanSet() {
		panic(Sprintf("set command deref'ed to an Unsetable Location: %q", Showv(argv)))
	}

	x := argv[n-1] // The value to be assigned.
	zv := AdaptToValue(fr, x, loc.Type())
	log.Printf(".... Reflect Set loc <%s> %s = %s", zv.Kind(), zv.Type(), Show(x))
	loc.Set(zv)
	return x
}

func cmdGoGet(fr *Frame, argv []T) T {
	varT := Arg1(argv)
	return MkT(RVars[varT.String()])
}

// TODO: test.
func cmdGoMapToHash(fr *Frame, argv []T) T {
	mapT := Arg1(argv)
	mapV := mapT.QuickReflectValue()
	keysVs := mapV.MapKeys()
	h := make(Hash, 4)
	for _, keyV := range keysVs {
		k := MkT(keyV.Interface()).String()
		v := MkT(mapV.MapIndex(keyV).Interface())
		h[k] = v
	}
	return MkHash(h)
}

func init() {
	if Unsafes == nil {
	    Unsafes = make(map[string]Command, 333)
	}

	Unsafes["go-call"] = cmdCall
	Unsafes["go-send"] = cmdSend
	Unsafes["go-getf"] = cmdGetf
	Unsafes["go-setf"] = cmdSetf
	Unsafes["go-get"] = cmdGoGet

	Unsafes["go-elem"] = cmdElem
	Unsafes["go-index"] = cmdIndex
	Unsafes["go-map-to-hash"] = cmdGoMapToHash

	RFuncs = make(map[string]interface{})
	RVars = make(map[string]interface{})
	RTypes = make(map[string]interface{})
	RConsts = make(map[string]interface{})
}
