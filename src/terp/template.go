package terp

import (
	"reflect"
	"bytes"
	"log"
)

func (fr *Frame) initTemplateCmds() {
	Builtins["template"] = cmdTemplate
}

func cmdTemplate(fr *Frame, argv []T) T {
	var buf *bytes.Buffer

	defer func() {
		if r := recover(); r != nil {
			// The buffer doesn't exist already, so create one.
			buf = new(bytes.Buffer)
			fr.SetVar("templbuf", MkValue(reflect.ValueOf(&buf)))
		}
	}()

	// Attempt to use the existing template buffer.
//	buf = fr.GetVar("templbuf").Raw().(*bytes.Buffer)

	buf = new(bytes.Buffer)

	template := Arg1(argv)
	t := template.String()

	fr.ParseTemplate(buf, t)

	return MkString(buf.String())
}

func (fr *Frame) ParseTemplate(buf *bytes.Buffer, s string) {
	log.Printf("ParseTemplate <- %q", s)
	i := 0
	n := len(s)

Loop:
	for i < n {
		c := s[i]

		if c == '{' && i+2 < n {
			peek := s[i+1]

			switch peek {
			case '{':
				s = fr.ParseTemplatePuts(buf, s[i+2:])
				n = len(s)
				i = 0
				continue Loop
			case '%':
				s = fr.ParseTemplateEval(s[i+2:])
				n = len(s)
				i = 0
				continue Loop
			case '#':
				s = fr.ParseTemplateComment(s[i+2:])
				n = len(s)
				i = 0
				continue Loop
			}
		}

		buf.WriteByte(c)
		i++
	}
}

func (fr *Frame) ParseTemplateEval(s string) string {
	log.Printf("ParseTemplateEval <- %q", s)
	i := 0
	n := len(s)
	toEval := new(bytes.Buffer)

Loop:
	for i < n {
		c := s[i]

		switch {
		case c == '%' && i+1 < n:
			peek := s[i+1]

			if peek == '}' {
				// Time to eval.
				e := MkString(toEval.String())
				log.Printf("ParseTemplateEval : Eval <- %q", e.String())
				fr.Eval(e)

				// Return the rest of the template.
				if i+2 < n {
					return s[i+2:]
				} else {
					return ""
				}
			}

		case c == '=' && i+1 < n:
			peek := s[i+1]

			if peek == '}' {
				block, rest := fr.ParseTemplateBlock(s[i+2:])
				toEval.WriteString(" [template {")
				toEval.WriteString(block)
				toEval.WriteString("}] ")
				s = rest
				i = 0
				n = len(s)
				continue Loop
			}
		}

		toEval.WriteByte(c)
		i++
	}

	panic("ParseTemplateEval: Missing eval closer \"%}\".")
}

func (fr *Frame) ParseTemplateBlock(s string) (string, string) {
	log.Printf("ParseTemplateBlock <- %q", s)
	i := 0
	n := len(s)
	blockBuf := new(bytes.Buffer)

	for i < n {
		c := s[i]

		if c == '{' && i+1 < n {
			peek := s[i+1]

			switch peek {
			case '=':
				return blockBuf.String(), s[i+2:]
			case '%':
				return blockBuf.String(), s[i+1:] // keep trailing %
			}
		}

		blockBuf.WriteByte(c)
		i++
	}

	panic("ParseTemplateEval: Missing block closer \"{=\" or \"{%}\".")
}

func (fr *Frame) ParseTemplatePuts(buf *bytes.Buffer, s string) string {
	log.Printf("ParseTemplatePuts <- %q", s)
	i := 0
	n := len(s)
	toEval := new(bytes.Buffer)

	for i < n {
		c := s[i]

		if c == '}' && i+1 < n {
			peek := s[i+1]

			if peek == '}' {
				// Time to eval and put the result into our template output.
				e := MkString(toEval.String())
				r := fr.Eval(e)
				log.Printf("ParseTemplatePuts : Eval <- %q", r.String())
				buf.WriteString(r.String())

				// Return the rest of the template.
				if i+2 < n {
					return s[i+2:]
				} else {
					return ""
				}
			}
		}

		toEval.WriteByte(c)
		i++
	}

	panic("ParseTemplatePuts: Missing puts closer \"}}\".")
}

func (fr *Frame) ParseTemplateComment(s string) string {
	log.Printf("ParseTemplateComment <- %q", s)
	i := 0
	n := len(s)

	for i < n {
		c := s[i]

		if c == '#' && i+1 < n {
			peek := s[i+1]

			if peek == '}' && i+2 < n {
				if i+2 < n {
					return s[i+2:]
				} else {
					return ""
				}
			}
		}

		i++
	}

	panic("ParseTemplateComment: Missing comment closer \"#}\".")
}

