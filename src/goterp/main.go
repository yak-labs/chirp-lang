package main

import (
	. "fmt"
	"os"

	"terp"
)

func main() {
	fr := terp.New()
	for _, a := range os.Args[1:] {
		Printf("<<< %q\n", a)
		z := fr.Eval(a)
		Printf(">>> %q\n", z)
	}
}
