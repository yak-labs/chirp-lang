all: ci goterp test


goterp: src/generated/reflections.go FORCE
	GOPATH=$$(pwd) go build -x goterp 
test: FORCE
	GOPATH=$$(pwd) go test -i src/terp/*.go
	GOPATH=$$(pwd) go test src/terp/*.go

_externs.txt: src/externs/externs.go
	go run src/externs/externs.go $$( eval $$(go env) ; find $$GOROOT/src/pkg/* -type d ) > _externs.txt

src/generated/reflections.go : src/externs/generate.tcl _externs.txt
	tclsh src/externs/generate1.tcl < _externs.txt  | grep -v // | grep -v _test | grep -v Test > src/generated/reflections.go
	tclsh src/externs/generate2.tcl < _externs.txt  | grep -v // | grep -v _test | grep -v Test >> src/generated/reflections.go
ci:
	ci -l -q -m/dev/null -t/dev/null src/*/*.go 
fmt:
	gofmt -w src/*/*.go 

demo1:
	./goterp 'must 3 [+ 1 2]; must 1015 [+ 0 1 2 3 [ + 2 2 ] [ - 8 3 ] [+ 100 900]]'

clean:
	rm -rf ./goterp _externs.txt src/generated/reflections.go pkg/*

FORCE:
