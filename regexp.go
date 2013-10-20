package chirp

import (
	"fmt"
	"regexp"
)

var regexpCache = make(map[string]*regexp.Regexp)

func regexpNoCase(exp string) string {
	return fmt.Sprintf("(?i)%s", exp)
}

func regexpFromCache(exp string) *regexp.Regexp {
	r, exists := regexpCache[exp]
	if !exists {
		r = regexp.MustCompile(exp)
		regexpCache[exp] = r
	}
	return r
}

func Regexp(exp string, nocase bool) *regexp.Regexp {
	if nocase {
		exp = regexpNoCase(exp)
	}

	return regexpFromCache(exp)
}

func RegexpMatch(exp, str string, nocase bool) bool {
	return Regexp(exp, nocase).MatchString(str)
}

func RegexpFindMatch(exp, str string, nocase bool) (bool, string) {
	r := Regexp(exp, nocase)
	return r.MatchString(str), r.FindString(str)
}

func RegexpFindSubmatch(exp, str string, nocase bool) (bool, []string) {
	r := Regexp(exp, nocase)
	return r.MatchString(str), r.FindStringSubmatch(str)
}

func cmdRegexp(fr *Frame, argv []T) T {
	nocase := false
	var exp string
	var str string

	if len(argv) < 2+1 {
		panic(fmt.Sprintf(
			"Expected 2 or more arguments, but got argv=%s",
			Showv(argv)))
	}

	arg_idx := 1
	if argv[arg_idx].String() == "-nocase" {
		nocase = true
		arg_idx++
		exp = argv[arg_idx].String()
		arg_idx++
		str = argv[arg_idx].String()
		arg_idx++
	} else {
		exp = argv[arg_idx].String()
		arg_idx++
		str = argv[arg_idx].String()
		arg_idx++
	}

	if len(argv) == arg_idx {
		return MkBool(RegexpMatch(exp, str, nocase))
	} else if len(argv) == arg_idx+1 {
		isMatch, match := RegexpFindMatch(exp, str, nocase)

		if len(match) == 0 {
			fr.SetVar(argv[arg_idx].String(), Empty)
		} else {
			fr.SetVar(argv[arg_idx].String(), MkString(match))
		}
		return MkBool(isMatch)
	} else {
		isMatch, submatches := RegexpFindSubmatch(exp, str, nocase)

		sub_idx := 0
		for arg_idx < len(argv) {
			if sub_idx < len(submatches) {
				fr.SetVar(argv[arg_idx].String(), MkString(submatches[sub_idx]))
			} else {
				fr.SetVar(argv[arg_idx].String(), Empty)
			}
			sub_idx++
			arg_idx++
		}

		return MkBool(isMatch)
	}
}

func init() {
	if Safes == nil {
		Safes = make(map[string]Command, 333)
	}

	Safes["regexp"] = cmdRegexp
}
