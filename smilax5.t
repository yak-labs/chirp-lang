# smilax5.t -- the trusted (unsafe) part of Smilax5, that configs & launches the safe interps.

#### New Storage Hierarchy:
# s: site (old bundle)
# v: volume (old dir)
# p: page (old file)
# f: file.  Special files: "wiki", "src", "db".
# r: revision; also v: varient.

set SiteNameRx     [/regexp/MustCompile {^[a-z][a-z0-9_]*$}]
set VolNameRx      [/regexp/MustCompile {^[A-Z]+[a-z0-9_]*$}]
set PageNameRx     [/regexp/MustCompile {^[A-Z]+[a-z]+[A-Z][A-Za-z0-9_]*$}]
set FileNameRx     [/regexp/MustCompile {^[A-Za-z0-9_.@%~][-A-Za-z0-9_.@%~]*[.][-A-Za-z0-9_.@%~]+$}]

set SiteDirRx     [/regexp/MustCompile {^s[.]([-A-Za-z0-9_.]+)$}]
set VolDirRx      [/regexp/MustCompile {^v[.]([-A-Za-z0-9_.]+)$}]
set PageDirRx     [/regexp/MustCompile {^p[.]([-A-Za-z0-9_.]+)$}]
set FileDirRx     [/regexp/MustCompile {^f[.]([-A-Za-z0-9_.]+)$}]
set RevFileRx     [/regexp/MustCompile {^r[.]([-A-Za-z0-9_.]+)$}]
set VarientFileRx [/regexp/MustCompile {^v[.]([-A-Za-z0-9_.]+)$}]

yproc ListSites { } {
	foreach f [/io/ioutil/ReadDir "data"] {
		set fname [send $f Name]
		set m [send $SiteDirRx FindStringSubmatch $fname]
		if {[notnull $m]} {
			yield [lat $m 1]
		}
	}
}

yproc ListVols { site } {
	foreach f [/io/ioutil/ReadDir "data/s.$site"] {
		set fname [send $f Name]
		set m [send $VolDirRx FindStringSubmatch $fname]
		if {[notnull $m]} {
			yield [lat $m 1]
		}
	}
}

yproc ListPages { site vol } {
	foreach f [/io/ioutil/ReadDir "data/s.$site/v.$vol"] {
		set fname [send $f Name]
		set m [send $PageDirRx FindStringSubmatch $fname]
		if {[notnull $m]} {
			yield [lat $m 1]
		}
	}
}

yproc ListFiles { site vol page } {
	foreach f [/io/ioutil/ReadDir "data/s.$site/v.$vol/p.$page"] {
		set fname [send $f Name]
		set m [send $FileDirRx FindStringSubmatch $fname]
		if {[notnull $m]} {
			yield [lat $m 1]
		}
	}
}

yproc ListRevs { site vol page file } {
	foreach f [/io/ioutil/ReadDir "data/s.$site/v.$vol/p.$page/f.$file"] {
		set fname [send $f Name]
		set m [send $RevFileRx FindStringSubmatch $fname]
		if {[notnull $m]} {
			yield [lat $m 1]
		}
	}
}

proc ReadFile { site vol page file } {
	set revs [concat [ListRevs $site $vol $page $file]]
	set rev [lat $revs 0]

	return [/io/ioutil/ReadFile "data/s.$site/v.$vol/p.$page/f.$file/r.$rev"]
}

proc WriteFile { site vol page file contents } {
	/os/MkDirAll "data/s.$site/v.$vol/p.$page/f.$file" 448
	/io/ioutil/WriteFile "data/s.$site/v.$vol/p.$page/f.$file/r.0" $contents 384
}

proc Route { path query } {
	/fmt/Fprintf $W %s "This is the base Router.  Replace me."
	/fmt/Fprintf $W {path: %s | query: %s} $path $query
}

proc RxCompile { pattern } {
	/regexp/MustCompile $pattern
}

set DB [db-scan data]

proc RequestBasicAuth {realm} {
	set h [send $W Header]
	send $h Set "WWW-Authenticate" "Basic realm=\"$realm\""
	send $W WriteHeader 401
}

proc ZygoteHandler {w r} {
	set clone [send $Zygote Clone]
	send $clone Eval [ list set W $w ]
	send $clone Eval [ list set R $r ]
	send $clone Eval [ list Route [getf $r URL Path] [send [getf $r URL] Query] ]
}

set Zygote [interp]
send $Zygote Alias - Route Route
send $Zygote Alias - ListSites ListSites
send $Zygote Alias - ListVols ListVols
send $Zygote Alias - ListPages ListPages
send $Zygote Alias - ListFiles ListFiles
send $Zygote Alias - ListRevs ListRevs
send $Zygote Alias - ReadFile ReadFile
send $Zygote Alias - WriteFile WriteFile
send $Zygote Alias - RxCompile RxCompile
send $Zygote Alias - DB "set DB"

# -- Load our mixins into our sub-interpreter
set mixins [ListPages root Mixin]
foreach m $mixins {
	send $Zygote Eval [list mixin $m [ReadFile root Mixin $m src]]
}

/net/http/HandleFunc / [http_handler ZygoteHandler]

/net/http/ListenAndServe 127.0.0.1:8080 ""
