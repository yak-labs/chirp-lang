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
	usage := `?-nocase? expression string ?matchVar ?groupVars...?? -> bool`
	dashes, expT, strT, vars := ArgDash2vUsage(argv, usage)

	exp := expT.String()
	str := strT.String()
	nocase := false

	for _, dash := range dashes {
		switch dash {
		case "-nocase":
			nocase = true
		default:
			panic("Unknown dash option: " + dash)
		}
	}

	var_idx := 0
	if len(vars) == var_idx {
		return MkBool(RegexpMatch(exp, str, nocase))
	} else if len(vars) == var_idx+1 {
		isMatch, match := RegexpFindMatch(exp, str, nocase)

		if len(match) == 0 {
			fr.SetVar(vars[var_idx].String(), Empty)
		} else {
			fr.SetVar(vars[var_idx].String(), MkString(match))
		}
		return MkBool(isMatch)
	}
	// else:

	isMatch, submatches := RegexpFindSubmatch(exp, str, nocase)

	sub_idx := 0
	for var_idx < len(vars) {
		if sub_idx < len(submatches) {
			fr.SetVar(vars[var_idx].String(), MkString(submatches[sub_idx]))
		} else {
			fr.SetVar(vars[var_idx].String(), Empty)
		}
		sub_idx++
		var_idx++
	}

	return MkBool(isMatch)
}

func init() {
	if Safes == nil {
		Safes = make(map[string]Command, 333)
	}

	Safes["regexp"] = cmdRegexp
}
