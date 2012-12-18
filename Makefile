all: ci goterp demo2 test


goterp: src/generated/reflections.go src/*/*.go
	GOPATH=$$(pwd) time go build -x goterp 
test:
	GOPATH=$$(pwd) time go test -i src/terp/*.go

src/generated/reflections.go : src/mkreflections.go
	mkdir -p src/generated
	GOPATH=$$PWD go  run  src/mkreflections.go > src/generated/reflections.go $$(cd $$(go env GOROOT)/pkg/$$(go env GOHOSTOS)_$$(go env GOHOSTARCH)/ ; find * -name \*.a ! -name cgo.a | sed 's/[.]a$$//' | LANG=C sort)
ci:
	ci -l -q -m/dev/null -t/dev/null src/*/*.go 
fmt:
	gofmt -w src/*/*.go 

demo1: goterp
	./goterp 'must 3 [+ 1 2]; must 1015 [+ 0 1 2 3 [ + 2 2 ] [ - 8 3 ] [+ 100 900]]'
demo2: goterp
	time ./goterp ' call /fmt/Printf "hey%gthere%gyou---" [+ 3] [+ 4] '
web: goterp
	./goterp "`cat demo/hello_web.t`" 

clean:
	rm -rf ./goterp ./demo-types-1 src/generated/reflections.go pkg/*

FORCE:
