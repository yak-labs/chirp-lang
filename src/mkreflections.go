/*
	mkexterns.go

	Outputs generated/reflections.go
*/
package main

import (
	. "fmt"
	"exp/types"
	"go/ast"
	"strings"
)

var imports = make(map[string]*ast.Object)

var packages = []string{
	"bufio", "io", "os", "bytes", "strings", "fmt", "net/http",
}

func RemarkScope(pack string, sc *ast.Scope) {
	for k, v := range sc.Objects {
		if ast.IsExported(k) {
			Printf("// %8s %s.%s\n", v.Kind, pack, k)
		}
	}
}

func DoScope(pack string, sc *ast.Scope) {
	for k, v := range sc.Objects {
		if ast.IsExported(k) {
			short := pack
			j := strings.LastIndex(pack, "/")
			if j > 0 {
				short = pack[j+1:]
			}
			switch v.Kind.String() {
			case "type":
				Printf("Types[\"/%s/%s\"] = new(*%s.%s)\n", pack, k, short, k)
				//Printf("Converters[\"/%s/%s\"] = nil\n", pack, k)
			case "func":
				Printf("Funcs[\"/%s/%s\"] = %s.%s\n", pack, k, short, k)
			case "var":
				Printf("Vars[\"/%s/%s\"] = &%s.%s\n", pack, k, short, k)
			case "const":
				Printf("Consts[\"/%s/%s\"] = %s.%s\n", pack, k, short, k)
			}
		}
	}
}

func main() {
	var err error
	var p *ast.Object
	var pp []*ast.Object = make([]*ast.Object, 0)

	Println("package generated")
	Println()
	for _, pack := range packages {
		p, err = types.GcImport(imports, pack)
		if err != nil {
			panic(err)
		}
		pp = append(pp, p)
		Printf("import \"%s\"\n", pack)
	}

	Println()
	Println("var Types map[string]interface{} = make(map[string]interface{}, 0)")
	Println("var Funcs map[string]interface{} = make(map[string]interface{}, 0)")
	Println("var Vars map[string]interface{} = make(map[string]interface{}, 0)")
	Println("var Consts map[string]interface{} = make(map[string]interface{}, 0)")

	//Println("type TypeConverter func (interface{}) interface{}")
	//Println("var Converters map[string]TypeConverter = make(map[string]TypeConverter, 0)")

	Println()
	Println("func init() {")

	for i, pkg := range pp {
		// ast.Print(nil, pkg)
		Println()
		RemarkScope(packages[i], pkg.Data.(*ast.Scope))
		Println()
		DoScope(packages[i], pkg.Data.(*ast.Scope))
		Println()
	}

	Println("} // END")
}
