# Chirp script to unbundle scenarios into data directory.

set TEXT_rx [/regexp/MustCompile {^/TEXT (.*)$}]
set all [/io/ioutil/ReadFile /dev/stdin]

proc Create path {
	set words [dropnull [split $path /]]
	set s [lat $words 0]
	set v [lat $words 1]
	set p [lat $words 2]
	set f [lat $words 3]
	set dir "data/s.$s/v.$v/p.$p/f.$f"
	set mode777 511
	/os/MkdirAll $dir $mode777
	return [/os/Create $dir/r.0]
}

set f ""
foreach line [split $all \n] {
	set m [send $TEXT_rx FindStringSubmatch $line]
	if [notnull $m] {
		# is a magic /TEXT line
		set path [lindex $m 1]
		if [notnull $f] {
			send $f Close
		}
		set f [Create $path]
	} else {
		# not a magic /TEXT line
		if [notnull $f] {
			send $f WriteString "$line\n"
		}
	}
}

if [notnull $f] {
	send $f Close
}
