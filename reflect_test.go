package chirp_test

import . "github.com/yak-labs/chirp-lang"
import _ "github.com/yak-labs/chirp-lang/goapi/default"

import (
	"testing"
)

var reflectTests = `
	proc doto {creator args} {
		set z [uplevel 1 $creator]
		foreach a $args {
			uplevel 1 [list $z] $a
		}
		set z
	}
	set cp [/net/http/Cookie new]
	must ptr [$cp kind]
	must struct [$cp * kind]
	must *http.Cookie [[$cp type] String]

	set cv [[/net/http/Cookie ptrtype] mkslice 3]
	must 3 [llength $cv]
	must "" [$cv @ 0]
	must "" [$cv @ 1]
	must "" [$cv @ 2]
	# TODO: why didn't doto work?
	# $cv @ 1 = [doto [/net/http/Cookie new] {. Name = COLOR} {. Value = RED}]
	$cv @ 1 = [/net/http/Cookie new]
	$cv @ 1 . Name = COLOR_COLOR
	$cv @ 1 . Value = RED_RED
	must "COLOR_COLOR=RED_RED" [[$cv @ 1] String]

	$cv @ 1 . Name = bug
	$cv @ 1 . Value = f00f
	must bug=f00f [$cv @ 1 String]

	must	abcdefghi	[/fmt/Sprintf {abc%sghi} def]

	set buf [/bytes/NewBufferString Hello]
	must 5 [$buf Len]
	must Hello [$buf String]
	$buf WriteString " World!"
	must 12 [$buf Len]
	must "Hello World!" [$buf String]
	must [list [+ 64 8] 1] [$buf ReadRune]
	$buf Truncate 3
	must "ell" [$buf String]

    yproc yrange n {
        set i 0
        while {[< $i $n]} {
                yield $i
                set i [+ $i 1]
        }
    }

	list -- Create Googol by appending 100 0's to a "1".
	set Googol 1
	foreach i [yrange 100] {
		set Googol "[set Googol]0"
	}

    list -- Now raise 10 to the 100 power.
	set z [/math/big/NewInt 0]
	set x [/math/big/NewInt 10 ]
	set y [/math/big/NewInt 100 ]
	$z Exp $x $y "" 
	must $Googol [$z String] 

    list -- Factorial function using /math/big numbers.
	proc big_factorial n {
	  if {[< $n 1]} {
	  	return [/math/big/NewInt 1]
	  }
	  set z [/math/big/NewInt 0]
	  $z Mul [/math/big/NewInt $n] [big_factorial [- $n 1]]
	  set z
	}
	list -- Print a few for fun.
	foreach n [yrange 5] {
		puts "fact $n == [[big_factorial $n] String]"
	}

	list -- fact70 is about 20% bigger than Googol.
	set fact70 [[big_factorial 70] String]
	puts Googol=$Googol
	puts fact70=$fact70
	must [string length $Googol] [string length $fact70]
`

func TestReflect(a *testing.T) {
	SetDebugFromEnv()
	NewInterpreter().Eval(MkString(reflectTests))
}
