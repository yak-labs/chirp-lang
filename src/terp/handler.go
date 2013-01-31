package terp

import (
	// "bytes"
	"encoding/base64"
	// . "fmt"
	// "html"
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
		fr2.Eval(body)
	}

	return MkValue(R.ValueOf(handler))
}

func init() {
	if Unsafes == nil {
		Unsafes = make(map[string]Command, 333)
	}

	Unsafes["http-handler-lambda"] = cmdHttpHandlerLambda
}
