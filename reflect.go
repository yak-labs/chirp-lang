package chirp

import (
	"errors"
	. "fmt"
	"log"
	R "reflect"
	"unsafe"
)

var Roots map[string]Applier = make(map[string]Applier)

var errorInterfaceType R.Type = R.TypeOf(errors.New).Out(0)

// TODO: This interface could replace Command, or could Command replace this interface?
type Applier interface {
	Apply(fr *Frame, args []T) T
}
type FuncRoot struct{ Func R.Value }
type VarRoot struct{ Var R.Value }
type TypeRoot struct{ Type R.Type }
type ConstRoot struct{ Const interface{} }

func (r FuncRoot) Apply(fr *Frame, args []T) T {
	if fr.G.IsSafe {
		panic("FuncRoot.Apply not allowed in Safe Interpreter.")
	}
	return commonCall(fr, args[0].String(), r.Func, args, 1)
}

func (r VarRoot) Apply(fr *Frame, args []T) T {
	if fr.G.IsSafe {
		panic("VarRoot.Apply not allowed in Safe Interpreter.")
	}
	return MkValue(r.Var.Elem())
}

func (r TypeRoot) Apply(fr *Frame, args []T) T {
	if fr.G.IsSafe {
		panic("TypeRoot.Apply not allowed in Safe Interpreter.")
	}
	if Debug['r'] {
		Say("ApplyToType <<<", r.Type)
		Sayf("............... %s", Showv(args))
	}
	z := ApplyToType(fr, r.Type, args)
	if Debug['r'] {
		Say("ApplyToType >>>", z)
	}
	return z
}

func (r ConstRoot) Apply(fr *Frame, args []T) T {
	if fr.G.IsSafe {
		panic("ConstRoot.Apply not allowed in Safe Interpreter.")
	}
	return MkT(r.Const)
}

func ApplyToType(fr *Frame, typ R.Type, args []T) T {
	if fr.G.IsSafe {
		panic("ApplyToType not allowed in Safe Interpreter.")
	}
	switch len(args) {
	case 1:
		// Backwards Compat
		return MkValue(R.ValueOf(typ))
	case 2:
		switch args[1].String() {
		case "type":
			return MkValue(R.ValueOf(typ))
		case "ptrtype":
			return MkValue(R.ValueOf(R.PtrTo(typ)))
		case "slicetype":
			return MkValue(R.ValueOf(R.SliceOf(typ)))
		case "chantype":
			return MkValue(R.ValueOf(R.ChanOf(R.BothDir, typ)))
		case "new":
			val := R.New(typ)
			if Debug['z'] {
				Say("new: CanSet:", val, val.CanSet(), val.Elem().CanSet())
			}
			z := MkValue(val)
			if Debug['z'] {
				Say("new: CanSet:", z, z.v.Elem().CanSet())
			}
			return z
		case "zero":
			return MkValue(R.Zero(typ))
		case "mkslice":
			return MkValue(R.ValueOf(R.MakeSlice(R.SliceOf(typ), 0, 4)))
		}
	case 3:
		switch args[1].String() {
		case "chantype":
			var dir R.ChanDir
			switch args[2].String()[0] {
			case 'R', 'r', '1':
				dir = R.RecvDir
			case 'S', 's', '2':
				dir = R.SendDir
			case 'B', 'b', '3':
				dir = R.BothDir
			default:
				panic(Sprintf("Bad ChanDir in chantype"))
			}
			return MkValue(R.ValueOf(R.ChanOf(dir, typ)))

		case "maptype":
			elemType, ok := args[2].QuickReflectValue().Interface().(R.Type)
			if !ok {
				panic(Sprintf("Bad element type in maptype"))
			}
			return MkValue(R.ValueOf(R.MapOf(typ, elemType)))

		case "mkmap":
			typ2, ok := args[2].Raw().(R.Type) // Element type.
			if !ok {
				panic(Sprintf("ApplyToType: Expected an element type as argument to 'mkmap' on key type %v", typ))
			}
			return MkValue(R.Value(R.MakeMap(R.MapOf(typ, typ2))))
		case "mkslice":
			k := int(args[2].Int())
			return MkValue(R.Value(R.MakeSlice(R.SliceOf(typ), k, 4)))
		}
	}
	panic(Sprintf("ApplyToType: Bad usage of Reflected Type %v", typ))
}

