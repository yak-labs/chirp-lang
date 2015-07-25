set a [chirp int 42]
set pi [chirp float 3.14]
must 45 [expr {$a + 3}]
must 6.28 [expr {$pi * 2}]

set triad [/regexp/MustCompile {...}]
set vv [$triad FindAllIndex "abcrstxyz0" -1]
say A $vv
say A [chirp list $vv]
set z {}
foreach v [chirp list $vv] {
  say B $v
  say B [chirp list $v]
  foreach r [chirp list $v] {
    say C $r
    say C [chirp int $r]
    lappend z [chirp int $r]
  }
}
must {0 3 3 6 6 9} $z
