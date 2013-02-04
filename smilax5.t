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
	check-site $site
	foreach f [go-call /io/ioutil/ReadDir "data/s.$site"] {
		set fname [go-send $f Name]
		set m [go-send $VolDirRx FindStringSubmatch $fname]
		if {[notnull $m]} {
			yield [lindex $m 1]
		}
	}
}

yproc @ListPages { site vol } {
	check-site $site
	foreach f [go-call /io/ioutil/ReadDir "data/s.$site/v.$vol"] {
		set fname [go-send $f Name]
		set m [go-send $PageDirRx FindStringSubmatch $fname]
		if {[notnull $m]} {
			yield [lindex $m 1]
		}
	}
}

yproc @ListFiles { site vol page } {
	check-site $site
	foreach f [go-call /io/ioutil/ReadDir "data/s.$site/v.$vol/p.$page"] {
		set fname [go-send $f Name]
		set m [go-send $FileDirRx FindStringSubmatch $fname]
		if {[notnull $m]} {
			yield [lindex $m 1]
		}
	}
}

yproc @ListRevs { site vol page file } {
	check-site $site
	foreach f [go-call /io/ioutil/ReadDir "data/s.$site/v.$vol/p.$page/f.$file"] {
		set fname [go-send $f Name]
		set m [go-send $RevFileRx FindStringSubmatch $fname]
		if {[notnull $m]} {
			yield [lindex $m 1]
		}
	}
}

proc @ReadFile { site vol page file } {
	check-site $site
  # TODO: Use [lsort -decreasing] so we can do this in less commands.
	set revs [lsort [concat [@ListRevs $site $vol $page $file]]]
	set rev [lindex $revs [expr [llength $revs] - 1]]

	return [go-call /io/ioutil/ReadFile "data/s.$site/v.$vol/p.$page/f.$file/r.$rev"]
}

proc @WriteFile { site vol page file contents } {
  check-site $site
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
	go-call /fmt/Fprintf [cred w] %s "This is the base Router.  Replace me."
	go-call /fmt/Fprintf [cred w] {path: %s | query: %s} $path $query
}

proc @RxCompile { pattern } {
	go-call /regexp/MustCompile $pattern
}

proc @FindStringSubmatch { rx str } {
	go-send $rx FindStringSubmatch $str
}

proc check-site site {
	set s [cred site]
	# If site empty, then no enforcement yet.
	if [null $s] return
	# If correct site, OK.
	if [eq $s $site] return
	# Must be super-user to access a different site.
	@auth-require-level 90
}

proc @EntityGet {site table id field tag} {
	check-site $site
	entity-get $site $table $id $field $tag
}

proc @EntityPut {site table id field tag values} {
	check-site $site
	entity-put $site $table $id $field $tag $values
}

proc @EntityLike {site table field tag value} {
	check-site $site
	entity-like $site $table $field $tag $value
}

proc @EntityTriples {site table id field tag value} {
	check-site $site
	entity-triples $site $table $id $field $tag $value
}

proc @ShowValue v { 
    go-call /fmt/Sprintf %v $v
}

proc @Puts { str} {
    go-call /fmt/Fprintf [cred w] %s $str
}

proc @ModeHtml {} {
    go-send [go-send [cred w] Header] Set "Content-Type" "text/html"
}

proc @TemporaryRedirect url {
	set rh [go-call /net/http/RedirectHandler $url 307]
	go-send $rh ServeHTTP [cred w] [cred r]
}

proc @auth-require-level {level} {
	if {[cred level] < $level} {
		@RequestBasicAuth
	}
}

proc @RequestBasicAuth {} {
	set h [go-send [cred w] Header]
	go-send $h Set "WWW-Authenticate" "Basic realm=\"[cred site]\""
	go-send [cred w] WriteHeader 401
}

proc @user {} {
	cred user
}
proc @host {} {
	cred host
}
proc @level {} {
	cred level
}
proc @site {} {
	cred site
}
proc @r {} {
	cred r
}
proc @w {} {
	cred w
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
	foreach r [db-select-like [cred site] pw Sys PassWord "$user:$pw" *] {
		return [lindex [go-getf $r Values] 1]
	}
	return 0
}

proc lookup-site {} {
	foreach r [db-select-like root serve Sys ServeSite [cred host] *] {
		return [lindex [go-getf $r Values] 1]
	}
	error "Unknown Site for HOST=[cred host]"
}

# NOW HANDLE REQUESTS
proc ZygoteHandler {w r} {
	#cred-put w $w
	#cred-put r $r
	cred-put site [lookup-site]

	set headers [go-getf $r Header]
	set authorization [go-send $headers Get Authorization]
	if [notnull $authorization] {
		set obfuscated [lindex $authorization 1]
		set decoded [go-send $BASE64 DecodeString $obfuscated]
		set m [go-send $BasicAuthUserPwSplitterRx FindStringSubmatch $decoded]
		if [notnull $m] {
			set _,USER,PASSWORD $m
		}
		set level [gold-level $USER $PASSWORD]
		if {$level <= 0} {
			set level [gold-level * $PASSWORD]
		}
	} else {
		set level [gold-level * *]
	}

	if {$level <= 0} {
	}
	cred-put level $level

	set clone [go-send $Zygote Clone]
	go-send $clone CopyCredFrom -
	go-send $clone Eval [ list Route [go-getf $r URL Path] [go-send [go-getf $r URL] Query] ]
}
go-call /net/http/HandleFunc / [ http-handler-lambda {w r who} {set Who $who; ZygoteHandler $w $r} ] 
go-call /net/http/ListenAndServe 127.0.0.1:8080 ""
