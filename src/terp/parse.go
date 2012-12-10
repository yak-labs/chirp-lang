package terp

import (
	"bytes"
	. "fmt"
	"log"
)

func White(ch uint8) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

func WhiteNotNewline(ch uint8) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}

func WhiteOrSemi(ch uint8) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n' || ch == ';'
}

func (fr *Frame) Eval(a Any) (result Any) {
	result = "" // In case there are no commands.
	log.Printf("< Eval < %#v\n", a)

	if v, ok := a.(List); ok {
		return fr.Apply(v)
	}

	rest := Str(a)
Loop:
	for {
		var words List
		words, rest = fr.ParseCmd(rest)
		if len(words) == 0 {
			break Loop
		}
		result = fr.Apply(words)
	}
	if len(rest) > 0 {
		panic(Sprintf("Eval: Did not eval entire string: rest=<%q>", rest))
	}
	log.Printf("> Eval > %#v\n", result)
	return
}

// Parse nested curlies, returning contents and new position
func (fr *Frame) ParseCurly(s string) (result Any, rest string) {
	if s[0] != '{' {
		panic("ParseCurly should begin at open curly")
	} // '}'
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
	if c != '}' {
		panic("ParseCurly: missing end brace:" + Repr(c))
	}
	i++

	result = buf.String()
	rest = s[i:]
	return
}

// TODO: ParseSquare is too much like Eval.
// Parse Square Bracketed subcommand, returning result and new position
func (fr *Frame) ParseSquare(s string) (result Any, rest string) {
	log.Printf("< ParseSquare < %#v\n", s)
	result = "" // In case there are no commands.
	if s[0] != '[' {
		panic("ParseSquare should begin at open square")
	}
	rest = s[1:]

Loop:
	for {
		var words List
		words, rest = fr.ParseCmd(rest)
		if len(words) == 0 {
			break Loop
		}
		result = fr.Apply(words)
	}
	if len(rest) == 0 || rest[0] != ']' {
		panic("ParseSquare: missing end bracket")
	}
	rest = rest[1:]
	log.Printf("> ParseSquare > %#v > %q\n", result, rest)
	return
}

func (fr *Frame) ParseQuote(s string) (result Any, rest string) {
	log.Printf("< ParseQuote < %#v\n", s)
	if s[0] != '"' {
		panic("ParseCurly should begin at open curly")
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
			result, rest := fr.ParseSquare(s[i:])
			buf.WriteString(Str(result))
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
	result = buf.String()
	rest = s[i:]
	log.Printf("> ParseQuote > %#v > %q\n", result, rest)
	return
}

// Parse a bareword, returning result and new position
func (fr *Frame) ParseWord(s string) (result Any, rest string) {
	log.Printf("< ParseWord < %#v\n", s)
	i := 0
	n := len(s)
	buf := bytes.NewBuffer(nil)
Loop:
	for i < n {
		c := s[i]
		switch c {
		case '[':
			// Mid-word, squares should return stringlike result.
			result, rest := fr.ParseSquare(s)
			buf.WriteString(Str(result))
			s = rest
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
	result = buf.String()
	rest = s[i:]
	log.Printf("> ParseWord > %#v > %q\n", result, rest)
	return
}

// Might return nonempty <rest> if it finds ']'
// Returns next command as List (may be empty) (substituting as needed) and remaining string.
func (fr *Frame) ParseCmd(str string) (z List, s string) {
	s = str
	log.Printf("< ParseCmd < %#v\n", s)
	z = make(List, 0, 8)
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
		log.Printf("* ParseCmd * TopLoop * z=%#v * s=%q\n", z, s)
		// found non-white
		switch s[0] {
		case ']':
			break Loop
		case '{':
			result, rest := fr.ParseCurly(s)
			z = append(z, result)
			s = rest
		case '[':
			result, rest := fr.ParseSquare(s)
			z = append(z, result)
			s = rest
		case '"':
			result, rest := fr.ParseQuote(s)
			z = append(z, result)
			s = rest
		default:
			result, rest := fr.ParseWord(s)
			z = append(z, result)
			s = rest
		}

		// skip white
		log.Printf("* ParseCmd * skip white * z=%#v * s=%q\n", z, s)
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
		log.Printf("* ParseCmd * End Loop * z=%#v * s=%q\n", z, s)
	} // End Loop
	log.Printf("* ParseCmd * Break Loop * z=%#v * s=%q\n", z, s)

	log.Printf("> ParseCmd > %#v > %q\n", z, s)
	return
}

func (fr *Frame) ParseList(a Any) List {
	if v, ok := a.(List); ok {
		return v
	}
	s := Str(a)
	n := len(s)
	i := 0
	z := make(List, 0, 8)

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
		z = append(z, buf.String())
	}
	return z
}
