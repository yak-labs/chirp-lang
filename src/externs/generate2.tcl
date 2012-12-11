proc doLine line {
  foreach {_ pkg x} $line break 
  puts "// $pkg $x"
  set kind [lindex $x 0]
  set name [lindex $x 1]
  set alias [string map {/ _} $pkg]
  set ::A($alias) $name
  switch $kind {
    TYPE {
      puts "Types\[ `/$pkg/$name` \] = new( *$alias.$name )"
      return $name
    }
    FUNC {
      puts "Funcs\[ `/$pkg/$name` \] = $alias.$name"
      return $name
    }
    VALUE {
      if {[lindex $x 3] eq "!"} return
      puts "Values\[ `/$pkg/$name` \] = &$alias.$name"
      return $name
    }
  }
}

puts {
var Types map[string]interface{} = make(map[string]interface{}, 0)
var Funcs map[string]interface{} = make(map[string]interface{}, 0)
var Values map[string]interface{} = make(map[string]interface{}, 0)
var Members map[string]string = make(map[string]string, 0)
}
puts "func init() \{"

while {[gets stdin line] >= 0} {
  # if [regexp {^@@ [io][os] } $line]
  if [regexp {^@@ ([a-z0-9/]+)} $line - pkg] {

    if [regexp {^builtin$} $pkg] continue
    if [regexp {^syscall$} $pkg] continue
    if [regexp {^runtime/cgo$} $pkg] continue
    if [regexp {zip} $pkg] continue
    if [regexp {archive/tar} $pkg] continue
    if [regexp {.*test$} $pkg] continue
    if [regexp {.*main$} $pkg] continue
    if [regexp {^go/doc/.*} $pkg] continue
    if [regexp {^go/printer/parser} $pkg] continue

    if [lsearch {io bufio fmt net/http} $pkg]<0 continue

    set name [doLine $line]
    if ![string length $name] continue
    if [regexp {.*main$} $pkg] continue
    if [regexp {^go/doc/.*} $pkg] continue
    if [regexp {^go/printer/parser} $pkg] continue

    set name [doLine $line]
    if [string length $name] {
      lappend P($pkg) $name
    }
  }
}

puts ""
foreach pkg [lsort [array names P]] {
  puts "Members\[ `$pkg` \] = `[lsort $P($pkg)]`"
}
puts "\}"
