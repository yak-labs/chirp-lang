proc home {w r} {
	call /fmt/Fprintf $w "Hello Web!\n  w=%#q\n  r=%#q\n" w r
}

call /net/http/HandleFunc / [http_handler home]

call /net/http/ListenAndServe 127.0.0.1:8080 [nil]
