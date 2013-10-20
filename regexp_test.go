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

func TestRegexpMatchVar(a *testing.T) {
	fr := New()
	fr.Eval(MkString(`
		must 1 [regexp owl {cow owl dog bear} owlvar]
		must owl $owlvar
	`))

	fr.Eval(MkString(`
		must 0 [regexp cat {cow owl dog bear} catvar]
		must {} $catvar
	`))
}

func TestRegexpSubmatchVar(a *testing.T) {
	fr := New()
	fr.Eval(MkString(`
		must 1 [regexp http://(yak\.net) {this is http://yak.net} url domain]
		must http://yak.net $url
		must yak.net $domain
	`))

	fr.Eval(MkString(`
		must 1 [regexp {http://([^/]+)/} {http://yak.net/} url domain path]
		must http://yak.net/ $url
		must yak.net $domain
		must {} $path
	`))

	fr.Eval(MkString(`
		must 1 [regexp {http://([^/]+)/(.+)} {http://yak.net/asdf} url domain path]
		must http://yak.net/asdf $url
		must yak.net $domain
		must asdf $path
	`))
}
