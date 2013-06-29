package terp_test

import (
	"testing"

	. "terp"
)

var cmdTests = `
  #must ignore comments

  # dumbCompileSequence can compile this:
  # Note that "args" means varargs,
  # and that eval concatenates lists.
  proc sum {args} { eval + $args }

  must 0 [+ ]
  must 5 [+ 5]
  must 9 [+ 4 5]
  must 9.5 [+ 4 5.5]
  must 15 [+ 5 4 3 2 1 0]
  must 15 [sum 5 [+ 2 2] [+ 1 2] 2 1 [- 99 99]]

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

  must cc [lindex [list aa bb cc dd ee] 2]
  must cc [lindex " aa bb cc dd ee" 2]
  must cc [lindex { aa bb cc dd ee } 2]
  must c [string index abcde 2]

  must 5 [llength [list aa bb cc dd ee] ]
  must 1 [string length [llength [list aa bb cc dd ee] ]]
  must 2 [string length [llength [list aa bb cc dd ee 1 2 3 4 5] ]]
  must 0 [string length ""]

  must {B2 a1 a10 a2 b1} [lsort {a10 B2 b1 a1 a2}]

  # dumbCompileSequence can compile this:
  proc double {x} {+ $x $x}
  # dumbCompileSequence can compile this:
  proc square {x} {* $x $x}

  must 16 [double 8]
  must 16 [square 4]

  proc triang n {
    if {[< [set n] 1]} {
      sum 0
    } else {
      sum [set n] [triang [- [set n] 1]]
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
          set i [sum $i 1]
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
  # dumbCompileSequence can compile this:
  proc UpSet {name x} {
  	upvar 1 $name a
	set a $x
  }
  # dumbCompileSequence can compile this:
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
  must M [lindex [hget $h pigs] 1]
  hset $h color red
  must red [hget $h color]
  must "color pigs" [hkeys $h]
  hdel $h color
  must "pigs" [hkeys $h]

  proc F s {
  	return "$s 0"
  }
  mixin One {
      # dumbCompileSequence can compile this:
      proc mix_number {} { return 1 } ; list -- mixin-local proc.

	  proc F s {
		return "$s [mix_number] [super F $s]"
	  }
  }
  mixin Two {
      # dumbCompileSequence can compile this:
      proc mix_number {} { return 2 } ; list -- mixin-local proc.

	  proc F s {
		return "$s [mix_number] [super F $s]"
	  }
  }
  mixin Three {
      # dumbCompileSequence can compile this:
      proc mix_number {} { return 3 } ; list -- mixin-local proc.

	  proc F s {
		return "$s [mix_number] [super F $s]"
	  }
  }
  must "foo 3 foo 2 foo 1 foo 0" [F "foo"]

  # dumbCompileSequence can compile this:
  proc demo1 { a b c d e } { list $a $b $c $d $e }

  proc helloArgs { x args } {
  	return [llength $args]
  }

  must 3 [helloArgs a b c d]
  must 4 [helloArgs a b c d e]
  must 0 [helloArgs a]

  proc helloArgs2 { x args } {
  	return [list $x $args]
  }
  must "a {b c d e}" [helloArgs2 a b c d e]

  must abcdefg [string range abcdefghij -99 6]
  must abcdefghij [string range abcdefghij -99 99]
  must "" [string range abcdefghij 3 2]
  must d [string range abcdefghij 3 3]
  must de [string range abcdefghij 3 4]
  must defg [string range abcdefghij 3 6]
  must defghij [string range abcdefghij 3 ""]
  must defghij [string range abcdefghij 3 end]
  must defgh [string range abcdefghij 3 -2]
  must j [string range abcdefghij 9 10]


  must abcdef [string slice abcdefghij -99 6]
  must abcdefghij [string slice abcdefghij -99 99]
  must "" [string slice abcdefghij 3 2]
  must "" [string slice abcdefghij 3 3]
  must d [string slice abcdefghij 3 4]
  must def [string slice abcdefghij 3 6]
  must defghij [string slice abcdefghij 3 ""]
  must defghij [string slice abcdefghij 3 end]
  must defgh [string slice abcdefghij 3 -2]
  must j [string slice abcdefghij 9 10]

  must {{} a b {} c {}} [split /a/b//c/ /]
  must {a b c} [dropnull [split /a/b/c /]]
  must {/a/b/c d e f} [split "/a/b/c d e f"]
  must {a b c} [join {  a   b   c }]
  must {a:=b:=c} [join {  a   b   c } :=]

	set iincr 0
	must 1 [incr iincr 1]
	must 2 [incr iincr 1]
	must 3 [incr iincr 1]
	must 3 [set iincr]
	must 1 [incr iincr -2]
	must 1 [set iincr]

	must 10 [string first a 0123456789abcdef]
	must -1 [string first z 0123456789abcdef]

	set str {}
	must asdf [append str asdf]
	must asdfqwer123 [append str qwer 123]
	must asdfqwer123 [append str]
	must asdfqwer123 [set str]

	set str 0123456789abcdef
	must a [string index $str 10]
	must {} [string index $str -1]
	must {} [string index $str 100]

	set key one
	set arr(one) foo
	must foo $arr($key)
	set arr(b) one
	must foo $arr($arr(b))
	set arr(z) n
	must foo $arr(o$arr(z)e)
	must 1 [catch {list $arr(aaa$arr(z)zzz) } what]

	set link [ tag a  href http://www.foo.com/bar  alt FooBar  click ]
	set ht [ ht [htraw "&lt;"] <One> \040 Two&Three \  $link ]
	set expect {&lt;&lt;One&gt; Two&amp;Three <a href="http://www.foo.com/bar" alt="FooBar">click</a
>}

	# destructuring list assignment:
	set nada,uno,dos,tres [yrange 10]
	must 0/1/2/3 $nada/$uno/$dos/$tres

	set nada,uno,dos,tres [yrange 2]
	must 0/1// $nada/$uno/$dos/$tres

	foreach 1,2 [list "10 100" "20 200"] {
		must $2 [expr $1 * 10]
	}

	# Propagation of error from yproc to consumer.
	yproc barfer {} {error BARF}
	must 1 [catch {concat [barfer]} what]
	must 1 [string match BARF* $what]
	must 1 [catch {foreach x [barfer] {error NOTREACHED}} what]
	must 1 [string match BARF* $what]

    # Make class Obj.
	set Obj [subclass ""]
    # Make class A.
	set A [subclass $Obj]
    # Instantiate an A.
	set a [new $A]
	# Underscore variables are members of object.
	meth $A mulx y {expr $_x*$y}
	meth $A setx x {set _x $x}
	meth $A getx {} {return $_x}
	# "on" sends message.
	on $a setx 100
	must 100 [on $a getx]
	must 400 [on $a mulx 4]
	# TODO: inheritance.
	# TODO: super.

	set a 111; set B 222

	must {aaa$a foo[expr 3*8]bar$B.zzz\\0105} [
		subst -noback -nocommand -novar {aaa$a foo[expr 3*8]bar$B.zzz\0105}]

	must "aaa\$a foo\[expr 3*8\]bar\$B.zzz\b5" [
		subst -noc -nov {aaa$a foo[expr 3*8]bar$B.zzz\0105}]

	must "aaa111 foo\[expr 3*8\]bar222.zzz\b5" [
		subst -noc {aaa$a foo[expr 3*8]bar$B.zzz\0105}]

	must "aaa111 foo24bar222.zzz\b5" [
		subst {aaa$a foo[expr 3*8]bar$B.zzz\0105}]

	# http://wiki.yak.net/1030
	must 3 [case abc in {a b} {list 1} default {list 2} a* {list 3}]
	must 1 [case a in {   {a b} {list 1}   default {list 2}   a* {list 3} }]
	must 2 [case xyz {   {a b}     {list  1}    default {list 2}   a*     {list 3} }]
`

