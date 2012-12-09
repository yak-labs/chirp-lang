package terp

import (
	. "fmt"
	"strconv"
)

var Builtins map[string]Command = make(map[string]Command, 0)

func init() {
	Builtins["+"] = cmdPlus
	Builtins["must"] = cmdMust
}

func cmdPlus(t *Terp, args []Any) Any {
	var z float64 = 0
	for _, a := range args {
		s := Str(a)
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(Sprintf("cmdPlus: Cannot ParseFloat: %q", s))
		}
		z += f
	}
	return z //Str(z)
}

func cmdMust(t *Terp, args []Any) Any {
	x := Str(args[0])
	y := Str(args[1])

	if x == y {
		return args[1]
	}

	panic("FAILED: must: " + Repr(args))
}
