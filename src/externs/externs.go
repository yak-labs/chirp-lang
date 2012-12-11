package main

import . "fmt"
import "go/ast"
import "go/parser"
import "go/token"
import "os"
import "strings"

var fset = token.NewFileSet()

// Format any object as string.
func Str(x interface{}) string {
  return Sprintf("%v", x)
}

// Executable-like Representation.
func Repr(x interface{}) string {
  return Sprintf("%#v", x)
}

// Like a ? : operator for strings (but both eval).
func Cond(b bool, t string, f string) string {
  if b {
    return t
  }
  return f
}

func AfterFirstDelim(s string, delim string, otherwise string) string {
  j := strings.Index(s, delim)
  if j < 0 {
    return otherwise
  }
  return s[ j + len(delim) : ]
}

func BeforeFinalDelim(s string, delim string, otherwise string) string {
  j := strings.LastIndex(s, delim)
  if j < 0 {
    return otherwise
  }
  return s[ : j]
}

func EndsWith(s string, tail string) bool {
  n := len(s)
  t := len(tail)
  if n < t {
    return false
  }
  return s[n-t: ] == tail
}

func FilterDotGo(info os.FileInfo) bool {
  s := info.Name()
  n := len(s)
  // return strings.Contains(s, ".go")
  return n>3 && s[n-3]=='.' && s[n-2]=='g' && s[n-1]=='o'
}

func GrokDir(dir string) {
  pkgs, err := parser.ParseDir(fset, dir, FilterDotGo, parser.Mode(0))
  if err != nil {
    panic(Sprintf("ERROR <%q> IN DIR <%s>", err, dir))
  }

  // Find the subdir that follows /pkg/
  subdir := AfterFirstDelim(dir, "/pkg/", "")
  if len(subdir) == 0 {
    panic("dir does not contain /pkg/ plus a tail: " + dir)
  }

  // Trim trailing /, if any.
  if subdir[ len(subdir) - 1 ] == '/' {
    subdir = subdir[ : len(subdir) - 1 ]
  }

  for pk, pv := range pkgs {
    // For dirpkg, remove final piece of subdir
    // and replace with stated package name.
    // Often these are the same, but not always.
    updir := BeforeFinalDelim(subdir, "/", "")
    dirpkg := Cond(updir == "", pk, updir + "/" + pk)

    Printf("#dirpkg %#v\n", dirpkg);
    Printf("#pv %#v\n", pv);

    for _, fv := range pv.Files {
      for i, dcl := range fv.Decls {
        doDecl(dcl, i, dirpkg)
      }
    }
  }
}

func doDecl(dcl interface{}, i int, dirpkg string) {
	switch x := dcl.(type) {
	  case (*ast.FuncDecl):
            if !ast.IsExported(x.Name.Name) {
	      return
	    }
            fstr := funcDeclStr(x)
            Printf("@@ %s %s\n", dirpkg, fstr)

            if strings.Contains(fstr, "?") {
	      panic(fstr)
            }
	  case (*ast.GenDecl):
	    for sk, sv := range x.Specs {
	      doSpec(sk, sv, dirpkg)
	    }
	  default:
	    panic(Repr(dcl))
	}
}

func doSpec(sk int, sv interface{}, dirpkg string) {
  //Printf("doSpec: %#v\n", sv)
  switch x := sv.(type) {
  case (*ast.ImportSpec):
    path := x.Path.Value
    path = path[1 : len(path)-1]  // Omit initial & final <">
    name := path
    j := strings.LastIndex(name, "/")
    if j > 0 {
      name = name[j+1 : ]
    }
    if x.Name != nil {
      name = x.Name.Name
    }
    //Printf("ImportSpec: NAME %s PATH %s\n", name, path)
    Printf("@@ %s { IMPORT %s path: %s }\n", dirpkg, name, path)
  case (*ast.TypeSpec):
    name := x.Name.Name
    if !ast.IsExported(x.Name.Name) {
      return
    }
    Printf("@@ %s { TYPE %s type: %s }\n", dirpkg, name, typeStr(x.Type))
  case (*ast.ValueSpec):
    for _, nv := range x.Names {
      //Printf("    ValueSPEC [%d] %s : %s\n", nk, nv.Name, Cond(x.Type == nil, "nil", typeStr(x.Type)))
      if !ast.IsExported(nv.Name) {
        return
      }
      Printf("@@ %s { VALUE %s : %s }\n", dirpkg, nv.Name, Cond(x.Type == nil, "!", typeStr(x.Type)))
    }
  default:
    panic(Sprintf("doSpec: DEFAULT: %#v\n", x))
  }
}

