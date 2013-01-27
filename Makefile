all: ci goterp test

big: clean all demos

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
	ci -l -q -m/dev/null -t/dev/null src/*.go src/*/*.go *.t *.txt demo/*.t
fmt:
	gofmt -w src/*/*.go 

run1: all
	rm -rf ./data
	./goterp unbundle.t < run1.txt
	./goterp smilax5.t

demos: goterp
	sh demo/all.sh
demoweb: goterp
	./goterp demo/hello_web.t
demowiki: goterp
	( cd demo ; ../goterp wiki.t )

clean:
	rm -rf ./goterp src/generated/reflections.go pkg/* demo/_*.err ./data

FORCE:
