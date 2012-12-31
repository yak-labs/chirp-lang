package terp_test

import (
	"testing"

	. "terp"
)

func TestSimpleList3(t *testing.T) {
	a := ParseList("  one  two three  ")
	if len(a) != 3 {
		panic("len not 3")
	}
	MustA(3, len(a))
	MustST("one", a[0])
	MustST("three", a[2])
}

func TestT2(t *testing.T) {
	a := ParseList("  one  { number two } three  ")
	if len(a) != 3 {
		panic("len not 3")
	}
	println(Repr(a))
	MustA(3, len(a))
	MustST("one", a[0])
	MustST(" number two ", a[1])
	MustST("three", a[2])
}

/*
func TestT3(t *testing.T) {
	fr := New()
	a := fr.Eval(MkString("list abc[list def]ghi"))
	MustA("xabcdefghi", a.String())
}
*/
