wrap.go : ../goapi_wrap.tcl Makefile
	tclsh ../goapi_wrap.tcl < "`go env GOROOT`/api/go1.txt" > wrap.go
	go build
clean:
	-rm wrap.go
