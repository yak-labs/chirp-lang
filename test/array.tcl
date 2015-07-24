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
