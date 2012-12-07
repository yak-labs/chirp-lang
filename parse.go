package terp

import (
	"bytes"
)

func White(ch uint8) bool {
  return ch==' ' || ch=='\t' || ch=='\r' || ch=='\n'
}

func ParseList(s string) List {
	n := len(s)
	i := 0
	z := make(List, 0, 8)

	for i < n {
		var c uint8

		// skip space
		for i < n {
			c = s[i]
			if !White(s[i]) { break }
			i++
		}
		if i == n { break }

		buf := bytes.NewBuffer(nil)

		// found non-white
		if c == '{' {
			i++
			c = s[i]
			b := 1
			for i < n {
				c = s[i]
				switch c {
				case '{': b++
				case '}': b--
				}
				if b==0 { break }
				buf.WriteByte(c)
				i++
			}
			if c != '}' {
				panic("ParseList: missing end brace:" + Repr(c))
			}
			i++
		} else {
			for i < n {
				c = s[i]
				if White(s[i]) { break }
				buf.WriteByte(c)
				i++
			}
		}
		z = append(z, buf.String())
	}
	return z
}




		
		



