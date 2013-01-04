
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

proc ZygoteHandler {w r} {
	set dirs [ListDirs root]
	interp-eval-in-clone zygote [
		list set W $w
	] [
		list set R $r
	] [
		list Route [getf $r URL Path] [send [getf $r URL] Query]
	]
}

interp-new zygote
interp-alias zygote rem rem
interp-alias zygote Route Route
interp-alias zygote ListDirs ListDirs
interp-alias zygote ListFiles ListFiles
interp-alias zygote ListRevs ListRevs
interp-alias zygote ReadFile ReadFile
interp-alias zygote WriteFile WriteFile

rem -- Load our mixins into our sub-interpreter
set mixins [ListFiles root Mixin]
foreach m $mixins {
	interp-eval zygote [list mixin $m [ReadFile root Mixin $m]]
}

/net/http/HandleFunc / [http_handler ZygoteHandler]

/net/http/ListenAndServe 127.0.0.1:8080 ""
