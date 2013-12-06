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
		if {[regexp {^pkg ([-A-Za-z0-9_/]+),} $line _ pkg]} {
			if {[llength $argv] == 0 || [lsearch -glob $argv $pkg] >= 0} {
				doLine $line
			}
		}
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

	# NEW STYLE
	foreach pkg [lsort [array names Funcs]] {
		set tag $Imports($pkg)
		foreach name [lsort $Funcs($pkg)] {
			puts "\tRoots\[`/$pkg/$name`\] = FuncRoot{ Func: reflect.ValueOf($tag.$name) }"
		}
	}

	foreach pkg [lsort [array names Vars]] {
		set tag $Imports($pkg)
		foreach name [lsort $Vars($pkg)] {
			puts "\tRoots\[`/$pkg/$name`\] = VarRoot{ Var: reflect.ValueOf(&$tag.$name) }"
		}
	}

	foreach pkg [lsort [array names Types]] {
		set tag $Imports($pkg)
		foreach name [lsort $Types($pkg)] {
			# puts "\tRoots\[`/$pkg/$name`\] = TypeRoot{ Type: reflect.ValueOf(new(*$tag.$name)).Type().Elem().Elem() }"
			

			puts "	{"
			puts "	var tmp *$tag.$name"
			puts "\tRoots\[`/$pkg/$name`\] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }"
			puts "	}"
		}
	}

	foreach pkg [lsort [array names Consts]] {
		set tag $Imports($pkg)
		foreach name [lsort $Consts($pkg)] {

			set type $ConstTypes($pkg,$name)

			# Heuristic for finding Uint64 consts with their high bit set.
			if { $type == "ideal-int"
					&& ( [string match {*Uint64} $name] || [string match {*64} $pkg] ) } {
				# Hope none of these are negative; in version 1.0 they are not.
				puts "\tRoots\[`/$pkg/$name`\] = ConstRoot{ Const: uint64($tag.$name) }"
			} elseif { $type == "ideal-int" } {
				# Specify int64 so we can compile when int==int32.
				puts "\tRoots\[`/$pkg/$name`\] = ConstRoot{ Const: int64($tag.$name) }"
			} else {
				puts "\tRoots\[`/$pkg/$name`\] = ConstRoot{ Const: $tag.$name }"
			}
		}
	}

	puts "}"
}

main $argv
