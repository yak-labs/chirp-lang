
case "$1" in
  on )
	set src/terp/*.go
	for f
	do
	  ex - "$f" <<END
	    1,\$ s;//nolog[.];log.;
	    w
	    q
END
	  ex - "$f" <<END
	    1,\$ s;^	// "log"$;	"log";
	    w
	    q
END
	  done
  ;;
  off )
	set src/terp/*.go
	for f
	do
	  ex - "$f" <<END
	    1,\$ s;\<log[.];//nolog.;
	    w
	    q
END
	  ex - "$f" <<END
	    1,\$ s;^	"log"$;	// "log";
	    w
	    q
END
	  done
  ;;
esac
