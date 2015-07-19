// Package ryba is a RYe Bridge Adapater for chirp (but it sounds fishy).

package ryba

import (
	"github.com/strickyak/rye"
	"github.com/yak-labs/chirp-lang"

	. "fmt"
	R "reflect"
)

func P2T(p rye.P) chirp.T {
	var z chirp.T
	c := p.Contents()
	println(Sprintf("P2T <<< %v <<< %v", c, p))

	switch t := c.(type) {
	case []rye.P:
		n := len(t)
		tt := make([]chirp.T, n, n)
		for i, e := range t {
			tt[i] = P2T(e) // Recurse.
		}
		z = chirp.MkList(tt)

	default:
		z = chirp.MkT(c)
	}

	println(Sprintf("P2T >>> %v", z))
	return z
}

func T2P(t chirp.T) rye.P {
	// Rye C_Objects and Go reflect.Values.
	if val := t.QuickReflectValue(); val.IsValid() {
		guts := val.Interface()
		if gs, ok := guts.(rye.GetSelfer); ok {
			return gs.GetSelf()
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

	panic(Sprintf("T2P: Bad Kind: %v", val.Kind()))
}

func cmdRyCall(fr *chirp.Frame, argv []chirp.T) chirp.T {
	fnT, argsT := chirp.Arg1v(argv)
	//fnP := T2P(fnT)
	n := len(argsT)
	argsP := make([]rye.P, n, n)
	for i, t := range argsT {
		argsP[i] = T2P(t)
	}

	if p, ok := fnT.Raw().(rye.ICallV); ok {
		z := p.CallV(argsP, nil, nil, nil)
		return P2T(z)
	}
	panic("rycall: fn not an ICallV")
}

func init() {
	if chirp.Unsafes == nil {
		chirp.Unsafes = make(map[string]chirp.Command, 333)
	}

	chirp.Unsafes["rycall"] = cmdRyCall
	// chirp.Unsafes["rysend"] = cmdRySend
}
