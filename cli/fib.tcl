proc fib {n} {
	if {$n < 2} {
		return $n
	} else {
		# set a [fib [expr {$n - 1}]]
		# set b [fib [expr {$n - 2}]]
		# expr {$a + $b}
		expr {[fib [expr {$n - 1}]]  + [fib [expr {$n - 2}]] }
	}
}

if {[llength $argv] > 1} {
	foreach x [lrange $argv 1 end] {
		puts "$x -> [fib $x]"
	}
} else {
	foreach x {3 6 9 12 15 18 21} {
		puts "$x -> [fib $x]"
	}
}
