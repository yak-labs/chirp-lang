package terp_test

import (
  "testing"
  . "terp"
)

var tests = `
  must 9 [expr { [+ 4 5] }]
`

func TestExpr(a *testing.T) {
  fr := New()
  fr.TEval(MkTs(tests))
}
