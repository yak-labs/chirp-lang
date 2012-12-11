puts {package generated}

while {[gets stdin line] >= 0} {
  #if [regexp {^@@ ([a-z0-9/]+) [{}] ([A-Z]+) (\\w+) } $line - pkg verb name] #
  if [regexp {^@@ ([a-z0-9/]+) } $line - pkg ] {
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
    
    set P($pkg) .
  }
}

puts ""
foreach pkg [lsort [array names P]] {
  set alias [string map {/ _} $pkg]
  puts "import $alias \"$pkg\""
}
