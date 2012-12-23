package terp_test

import (
	"testing"

	. "terp"
)

var reflectTests = `
	must	abcdefghi	[/fmt/Sprintf {abc%sghi} def]
`

func TestReflect(a *testing.T) {
	New().TEval(MkTs(reflectTests))
}
