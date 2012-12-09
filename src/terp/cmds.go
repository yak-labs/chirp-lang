package terp

import (
	. "fmt"
	"strconv"
)

var Builtins map[string]Command = make(map[string]Command, 0)

func init() {
	Builtins["+"] = cmdPlus
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
