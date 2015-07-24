proc iota n {
  set z {}
  set i 0
  while { $i < $n } {
    lappend z $i
    incr i
  }
  set z
}
  
##################

macro iota_ten {} { iota 10 }

proc Ten {} { iota_ten }
proc ATen {} { return "A[iota_ten]" }
proc BTen {} { return $B(x[iota_ten]y) }


must [iota 10] [Ten]
must A[iota 10] [ATen]

set B(x[iota 10]y) 444
must 444 [BTen]

##################

macro apply1 {X Y} { $X $Y }

proc aTen {} { apply1 iota 10 }
proc aATen {} { return "A[apply1 iota 10]" }
proc aBTen {} { return $B(x[apply1 iota 10]y) }

must [iota 10] [aTen]
must A[iota 10] [aATen]

must 444 [aBTen]

#############################

macro DOUBLET X { list $X $X }

macro STORE X { set Store $X }
macro RECALL {} { set Store }

proc test-foo-foo {} {
  STORE [DOUBLET foo]
  return [RECALL]
}

must "foo foo" [test-foo-foo]

#############################

# K combinater -- returns its first arg.
###macro kc {a ARGS} {eval {set z $a; list {*}$ARGS; set z}}
proc kc {a args} {set a}

proc step {x y} {
  upvar 1 $x r
  upvar 1 $y s
  kc "[set r],[set s]" [incr r] [incr s 10]
}
set a,b {100 200}
step a b
step b a
must 111,211 [step a b]
