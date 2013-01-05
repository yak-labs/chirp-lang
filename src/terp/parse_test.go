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

func TestT3(t *testing.T) {
	fr := New()
	a := fr.Eval(MkString("list xabc[list def]ghi"))
	MustA("xabcdefghi", a.String())
}

func TestParseEscaping(t *testing.T) {
	fr := New()
	s := `proc p {} {
		return "0\132\tNumber\n"
	} ; p`
	x := fr.Eval(MkString(s))
	if x.String() != "0\132\tNumber\n" {
		t.Errorf("Broken x was %q.", x.String())
	}
}

func TestSlashEscaping(t *testing.T) {
	s := "0\132\tNumber\n"
	x := MkString(s).ListElementString()
	if x != `{0Z\011Number\012}` {
		t.Errorf("Broken x was %q.", x)
	}
}

