package terp

import (
	"bytes"
	. "fmt"
	"log"
	"strconv"
	"strings"
)

func (fr *Frame) initExpr() {
	Builtins["expr"] = cmdExpr
}

// Concatenate the arguments, adding a space separator, before evaluating the
// expression.
func cmdExpr(fr *Frame, argv []T) T {
	strs := make([]string, len(argv))

	for i, t := range argv[1:] {
		strs[i] = t.String()
	}

	ex := strings.Join(strs, " ")

	return fr.ParseExpression(ex)
}

// Takes a single word that represents an expression and returns the result.
func (fr *Frame) EvalExpr(a T) (result T) {
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

// Takes the string that represents an expression and returns the result.
func (fr *Frame) ParseExpression(s string) (result T) {
	log.Printf("ParseExpression <- %q", s)

	t, _ := fr.ParseExprCond(s)

	return t
}

func (fr *Frame) ParseExprCond(s string) (T, string) {
	log.Printf("ParseExprCond <- %q", s)
	var z T
	z, s = fr.ParseExprDisjunct(s)
	n := len(s)
	i := 0

Loop:
	for i < n {
		c := s[i]

		switch {
		case c == '?':
			i++
			colon := strings.Index(s[i:], ":") // find the ':' character.
			if z.Truth() {
				z = fr.ParseExpression(s[i:colon+1])
				break Loop
			} else {
				z = fr.ParseExpression(s[colon+2:])
				break Loop
			}
		default:
			i++
		}
	}

	return z, s[i:]
}

func (fr *Frame) ParseExprDisjunct(s string) (T, string) {
	log.Printf("ParseExprDisjunct <- %q", s)
	i := 0
	n := len(s)
	var op [2]uint8
	var z T = Empty
	var lookForOp bool = false

Loop:
	for i < n {
		c := s[i]
		p := s[i+1]

		if lookForOp {
			switch {
			case c == '|' && p == '|':
				op = [2]uint8{c, p}
				lookForOp = false
				i += 2
			case White(c):
				i++
			default:
				break Loop
			}
		} else {
			t, rest := fr.ParseExprConjunct(s[i:])
			s = rest
			n = len(s)
			i = 0
			lookForOp = true

			if t == Empty {
				break Loop
			}

			if z == Empty {
				z = t
				if z.Truth() {
					break Loop // shortcircuit
				}
			} else {
				if op == [2]uint8{'|', '|'} {
					if t.Truth() {
						z = MkBool(true)
						break Loop // shortcircuit
					} else {
						z = MkBool(false)
					}
				} else {
					panic("Unexpected operator in ParseExprConjunct.")
				}
			}
		}
	}

	return z, s[i:]
}

func (fr *Frame) ParseExprConjunct(s string) (T, string) {
	log.Printf("ParseExprConjunct <- %q", s)
	i := 0
	n := len(s)
	var op [2]uint8
	var z T = Empty
	var lookForOp bool = false

Loop:
	for i < n {
		c := s[i]
		p := s[i+1]

		if lookForOp {
			switch {
			case c == '&' && p == '&':
				op = [2]uint8{c, p}
				lookForOp = false
				i += 2
			case White(c):
				i++
			default:
				break Loop
			}
		} else {
			t, rest := fr.ParseExprRel(s[i:])
			s = rest
			n = len(s)
			i = 0
			lookForOp = true

			if t == Empty {
				break Loop
			}

			if z == Empty {
				z = t
				if !z.Truth() {
					break Loop // shortcircuit
				}
			} else {
				if op == [2]uint8{'&', '&'} {
					if z.Truth() && t.Truth() {
						z = MkBool(true)
					} else {
						z = MkBool(false)
						break Loop // shortcircuit
					}
				} else {
					panic("Unexpected operator in ParseExprConjunct.")
				}
			}
		}
	}

	return z, s[i:]
}

func (fr *Frame) ParseExprRel(s string) (T, string) {
	log.Printf("ParseExprRel <- %q", s)
	i := 0
	n := len(s)
	var op [2]uint8
	var z T = Empty
	var lookForOp bool = false

Loop:
	for i < n {
		c := s[i]

		if lookForOp == true {
			switch {
			case c == '!', c == '=', c == '<',  c == '>':
				i++
				peek := s[i]
				if peek == '=' {
					op = [2]uint8{c, peek}
					i++
				} else {
					op = [2]uint8{c, 0}
				}
				lookForOp = false
			case White(c):
				i++
			default:
				break Loop
			}
		} else {
			t, rest := fr.ParseExprTerm(s[i:])
			s = rest
			n = len(s)
			i = 0
			lookForOp = true

			if t == Empty {
				break Loop
			}

			if z == Empty {
				z = t
			} else {
				switch op {
				case [2]uint8{'!', '='}:
					z = MkBool(z.Float() != t.Float())

				case [2]uint8{'=', '='}:
					z = MkBool(z.Float() == t.Float())

				case [2]uint8{'<', 0}:
					z = MkBool(z.Float() < t.Float())

				case [2]uint8{'<', '='}:
					z = MkBool(z.Float() <= t.Float())

				case [2]uint8{'>', 0}:
					z = MkBool(z.Float() > t.Float())

				case [2]uint8{'>', '='}:
					z = MkBool(z.Float() >= t.Float())
				}
			}
		}
	}

	return z, s[i:]
}

func (fr *Frame) ParseExprTerm(s string) (T, string) {
	log.Printf("ParseExprTerm <- %q", s)
	i := 0
	n := len(s)
	var op uint8 = 0
	var z T = Empty
	var lookForOp bool = false

Loop:
	for i < n {
		c := s[i]

		if lookForOp == true {
			switch {
			case c == '+', c == '-', c == '|',  c == '^':
				if c == '|' {
					peek := s[i+1]
					if peek == '|' {
						// break if we found "||"
						break Loop
					}
				}
				i++
				op = c
				lookForOp = false
			case White(c):
				i++
			default:
				break Loop
			}
		} else {
			t, rest := fr.ParseExprFactor(s[i:])
			s = rest
			n = len(s)
			i = 0
			lookForOp = true

			if t == Empty {
				break Loop
			}

			if z == Empty {
				z = t
			} else {
				switch op {
				case '+':
					z = MkFloat(z.Float() + t.Float())

				case '-':
					z = MkFloat(z.Float() - t.Float())

				case '|':
					z = z // TODO

				case '^':
					z = z // TODO
				}
			}
		}
	}

	return z, s[i:]
}

func (fr *Frame) ParseExprFactor(s string) (T, string) {
	log.Printf("ParseExprFactor <- %q", s)
	i := 0
	n := len(s)
	var op uint8 = 0
	var z T = Empty
	var lookForOp bool = false

Loop:
	for i < n {
		c := s[i]

		if lookForOp == true {
			switch {
			case c == '*', c == '/', c == '%',  c == '&':
				if c == '&' {
					peek := s[i+1]
					if peek == '&' {
						// break if we found "&&"
						break Loop
					}
				}
				i++
				op = c
				lookForOp = false
			case White(c):
				i++
			default:
				break Loop
			}
		} else {
			t, rest := fr.ParseExprPrim(s[i:])
			s = rest
			n = len(s)
			i = 0
			lookForOp = true

			if t == Empty {
				break Loop
			}

			if z == Empty {
				z = t
			} else {
				switch op {
				case '*':
					z = MkFloat(z.Float() * t.Float())

				case '/':
					z = MkFloat(z.Float() / t.Float())

				case '%':
					z = z // TODO

				case '&':
					z = z // TODO
				}
			}
		}
	}

	return z, s[i:]
}

func (fr *Frame) ParseExprPrim(s string) (T, string) {
	log.Printf("ParseExprPrim <- %q", s)
	i := 0
	n := len(s)

	for i < n {
		c := s[i]

		switch {
		case c == '[':
			return fr.ParseSquare(s[i:])

		case c == '{':
			return fr.ParseCurly(s[i:])

		case c == '$':
			return fr.ParseDollar(s[i:])

		case c == '"':
			return fr.ParseQuote(s[i:])

		case isNumeric(c) || c == '.' && isNumeric(s[i+1]):
			return fr.ParseNumber(s[i:])

		case isOperator(c):
			return Empty, s[i:]
		}

		i++
	}

	panic("ParseExprPrim: No primitive found.")
}

func (fr *Frame) ParseNumber(s string) (T, string) {
	log.Printf("ParseNumber <- %q", s)
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
		return MkFloat(v), s[i:]
	}

	panic(Sprintf("ParseNumber: Could not parse float from: %s", vstr))
}
