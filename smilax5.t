# smilax5.t -- the trusted (unsafe) part of Smilax5, that configs & launches the safe interps.

#### New Storage Hierarchy:
# s: site (old bundle)
# v: volume (old dir)
# p: page (old file)
# f: file.  Special files: "@wiki".
# r: revision; also v: varient.

set SiteNameRx     [go-call /regexp/MustCompile {^[a-z][a-z0-9_]*$}]
set VolNameRx      [go-call /regexp/MustCompile {^[A-Z]+[a-z0-9_]*$}]
set PageNameRx     [go-call /regexp/MustCompile {^[A-Z]+[a-z]+[A-Z][A-Za-z0-9_]*$}]
set FileNameRx     [go-call /regexp/MustCompile {^[A-Za-z0-9_.@%~][-A-Za-z0-9_.@%~]*[.][-A-Za-z0-9_.@%~]+$}]

set SiteDirRx     [go-call /regexp/MustCompile {^s[.]([-A-Za-z0-9_.]+)$}]
set VolDirRx      [go-call /regexp/MustCompile {^v[.]([-A-Za-z0-9_.]+)$}]
set PageDirRx     [go-call /regexp/MustCompile {^p[.]([-A-Za-z0-9_.]+)$}]
set FileDirRx     [go-call /regexp/MustCompile {^f[.]([-A-Za-z0-9_.]+)$}]
set RevFileRx     [go-call /regexp/MustCompile {^r[.]([-A-Za-z0-9_.]+)$}]
set VarientFileRx [go-call /regexp/MustCompile {^v[.]([-A-Za-z0-9_.]+)$}]

set MarkForSubinterpRx [go-call /regexp/MustCompile {^[@]([-A-Za-z0-9_]+)$}]
set BasicAuthUserPwSplitterRx [go-call /regexp/MustCompile {^([^:]*)[:](.*)$}]
set BASE64 [go-elem [go-get /encoding/base64/StdEncoding]]

yproc @ListSites {} {
	foreach f [go-call /io/ioutil/ReadDir "data"] {
		set fname [go-send $f Name]
		set m [go-send $SiteDirRx FindStringSubmatch $fname]
		if {[notnull $m]} {
			yield [lindex $m 1]
		}
	}
}

yproc @ListVols { site } {
	foreach f [go-call /io/ioutil/ReadDir "data/s.$site"] {
		set fname [go-send $f Name]
		set m [go-send $VolDirRx FindStringSubmatch $fname]
		if {[notnull $m]} {
			yield [lindex $m 1]
		}
	}
}

yproc @ListPages { site vol } {
	foreach f [go-call /io/ioutil/ReadDir "data/s.$site/v.$vol"] {
		set fname [go-send $f Name]
		set m [go-send $PageDirRx FindStringSubmatch $fname]
		if {[notnull $m]} {
			yield [lindex $m 1]
		}
	}
}

yproc @ListFiles { site vol page } {
	foreach f [go-call /io/ioutil/ReadDir "data/s.$site/v.$vol/p.$page"] {
		set fname [go-send $f Name]
		set m [go-send $FileDirRx FindStringSubmatch $fname]
		if {[notnull $m]} {
			yield [lindex $m 1]
		}
	}
}

yproc @ListRevs { site vol page file } {
	foreach f [go-call /io/ioutil/ReadDir "data/s.$site/v.$vol/p.$page/f.$file"] {
		set fname [go-send $f Name]
		set m [go-send $RevFileRx FindStringSubmatch $fname]
		if {[notnull $m]} {
			yield [lindex $m 1]
		}
	}
}

proc @ReadFile { site vol page file } {
  # TODO: Use [lsort -decreasing] so we can do this in less commands.
	set revs [lsort [concat [@ListRevs $site $vol $page $file]]]
	set rev [lindex $revs [expr [llength $revs] - 1]]

	return [go-call /io/ioutil/ReadFile "data/s.$site/v.$vol/p.$page/f.$file/r.$rev"]
}

proc @WriteFile { site vol page file contents } {
  set now [go-call /time/Now]
  set nowUnix [go-send $now Unix]

  # Need to use strconv, otherwise the int64 gets turned into a float and the
  # timestamp will get represented as scientific notation.
  set timestamp [go-call /strconv/FormatInt $nowUnix 10]

	go-call /os/MkdirAll "data/s.$site/v.$vol/p.$page/f.$file" 448
	go-call /io/ioutil/WriteFile "data/s.$site/v.$vol/p.$page/f.$file/r.$timestamp" $contents 384

	# Save no records, but stupid side-effect is to reread all files.
	db-save-records "data" {}
}

