set x  {{one 1} {two 2} {three 3} {ten 10} {hundred 100} {twenty 20}}

must [lsort $x]              {{hundred 100} {one 1} {ten 10} {three 3} {twenty 20} {two 2}}
must [lsort -desc $x]        {{two 2} {twenty 20} {three 3} {ten 10} {one 1} {hundred 100}}
must [lsort -int -ind 1 $x]  {{one 1} {two 2} {three 3} {ten 10} {twenty 20} {hundred 100}}
must [lsort -rea -ind 1 $x]  {{one 1} {two 2} {three 3} {ten 10} {twenty 20} {hundred 100}}
