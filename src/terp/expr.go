package terp

import (
	"bytes"
)

func (fr *Frame) initExpr() {
	TBuiltins["expr"] = tcmdExpr
}

func tcmdExpr(fr *Frame, argv []T) T {
	// Just support 1 arg expressions for now.  We'll concat later.
	ex := TArgv1(argv)

	return fr.TEvalExpr(ex)
}

// Takes a single word that represents an expression and returns the result.
func (fr *Frame) TEvalExpr(a T) (result T) {
	if v, ok := a.(Ts); ok {
		return fr.ParseExpression(v.String())
	}

	panic("TEvalExpr: Expected a Ts as its arguement.")
}

// Takes the string that represents an expression and returns the result.
func (fr *Frame) ParseExpression(s string) (result T) {
	i := 0
	n := len(s)
	buff := bytes.NewBuffer(nil)

	for i < n {
		c := s[i]
		switch c {
		case '[':
			sqResult, rest := fr.ParseSquare(s[i:])
			buff.WriteString(sqResult.String())

			s = rest
			n = len(s)
			i = 0
		case ']':
			panic("ParseExpression: CloseSquareBracket inside Expression")
		default:
			i++
		}
	}

	return MkTs(buff.String())
}
