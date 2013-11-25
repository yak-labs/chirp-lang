package chirp

import (
	"bytes"
	. "fmt"
	"regexp"
)

type Token uint8

const (
	TokEnd Token = iota // end of string
	TokNewline
	TokAlfaNum
	TokNumber
	TokOther // Non-Ascii or not special.

	TokBoolOr  // ||
	TokBoolAnd // &&

	TokNumEq // ==
	TokNumNe
	TokNumLt
	TokNumLe
	TokNumGt
	TokNumGe

	TokStrEq // eq
	TokStrNe
	TokStrLt
	TokStrLe
	TokStrGt
	TokStrGe

	TokShiftLeft  // <<
	TokShiftRight // >>
)

type Lex struct {
	Str  string
	Len  int   // len(s)
	Next int   // next position to scan
	Pos  int   // position of token
	Tok  Token // type of token
}

func (x *Lex) Show() string {
	return Sprintf("<<<%s<<<%s>>>%s>>>", x.Str[:x.Pos], x.Str[x.Pos:x.Next], x.Str[x.Next:])
}

func NewLex(s string) *Lex {
	Say("NewLex <- %v", s)
	t := &Lex{Str: s, Len: len(s)}
	t.Advance()
	return t
}

func MustTok(a, b Token) {
	if a != b {
		Say("MustTok BAD: " + string(a) + " .vs. " + string(b))
		panic(Sprintf("Wrong Token in MustTok: %u vs %u", a, b))
	}
}

var numberRegexp *regexp.Regexp = regexp.MustCompile("^[-]?[0-9]+[.]?[0-9]*([-+]?[Ee][0-9]+)?")
var alfaNumRegexp *regexp.Regexp = regexp.MustCompile("^[A-Za-z0-9_]+")

func (x *Lex) Current() string {
	return x.Str[x.Pos:x.Next]
}

func (x *Lex) SkipComment() {
	Say("Lex SkipComment")
	if x.Tok != Token('#') {
		panic("not # in SkipComment")
	}
	for x.Next < x.Len {
		if x.Next == '\n' {
			break
		}
		x.Next++
	}
	x.Advance()
}

func (x *Lex) PeekNext() byte {
	if x.Next == x.Len {
		return 0
	}
	return x.Str[x.Next]
}

// AdvanceIfAlfaNum will either take a TokAlfaNum (no white space first), or not advance.
func (x *Lex) AdvanceIfAlfaNum() {
	Say("Lex AdvanceIfAlfaNum", x)
	bounds := alfaNumRegexp.FindStringIndex(x.Str[x.Next:])
	if bounds == nil {
		return
	}
	x.Tok = TokAlfaNum
	x.Pos = x.Next
	x.Next += bounds[1]
}

func (x *Lex) Advance() {
	Say("Lex Advance", x)
	// Skip over white space.
	for x.Next < x.Len {
		c := x.Str[x.Next]
		// Return on newline or non-white space.
		// Tcl's definition of White Space includes \v; go's does not.
		if c > ' ' || (c != ' ' && c != '\t' && c != '\r' && c != '\v') {
			break
		}
		x.Next++
	}
	// Set starting position first.
	x.Pos = x.Next

	// Handle End of String.
	if x.Next == x.Len {
		x.Tok = TokEnd
		return
	}

	// Get Next two chars in c, d
	var c uint8 = x.Str[x.Next]
	var d uint8 = 0
	if x.Next+1 < x.Len {
		d = x.Str[x.Next+1]
	}

	var bounds []int
	Say("lex-advance-switch: ", x)
	Say("lex-advance-switch: " + string(c))
	switch c {
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		bounds = numberRegexp.FindStringIndex(x.Str[x.Next:])
		// Say("number bounds", bounds)
		if bounds != nil {
			x.Next += bounds[1]
			x.Tok = TokNumber
			return
		}
		if c == '-' {
			// Say("minus", c)
			goto single
		}
		fallthrough
	case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '_',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
		bounds = alfaNumRegexp.FindStringIndex(x.Str[x.Next:])
		// Say("alfa bounds", c)
		if bounds == nil {
			panic("alfaNumRegexp.FindStringIndex bounds cannot be nil.")
		}
		if bounds[0] != 0 {
			panic("alfaNumRegexp.FindStringIndex bounds[0] must be 0.")
		}
		x.Next += bounds[1]
		x.Tok = TokAlfaNum
		return

	case '!':
		if d == '=' {
			x.Tok = TokNumNe
			goto pair
		}
		goto single
	case '=':
		if d == '=' {
			x.Tok = TokNumEq
			goto pair
		}
		goto other
	case '<':
		if d == '<' {
			x.Tok = TokShiftLeft
			goto pair
		}
		if d == '=' {
			x.Tok = TokNumLe
			goto pair
		}
		goto single
	case '>':
		if d == '>' {
			x.Tok = TokShiftRight
			goto pair
		}
		if d == '=' {
			x.Tok = TokNumGe
			goto pair
		}
		goto single

	case '\n':
		// Newlines are significant for separating statements.
		goto newline

	default:
		if '!' <= c && c <= '~' {
			// Other printable ASCII stuff returns theselves as Tok.
			goto single
		}
		goto other
	}

other: // Nonprintable or nonASCII, consuming 1 char.
	x.Tok = TokOther
	x.Next++
	return

single:
	x.Tok = Token(c) // Use actual char value as Token.
	x.Next++
	return

pair: // Consume a pair of chars.
	x.Next += 2
	return

newline:
	x.Tok = TokNewline
	x.Next++
	return
}