// ApplyToReflectedValue applies args, starting at i, to terpValue t.
func ApplyToReflectedValue(fr *Frame, v R.Value, args []T, i int) T {
	if fr.G.IsSafe {
		panic("ApplyToReflectedValue not allowed in Safe Interpreter.")
	}
	if Debug['r'] {
		Say("ApplyToReflectedValue <<<", v, i)
		Sayf("......................... %s", Showv(args))
	}
	z := applyToReflectedValueRecur(fr, v, args, i)
	if Debug['r'] {
		Say("ApplyToReflectedValue >>>", z)
	}
	return z
}
func applyToReflectedValueRecur(fr *Frame, v R.Value, args []T, i int) T {
	if Debug['z'] {
		Say("applyToReflectedValueRecur  CanSet?", v, v.CanSet())
	}
	n := len(args)
	if n-1 < i { // No more operations left to do.
		return MkT(v.Interface())
	}
	methName := args[i].String()
	typ := v.Type()
	kind := v.Kind()
	switch methName {
	case "":
		if len(methName) < 1 {
			panic(Sprintf("Empty method name called on Reflected %q", v))
		}
	case ".":
		if Debug['z'] {
			Say("Try .")
		}
		v2 := v
		switch kind {
		case R.Ptr, R.Interface:
			if v.IsNil() {
				panic(Sprintf("Cannot use field %q on NIL pointer", args[i+1].String()))
			}
			v2 = v.Elem()
			if Debug['z'] {
				Say("Deref . ", v2)
			}
		}

		switch v2.Kind() {
		case R.Struct:
			if Debug['z'] {
				Say("case . struct")
			}
			if n-1 < i+1 {
				panic(Sprintf("Missing fieldName after '.' on Reflected %q", v2))
			}
			fieldName := args[i+1].String()
			field := v2.FieldByName(fieldName)
			if !field.IsValid() {
				panic(Sprintf("No field named %q on Reflected %q", fieldName, v2))
			}

			if Debug['z'] {
				Say("n i args", n, i, args)
			}
			if n > i+3 && args[i+2].String() == "=" {
				if Debug['z'] {
					Say("case . struct =")
				}
				if n-1 != i+3 {
					panic(Sprintf("Extra args not allowed after: . field = x"))
				}
				// case Set Field.
				x := AdaptToValue(fr, args[i+3], field.Type())
				if Debug['z'] {
					Say("AdaptToValue arg field -> x", args[i+3], field, x)
				}
				if !x.IsValid() {
					panic(Sprintf("Cannot Adapt value %s to field named %q on Reflected %q", fieldName, v2))
				}
				if Debug['z'] {
					Say("Setting X, CanSet?", v, field, v.CanSet(), v2.CanSet(), field.CanSet())
				}
				field.Set(x)
				if Debug['z'] {
					Say("Did Set X")
				}
				return MkValue(x)

			} else {
				if Debug['z'] {
					Say("case . struct ! =")
				}
				// case Get Field.
				return applyToReflectedValueRecur(fr, field, args, i+2)
			}
		}
		if Debug['z'] {
			Say("drop out .")
		}

	case "@":
		if Debug['z'] {
			Say("try @")
		}
		v2 := v
		switch kind {
		case R.Ptr, R.Interface:
			if v.IsNil() {
				panic(Sprintf("Cannot use @ on NIL pointer"))
			}
			v2 = v.Elem()
			if Debug['z'] {
				Say("Deref @", v2)
			}
		}

		if Debug['z'] {
			Say("more @", v2.Kind().String())
		}
		switch v2.Kind() {
		case R.Slice:
			if Debug['z'] {
				Say("more @ case Slice")
			}
			if n-1 < i+1 {
				panic(Sprintf("Missing fieldName after '.' on Reflected %q", v2))
			}
			k := int(args[i+1].Int())
			target := v2.Index(k)
			if !target.IsValid() {
				panic(Sprintf("Failed index %d on Reflected %q", k, v2))
			}
			//////

			if n > i+3 && args[i+2].String() == "=" {
				if Debug['z'] {
					Say("case @ struct =")
				}
				if n-1 != i+3 {
					panic(Sprintf("Extra args not allowed after: . field = x"))
				}
				// case Set Field.
				x := AdaptToValue(fr, args[i+3], target.Type())
				if Debug['z'] {
					Say("AdaptToValue @ arg target -> x", args[i+3], target, x)
				}
				if !x.IsValid() {
					panic(Sprintf("Cannot Adapt value %s to target on Reflected %q", args[i+3], v2))
				}
				if Debug['z'] {
					Say("Setting @ X")
				}
				target.Set(x)
				return MkValue(x)
			}
			//////
			return ApplyToReflectedValue(fr, target, args, i+2)

		case R.Map:
			if Debug['z'] {
				Say("more @ case Map")
			}
			if n-1 < i+1 {
				panic(Sprintf("Missing fieldName after '.' on Reflected %q", v2))
			}
			k := R.ValueOf(args[i+1].Raw())
			target := v2.MapIndex(k)
			if !target.IsValid() {
				panic(Sprintf("Failed MapIndex() on Reflected %q", v2))
			}
			//////

			if n > i+3 && args[i+2].String() == "=" {
				if Debug['z'] {
					Say("case map@ struct =")
				}
				if n-1 != i+3 {
					panic(Sprintf("Extra args not allowed after: . field = x"))
				}
				// case Set Field.
				x := AdaptToValue(fr, args[i+3], target.Type())
				if Debug['z'] {
					Say("AdaptToValue map@ arg target -> x", args[i+3], target, x)
				}
				if !x.IsValid() {
					panic(Sprintf("Cannot Adapt value %s to target on Reflected %q", args[i+3], v2))
				}
				if Debug['z'] {
					Say("Setting map@ X")
				}
				target.Set(x)
				return MkValue(x)
			}
			//////
			return applyToReflectedValueRecur(fr, target, args, i+2)
		}
		if Debug['z'] {
			Say("drop out @")
		}
	case "*":
		switch kind {
		case R.Ptr, R.Interface:
			if v.IsNil() {
				panic(Sprintf("Cannot dereference NIL pointer"))
			}
			z := v.Elem()
			if !z.IsValid() {
				panic(Sprintf("Failed Elem() on Reflected %q", v))
			}
			return applyToReflectedValueRecur(fr, z, args, i+1)
		}
		if Debug['z'] {
			Say("drop out *")
		}

	case "kind":
		return MkString(kind.String())
	case "type":
		return MkValue(R.ValueOf(typ))
	}

	// default:
	fn := v.MethodByName(methName)
	if !fn.IsValid() {
		v3 := v

		if Debug['z'] {
			Say("v:", v3)
		}
		if Debug['z'] {
			Say("v3:", v3)
		}
		if Debug['z'] {
			Say("v3.Kind():", v3.Kind())
		}
		if Debug['z'] {
			Say("v3.Type():", v3.Type())
		}
		if Debug['z'] {
			Say("v3.Interface():", v3.Interface())
		}
		typeObj, ok := v3.Interface().(R.Type)
		if ok {
			if Debug['z'] {
				Say("ok:ApplyToType", typeObj, args)
			}
			return ApplyToType(fr, typeObj, args)
		}
		if Debug['z'] {
			Say("notok", args)
		}

		switch v.Kind() {
		case R.Ptr, R.Interface:
			if v.IsNil() {
				panic(Sprintf("Cannot call method %q on NIL pointer", methName))
			}
			v3 = v.Elem()
		}
		fn = v3.MethodByName(methName)
		if !fn.IsValid() {
			panic(Sprintf("No such method %q on Reflected %q", methName, v))
		}
	}
	return commonCall(fr, methName, fn, args, i+1)
}

