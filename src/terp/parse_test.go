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
	Must(3, len(a))
	MustT("one", a[0])
	MustT("three", a[2])
}

func TestT2(t *testing.T) {
	a := ParseList("  one  { number two } three  ")
	if len(a) != 3 {
		panic("len not 3")
	}
	println(Repr(a))
	Must(3, len(a))
	MustT("one", a[0])
	MustT(" number two ", a[1])
	MustT("three", a[2])
}

/*
func TestT3(t *testing.T) {
	fr := New()
	a := fr.TEval(MkTs("list abc[list def]ghi"))
	Must("xabcdefghi", a.String())
}
*/
