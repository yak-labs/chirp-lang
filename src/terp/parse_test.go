package terp

import (
	"testing"
)

func Test1(t *testing.T) {
	a := ParseList("  one  two three  ")
	if len(a) != 3 {
		panic("len not 3")
	}
	Must(3, len(a))
	Must("one", a[0])
	Must("three", a[2])
}

func Test2(t *testing.T) {
	a := ParseList("  one  { number two } three  ")
	//if len(a) != 3 { panic("len not 3") }
	println(Repr(a))
	Must(3, len(a), a)
	Must("one", a[0])
	Must(" number two ", a[1])
	Must("three", a[2])
}
