proc iota n {
  set z {}
  set i 0
  while { $i < $n } {
    lappend z $i
    incr i
  }
  set z
}
  
macro iota_ten {X Y} { $X $Y }

proc Ten {} { iota_ten iota 10 }
proc ATen {} { return "A[iota_ten iota 10]" }
proc BTen {} { return $B(x[iota_ten iota 10]y) }

must [iota 10] [Ten]
must A[iota 10] [ATen]

set B(x[iota 10]y) 444
must 444 [BTen]
