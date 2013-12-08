package chirp

import (
	"testing"
)

var opaqueTests = `
	set obj [opaque wrap cover1 secret1]
	must cover1 [opaque show $obj ]
	must secret1 [opaque unwrap $obj ]

	set t1 [interp]
	must cover1 [$t1 Eval [list opaque show $obj]]
	must 1 [catch {$t1 Eval [list opaque unwrap $obj]} what]
`

func TestOpaque(a *testing.T) {
	fr := NewInterpreter()
	fr.Eval(MkString(opaqueTests))
}
