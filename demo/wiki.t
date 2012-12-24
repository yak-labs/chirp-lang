proc GetQuery {r key} {
	send [send [get r URL] Query] Get $key
}

proc ModeHtml {w} {
	send [send $w Header] Set "Content-Type" "text/html"
}

proc starter {w r} {
	ModeHtml $w

	/fmt/Fprintf $w "<ul>"
	/fmt/Fprintf $w "<li>(xyz=( %s ))  " [GetQuery $r xyz]
	/fmt/Fprintf $w "<li>(Path=( %s ))  " [get r URL Path]
	/fmt/Fprintf $w "<li>(RawQuery=( %s ))  " [get r URL RawQuery]
	/fmt/Fprintf $w "<li>(User=( %s ))  " [get r URL User]
	/fmt/Fprintf $w "<li>(Host=( %s ))  " [get r URL Host]
	/fmt/Fprintf $w "<li>(Scheme=( %s ))  " [get r URL Scheme]
	/fmt/Fprintf $w "</ul>"
}

proc home {w r} {
	starter $w $r
	display "HomePage" $w
}

proc display {page w} {
	
	set pagef "[get page].wik"
	/fmt/Fprintf $w "(display( %s ))" [/io/ioutil/ReadFile $pagef]
}

proc view {w r} {
	starter $w $r
	set page [GetQuery $r page]
	if {catch {/io/ioutil/ReadFile "$page.wik"} junk } {
		/fmt/Fprintf $w "Page Not Found: $page"
	} else {
		display $page $w
	}
}

proc list {w r} {
	starter $w $r
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
