all: ci goterp test

goterp: FORCE
	GOPATH=$$(pwd) go build -x goterp 
test: FORCE
	GOPATH=$$(pwd) go test -i src/terp/*.go
	GOPATH=$$(pwd) go test src/terp/*.go
ci:
	ci -l -q -m/dev/null -t/dev/null src/*/*.go 
fmt:
	gofmt -w src/*/*.go 

demo1:
	./goterp 'must 3 [+ 1 2]; must 1015 [+ 0 1 2 3 [ + 2 2 ] [ - 8 3 ] [+ 100 900]]'

FORCE:
