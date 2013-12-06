wrap.go : ../goapi_wrap.tcl Makefile
	tclsh ../goapi_wrap.tcl < "`go env GOROOT`/api/go1.txt" > wrap.go reflect bytes strings strconf fmt bufio os net/http regexp  encoding/base64 io/ioutil math math/big
	go build
clean:
	-rm wrap.go
