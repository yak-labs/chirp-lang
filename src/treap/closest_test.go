package treap

import (
	. "fmt"
	"math/rand"
	"testing"
)

func init() {
	// so that every run is the same seq of rand numbers
	rand.Seed(0)
}

func TestClosest(t *testing.T) {
	tree := NewTree(StringLess8)

	for i := 0; i < 10000; i++ {
		tree.Insert(Sprintf("%d", i), Sprintf("%d", i*i))
	}

	s := tree.ClosestKey("!").(string)
	if s != "0" {
		t.Errorf("expected 0")
	}
	
	s = tree.ClosestKey("0").(string)
	if s != "0" {
		t.Errorf("expected 0 again")
	}
	
	s = tree.ClosestKey("00").(string)
	if s != "1" {
		t.Errorf("expected 1")
	}

	for i := 0; i < 10000; i++ {
		k := Sprintf("%d", i)

		s = tree.ClosestKey(k).(string)
		if s != k {
			t.Errorf("expected " + k)
		}
	}

	p := tree.ClosestKey("z")
	if p != nil {
		t.Errorf("expected nil")
	}

	reverse := make(map[string]string)
	for i := 0; i < 10000; i++ {
		k := Sprintf("%d", i)
		v := Sprintf("%d", i*i)
		if k >= "5" {
			reverse[v] = k
		} else {
			reverse[v] = "bad"
		}
	}
	c := tree.IterAscendStartingAt("49999999999")
	prev := ""
	for {
		x := <-c
		if x == nil { break }
		xx := x.(string)
		rev := reverse[xx]
		println(rev)

		if rev == "bad" {
			t.Errorf("bad " + xx)
		}
		if rev == "" {
			t.Errorf("emtpy " + xx)
		}
		if rev <= prev {
			t.Errorf("backwards " + xx)
		}
		prev = rev
	}
}	

func StringLess8(p, q interface{}) bool {
	return p.(string) < q.(string)
}

func IntLess8(p, q interface{}) bool {
	return p.(int) < q.(int)
}

