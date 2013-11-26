package chirp

import (
	"bytes"
	. "fmt"
	"strings"
)

func White(ch uint8) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

func WhiteOrSemi(ch uint8) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n' || ch == ';'
}

func (fr *Frame) Eval(a T) (result T) {
	FrameEvalTCounter.Incr()
	defer func() {
		if r := recover(); r != nil {
			if re, ok := r.(error); ok {
				r = re.Error() // Convert error to string.
			}
			if rs, ok := r.(string); ok {
				// TODO: Require debug level for the Eval arg.
				as := a.String()
				if len(as) > 100 {
					as = as[:100] + "..."
				}
				r = rs + Sprintf("\n\tin Eval\n\t\t%q", as)
			}
			panic(r)
		}
	}()

	result = Empty // In case of empty eval.

	if a.IsPreservedByList() {
		FrameEvalTQuickApplyCounter.Incr()
		return fr.Apply(a.List())
	}

	if true {
		// New code, using parse2.

		src := a.String()
		lex := NewLex(src)
		seq := Parse2Seq(lex)
		if lex.Tok != TokEnd {
			Sayf("INCOMPLETE PARSE: Non-Empty rest: <%q>", src)
			return nil
		}

		result = seq.Eval(fr)
		return

	} else {
		// Old code, using parse.
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
			FrameEvalTShortCounter.Incr()
			panic(Sprintf("Eval: Did not eval entire string: rest=<%q>", rest))
		}
		return
	}
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
		case '\\':
			// In curly braces, only 3 specific things can be backslash-escaped.
			// Followed by anything else, no escaping happens.
			switch s[i+1] {
			case '\\':
				c = '\\'
				i++
			case '{':
				c = '{'
				i++
			case '}':
				c = '}'
				i++
			default:
				// Keep that backslash, it's real.
			}
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
	return
}

func (fr *Frame) ParseQuote(s string) (T, string) {
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
	return MkString(buf.String()), s[i:]
}

// Parse a bareword, returning result and new position
func (fr *Frame) ParseWord(s string) (T, string) {
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
	return MkString(buf.String()), s[i:]
}

// Parse the Key for a Dollar with Parens, e.g. $x(key).
// Dollar, Square, and Backslash substitutions occur.
// White space and DQuotes are not special.
// Terminates with ")".
func (fr *Frame) ParseDollarKey(s string) (T, string) {
	i := 0
	n := len(s)
	buf := bytes.NewBuffer(nil)

Loop:
	for i < n {
		c := s[i]
		switch c {
		case ')':
			i++ // Skip over ')'
			break Loop
		case '[':
			// Mid-word, squares should return stringlike result.
			newresult2, rest2 := fr.ParseSquare(s[i:])
			buf.WriteString(newresult2.String())
			s = rest2
			n = len(s)
			i = 0
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
		case '\\':
			c, i = consumeBackslashEscaped(s, i)
			buf.WriteByte(c)
		default:
			buf.WriteByte(c)
			i++
		}
	}
	result := MkString(buf.String())
	rest := s[i:]
	return result, rest
}

// Parse a variable name after a '$', returning result and new position
func (fr *Frame) ParseDollar(s string) (T, string) {
	var key T // nil unless Key exists.
	if '$' != s[0] {
		panic("Expected $ at beginning of ParseDollar")
	}
	i := 1
	n := len(s)
	buf := bytes.NewBuffer(nil)
Loop:
	for i < n {
		c := s[i]
		if c == '(' {
			i++
			key, s = fr.ParseDollarKey(s[i:])
			n = len(s)
			i = 0
			break Loop
		} else if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' || c == '_' {
			buf.WriteByte(c)
			i++
		} else {
			break Loop
		}
	}

	varName := buf.String()
	if len(varName) < 1 {
		panic(Sprintf("Empty Variable Name after $ here: %q", s))
	}
	z := fr.GetVar(varName)
	if key != nil {
		z = z.GetAt(key)
		if z == nil {
			panic(Sprintf("ParseDollar: key not found: varName=%q key=%s", varName, Show(key)))
		}
	}
	return z, s[i:]
}

