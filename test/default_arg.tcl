proc sum { a {b 10} {c 1}} {
  expr {$a + $b + $c}
}

must 567 [sum 500 60 7]
must 561 [sum 500 60]
must 511 [sum 500]
