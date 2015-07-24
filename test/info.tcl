set Foo 1
set Bar 2
set Array(color) red
set Array(bg) black

must 1 [expr {[llength [info globals]] > 1}] globals
must 1 [expr {[llength [info locals]] > 1}] locals
must 1 [expr {[llength [info commands]] > 20}] commands

must 0 [info exists Foof] F00f
must 1 [info exists Bar] Bar
must 0 [info exists Array(level)] level
must 1 [info exists Array(color)] color
