package main

import (
	. "fmt"
	"exp/types"
	"go/ast"
)

var imports = make(map[string]*ast.Object)

func main() {
	var pkg *ast.Object
	var err error

	pkg, err = types.GcImport(imports, "bufio")
	if err != nil { panic(err)}

	Printf("pkg=%#v\n", pkg)
	Printf("pkg.Data=%#v\n", pkg.Data)
}
