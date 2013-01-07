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

  set conjLazy 123
  must 1 [expr 3 >= 3 && 4 > 3]
  must 0 [expr 3 >= 3 && 4 < 3]
  must 0 [expr 3 > 3 && 4 > 3]
  must 0 [expr 3 > 3 && 4 < 3]
  must 1 [expr {$conjLazy == 123 && [set conjLazy 567] == 567}]
  must 0 [expr {$conjLazy == 123 && [set conjLazy 789] == 789}]
  must 567 $conjLazy

  set disjLazy 123
  must 1 [expr 3 >= 3 || 4 > 3]
  must 1 [expr 3 >= 3 || 4 < 3]
  must 1 [expr 3 > 3 || 4 > 3]
  must 0 [expr 3 > 3 || 4 < 3]
  must 1 [expr {$disjLazy != 123 || [set disjLazy 567] == 567}]
  must 1 [expr {$disjLazy == 567 || [set disjLazy 789] == 789}]
  must 567 $disjLazy

  must 123 [expr 1 ? 123 : 456]
  must 456 [expr 0 ? 123 : 456]
  must 456 [expr 0 ? 123 : 1 ? 456 : 789]
  must 789 [expr 0 ? 123 : 0 ? 456 : 789]

  must 0 [expr {{word one} lt "word $a"}]
  must 1 [expr {"word $a" lt {word one}}]
  must 1 [expr {{word one} gt "word $a"}]
  must 0 [expr {"word $a" gt {word one}}]
  must 0 [expr {"asdf" eq {qwer}}]
  must 1 [expr {"qwer" eq {qwer}}]
  must 1 [expr {"asdf" ne {qwer}}]
  must 0 [expr {"qwer" ne {qwer}}]
`

func TestExpr(a *testing.T) {
	fr := New()
	fr.Eval(MkString(exprTests))
}
