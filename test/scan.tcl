must {abc 10 def 100 ghi} [format {abc %d def %d ghi} 10.0 100.0]
must {abc 10 def 100 ghi} [format {abc %g def %g ghi} 10.0 100.0]
must {abc 10.0 def 100.0 ghi} [format {abc %s def %s ghi} 10.0 100.0]
must 2 [scan "abc 20 def 40 ghi" {abc %d def %d ghi} x y]
must 2 [scan "20 40 foo" {%d %d} x y]
must 2 [scan "20 40" {%d %d} x y]
must 1 [catch {scan "20 " {%d %d} x y}]
must 1 [catch {scan "20" {%d %d} x y}]
must 1 [catch {scan " " {%d %d} x y}]
must 1 [catch {scan "" {%d %d} x y}]
