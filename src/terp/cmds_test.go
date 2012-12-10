package terp_test

import (
	"testing"

	. "terp"
)

var tests = `
  must 0 [+ ]
  must 5 [+ 5]
  must 9 [+ 4 5]
  must 9.5 [+ 4 5.5]
  must 15 [+ 5 4 3 2 1 0]
  must 15 [+ 5 [+ 2 2] [+ 1 2] 2 1 [- 99 99]]

  must 9 [- 12 3]
  must 8.5 [- 12 3.5]
  must 8.5 [+ 12 -3.5]
  must -8.5 [+ -12 3.5]

  must 1 [*]
  must 100 [* 100]
  must 200 [* 100 2]
  must 500 [* 100 2 2.5]

  must 1 [== [+ 1 3] [+ 2 2]]
  must 1 [> 0.334 [/ 1 3]]
  must 1 [< 0.333 [/ 1 3]]
  must 23 [if {== 0.5 [/ 1 2]} {+ 20 3} else {+ 40 2}]

  set a 100; set b 20; must 120 [+ [get a] [get b]]
`

func TestFoo(a *testing.T) {
	fr := New()
	fr.Eval(tests)
}
