package terp

import (
	// "bytes"
	"encoding/base64"
	. "fmt"
	// "html"
	"log"
	"net/http"
	"net/url"
	R "reflect"
	"regexp"
)

var BasicAuthRx = regexp.MustCompile(`^Basic ([A-Za-z0-9+/=]+)`)
var BasicAuthUserPwSplitterRx = regexp.MustCompile(`^Basic ([^:]*)[:](.*)$`)

type Who struct {
	user       string
	password   string
	remoteAddr string
	host       string
	site       string
	level      int
	query      url.Values
	form       url.Values
}

func cmdHttpHandlerLambda(fr *Frame, argv []T) T {
	args, body := Arg2(argv)

	argList := args.List()
	if len(argList) != 3 {
		panic("Expected 3 names in http-handler arg list, the http.ResponseWriter & *http.Request & *Who")
	}
	wName := argList[0].String()
	rName := argList[1].String()
	whoName := argList[2].String()

	handler := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		query := r.URL.Query()
		form := r.Form
		hdr := r.Header

		for k, v := range hdr {
			log.Printf("HDR %q : %q", k, v)
		}

		// Decode Authorization, if present.
		var user, pw string
		auth := hdr.Get("Authorization")
		if auth != "" {
			m1 := BasicAuthRx.FindStringSubmatch(auth)
			if m1 != nil {
				decoded, berr := base64.StdEncoding.DecodeString(m1[1])
				if berr != nil {
					m2 := BasicAuthUserPwSplitterRx.FindSubmatch(decoded)
					if m2 != nil {
						user, pw = string(m2[1]), string(m2[2])
					}
				}
			}
		}

		who := &Who{
			user:     user,
			password: pw,
			host:     hdr.Get("Host"),
			query:    query,
			form:     form,
		}

		fr2 := fr.G.Fr.NewFrame() // New frame from Global level.
		fr2.SetVar(wName, MkValue(R.ValueOf(w)))
		fr2.SetVar(rName, MkValue(R.ValueOf(r)))
		fr2.SetVar(whoName, MkValue(R.ValueOf(who)))

		fr2.Cred = make(Hash)
		fr2.Cred["user"] = MkString(user)
		fr2.Cred["password"] = MkString(pw)
		fr2.Cred["method"] = MkString(r.Method)
		fr2.Cred["host"] = MkString(r.Host)

		fr2.Cred["w"] = MkValue(R.ValueOf(w))
		fr2.Cred["r"] = MkValue(R.ValueOf(r))

		qh := MkHash()
		for k, v := range query {
			qh.h[k] = MkString(v[0])
		}
		fr2.Cred["query"] = qh

		fh := MkHash()
		for k, v := range form {
			fh.h[k] = MkString(v[0])
		}
		fr2.Cred["form"] = fh

		hh := MkHash()
		for k, v := range hdr {
			hh.h[k] = MkString(v[0])
		}
		fr2.Cred["header"] = hh

		fr2.Eval(body)
	}

	return MkValue(R.ValueOf(handler))
}


// Getting cred is safe.  Setting it is unsafe.  This is the setter.
func cmdCredPut(fr *Frame, argv []T) T {
	name, value := Arg2(argv)
	if fr.Cred == nil {
		fr.Cred = make(Hash, 4)
	}
	key := name.String()
	if _, ok := fr.Cred[key]; ok {
		panic(Sprintf("cred-set refuses to redefine key %q", key))
	}
	fr.Cred[key] = value
	return value
}


func init() {
	if Unsafes == nil {
		Unsafes = make(map[string]Command, 333)
	}

	Unsafes["http-handler-lambda"] = cmdHttpHandlerLambda
	Unsafes["cred-put"] = cmdCredPut
}
