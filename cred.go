package chirp

import (
	. "fmt"
)

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
	if Unsafes == nil {
		Unsafes = make(map[string]Command, 333)
	}

	Unsafes["cred-put"] = cmdCredPut
}
