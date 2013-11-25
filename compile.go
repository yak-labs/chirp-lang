package chirp

/*
// To see this simple compiler in action, try
// $ GOPATH=$(pwd) go test -test.v src/terp/*test.go 2>&1 | grep PRECOMP
*/

import (
	. "fmt"
	// "log"
	"regexp"
	"strings"
)

// interfaces

// A Evaler is some piece of code that can be evaluated.
// It may be a single command (Words) or a sequence
// of commands (Seq).
type Evaler interface {
	Eval(fr *Frame) T
}

// A word becomes one argument, or the command word itself.
// Substitutions take place each time it is to be used.
type Word interface {
	Subst(fr *Frame) T
}

// structs

// Seq is a list of Stmts that can be evaluated.
// The value is usually the value of the last one.
// If empty, the value is Empty.
type Seq struct { // implements Evaler
	stmts []Evaler
}

// Words is a single command statement;
// it cannot be an empty list.
// The first word usually names the command.
type Words struct { // * implements Evaler
	words []Word
}

type WordParts struct { // * implements Word // TODO
	parts []Word
}

type BareWord struct { // * implements Word
	word T
}

type Dollar1 struct { // * implements Word
	varName string
}

type Dollar2 struct { // * implements Word // TODO
	varName   string
	subscript Word
}

type Square struct { // * implements Word // TODO
	body Evaler
}

// Eval each member of the sequence, returning the last result.
// If sequence is empty, return Empty.
func (me *Seq) Eval(fr *Frame) T {
	// log.Printf("PRECOMPILED SEQUENCE << %v", me.stmts)
	var z T = Empty
	for _, stmt := range me.stmts {
		z = stmt.Eval(fr)
	}
	// log.Printf("PRECOMPILED SEQUENCE >> %s", Show(z))
	return z
}

// Eval Words by substituting each word, and then Applying the resulting list.
func (me *Words) Eval(fr *Frame) T {
	// log.Printf("PRECOMPILED WORDS << %v", me.words)
	cmd := make([]T, len(me.words))
	for i, w := range me.words {
		cmd[i] = w.Subst(fr)
	}
	z := fr.Apply(cmd)
	// log.Printf("PRECOMPILED WORDS >> %s", Show(z))
	return z
}

// Subst on a BareWord just returns the bare word.
func (me *BareWord) Subst(fr *Frame) T {
	return me.word
}

// Subst on a Dollar1 looks up the variable.
func (me *Dollar1) Subst(fr *Frame) T {
	return fr.GetVar(me.varName)
}

var InertChars = "[!%-/0-9:<-@A-Z^_`a-z|~]"
var BareWordPattern = "(" + InertChars + "+)"

var AlphanumChars = "[A-Za-z0-9_]"
var DumbDollarPattern = "[$](" + AlphanumChars + "+)"

var MatchBareWord = regexp.MustCompile("^" + BareWordPattern + "$")
var MatchDumbDollar = regexp.MustCompile("^" + DumbDollarPattern + "$")

// Very dumb compiler.
// Understands things made only of inert bare words and simple dollar words.
// e.g.   proc doublePlus11 {x} { set a 10 ; + $x $x 1 $a }
//    or  proc f {x} { /fmt/Sprintf %e $x }
// If it finds something it doesn't grok, it panics.
func dumbCompileSequence(fr *Frame, s string) Evaler {
	s = strings.Replace(s, "\t", " ", -1) // get rid of tabs
	s = strings.Replace(s, ";", "\n", -1) // get rid of semicolons

	seq := &Seq{stmts: make([]Evaler, 0, 4)}
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		words := &Words{words: make([]Word, 0, 4)}

		v := strings.Split(line, " ")
		for _, e := range v {
			if e == "" {
				continue
			}
			mw := MatchBareWord.FindStringSubmatch(e)
			if mw != nil {
				words.words = append(words.words, &BareWord{word: MkString(e)})
				continue
			}
			md := MatchDumbDollar.FindStringSubmatch(e)
			if md != nil {
				words.words = append(words.words, &Dollar1{varName: md[1]})
				continue
			}
			panic(Sprintf("Not Handled by CommpileSequence: %q", e))
		}

		if len(words.words) > 0 {
			seq.stmts = append(seq.stmts, words)
		}
	}

	// log.Printf("dumbCompileSequence: SUCCESSFULLY PRECOMPILED: %v <- %q", *seq, s)
	return seq
}

// CompileSequence uses dumbCompileSequence.
// If dumbCompileSequence panics, we recover and return nil.
func CompileSequence(fr *Frame, s string) (z Evaler) {
	defer func() {
		ex := recover() // z stays nil
		if ex != nil {
			Say("CompileSequence BAD", s)
		}
	}()

	// Older Dumber version:
	// return dumbCompileSequence(fr, s)

	lex := NewLex(s)
	z = Parse2Seq(lex)
	if lex.Tok != TokEnd {
		Sayf("CompileSequence Non-Empty rest: %q", s)
		return nil
	}
	Say("CompileSequence GOOD", s, z)
	return
}
