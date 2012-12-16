proc home {w r} {
	call /fmt/Fprintf [get w] "Hello Web!  ###########  w=%#q  ###########  r=%#q\n" [get w] [get r]
}

call /net/http/HandleFunc / [http_handler home]

call /net/http/ListenAndServe 127.0.0.1:8080 [nil]
