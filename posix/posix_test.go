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

  set f [open $path "a"]
  puts $f One
  puts $f Two
  puts -nonewline $f Three
  puts $f AndAHalf
  flush $f

  set f [open $path "r"]
  must One [gets $f]
  must Two [gets $f x ; set x]
  must ThreeAndAHalf [gets $f]
  close $f
`

func TestHt(a *testing.T) {
	fr := New()
	fr.Eval(MkString(tests))
}
