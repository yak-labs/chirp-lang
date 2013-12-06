proc minus1 n { set e [expr {[llength $n] - 1}]; lrange $n 1 $e }
proc minus2 n { set e [expr {[llength $n] - 1}]; lrange $n 2 $e }
proc plus {a b} { concat $a $b }
proc lt2 n { expr {[llength $n] < 2} }

proc fib {n} {
	# puts "fib <<< $n <<<"
	if [lt2 $n] {
		set z $n
	} else {
			# puts "$n >>minus1>> [minus1 $n]"
		set a [fib [minus1 $n]]
			# puts "a >>> $a"
			# puts "$n >>minus2>> [minus2 $n]"
		set b [fib [minus2 $n]]
			# puts "b >>> $b"
		set z [plus $a $b]
	}
	# puts "fib >>> $n >>> $z"
	set z
}

# foreach x {1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20}
foreach x {1 2 3 4 5 6 7 8 9 10 11 12 13 14 } {
	lappend n I
	puts "$x -> [llength [fib $n]]"
}
