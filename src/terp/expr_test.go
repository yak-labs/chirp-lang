package terp_test

import (
	. "terp"
	"testing"
)

var exprTests = `
  must 9 [expr [+ 4 5]]
  must 9 [expr 4 + 5]
  must 20 [expr 4 * 5]
  must 16 [expr 4 * 5 + 3 - 7]
  must 16 [expr 16]
  must 44 [expr { 3 * 2 * 4 + 5 * 4} ]

  set a 3
  set b 6
  must 6.1 [expr { 3.1 + $a }]
  must 5.6 [expr { 2 + "$a.$b" }]
  must 8 [expr { 4*[llen "6 2"] }]

  must 0 [expr { 4 == 3 }]
  must 1 [expr { 3 == 3 }]
  must 0 [expr { 3 != 3 }]
  must 1 [expr { 4 != 3 }]
  must 0 [expr { 4 < 3 }]
  must 1 [expr { 3 < 4 }]
  must 0 [expr { 3 < 3 }]
  must 0 [expr { 4 <= 3 }]
  must 1 [expr { 3 <= 4 }]
  must 1 [expr { 3 <= 3 }]
  must 1 [expr { 4 > 3 }]
  must 0 [expr { 3 > 4 }]
  must 0 [expr { 3 > 3 }]
  must 1 [expr { 4 >= 3 }]
  must 0 [expr { 3 >= 4 }]
  must 1 [expr { 3 >= 3 }]
`

func TestExpr(a *testing.T) {
	fr := New()
	fr.Eval(MkString(exprTests))
}
