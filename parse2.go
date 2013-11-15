/*
	parse2 will parse to an ADT and support Visitor Pattern on it.
	This will be useful for bytecoes, compilers, tools.
*/

package chirp

import (
	"bytes"
	. "fmt"
	_ "strings"
)

// Any piece of tcl code, a sequence of commands.
type PSeq struct {
	Cmds []*PCmd
}

func (me *PSeq) Show() string {
	z := "PSeq{ "
	for _, e := range me.Cmds {
		z += e.Show()
	}
	z += "} "
	return z
}

// One command made of one or more words.
type PCmd struct {
	Words []*PWord
}

func (me *PCmd) Show() string {
	z := "PCmd{ "
	for _, e := range me.Words {
		z += e.Show()
	}
	z += "} "
	return z
}

// One words, composed of parts that may require substitions.
type PWord struct {
	Parts []*PPart
}

func (me *PWord) Show() string {
	z := "PWord{ "
	for _, e := range me.Parts {
		z += e.Show()
	}
	z += "} "
	return z
}

type PartType int

const (
	BARE    PartType = iota + 1 // Does not need substitions (backslash subs aready done).
	DOLLAR1                     // $x, variable subs without index
	DOLLAR2                     // $x(...), variable subs with index
	SQUARE                      // [...], subcommand eval and replace.
)

type PPart struct {
	Str  string
	Word PWord // for DOLLAR2
	Seq  PSeq  // for SQUARE
	Type PartType
}

func (me *PPart) Show() string {
	switch me.Type {
	case BARE:
		return Sprintf("BARE{%#v} ", me.Str)
	case DOLLAR1:
		return Sprintf("DOLLAR1{%#v} ", me.Str)
	case DOLLAR2:
		return Sprintf("DOLLAR2{%#v,%s} ", me.Str, me.Word.Show())
	case SQUARE:
		return Sprintf("SQUARE{ %s } ", me.Seq.Show())
	}
	return "PPart{?} "
}

// Parse nested curlies, returning contents and new position
func (fr *Frame) Parse2Curly(s string) (*PWord, string) {
	if s[0] != '{' {
		panic("Parse2Curly should begin at open curly")
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
		panic("Parse2Curly: missing end curly:" + Repr(c))
	}
	i++

	result := &PWord{
		Parts: []*PPart{
			&PPart{
				Str:  buf.String(),
				Type: BARE,
			},
		},
	}
	return result, s[i:]
}

/*

// Parse Square Bracketed subcommand, returning result and new position
func (fr *Frame) Parse2Square(s string) (result T, rest string) {
	if s[0] != '[' {
		panic("Parse2Square should begin at open square")
	}
	rest = s[1:]
	result = Empty // In case there are no commands.

Loop:
	for {
		var words []T
		words, rest = fr.Parse2Cmd(rest)
		if len(words) == 0 {
			break Loop
		}
		result = fr.Apply(words)
	}

	if len(rest) == 0 || rest[0] != ']' {
		panic("Parse2Square: missing end bracket: s=" + Repr(s))
	}
	rest = rest[1:]
	return
}

func (fr *Frame) Parse2Quote(s string) (T, string) {
	if s[0] != '"' {
		panic("Parse2Quote should begin at open Quote")
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
			newresult, rest := fr.Parse2Square(s[i:])
			buf.WriteString(newresult.String())
			s = rest
			n = len(s)
			i = 0
		case ']':
			panic("Parse2Quote: CloseSquareBracket inside Quote")
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
			newresult2, rest2 := fr.Parse2Square(s[i:])
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
			newresult2, rest2 := fr.Parse2Square(s[i:])
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
func (fr *Frame) Parse2Cmd(str string) (z *PCmd, s string) {
	s = str
	z = &PCmd{
		words: make([]*PWord, 0, 4),
	}
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
			newresult, rest := fr.Parse2Curly(s)
			z.words = append(z.words, newresult)
			s = rest
		case '[': // stupid vim: ']'
			newresult, rest := fr.Parse2Square(s)
			z.words = append(z.words, newresult)
			s = rest
		case '"':
			newresult, rest := fr.Parse2Quote(s)
			z.words = append(z.words, newresult)
			s = rest
		case '#':
			if len(z.words) == 0 {
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
			z.words = append(z.words, newresult)
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

func (fr *Frame) Parse2Seq(str string) (*PSeq, string) {
	rest := str
	z := &PSeq{
		cmds: make([]*PCmd, 0, 4),
	}
Loop:
	for {
		var cmd *PCmd
		cmd, rest = fr.Parse2Cmd(rest)
		if cmd == nil {
			break Loop
		}
		z.cmds = append(z.cmds, cmd)
	}
	return z, rest
}

*/
