package goapi

import . `github.com/yak-labs/chirp-lang`
import (
	bytes `bytes`
	fmt `fmt`
	math_big `math/big`
	net_http `net/http`
	os `os`
	reflect `reflect`
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
	Roots[`/math/big/NewInt`] = FuncRoot{ Func: reflect.ValueOf(math_big.NewInt) }
	Roots[`/math/big/NewRat`] = FuncRoot{ Func: reflect.ValueOf(math_big.NewRat) }
	Roots[`/net/http/CanonicalHeaderKey`] = FuncRoot{ Func: reflect.ValueOf(net_http.CanonicalHeaderKey) }
	Roots[`/net/http/DetectContentType`] = FuncRoot{ Func: reflect.ValueOf(net_http.DetectContentType) }
	Roots[`/net/http/Error`] = FuncRoot{ Func: reflect.ValueOf(net_http.Error) }
	Roots[`/net/http/FileServer`] = FuncRoot{ Func: reflect.ValueOf(net_http.FileServer) }
	Roots[`/net/http/Get`] = FuncRoot{ Func: reflect.ValueOf(net_http.Get) }
	Roots[`/net/http/Handle`] = FuncRoot{ Func: reflect.ValueOf(net_http.Handle) }
	Roots[`/net/http/HandleFunc`] = FuncRoot{ Func: reflect.ValueOf(net_http.HandleFunc) }
	Roots[`/net/http/Head`] = FuncRoot{ Func: reflect.ValueOf(net_http.Head) }
	Roots[`/net/http/ListenAndServe`] = FuncRoot{ Func: reflect.ValueOf(net_http.ListenAndServe) }
	Roots[`/net/http/ListenAndServeTLS`] = FuncRoot{ Func: reflect.ValueOf(net_http.ListenAndServeTLS) }
	Roots[`/net/http/MaxBytesReader`] = FuncRoot{ Func: reflect.ValueOf(net_http.MaxBytesReader) }
	Roots[`/net/http/NewFileTransport`] = FuncRoot{ Func: reflect.ValueOf(net_http.NewFileTransport) }
	Roots[`/net/http/NewRequest`] = FuncRoot{ Func: reflect.ValueOf(net_http.NewRequest) }
	Roots[`/net/http/NewServeMux`] = FuncRoot{ Func: reflect.ValueOf(net_http.NewServeMux) }
	Roots[`/net/http/NotFound`] = FuncRoot{ Func: reflect.ValueOf(net_http.NotFound) }
	Roots[`/net/http/NotFoundHandler`] = FuncRoot{ Func: reflect.ValueOf(net_http.NotFoundHandler) }
	Roots[`/net/http/ParseHTTPVersion`] = FuncRoot{ Func: reflect.ValueOf(net_http.ParseHTTPVersion) }
	Roots[`/net/http/Post`] = FuncRoot{ Func: reflect.ValueOf(net_http.Post) }
	Roots[`/net/http/PostForm`] = FuncRoot{ Func: reflect.ValueOf(net_http.PostForm) }
	Roots[`/net/http/ProxyFromEnvironment`] = FuncRoot{ Func: reflect.ValueOf(net_http.ProxyFromEnvironment) }
	Roots[`/net/http/ProxyURL`] = FuncRoot{ Func: reflect.ValueOf(net_http.ProxyURL) }
	Roots[`/net/http/ReadRequest`] = FuncRoot{ Func: reflect.ValueOf(net_http.ReadRequest) }
	Roots[`/net/http/ReadResponse`] = FuncRoot{ Func: reflect.ValueOf(net_http.ReadResponse) }
	Roots[`/net/http/Redirect`] = FuncRoot{ Func: reflect.ValueOf(net_http.Redirect) }
	Roots[`/net/http/RedirectHandler`] = FuncRoot{ Func: reflect.ValueOf(net_http.RedirectHandler) }
	Roots[`/net/http/Serve`] = FuncRoot{ Func: reflect.ValueOf(net_http.Serve) }
	Roots[`/net/http/ServeContent`] = FuncRoot{ Func: reflect.ValueOf(net_http.ServeContent) }
	Roots[`/net/http/ServeFile`] = FuncRoot{ Func: reflect.ValueOf(net_http.ServeFile) }
	Roots[`/net/http/SetCookie`] = FuncRoot{ Func: reflect.ValueOf(net_http.SetCookie) }
	Roots[`/net/http/StatusText`] = FuncRoot{ Func: reflect.ValueOf(net_http.StatusText) }
	Roots[`/net/http/StripPrefix`] = FuncRoot{ Func: reflect.ValueOf(net_http.StripPrefix) }
	Roots[`/net/http/TimeoutHandler`] = FuncRoot{ Func: reflect.ValueOf(net_http.TimeoutHandler) }
	Roots[`/os/Chdir`] = FuncRoot{ Func: reflect.ValueOf(os.Chdir) }
	Roots[`/os/Chmod`] = FuncRoot{ Func: reflect.ValueOf(os.Chmod) }
	Roots[`/os/Chown`] = FuncRoot{ Func: reflect.ValueOf(os.Chown) }
	Roots[`/os/Chtimes`] = FuncRoot{ Func: reflect.ValueOf(os.Chtimes) }
	Roots[`/os/Clearenv`] = FuncRoot{ Func: reflect.ValueOf(os.Clearenv) }
	Roots[`/os/Create`] = FuncRoot{ Func: reflect.ValueOf(os.Create) }
	Roots[`/os/Environ`] = FuncRoot{ Func: reflect.ValueOf(os.Environ) }
	Roots[`/os/Exit`] = FuncRoot{ Func: reflect.ValueOf(os.Exit) }
	Roots[`/os/Expand`] = FuncRoot{ Func: reflect.ValueOf(os.Expand) }
	Roots[`/os/ExpandEnv`] = FuncRoot{ Func: reflect.ValueOf(os.ExpandEnv) }
	Roots[`/os/FindProcess`] = FuncRoot{ Func: reflect.ValueOf(os.FindProcess) }
	Roots[`/os/Getegid`] = FuncRoot{ Func: reflect.ValueOf(os.Getegid) }
	Roots[`/os/Getenv`] = FuncRoot{ Func: reflect.ValueOf(os.Getenv) }
	Roots[`/os/Geteuid`] = FuncRoot{ Func: reflect.ValueOf(os.Geteuid) }
	Roots[`/os/Getgid`] = FuncRoot{ Func: reflect.ValueOf(os.Getgid) }
	Roots[`/os/Getgroups`] = FuncRoot{ Func: reflect.ValueOf(os.Getgroups) }
	Roots[`/os/Getpagesize`] = FuncRoot{ Func: reflect.ValueOf(os.Getpagesize) }
	Roots[`/os/Getpid`] = FuncRoot{ Func: reflect.ValueOf(os.Getpid) }
	Roots[`/os/Getppid`] = FuncRoot{ Func: reflect.ValueOf(os.Getppid) }
	Roots[`/os/Getuid`] = FuncRoot{ Func: reflect.ValueOf(os.Getuid) }
	Roots[`/os/Getwd`] = FuncRoot{ Func: reflect.ValueOf(os.Getwd) }
	Roots[`/os/Hostname`] = FuncRoot{ Func: reflect.ValueOf(os.Hostname) }
	Roots[`/os/IsExist`] = FuncRoot{ Func: reflect.ValueOf(os.IsExist) }
	Roots[`/os/IsNotExist`] = FuncRoot{ Func: reflect.ValueOf(os.IsNotExist) }
	Roots[`/os/IsPathSeparator`] = FuncRoot{ Func: reflect.ValueOf(os.IsPathSeparator) }
	Roots[`/os/IsPermission`] = FuncRoot{ Func: reflect.ValueOf(os.IsPermission) }
	Roots[`/os/Lchown`] = FuncRoot{ Func: reflect.ValueOf(os.Lchown) }
	Roots[`/os/Link`] = FuncRoot{ Func: reflect.ValueOf(os.Link) }
	Roots[`/os/Lstat`] = FuncRoot{ Func: reflect.ValueOf(os.Lstat) }
	Roots[`/os/Mkdir`] = FuncRoot{ Func: reflect.ValueOf(os.Mkdir) }
	Roots[`/os/MkdirAll`] = FuncRoot{ Func: reflect.ValueOf(os.MkdirAll) }
	Roots[`/os/NewFile`] = FuncRoot{ Func: reflect.ValueOf(os.NewFile) }
	Roots[`/os/NewSyscallError`] = FuncRoot{ Func: reflect.ValueOf(os.NewSyscallError) }
	Roots[`/os/Open`] = FuncRoot{ Func: reflect.ValueOf(os.Open) }
	Roots[`/os/OpenFile`] = FuncRoot{ Func: reflect.ValueOf(os.OpenFile) }
	Roots[`/os/Pipe`] = FuncRoot{ Func: reflect.ValueOf(os.Pipe) }
	Roots[`/os/Readlink`] = FuncRoot{ Func: reflect.ValueOf(os.Readlink) }
	Roots[`/os/Remove`] = FuncRoot{ Func: reflect.ValueOf(os.Remove) }
	Roots[`/os/RemoveAll`] = FuncRoot{ Func: reflect.ValueOf(os.RemoveAll) }
	Roots[`/os/Rename`] = FuncRoot{ Func: reflect.ValueOf(os.Rename) }
	Roots[`/os/SameFile`] = FuncRoot{ Func: reflect.ValueOf(os.SameFile) }
	Roots[`/os/Setenv`] = FuncRoot{ Func: reflect.ValueOf(os.Setenv) }
	Roots[`/os/StartProcess`] = FuncRoot{ Func: reflect.ValueOf(os.StartProcess) }
	Roots[`/os/Stat`] = FuncRoot{ Func: reflect.ValueOf(os.Stat) }
	Roots[`/os/Symlink`] = FuncRoot{ Func: reflect.ValueOf(os.Symlink) }
	Roots[`/os/TempDir`] = FuncRoot{ Func: reflect.ValueOf(os.TempDir) }
	Roots[`/os/Truncate`] = FuncRoot{ Func: reflect.ValueOf(os.Truncate) }
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
	Roots[`/bytes/ErrTooLarge`] = VarRoot{ Var: reflect.ValueOf(&bytes.ErrTooLarge) }
	Roots[`/net/http/DefaultClient`] = VarRoot{ Var: reflect.ValueOf(&net_http.DefaultClient) }
	Roots[`/net/http/DefaultServeMux`] = VarRoot{ Var: reflect.ValueOf(&net_http.DefaultServeMux) }
	Roots[`/net/http/DefaultTransport`] = VarRoot{ Var: reflect.ValueOf(&net_http.DefaultTransport) }
	Roots[`/net/http/ErrBodyNotAllowed`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrBodyNotAllowed) }
	Roots[`/net/http/ErrBodyReadAfterClose`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrBodyReadAfterClose) }
	Roots[`/net/http/ErrContentLength`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrContentLength) }
	Roots[`/net/http/ErrHandlerTimeout`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrHandlerTimeout) }
	Roots[`/net/http/ErrHeaderTooLong`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrHeaderTooLong) }
	Roots[`/net/http/ErrHijacked`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrHijacked) }
	Roots[`/net/http/ErrLineTooLong`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrLineTooLong) }
	Roots[`/net/http/ErrMissingBoundary`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrMissingBoundary) }
	Roots[`/net/http/ErrMissingContentLength`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrMissingContentLength) }
	Roots[`/net/http/ErrMissingFile`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrMissingFile) }
	Roots[`/net/http/ErrNoCookie`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrNoCookie) }
	Roots[`/net/http/ErrNoLocation`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrNoLocation) }
	Roots[`/net/http/ErrNotMultipart`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrNotMultipart) }
	Roots[`/net/http/ErrNotSupported`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrNotSupported) }
	Roots[`/net/http/ErrShortBody`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrShortBody) }
	Roots[`/net/http/ErrUnexpectedTrailer`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrUnexpectedTrailer) }
	Roots[`/net/http/ErrWriteAfterFlush`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrWriteAfterFlush) }
	Roots[`/os/Args`] = VarRoot{ Var: reflect.ValueOf(&os.Args) }
	Roots[`/os/ErrExist`] = VarRoot{ Var: reflect.ValueOf(&os.ErrExist) }
	Roots[`/os/ErrInvalid`] = VarRoot{ Var: reflect.ValueOf(&os.ErrInvalid) }
	Roots[`/os/ErrNotExist`] = VarRoot{ Var: reflect.ValueOf(&os.ErrNotExist) }
	Roots[`/os/ErrPermission`] = VarRoot{ Var: reflect.ValueOf(&os.ErrPermission) }
	Roots[`/os/Interrupt`] = VarRoot{ Var: reflect.ValueOf(&os.Interrupt) }
	Roots[`/os/Kill`] = VarRoot{ Var: reflect.ValueOf(&os.Kill) }
	Roots[`/os/Stderr`] = VarRoot{ Var: reflect.ValueOf(&os.Stderr) }
	Roots[`/os/Stdin`] = VarRoot{ Var: reflect.ValueOf(&os.Stdin) }
	Roots[`/os/Stdout`] = VarRoot{ Var: reflect.ValueOf(&os.Stdout) }
	{
	var tmp *bytes.Buffer
	Roots[`/bytes/Buffer`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *bytes.Reader
	Roots[`/bytes/Reader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *math_big.Int
	Roots[`/math/big/Int`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *math_big.Rat
	Roots[`/math/big/Rat`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Client
	Roots[`/net/http/Client`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Client
	Roots[`/net/http/Client`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Client
	Roots[`/net/http/Client`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Client
	Roots[`/net/http/Client`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.ProtocolError
	Roots[`/net/http/ProtocolError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.ProtocolError
	Roots[`/net/http/ProtocolError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.ServeMux
	Roots[`/net/http/ServeMux`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Server
	Roots[`/net/http/Server`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Server
	Roots[`/net/http/Server`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Server
	Roots[`/net/http/Server`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Server
	Roots[`/net/http/Server`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Server
	Roots[`/net/http/Server`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Server
	Roots[`/net/http/Server`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Server
	Roots[`/net/http/Server`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Transport
	Roots[`/net/http/Transport`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Transport
	Roots[`/net/http/Transport`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Transport
	Roots[`/net/http/Transport`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Transport
	Roots[`/net/http/Transport`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Transport
	Roots[`/net/http/Transport`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Transport
	Roots[`/net/http/Transport`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Transport
	Roots[`/net/http/Transport`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.File
	Roots[`/os/File`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.LinkError
	Roots[`/os/LinkError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.LinkError
	Roots[`/os/LinkError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.LinkError
	Roots[`/os/LinkError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.LinkError
	Roots[`/os/LinkError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.LinkError
	Roots[`/os/LinkError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.PathError
	Roots[`/os/PathError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.PathError
	Roots[`/os/PathError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.PathError
	Roots[`/os/PathError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.PathError
	Roots[`/os/PathError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.ProcAttr
	Roots[`/os/ProcAttr`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.ProcAttr
	Roots[`/os/ProcAttr`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.ProcAttr
	Roots[`/os/ProcAttr`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.ProcAttr
	Roots[`/os/ProcAttr`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.ProcAttr
	Roots[`/os/ProcAttr`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.Process
	Roots[`/os/Process`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.Process
	Roots[`/os/Process`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.ProcessState
	Roots[`/os/ProcessState`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.SyscallError
	Roots[`/os/SyscallError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.SyscallError
	Roots[`/os/SyscallError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.SyscallError
	Roots[`/os/SyscallError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
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
	Roots[`/bytes/MinRead`] = ConstRoot{ Const: int64(bytes.MinRead) }
	Roots[`/math/big/MaxBase`] = ConstRoot{ Const: int64(math_big.MaxBase) }
	Roots[`/net/http/DefaultMaxHeaderBytes`] = ConstRoot{ Const: int64(net_http.DefaultMaxHeaderBytes) }
	Roots[`/net/http/DefaultMaxIdleConnsPerHost`] = ConstRoot{ Const: int64(net_http.DefaultMaxIdleConnsPerHost) }
	Roots[`/net/http/StatusAccepted`] = ConstRoot{ Const: int64(net_http.StatusAccepted) }
	Roots[`/net/http/StatusBadGateway`] = ConstRoot{ Const: int64(net_http.StatusBadGateway) }
	Roots[`/net/http/StatusBadRequest`] = ConstRoot{ Const: int64(net_http.StatusBadRequest) }
	Roots[`/net/http/StatusConflict`] = ConstRoot{ Const: int64(net_http.StatusConflict) }
	Roots[`/net/http/StatusContinue`] = ConstRoot{ Const: int64(net_http.StatusContinue) }
	Roots[`/net/http/StatusCreated`] = ConstRoot{ Const: int64(net_http.StatusCreated) }
	Roots[`/net/http/StatusExpectationFailed`] = ConstRoot{ Const: int64(net_http.StatusExpectationFailed) }
	Roots[`/net/http/StatusForbidden`] = ConstRoot{ Const: int64(net_http.StatusForbidden) }
	Roots[`/net/http/StatusFound`] = ConstRoot{ Const: int64(net_http.StatusFound) }
	Roots[`/net/http/StatusGatewayTimeout`] = ConstRoot{ Const: int64(net_http.StatusGatewayTimeout) }
	Roots[`/net/http/StatusGone`] = ConstRoot{ Const: int64(net_http.StatusGone) }
	Roots[`/net/http/StatusHTTPVersionNotSupported`] = ConstRoot{ Const: int64(net_http.StatusHTTPVersionNotSupported) }
	Roots[`/net/http/StatusInternalServerError`] = ConstRoot{ Const: int64(net_http.StatusInternalServerError) }
	Roots[`/net/http/StatusLengthRequired`] = ConstRoot{ Const: int64(net_http.StatusLengthRequired) }
	Roots[`/net/http/StatusMethodNotAllowed`] = ConstRoot{ Const: int64(net_http.StatusMethodNotAllowed) }
	Roots[`/net/http/StatusMovedPermanently`] = ConstRoot{ Const: int64(net_http.StatusMovedPermanently) }
	Roots[`/net/http/StatusMultipleChoices`] = ConstRoot{ Const: int64(net_http.StatusMultipleChoices) }
	Roots[`/net/http/StatusNoContent`] = ConstRoot{ Const: int64(net_http.StatusNoContent) }
	Roots[`/net/http/StatusNonAuthoritativeInfo`] = ConstRoot{ Const: int64(net_http.StatusNonAuthoritativeInfo) }
	Roots[`/net/http/StatusNotAcceptable`] = ConstRoot{ Const: int64(net_http.StatusNotAcceptable) }
	Roots[`/net/http/StatusNotFound`] = ConstRoot{ Const: int64(net_http.StatusNotFound) }
	Roots[`/net/http/StatusNotImplemented`] = ConstRoot{ Const: int64(net_http.StatusNotImplemented) }
	Roots[`/net/http/StatusNotModified`] = ConstRoot{ Const: int64(net_http.StatusNotModified) }
	Roots[`/net/http/StatusOK`] = ConstRoot{ Const: int64(net_http.StatusOK) }
	Roots[`/net/http/StatusPartialContent`] = ConstRoot{ Const: int64(net_http.StatusPartialContent) }
	Roots[`/net/http/StatusPaymentRequired`] = ConstRoot{ Const: int64(net_http.StatusPaymentRequired) }
	Roots[`/net/http/StatusPreconditionFailed`] = ConstRoot{ Const: int64(net_http.StatusPreconditionFailed) }
	Roots[`/net/http/StatusProxyAuthRequired`] = ConstRoot{ Const: int64(net_http.StatusProxyAuthRequired) }
	Roots[`/net/http/StatusRequestEntityTooLarge`] = ConstRoot{ Const: int64(net_http.StatusRequestEntityTooLarge) }
	Roots[`/net/http/StatusRequestTimeout`] = ConstRoot{ Const: int64(net_http.StatusRequestTimeout) }
	Roots[`/net/http/StatusRequestURITooLong`] = ConstRoot{ Const: int64(net_http.StatusRequestURITooLong) }
	Roots[`/net/http/StatusRequestedRangeNotSatisfiable`] = ConstRoot{ Const: int64(net_http.StatusRequestedRangeNotSatisfiable) }
	Roots[`/net/http/StatusResetContent`] = ConstRoot{ Const: int64(net_http.StatusResetContent) }
	Roots[`/net/http/StatusSeeOther`] = ConstRoot{ Const: int64(net_http.StatusSeeOther) }
	Roots[`/net/http/StatusServiceUnavailable`] = ConstRoot{ Const: int64(net_http.StatusServiceUnavailable) }
	Roots[`/net/http/StatusSwitchingProtocols`] = ConstRoot{ Const: int64(net_http.StatusSwitchingProtocols) }
	Roots[`/net/http/StatusTeapot`] = ConstRoot{ Const: int64(net_http.StatusTeapot) }
	Roots[`/net/http/StatusTemporaryRedirect`] = ConstRoot{ Const: int64(net_http.StatusTemporaryRedirect) }
	Roots[`/net/http/StatusUnauthorized`] = ConstRoot{ Const: int64(net_http.StatusUnauthorized) }
	Roots[`/net/http/StatusUnsupportedMediaType`] = ConstRoot{ Const: int64(net_http.StatusUnsupportedMediaType) }
	Roots[`/net/http/StatusUseProxy`] = ConstRoot{ Const: int64(net_http.StatusUseProxy) }
	Roots[`/net/http/TimeFormat`] = ConstRoot{ Const: net_http.TimeFormat }
	Roots[`/os/DevNull`] = ConstRoot{ Const: os.DevNull }
	Roots[`/os/ModeAppend`] = ConstRoot{ Const: os.ModeAppend }
	Roots[`/os/ModeCharDevice`] = ConstRoot{ Const: os.ModeCharDevice }
	Roots[`/os/ModeDevice`] = ConstRoot{ Const: os.ModeDevice }
	Roots[`/os/ModeDir`] = ConstRoot{ Const: os.ModeDir }
	Roots[`/os/ModeExclusive`] = ConstRoot{ Const: os.ModeExclusive }
	Roots[`/os/ModeNamedPipe`] = ConstRoot{ Const: os.ModeNamedPipe }
	Roots[`/os/ModePerm`] = ConstRoot{ Const: os.ModePerm }
	Roots[`/os/ModeSetgid`] = ConstRoot{ Const: os.ModeSetgid }
	Roots[`/os/ModeSetuid`] = ConstRoot{ Const: os.ModeSetuid }
	Roots[`/os/ModeSocket`] = ConstRoot{ Const: os.ModeSocket }
	Roots[`/os/ModeSticky`] = ConstRoot{ Const: os.ModeSticky }
	Roots[`/os/ModeSymlink`] = ConstRoot{ Const: os.ModeSymlink }
	Roots[`/os/ModeTemporary`] = ConstRoot{ Const: os.ModeTemporary }
	Roots[`/os/ModeType`] = ConstRoot{ Const: os.ModeType }
	Roots[`/os/O_APPEND`] = ConstRoot{ Const: os.O_APPEND }
	Roots[`/os/O_CREATE`] = ConstRoot{ Const: os.O_CREATE }
	Roots[`/os/O_EXCL`] = ConstRoot{ Const: os.O_EXCL }
	Roots[`/os/O_RDONLY`] = ConstRoot{ Const: os.O_RDONLY }
	Roots[`/os/O_RDWR`] = ConstRoot{ Const: os.O_RDWR }
	Roots[`/os/O_SYNC`] = ConstRoot{ Const: os.O_SYNC }
	Roots[`/os/O_TRUNC`] = ConstRoot{ Const: os.O_TRUNC }
	Roots[`/os/O_WRONLY`] = ConstRoot{ Const: os.O_WRONLY }
	Roots[`/os/PathListSeparator`] = ConstRoot{ Const: os.PathListSeparator }
	Roots[`/os/PathSeparator`] = ConstRoot{ Const: os.PathSeparator }
	Roots[`/os/SEEK_CUR`] = ConstRoot{ Const: os.SEEK_CUR }
	Roots[`/os/SEEK_END`] = ConstRoot{ Const: os.SEEK_END }
	Roots[`/os/SEEK_SET`] = ConstRoot{ Const: os.SEEK_SET }
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
}
