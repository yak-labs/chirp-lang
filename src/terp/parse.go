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

func (fr *Frame) Eval(a T) (result T) {
	result = MkString("") // In case of empty eval.
	log.Printf("< Eval < (%T) ## %#v ## %q\n", a, a, a.String())

	if a.IsPreservedByList() {
		return fr.Apply(a.List())
	}

	rest := a.String()
Loop:
	for {
		var words []T
		words, rest = fr.ParseCmd(rest)
		if len(words) == 0 {
			break Loop
		}
		result = fr.Apply(words)
	}
	if len(rest) > 0 {
		panic(Sprintf("Eval: Did not eval entire string: rest=<%q>", rest))
	}
	log.Printf("> Eval > (%T) ## %#v ## %q\n", result, result, result.String())
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

	result = MkString(buf.String())
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
		result = fr.Apply(words)
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
		case '$':
			newresult3, rest3 := fr.ParseDollar(s[i:])
			buf.WriteString(newresult3.String())
			s = rest3
			n = len(s)
			i = 0
		case '"':
			i++
			break Loop
		case '\\':
			c, i = consumeBackslashEscaped(s, i)
			buf.WriteByte(c)
		default:
			buf.WriteByte(c)
			i++
		}
	}
	//- log.Printf("> ParseQuote > %#v > %q\n", result, rest)
	return MkString(buf.String()), s[i:]
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
		case '$':
			newresult3, rest3 := fr.ParseDollar(s[i:])

			// Special case, the entire word is dollar-substituted. 
			if i == 0 && buf.Len() == 0 && (len(rest3) == 0 || WhiteOrSemi(rest3[0]) || rest3[0] == ']') {
				return newresult3, rest3
			}

			buf.WriteString(newresult3.String())
			s = rest3
			n = len(s)
			i = 0
		case ' ', '\t', '\n', '\r', ';':
			break Loop
		case '"':
			panic("ParseWord: DoubleQuote inside word")
		case '\\':
			c, i = consumeBackslashEscaped(s, i)
			buf.WriteByte(c)
		default:
			buf.WriteByte(c)
			i++
		}
	}
	// result = MkString(buf.String())
	// rest = s[i:]
	//- log.Printf("> ParseWord > %#v > %q\n", result, rest)
	return MkString(buf.String()), s[i:]
}

// Parse a variable name after a '$', returning result and new position
func (fr *Frame) ParseDollar(s string) (T, string) {
	//- log.Printf("< ParseDollar < %#v\n", s)
	if '$' != s[0] {
		panic("Expected $ at beginning of ParseDollar")
	}
	i := 1
	n := len(s)
	buf := bytes.NewBuffer(nil)
Loop:
	for i < n {
		c := s[i]
		if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' || c == '_' {
			buf.WriteByte(c)
		} else {
			break Loop
		}
		i++
	}

	varName := buf.String()
	if len(varName) < 1 {
		panic(Sprintf("Empty Variable Name after $ here: %q", s))
	}
	return fr.GetVar(varName), s[i:]
}

// Might return nonempty <rest> if it finds ']'
// Returns next command as List (may be empty) (substituting as needed) and remaining string.
func (fr *Frame) ParseCmd(str string) (zwords []T, s string) {
	s = str
	//- log.Printf("< ParseCmd < %#v\n", s)
	zwords = make([]T, 0, 4)
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
		// found non-white
		switch s[0] {
		case ']':
			break Loop
		case '{':
			newresult, rest := fr.ParseCurly(s)
			zwords = append(zwords, newresult)
			s = rest
		case '[':
			newresult, rest := fr.ParseSquare(s)
			zwords = append(zwords, newresult)
			s = rest
		case '"':
			newresult, rest := fr.ParseQuote(s)
			zwords = append(zwords, newresult)
			s = rest
		default:
			newresult, rest := fr.ParseWord(s)
			zwords = append(zwords, newresult)
			s = rest
		}

		// skip white
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
	} // End Loop
	return
}

func consumeBackslashEscaped(s string, i int) (byte, int) {
	switch s[i+1] {
		case 'n':
			return '\n', i+2
		case 'r':
			return '\r', i+2
		case 't':
			return '\t', i+2
	}
	if s[i+1] < '0' || s[i+1] > '7' {
		panic(Sprintf("First character after backslash is not octal %q.", s[i:i+2]))
	}
	if s[i+2] < '0' || s[i+2] > '7' {
		panic(Sprintf("Second character after backslash is not octal %q.", s[i:i+3]))
	}
	if s[i+3] < '0' || s[i+3] > '7' {
		panic(Sprintf("Third character after backslash is not octal %q.", s[i:i+4]))
	}
	a := s[i+1] - '0'
	b := s[i+2] - '0'
	c := s[i+3] - '0'
	return byte(a*64 + b*8 + c), i+4
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
				case '\\':
					c, i = consumeBackslashEscaped(s, i)
					i -= 1
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
				if c == '\\' {
					c, i = consumeBackslashEscaped(s, i)
					i -= 1
				}
				buf.WriteByte(c)
				i++
			}
		}
		z = append(z, MkString(buf.String()))
	}
	return z
}