// AdvanceCurly returns the string decoded from the Curly clause and leaves Next on the close-curly.
func (x *Lex) AdvanceCurly() string {
	if x.Tok != Token('{') {
		panic("AdvanceCurly should begin at open curly")
	} // vim: '}'
	// x.Next++

	buf := bytes.NewBuffer(nil)
	b := 1 // brace depth

Loop:
	for x.Next < x.Len {
		c := x.Str[x.Next]
		switch c {
		case '{':
			b++
		case '}':
			b--
		case '\\':
			// In curly braces, only 3 specific things can be backslash-escaped.
			// Followed by anything else, no escaping happens.
			// BUG: if \ and EOS
			switch x.Str[x.Next+1] {
			case '\\':
				c = '\\'
				x.Next++
			case '{':
				c = '{'
				x.Next++
			case '}':
				c = '}'
				x.Next++
			default:
				// Keep that backslash, it's real.
			}
		}
		if b == 0 {
			break Loop
		}
		buf.WriteByte(c)
		x.Next++
	}

	// vim: '{'
	// BUG: if EOS
	if x.Str[x.Next] != '}' {
		panic("AdvanceCurly: missing end curly: " + Repr(x.Str[x.Next]))
	}
	return buf.String()
}

func (lex *Lex) Stretch1() {
	lex.Next++
}

// StretchBackslashEscaped stretches the Next pointer across \C or \ooo for octal.
func (lex *Lex) StretchBackslashEscaped() byte {
	s := lex.Str
	if lex.Next+1 >= lex.Len {
		panic("EOS after escaping backslash")
	}
	switch s[lex.Next+1] {
	case 'a':
		lex.Next += 2
		return '\a'
	case 'b':
		lex.Next += 2
		return '\b'
	case 'f':
		lex.Next += 2
		return '\f'
	case 'n':
		lex.Next += 2
		return '\n'
	case 'r':
		lex.Next += 2
		return '\r'
	case 't':
		lex.Next += 2
		return '\t'
	case 'v':
		lex.Next += 2
		return '\v'
	case 'x':
		panic("Hexadecimal Backslash Escapes not supported (yet)")
	}
	if s[lex.Next+1] < '0' || s[lex.Next+1] > '7' {
		lex.Next += 2
		return s[lex.Next+1] // Default for all other cases is the escaped char.
	}
	if lex.Next+3 >= lex.Len {
		panic("EOS in octal escape")
	}
	if s[lex.Next+2] < '0' || s[lex.Next+2] > '7' {
		panic(Sprintf("Second character after backslash is not octal %q.", s[lex.Next:lex.Next+3]))
	}
	if s[lex.Next+3] < '0' || s[lex.Next+3] > '7' {
		panic(Sprintf("Third character after backslash is not octal %q.", s[lex.Next:lex.Next+4]))
	}
	a := s[lex.Next+1] - '0'
	b := s[lex.Next+2] - '0'
	c := s[lex.Next+3] - '0'
	lex.Next += 4
	return byte(a*64 + b*8 + c)
}

func cmdLex(fr *Frame, argv []T) T {
	strT := Arg1(argv)
	str := strT.String()
	x := NewLex(str)
	z := make([]T, 0, 8)
	for x.Tok != TokEnd {
		z = append(z, MkUint(uint64(x.Tok)))
		z = append(z, MkString(x.Current()))
		x.Advance()
	}
	return MkList(z)
}

func init() {
	if Safes == nil {
		Safes = make(map[string]Command, 333)
	}

	Safes["lex"] = cmdLex
}