func TestFoo(a *testing.T) {
	fr := New()
	fr.Eval(MkString(cmdTests))
}

func TestStringMatchExact(t *testing.T) {
	// Positive pattern matches.
	ppats, pstrs := make([]string, 0, 4), make([]string, 0, 4)

	// Negative pattern matches.
	npats, nstrs := make([]string, 0, 4), make([]string, 0, 4)

	ppats, pstrs = append(ppats, "asdfqwer1234"), append(pstrs, "asdfqwer1234")

	npats, nstrs = append(npats, "asdf"), append(nstrs, "asdfqwer")
	npats, nstrs = append(npats, "asdfqwer"), append(nstrs, "asdf")
	npats, nstrs = append(npats, "z"), append(nstrs, "")

	for i, p := range ppats {
		matched := StringMatch(p, pstrs[i])
		if !matched {
			t.Errorf("StringMatch should have matched pattern %q with string %q", p, pstrs[i])
		}
	}

	for i, p := range npats {
		matched := StringMatch(p, nstrs[i])
		if matched {
			t.Errorf("StringMatch should NOT have matched pattern %q with string %q", p, nstrs[i])
		}
	}
}

func TestStringMatchBackSlash(t *testing.T) {
	pat, str := "\\*\\?\\[\\]\\\\", "*?[]\\"

	if !StringMatch(pat, str) {
		t.Errorf("StringMatch should have matched pattern %q with string %q", pat, str)
	}
}