// Might return nonempty <rest> if it finds ']'
// Returns next command as List (may be empty) (substituting as needed) and remaining string.
func (fr *Frame) ParseCmd(str string) (zwords []T, s string) {
	ParseCmdCounter.Incr()
	s = str
	zwords = make([]T, 0, 4)
	var c uint8

Restart:
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
		case '{': // stupid vim: '}'
			newresult, rest := fr.ParseCurly(s)
			zwords = append(zwords, newresult)
			s = rest
		case '[': // stupid vim: ']'
			newresult, rest := fr.ParseSquare(s)
			zwords = append(zwords, newresult)
			s = rest
		case '"':
			newresult, rest := fr.ParseQuote(s)
			zwords = append(zwords, newresult)
			s = rest
		case '#':
			if len(zwords) == 0 {
				eol := strings.Index(s, "\n")
				if eol >= 0 {
					s = s[eol:]
					goto Restart
				}
				s = ""
				goto Restart
			}
			fallthrough
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
	case 'a':
		return '\a', i + 2
	case 'b':
		return '\b', i + 2
	case 'f':
		return '\f', i + 2
	case 'n':
		return '\n', i + 2
	case 'r':
		return '\r', i + 2
	case 't':
		return '\t', i + 2
	case 'v':
		return '\v', i + 2
	case 'x':
		panic("Hexadecimal Backslash Escapes not supported (yet)")
	}
	if s[i+1] < '0' || s[i+1] > '7' {
		return s[i+1], i + 2 // Default for all other cases is the escaped char.
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
	return byte(a*64 + b*8 + c), i + 4
}

type SubstFlags int

const (
	NoDollar SubstFlags = 1 << iota
	NoSquare
	NoBackslash
)

// SubstString does Square, Dollar, and Backslash substitutions on a string.
func (fr *Frame) SubstString(s string, flags SubstFlags) string {
	i := 0
	n := len(s)
	buf := bytes.NewBuffer(nil)

Loop:
	for i < n {
		c := s[i]
		switch c {
		case '[':
			if (flags & NoSquare) == NoSquare {
				goto Default
			}
			// Mid-word, squares should return stringlike result.
			newresult2, rest2 := fr.ParseSquare(s[i:])
			buf.WriteString(newresult2.String())
			s = rest2
			n = len(s)
			i = 0
			continue Loop
		case ']':
			if (flags & NoSquare) == NoSquare {
				goto Default
			}
			break Loop
		case '$':
			if (flags & NoDollar) == NoDollar {
				goto Default
			}
			newresult3, rest3 := fr.ParseDollar(s[i:])
			buf.WriteString(newresult3.String())
			s = rest3
			n = len(s)
			i = 0
			continue Loop
		case '\\':
			if (flags & NoBackslash) == NoBackslash {
				goto Default
			}
			c, i = consumeBackslashEscaped(s, i)
			buf.WriteByte(c)
			continue Loop
		}
	Default:
		buf.WriteByte(c)
		i++
	}
	if i != n {
		panic(Sprintf("Syntax error in subst: %q", s))
	}
	z := buf.String()
	return z
}

func ParseListOrRecover(s string) (recs []T, err interface{}) {
	defer func() {
		err = recover()
	}()
	recs = ParseList(s)
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

var ParseCmdCounter Counter
var SubstStringCounter Counter
var FrameEvalTCounter Counter
var FrameEvalTQuickApplyCounter Counter
var FrameEvalTShortCounter Counter
var ParseListCounter Counter

func init() {
	ParseCmdCounter.Register("ParseCmd")
	SubstStringCounter.Register("SubstString")
	FrameEvalTCounter.Register("FrameEvalT")
	FrameEvalTQuickApplyCounter.Register("FrameEvalTQuickApply")
	FrameEvalTShortCounter.Register("FrameEvalTShort")
	ParseListCounter.Register("ParseList")
}
