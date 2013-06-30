package chirp_test

import (
	"testing"
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

func TestDollarKey(t *testing.T) {
	fr := New()
	a := fr.Eval(MkString("set h [hash]; hset $h foo bar;  list aaa$h(foo)zzz"))
	MustA("aaabarzzz", a.String())
}

func TestComment(t *testing.T) {
	fr := New()
	a := fr.Eval(MkString("#list 1; list 2 \n list 3 \n list 4"))
	MustA("4", a.String())

	a = fr.Eval(MkString("#list 1; list 2 \n list 3 \n #list 4"))
	MustA("3", a.String())

	a = fr.Eval(MkString("list #1; list #2 \n list #3 \n #list 4"))
	MustA("#3", a.String())

	a = fr.Eval(MkString("# xyzzy \n # plough"))
	MustA("", a.String())

	a = fr.Eval(MkString("list 8 # xyzzy ; ; ; \n ; \n # plough"))
	MustA("8 # xyzzy", a.String())

	a = fr.Eval(MkString("proc #foo {} {return 777} ; \"#foo\" "))
	MustA("777", a.String())

	a = fr.Eval(MkString("proc #bar {} {return 888} ; {#bar} "))
	MustA("888", a.String())
}
