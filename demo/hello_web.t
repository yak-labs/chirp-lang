proc home {w r} {
	go-call /fmt/Fprintf $w "Hello Web!  ###########  w=%s  ###########  r=%s\n" $w $r
}

go-call /net/http/HandleFunc / [http_handler home]

go-call /net/http/ListenAndServe 127.0.0.1:8080 ""
