
proc rem {args} {
	return ""
}

set InternalFileName_rx [/regexp/MustCompile {^[a-z][.]([-A-Za-z0-9_.]+)$}]

yproc ListDirs { bundle } {
	foreach f [/io/ioutil/ReadDir "./data/b.$bundle"] {
		set fname [send $f Name]
		set m [send $InternalFileName_rx FindStringSubmatch $fname]
		if {[set m]} {
			yield [lat $m 1]
		}
	}
}

yproc ListFiles { bundle dir } {
	foreach f [/io/ioutil/ReadDir "./data/b.$bundle/d.$dir"] {
		set fname [send $f Name]
		set m [send $InternalFileName_rx FindStringSubmatch $fname]
		if {[set m]} {
			yield [lat $m 1]
		}
	}
}

yproc ListRevs { bundle dir file } {
	foreach f [/io/ioutil/ReadDir "./data/b.$bundle/d.$dir/f.$file"] {
		set fname [send $f Name]
		set m [send $InternalFileName_rx FindStringSubmatch $fname]
		if {[set m]} {
			yield [lat $m 1]
		}
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

proc Route { path query } {
	/fmt/Fprintf $W %s "This is the base Router.  Replace me."
	/fmt/Fprintf $W {path: %s | query: %s} $path $query
}

proc RxCompile { pattern } {
	/regexp/MustCompile $pattern
}

set DB [db-scan ./data]

proc ZygoteHandler {w r} {
	set clone [send $Zygote Clone]
	send $clone Eval [ list set W $w ]
	send $clone Eval [ list set R $r ]
	send $clone Eval [ list Route [getf $r URL Path] [send [getf $r URL] Query] ]
}

set Zygote [interp]
send $Zygote Alias - rem rem
send $Zygote Alias - Route Route
send $Zygote Alias - ListDirs ListDirs
send $Zygote Alias - ListFiles ListFiles
send $Zygote Alias - ListRevs ListRevs
send $Zygote Alias - ReadFile ReadFile
send $Zygote Alias - WriteFile WriteFile
send $Zygote Alias - RxCompile RxCompile
send $Zygote Alias - DB "set DB"

rem -- Load our mixins into our sub-interpreter
set mixins [ListFiles root Mixin]
foreach m $mixins {
	send $Zygote Eval [list mixin $m [ReadFile root Mixin $m]]
}

/net/http/HandleFunc / [http_handler ZygoteHandler]

/net/http/ListenAndServe 127.0.0.1:8080 ""
