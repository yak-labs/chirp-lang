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
