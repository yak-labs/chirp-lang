package posix

import (
	. "github.com/yak-labs/chirp-lang"
	"testing"
)

var tests = `
  # TODO: portable way of finding the /tmp directory.
  set path "[file tempdir][file separator]tmp.posix_test.go.tmp"
  set f [open $path "w"]
  log 0 {first open -> $f}
  set g [open $path "r"]
  log 0 {second open -> $g}
  close $f
  close $g
`

func TestHt(a *testing.T) {
	fr := New()
	fr.Eval(MkString(tests))
}
