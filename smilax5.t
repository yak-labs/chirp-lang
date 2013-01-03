
proc rem {args} {
	return ""
}

yproc ListDirs { bundle } {
	foreach f [/io/ioutil/ReadDir "./data/b.$bundle"] {
		set fname [send $f Name]
		yield [string-range $fname 2 ""]
	}
}

yproc ListFiles { bundle dir } {
	foreach f [/io/ioutil/ReadDir "./data/b.$bundle/d.$dir"] {
		set fname [send $f Name]
		yield [string-range $fname 2 ""]
	}
}

yproc ListRevs { bundle dir file } {
	foreach f [/io/ioutil/ReadDir "./data/b.$bundle/d.$dir/f.$file"] {
		set fname [send $f Name]
		yield [string-range $fname 2 ""]
	}
}

proc ReadFile { bundle dir file } {
	set revs [concat [ListRevs $bundle $dir $file]]
	set rev [lat $revs 0]

	return [/io/ioutil/ReadFile "./data/b.$bundle/d.$dir/f.$file/r.$rev"]
}

proc WriteFile { bundle dir file contents } {
	/os/MkDirAll "./data/b.$bundle/d.$dir/f.$file" 448
	/io/ioutil/WriteFile "./data/b.$bundle/d.$dir/f.$file/r.0" $contents 384
}

proc Route { w r path query } {
	/fmt/Fprintf $w {path: %s | query: %s} $path $query
}

proc home {w r} {
	set dirs [ListDirs root]
	/fmt/Fprintf $w %s [concat $dirs]
	Route $w $r [getf $r URL Path] [send [getf $r URL] Query]
}

rem -- Load our mixins
set mixins [ListFiles root Mixin]
foreach m $mixins {
	mixin $m [ReadFile root Mixin $m]
}

/net/http/HandleFunc / [http_handler home]

/net/http/ListenAndServe 127.0.0.1:8080 ""
