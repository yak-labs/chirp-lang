all: ci goterp test


goterp: src/generated/reflections.go src/*/*.go
	GOPATH=$$(pwd) time go build -x goterp 
	: : : : : COMPILED OK
test:
	GOPATH=$$(pwd) time go test -i src/terp/*test.go
	: : : DEPENDENT PACKAGES INSTALLED.
	GOPATH=$$(pwd) time go test src/terp/*test.go
	: : : : : TESTED OK

src/generated/reflections.go : src/mkreflections.go
	mkdir -p src/generated
	GOPATH=$$PWD go  run  src/mkreflections.go > src/generated/reflections.go $$(cd $$(go env GOROOT)/pkg/$$(go env GOHOSTOS)_$$(go env GOHOSTARCH)/ ; find * -name \*.a ! -name cgo.a | sed 's/[.]a$$//' | LANG=C sort)
ci:
	ci -l -q -m/dev/null -t/dev/null src/*/*.go 
fmt:
	gofmt -w src/*/*.go 

demo1: goterp
	./goterp -c 'must 3 [+ 1 2]; must 1015 [+ 0 1 2 3 [ + 2 2 ] [ - 8 3 ] [+ 100 900]]'
demo2: goterp
	time ./goterp -c ' call /fmt/Printf "hey%gthere%gyou---" [+ 3] [+ 4] '
web: goterp
	./goterp demo/hello_web.t

clean:
	rm -rf ./goterp ./demo-types-1 src/generated/reflections.go pkg/*

FORCE:
