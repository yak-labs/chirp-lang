package chirp

import (
	"testing"
)

func TestRegexp(a *testing.T) {
	fr := New()
	fr.Eval(MkString(`
		must 1 [regexp "asdf" "asdf"]
	`))

	fr.Eval(MkString(`
		must 1 [regexp "a..fq?" "asdf"]
	`))

	fr.Eval(MkString(`
		must 0 [regexp "qwera..fq?" "asdf"]
	`))
}

func TestRegexpNoCase(a *testing.T) {
	fr := New()
	fr.Eval(MkString(`
		must 0 [regexp "asdf" "ASDF"]
	`))

	fr.Eval(MkString(`
		must 1 [regexp -nocase "asdf" "ASDF"]
	`))
}
