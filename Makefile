all: goterp ci

goterp: FORCE
	GOPATH=$$(pwd) go build -x goterp 
ci:
	ci -l -q -m/dev/null -t/dev/null src/*/*.go 

FORCE:
