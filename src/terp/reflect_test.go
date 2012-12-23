package terp_test

import (
	"testing"

	. "terp"
)

var reflectTests = `
	must	abcdefghi	[/fmt/Sprintf {abc%sghi} def]

	set buf [call /bytes/NewBufferString Hello]
	must 5 [send $buf Len]
	must Hello [send $buf String]
	send $buf WriteString " World!"
	must 12 [send $buf Len]
	must "Hello World!" [send $buf String]
	must [list [+ 64 8] 1] [send $buf ReadRune]
	send $buf Truncate 3
	must "ell" [send $buf String]
`

func TestReflect(a *testing.T) {
	New().TEval(MkTs(reflectTests))
}
