all: goterp ci demo1

goterp: FORCE
	GOPATH=$$(pwd) go build -x goterp 
ci:
	ci -l -q -m/dev/null -t/dev/null src/*/*.go 
fmt:
	gofmt -w src/*/*.go 

demo1:
	./goterp 'must 3 [+ 1 2]; must 1015 [+ 0 1 2 3 [ + 2 2 ] 5 [+ 100 900]]'

FORCE:
