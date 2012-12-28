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

foreach i [range 10] {
	puts "The triang of $i is [triang $i]"
	puts "The factorial of $i is [factorial $i]"
}

foreach i [yrange 10] {
	puts "The triang of $i is [triang $i]"
}

foreach t [ytriangs [yrange 10]] {
	puts "The triang is $t"
}
