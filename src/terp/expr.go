package terp

import (
	"bytes"
	. "fmt"
	"strconv"
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

func isNumeric(ch uint8) bool {
	return '0' <= ch && ch <= '9'
}

func isOperator(ch uint8) bool {
	return ch == '+' || ch == '-' || ch == '~' || ch == '!' || ch == '*' ||
		ch == '/' || ch == '%' || ch == '<' || ch == '>' || ch == '=' ||
		ch == '&' || ch == '|' || ch == '^' || ch == '?' || ch == ':'
}

// Binary operation
type BinOp func(a, b float64) float64

// Binary expression tree node.
type BinExpr struct {
	left   *BinExpr
	right  *BinExpr
	op     BinOp
	value  float64
}

// Initializes a new instance of a binary expression.
func NewBinExpr() *BinExpr {
	return &BinExpr{left: nil, right: nil, op: nil, value: 0.0}
}

// Resolves the value of a binary expression.
func (bex *BinExpr) Value() float64 {
	if bex.left == nil && bex.right == nil && bex.op == nil {
		return bex.value
	}

	if bex.left == nil && bex.right == nil && bex.op == nil {
		panic("BinExpr.Value(): Incomplete binary expression.")
	}

	return bex.op(bex.left.Value(), bex.right.Value())
}

// Takes the string that represents an expression and returns the result.
func (fr *Frame) ParseExpression(s string) (result T) {
	i := 0
	n := len(s)
	result = Empty
	bex := NewBinExpr()

	for i < n {
		c := s[i]

		switch {
		case c == '[':
			sqResult, rest := fr.ParseSquare(s[i:])
			result = sqResult

			s = rest
			n = len(s)
			i = 0

		case c == ']':
			panic("ParseExpression: CloseSquareBracket inside Expression")

		case isNumeric(c) || c == '.' && isNumeric(s[i+1]):
			num, rest := fr.ParseNumber(s[i:])

			if bex.left == nil {
				bex.left = NewBinExpr()
				bex.left.value = num
			}

			if bex.left != nil && bex.op != nil && bex.right == nil {
				bex.right = NewBinExpr()
				bex.right.value = num
			}

			s = rest
			n = len(s)
			i = 0

		case isOperator(c):
			op, rest := fr.ParseOperator(s[i:])

			switch {
			case bex.op == nil:
				bex.op = op

			case bex.op != nil:
				newNode := NewBinExpr()
				newNode.left = bex
				newNode.op = op
				bex = newNode
			}

			s = rest
			n = len(s)
			i = 0

		default:
			i++
		}
	}

	if result == Empty {
		result = MkTf(bex.Value())
	}

	return result
}

func (fr *Frame) ParseNumber(s string) (float64, string) {
	i := 0
	n := len(s)
	decimal := false // only allow one decimal per number
	buf := bytes.NewBuffer(nil)

	for i < n {
		c := s[i]

		switch {
		case isNumeric(c) || c == '.' && decimal == false:
			if c == '.' {
				decimal = true
			}

			buf.WriteByte(c)
			i++

		// An operator or whitespace signifies the end of the number.
		case isOperator(c) || White(c):
			vstr := buf.String()

			if v, ok := strconv.ParseFloat(vstr, 64); ok == nil {
				return v, s[i:]
			}

			panic(Sprintf("ParseNumber: Could not parse float from: %s", vstr))

		default:
			panic(Sprintf("ParseNumber: Unexpected character, '%c', found while parsing number.", c))
		}
	}

	panic("ParseNumber: No number found.")
}

func (fr *Frame) ParseOperator(s string) (BinOp, string) {
	i := 0
	n := len(s)
	buf := bytes.NewBuffer(nil)

	for i < n {
		c := s[i]

		switch {
		case isOperator(c):
			buf.WriteByte(c)
			i++

		// A number or whitespace signifies the end of the operator.
		case isNumeric(c) || White(c):
			opstr := buf.String()

			switch opstr {
			case "+":
				return func(a, b float64) float64 { return a + b }, s[i:]
			case "-":
				return func(a, b float64) float64 { return a - b }, s[i:]
			case "*":
				return func(a, b float64) float64 { return a * b }, s[i:]
			case "/":
				return func(a, b float64) float64 { return a / b }, s[i:]
			}

			panic(Sprintf("ParseOperator: Operator \"%s\" is not recognized.", opstr))

		default:
			panic(Sprintf("ParseOperator: Unexpected character, '%c', found while parsing operator.", c))
		}
	}

	panic("ParseOperator: No operator found.")
}
