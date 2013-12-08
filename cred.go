package chirp

import (
	. "fmt"
	R "reflect"
)

// Opaque is a new idea for things like credentials.  Fields must be private.
type Opaque struct {
	cover    string // Enforce only a string may be seen as the cover.
	contents T
	owner    *Global
}

func cmdOpaqueWrap(fr *Frame, argv []T) T {
	cover, contents := Arg2(argv)
	obj := &Opaque{
		cover:    cover.String(),
		contents: contents,
		owner:    fr.G,
	}
	return MkValue(R.ValueOf(obj))
}

func cmdOpaqueUnwrap(fr *Frame, argv []T) T {
	obj := Arg1(argv)
	switch p := obj.Raw().(type) {
	case *Opaque:
		if p.owner != fr.G {
			panic(Sprintf("Cannot unwrap opaque object in wrong interpreter: %v", p))
		}
		return p.contents
	}
	panic(Sprintf("Not an opaque object: %v", obj.Raw()))
}

func cmdOpaqueShow(fr *Frame, argv []T) T {
	obj := Arg1(argv)
	switch p := obj.Raw().(type) {
	case *Opaque:
		return MkString(p.cover)
	}
	panic(Sprintf("Not an opaque object: %v", obj.Raw()))
}

var opaqueEnsemble = []EnsembleItem{
	EnsembleItem{Name: "wrap", Cmd: cmdOpaqueWrap},
	EnsembleItem{Name: "unwrap", Cmd: cmdOpaqueUnwrap},
	EnsembleItem{Name: "show", Cmd: cmdOpaqueShow},
}

// Cred

// Getting cred is safe.  Setting it is unsafe.  This is the setter.
func cmdCredPut(fr *Frame, argv []T) T {
	name, value := Arg2(argv)
	if fr.Cred == nil {
		fr.Cred = make(Hash, 4)
	}
	key := name.String()
	if _, ok := fr.Cred[key]; ok {
		panic(Sprintf("cred-set refuses to redefine key %q", key))
	}
	fr.Cred[key] = value
	return value
}

func init() {
	if Safes == nil {
		Safes = make(map[string]Command, 333)
	}

	Safes["opaque"] = MkEnsemble(opaqueEnsemble)

	if Unsafes == nil {
		Unsafes = make(map[string]Command, 333)
	}

	Unsafes["cred-put"] = cmdCredPut
}
