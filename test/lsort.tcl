set x  {{one 1} {two 2} {three 3} {ten 10} {hundred 100} {twenty 20}}

must [lsort $x]              {{hundred 100} {one 1} {ten 10} {three 3} {twenty 20} {two 2}}
must [lsort -desc $x]        {{two 2} {twenty 20} {three 3} {ten 10} {one 1} {hundred 100}}
must [lsort -int -ind 1 $x]  {{one 1} {two 2} {three 3} {ten 10} {twenty 20} {hundred 100}}
must [lsort -rea -ind 1 $x]  {{one 1} {two 2} {three 3} {ten 10} {twenty 20} {hundred 100}}

set n [list 384 938 288 3771 38 9948 2 -4 -33 -3 -333 0]

must [lsort $n] {-3 -33 -333 -4 0 2 288 3771 38 384 938 9948}
must [lsort -int $n] {-333 -33 -4 -3 0 2 38 288 384 938 3771 9948}
must [lsort -real $n] {-333 -33 -4 -3 0 2 38 288 384 938 3771 9948}
must [lsort -int -desc $n]  [lreverse {-333 -33 -4 -3 0 2 38 288 384 938 3771 9948}]
must [lsort -real -desc $n] [lreverse {-333 -33 -4 -3 0 2 38 288 384 938 3771 9948}]
