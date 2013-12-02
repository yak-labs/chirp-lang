# This Tcl script reads in a GO API text file
# and writes a .go file for a package "goapi"
# that, on init, installs all the exported items in the
# RFuncs, RVars, RTypes, & RConsts maps.
#
# Known to work for API Version 1.0: $GOROOT/api/go1.txt
#
# If this links in too much stuff for your tastes,
# you may copy the output .go file, edit it down to what
# you like, and import that instead, from your main package.
# Or leave it out entirely, if you don't care about calling GO.

proc doImport pkg {
	global Imports
	if {![info exists Imports($pkg)]} {
		regsub -all {/} $pkg {_} tag
		set Imports($pkg) ${tag}
	}
}

proc doLine line {
	global Funcs Vars Types Consts ConstTypes

	# Omit unsafe and testing.
	if {[regexp {^pkg unsafe,} $line]} return
	if {[regexp {^pkg testing[/,]} $line]} return

	if {[regexp {^pkg ([-A-Za-z0-9_/]+), func ([A-Za-z0-9_]+)} $line _ pkg name]} {
		doImport $pkg
		lappend Funcs($pkg) $name
	}

	if {[regexp {^pkg ([-A-Za-z0-9_/]+), var ([A-Za-z0-9_]+)} $line _ pkg name]} {
		doImport $pkg
		lappend Vars($pkg) $name
	}

	if {[regexp {^pkg ([-A-Za-z0-9_/]+), type ([A-Za-z0-9_]+) struct} $line _ pkg name]} {
		doImport $pkg
		lappend Types($pkg) $name
	}

	if {[regexp {^pkg ([-A-Za-z0-9_/]+), const ([A-Za-z0-9_]+) ([-A-Za-z0-9_]*)} $line _ pkg name type]} {
		doImport $pkg
		lappend Consts($pkg) $name
		lappend ConstTypes($pkg,$name) $type
	}
}

proc main argv {
	global Imports Funcs Ins Outs Vars Types Consts ConstTypes
	while {[gets stdin line] >= 0} {
		doLine $line
	}

	puts "package goapi"
	puts ""

	puts "import . `github.com/yak-labs/chirp-lang`"
	puts "import ("
	foreach i [lsort [array names Imports]] {
		puts "\t$Imports($i) `$i`"
	}
	puts ")"
	puts ""

	puts "func init() {"

	foreach pkg [lsort [array names Funcs]] {
		set tag $Imports($pkg)
		foreach name [lsort $Funcs($pkg)] {
			puts "\tRFuncs\[`/$pkg/$name`\] = $tag.$name"
		}
	}

	foreach pkg [lsort [array names Vars]] {
		set tag $Imports($pkg)
		foreach name [lsort $Vars($pkg)] {
			puts "\tRVars\[`/$pkg/$name`\] = &$tag.$name"
		}
	}

	foreach pkg [lsort [array names Types]] {
		set tag $Imports($pkg)
		foreach name [lsort -unique $Types($pkg)] {
			puts "\tRTypes\[`/$pkg/$name`\] = new($tag.$name)"
		}
	}

	foreach pkg [lsort [array names Consts]] {

		if { $pkg == "hash/crc64" } {
			# Problems e.g. constant 14514072000185962306 overflows int.
			##continue
		}

		set tag $Imports($pkg)
		foreach name [lsort $Consts($pkg)] {
			set type $ConstTypes($pkg,$name)

			# Heuristic for finding Uint64 consts with their high bit set.
			if { $type == "ideal-int"
					&& ( [string match {*Uint64} $name] || [string match {*64} $pkg] ) } {
				# Hope none of these are negative; in version 1.0 they are not.
				puts "\tRConsts\[`/$pkg/$name`\] = uint64($tag.$name)"
			} elseif { $type == "ideal-int" } {
				# Specify int64 so we can compile when int==int32.
				puts "\tRConsts\[`/$pkg/$name`\] = int64($tag.$name)"
			} else {
				puts "\tRConsts\[`/$pkg/$name`\] = $tag.$name"
			}
		}
	}

	puts "}"
}

main $argv
