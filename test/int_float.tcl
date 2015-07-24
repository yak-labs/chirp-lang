set a 100
set b 3
set aa 100.0
set bb 3.33

must "33" [expr {$a/$b}]
must "33.3333333333333" [expr {$aa/$b}]
must "30.03003003003" [expr {$aa/$bb}]

