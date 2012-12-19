package main

import (
	. "fmt"
	"os"

	"terp"
)

func main() {
	fr := terp.New()
	for _, a := range os.Args[1:] {
		Printf("<<< %#v\n", a)
		var z terp.Any
		if a[0] == ',' {
			z = fr.TEval(terp.MkTs(a[1:]))
		} else {
			z = fr.Eval(a)
		}
		Printf(">>> %#v\n", z)
	}
}
