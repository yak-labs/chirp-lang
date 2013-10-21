package posix

import (
	"bufio"
	. "fmt"
	. "github.com/yak-labs/chirp-lang"
	"io"
	"os"
	R "reflect"
	"strings"
)

type terpFile struct {
	f *os.File
	r *bufio.Reader
	w *bufio.Writer
}

func MkFile(f *os.File) *terpFile {
	return &terpFile{f: f}
}

// *terpFile implements T

func (t *terpFile) Raw() interface{} {
	return t.f
}
func (t *terpFile) String() string {
	return Sprintf("file%d", t.f.Fd())
}
func (t *terpFile) Float() float64 {
	panic("not implemented on terpFile (Float)")
}
func (t *terpFile) Int() int64 {
	panic("not implemented on terpFile (Int)")
}
func (t *terpFile) Uint() uint64 {
	panic("not implemented on terpFile (Uint)")
}
func (t *terpFile) ListElementString() string {
	return t.String()
}
func (t *terpFile) Bool() bool {
	panic("terpFile cannot be used as Bool")
}
func (t *terpFile) IsEmpty() bool {
	return false
}
func (t *terpFile) IsPreservedByList() bool { return true }
func (t *terpFile) IsQuickNumber() bool     { return false }
func (t *terpFile) List() []T {
	return []T{t}
}
func (t *terpFile) HeadTail() (hd, tl T) {
	return MkList(t.List()).HeadTail()
}
func (t *terpFile) Hash() Hash {
	panic("a terpFile is not a Hash")
}
func (t *terpFile) GetAt(key T) T {
	panic("a terpFile cannot GetAt")
}
func (t *terpFile) PutAt(value T, key T) {
	panic("a terpFile cannot PutAt")
}
func (t *terpFile) QuickReflectValue() R.Value {
	panic("a terpFile cannot QuickReflectValue")
}

func cmdOpen(fr *Frame, argv []T) T {
	nameT, args := Arg1v(argv)
	name := nameT.String()

	// access defaults to "r" if no extra arg.
	access := "r"
	if len(args) > 0 {
		access = args[0].String()
	}

	return Open(name, access)
}

func Open(name string, access string) T {
	var f *os.File
	var err error
	switch access {
	case "r":
		f, err = os.OpenFile(name, os.O_RDONLY, 0666)
	case "r+":
		f, err = os.OpenFile(name, os.O_RDWR, 0666)
	case "w":
		f, err = os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	case "w+":
		f, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0666)
	case "a":
		f, err = os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	case "a+":
		f, err = os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	default:
		panic(Sprintf(`Unknown access mode in "open" command: %q`, access))
	}

	if err != nil {
		panic(Sprintf(`Cannot "open" file %q because %q`, name, err.Error()))
	}

	return &terpFile{f: f}
}

func cmdFlush(fr *Frame, argv []T) T {
	fileT := Arg1(argv)
	tf := fileT.(*terpFile)
	Flush(tf)
	return Empty
}

func Flush(tf *terpFile) {
	if tf.w != nil {
		tf.w.Flush()
	}
}

func cmdClose(fr *Frame, argv []T) T {
	fileT := Arg1(argv)
	tf := fileT.(*terpFile)
	Close(tf)
	return Empty
}

func Close(tf *terpFile) {
	if tf.w != nil {
		tf.w.Flush()
	}
	if tf.f != nil {
		tf.f.Close()
	}

	tf.f = nil
	tf.r = nil
	tf.w = nil
}

func cmdGets(fr *Frame, argv []T) T {
	fileT, args := Arg1v(argv)
	var varName string
	if len(args) > 1 {
		panic(`Too many args to "gets"`)
	}
	if len(args) > 0 {
		varName = args[0].String()
	}
	f := fileT.(*terpFile)

	if f.r == nil {
		f.r = bufio.NewReader(f.f)
	}

	data, err := f.r.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(Sprintf(`Error duing "gets": %s`, err.Error()))
	}
	if len(data) > 0 {
		if data[len(data)-1] == '\n' {
			data = data[:len(data)-1]
		}
	}
	dataT := MkString(data)

	if len(varName) > 0 {
		fr.SetVar(varName, dataT)
		return MkInt(int64(len(data)))
	}
	// else:
	return dataT
}

func cmdPuts(fr *Frame, argv []T) T {
	i := 1
	noNewLine := false
	if len(argv) > i {
		if strings.HasPrefix(argv[i].String(), "-n") && strings.HasPrefix("-nonewline", argv[i].String()) {
			noNewLine = true
			i++
		}
	}

	var t *terpFile
	var data string
	switch len(argv) {
	case i + 1:
		data = argv[i].String()
	case i + 2:
		var ok bool
		t, ok = argv[i].(*terpFile)
		if !ok {
			panic(Sprintf(`Bad args to "puts". Expected file as arg %d.`, i))
		}
		data = argv[i+1].String()
	default:
		panic(`Bad args to "puts"`)
	}

	Puts(noNewLine, t, data)
	return Empty
}

func Puts(noNewLine bool, t *terpFile, data string) {
	var err error
	if t == nil {
		if noNewLine {
			_, err = Print(data)
		} else {
			_, err = Println(data)
		}
	} else {
		if t.w == nil {
			t.w = bufio.NewWriter(t.f)
		}
		if noNewLine {
			_, err = Fprint(t.w, data)
		} else {
			_, err = Fprintln(t.w, data)
		}
	}
	if err != nil {
		panic(Sprintf(`Error during "puts": %s`, err.Error()))
	}
}

var fileEnsemble = []EnsembleItem{
	EnsembleItem{Name: "separator", Cmd: cmdFileSeparator},
	EnsembleItem{Name: "tempdir", Cmd: cmdFileTempdir},
	EnsembleItem{Name: "join", Cmd: cmdFileJoin},
}

func cmdFileSeparator(fr *Frame, argv []T) T {
	Arg0(argv)
	return MkString(string(os.PathSeparator))
}

func cmdFileTempdir(fr *Frame, argv []T) T {
	Arg0(argv)
	return MkString(os.TempDir())
}

func cmdFileJoin(fr *Frame, argv []T) T {
	panic("TODO")
}

func init() {
	if Unsafes == nil {
		Unsafes = make(map[string]Command, 333)
	}

	Unsafes["open"] = cmdOpen
	Unsafes["close"] = cmdClose
	Unsafes["file"] = MkEnsemble(fileEnsemble)
	Unsafes["gets"] = cmdGets
	Unsafes["puts"] = cmdPuts
	Unsafes["flush"] = cmdFlush
}
