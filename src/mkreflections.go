/*
	mkexterns.go

	Outputs generated/reflections.go
*/
package main

import (
	"exp/types"
	. "fmt"
	"go/ast"
	"os"
	"runtime"
	"path/filepath"
	"strings"
)

var imports = make(map[string]*ast.Object)

var packages = make([]string, 0)

var pkgbasepath, _ = filepath.Abs(Sprintf("%s/pkg/%s_%s",
	os.Getenv("GOROOT"),
	runtime.GOOS,
	runtime.GOARCH))

func RemarkScope(pack string, sc *ast.Scope) {
	for k, v := range sc.Objects {
		if ast.IsExported(k) {
			Printf("// %8s %s.%s\n", v.Kind, pack, k)
		}
	}
}

var Uniq map[string]int = make(map[string]int, 5000)

func DoScope(pack string, sc *ast.Scope) {
	for k, v := range sc.Objects {
		if ast.IsExported(k) {
			snake := strings.Replace(pack, "/", "_", -1)
			snake = strings.Replace(snake, ".", "_", -1)
			snake = strings.Replace(snake, "-", "_", -1)
			short := pack
			j := strings.LastIndex(pack, "/")
			if j > 0 {
				short = pack[j+1:]
			}
			switch v.Kind.String() {
			case "type":
				Printf("Types[\"/%s/%s\"] = new(*%s.%s)\n", pack, k, snake, k)
				//Printf("Converters[\"/%s/%s\"] = nil\n", pack, k)
			case "func":
				Printf("Funcs[\"/%s/%s\"] = %s.%s\n", pack, k, snake, k)
			case "var":
				Printf("Vars[\"/%s/%s\"] = &%s.%s\n", pack, k, snake, k)
			case "const":
				Printf("//? Consts[\"/%s/%s\"] = %s.%s\n", pack, k, snake, k)
			}

			pre := short[:1]
			nom := Sprintf("{%s%s}", pre, k)
			Printf("//# OLD <%s> %d\n", nom, Uniq[nom])
			Uniq[nom]++
			Printf("//# NEW <%s> %d\n", nom, Uniq[nom])
		}
	}
}

// Finds the packages to use specific to this
func findPackages(path string, info os.FileInfo, err error) error {
	if (strings.HasSuffix(path, ".a") && !strings.HasSuffix(path, "cgo.a")) {
		var pkg, _ = filepath.Rel(pkgbasepath, path)
		pkg = filepath.ToSlash(pkg) // Convert all dir seperators to correct slash
		pkg = pkg[0 : len(pkg) - 2] // Take out the extension
		packages = append(packages, pkg)
	}

	return nil
}

func main() {
	var err error
	var p *ast.Object
	var pp []*ast.Object = make([]*ast.Object, 0)

	Println("package generated")
	Println()

	if len(os.Args) > 1 {
		packages = os.Args[1:]
	} else {
		// Walk the packages directory to find our packages
		filepath.Walk(pkgbasepath, findPackages)
	}
	for _, pack := range packages {
		p, err = types.GcImport(imports, pack)
		if err != nil {
			panic(err)
		}
		pp = append(pp, p)
		snake := strings.Replace(pack, "/", "_", -1)
		snake = strings.Replace(snake, ".", "_", -1)
		snake = strings.Replace(snake, "-", "_", -1)
		Printf("import %s \"%s\"\n", snake, pack)
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
		if packages[i] == "sort" {
			// ast.Print(nil, pkg)
		}
		Println()
		RemarkScope(packages[i], pkg.Data.(*ast.Scope))
		Println()
		DoScope(packages[i], pkg.Data.(*ast.Scope))
		Println()
	}

	Println("} // END")

	for uk, uv := range Uniq {
		if uv > 1 {
			Printf("//# ! Uniq <%s> %d\n", uk, uv)
		}
	}
}