func typeStr(a interface{}) string {
  if a == nil {
    return "nil"
  }
  switch t := a.(type) {
  case (*ast.Ident):
    return "'" + t.Name + " "
  case (*ast.ArrayType):
    return "{ ARRAY " + typeStr(t.Elt) + " } "
  case (*ast.StarExpr):
    return "{ STAR " + typeStr(t.X) + " } "
  case (*ast.Ellipsis):
    return "{ ELLIPSIS " + typeStr(t.Elt) + " } "
  case (*ast.SelectorExpr):
    return "{ SEL " + typeStr(t.X) + " dot: " + typeStr(t.Sel) + " } "
  case (*ast.InterfaceType):
    return "{ INTERFACE todo: ...  } "
  case (*ast.FuncType):
    return "{ FN " + funcParamsResults(t) +  " } "
  case (*ast.MapType):
    return "{ MAP " + typeStr(t.Key) + " to: " + typeStr(t.Value) + " } "
  case (*ast.ChanType):
    return "{ CHAN " + Str(t.Dir) + " " + typeStr(t.Value) + " } "

  case (*ast.StructType):
    z := Sprintf("{ STRUCT complete: %s { ", Cond(t.Incomplete, "F", "T"))
    Printf("#    %#v\n", t)
    Printf("#    %#v\n", t.Fields)
    for fk, fv := range t.Fields.List {
      Printf("#    [%v]  %#v\n", fk, fv)
      if fv.Names == nil {
        z = Sprintf("%s , ! : %s ", z, typeStr(fv.Type))
      }
      for _, nv := range fv.Names {
        Printf("#      Name: %s  Type: %s\n", nv.Name, typeStr(fv.Type))
        if !ast.IsExported(nv.Name) {
	  continue
	}
        z = Sprintf("%s , %s : %s ", z, nv.Name, typeStr(fv.Type))
      }
    }
    return z + " } } "
  }

  panic(Sprintf("{ ?WHAT? %#v } ", a))
}


func funcDeclStr(f *ast.FuncDecl) string {
  fstr := "{ "
  if f.Recv != nil {
    if len(f.Recv.List) != 1 { panic("f.Recv.List") }
    fstr += "METH " + f.Name.Name + " recv: " + typeStr(f.Recv.List[0].Type)
  } else {
    fstr += "FUNC " + f.Name.Name
  }
  fstr += " " + funcParamsResults(f.Type) + " } "
  return fstr
}

func funcParamsResults(f *ast.FuncType) string {
  z := "args: { "

  // list of parameters
  params := f.Params
  if params == nil { panic("params is never nil but it is") }
  for _, lv := range params.List {
    pname := "_"
    if len(lv.Names) > 0 {
      pname = lv.Names[0].Name
    }
    tname := typeStr(lv.Type)
    z +=  ", " + pname + " : " + tname
  }

  z += "} "
  z += "results: { "

  // list of parameters
  rr := f.Results
  if rr != nil {
    for _, lv := range rr.List {
      pname := "_"
      if len(lv.Names) > 0 {
        pname = lv.Names[0].Name
      }
      tname := typeStr(lv.Type)
      z +=  ", " + pname + " : " + tname
    }
  }

  z += "} "
  return z
}

func main() {
  for _, dir := range os.Args[1:] {
    GrokDir(dir)
  }
}
