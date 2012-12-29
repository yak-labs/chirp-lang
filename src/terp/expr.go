package terp

import (
	"bytes"
	. "fmt"
	"log"
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
	return fr.ParseExpression(a.String())
}

func isNumeric(ch uint8) bool {
	return '0' <= ch && ch <= '9'
}

func isOperator(ch uint8) bool {
	return ch == '+' || ch == '-' || ch == '~' || ch == '!' || ch == '*' ||
		ch == '/' || ch == '%' || ch == '<' || ch == '>' || ch == '=' ||
		ch == '&' || ch == '|' || ch == '^' || ch == '?' || ch == ':'
}

// Floating point binary operation
type FloatBinOp func(a, b float64) float64

// Binary expression tree node.
type BinExpr interface {
	Value() float64
}

// Holds a binary floating point operation.
type FloatOpExpr struct {
	left  BinExpr
	right BinExpr
	op    FloatBinOp
}

// Initializes a new instance of a binary expression.
func NewFloatOpExpr(op FloatBinOp) *FloatOpExpr {
	return &FloatOpExpr{left: nil, right: nil, op: op}
}

// Resolves the value of a binary expression.
func (flop FloatOpExpr) Value() float64 {
	if flop.left == nil || flop.right == nil || flop.op == nil {
		panic("FloatOpExpr.Value(): Incomplete binary expression.")
	}

	return flop.op(flop.left.Value(), flop.right.Value())
}

// Holds a floating point value in the binary expression tree.
type FloatVal struct {
	val float64
}

// Initializes a new instance of a float value.
func NewFloatVal(num float64) *FloatVal {
	return &FloatVal{val: num}
}

// Returns the value held in the node.
func (v FloatVal) Value() float64 {
	return v.val
}

// Takes the string that represents an expression and returns the result.
func (fr *Frame) ParseExpression(s string) (result T) {
	log.Printf("ParseExpression <- %q", s)
	i := 0
	n := len(s)
	result = Empty
	var bex BinExpr = nil

Loop:
	for i < n {
		c := s[i]

		switch {
		case c == '[':
			sqResult, rest := fr.ParseSquare(s[i:])
			s = rest
			n = len(s)
			i = 0

			result = sqResult

			// BEGIN temporary fix -- TODO FIXME
			log.Printf("ParseExpression ==> %s", Show(result))
			return result
			// END temporary fix -- TODO FIXME

		case c == ']':
			panic("ParseExpression: CloseSquareBracket inside Expression")

		case isNumeric(c) || c == '.' && isNumeric(s[i+1]):
			num, rest := fr.ParseNumber(s[i:])
			s = rest
			n = len(s)
			i = 0

			if bex == nil {
				bex = NewFloatVal(num)
				continue Loop
			}

			switch t := bex.(type) {
			case *FloatOpExpr:
				switch {
				case t.left != nil && t.right != nil:
					panic("ParseExpression: No place to put operand.")
				case t.left == nil:
					t.left = NewFloatVal(num)
				case t.right == nil:
					t.right = NewFloatVal(num)
				}

			default:
				panic("ParseExpression: Unexpected expression node type.")
			}

		case isOperator(c):
			op, rest := fr.ParseOperator(s[i:])
			s = rest
			n = len(s)
			i = 0

			if bex == nil {
				panic("ParseExpression: Operator missing an operand.")
			}

			newNode := NewFloatOpExpr(op)
			newNode.left = bex
			bex = newNode

		default:
			i++
		}
	}

	log.Printf("ParseExpression ============== result= %s", Show(result))
	if result.IsEmpty() {
		result = MkTf(bex.Value())
	}

	log.Printf("ParseExpression -> %s", Show(result))
	return result
}

func (fr *Frame) ParseNumber(s string) (float64, string) {
	i := 0
	n := len(s)
	decimal := false // only allow one decimal per number
	buf := bytes.NewBuffer(nil)

Loop:
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
			break Loop

		default:
			panic(Sprintf("ParseNumber: Unexpected character, '%c', found while parsing number.", c))
		}
	}

	vstr := buf.String()

	if v, ok := strconv.ParseFloat(vstr, 64); ok == nil {
		return v, s[i:]
	}

	panic(Sprintf("ParseNumber: Could not parse float from: %s", vstr))
}

func (fr *Frame) ParseOperator(s string) (FloatBinOp, string) {
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
