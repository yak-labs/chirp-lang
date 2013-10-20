package posix

import (
	"bufio"
	. "fmt"
	. "github.com/yak-labs/chirp-lang"
	"os"
	R "reflect"
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
	// Unsafes["gets"] = cmdGets
	// Unsafes["puts"] = cmdPuts
}
