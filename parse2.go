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

func (me *PSeq) Eval(fr *Frame) T {
	var z T = Empty
	for _, cmd := range me.Cmds {
		z = cmd.Eval(fr)
	}
	return z
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

func (me *PCmd) Eval(fr *Frame) T {
	words := make([]T, len(me.Words))
	for i, w := range me.Words {
		words[i] = w.Eval(fr)
	}
	return fr.Apply(words)
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
	Multi *terpMulti // If not null, value is fixed and precompiled.
}

func (me *PWord) Eval(fr *Frame) T {
	switch len(me.Parts) {
	case 0:
		return Empty
	case 1:
		if me.Multi != nil {
			//Say("me.Parts[0]: ", me.Parts[0])
			//Say("  me.Multi: ", *me.Multi)
			//if me.Multi.l != nil {
			//	Say("    me.Multi.l", *me.Multi.l)
			//}
			return *me.Multi
		}
		return me.Parts[0].Eval(fr)
	}
	buf := bytes.NewBuffer(nil)
	for _, part := range me.Parts {
		if part.Type == BARE { // Optimization: avoid creating a T.
			buf.WriteString(part.Str)
		} else {
			buf.WriteString(part.Eval(fr).String())
		}
	}
	return MkString(buf.String())
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

func (me *PPart) Eval(fr *Frame) T {
	switch me.Type {
	case BARE:
		return MkString(me.Str)
	case SQUARE:
		return me.Seq.Eval(fr)
	case DOLLAR1:
		return fr.GetVar(me.Str)
	case DOLLAR2:
		v := fr.GetVar(me.Str)
		h := v.Hash()
		if h == nil {
			panic(Sprintf("(*PWord.Eval.DOLLAR2*) Variable %q is not a hash.", me.Str))
		}
		k := me.Seq.Eval(fr).String()
		z := h[k]
		if z == nil {
			panic(Sprintf("(*PWord.Eval.DOLLAR2*) Variable %q: Key not found", me.Str))
		}
		return z
	}
	panic(Sprintf("(*PWord.Eval*) Unknown PartType: %d", me.Type))
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

// Parse nested curlies, returning contents.
func Parse2Curly(lex *Lex) *PWord {
	Say("Enter Parse2Curly", lex)
	if lex.Tok != '{' {
		panic("Parse2Curly should begin at open curly")
	} // vim: '}'

	x := lex.AdvanceCurly()
	// Next is now on the close-curly.
	// Let set Tok to the close-curly:
	lex.Advance()

	multi := MkMulti(x)
	result := &PWord{
		Parts: []*PPart{
			&PPart{
				Str:  x,
				Type: BARE,
			},
		},
		Multi: &multi,
	}
	return result
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
		parts = append(parts, bare)
		Say("finishBarePart +->", parts)
		return parts, bytes.NewBuffer(nil)
	}
	Say("finishBarePart ---", parts)
	return parts, buf
}

// Parse Square Bracketed subcommand, returning result and new position
func Parse2Square(lex *Lex) *PPart {
	Say("Enter Parse2Square", lex)
	if lex.Tok != Token('[') {
		panic("Parse2Square should begin at open square")
	}
	lex.Advance()
	cmds := make([]*PCmd, 0)

Loop:
	for {
		var cmd *PCmd
		cmd = Parse2Cmd(lex)
		if cmd == nil {
			break Loop
		}
		cmds = append(cmds, cmd)
	}

	if lex.Tok != Token(']') {
		panic(Sprintf("Parse2Square: missing end bracket: rest=%q" + lex.Str[lex.Next:]))
	}
	return &PPart{Type: SQUARE, Seq: &PSeq{Cmds: cmds}}
}

func Parse2Quote(lex *Lex) *PWord {
	Say("Enter Parse2Quote", lex)
	if lex.Tok != Token('"') {
		panic("@@@@@ Parse2Quote should begin at open Quote")
	}
	buf := bytes.NewBuffer(nil)
	parts := make([]*PPart, 0)

Loop:
	for lex.Next < lex.Len {
		c := lex.Str[lex.Next]
		Say("Switch Parse2Quote:", string(c))
		switch c {
		case '"':
			// lex.Stretch1()
			lex.Advance() // focus on close-quote.
			break Loop
		case '[':
			parts, buf = finishBarePart(parts, buf)
			// Mid-word, squares should return stringlike result.
			lex.Advance() // to Token('[')
			MustTok(Token('['), lex.Tok)
			r := Parse2Square(lex)
			MustTok(Token(']'), lex.Tok)
			parts = append(parts, r)
		case ']':
			panic("Parse2Quote: CloseSquareBracket inside Quote")
		case '$':
			parts, buf = finishBarePart(parts, buf)
			lex.Advance() // to Token('$')
			MustTok(Token('$'), lex.Tok)
			r := Parse2Dollar(lex)
			parts = append(parts, r)
		case '\\':
			c = lex.StretchBackslashEscaped()
			buf.WriteByte(c)
		default:
			buf.WriteByte(c)
			lex.Stretch1()
		}
	}
	Say("Exit Parse2Quote", lex)
	parts, buf = finishBarePart(parts, buf)
	return &PWord{Parts: parts}
}

// Parse a word, returning result and new position
func Parse2Word(lex *Lex) *PWord {
	Say("Enter Parse2Word", lex)
	buf := bytes.NewBuffer(nil)
	parts := make([]*PPart, 0)

Loop:
	for lex.Tok != TokEnd {
		Say("..... Parse2Word case", lex.Tok)

		// Use the normal Tok lexer to start,
		// but finish with Next pointing to the next unconsumed thing,
		// either more parts to the word, or white space.
		switch lex.Tok {
		case '[':
			parts, buf = finishBarePart(parts, buf)
			// Mid-word, squares should return stringlike result.
			r := Parse2Square(lex)
			parts = append(parts, r)
			MustTok(Token(']'), lex.Tok)
			lex.Stretch1() // Step just past the ]

		case ']':
			// The close-bracket is not part of the word;
			// it terminates an outer seq of cmds.
			// Break with focus still on the bracket.
			break Loop

		case '$':
			parts, buf = finishBarePart(parts, buf)
			r := Parse2Dollar(lex)
			parts = append(parts, r)
			// Next should be just after the whole DOLLAR thing.

		case ' ', '\t', '\n', '\r', ';':
			parts, buf = finishBarePart(parts, buf)
			lex.Advance() // Advance to the next word.
			break Loop

		case '"':
			panic("Parse2Word: DoubleQuote inside word")

		case '\\':
			c := lex.StretchBackslashEscaped()
			buf.WriteByte(c)

		default:
			buf.WriteString(lex.Current())
		}

		// Peek to see if white space is next.
		if lex.Next < lex.Len {
			switch lex.Str[lex.Next] {
			case ' ', '\t', '\n', '\r', ';':
				// If white space, then this Word is over.  Break.
				parts, buf = finishBarePart(parts, buf)
				lex.Advance() // Advance past this white space to next thing before break.
				break Loop
			}
		}

		// Not at white space; there is more.
		// Pos will become Next.  Already made sure not at white space.
		// This will focus on the next thing in the Word.
		lex.Advance() // MAYBE.
	}

	parts, buf = finishBarePart(parts, buf)
	z := &PWord{Parts: parts}
	if len(parts) == 1 && parts[0].Type == BARE {
		multi := MkMulti(parts[0].Str) // Optimize for fixed bare value.
		z.Multi = &multi
	}
	return z
}

// Parse the Key for a Dollar with Parens, e.g. $x(key).
// Dollar, Square, and Backslash substitutions occur.
// White space and DQuotes are not special.
// Terminates with ")".
func Parse2DollarKey(lex *Lex) *PWord {
	Say("Enter Parse2DollarKey", lex)
	buf := bytes.NewBuffer(nil)
	parts := make([]*PPart, 0)

	MustB('(', lex.Str[lex.Next])
	lex.Stretch1()

Loop:
	for lex.Next < lex.Len {
		c := lex.PeekNext()
		switch c {
		case ')':
			break Loop
		case '[':
			parts, buf = finishBarePart(parts, buf)
			// Mid-word, squares should return stringlike result.
			lex.Advance()
			MustTok(Token('['), lex.Tok)
			r := Parse2Square(lex)
			parts = append(parts, r)
			MustTok(Token(']'), lex.Tok)
		case '$':
			parts, buf = finishBarePart(parts, buf)
			lex.Advance()
			MustTok(Token('$'), lex.Tok)
			r := Parse2Dollar(lex)
			parts = append(parts, r)
		case '\\':
			c = lex.StretchBackslashEscaped()
			buf.WriteByte(c)
		default:
			buf.WriteByte(c)
			lex.Stretch1()
		}
	}
	lex.Advance()
	MustTok(Token(')'), lex.Tok)

	parts, buf = finishBarePart(parts, buf)
	return &PWord{Parts: parts}
}

// Parse a variable name after a '$', returning *PPart.
func Parse2Dollar(lex *Lex) *PPart {
	Say("Enter Parse2Dollar", lex)
	if lex.Tok != Token('$') {
		panic("Expected $ at beginning of Parse2Dollar")
	}

	lex.AdvanceIfAlfaNum()
	if lex.Tok != TokAlfaNum {
		panic("Expected a varname after $")
	}
	name := lex.Current()

	// lex.Advance() // DONT Advance-- white matters.

	if lex.PeekNext() == '(' {
		key := Parse2DollarKey(lex)
		MustTok(Token(')'), lex.Tok)
		// DONT lex.Stretch1()
		return &PPart{
			Type: DOLLAR2,
			Str:  name,
			Word: key,
		}
	}
	// else:
	return &PPart{
		Type: DOLLAR1,
		Str:  name,
	}
}

// Returns next command, or else nil.
func Parse2Cmd(lex *Lex) *PCmd {
	Say("Enter Parse2Cmd", lex)
	words := make([]*PWord, 0)

Restart:
	// skip initial newlines and ';'s (as well as white space)
	for lex.Tok == Token(';') || lex.Tok == TokNewline {
		lex.Advance()
	}

Loop:
	// Ways break Loop: TokEnd, TokNewline, Token(';'), Token(']').
	// We leave the lex.Tok at one of those four when we return.
	for lex.Tok != TokEnd {
		switch lex.Tok {
		case TokNewline, Token(';'):
			lex.Advance()
			break Loop
		case Token(']'):
			// DONT lex.Advance() -- leave this at the end-bracket.
			break Loop
		case Token('{'): // vim: '}'
			r := Parse2Curly(lex)
			words = append(words, r)
			// vim: '{'
			MustTok(Token('}'), lex.Tok)
			// TODO: Must Be Followed By White Or End
			lex.Advance()
		case Token('['):
			part := Parse2Square(lex)
			// TODO: Bug if word is [foo][bar]
			words = append(words, &PWord{Parts: []*PPart{part}})
			// vim: '['
			MustTok(Token(']'), lex.Tok)
			// TODO: MIGHT NOT Be Followed By White Or End
			lex.Advance()
		case Token('"'):
			r := Parse2Quote(lex)
			words = append(words, r)
			Say("DID Parse2Quote GOT", r, lex)
			MustTok(Token('"'), lex.Tok)
			lex.Advance()
		case Token('#'):
			if len(words) == 0 {
				lex.SkipComment()
				goto Restart
			}
			fallthrough // # is not special if not at beginning of command.
		default:
			r := Parse2Word(lex)
			words = append(words, r)
		}
	} // End Loop

	if len(words) == 0 {
		return nil
	}
	return &PCmd{Words: words}
}

func Parse2Seq(lex *Lex) *PSeq {
	Say("Enter Parse2Seq", lex)
	z := &PSeq{
		Cmds: make([]*PCmd, 0),
	}
Loop:
	for {
		cmd := Parse2Cmd(lex)
		if cmd == nil {
			break Loop
		}
		z.Cmds = append(z.Cmds, cmd)
		Say("Append Parse2Seq z", z)
	}
	return z
}
