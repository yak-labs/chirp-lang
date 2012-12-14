all: ci goterp demo2 test


goterp: src/generated/reflections.go FORCE
	GOPATH=$$(pwd) go build -x goterp 
test: FORCE
	GOPATH=$$(pwd) go test -i src/terp/*.go
	GOPATH=$$(pwd) go test src/terp/*.go

src/generated/reflections.go : src/mkreflections.go
	mkdir -p src/generated
	GOPATH=$$PWD go  run  src/mkreflections.go > src/generated/reflections.go
ci:
	ci -l -q -m/dev/null -t/dev/null src/*/*.go 
fmt:
	gofmt -w src/*/*.go 

demo1:
	./goterp 'must 3 [+ 1 2]; must 1015 [+ 0 1 2 3 [ + 2 2 ] [ - 8 3 ] [+ 100 900]]'
demo2:
	./goterp ' call /fmt/Printf "hey%gthere%gyou---" [+ 3] [+ 4] '

clean:
	rm -rf ./goterp ./demo-types-1 src/generated/reflections.go pkg/*

FORCE:
