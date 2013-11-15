/*
	parse2 will parse to an ADT and support Visitor Pattern on it.
	This will be useful for bytecoes, compilers, tools.
*/

package chirp

import (
	"bytes"
	. "fmt"
	"strings"
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
	Word *PWord // for DOLLAR2
	Seq  *PSeq  // for SQUARE
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
	return Sprintf("PPart{?%d?} ", me.Type)
}

// Parse nested curlies, returning contents and new position
func Parse2Curly(s string) (*PWord, string) {
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

// finishBarePart turns the buffer into a BARE PPart and appends it to the slice, unless the buffer is empty.
// It returns a new empty buffer (or the input buffer, if it was empty).
// This pattern is used by most of the Parsers.
func finishBarePart(parts []*PPart, buf *bytes.Buffer) ([]*PPart, *bytes.Buffer) {
	if buf.Len() > 0 {
		bare := &PPart{
			Type: BARE,
			Str:  buf.String(),
		}
		return append(parts, bare), bytes.NewBuffer(nil)
	}
	return parts, buf
}

// Parse Square Bracketed subcommand, returning result and new position
func Parse2Square(s string) (*PPart, string) {
	if s[0] != '[' {
		panic("Parse2Square should begin at open square")
	}
	rest := s[1:]
	cmds := make([]*PCmd, 0)

Loop:
	for {
		var cmd *PCmd
		cmd, rest = Parse2Cmd(rest)
		if cmd == nil {
			break Loop
		}
		cmds = append(cmds, cmd)
	}

	if len(rest) == 0 || rest[0] != ']' {
		panic("Parse2Square: missing end bracket: s=" + Repr(s))
	}
	rest = rest[1:]
	return &PPart{Type: SQUARE, Seq: &PSeq{Cmds: cmds}}, rest
}

func Parse2Quote(s string) (*PWord, string) {
	if s[0] != '"' {
		panic("Parse2Quote should begin at open Quote")
	}
	i := 1
	n := len(s)
	buf := bytes.NewBuffer(nil)
	parts := make([]*PPart, 0)
Loop:
	for i < n {
		c := s[i]
		switch c {
		case '[':
			parts, buf = finishBarePart(parts, buf)
			// Mid-word, squares should return stringlike result.
			newresult, rest := Parse2Square(s[i:])
			parts = append(parts, newresult)
			s = rest
			n = len(s)
			i = 0
		case ']':
			panic("Parse2Quote: CloseSquareBracket inside Quote")
		case '$':
			parts, buf = finishBarePart(parts, buf)
			newresult3, rest3 := Parse2Dollar(s[i:])
			parts = append(parts, newresult3)
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
	parts, buf = finishBarePart(parts, buf)
	return &PWord{Parts: parts}, s[i:]
}

// Parse a word, returning result and new position
func Parse2Word(s string) (*PWord, string) {
	i := 0
	n := len(s)
	buf := bytes.NewBuffer(nil)
	parts := make([]*PPart, 0)

Loop:
	for i < n {
		c := s[i]
		switch c {
		case '[':
			parts, buf = finishBarePart(parts, buf)
			// Mid-word, squares should return stringlike result.
			newresult2, rest2 := Parse2Square(s[i:])
			parts = append(parts, newresult2)

			s = rest2
			n = len(s)
			i = 0
		case ']':
			break Loop
		case '$':
			parts, buf = finishBarePart(parts, buf)
			newresult3, rest3 := Parse2Dollar(s[i:])
			parts = append(parts, newresult3)

			s = rest3
			n = len(s)
			i = 0
		case ' ', '\t', '\n', '\r', ';':
			parts, buf = finishBarePart(parts, buf)
			break Loop
		case '"':
			panic("Parse2Word: DoubleQuote inside word")
		case '\\':
			c, i = consumeBackslashEscaped(s, i)
			buf.WriteByte(c)
		default:
			buf.WriteByte(c)
			i++
		}
	}
	parts, buf = finishBarePart(parts, buf)
	return &PWord{Parts: parts}, s[i:]
}

// Parse the Key for a Dollar with Parens, e.g. $x(key).
// Dollar, Square, and Backslash substitutions occur.
// White space and DQuotes are not special.
// Terminates with ")".
func Parse2DollarKey(s string) (*PWord, string) {
	i := 0
	n := len(s)
	buf := bytes.NewBuffer(nil)
	parts := make([]*PPart, 0)

Loop:
	for i < n {
		c := s[i]
		switch c {
		case ')':
			i++ // Skip over ')'
			break Loop
		case '[':
			parts, buf = finishBarePart(parts, buf)
			// Mid-word, squares should return stringlike result.
			newresult2, rest2 := Parse2Square(s[i:])
			parts = append(parts, newresult2)
			s = rest2
			n = len(s)
			i = 0
		case '$':
			parts, buf = finishBarePart(parts, buf)
			newresult3, rest3 := Parse2Dollar(s[i:])
			parts = append(parts, newresult3)
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
	parts, buf = finishBarePart(parts, buf)
	return &PWord{Parts: parts}, s[i:]
}

// Parse a variable name after a '$', returning result and new position
func Parse2Dollar(s string) (*PPart, string) {
	var key *PWord // nil unless Key exists.
	if '$' != s[0] {
		panic("Expected $ at beginning of Parse2Dollar")
	}
	i := 1
	n := len(s)
	buf := bytes.NewBuffer(nil)
	typ := DOLLAR1
Loop:
	for i < n {
		c := s[i]
		if c == '(' {
			i++
			key, s = Parse2DollarKey(s[i:])
			n = len(s)
			i = 0
			typ = DOLLAR2
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
		panic(Sprintf("Parse2Dollar: Empty Variable Name after $ here: %q", s))
	}

	z := &PPart{
		Type: typ,
		Str:  varName,
		Word: key,
	}
	return z, s[i:]
}

// Returns next command (or else nil) and remaining string.
func Parse2Cmd(str string) (*PCmd, string) {
	s := str
	words := make([]*PWord, 0)
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
			newresult, rest := Parse2Curly(s)
			words = append(words, newresult)
			s = rest
		case '[': // stupid vim: ']'
			part, rest := Parse2Square(s)
			// TODO: Bug if word is [foo][bar]
			words = append(words, &PWord{Parts: []*PPart{part}})
			s = rest
		case '"':
			newresult, rest := Parse2Quote(s)
			words = append(words, newresult)
			s = rest
		case '#':
			if len(words) == 0 {
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
			newresult, rest := Parse2Word(s)
			words = append(words, newresult)
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
	if len(words) == 0 {
		return nil, s
	}
	return &PCmd{Words: words}, s
}

func Parse2Seq(str string) (*PSeq, string) {
	rest := str
	z := &PSeq{
		Cmds: make([]*PCmd, 0, 4),
	}
Loop:
	for {
		var cmd *PCmd
		cmd, rest = Parse2Cmd(rest)
		if cmd == nil {
			break Loop
		}
		z.Cmds = append(z.Cmds, cmd)
	}
	return z, rest
}
