
old="$1"
new="$2"
shift
shift

for f
do
  ex - "$f" <<END
1,\$ s;\<$old\>;$new;g
w
q
END
done
