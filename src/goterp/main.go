package main

import (
	. "fmt"
	"os"

	. "terp"
)

func main() {
	t := NewTerp()
	for _, a := range os.Args[1:] {
		Printf("<<< %q\n", a)
		z := t.Eval(a)
		Printf(">>> %q\n", z)
	}
}
