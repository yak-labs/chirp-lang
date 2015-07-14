# Expand a subcommand.
set {*}[list foo bar]
must $foo bar

# Expand a variable.
set yum [list ice cream]
set {*}$yum
must $ice cream

# Expand a macro parameter.
macro Set2 Pair {set {*}$Pair}
proc testSet2 {} {
  Set2 "abc xyz"
  must $abc xyz
}
testSet2
