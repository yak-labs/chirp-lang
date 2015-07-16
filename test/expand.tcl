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

macro LIST {First Second ARGS} {
  list $First $Second {*}$ARGS
}
proc testLIST {} {
  must [list abc "def ghi" ] [LIST abc "def ghi" ] 
  must [list abc "def ghi" xyz ] [LIST abc "def ghi" xyz ] 
  must [list abc "def ghi" xyz zzz] [LIST abc "def ghi" xyz zzz] 
  must [list abc "def ghi" xyz zzz aaa] [LIST abc "def ghi" xyz zzz aaa] 
}
testLIST