func LookupRootAndApply(fr *Frame, rootName string, args []T) T {
	r, ok := Roots[rootName]
	if !ok {
		panic(Sprintf("Root not found: %q", rootName))
	}
	return r.Apply(fr, args)
}

func AdaptToValue(fr *Frame, a T, ty R.Type) R.Value {
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
		if ty == TypeT { // TODO: why doesn't this work?
			return R.ValueOf(a)
		}
		if ty.AssignableTo(TypeT) { // TODO: why doesn't this work?
			return R.ValueOf(a)
		}
		if ty.String() == "chirp.T" { // TODO: This is fragile & lame, but it works.
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
	if Debug['r'] {
		log.Printf("AdaptToValue: Default: for type <%s>: %s", ty, Show(a))
	}
	return R.ValueOf(a.Raw())
}

// commonCall is common to both "go-call" and "go-send".
// args already has function name or receiver and message name stripped; it's just the args.
func commonCall(fr *Frame, funcName string, fn R.Value, args []T, numFrontArgs int) T {
	if fr.G.IsSafe {
		panic("Calling reflected function or method is not allowed in Safe Interpreter.")
	}
	ty := fn.Type()

	// log.Printf("... fn <%s> type: <%s> %s", funcName, ty.Kind(), ty)

	nin := ty.NumIn()
	nout := ty.NumOut()
	vari := ty.IsVariadic() // Last arg has "..." and presents as slice type

	// Log the function's expected arguments.
	for i := 0; i < nin; i++ {
		// log.Printf("... fn expect in[%d] : <%s> %s", i, ty.In(i).Kind(), ty.In(i))
	}
	for i := 0; i < nout; i++ {
		// log.Printf("... fn returns out[%d] : <%s>  %s", i, ty.Out(i).Kind(), ty.Out(i))
	}

	// Check num args
	if vari {
		if len(args)-numFrontArgs < nin-1 {
			panic(Sprintf("command <%s> expected at least %d args but got %d.", funcName, nin-1, len(args)-numFrontArgs))
		}
	} else {
		if len(args)-numFrontArgs != nin {
			panic(Sprintf("command <%s> expected %d args but got %d.", funcName, nin, len(args)-numFrontArgs))
		}
	}

	// Convert actual args to Values, into pp.
	pp := make([]R.Value, len(args)-numFrontArgs)
	for i, p := range args[numFrontArgs:] {
		var target R.Type
		if vari && i >= nin-1 {
			target = ty.In(nin - 1).Elem() // element type of variadic args...t
		} else {
			target = ty.In(i)
		}
		pp[i] = AdaptToValue(fr, p, target)
		// log.Printf("........ passing [%d] %s as (%s) %s", i, Show(p), pp[i].Kind().String(), pp[i].Type().String())

	}

	// Call it.
	// log.Printf("...(calling)...  %v  (  %v  )", funcName, pp)
	xx := fn.Call(pp)
	// log.Printf("...(called)...")

	// Convert results from Values into zz.
	zz := make([]T, len(xx))
	for i, x := range xx {
		z := x.Interface()
		zz[i] = MkT(z)
		// log.Printf("........ result [%d] %s from (%s) %s", i, Show(zz[i]), x.Kind().String(), x.Type().String())
	}

	// If last arg is type error, remove it if nil, or panic it.
	if nout > 0 && ty.Out(nout-1).Implements(errorInterfaceType) {
		// log.Printf("Last result implements error; checking it: %#v", zz[nout-1])
		if !xx[nout-1].IsNil() {
			panic(Sprintf("call %q returns error: %q", funcName, xx[nout-1].Interface().(error).Error()))
		}
		// log.Printf("Last result implements error; was nil; dropping it.")
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

	Unsafes["go-map-to-hash"] = cmdGoMapToHash
}
