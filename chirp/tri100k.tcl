proc tri {n} {
	if {$n < 1} {
		return 0
	} else {
		expr {$n + [tri [expr {$n - 1}]]}
	}
}

puts [tri 100000]
