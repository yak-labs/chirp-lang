package terp

import (
	"bytes"
	. "fmt"
	"log"
)

func White(ch uint8) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

func WhiteOrSemi(ch uint8) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n' || ch == ';'
}

func (fr *Frame) TEval(a T) (result T) {
	result = MkTs("")  // In case of empty eval.
	log.Printf("< TEval < (%T) ## %#v ## %q\n", a, a, a.String())

	if v, ok := a.(Tl); ok {
		return fr.TApply(v.l)
	}

	rest := a.String()
Loop:
	for {
		var words []T
		words, rest = fr.ParseCmd(rest)
		if len(words) == 0 {
			break Loop
		}
		result = fr.TApply(words)
	}
	if len(rest) > 0 {
		panic(Sprintf("TEval: Did not eval entire string: rest=<%q>", rest))
	}
	log.Printf("> TEval > (%T) ## %#v ## %q\n", result, result, result.String())
	return
}

// Parse nested curlies, returning contents and new position
func (fr *Frame) ParseCurly(s string) (result T, rest string) {
	if s[0] != '{' {
		panic("ParseCurly should begin at open curly")
	} // vim: '}'
	n := len(s)
	i := 1

	buf := bytes.NewBuffer(nil)
	c := s[i]
	b := 1 // brace depth
Loop:
	for i < n {
		c = s[i]
		switch c {
		case '{':
			b++
		case '}':
			b--
		}
		if b == 0 {
			break Loop
		}
		buf.WriteByte(c)
		i++
	}
	// vim: '{'
	if c != '}' {
		panic("ParseCurly: missing end curly:" + Repr(c))
	}
	i++

	result = MkTs(buf.String())
	rest = s[i:]
	return
}

// TODO: ParseSquare is too much like Eval.
// Parse Square Bracketed subcommand, returning result and new position
func (fr *Frame) ParseSquare(s string) (result T, rest string) {
	//- log.Printf("< ParseSquare < %#v\n", s)
	if s[0] != '[' {
		panic("ParseSquare should begin at open square")
	}
	rest = s[1:]
	result = Empty // In case there are no commands.

Loop:
	for {
		var words []T
		words, rest = fr.ParseCmd(rest)
		if len(words) == 0 {
			break Loop
		}
		result = fr.TApply(words)
	}

	if len(rest) == 0 || rest[0] != ']' {
		panic("ParseSquare: missing end bracket: s=" + Repr(s))
	}
	rest = rest[1:]
	//- log.Printf("> ParseSquare > %#v > %q\n", result, rest)
	return
}

func (fr *Frame) ParseQuote(s string) (T, string) {
	//- log.Printf("< ParseQuote < %#v\n", s)
	if s[0] != '"' {
		panic("ParseQuote should begin at open Quote")
	}
	i := 1
	n := len(s)
	buf := bytes.NewBuffer(nil)
Loop:
	for i < n {
		c := s[i]
		switch c {
		case '[':
			// Mid-word, squares should return stringlike result.
			newresult, rest := fr.ParseSquare(s[i:])
			buf.WriteString(newresult.String())
			s = rest
			n = len(s)
			i = 0
		case ']':
			panic("ParseQuote: CloseSquareBracket inside Quote")
		case '"':
			i++
			break Loop
		default:
			buf.WriteByte(c)
			i++
		}
	}
	//- log.Printf("> ParseQuote > %#v > %q\n", result, rest)
	return MkTs(buf.String()), s[i:]
}

// Parse a bareword, returning result and new position
func (fr *Frame) ParseWord(s string) (T, string) {
	//- log.Printf("< ParseWord < %#v\n", s)
	i := 0
	n := len(s)
	buf := bytes.NewBuffer(nil)
Loop:
	for i < n {
		c := s[i]
		switch c {
		case '[':
			// Mid-word, squares should return stringlike result.
			newresult2, rest2 := fr.ParseSquare(s[i:])
			buf.WriteString(newresult2.String())
			s = rest2
			n = len(s)
			i = 0
		case ']':
			break Loop
		case ' ', '\t', '\n', '\r', ';':
			break Loop
		case '"':
			panic("ParseWord: DoubleQuote inside word")
		default:
			buf.WriteByte(c)
			i++
		}
	}
	// result = MkTs(buf.String())
	// rest = s[i:]
	//- log.Printf("> ParseWord > %#v > %q\n", result, rest)
	return MkTs(buf.String()), s[i:]
}

// Might return nonempty <rest> if it finds ']'
// Returns next command as List (may be empty) (substituting as needed) and remaining string.
func (fr *Frame) ParseCmd(str string) (zcmd []T, s string) {
	s = str
	//- log.Printf("< ParseCmd < %#v\n", s)
	zcmd = make([]T, 0, 4)
	var c uint8

	// skip space or ;
	i := 0
	n := len(s)
	for i < n {
		c = s[i]
		if !WhiteOrSemi(s[i]) {
			break
		}
		i++
	}
	s = s[i:]

Loop:
	for len(s) > 0 {
		//- log.Printf("* ParseCmd * TopLoop * zcmd=%#v * s=%q\n", zcmd, s)
		// found non-white
		switch s[0] {
		case ']':
			break Loop
		case '{':
			newresult, rest := fr.ParseCurly(s)
			zcmd = append(zcmd, newresult)
			s = rest
		case '[':
			newresult, rest := fr.ParseSquare(s)
			zcmd = append(zcmd, newresult)
			s = rest
		case '"':
			newresult, rest := fr.ParseQuote(s)
			zcmd = append(zcmd, newresult)
			s = rest
		default:
			newresult, rest := fr.ParseWord(s)
			zcmd = append(zcmd, newresult)
			s = rest
		}

		// skip white
		//- log.Printf("* ParseCmd * skip white * zcmd=%#v * s=%q\n", zcmd, s)
		n = len(s)
		i = 0
	Skip:
		for i < n {
			switch s[i] {
			case ' ', '\t', '\r':
				i++
				continue Skip
			case ';', '\n':
				break Skip
			default:
				break Skip
			}
		}
		s = s[i:]
		if len(s) == 0 {
			break Loop // end of string
		}
		c = s[0]
		if c == ';' || c == '\n' {
			s = s[1:]  // Omit the semicolon or newline
			break Loop // end of cmd
		}
		//- log.Printf("* ParseCmd * End Loop * zcmd=%#v * s=%q\n", zcmd, s)
	} // End Loop
	//- log.Printf("* ParseCmd * Break Loop * zcmd=%#v * s=%q\n", zcmd, s)

	//- log.Printf("> ParseCmd > %#v > %q\n", zcmd, s)
	return
}

func ParseList(s string) []T {
	n := len(s)
	i := 0
	z := make([]T, 0, 4)

	for i < n {
		var c uint8

		// skip space
		for i < n {
			c = s[i]
			if !White(s[i]) {
				break
			}
			i++
		}
		if i == n {
			break
		}

		buf := bytes.NewBuffer(nil)

		// found non-white
		if c == '{' {
			i++
			c = s[i]
			b := 1
			for i < n {
				c = s[i]
				switch c {
				case '{':
					b++
				case '}':
					b--
				}
				if b == 0 {
					break
				}
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
				if White(s[i]) {
					break
				}
				buf.WriteByte(c)
				i++
			}
		}
		z = append(z, MkTs(buf.String()))
	}
	return z
}
