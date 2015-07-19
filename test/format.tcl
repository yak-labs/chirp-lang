must foo [format foo]
must foobarbaz [format foo%s%s bar baz]
must "  foo" [format %5s foo]

must "abc 10 def 100 ghi" [format {abc %d def %d ghi} 10.0 100.0]
must "abc 10 def 100 ghi" [format {abc %g def %g ghi} 10.0 100.0]
must "abc 10.0 def 100.0 ghi" [format {abc %s def %s ghi} 10.0 100.0]
