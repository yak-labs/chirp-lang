proc home {w r} {
	display "HomePage" [get w]
}

proc display {page w} {
	set pagef "[get page].wik"
	/fmt/Fprintf [get w] "(display( %s ))" [/io/ioutil/ReadFile [get pagef]]
}

proc view {w r} {
	set page [
		getq [get r] page
	]

	display [get page] [get w]
}

proc list {w r} {
	set dinfo [/io/ioutil/ReadDir .]
	set finfo [index [get dinfo] 0]
	set fname [send [get finfo] Name]
	/fmt/Fprintf [get w] "( %s ) .... " [get fname]

	set flist [tolist [get dinfo]]
	foreach f [get flist] {
		set fname [send [get f] Name]
		/fmt/Fprintf [get w] "{ %s }  " [get fname]
	}
}

/net/http/HandleFunc / [http_handler home]
/net/http/HandleFunc /view [http_handler view]
/net/http/HandleFunc /list [http_handler list]

/net/http/ListenAndServe 127.0.0.1:8080 ""
