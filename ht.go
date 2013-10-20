package chirp

import (
	"bytes"
	. "fmt"
	"html"
	R "reflect"
	"regexp"
)

var esc = html.EscapeString

// Nice things should start with letter or "_",
// and may also contain digits, ":", and "-".
var nice_rx = regexp.MustCompile("^[A-Za-z_][-A-Za-z0-9_:]*$")

// nice asserts that tags and attributes are "nice", conforming to nice_rx.
// returning its argument, or panics if bad.
func nice(s string) string {
	if nice_rx.FindString(s) == s {
		return s
	}
	panic(Sprintf("ht or ht-tag: Bad identifier: %q", s))
}

var nice_entity_rx = regexp.MustCompile("^#?[A-Za-z0-9]+$")

// nice asserts that tags and attributes are "nice", conforming to nice_rx.
// returning its argument, or panics if bad.
func niceEntity(s string) string {
	if nice_entity_rx.FindString(s) == s {
		return s
	}
	panic(Sprintf("ht-entity: Bad entity syntax: %q", s))
}

// HTML contains properly formatted & escaped HTML
type HTML string

// In script:  puts [go-send $ht Html]
func (p HTML) Html() string {
	return string(p)
}

// HtRaw trusts its input string is proper HTML, and casts it to HTML.
func HtRaw(s string) HTML {
	return HTML(s)
}

// Ht concats its arguments, converting each to HTML if it is not already HTML.
func Ht(things []T) HTML {
	buf := bytes.NewBuffer(make([]byte, 0, 40))
	for _, thing := range things {
		switch a := thing.Raw().(type) {
		case HTML:
			buf.WriteString(string(a))
		case string:
			buf.WriteString(esc(a))
		case Stringer:
			buf.WriteString(esc(a.String()))
		case error:
			buf.WriteString(esc(a.Error()))
		case Hash:
			buf.WriteString(esc(thing.String()))
		default:
			buf.WriteString(esc(Sprintf("%v", a)))
		}
	}
	return HTML(buf.String())
}

// Tag takes a HTML tag name, a slice of attr keys and values (stride is 2), and a body for the tag.
func Tag(tag T, attrs []T, body T) HTML {
	tg := nice(tag.String())
	buf := bytes.NewBuffer(nil)
	buf.WriteString(`<`)
	buf.WriteString(tg)
	for i := 0; i+1 < len(attrs); i += 2 {
		k := nice(attrs[i].String())
		v := esc(attrs[i+1].String())
		buf.WriteString(Sprintf(" %s=\"%s\"", k, v))
	}
	var x string
	if ht, ok := body.Raw().(HTML); ok {
		// Body is already a HTML; use it raw.
		x = string(ht)
	} else {
		// Must escape body.
		x = esc(body.String())
	}
	if len(x) == 0 {
		buf.WriteString("/\n>")
	} else {
		buf.WriteString(Sprintf(">%s</%s\n>", x, esc(tg)))
	}
	return HTML(buf.String())
}

func cmdHt(fr *Frame, argv []T) T {
	args := Arg0v(argv)
	return MkValue(R.ValueOf(Ht(args)))
}

func cmdHtRaw(fr *Frame, argv []T) T {
	html := Arg1(argv)
	return MkValue(R.ValueOf(HtRaw(html.String())))
}

func HtEntity(entity string) T {
    s := "&" + niceEntity(entity) + ";"
	return MkValue(R.ValueOf(HtRaw(s)))
}
func cmdHtEntity(fr *Frame, argv []T) T {
	entity := Arg1(argv)
	return HtEntity(entity.String())
}

func HtPeek(ht HTML) string {
    return string(ht)
}
func cmdHtPeek(fr *Frame, argv []T) T {
	h := Arg1(argv)
	ht, ok := h.Raw().(HTML)
	if (!ok) {
	  panic("ht-peek expects an ht object argument")
	}
	return MkString(HtPeek(ht))
}

func cmdHtTag(fr *Frame, argv []T) T {
	tag, args := Arg1v(argv)
	n := len(args)
	var body T = Empty

	if n%2 == 1 {
		// If odd arg, remove last arg to be body.
		body = args[n-1]
		args = args[:n-1]
	}
	return MkValue(R.ValueOf(Tag(tag, args, body)))
}

func init() {
	if Safes == nil {
		Safes = make(map[string]Command, 333)
	}

	Safes["ht"] = cmdHt
	Safes["ht-raw"] = cmdHtRaw
	Safes["ht-entity"] = cmdHtEntity
	Safes["ht-tag"] = cmdHtTag
	Safes["ht-peek"] = cmdHtPeek
}
