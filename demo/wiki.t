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

/net/http/HandleFunc / [http_handler home]
/net/http/HandleFunc /view [http_handler view]

/net/http/ListenAndServe 127.0.0.1:8080 [nil]
