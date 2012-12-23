proc home {w r} {
	display "HomePage" $w
}

proc display {page w} {
	set pagef "[get page].wik"
	/fmt/Fprintf $w "(display( %s ))" [/io/ioutil/ReadFile $pagef]
}

proc view {w r} {
	set page [
		getq $r page
	]

	display $page $w
}

proc list {w r} {
	set dinfo [/io/ioutil/ReadDir .]
	set finfo [index $dinfo 0]
	set fname [send $finfo Name]
	/fmt/Fprintf $w "( %s ) .... " $fname

	set flist  $dinfo
	foreach f  $flist {
		set fname [send $f Name]
		/fmt/Fprintf $w "{ %s }  " $fname
	}
}

/net/http/HandleFunc / [http_handler home]
/net/http/HandleFunc /view [http_handler view]
/net/http/HandleFunc /list [http_handler list]

/net/http/ListenAndServe 127.0.0.1:8080 ""
