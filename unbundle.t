# Chirp script to unbundle scenarios into data directory.

set TEXT_rx [go-call /regexp/MustCompile {^/TEXT (.*)$}]
set all [go-call /io/ioutil/ReadFile /dev/stdin]

proc Create path {
	set s,v,p,f [dropnull [split $path /]]
	set dir "data/s.$s/v.$v/p.$p/f.$f"
	set mode777 511
	go-call /os/MkdirAll $dir $mode777
	return [go-call /os/Create $dir/r.0]
}

set f ""
foreach line [split $all \n] {
	set m [go-send $TEXT_rx FindStringSubmatch $line]
	if [notnull $m] {
		# is a magic /TEXT line
		set path [lindex $m 1]
		if [notnull $f] {
			go-send $f Close
		}
		set f [Create $path]
	} else {
		# not a magic /TEXT line
		if [notnull $f] {
			go-send $f WriteString "$line\n"
		}
	}
}

if [notnull $f] {
	go-send $f Close
}
