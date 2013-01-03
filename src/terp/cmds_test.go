package terp_test

import (
	"testing"

	. "terp"
)

var cmdTests = `
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
  must 23 [if {[== 0.5 [/ 1 2]]} {+ 20 3} else {+ 40 2}]

  set a 100; set b 20; must 120 [+ [set a] [set b]]
  must 81028 [+ "8[set a]8" [set b]]
  must 81028 [+ 8[set a]8 [set b]]

  must cc [lat [list aa bb cc dd ee] 2]
  must cc [lat " aa bb cc dd ee" 2]
  must cc [lat { aa bb cc dd ee } 2]
  must c [sat abcde 2]

  must 5 [llen [list aa bb cc dd ee] ]
  must 1 [slen [llen [list aa bb cc dd ee] ]]
  must 2 [slen [llen [list aa bb cc dd ee 1 2 3 4 5] ]]
  must 0 [slen ""]

  proc double {x} {+ [set x] [set x]}
  proc square {x} {* [set x] [set x]}
  must 16 [double 8]
  must 16 [square 4]
  proc triang n {
    if {[< [set n] 1]} {
      + 0
    } else {
      + [set n] [triang [- [set n] 1]]
    }
  }
  must 15 [triang 5]

  proc range n {
    if {[< $n 1]} {
      list
    } else {
      set m [- $n 1]
      concat [range $m] [list $m]
    }
  }
  must "" [range 0]
  must "0" [range 1]
  must "0 1 2 3 4" [range 5]

  yproc yrange n {
      set i 0
      while {[< $i $n]} {
          yield $i 
          set i [+ $i 1]
      }
  }
  must "" [concat [yrange 0]]
  must "0" [concat [yrange 1]]
  must "0 1 2 3 4" [concat [yrange 5]]

  yproc ytriangs nums {
      foreach n $nums {
          yield [triang $n]
      }
  }
  must "1 3 6 10 15" [concat [ytriangs "1 2 3 4 5"]]

  yproc naturals {} {
      set i 0
      while {[+ 1]} {
          yield $i 
          set i [+ $i 1]
      }
  }
  yproc ytriangs_lt n {
    catch {
  	  foreach i [naturals] {
		set x [triang $i]
		if {[>= $x $n]} {error RETURN}
		yield $x
	  }
	} what
  }
  must [list 0 1 3 6 10 15 21 28 36 45 55 66 78 91] [concat [ytriangs_lt 100]]

  proc factorial_with_while n {
  	set z 1
	while {[> $n 0]} {
		set z [* $z $n]
		set n [- $n 1]
	}
	set z
  }
  must 120 [factorial_with_while 5]
  must 1 [factorial_with_while 1]
  must 1 [factorial_with_while 0]

  must 4 [catch continue x]
  must 3 [catch break x]
  must 2 [catch "return foo" x]
  must foo $x

  list -- Test of "return"
  proc and list {
	foreach cmd $list {
		set b [eval $cmd]
		if [== 0 $b] {
			return 0
		}
	}
	return 1
  }
  must 1 [and [list {< 2 4} {< 4 8} {< 8 16}]]
  must 0 [and [list {< 2 4} {> 4 8} {< 8 16}]]

  list -- Test of "break"
  proc five {   } {
  	foreach i [naturals] {
		if {[== $i 5]} break
	}
	return $i
  }
  must 5 [five]

  list -- Test of "continue"
  proc six {
  } {
  	foreach i [naturals] {
		if {[< $i 6]} continue
		break
	}
	return $i
  }
  must 6 [six]

  list -- test of "upvar"
  proc UpSet {name x} {
  	upvar 1 $name a
	set a $x
  }
  proc UpGet {name} {
  	upvar 1 $name a
	set a
  }
  set bar 42
  must 42 [UpGet bar]
  UpSet bar 54
  must 54 $bar

  list -- test of "hash"
  set h [hash]
  hset $h color purple
  must purple [hget $h color]
  hset $h pigs [list S M L]
  must M [lat [hget $h pigs] 1]
  hset $h color red
  must red [hget $h color]
  must "color pigs" [hkeys $h]
  hdel $h color
  must "pigs" [hkeys $h]

  proc F s {
  	return "$s 0"
  }
  mixin One {
	  proc F s {
		return "$s 1 [super F $s]"
	  }
  }
  mixin Two {
	  proc F s {
		return "$s 2 [super F $s]"
	  }
  }
  mixin Three {
	  proc F s {
		return "$s 3 [super F $s]"
	  }
  }
  must "foo 3 foo 2 foo 1 foo 0" [F "foo"]

  proc demo1 { a b c d e } { list $a $b $c $d $e }
  interp-new S1
  interp-alias S1 demo2 "demo1 1 2 3"
  must [list 1 2 3 x y] [interp-eval S1 {demo2 x y}]

  proc helloArgs { x args } {
  	return [llen $args]
  }

  must 3 [helloArgs a b c d]
  must 4 [helloArgs a b c d e]
  must 0 [helloArgs a]
 
  proc helloArgs2 { x args } {
  	return [list $x $args]
  }
  must "a {b c d e}" [helloArgs2 a b c d e]

`

func TestFoo(a *testing.T) {
	fr := New()
	fr.Eval(MkString(cmdTests))
}
