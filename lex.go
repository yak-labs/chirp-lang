package chirp

import (
	"regexp"
)

type Token uint8

const (
	TokEnd Token = iota // end of string
	TokAlfaNum
	TokNumber
	TokOther // Non-Ascii or not special.

	// " "
	// [ ]
	// { }
	// ( )

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

	// !
	// ~

	// +
	// -
	// *
	// /
	// %

	// &
	// |
	// ^
	TokShiftLeft
	TokShiftRight
)

type Lex struct {
	Str  string
	Len  int   // len(s)
	Next int   // next position to scan
	Pos  int   // position of token
	Type Token // type of token
}

func NewLex(s string) *Lex {
	t := &Lex{Str: s, Len: len(s)}
	t.Advance()
	return t
}

var numberRegexp *regexp.Regexp = regexp.MustCompile("^[-]?[0-9]+[.]?[0-9]*([-+]?[Ee][0-9]+)?")
var alphaNumRegexp *regexp.Regexp = regexp.MustCompile("^[A-Za-z0-9_]+")

func (x *Lex) Current() string {
	return x.Str[x.Pos:x.Next]
}

func (x *Lex) Advance() {
	// Skip over white space.
	// TODO:  are newlines white space?
	for x.Next < x.Len && White(x.Str[x.Next]) {
		x.Next++
	}
	// Set starting position first.
	x.Pos = x.Next
	// Handle End of String.
	if x.Next == x.Len {
		x.Type = TokEnd
		return
	}

	// Get Next two chars in c, d
	var c uint8 = x.Str[x.Next]
	var d uint8 = 0
	if x.Next+1 < x.Len {
		d = x.Str[x.Next+1]
	}

	var bounds []int
	// Say("switch", c)
	switch c {
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		bounds = numberRegexp.FindStringIndex(x.Str[x.Next:])
		// Say("number bounds", bounds)
		if bounds != nil {
			x.Next += bounds[1]
			x.Type = TokNumber
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
		bounds = alphaNumRegexp.FindStringIndex(x.Str[x.Next:])
		// Say("alpha bounds", c)
		if bounds == nil {
			panic("alphaNumRegexp.FindStringIndex bounds cannot be nil.")
		}
		if bounds[0] != 0 {
			panic("alphaNumRegexp.FindStringIndex bounds[0] must be 0.")
		}
		x.Next += bounds[1]
		x.Type = TokAlfaNum
		return
	case '"', '[', ']', '{', '}', '(', ')':
		goto single

	case '+', '*', '/', '%', '^', '~':
		goto single

	case '!':
		if d == '=' {
			x.Type = TokNumNe
			goto pair
		}
		goto single
	case '=':
		if d == '=' {
			x.Type = TokNumEq
			goto pair
		}
		goto other
	case '<':
		if d == '<' {
			x.Type = TokShiftLeft
			goto pair
		}
		if d == '=' {
			x.Type = TokNumLe
			goto pair
		}
		goto single
	case '>':
		if d == '>' {
			x.Type = TokShiftRight
			goto pair
		}
		if d == '=' {
			x.Type = TokNumGe
			goto pair
		}
		goto single
	}
	// Fall through to other.

other: // Anything else goes here, consuming 1 char.
	x.Type = TokOther
	x.Next++
	return

single:
	x.Type = Token(c) // Use actual char value as Token.
	x.Next++
	return

pair: // Consume a pair of chars.
	x.Next += 2
	return
}

func cmdLex(fr *Frame, argv []T) T {
	strT := Arg1(argv)
	str := strT.String()
	x := NewLex(str)
	z := make([]T, 0, 8)
	for x.Type != TokEnd {
		z = append(z, MkUint(uint64(x.Type)))
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
