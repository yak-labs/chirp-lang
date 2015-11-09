set Goose 15
set Duck(name) Donald
set Swan [hash]

must 0 [array exists Bird] Bird
must 0 [array exists Goose] Goose
must 1 [array exists Duck] Duck
must 1 [array exists Swan] Swan

must name [array names Duck]
must "" [array names Swan]

set Duck(huey) H
set Duck(duey) D
set Duck(luey) L

must {duey huey luey name} [array names Duck] {names}
must [list duey D huey H luey L name Donald] $Duck

array set A {2 3 6 7 9 0 22 33 77 88}
echo [array get A]
must 5 [array size A]
must 0 [array size DoesNotExist]
must 88 $A(77)
must [array names A] [list 2 22 6 77 9]
array set B [array get A]
must 5 [array size B]
must [array names B] [list 2 22 6 77 9]
hdel $B 2
must [array names B] [list 22 6 77 9]
must 4 [array size B]
