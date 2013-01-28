proc triang n {
	if {[< $n 1]} {
		+ 0 
	} else {
		+ $n [triang [- $n 1]]
	}
}

yproc ytriangs nums {
	foreach n $nums {
		yield [triang $n]
	}
}

proc factorial n {
	if {[< $n 1]} {
		* 1 
	} else {
		* $n [factorial [- $n 1]]
	}
}

list -- Factorial function using /math/big numbers.
proc big_factorial n {
  if {[< $n 1]} {
    return [go-call /math/big/NewInt 1]
  }
  set z [go-call /math/big/NewInt 0]
  go-send $z Mul [go-call /math/big/NewInt $n] [big_factorial [- $n 1]]
  set z
}

proc range n {
	if {[< $n 1]} {
		list
	} else {
		set m [- $n 1]
		concat [range $m] [list $m]
	}
}

yproc yrange n {
	set i 0
	while {[< $i $n]} {
		yield $i
		set i [+ $i 1]
	}
}

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

foreach i [range 10] {
	puts "The triang of $i is [triang $i]"
	puts "The factorial of $i is [factorial $i]"
}

foreach i [yrange 11] {
	set n [* 10 $i]
	puts "The big_factorial of $n is [go-send [big_factorial $n] String]"
}

foreach t [ytriangs [yrange 10]] {
	puts "The triang is $t"
}

puts "The triangs under 100 are [concat [ytriangs_lt 100]]"
