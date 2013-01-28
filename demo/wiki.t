proc GetQuery {r key} {
	go-send [go-send [go-getf $r URL] Query] Get $key
}

proc ModeHtml {w} {
	go-send [go-send $w Header] Set "Content-Type" "text/html"
}

proc starter {w r} {
	ModeHtml $w

	go-call /fmt/Fprintf $w "<ul>"
	go-call /fmt/Fprintf $w {<li> <a href="/list">List Of Pages</a>}
	go-call /fmt/Fprintf $w {<li> <a href="/">Home Page</a>}
	go-call /fmt/Fprintf $w "<li>(xyz=( %s ))  " [GetQuery $r xyz]
	go-call /fmt/Fprintf $w "<li>(Path=( %s ))  " [go-getf $r URL Path]
	go-call /fmt/Fprintf $w "<li>(RawQuery=( %s ))  " [go-getf $r URL RawQuery]
	go-call /fmt/Fprintf $w "<li>(User=( %s ))  " [go-getf $r URL User]
	go-call /fmt/Fprintf $w "<li>(Host=( %s ))  " [go-getf $r URL Host]
	go-call /fmt/Fprintf $w "<li>(Scheme=( %s ))  " [go-getf $r URL Scheme]
	go-call /fmt/Fprintf $w "</ul>"
}

proc Home {w r} {
	starter $w $r
	Display "HomePage" $w
}

proc Display {page w} {
	set pagef "$page.wik"
	go-call /fmt/Fprintf $w "(Display( %s ))" [go-call /io/ioutil/ReadFile $pagef]
}

proc View {w r} {
	starter $w $r
	set page [GetQuery $r page]
	if {[catch {go-call /io/ioutil/ReadFile "$page.wik"} junk]} {
		go-call /fmt/Fprintf $w "Page Not Found: $page"
	} else {
		Display $page $w
	}
}

proc List {w r} {
	starter $w $r
	go-call /fmt/Fprintf $w "<ul>"
	foreach f [go-call /io/ioutil/ReadDir .] {
		set fname [go-send $f Name]
		set m [go-send $ValidName FindStringSubmatch $fname]
		if {[null $m]} {
			go-call /fmt/Fprintf $w {<li> Invalid Filename: %s} $fname
		} else {
			set _,name $m
			go-call /fmt/Fprintf $w {<li> <a href="/view?page=%s">"%s"</a> for filename %s} $name $name $fname
		}
	}
	go-call /fmt/Fprintf $w "</ul>"
}

set ValidName [go-call /regexp/MustCompile {^([A-Z]+[a-z]+[A-Z][A-Za-z0-9_]*)[.]wik$}]

go-call /net/http/HandleFunc / [http_handler Home]
go-call /net/http/HandleFunc /view [http_handler View]
go-call /net/http/HandleFunc /list [http_handler List]

go-call /net/http/ListenAndServe 127.0.0.1:8080 ""
