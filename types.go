package chirp

import (
	"bytes"
	. "fmt"
	R "reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// T is an interface to any Tcl value.
// Use them only through these methods, or fix these methods.
type T interface {
	ChirpFlavor() string
	Raw() interface{}
	String() string
	Float() float64
	Int() int64
	Uint() uint64
	ListElementString() string
	IsQuickString() bool
	IsQuickList() bool
	IsQuickHash() bool
	Bool() bool    // Like Python, empty values and 0 values are false.
	IsEmpty() bool // Would String() return ""?
	List() []T
	IsPreservedByList() bool
	IsQuickInt() bool
	IsQuickNumber() bool
	HeadTail() (hd, tl T)
	Hash() (Hash, *sync.Mutex)
	GetAt(key T) T
	PutAt(value T, key T)
	QuickReflectValue() R.Value
	EvalSeq(fr *Frame) T
	EvalExpr(fr *Frame) T
	Apply(fr *Frame, args []T) T
}

// terpInt is a Tcl value holding a int64.
type terpInt struct { // Implements T.
	i int64
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

// *terpMulti is a Tcl value holding several pre-compiled representations,
// which were parsed from a string.
type terpMulti struct { // Implements T.
	s               terpString
	preservedByList bool
	i               *terpInt
	f               *terpFloat
	l               *terpList
	seq             *PSeq
	expr            *PExpr
	command         Command
}

func (o *terpMulti) Show() string {
	return Sprintf("MULTI{ s: {%q} i:%v f:%v p:%v seq:%s expr:%s } ", o.s, (o.i != nil), (o.f != nil), o.preservedByList, ShowSeqUnlessNull(o.seq), ShowExprUnlessNull(o.expr))
}
func ShowSeqUnlessNull(seq *PSeq) string {
	if seq == nil {
		return "*nil*"
	}
	return seq.Show()
}
func ShowExprUnlessNull(expr *PExpr) string {
	if expr == nil {
		return "*nil*"
	}
	return expr.Show()
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

// *terpHash holds a Hash.
type terpHash struct { // Implements T.
	h  Hash
	Mu sync.Mutex
}

func MkHash(h Hash) *terpHash {
	MkHashCounter.Incr()
	if h == nil {
		return &terpHash{h: make(Hash, 4)}
	}
	return &terpHash{h: h}
}
func MkGenerator(readerChan <-chan Either) terpGenerator {
	MkGeneratorCounter.Incr()
	return terpGenerator{guts: &terpGeneratorGuts{readerChan: readerChan}}
}
func MkBool(a bool) T {
	MkBoolCounter.Incr()
	if a {
		return True
	}
	return False
}
func MkNum(s string) T {
	MkNumCounter.Incr()
	if strings.IndexByte(s, '.') >= 0 {
		return MkFloat(MkString(s).Float())
	}
	return MkInt(MkString(s).Int())
}
func MkFloat(a float64) terpFloat {
	MkFloatCounter.Incr()
	return terpFloat{f: a}
}
func MkInt(a int64) terpInt {
	MkIntCounter.Incr()
	return terpInt{i: int64(a)}
}
func MkUint(a uint64) terpInt {
	// We dont have a terpUint type; jam it into a signed terpInt.
	MkUintCounter.Incr()
	return terpInt{i: int64(a)}
}
func MkString(a string) terpString {
	MkStringCounter.Incr()
	return terpString{s: a}
}
func MkList(a []T) terpList {
	MkListCounter.Incr()
	return terpList{l: a}
}
func MkStringList(a []string) terpList {
	MkStringListCounter.Incr()
	z := make([]T, len(a))
	for i, e := range a {
		z[i] = MkString(e)
	}
	return terpList{l: z}
}
func MkValue(a R.Value) terpValue {
	MkValueCounter.Incr()
	return terpValue{v: a}
}
func MaybeCompileSequence(fr *Frame, s string) (seq *PSeq) {
	defer func() {
		recover()
	}()
	seq = CompileSequence(fr, s)
	return
}
func MkMultiFr(fr *Frame, a *terpMulti) *terpMulti {
	//println("MkMultiFr <<<<<<", a.Show())
	m := MkMulti(a.s.s)
	m.seq = MaybeCompileSequence(fr, a.s.s)
	//println("MkMultiFr <<<<<<", a.Show(), ">>>>>>", m.Show())
	return m
}
func MkMulti(s string) *terpMulti {
	MkMultiCounter.Incr()
	var ts terpString = MkString(s)
	m := &terpMulti{
		s:               ts,
		preservedByList: ts.IsPreservedByList(),
	}

	func() {
		defer func() {
			_ = recover()
		}()
		x := MkInt(ts.Int())
		m.i = &x
	}()

	func() {
		defer func() {
			_ = recover()
		}()
		x := MkFloat(ts.Float())
		m.f = &x
	}()

	func() {
		defer func() {
			_ = recover()
		}()
		x := MkList(ts.List())
		m.l = &x
	}()

	// This is why you cannot rename builtins.
	m.command = Safes[ts.s]

	return m
}

func MkT(a interface{}) T {
	MkTCounter.Incr()
	// Very specific type cases.
	switch x := a.(type) {
	case T:
		// panic(Sprintf("Calling MkT() on a T: <%T> <%#v> %s", x, x, x.String()))
		return x

	// case R.Value:
	// 	// Some day we'll allow this, but for now, flag an error.
	// 	panic(Sprintf("Calling MkT() on a R.Value: <%T> <%#v> %s", x, x, x.String()))

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

		//	// This will convert slices to lists.
		//	// Is this a good idea?
		//	return terpValue{v: v}.terpList()

		switch v.Type().Elem() {
		case TypeT:
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

// *terpHash implements T

func (t *terpHash) ChirpFlavor() string { return "Hash" }
func (t *terpHash) Raw() interface{}    { return t.h }
func (t *terpHash) String() string {
	return MkList(t.List()).String()
}
func (t *terpHash) Float() float64 {
	panic("not implemented on terpHash (Float)")
}
func (t *terpHash) Int() int64 {
	panic("not implemented on terpHash (Int)")
}
func (t *terpHash) Uint() uint64 {
	panic("not implemented on terpHash (Uint)")
}
func (t *terpHash) ListElementString() string {
	return MkString(t.String()).ListElementString()
}
func (t *terpHash) IsQuickString() bool {
	return false
}
func (t *terpHash) IsQuickList() bool {
	return false
}
func (t *terpHash) IsQuickHash() bool {
	return true
}
func (t *terpHash) Bool() bool {
	panic("terpHash cannot be used as Bool")
}
func (t *terpHash) IsEmpty() bool {
	t.Mu.Lock()
	z := (len(t.h) == 0)
	t.Mu.Unlock()
	return z
}

type SortListByStringTSlice []T

func (p SortListByStringTSlice) Len() int           { return len(p) }
func (p SortListByStringTSlice) Less(i, j int) bool { return p[i].String() < p[j].String() }
func (p SortListByStringTSlice) Swap(i, j int)      { p[j], p[i] = p[i], p[j] }

// SortListByString is used by smilax-web/db.
func SortListByString(list []T) {
	sort.Sort(SortListByStringTSlice(list))
}

func SortedKeysOfHash(h Hash) []string {
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

func (t *terpHash) IsPreservedByList() bool { return true }
func (t *terpHash) IsQuickInt() bool        { return false }
func (t *terpHash) IsQuickNumber() bool     { return false }
func (t *terpHash) List() []T {
	t.Mu.Lock()
	keys := SortedKeysOfHash(t.h)
	z := make([]T, 0, 2*len(keys))

	for _, k := range keys {
		v := t.h[k]
		z = append(z, MkString(k), v)
	}
	t.Mu.Unlock()
	return z
}
func (t *terpHash) HeadTail() (hd, tl T) {
	return MkList(t.List()).HeadTail()
}
func (t *terpHash) Hash() (Hash, *sync.Mutex) {
	return t.h, &t.Mu
}
func (t *terpHash) GetAt(key T) T {
	k := key.String()

	t.Mu.Lock()
	z := t.h[k]
	t.Mu.Unlock()

	return z
}
func (t *terpHash) PutAt(value T, key T) {
	k := key.String()

	t.Mu.Lock()
	t.h[k] = value
	t.Mu.Unlock()
}
func (t *terpHash) QuickReflectValue() R.Value  { return InvalidValue }
func (t *terpHash) EvalSeq(fr *Frame) T         { return Parse2EvalSeqStr(fr, t.String()) }
func (t *terpHash) EvalExpr(fr *Frame) T        { return Parse2EvalExprStr(fr, t.String()) }
func (t *terpHash) Apply(fr *Frame, args []T) T { panic("Cannot apply terpHash as command") }

// terpGenerator implements T

func (t terpGenerator) ChirpFlavor() string { return "Generator" }
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
func (t terpGenerator) IsQuickString() bool {
	return false
}
func (t terpGenerator) IsQuickList() bool {
	return true
}
func (t terpGenerator) IsQuickHash() bool {
	return false
}
func (t terpGenerator) Bool() bool {
	panic("terpGenerator cannot be used as Bool")
}
func (t terpGenerator) IsEmpty() bool {
	hd, _ := t.HeadTail()
	return hd == nil
}
func (t terpGenerator) IsPreservedByList() bool { return true }
func (t terpGenerator) IsQuickInt() bool        { return false }
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
func (t terpGenerator) Hash() (Hash, *sync.Mutex) {
	panic("terpGenerator is not a Hash")
}
func (t terpGenerator) GetAt(key T) T {
	panic("terpGenerator is not a Hash")
}
func (t terpGenerator) PutAt(value T, key T) {
	panic("terpGenerator is not a Hash")
}
func (t terpGenerator) QuickReflectValue() R.Value  { return InvalidValue }
func (t terpGenerator) EvalSeq(fr *Frame) T         { return Parse2EvalSeqStr(fr, t.String()) }
func (t terpGenerator) EvalExpr(fr *Frame) T        { return Parse2EvalExprStr(fr, t.String()) }
func (t terpGenerator) Apply(fr *Frame, args []T) T { panic("Cannot apply terpGenerator as command") }

// terpInt implements T

func (t terpInt) ChirpFlavor() string { return "Float" }
func (t terpInt) Raw() interface{} {
	return t.i
}
func (t terpInt) String() string {
	return Sprintf("%d", t.i)
}
func (t terpInt) ListElementString() string {
	return t.String()
}
func (t terpInt) IsQuickString() bool {
	return false
}
func (t terpInt) IsQuickList() bool {
	return false
}
func (t terpInt) IsQuickHash() bool {
	return false
}
func (t terpInt) Bool() bool {
	return t.i != 0
}
func (t terpInt) IsEmpty() bool {
	return false
}
func (t terpInt) Float() float64 {
	return float64(t.i)
}
func (t terpInt) Int() int64 {
	return int64(t.i)
}
func (t terpInt) Uint() uint64 {
	return uint64(t.i)
}
func (t terpInt) IsPreservedByList() bool { return true }
func (t terpInt) IsQuickInt() bool        { return true }
func (t terpInt) IsQuickNumber() bool     { return true }
func (t terpInt) List() []T {
	return []T{t}
}
func (t terpInt) HeadTail() (hd, tl T) {
	return MkList(t.List()).HeadTail()
}
func (t terpInt) Hash() (Hash, *sync.Mutex) {
	panic(" is not a Hash")
}
func (t terpInt) GetAt(key T) T {
	panic("terpInt is not a Hash")
}
func (t terpInt) PutAt(value T, key T) {
	panic("terpInt is not a Hash")
}
func (t terpInt) QuickReflectValue() R.Value  { return InvalidValue }
func (t terpInt) EvalSeq(fr *Frame) T         { return Parse2EvalSeqStr(fr, t.String()) }
func (t terpInt) EvalExpr(fr *Frame) T        { return t } // Numbers are self-Expr-eval'ing.
func (t terpInt) Apply(fr *Frame, args []T) T { return fr.Apply(args) }

// terpFloat implements T

func (t terpFloat) ChirpFlavor() string { return "Float" }
func (t terpFloat) Raw() interface{} {
	return t.f
}
func (t terpFloat) String() string {
	return Sprintf("%.15g", t.f)
}
func (t terpFloat) ListElementString() string {
	return t.String()
}
func (t terpFloat) IsQuickString() bool {
	return false
}
func (t terpFloat) IsQuickList() bool {
	return false
}
func (t terpFloat) IsQuickHash() bool {
	return false
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
func (t terpFloat) IsQuickInt() bool        { return false }
func (t terpFloat) IsQuickNumber() bool     { return true }
func (t terpFloat) List() []T {
	return []T{t}
}
func (t terpFloat) HeadTail() (hd, tl T) {
	return MkList(t.List()).HeadTail()
}
func (t terpFloat) Hash() (Hash, *sync.Mutex) {
	panic(" is not a Hash")
}
func (t terpFloat) GetAt(key T) T {
	panic("terpFloat is not a Hash")
}
func (t terpFloat) PutAt(value T, key T) {
	panic("terpFloat is not a Hash")
}
func (t terpFloat) QuickReflectValue() R.Value  { return InvalidValue }
func (t terpFloat) EvalSeq(fr *Frame) T         { return Parse2EvalSeqStr(fr, t.String()) }
func (t terpFloat) EvalExpr(fr *Frame) T        { return t } // Numbers are self-Expr-eval'ing.
func (t terpFloat) Apply(fr *Frame, args []T) T { return fr.Apply(args) }

// terpString implements T

func (t terpString) ChirpFlavor() string { return "String" }
func (t terpString) Raw() interface{} {
	return t.s
}
func (t terpString) String() string {
	return t.s
}
func (t terpString) ListElementString() string {
	return ToListElementString(t.s)
}
func (t terpString) IsQuickString() bool {
	return true
}
func (t terpString) IsQuickList() bool {
	return false
}
func (t terpString) IsQuickHash() bool {
	return false
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
	z, err := strconv.ParseInt(t.s, 10, 64)
	if err != nil {
		panic(err)
	}
	return z
}
func (t terpString) Uint() uint64 {
	// First try unsigned parse.
	z, err := strconv.ParseUint(t.s, 10, 64)
	if err == nil {
		return z
	}
	// Then try signed parse.
	return uint64(t.Int())
}
func (t terpString) IsQuickInt() bool    { return false }
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
func (t terpString) Hash() (Hash, *sync.Mutex) {
	panic("A string is not a Hash")
}
func (t terpString) GetAt(key T) T {
	panic("terpString is not a Hash")
}
func (t terpString) PutAt(value T, key T) {
	panic("terpString is not a Hash")
}
func (t terpString) QuickReflectValue() R.Value  { return InvalidValue }
func (t terpString) EvalSeq(fr *Frame) T         { return Parse2EvalSeqStr(fr, t.String()) }
func (t terpString) EvalExpr(fr *Frame) T        { return Parse2EvalExprStr(fr, t.String()) }
func (t terpString) Apply(fr *Frame, args []T) T { return fr.Apply(args) }

// terpList implements T

func (t terpList) ChirpFlavor() string { return "List" }
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
func (t terpList) IsQuickString() bool {
	return false
}
func (t terpList) IsQuickList() bool {
	return false
}
func (t terpList) IsQuickHash() bool {
	return false
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
func (t terpList) IsQuickInt() bool {
	if len(t.l) == 1 {
		return t.l[0].IsQuickInt()
	}
	return false
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
func (t terpList) Hash() (Hash, *sync.Mutex) {
	panic("A List is not a Hash")
}
func (t terpList) GetAt(key T) T {
	panic("terpList is not a Hash")
}
func (t terpList) PutAt(value T, key T) {
	panic("terpList is not a Hash")
}
func (t terpList) QuickReflectValue() R.Value { return InvalidValue }

// Bug.3
//NO// func (t terpList) EvalSeq(fr *Frame) T         { return Parse2EvalSeqStr(fr, t.String()) }
func (t terpList) EvalSeq(fr *Frame) T { return fr.Apply(t.l) }

func (t terpList) EvalExpr(fr *Frame) T        { return Parse2EvalExprStr(fr, t.String()) }
func (t terpList) Apply(fr *Frame, args []T) T { return fr.Apply(args) }

// terpValue implements T

func (t terpValue) ChirpFlavor() string { return "Value" }
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
func (t terpValue) IsQuickString() bool {
	return false
}
func (t terpValue) IsQuickList() bool {
	return false
}
func (t terpValue) IsQuickHash() bool {
	return false
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

var Float64Type R.Type = R.TypeOf(*new(float64))
var Int64Type R.Type = R.TypeOf(*new(int64))
var Uint64Type R.Type = R.TypeOf(*new(uint64))

func (t terpValue) Float() float64 {
	return t.v.Convert(Float64Type).Interface().(float64)
}
func (t terpValue) Int() int64 {
	return t.v.Convert(Int64Type).Interface().(int64)
}
func (t terpValue) Uint() uint64 {
	return t.v.Convert(Uint64Type).Interface().(uint64)
}
func (t terpValue) IsQuickInt() bool {
	return t.v.Type().ConvertibleTo(Int64Type)
}
func (t terpValue) IsQuickNumber() bool {
	return t.v.Type().ConvertibleTo(Float64Type)
}
func (t terpValue) IsPreservedByList() bool { return false }
func (t terpValue) List() []T {
	switch t.v.Kind() {

	/*
		// Treat Pointer and Interface as a singleton list.
		case R.Ptr, R.Interface:
			x := MkT(t.v.Elem().Interface())
			return []T{x}
	*/

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
	panic("terpValue: not a list")
}
func (t terpValue) HeadTail() (hd, tl T) {
	return MkList(t.List()).HeadTail()
}

///////////////////////////////////////////////////////////////////////
// *terpMulti implements T

func (t terpMulti) ChirpFlavor() string { return "Multi" }
func (t *terpMulti) Raw() interface{} {
	return t.s.Raw()
}
func (t *terpMulti) String() string {
	return t.s.String()
}
func (t *terpMulti) ListElementString() string {
	return t.s.ListElementString()
}
func (t *terpMulti) IsQuickString() bool {
	return true
}
func (t *terpMulti) IsQuickList() bool {
	return t.l != nil
}
func (t *terpMulti) IsQuickHash() bool {
	return false
}
func (t *terpMulti) Bool() bool {
	if t.f != nil {
		return t.f.Bool()
	}
	return t.s.Bool()
}
func (t *terpMulti) IsEmpty() bool {
	if t.l != nil {
		return t.l.IsEmpty()
	}
	return t.s.IsEmpty()
}
func (t *terpMulti) Float() float64 {
	if t.f != nil {
		return t.f.Float()
	}
	return t.s.Float()
}
func (t *terpMulti) Int() int64 {
	if t.f != nil {
		return t.f.Int()
	}
	return t.s.Int()
}
func (t *terpMulti) Uint() uint64 {
	if t.f != nil {
		return t.f.Uint()
	}
	return t.s.Uint()
}
func (t *terpMulti) IsQuickInt() bool {
	if t.i != nil {
		return t.i.IsQuickInt()
	}
	return t.s.IsQuickInt()
}
func (t *terpMulti) IsQuickNumber() bool {
	if t.f != nil {
		return t.f.IsQuickNumber()
	}
	return t.s.IsQuickNumber()
}
func (t *terpMulti) IsPreservedByList() bool {
	return t.preservedByList
}
func (t *terpMulti) List() []T {
	if t.l != nil {
		return t.l.List()
	}
	return t.s.List()
}
func (t *terpMulti) HeadTail() (hd, tl T) {
	if t.l != nil {
		return t.l.HeadTail()
	}
	return t.s.HeadTail()
}
func (t *terpMulti) Hash() (Hash, *sync.Mutex) {
	panic("terpMulti: is not a Hash")
}
func (t *terpMulti) GetAt(key T) T {
	panic("terpMulti: is not a Hash")
}
func (t *terpMulti) PutAt(value T, key T) {
	panic("terpMulti: is not a Hash")
}
func (t *terpMulti) QuickReflectValue() R.Value { return InvalidValue }
func (t *terpMulti) EvalSeq(fr *Frame) T {
	MultiEvalSeqCounter.Incr()
	if t.seq == nil {
		MultiEvalSeqCompileCounter.Incr()
		// Lazily compile the first time it is eval'ed as a Seq.
		t.seq = Parse2SeqStr(t.s.s)
	}
	return fr.EvalSeqWithErrorLocation(t.seq)
}
func (t *terpMulti) EvalExpr(fr *Frame) T {
	MultiEvalExprCounter.Incr()
	if t.expr == nil {
		MultiEvalExprCompileCounter.Incr()
		// Lazily compile the first time it is eval'ed as an Expr.
		t.expr = Parse2ExprStr(t.s.s)
	}
	return t.expr.Eval(fr)
}
func (t terpMulti) Apply(fr *Frame, args []T) T {
	if t.command != nil {
		defer func() {
			if r := recover(); r != nil {
				if re, ok := r.(error); ok {
					r = re.Error() // Convert error to string.
				}
				if rs, ok := r.(string); ok {
					rs = rs + Sprintf("\n\tin (terpMulti)Apply\n\t\t%q", args[0])

					// TODO: Require debug level for the args.
					for _, ae := range args[1:] {
						as := ae.String()
						if len(as) > 40 {
							as = as[:40] + "..."
						}
						rs = rs + Sprintf(" %q", as)
					}

					if fr.MixinLevel > 0 {
						rs = rs + Sprintf("\n\t\t(frame's MixinLevel=%d)", fr.MixinLevel)
					}
					if len(fr.MixinName) > 0 {
						rs = rs + Sprintf("\n\t\t(frame's MixinName=%q)", fr.MixinName)
					}
					r = rs
				}
				panic(r)
			}
		}()

		return t.command(fr, args)
	}
	return fr.Apply(args)
}

///////////////////////////////////////////////////////////////////////

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
func (t terpValue) Hash() (Hash, *sync.Mutex) {
	panic("A GoValue is not a Hash")
}
func (t terpValue) GetAt(key T) T {
	panic("terpValue is not a Hash")
}
func (t terpValue) PutAt(value T, key T) {
	panic("terpValue is not a Hash")
}
func (t terpValue) QuickReflectValue() R.Value  { return t.v }
func (t terpValue) EvalSeq(fr *Frame) T         { return Parse2EvalSeqStr(fr, t.String()) }
func (t terpValue) EvalExpr(fr *Frame) T        { return Parse2EvalExprStr(fr, t.String()) }
func (t terpValue) Apply(fr *Frame, args []T) T { return ApplyToReflectedValue(fr, t.v, args, 1) }

////////////////////////////////////////

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

func ShowEnsembleItems(items []EnsembleItem) string {
	z := ""
	for _, e := range items {
		z += " " + e.Name
	}
	return z
}

func MkEnsemble(items []EnsembleItem) Command {
	cmd := func(fr *Frame, argv []T) T {
		switch len(argv) {
		case 0:
			panic("TODO: doc string")
		case 1:
			panic(Sprintf("Ensemble options: %s", ShowEnsembleItems(items)))
		}
		subName := argv[1].String()
		// Try for exact match.
		for _, e := range items {
			if e.Name == subName {
				return e.Cmd(fr, argv[1:])
			}
		}
		// Failing exact match, try for prefix match.
		found := -1
		for i, e := range items {
			if len(subName) < len(e.Name) && e.Name[:len(subName)] == subName {
				if found < 0 {
					found = i
				} else {
					panic(Sprintf("Ensemble subcommand ambiguous: %#v Options: %s",
						subName, ShowEnsembleItems(items)))
				}
			}
		}
		if found >= 0 {
			return items[found].Cmd(fr, argv[1:])
		}
		panic(Sprintf("Ensemble subcommand not found: %#v Options: %s",
			subName, ShowEnsembleItems(items)))
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

var MultiEvalSeqCounter Counter
var MultiEvalSeqCompileCounter Counter
var MultiEvalExprCounter Counter
var MultiEvalExprCompileCounter Counter
var MkHashCounter Counter
var MkGeneratorCounter Counter
var MkBoolCounter Counter
var MkNumCounter Counter
var MkFloatCounter Counter
var MkHackFloatCounter Counter
var MkIntCounter Counter
var MkUintCounter Counter
var MkStringCounter Counter
var MkListCounter Counter
var MkStringListCounter Counter
var MkValueCounter Counter
var MkMultiCounter Counter
var MkTCounter Counter

func init() {
	MultiEvalSeqCounter.Register("MultiEvalSeq")
	MultiEvalSeqCompileCounter.Register("MultiEvalSeqCompile")
	MultiEvalExprCounter.Register("MultiEvalExpr")
	MultiEvalExprCompileCounter.Register("MultiEvalExprCompile")
	MkHashCounter.Register("MkHash")
	MkGeneratorCounter.Register("MkGenerator")
	MkBoolCounter.Register("MkBool")
	MkNumCounter.Register("MkNum")
	MkFloatCounter.Register("MkFloat")
	MkHackFloatCounter.Register("MkHackFloat")
	MkIntCounter.Register("MkInt")
	MkUintCounter.Register("MkUint")
	MkStringCounter.Register("MkString")
	MkListCounter.Register("MkList")
	MkStringListCounter.Register("MkStringList")
	MkValueCounter.Register("MkValue")
	MkMultiCounter.Register("MkMulti")
	MkTCounter.Register("MkT")
}
