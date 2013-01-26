package terp

import (
	. "fmt"
	R "reflect"
)

type Obj struct {
	Class	*Obj
	Super	*Obj
	Slots	Hash
}

func cmdSubclass(fr *Frame, argv []T) T {
	super := Arg1(argv)
	var sup *Obj
	if ! super.IsEmpty() {
		sup = super.Raw().(*Obj)
	}
	z := &Obj{
		Super: sup,
		Slots: make(Hash),
	}
	return MkValue(R.ValueOf(z))
}

var methSerial = 100
func cmdMeth(fr *Frame, argv []T) T {
	cls, name, arglist, body := Arg4(argv)
	super := cls.Raw().(*Obj).Super
	slots := cls.Raw().(*Obj).Slots

	methSerial++
	tmpName := Sprintf("%s_%d", name, methSerial)
	procArgs := []T{MkString("proc"), MkString(tmpName), arglist, body}

	procOrYProc(fr, procArgs, false, super)
	node := fr.G.Cmds[tmpName]
	slots[name.String()] = MkValue(R.ValueOf(node.Fn))
	return Empty
}

func cmdNew(fr *Frame, argv []T) T {
	class := Arg1(argv)
	cls := class.Raw().(*Obj)
	z := &Obj{
		Class: cls,
		Slots: make(Hash),
	}
	return MkValue(R.ValueOf(z))
}

func cmdOn(fr *Frame, argv []T) T {
	rcvr, msg, _ := Arg2v(argv)
	cls := rcvr.Raw().(*Obj).Class
	cmdT, ok := cls.Slots[msg.String()]
	if !ok {
		panic(Sprintf("Cannot find message: %q", msg.String()))
	}
	cmd := cmdT.Raw().(Command)
	return cmd(fr, argv[1:])
}

func (fr *Frame) initClass() {
	Builtins["subclass"] = cmdSubclass
	Builtins["meth"] = cmdMeth
	Builtins["new"] = cmdNew
	Builtins["on"] = cmdOn
}
