proc doImport pkg {
	global Imports
	if {![info exists Imports($pkg)]} {
		regsub -all {/} $pkg {_} tag
		set Imports($pkg) ${tag}
	}
}

proc doLine line {
	global Funcs Vars Types Consts

	# Omit unsafe.
	if {[regexp {^pkg unsafe,} $line]} return

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

	if {[regexp {^pkg ([-A-Za-z0-9_/]+), const ([A-Za-z0-9_]+)} $line _ pkg name]} {
		doImport $pkg
		lappend Consts($pkg) $name
	}
}

proc main argv {
	global Imports Funcs Ins Outs Vars Types Consts
	while {[gets stdin line] >= 0} {
		doLine $line
	}

	puts "package goapi"
	puts ""

	foreach i [lsort [array names Imports]] {
		puts "import $Imports($i) \"$i\""
	}
	puts ""

	puts {var Functions map[string]interface{}}
	puts {var Vars map[string]interface{}}
	puts {var Types map[string]interface{}}
	puts {var Consts map[string]interface{}}

	puts ""
	puts "func init() {"

	puts {	Functions = make(map[string]interface{})}
	puts {	Vars = make(map[string]interface{})}
	puts {	Types = make(map[string]interface{})}
	puts {	Consts = make(map[string]interface{})}

	foreach pkg [lsort [array names Funcs]] {
		set tag $Imports($pkg)
		foreach name [lsort $Funcs($pkg)] {
			puts "\tFunctions\[`/$pkg/$name`\] = $tag.$name"
		}
	}

	foreach pkg [lsort [array names Vars]] {
		set tag $Imports($pkg)
		foreach name [lsort $Vars($pkg)] {
			puts "\tVars\[`/$pkg/$name`\] = &$tag.$name"
		}
	}

	foreach pkg [lsort [array names Types]] {
		set tag $Imports($pkg)
		foreach name [lsort -unique $Types($pkg)] {
			puts "\tTypes\[`/$pkg/$name`\] = new($tag.$name)"
		}
	}

	foreach pkg [lsort [array names Consts]] {
		if { $pkg == "hash/crc64" } {
			# Problems e.g. constant 14514072000185962306 overflows int.
			continue
		}

		set tag $Imports($pkg)
		foreach name [lsort $Consts($pkg)] {
			if { "$pkg/$name" == "math/MaxUint64" } continue
			puts "\tConsts\[`/$pkg/$name`\] = $tag.$name"
		}
	}

	puts "}"
}

main $argv
