package chirp

import (
	"testing"
)

func nextMustLex(x *Lex, t Token, s string) {
	MustA(t, x.Tok)
	MustA(s, x.Current())
	x.Advance()
}

func TestLex1(a *testing.T) {
	x := NewLex(` 10 + 20 - 30 * nil <<(4!=3)`)

	MustA(TokNumber, x.Tok)
	MustA(1, x.Pos)
	MustA(3, x.Next)
	x.Advance()

	nextMustLex(x, Token('+'), "+")
	nextMustLex(x, TokNumber, "20")
	nextMustLex(x, Token('-'), "-")
	nextMustLex(x, TokNumber, "30")
	nextMustLex(x, Token('*'), "*")
	nextMustLex(x, TokAlfaNum, "nil")
	nextMustLex(x, TokShiftLeft, "<<")
	nextMustLex(x, Token('('), "(")
	nextMustLex(x, TokNumber, "4")
	nextMustLex(x, TokNumNe, "!=")
	nextMustLex(x, TokNumber, "3")
	nextMustLex(x, Token(')'), ")")
}
