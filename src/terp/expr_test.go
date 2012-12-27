package terp_test

import (
  "testing"
  . "terp"
)

var exprTests = `
  must 9 [expr { [+ 4 5] }]
  must 9 [expr { 4 + 5 }]
  must 20 [expr { 4 * 5 }]
  must 16 [expr { 4 * 5 + 3 - 7 }]
`

func TestExpr(a *testing.T) {
  fr := New()
  fr.TEval(MkTs(exprTests))
}
