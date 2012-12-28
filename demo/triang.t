proc triang n {
	if {[< $n 1]} {
		+ 0 
	} else {
		+ $n [triang [- $n 1]]
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

foreach i [range 20] {
	puts "The triang of $i is [triang $i]"
	puts "The factorial of $i is [factorial $i]"
}
