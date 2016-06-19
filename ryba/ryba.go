// Package ryba is a RYe Bridge Adapater for chirp (but it sounds fishy).

package ryba

import (
	"github.com/strickyak/rye"
	"github.com/yak-labs/chirp-lang"

	. "fmt"
	R "reflect"
)

func T2B(t chirp.T) rye.M {
	// Rye C_Objects and Go reflect.Values.
	if val := t.QuickReflectValue(); val.IsValid() {
		guts := val.Interface()
		if gs, ok := guts.(rye.GetSelfer); ok {
			return gs.GetSelf().M()
		}
		return rye.MkValue(val)
	}

	raw := t.Raw()
	val := R.ValueOf(raw)
	switch val.Kind() {
	case R.Invalid:
		return rye.None
	case R.Bool:
		return rye.MkBool(val.Bool())
	case R.Int:
		return rye.MkInt(int64(val.Int()))
	case R.Int8:
		return rye.MkInt(int64(val.Int()))
	case R.Int16:
		return rye.MkInt(int64(val.Int()))
	case R.Int32:
		return rye.MkInt(int64(val.Int()))
	case R.Int64:
		return rye.MkInt(int64(val.Int()))
	case R.Uint:
		return rye.MkInt(int64(val.Uint()))
	case R.Uint8:
		return rye.MkInt(int64(val.Uint()))
	case R.Uint16:
		return rye.MkInt(int64(val.Uint()))
	case R.Uint32:
		return rye.MkInt(int64(val.Uint()))
	case R.Uint64:
		return rye.MkInt(int64(val.Uint()))
	case R.Uintptr:
		return rye.MkInt(int64(val.Uint()))
	case R.Float32:
		return rye.MkFloat(float64(val.Float()))
	case R.Float64:
		return rye.MkFloat(float64(val.Float()))
	case R.Complex64:
		panic("cannot complex")
	case R.Complex128:
		panic("cannot complex")
	case R.Array:
	case R.Chan:
	case R.Func:
	case R.Interface:
	case R.Map:
	case R.Ptr:
	case R.Slice:
	case R.String:
		return rye.MkStr(val.String())
	case R.Struct:
	case R.UnsafePointer:
	}

	panic(Sprintf("T2B: Bad Kind: %v", val.Kind()))
}

func cmdRyCall(fr *chirp.Frame, argv []chirp.T) chirp.T {
	fnT, argsT := chirp.Arg1v(argv)
	n := len(argsT)

	argsVV := make([]R.Value, n)
	for i, t := range argsT {
		argsVV[i] = R.ValueOf(rye.Mk(t.Raw()))
	}

	//println(Sprintf("qqq<<< cmdRyCall: fnT %T %T %#v", fnT, fnT.Raw(), fnT.Raw()))

	var fn interface{} = fnT.Raw()
	fnV := R.ValueOf(fn)
	resultsVV := fnV.Call(argsVV)
	if len(resultsVV) != 1 {
		panic("bad len resultsVV")
	}

	//println(Sprintf("qqq>>> cmdRyCall: fnT %T %T %#v", resultsVV[0], resultsVV[0].Interface(), resultsVV[0].Interface()))

	// Note:  This may not be quite right; we used to have some extra code for slices.
	return chirp.MkT(resultsVV[0].Interface().(rye.M).Contents())
}

func init() {
	if chirp.Unsafes == nil {
		chirp.Unsafes = make(map[string]chirp.Command, 333)
	}

	chirp.Unsafes["rycall"] = cmdRyCall
	// chirp.Unsafes["rysend"] = cmdRySend
}
