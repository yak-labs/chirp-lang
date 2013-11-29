package chirp

import (
	"bytes"
	. "fmt"
	"strings"
)

var _ = strings.Index

func White(ch uint8) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

func WhiteOrSemi(ch uint8) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n' || ch == ';'
}

func (fr *Frame) EvalSeqWithErrorLocation(seq *PSeq) (result T) {
	EvalSeqWithErrorLocationCounter.Incr()
	defer func() {
		if r := recover(); r != nil {
			if re, ok := r.(error); ok {
				r = re.Error() // Convert error to string.
			}
			if rs, ok := r.(string); ok {
				// TODO: Require debug level for the Eval arg.
				src := seq.Src
				if len(src) > 100 {
					src = src[:100] + "..."
				}
				r = rs + Sprintf("\n\tin Eval\n\t\t%q", src)
			}
			panic(r)
		}
	}()

	return seq.Eval(fr)
}

func (fr *Frame) Eval(a T) (result T) {
	return a.EvalSeq(fr)
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

func (fr *Frame) SubstString(s string, flags SubstFlags) string {
	SubstStringCounter.Incr()
	lex := NewLex(s)
	lex.Pos = 0  // Rewind after the builtin Advance()
	lex.Next = 0 // Rewind after the builtin Advance()
	buf := bytes.NewBuffer(nil)
	for lex.Next < lex.Len {
		var c byte = lex.Str[lex.Next]
		if c == '[' && (flags&NoSquare) == 0 {
			lex.Pos = lex.Next // Refocus on the open square.
			lex.Next = lex.Pos + 1
			lex.Tok = '['
			part := Parse2Square(lex)
			MustTok(Token(']'), lex.Tok)
			str := part.Eval(fr).String()
			buf.WriteString(str)
		} else if c == '$' && (flags&NoDollar) == 0 {
			lex.Pos = lex.Next // Refocus on the dollar.
			lex.Next = lex.Pos + 1
			lex.Tok = '$'
			part := Parse2Dollar(lex)
			str := part.Eval(fr).String()
			buf.WriteString(str)
		} else if c == '\\' && (flags&NoBackslash) == 0 {
			c = lex.StretchBackslashEscaped()
			buf.WriteByte(c)
		} else {
			lex.Stretch1()
			buf.WriteByte(c)
		}
	}
	return buf.String()
}

func ParseListOrRecover(s string) (recs []T, err interface{}) {
	defer func() {
		err = recover()
	}()
	recs = ParseList(s)
	return
}

func ParseList(s string) []T {
	ParseListCounter.Incr()
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

var SubstStringCounter Counter
var ParseListCounter Counter
var EvalSeqWithErrorLocationCounter Counter

func init() {
	SubstStringCounter.Register("SubstString")
	ParseListCounter.Register("ParseList")
	EvalSeqWithErrorLocationCounter.Register("EvalSeqWithErrorLocation")
}
