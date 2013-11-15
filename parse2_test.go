package chirp

import (
	"testing"
)

func Test2Dollar1(t *testing.T) {
	expect := "PCmd{ PWord{ BARE{\"one\"} } PWord{ DOLLAR1{\"two\"} } PWord{ BARE{\"three\"} DOLLAR1{\"four\"} } } "

	word, s := Parse2Cmd(" one $two three$four ")
	MustNoSp(expect, word.Show())
	MustA("", s)
}

func Test2Dollar2(t *testing.T) {
	expect := "PCmd{ PWord{ BARE{\"one\"} } PWord{ DOLLAR2{\"two\",PWord{ BARE{\"ab\"} DOLLAR1{\"three\"} } } DOLLAR2{\"xyz\",PWord{ } } } } "

	word, s := Parse2Cmd(" one $two(ab$three)$xyz() ")
	MustNoSp(expect, word.Show())
	MustA("", s)
}

func Test2Curly(t *testing.T) {
	expect := "PWord{ BARE{\" one {two three {}} \123 five\"} }"

	word, s := Parse2Curly("{ one {two three {}} \123 five} six")
	MustNoSp(expect, word.Show())
	MustA(" six", s)
}

func Test2Square(t *testing.T) {
	expect := "SQUARE{ PSeq{ PCmd{ PWord{ BARE{\"one\"} } PWord{ BARE{\"two three {}\"} } PWord{ BARE{\"S\"} } PWord{ BARE{\"five\"} } } }  } "

	word, s := Parse2Square("[ one {two three {}} \123 five] six")
	MustNoSp(expect, word.Show())
	MustA(" six", s)
}

func Test2Cmd(t *testing.T) {
	word, s := Parse2Cmd("one two three")
	MustNoSp("PCmd{ PWord{ BARE{\"one\"} } PWord{ BARE{\"two\"} } PWord{ BARE{\"three\"}} } ", word.Show())
	MustA("", s)

	word, s = Parse2Cmd(" one two three ")
	MustNoSp("PCmd{ PWord{ BARE{\"one\"} } PWord{ BARE{\"two\"} } PWord{ BARE{\"three\"}} } ", word.Show())
	MustA("", s)

	word, s = Parse2Cmd(" one two three ;HEY")
	MustNoSp("PCmd{ PWord{ BARE{\"one\"} } PWord{ BARE{\"two\"} } PWord{ BARE{\"three\"}} } ", word.Show())
	MustA("HEY", s)
}

func Test2Seq(t *testing.T) {
	expect := "PSeq{ PCmd{ PWord{ BARE{\"one\"} } PWord{ BARE{\"two\"} } } PCmd{ PWord{ BARE{\"three\"} } } } "
	word, s := Parse2Seq("one two ; three")
	MustNoSp(expect, word.Show())
	MustA("", s)

	word, s = Parse2Seq(" one two\nthree ]FOO")
	MustNoSp(expect, word.Show())
	MustA("]FOO", s)

	word, s = Parse2Seq(" one two\nthree] FOO")
	MustNoSp(expect, word.Show())
	MustA("] FOO", s)
}
