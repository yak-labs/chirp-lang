proc iota n {
  set z {}
  set i 0
  while { $i < $n } {
    lappend z $i
    incr i
  }
  set z
}
  
macro ten {X Y} { iota 10 }

proc Ten {} { ten foo bar }

must [iota 10] [Ten]
