package chirp

import (
	"bytes"
	. "fmt"
	R "reflect"
	"sort"
	"strconv"
	"strings"
)

// T is an interface to any Tcl value.
// Use them only through these methods, or fix these methods.
type T interface {
	Raw() interface{}
	String() string
	Float() float64
	Int() int64
	Uint() uint64
	ListElementString() string
	Bool() bool    // Like Python, empty values and 0 values are false.
	IsEmpty() bool // Would String() return ""?
	List() []T
	IsPreservedByList() bool
	IsQuickNumber() bool
	HeadTail() (hd, tl T)
	Hash() Hash
	GetAt(key T) T
	PutAt(value T, key T)
	QuickReflectValue() R.Value
}

// terpFloat is a Tcl value holding a float64.
type terpFloat struct { // Implements T.
	f float64
}

// terpString is a Tcl value holding a string.
type terpString struct { // Implements T.
	s string
}

// terpList is a Tcl value holding a List.
type terpList struct { // Implements T.
	l []T
}

// terpValue is a Tcl value holding a Go reflect.Value.
// It is a handle to non-Tcl Go objets.
type terpValue struct { // Implements T.
	v R.Value
}

// terpGenerator holds a channel for reading from a generator (yproc command).
type terpGeneratorGuts struct { // Mutable.
	readerChan <-chan Either
	hd         T
	tl         T
}
type terpGenerator struct { // Implements T.
	guts *terpGeneratorGuts // Pointer to mutable.
}

// terpHash holds a Hash.
type terpHash struct { // Imlements T.
	H Hash
}

func MkHash(h Hash) terpHash {
	if h == nil {
		return terpHash{H: make(Hash, 4)}
	}
	return terpHash{H: h}
}
func MkGenerator(readerChan <-chan Either) terpGenerator {
	return terpGenerator{guts: &terpGeneratorGuts{readerChan: readerChan}}
}
func MkBool(a bool) terpFloat {
	if a {
		return True
	}
	return False
}
func MkFloat(a float64) terpFloat {
	return terpFloat{f: a}
}
func MkInt(a int64) terpFloat {
	return terpFloat{f: float64(a)}
}
func MkUint(a uint64) terpFloat {
	return terpFloat{f: float64(a)}
}
func MkString(a string) terpString {
	return terpString{s: a}
}
func MkList(a []T) terpList {
	return terpList{l: a}
}
func MkStringList(a []string) terpList {
	z := make([]T, len(a))
	for i, e := range a {
		z[i] = MkString(e)
	}
	return terpList{l: z}
}
func MkValue(a R.Value) terpValue {
	return terpValue{v: a}
}
func MkT(a interface{}) T {
	// Very specific type cases.
	switch x := a.(type) {
	case T:
		// panic(Sprintf("Calling MkT() on a T: <%T> <%#v> %s", x, x, x.String()))
		return x

	case R.Value:
		// Some day we'll allow this, but for now, flag an error.
		panic(Sprintf("Calling MkT() on a R.Value: <%T> <%#v> %s", x, x, x.String()))

	case nil:
		return Empty

	case string:
		if len(x) > 6 && x[:6] == "Value:" {
			panic(666) // TODO: TEMPORARY
		}
	}

	// Use reflection to figure it out.
	v := R.ValueOf(a)
	switch v.Kind() {

	case R.Bool:
		return MkBool(v.Bool())

	case R.Int:
		return MkInt(v.Int())
	case R.Int8:
		return MkInt(v.Int())
	case R.Int16:
		return MkInt(v.Int())
	case R.Int32:
		return MkInt(v.Int())
	case R.Int64:
		return MkInt(v.Int())

	case R.Uint:
		return MkUint(v.Uint())
	case R.Uint8:
		return MkUint(v.Uint())
	case R.Uint16:
		return MkUint(v.Uint())
	case R.Uint32:
		return MkUint(v.Uint())
	case R.Uint64:
		return MkUint(v.Uint())
	case R.Uintptr:
		return MkUint(v.Uint())

	case R.Float32:
		return MkFloat(v.Float())
	case R.Float64:
		return MkFloat(v.Float())

	case R.Complex64:
	case R.Complex128:

	case R.Array:
	case R.Chan:
		if v.IsNil() {
			return Empty
		}
	case R.Func:
		if v.IsNil() {
			return Empty
		}
	case R.Interface:
		if v.IsNil() {
			return Empty
		}
	case R.Map:
		if v.IsNil() {
			return Empty
		}
	case R.Ptr:
		if v.IsNil() {
			return Empty
		}
	case R.Slice:
		if v.IsNil() {
			return Empty
		}

		/*
			// This will convert slices to lists.
			// Is this a good idea?
			return terpValue{v: v}.terpList()
		*/

		var pointerToT *T
		switch v.Type().Elem() {
		case R.TypeOf(pointerToT).Elem(): // i.e. case T
			return MkList(v.Interface().([]T))
		}
		switch v.Type().Elem().Kind() {
		case R.Uint8:
			return MkString(string(v.Interface().([]byte)))
		}

	case R.String:
		return MkString(v.String())
	case R.Struct:
	case R.UnsafePointer:
	}

	// Everything else becomes a terpValue
	return MkValue(v)
}

