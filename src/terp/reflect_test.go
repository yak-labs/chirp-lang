package terp_test

import (
	"testing"

	. "terp"
)

var reflectTests = `
	must	abcdefghi	[go-call /fmt/Sprintf {abc%sghi} def]

	set buf [go-call /bytes/NewBufferString Hello]
	must 5 [go-send $buf Len]
	must Hello [go-send $buf String]
	go-send $buf WriteString " World!"
	must 12 [go-send $buf Len]
	must "Hello World!" [go-send $buf String]
	must [list [+ 64 8] 1] [go-send $buf ReadRune]
	go-send $buf Truncate 3
	must "ell" [go-send $buf String]

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
	set z [go-call /math/big/NewInt 0]
	set x [go-call /math/big/NewInt 10 ]
	set y [go-call /math/big/NewInt 100 ]
	go-send $z Exp $x $y "" 
	must $Googol [go-send $z String] 

    list -- Factorial function using /math/big numbers.
	proc big_factorial n {
	  if {[< $n 1]} {
	  	return [go-call /math/big/NewInt 1]
	  }
	  set z [go-call /math/big/NewInt 0]
	  go-send $z Mul [go-call /math/big/NewInt $n] [big_factorial [- $n 1]]
	  set z
	}
	list -- Print a few for fun.
	foreach n [yrange 5] {
		puts "fact $n == [go-send [big_factorial $n] String]"
	}

	list -- fact70 is about 20% bigger than Googol.
	set fact70 [go-send [big_factorial 70] String]
	puts Googol=$Googol
	puts fact70=$fact70
	must [string length $Googol] [string length $fact70]
`

func TestReflect(a *testing.T) {
	New().Eval(MkString(reflectTests))
}
