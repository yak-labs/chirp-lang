wrap.go : ../goapi_wrap.tcl Makefile
	tclsh ../goapi_wrap.tcl < ../api/go1.txt > wrap.go
	go build
clean:
	-rm wrap.go