func TestStringMatchStar(t *testing.T) {
	// Positive pattern matches.
	ppats, pstrs := make([]string, 0, 4), make([]string, 0, 4)

	// Negative pattern matches.
	npats, nstrs := make([]string, 0, 4), make([]string, 0, 4)

	ppats, pstrs = append(ppats, "*"), append(pstrs, "")
	ppats, pstrs = append(ppats, "*"), append(pstrs, "asdfqwer1234")
	ppats, pstrs = append(ppats, "*.jpg"), append(pstrs, "picture.jpg")
	ppats, pstrs = append(ppats, "house*dog*cat"), append(pstrs, "house123dog456cat")
	ppats, pstrs = append(ppats, "thisismytest*"), append(pstrs, "thisismytestthisoneismine")
	ppats, pstrs = append(ppats, "house*\\*dog*\\*cat"), append(pstrs, "house123*dog456*cat")

	npats, nstrs = append(npats, "abc*z"), append(nstrs, "abc")
	npats, nstrs = append(npats, "abc*z"), append(nstrs, "abcdefg")

	for i, p := range ppats {
		matched := StringMatch(p, pstrs[i])
		if !matched {
			t.Errorf("StringMatch should have matched pattern %q with string %q", p, pstrs[i])
		}
	}

	for i, p := range npats {
		matched := StringMatch(p, nstrs[i])
		if matched {
			t.Errorf("StringMatch should NOT have matched pattern %q with string %q", p, nstrs[i])
		}
	}
}

func TestStringMatchQuestion(t *testing.T) {
	pat, str := "qwer?asdf", "qwer0asdf"
	pat2, str2 := "qwer?asdf", "qwer01asdf"

	if !StringMatch(pat, str) {
		t.Errorf("StringMatch should have matched pattern %q with string %q", pat, str)
	}

	if StringMatch(pat2, str2) {
		t.Errorf("StringMatch should NOT have matched pattern %q with string %q", pat2, str2)
	}
}

func TestStringMatchBracket(t *testing.T) {
	// Positive pattern matches.
	ppats, pstrs := make([]string, 0, 4), make([]string, 0, 4)

	// Negative pattern matches.
	npats, nstrs := make([]string, 0, 4), make([]string, 0, 4)

	ppats, pstrs = append(ppats, "[a-z]"), append(pstrs, "a")
	ppats, pstrs = append(ppats, "[a-z]"), append(pstrs, "i")
	npats, nstrs = append(npats, "[a-z]"), append(nstrs, "1")

	ppats, pstrs = append(ppats, "[ab]"), append(pstrs, "a")
	ppats, pstrs = append(ppats, "[ab]"), append(pstrs, "b")
	npats, nstrs = append(npats, "[ab]"), append(nstrs, "c")

	ppats, pstrs = append(ppats, "12[ab2-4cd]45"), append(pstrs, "12345")
	ppats, pstrs = append(ppats, "12[ab2-4cd]45"), append(pstrs, "12b45")
	ppats, pstrs = append(ppats, "12[ab2-4cd]45"), append(pstrs, "12d45")
	npats, nstrs = append(npats, "12[ab2-4cd]45"), append(nstrs, "12145")
	npats, nstrs = append(npats, "12[ab2-4cd]45"), append(nstrs, "12545")

	for i, p := range ppats {
		matched := StringMatch(p, pstrs[i])
		if !matched {
			t.Errorf("StringMatch should have matched pattern %q with string %q", p, pstrs[i])
		}
	}

	for i, p := range npats {
		matched := StringMatch(p, nstrs[i])
		if matched {
			t.Errorf("StringMatch should NOT have matched pattern %q with string %q", p, nstrs[i])
		}
	}
}
