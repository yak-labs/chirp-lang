# This is chirp, not tcl.

proc FetchDir hostport {
  echo hostport=$hostport
  set TCPAddr [/net/ResolveTCPAddr tcp4 $hostport]
  echo TCPAddr=$TCPAddr
  set Conn [/net/DialTCP tcp4 "" $TCPAddr]
  echo Conn=$Conn
  set n [$Conn Write "/\n"]
  echo wrote=$n

  set z ""
  catch {
      while {$n > 0} {
        set bb [[/byte type] mkslice 1000]
        set n [$Conn Read $bb]
        echo n=$n
        set s [[[/string type] convert $bb] reify]
        echo got=[string length $s]
        append z [string range $s 0 [expr {$n - 1}]]
      }
  }
  say Caught

  $Conn Close
  say Closed

  set zz {}
  say zz $zz
  foreach line [split $z "\n"] {
    set line [string trim $line]
    say line $line
    if {[string length $line] > 1} {
      set char [string index $line 0]
      set rest [string range $line 1 end]
      set vec [split $rest "\t"]
      set n [llength $vec]
      say n $n $vec
      if {$n == 4} {
        lappend zz [list $char {*}[lrange $vec 0 3] <->]
        say zz $zz
      }
      if {$n == 5} {
        lappend zz [list $char {*}[lrange $vec 0 4]]
        say zz $zz
      }
    }
  }

  say return $zz
  return $zz
}

set hostport [lindex $Argv 1]
set zz [FetchDir $hostport]
say zz $zz
foreach a,b,c,d,e,f $zz {
  echo "TYPE $a TITLE $b PATH $c HOST $d PORT $e EXTRA $f"
}
