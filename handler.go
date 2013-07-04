package chirp

import (
	// "bytes"
	"encoding/base64"
	. "fmt"
	// "html"
	"io/ioutil"
	// "mime/multipart"
	"net/http"
	// "net/url"
	R "reflect"
	"regexp"
)

var BasicAuthRx = regexp.MustCompile(`^Basic ([A-Za-z0-9+/=]+)`)
var BasicAuthUserPwSplitterRx = regexp.MustCompile(`^Basic ([^:]*)[:](.*)$`)

func cmdHttpHandlerLambda(fr *Frame, argv []T) T {
	args, body := Arg2(argv)

	// TODO: get rid of these "w r" args.
	argList := args.List()
	if len(argList) != 2 {
		panic("Expected 2 names in http-handler arg list, the http.ResponseWriter & *http.Request")
	}
	wName := argList[0].String()
	rName := argList[1].String()

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

		fr2 := fr.G.Fr.NewFrame() // New frame from Global level.
		fr2.SetVar(wName, MkValue(R.ValueOf(w)))
		fr2.SetVar(rName, MkValue(R.ValueOf(r)))

		fr2.Cred = make(Hash)
		fr2.Cred["user"] = MkString(user)
		fr2.Cred["password"] = MkString(pw)
		fr2.Cred["method"] = MkString(r.Method)
		fr2.Cred["host"] = MkString(r.Host)

		fr2.Cred["w"] = MkValue(R.ValueOf(w))
		fr2.Cred["r"] = MkValue(R.ValueOf(r))

		// QUERY
		qh := MkHash(nil)
		for k, v := range query {
			qh.h[k] = MkString(v[0])
		}
		fr2.Cred["query"] = qh

		// SIMPLE FORM
		fh := MkHash(nil)
		for k, v := range form {
			fh.h[k] = MkString(v[0])
		}
		fr2.Cred["form"] = fh // TODO: save this in "query" too.

		hh := MkHash(nil)
		for k, v := range hdr {
			hh.h[k] = MkString(v[0])
		}
		fr2.Cred["header"] = hh

		// MIME MULTIPART FORM
		mpReader, _ := r.MultipartReader()
		if mpReader == nil {
			// Not a multipart/form-data POST.
			fr2.Cred["values"] = MkHash(nil)
			fr2.Cred["files"] = MkHash(nil)
		} else {
			mpForm, mpFormErr := mpReader.ReadForm(1000000)
			if mpFormErr != nil {
				panic(mpFormErr)
			}

			// Not sure when we find "mpvalues" in the multipart form, but we support them.
			mpvalues := make(Hash, 0)
			for k, v := range mpForm.Value {
				mpvalues[k] = MkStringList(v)
			}
			fr2.Cred["mpvalues"] = MkHash(mpvalues)

			// These are the uploaded files.
			// Keys are the form field name.
			// Values is a list of 2 items, the filename and the contents.
			uploads := make(Hash, 0)
			for k, v := range mpForm.File {
				fileHeader := v[0] // Only use the first one.
				fileR, openErr := fileHeader.Open()
				if openErr != nil {
					panic(openErr)
				}
				defer fileR.Close()
				bb, readErr := ioutil.ReadAll(fileR)
				if readErr != nil {
					panic(readErr)
				}
				filenameT := MkString(fileHeader.Filename)
				contentsT := MkString(string(bb))
				uploads[k] = MkList([]T{filenameT, contentsT})
			}
			fr2.Cred["uploads"] = MkHash(uploads)
		}

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