// terpHash implements T

func (t terpHash) Raw() interface{} {
	return t.H
}
func (t terpHash) String() string {
	z := make([]T, 0, 2*len(t.H))
	for k, v := range t.H {
		if v == nil {
			continue
		}
		z = append(z, MkString(k))
		z = append(z, v)
	}
	return MkList(z).String()
}
func (t terpHash) Float() float64 {
	panic("not implemented on terpHash (Float)")
}
func (t terpHash) Int() int64 {
	panic("not implemented on terpHash (Int)")
}
func (t terpHash) Uint() uint64 {
	panic("not implemented on terpHash (Uint)")
}
func (t terpHash) ListElementString() string {
	panic("not implemented on terpHash (ListElementString)")
}
func (t terpHash) Bool() bool {
	panic("terpHash cannot be used as Bool")
}
func (t terpHash) IsEmpty() bool {
	return len(t.H) == 0
}

type SortListByStringTSlice []T

func (p SortListByStringTSlice) Len() int           { return len(p) }
func (p SortListByStringTSlice) Less(i, j int) bool { return p[i].String() < p[j].String() }
func (p SortListByStringTSlice) Swap(i, j int)      { p[j], p[i] = p[i], p[j] }

func SortListByString(list []T) {
	sort.Sort(SortListByStringTSlice(list))
}

