package terp

import (
	"bytes"
	"log"
)

func (fr *Frame) initTemplateCmds() {
	Builtins["template"] = cmdTemplate
}

func cmdTemplate(fr *Frame, argv []T) T {
	template := Arg1(argv)
	t := template.String()
	buf := new(bytes.Buffer)

	fr.ParseTemplate(buf, t)

	return MkString(buf.String())
}

func (fr *Frame) ParseTemplate(buf *bytes.Buffer, s string) {
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
				// TODO
				//s = fr.ParseTemplateEval()
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

