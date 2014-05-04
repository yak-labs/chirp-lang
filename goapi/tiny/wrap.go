package goapi

import . `github.com/yak-labs/chirp-lang`
import (
	bytes `bytes`
	fmt `fmt`
	reflect `reflect`
	strconv `strconv`
	strings `strings`
)

func init() {
	Roots[`/bytes/Compare`] = FuncRoot{ Func: reflect.ValueOf(bytes.Compare) }
	Roots[`/bytes/Contains`] = FuncRoot{ Func: reflect.ValueOf(bytes.Contains) }
	Roots[`/bytes/Count`] = FuncRoot{ Func: reflect.ValueOf(bytes.Count) }
	Roots[`/bytes/Equal`] = FuncRoot{ Func: reflect.ValueOf(bytes.Equal) }
	Roots[`/bytes/EqualFold`] = FuncRoot{ Func: reflect.ValueOf(bytes.EqualFold) }
	Roots[`/bytes/Fields`] = FuncRoot{ Func: reflect.ValueOf(bytes.Fields) }
	Roots[`/bytes/FieldsFunc`] = FuncRoot{ Func: reflect.ValueOf(bytes.FieldsFunc) }
	Roots[`/bytes/HasPrefix`] = FuncRoot{ Func: reflect.ValueOf(bytes.HasPrefix) }
	Roots[`/bytes/HasSuffix`] = FuncRoot{ Func: reflect.ValueOf(bytes.HasSuffix) }
	Roots[`/bytes/Index`] = FuncRoot{ Func: reflect.ValueOf(bytes.Index) }
	Roots[`/bytes/IndexAny`] = FuncRoot{ Func: reflect.ValueOf(bytes.IndexAny) }
	Roots[`/bytes/IndexByte`] = FuncRoot{ Func: reflect.ValueOf(bytes.IndexByte) }
	Roots[`/bytes/IndexFunc`] = FuncRoot{ Func: reflect.ValueOf(bytes.IndexFunc) }
	Roots[`/bytes/IndexRune`] = FuncRoot{ Func: reflect.ValueOf(bytes.IndexRune) }
	Roots[`/bytes/Join`] = FuncRoot{ Func: reflect.ValueOf(bytes.Join) }
	Roots[`/bytes/LastIndex`] = FuncRoot{ Func: reflect.ValueOf(bytes.LastIndex) }
	Roots[`/bytes/LastIndexAny`] = FuncRoot{ Func: reflect.ValueOf(bytes.LastIndexAny) }
	Roots[`/bytes/LastIndexFunc`] = FuncRoot{ Func: reflect.ValueOf(bytes.LastIndexFunc) }
	Roots[`/bytes/Map`] = FuncRoot{ Func: reflect.ValueOf(bytes.Map) }
	Roots[`/bytes/NewBuffer`] = FuncRoot{ Func: reflect.ValueOf(bytes.NewBuffer) }
	Roots[`/bytes/NewBufferString`] = FuncRoot{ Func: reflect.ValueOf(bytes.NewBufferString) }
	Roots[`/bytes/NewReader`] = FuncRoot{ Func: reflect.ValueOf(bytes.NewReader) }
	Roots[`/bytes/Repeat`] = FuncRoot{ Func: reflect.ValueOf(bytes.Repeat) }
	Roots[`/bytes/Replace`] = FuncRoot{ Func: reflect.ValueOf(bytes.Replace) }
	Roots[`/bytes/Runes`] = FuncRoot{ Func: reflect.ValueOf(bytes.Runes) }
	Roots[`/bytes/Split`] = FuncRoot{ Func: reflect.ValueOf(bytes.Split) }
	Roots[`/bytes/SplitAfter`] = FuncRoot{ Func: reflect.ValueOf(bytes.SplitAfter) }
	Roots[`/bytes/SplitAfterN`] = FuncRoot{ Func: reflect.ValueOf(bytes.SplitAfterN) }
	Roots[`/bytes/SplitN`] = FuncRoot{ Func: reflect.ValueOf(bytes.SplitN) }
	Roots[`/bytes/Title`] = FuncRoot{ Func: reflect.ValueOf(bytes.Title) }
	Roots[`/bytes/ToLower`] = FuncRoot{ Func: reflect.ValueOf(bytes.ToLower) }
	Roots[`/bytes/ToLowerSpecial`] = FuncRoot{ Func: reflect.ValueOf(bytes.ToLowerSpecial) }
	Roots[`/bytes/ToTitle`] = FuncRoot{ Func: reflect.ValueOf(bytes.ToTitle) }
	Roots[`/bytes/ToTitleSpecial`] = FuncRoot{ Func: reflect.ValueOf(bytes.ToTitleSpecial) }
	Roots[`/bytes/ToUpper`] = FuncRoot{ Func: reflect.ValueOf(bytes.ToUpper) }
	Roots[`/bytes/ToUpperSpecial`] = FuncRoot{ Func: reflect.ValueOf(bytes.ToUpperSpecial) }
	Roots[`/bytes/Trim`] = FuncRoot{ Func: reflect.ValueOf(bytes.Trim) }
	Roots[`/bytes/TrimFunc`] = FuncRoot{ Func: reflect.ValueOf(bytes.TrimFunc) }
	Roots[`/bytes/TrimLeft`] = FuncRoot{ Func: reflect.ValueOf(bytes.TrimLeft) }
	Roots[`/bytes/TrimLeftFunc`] = FuncRoot{ Func: reflect.ValueOf(bytes.TrimLeftFunc) }
	Roots[`/bytes/TrimRight`] = FuncRoot{ Func: reflect.ValueOf(bytes.TrimRight) }
	Roots[`/bytes/TrimRightFunc`] = FuncRoot{ Func: reflect.ValueOf(bytes.TrimRightFunc) }
	Roots[`/bytes/TrimSpace`] = FuncRoot{ Func: reflect.ValueOf(bytes.TrimSpace) }
	Roots[`/fmt/Errorf`] = FuncRoot{ Func: reflect.ValueOf(fmt.Errorf) }
	Roots[`/fmt/Fprint`] = FuncRoot{ Func: reflect.ValueOf(fmt.Fprint) }
	Roots[`/fmt/Fprintf`] = FuncRoot{ Func: reflect.ValueOf(fmt.Fprintf) }
	Roots[`/fmt/Fprintln`] = FuncRoot{ Func: reflect.ValueOf(fmt.Fprintln) }
	Roots[`/fmt/Fscan`] = FuncRoot{ Func: reflect.ValueOf(fmt.Fscan) }
	Roots[`/fmt/Fscanf`] = FuncRoot{ Func: reflect.ValueOf(fmt.Fscanf) }
	Roots[`/fmt/Fscanln`] = FuncRoot{ Func: reflect.ValueOf(fmt.Fscanln) }
	Roots[`/fmt/Print`] = FuncRoot{ Func: reflect.ValueOf(fmt.Print) }
	Roots[`/fmt/Printf`] = FuncRoot{ Func: reflect.ValueOf(fmt.Printf) }
	Roots[`/fmt/Println`] = FuncRoot{ Func: reflect.ValueOf(fmt.Println) }
	Roots[`/fmt/Scan`] = FuncRoot{ Func: reflect.ValueOf(fmt.Scan) }
	Roots[`/fmt/Scanf`] = FuncRoot{ Func: reflect.ValueOf(fmt.Scanf) }
	Roots[`/fmt/Scanln`] = FuncRoot{ Func: reflect.ValueOf(fmt.Scanln) }
	Roots[`/fmt/Sprint`] = FuncRoot{ Func: reflect.ValueOf(fmt.Sprint) }
	Roots[`/fmt/Sprintf`] = FuncRoot{ Func: reflect.ValueOf(fmt.Sprintf) }
	Roots[`/fmt/Sprintln`] = FuncRoot{ Func: reflect.ValueOf(fmt.Sprintln) }
	Roots[`/fmt/Sscan`] = FuncRoot{ Func: reflect.ValueOf(fmt.Sscan) }
	Roots[`/fmt/Sscanf`] = FuncRoot{ Func: reflect.ValueOf(fmt.Sscanf) }
	Roots[`/fmt/Sscanln`] = FuncRoot{ Func: reflect.ValueOf(fmt.Sscanln) }
	Roots[`/reflect/Append`] = FuncRoot{ Func: reflect.ValueOf(reflect.Append) }
	Roots[`/reflect/AppendSlice`] = FuncRoot{ Func: reflect.ValueOf(reflect.AppendSlice) }
	Roots[`/reflect/Copy`] = FuncRoot{ Func: reflect.ValueOf(reflect.Copy) }
	Roots[`/reflect/DeepEqual`] = FuncRoot{ Func: reflect.ValueOf(reflect.DeepEqual) }
	Roots[`/reflect/Indirect`] = FuncRoot{ Func: reflect.ValueOf(reflect.Indirect) }
	Roots[`/reflect/MakeChan`] = FuncRoot{ Func: reflect.ValueOf(reflect.MakeChan) }
	Roots[`/reflect/MakeMap`] = FuncRoot{ Func: reflect.ValueOf(reflect.MakeMap) }
	Roots[`/reflect/MakeSlice`] = FuncRoot{ Func: reflect.ValueOf(reflect.MakeSlice) }
	Roots[`/reflect/New`] = FuncRoot{ Func: reflect.ValueOf(reflect.New) }
	Roots[`/reflect/NewAt`] = FuncRoot{ Func: reflect.ValueOf(reflect.NewAt) }
	Roots[`/reflect/PtrTo`] = FuncRoot{ Func: reflect.ValueOf(reflect.PtrTo) }
	Roots[`/reflect/TypeOf`] = FuncRoot{ Func: reflect.ValueOf(reflect.TypeOf) }
	Roots[`/reflect/ValueOf`] = FuncRoot{ Func: reflect.ValueOf(reflect.ValueOf) }
	Roots[`/reflect/Zero`] = FuncRoot{ Func: reflect.ValueOf(reflect.Zero) }
	Roots[`/strconv/AppendBool`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendBool) }
	Roots[`/strconv/AppendFloat`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendFloat) }
	Roots[`/strconv/AppendInt`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendInt) }
	Roots[`/strconv/AppendQuote`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendQuote) }
	Roots[`/strconv/AppendQuoteRune`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendQuoteRune) }
	Roots[`/strconv/AppendQuoteRuneToASCII`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendQuoteRuneToASCII) }
	Roots[`/strconv/AppendQuoteToASCII`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendQuoteToASCII) }
	Roots[`/strconv/AppendUint`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendUint) }
	Roots[`/strconv/Atoi`] = FuncRoot{ Func: reflect.ValueOf(strconv.Atoi) }
	Roots[`/strconv/CanBackquote`] = FuncRoot{ Func: reflect.ValueOf(strconv.CanBackquote) }
	Roots[`/strconv/FormatBool`] = FuncRoot{ Func: reflect.ValueOf(strconv.FormatBool) }
	Roots[`/strconv/FormatFloat`] = FuncRoot{ Func: reflect.ValueOf(strconv.FormatFloat) }
	Roots[`/strconv/FormatInt`] = FuncRoot{ Func: reflect.ValueOf(strconv.FormatInt) }
	Roots[`/strconv/FormatUint`] = FuncRoot{ Func: reflect.ValueOf(strconv.FormatUint) }
	Roots[`/strconv/IsPrint`] = FuncRoot{ Func: reflect.ValueOf(strconv.IsPrint) }
	Roots[`/strconv/Itoa`] = FuncRoot{ Func: reflect.ValueOf(strconv.Itoa) }
	Roots[`/strconv/ParseBool`] = FuncRoot{ Func: reflect.ValueOf(strconv.ParseBool) }
	Roots[`/strconv/ParseFloat`] = FuncRoot{ Func: reflect.ValueOf(strconv.ParseFloat) }
	Roots[`/strconv/ParseInt`] = FuncRoot{ Func: reflect.ValueOf(strconv.ParseInt) }
	Roots[`/strconv/ParseUint`] = FuncRoot{ Func: reflect.ValueOf(strconv.ParseUint) }
	Roots[`/strconv/Quote`] = FuncRoot{ Func: reflect.ValueOf(strconv.Quote) }
	Roots[`/strconv/QuoteRune`] = FuncRoot{ Func: reflect.ValueOf(strconv.QuoteRune) }
	Roots[`/strconv/QuoteRuneToASCII`] = FuncRoot{ Func: reflect.ValueOf(strconv.QuoteRuneToASCII) }
	Roots[`/strconv/QuoteToASCII`] = FuncRoot{ Func: reflect.ValueOf(strconv.QuoteToASCII) }
	Roots[`/strconv/Unquote`] = FuncRoot{ Func: reflect.ValueOf(strconv.Unquote) }
	Roots[`/strconv/UnquoteChar`] = FuncRoot{ Func: reflect.ValueOf(strconv.UnquoteChar) }
	Roots[`/strings/Contains`] = FuncRoot{ Func: reflect.ValueOf(strings.Contains) }
	Roots[`/strings/ContainsAny`] = FuncRoot{ Func: reflect.ValueOf(strings.ContainsAny) }
	Roots[`/strings/ContainsRune`] = FuncRoot{ Func: reflect.ValueOf(strings.ContainsRune) }
	Roots[`/strings/Count`] = FuncRoot{ Func: reflect.ValueOf(strings.Count) }
	Roots[`/strings/EqualFold`] = FuncRoot{ Func: reflect.ValueOf(strings.EqualFold) }
	Roots[`/strings/Fields`] = FuncRoot{ Func: reflect.ValueOf(strings.Fields) }
	Roots[`/strings/FieldsFunc`] = FuncRoot{ Func: reflect.ValueOf(strings.FieldsFunc) }
	Roots[`/strings/HasPrefix`] = FuncRoot{ Func: reflect.ValueOf(strings.HasPrefix) }
	Roots[`/strings/HasSuffix`] = FuncRoot{ Func: reflect.ValueOf(strings.HasSuffix) }
	Roots[`/strings/Index`] = FuncRoot{ Func: reflect.ValueOf(strings.Index) }
	Roots[`/strings/IndexAny`] = FuncRoot{ Func: reflect.ValueOf(strings.IndexAny) }
	Roots[`/strings/IndexFunc`] = FuncRoot{ Func: reflect.ValueOf(strings.IndexFunc) }
	Roots[`/strings/IndexRune`] = FuncRoot{ Func: reflect.ValueOf(strings.IndexRune) }
	Roots[`/strings/Join`] = FuncRoot{ Func: reflect.ValueOf(strings.Join) }
	Roots[`/strings/LastIndex`] = FuncRoot{ Func: reflect.ValueOf(strings.LastIndex) }
	Roots[`/strings/LastIndexAny`] = FuncRoot{ Func: reflect.ValueOf(strings.LastIndexAny) }
	Roots[`/strings/LastIndexFunc`] = FuncRoot{ Func: reflect.ValueOf(strings.LastIndexFunc) }
	Roots[`/strings/Map`] = FuncRoot{ Func: reflect.ValueOf(strings.Map) }
	Roots[`/strings/NewReader`] = FuncRoot{ Func: reflect.ValueOf(strings.NewReader) }
	Roots[`/strings/NewReplacer`] = FuncRoot{ Func: reflect.ValueOf(strings.NewReplacer) }
	Roots[`/strings/Repeat`] = FuncRoot{ Func: reflect.ValueOf(strings.Repeat) }
	Roots[`/strings/Replace`] = FuncRoot{ Func: reflect.ValueOf(strings.Replace) }
	Roots[`/strings/Split`] = FuncRoot{ Func: reflect.ValueOf(strings.Split) }
	Roots[`/strings/SplitAfter`] = FuncRoot{ Func: reflect.ValueOf(strings.SplitAfter) }
	Roots[`/strings/SplitAfterN`] = FuncRoot{ Func: reflect.ValueOf(strings.SplitAfterN) }
	Roots[`/strings/SplitN`] = FuncRoot{ Func: reflect.ValueOf(strings.SplitN) }
	Roots[`/strings/Title`] = FuncRoot{ Func: reflect.ValueOf(strings.Title) }
	Roots[`/strings/ToLower`] = FuncRoot{ Func: reflect.ValueOf(strings.ToLower) }
	Roots[`/strings/ToLowerSpecial`] = FuncRoot{ Func: reflect.ValueOf(strings.ToLowerSpecial) }
	Roots[`/strings/ToTitle`] = FuncRoot{ Func: reflect.ValueOf(strings.ToTitle) }
	Roots[`/strings/ToTitleSpecial`] = FuncRoot{ Func: reflect.ValueOf(strings.ToTitleSpecial) }
	Roots[`/strings/ToUpper`] = FuncRoot{ Func: reflect.ValueOf(strings.ToUpper) }
	Roots[`/strings/ToUpperSpecial`] = FuncRoot{ Func: reflect.ValueOf(strings.ToUpperSpecial) }
	Roots[`/strings/Trim`] = FuncRoot{ Func: reflect.ValueOf(strings.Trim) }
	Roots[`/strings/TrimFunc`] = FuncRoot{ Func: reflect.ValueOf(strings.TrimFunc) }
	Roots[`/strings/TrimLeft`] = FuncRoot{ Func: reflect.ValueOf(strings.TrimLeft) }
	Roots[`/strings/TrimLeftFunc`] = FuncRoot{ Func: reflect.ValueOf(strings.TrimLeftFunc) }
	Roots[`/strings/TrimRight`] = FuncRoot{ Func: reflect.ValueOf(strings.TrimRight) }
	Roots[`/strings/TrimRightFunc`] = FuncRoot{ Func: reflect.ValueOf(strings.TrimRightFunc) }
	Roots[`/strings/TrimSpace`] = FuncRoot{ Func: reflect.ValueOf(strings.TrimSpace) }
	Roots[`/bytes/ErrTooLarge`] = VarRoot{ Var: reflect.ValueOf(&bytes.ErrTooLarge) }
	Roots[`/strconv/ErrRange`] = VarRoot{ Var: reflect.ValueOf(&strconv.ErrRange) }
	Roots[`/strconv/ErrSyntax`] = VarRoot{ Var: reflect.ValueOf(&strconv.ErrSyntax) }
	{
	var tmp *bytes.Buffer
	Roots[`/bytes/Buffer`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *bytes.Reader
	Roots[`/bytes/Reader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.Method
	Roots[`/reflect/Method`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.Method
	Roots[`/reflect/Method`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.Method
	Roots[`/reflect/Method`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.Method
	Roots[`/reflect/Method`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.Method
	Roots[`/reflect/Method`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.Method
	Roots[`/reflect/Method`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.SliceHeader
	Roots[`/reflect/SliceHeader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.SliceHeader
	Roots[`/reflect/SliceHeader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.SliceHeader
	Roots[`/reflect/SliceHeader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.SliceHeader
	Roots[`/reflect/SliceHeader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StringHeader
	Roots[`/reflect/StringHeader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StringHeader
	Roots[`/reflect/StringHeader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StringHeader
	Roots[`/reflect/StringHeader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.Value
	Roots[`/reflect/Value`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.ValueError
	Roots[`/reflect/ValueError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.ValueError
	Roots[`/reflect/ValueError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.ValueError
	Roots[`/reflect/ValueError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *strconv.NumError
	Roots[`/strconv/NumError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *strconv.NumError
	Roots[`/strconv/NumError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *strconv.NumError
	Roots[`/strconv/NumError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *strconv.NumError
	Roots[`/strconv/NumError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *strings.Reader
	Roots[`/strings/Reader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *strings.Replacer
	Roots[`/strings/Replacer`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	Roots[`/bytes/MinRead`] = ConstRoot{ Const: int64(bytes.MinRead) }
	Roots[`/reflect/Array`] = ConstRoot{ Const: reflect.Array }
	Roots[`/reflect/Bool`] = ConstRoot{ Const: reflect.Bool }
	Roots[`/reflect/BothDir`] = ConstRoot{ Const: reflect.BothDir }
	Roots[`/reflect/Chan`] = ConstRoot{ Const: reflect.Chan }
	Roots[`/reflect/Complex128`] = ConstRoot{ Const: reflect.Complex128 }
	Roots[`/reflect/Complex64`] = ConstRoot{ Const: reflect.Complex64 }
	Roots[`/reflect/Float32`] = ConstRoot{ Const: reflect.Float32 }
	Roots[`/reflect/Float64`] = ConstRoot{ Const: reflect.Float64 }
	Roots[`/reflect/Func`] = ConstRoot{ Const: reflect.Func }
	Roots[`/reflect/Int`] = ConstRoot{ Const: reflect.Int }
	Roots[`/reflect/Int16`] = ConstRoot{ Const: reflect.Int16 }
	Roots[`/reflect/Int32`] = ConstRoot{ Const: reflect.Int32 }
	Roots[`/reflect/Int64`] = ConstRoot{ Const: reflect.Int64 }
	Roots[`/reflect/Int8`] = ConstRoot{ Const: reflect.Int8 }
	Roots[`/reflect/Interface`] = ConstRoot{ Const: reflect.Interface }
	Roots[`/reflect/Invalid`] = ConstRoot{ Const: reflect.Invalid }
	Roots[`/reflect/Map`] = ConstRoot{ Const: reflect.Map }
	Roots[`/reflect/Ptr`] = ConstRoot{ Const: reflect.Ptr }
	Roots[`/reflect/RecvDir`] = ConstRoot{ Const: reflect.RecvDir }
	Roots[`/reflect/SendDir`] = ConstRoot{ Const: reflect.SendDir }
	Roots[`/reflect/Slice`] = ConstRoot{ Const: reflect.Slice }
	Roots[`/reflect/String`] = ConstRoot{ Const: reflect.String }
	Roots[`/reflect/Struct`] = ConstRoot{ Const: reflect.Struct }
	Roots[`/reflect/Uint`] = ConstRoot{ Const: reflect.Uint }
	Roots[`/reflect/Uint16`] = ConstRoot{ Const: reflect.Uint16 }
	Roots[`/reflect/Uint32`] = ConstRoot{ Const: reflect.Uint32 }
	Roots[`/reflect/Uint64`] = ConstRoot{ Const: reflect.Uint64 }
	Roots[`/reflect/Uint8`] = ConstRoot{ Const: reflect.Uint8 }
	Roots[`/reflect/Uintptr`] = ConstRoot{ Const: reflect.Uintptr }
	Roots[`/reflect/UnsafePointer`] = ConstRoot{ Const: reflect.UnsafePointer }
	Roots[`/strconv/IntSize`] = ConstRoot{ Const: int64(strconv.IntSize) }
}