func SortedKeysOfHash(h Hash) []string {
	// TODO: mutex
	keys := make([]string, 0, len(h))

	for k, v := range h {
		if v == nil {
			continue // Omit phantoms and deletions.
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (t terpHash) IsPreservedByList() bool { return true }
func (t terpHash) IsQuickNumber() bool     { return false }
func (t terpHash) List() []T {
	keys := SortedKeysOfHash(t.H)
	z := make([]T, 0, 2*len(keys))
	// TODO: mutex
	for _, k := range keys {
		v := t.H[k]
		if v == nil {
			continue // Omit phantoms and deletions.
		}
		z = append(z, MkString(k), v)
	}
	return z
}
func (t terpHash) HeadTail() (hd, tl T) {
	return MkList(t.List()).HeadTail()
}
func (t terpHash) Hash() Hash {
	return t.H
}
func (t terpHash) GetAt(key T) T {
	return t.H[key.String()]
}
func (t terpHash) PutAt(value T, key T) {
	t.H[key.String()] = value
}
func (t terpHash) QuickReflectValue() R.Value { return InvalidValue }

// terpGenerator implements T

func (t terpGenerator) Raw() interface{} {
	panic("not implemented on generator (terpGenerator)")
}
func (t terpGenerator) String() string {
	return Repr(t.guts)
}
func (t terpGenerator) Float() float64 {
	panic("not implemented on generator (terpGenerator)")
}
func (t terpGenerator) Int() int64 {
	panic("not implemented on generator (terpGenerator)")
}
func (t terpGenerator) Uint() uint64 {
	panic("not implemented on generator (terpGenerator)")
}
func (t terpGenerator) ListElementString() string {
	panic("not implemented on generator (terpGenerator)")
}
func (t terpGenerator) Bool() bool {
	panic("terpGenerator cannot be used as Bool")
}
func (t terpGenerator) IsEmpty() bool {
	hd, _ := t.HeadTail()
	return hd == nil
}
func (t terpGenerator) IsPreservedByList() bool { return true }
func (t terpGenerator) IsQuickNumber() bool     { return false }
func (t terpGenerator) List() []T {
	z := make([]T, 0, 4)
	for {
		ei := <-t.guts.readerChan
		if ei.Bad != nil {
			panic(ei.Bad)
		}
		if ei.Good == nil {
			break
		}
		z = append(z, ei.Good)
	}
	return z
}
func (t terpGenerator) HeadTail() (hd, tl T) {
	g := t.guts
	if g.readerChan == nil {
		return g.hd, g.tl
	}

	ei := <-g.readerChan
	if ei.Bad != nil {
		panic(ei.Bad)
	}

	g.hd = ei.Good
	if g.hd == nil {
		g.readerChan = nil
		return nil, nil
	}
	g.tl = terpGenerator{guts: &terpGeneratorGuts{readerChan: g.readerChan}}
	return g.hd, g.tl
}
func (t terpGenerator) Hash() Hash {
	panic("terpGenerator is not a Hash")
}
func (t terpGenerator) GetAt(key T) T {
	panic("terpGenerator is not a Hash")
}
func (t terpGenerator) PutAt(value T, key T) {
	panic("terpGenerator is not a Hash")
}
func (t terpGenerator) QuickReflectValue() R.Value { return InvalidValue }

// terpFloat implements T

func (t terpFloat) Raw() interface{} {
	return t.f
}
func (t terpFloat) String() string {
	return Sprintf("%g", t.f)
}
func (t terpFloat) ListElementString() string {
	return t.String()
}
func (t terpFloat) Bool() bool {
	return t.f != 0
}
func (t terpFloat) IsEmpty() bool {
	return false
}
func (t terpFloat) Float() float64 {
	return t.f
}
func (t terpFloat) Int() int64 {
	return int64(t.f)
}
func (t terpFloat) Uint() uint64 {
	return uint64(t.f)
}
func (t terpFloat) IsPreservedByList() bool { return true }
func (t terpFloat) IsQuickNumber() bool     { return true }
func (t terpFloat) List() []T {
	return []T{t}
}
func (t terpFloat) HeadTail() (hd, tl T) {
	return MkList(t.List()).HeadTail()
}
func (t terpFloat) Hash() Hash {
	panic(" is not a Hash")
}
func (t terpFloat) GetAt(key T) T {
	panic("terpFloat is not a Hash")
}
func (t terpFloat) PutAt(value T, key T) {
	panic("terpFloat is not a Hash")
}
func (t terpFloat) QuickReflectValue() R.Value { return InvalidValue }

// terpString implements T

func (t terpString) Raw() interface{} {
	return t.s
}
func (t terpString) String() string {
	return t.s
}
func (t terpString) ListElementString() string {
	return ToListElementString(t.s)
}
func (t terpString) Bool() bool {
	if t.s == "0" {
		return false
	}
	if t.s == "1" {
		return true
	}
	return MkFloat(t.Float()).Bool()
}
func (t terpString) IsEmpty() bool {
	return t.s == ""
}
func (t terpString) Float() float64 {
	z, err := strconv.ParseFloat(t.s, 64)
	if err != nil {
		panic(err)
	}
	return z
}
func (t terpString) Int() int64 {
	return int64(t.Float()) //TODO
}
func (t terpString) Uint() uint64 {
	return uint64(t.Float()) //TODO
}
func (t terpString) IsQuickNumber() bool { return false }
func (t terpString) IsPreservedByList() bool {
	return nil != MatchBareWord.FindStringSubmatch(t.s)
}
func (t terpString) List() []T {
	if t.IsPreservedByList() {
		return []T{t}
	}
	return ParseList(t.s)
}
func (t terpString) HeadTail() (hd, tl T) {
	return MkList(t.List()).HeadTail()
}
func (t terpString) Hash() Hash {
	panic("A string is not a Hash")
}
func (t terpString) GetAt(key T) T {
	panic("terpString is not a Hash")
}
func (t terpString) PutAt(value T, key T) {
	panic("terpString is not a Hash")
}
func (t terpString) QuickReflectValue() R.Value { return InvalidValue }

// terpList implements T

func (t terpList) Raw() interface{} {
	z := make([]interface{}, len(t.l))
	for i, e := range t.l {
		z[i] = e.Raw() // Recurse.
	}
	return z
}
func (t terpList) String() string {
	z := ""
	for k, v := range t.l {
		if k > 0 {
			z += " "
		}
		z += v.ListElementString()
	}
	return z
}
func (t terpList) ListElementString() string {
	return ToListElementString(t.String())
}
func (t terpList) Bool() bool {
	if len(t.l) == 1 {
		return t.l[0].Bool()
	}
	panic("terpList cannot be used as Bool")
}
func (t terpList) IsEmpty() bool {
	return len(t.l) == 0
}
func (t terpList) Float() float64 {
	if len(t.l) != 1 {
		panic("cant")
	}
	return t.l[0].Float()
}
func (t terpList) Int() int64 {
	if len(t.l) != 1 {
		panic("cant")
	}
	return t.l[0].Int()
}
func (t terpList) Uint() uint64 {
	if len(t.l) != 1 {
		panic("cant")
	}
	return t.l[0].Uint()
}
func (t terpList) IsQuickNumber() bool {
	if len(t.l) == 1 {
		return t.l[0].IsQuickNumber()
	}
	return false
}
func (t terpList) IsPreservedByList() bool { return true }
func (t terpList) List() []T {
	return t.l
}
func (t terpList) HeadTail() (hd, tl T) {
	if len(t.l) == 0 {
		return nil, nil
	}
	return t.l[0], MkList(t.l[1:])
}
func (t terpList) Hash() Hash {
	panic("A List is not a Hash")
}
func (t terpList) GetAt(key T) T {
	panic("terpList is not a Hash")
}
func (t terpList) PutAt(value T, key T) {
	panic("terpList is not a Hash")
}
func (t terpList) QuickReflectValue() R.Value { return InvalidValue }

// terpValue implements T

func (t terpValue) Raw() interface{} {
	return t.v.Interface()
}
func (t terpValue) String() string {
	s := Sprintf("Value:%s:%s", t.v.Kind(), t.v.Type())
	return s
}
func (t terpValue) ListElementString() string {
	return ToListElementString(t.String())
}
func (t terpValue) Bool() bool {
	panic("terpValue cannot be used as Bool")
}
func (t terpValue) IsEmpty() bool {
	switch t.v.Kind() {
	// IsNil() can only be called on this 6 Kinds:
	case R.Func, R.Interface, R.Ptr, R.Slice, R.Map, R.Chan:
		return t.v.IsNil()
	}
	// Strings, numbers, and bools should not be in terpValue so we don't look for emptiness or zeroness in them.
	return false
}
func (t terpValue) Float() float64 {
	panic("cant yet")
}
func (t terpValue) Int() int64 {
	panic("cant yet")
}
func (t terpValue) Uint() uint64 {
	panic("cant yet")
}
func (t terpValue) IsQuickNumber() bool {
	panic("cant yet")
}
func (t terpValue) IsPreservedByList() bool { return true }
func (t terpValue) List() []T {
	/***
		Is this a good idea?

		At times, it is really convenient to have a Raw Slice be a list.

		But other times, we want to edit that Raw Slice in place.

		Maybe this is right -- only when you explicitly ask for a List() do we explode it.

		Is treating Ptr and Iface like a Singleton List a good idea?
	***/
	switch t.v.Kind() {

	// Treat Pointer and Interface as a singleton list.
	case R.Ptr, R.Interface:
		x := MkT(t.v.Elem().Interface())
		return []T{x}

	// Slices and Arrays are naturally lists (unless they're bytes)
	case R.Slice, R.Array:
		if t.v.Type().Elem().Kind() == R.Uint8 {
			panic(Sprintf("Slice of Uint8 should not be in terpValue: %q", string(t.v.Interface().([]byte))))
		}
		n := t.v.Len()
		z := make([]T, n)
		for i := 0; i < n; i++ {
			z[i] = MkT(t.v.Index(i).Interface())
		}
		return z
	}
	/********/
	return []T{t}
}
func (t terpValue) HeadTail() (hd, tl T) {
	return MkList(t.List()).HeadTail()
}

func NeedsOctalEscape(b byte) bool {
	switch b {
	case '{':
		return true
	case '}':
		return true
	case '\\':
		return true
	}
	if b < ' ' {
		return true
	}
	return false
}

func OctalEscape(s string) string {
	needsEscaping := false
	for _, b := range []byte(s) {
		if NeedsOctalEscape(b) {
			needsEscaping = true
			break
		}
	}
	if !needsEscaping {
		return s
	}
	buf := bytes.NewBuffer(nil)
	for _, b := range []byte(s) {
		if NeedsOctalEscape(b) {
			buf.WriteString(Sprintf("\\%03o", b))
		} else {
			buf.WriteByte(b)
		}
	}
	return (buf.String())
}

func ToListElementString(s string) string {
	// TODO: Not perfect, but we are not doing \ yet.
	// TODO: Broken for mismatched {}.
	if s == "" {
		return "{}"
	}

	if strings.ContainsAny(s, " \t\n\r{}\\") {
		return "{" + OctalEscape(s) + "}"
	}
	return s
}
func (t terpValue) Hash() Hash {
	panic("A GoValue is not a Hash")
}
func (t terpValue) GetAt(key T) T {
	panic("terpValue is not a Hash")
}
func (t terpValue) PutAt(value T, key T) {
	panic("terpValue is not a Hash")
}
func (t terpValue) QuickReflectValue() R.Value { return t.v }

func (g *Global) MintMixinSerial() int {
	g.Mu.Lock()
	defer g.Mu.Unlock()

	g.MixinSerial++
	return g.MixinSerial
}

type EnsembleItem struct {
	Name string
	Cmd  Command
	Doc  string
}

func MkEnsemble(items []EnsembleItem) Command {
	cmd := func(fr *Frame, argv []T) T {
		switch len(argv) {
		case 0:
			panic("TODO: doc string")
		case 1:
			panic("Ensemble command needs a subcommand argument: " + Showv(argv))
		}
		subName := argv[1].String()
		for _, e := range items {
			if e.Name == subName {
				return e.Cmd(fr, argv[1:])
			}
		}
		panic("Ensemble subcommand not found: " + Showv(argv))
	}
	return cmd
}

func NonEmpty(v []string) []string {
	z := make([]string, 0, len(v))
	for _, e := range v {
		if len(e) > 1 {
			z = append(z, e)
		}
	}
	return z
}
