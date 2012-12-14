package main

import (
	. "fmt"
	"exp/types"
	"go/ast"
	"go/token"
)

var imports = make(map[string]*ast.Object)

var packages = []string{
	"bufio", "io", "os", "bytes", "strings", "fmt",
}

func main() {
	var err error
	var p *ast.Object
	var pp []*ast.Object = make([]*ast.Object, 0)

	for _, s := range packages {
		p, err = types.GcImport(imports, s)
		if err != nil {
			panic(err)
		}
		pp = append(pp, p)
	}

	fset := token.NewFileSet()
	for i, pkg := range pp {
		Printf("\n")
		Printf("%s: \n", packages[i])
		Printf("\n")
		Printf("pkg=%#v\n", pkg)
		Printf("\n")
		Printf("pkg.Data=%#v\n", pkg.Data)
		Printf("\n")
		ast.Print(fset, pkg)
		Printf("\n")
		Printf("\n")
	}

}
