package terp_test

import (
	"testing"

	. "terp"
)

var reflectTests = `
	must	abcdefghi	[/fmt/Sprintf {abc%sghi} def]

	set buf [call /bytes/NewBufferString Hello]
	must 5 [send $buf Len]
	must Hello [send $buf String]
	send $buf WriteString " World!"
	must 12 [send $buf Len]
	must "Hello World!" [send $buf String]
	must [list [+ 64 8] 1] [send $buf ReadRune]
	send $buf Truncate 3
	must "ell" [send $buf String]

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
	set x [ /math/big/NewInt 10 ]
	set y [ /math/big/NewInt 100 ]
	send $z Exp $x $y "" 
	must $Googol [send $z String] 

    list -- Factorial function using /math/big numbers.
	proc big_factorial n {
	  if {[< $n 1]} {
	  	return [/math/big/NewInt 1]
	  }
	  set z [/math/big/NewInt 0]
	  send $z Mul [/math/big/NewInt $n] [big_factorial [- $n 1]]
	  set z
	}
	list -- Print a few for fun.
	foreach n [yrange 5] {
		puts "fact $n == [send [big_factorial $n] String]"
	}

	list -- fact70 is about 20% bigger than Googol.
	set fact70 [send [big_factorial 70] String]
	puts Googol=$Googol
	puts fact70=$fact70
	must [slen $Googol] [slen $fact70]
`

func TestReflect(a *testing.T) {
	New().Eval(MkTs(reflectTests))
}