proc @Route { path query } {
	go-call /fmt/Fprintf $W %s "This is the base Router.  Replace me."
	go-call /fmt/Fprintf $W {path: %s | query: %s} $path $query
}

proc @RxCompile { pattern } {
	go-call /regexp/MustCompile $pattern
}

proc @FindStringSubmatch { rx str } {
	go-send $rx FindStringSubmatch $str
}

proc @EntityGet {site table id field tag} {
	entity-get $site $table $id $field $tag
}

proc @EntityPut {site table id field tag values} {
	entity-put $site $table $id $field $tag $values
}

proc @EntityLike {site table field tag value} {
	entity-like $site $table $field $tag $value
}

proc @EntityTriples {site table id field tag value} {
	entity-triples $site $table $id $field $tag $value
}

proc @ShowValue v { 
    go-call /fmt/Sprintf %v $v
}

proc @Puts { str} {
    go-call /fmt/Fprintf $W %s $str
}

proc @ModeHtml {w} {
    go-send [go-send $w Header] Set "Content-Type" "text/html"
}

proc @ParseForm r {
    go-send $r ParseForm
}

proc @GetQuery r {
    go-send [go-getf $r URL] Query
}
proc @GetForm r {
    go-getf $r Form
}

proc @XContent r {
    set content [go-send $r FormValue text]
}
proc @XContent2 r {
    set content2 [go-send [go-send [go-getf $r URL] Query] Get text]
}

proc @UNSAFEDbSelectAll {} {  # TODO remove UNSAFE call.
	db-select-like * * * * * *
}

proc @auth-require-level {level} {
}

proc @RequestBasicAuth {w realm} {
	set h [go-send $w Header]
	go-send $h Set "WWW-Authenticate" "Basic realm=\"$realm\""
	go-send $w WriteHeader 401
}

proc @user {} {
	return $USER
}
proc @host {} {
	return $HOST
}
proc @level {} {
	return $LEVEL
}
proc @site {} {
	return $SITE
}

# Dir name is "data"
db-rebuild "data"

######  DEFINE @-procs ABOVE.

set Zygote [interp]

foreach cmd [info commands] {
  set m [go-send $MarkForSubinterpRx FindStringSubmatch $cmd]
  if [notnull $m] {
    go-send $Zygote Alias - [lindex $m 1] $cmd
  }
}

go-send $Zygote Alias - DB "set DB"

# -- Load our mixins into our sub-interpreter
set mixins [@ListPages root Mixin]
foreach m $mixins {
	go-send $Zygote Eval [list mixin $m [@ReadFile root Mixin $m @wiki]]
}

proc gold-level {user pw} {
	foreach r [db-select-like $SITE pw Sys PassWord "$user:$pw" *] {
		return [lindex [go-getf $r Values] 1]
	}
	return 0
}

proc lookup-site {} {
	foreach r [db-select-like root serve Sys ServeSite $HOST *] {
		return [lindex [go-getf $r Values] 1]
	}
	error "Unknown Site for HOST=$HOST"
}

# NOW HANDLE REQUESTS
proc ZygoteHandler {w r} {
	set clone [go-send $Zygote Clone]

	set HOST [go-getf $r Host]
	set SITE [lookup-site]

	set headers [go-getf $r Header]
	set authorization [go-send $headers Get Authorization]
	if [notnull $authorization] {
		set obfuscated [lindex $authorization 1]
		set decoded [go-send $BASE64 DecodeString $obfuscated]
		set m [go-send $BasicAuthUserPwSplitterRx FindStringSubmatch $decoded]
		if [notnull $m] {
			set _,USER,PASSWORD $m
		}
		set LEVEL [gold-level $USER $PASSWORD]
		if {$LEVEL <= 0} {
			set LEVEL [gold-level * $PASSWORD]
		}
	} else {
		# @RequestBasicAuth $w 5.SMILAX.ORG
		set LEVEL [gold-level * *]
	}

	set W $w
	set R $r
	go-send $clone Eval [ list set W $w ]
	go-send $clone Eval [ list set R $r ]
	go-send $clone Eval [ list Route [go-getf $r URL Path] [go-send [go-getf $r URL] Query] ]
}
go-call /net/http/HandleFunc / [http_handler ZygoteHandler]
go-call /net/http/ListenAndServe 127.0.0.1:8080 ""
