package chirp

import (
	"testing"
)

var tests = `
  must "<body>Foo</body\n>" [ht-peek [ht-tag body Foo]]
  must "<body>&Eacute;</body\n>" [ht-peek [ht-tag body [ht-entity Eacute]]]
  must "One&amp;Two&lt;Three&gt;Four&#39;s<b>Five</b\n>" [ht-peek [ht One& Two<Three> Four's [ht-tag b Five]]]
`

func TestHt(a *testing.T) {
	fr := New()
	fr.Eval(MkString(tests))
}
