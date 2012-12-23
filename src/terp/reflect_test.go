package terp_test

import (
	"testing"

	. "terp"
)

var reflectTests = `
	must	abcdefghi	[/fmt/Sprintf {abc%sghi} def]

	set buf [call /bytes/NewBufferString Hello]
	must 5 [send $buf Len]
`

func TestReflect(a *testing.T) {
	New().TEval(MkTs(reflectTests))
}
