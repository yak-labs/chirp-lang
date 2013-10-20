package chirp

import (
	"fmt"
	"regexp"
)

var regexpCache = make(map[string]*regexp.Regexp)

func Regexp(exp, str string, nocase bool) bool {
	if nocase {
		exp = fmt.Sprintf("(?i)%s", exp)
	}

	r, exists := regexpCache[exp]
	if !exists {
		r = regexp.MustCompile(exp)
		regexpCache[exp] = r
	}
	return r.MatchString(str)
}

func cmdRegexp(fr *Frame, argv []T) T {
	nocase := false
	var exp string
	var str string

	if len(argv) < 2+1 || len(argv) > 3+1 {
		panic(fmt.Sprintf(
			"Expected 2 to 3 arguments, but got argv=%s",
			Showv(argv)))
	}

	if argv[1].String() == "-nocase" {
		nocase = true
		exp = argv[2].String()
		str = argv[3].String()
	} else {
		exp = argv[1].String()
		str = argv[2].String()
	}

	return MkBool(Regexp(exp, str, nocase))
}

func init() {
	if Safes == nil {
		Safes = make(map[string]Command, 333)
	}

	Safes["regexp"] = cmdRegexp
}
