proc home {w r} {
	/fmt/Fprintf [get w] "Hello Web!  ###########  w=%#q  ###########  r=%#q\n" [get w] [get r]
}

/net/http/HandleFunc / [http_handler home]

/net/http/ListenAndServe 127.0.0.1:8080 [nil]
