package chirp

import (
	"testing"
)

func testParse2Cmd(input, expect, rest string) {
	lex := NewLex(input)
	cmd := Parse2Cmd(lex)
	MustNoSp(expect, cmd.Show())
	MustA(rest, lex.Str[lex.Pos:])
}

func testParse2Seq(input, expect, rest string) {
	lex := NewLex(input)
	seq := Parse2Seq(lex)
	MustNoSp(expect, seq.Show())
	MustA(rest, lex.Str[lex.Pos:])
}

func testParse2Word(f func(*Lex) *PWord, input, expect, rest string) {
	lex := NewLex(input)
	cmd := f(lex)
	MustNoSp(expect, cmd.Show())
	MustA(rest, lex.Str[lex.Next:])
}

func testParse2Part(f func(*Lex) *PPart, input, expect, rest string) {
	lex := NewLex(input)
	cmd := f(lex)
	MustNoSp(expect, cmd.Show())
	MustA(rest, lex.Str[lex.Pos:])
}

func testParse2Expr(input, expect, rest string) {
	lex := NewLex(input)
	expr := Parse2ExprTop(lex)
	MustNoSp(expect, expr.Show())
	MustA(rest, lex.Str[lex.Pos:])
}

func test2EvalExpr(fr *Frame, input string, expect T) {
	lex := NewLex(input)
	expr := Parse2ExprTop(lex)
	MustTok(TokEnd, lex.Tok)
	z := expr.Eval(fr)
	Must(expect, z)
}

func Test2Dollar1(t *testing.T) {
	expect := "PCmd{ PWord{ BARE{\"one\"} } PWord{ DOLLAR1{\"two\"} } PWord{ BARE{\"three\"} DOLLAR1{\"four\"} } } "
	testParse2Cmd(" one $two three$four ", expect, "")
}

func Test2Dollar2(t *testing.T) {
	expect := "PCmd{ PWord{ BARE{\"one\"} } PWord{ DOLLAR2{\"two\",PWord{ BARE{\"ab\"} DOLLAR1{\"three\"} } } DOLLAR2{\"xyz\",PWord{ } } } } "
	testParse2Cmd(" one $two(ab$three)$xyz() ", expect, "")
}

func Test2Curly(t *testing.T) {
	expect := "PWord{ BARE{\" one {two three {}} \123 five\"} }"
	input := "{ one {two three {}} \123 five} six"
	testParse2Word(Parse2Curly, input, expect, " six")
}

func Test2Square(t *testing.T) {
	expect := "SQUARE{ PSeq{ PCmd{ PWord{ BARE{\"one\"} } PWord{ BARE{\"two three {}\"} } PWord{ BARE{\"S\"} } PWord{ BARE{\"five\"} } } }  } "
	input := "[ one {two three {}} \123 five] six"
	testParse2Part(Parse2Square, input, expect, "] six")
}

func Test2Cmd(t *testing.T) {
	input := "one two three"
	expect := "PCmd{ PWord{ BARE{\"one\"} } PWord{ BARE{\"two\"} } PWord{ BARE{\"three\"}} } "
	testParse2Cmd(input, expect, "")

	input = " one two three "
	expect = "PCmd{ PWord{ BARE{\"one\"} } PWord{ BARE{\"two\"} } PWord{ BARE{\"three\"}} } "
	testParse2Cmd(input, expect, "")

	input = " one two three ;HEY"
	expect = "PCmd{ PWord{ BARE{\"one\"} } PWord{ BARE{\"two\"} } PWord{ BARE{\"three\"}} } "
	testParse2Cmd(input, expect, "HEY")
}

func Test2Seq(t *testing.T) {
	expect := "PSeq{ PCmd{ PWord{ BARE{\"one\"} } PWord{ BARE{\"two\"} } } PCmd{ PWord{ BARE{\"three\"} } } } "
	input := "one two ; three"
	testParse2Seq(input, expect, "")

	input = " one two\nthree ]FOO"
	testParse2Seq(input, expect, "]FOO")

	input = " one two\nthree] FOO"
	testParse2Seq(input, expect, "] FOO")
}

func Test2ExprTop(t *testing.T) {
	expect := "PExpr{ op+ PExpr{ op\" PWord{ BARE{\"38\"} } } PExpr{ op\" PWord{ BARE{\"4\"} } } } "
	input := " 38 + 4 "
	testParse2Expr(input, expect, "")

	expect = "PExpr{op?PExpr{op7PExpr{op\"PWord{DOLLAR1{\"a\"}}}PExpr{op\"PWord{DOLLAR2{\"b\",PWord{BARE{\"z\"}}}}}}PExpr{op+PExpr{op\"PWord{BARE{\"38\"}}}PExpr{op\"PWord{BARE{\"4\"}}}}PExpr{op11PExpr{op\"PWord{BARE{\"abcde\"}}}PExpr{op\"PWord{BARE{\"wxyz\"}}}}}"
	input = " $a == $b(z) ? 38 + 4 : \"abcde\" eq {wxyz} "
	testParse2Expr(input, expect, "")

	fr := NewInterpreter()
	fr.SetVar("x", MkInt(10))
	fr.SetVar("y", MkInt(20))
	fr.SetVar("z", MkInt(30))
	input = " $x < $y ? {YES} : $z "
	test2EvalExpr(fr, input, MkString("YES"))
}
