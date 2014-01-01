/*
	package github.com/yak-labs/chirp-lang/rpc provides one command 'rpc'
	for making rpc servers and clients.

	Example Server:

		rpc server /path list
		/net/http/ListenAndServe :8080 ""

		"list" is a command that should take two arguments, a simple string verb,
		and a hash of parameters.

	Example client:

		set r [rpc connect http://localhost:8080/path]
		puts [rpc call $r verb [hash abc 123 xyz 789]]
		rpc close $r
*/

package rpc

import . "github.com/yak-labs/chirp-lang"

import (
	. "fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

const SPECIAL_VERB_KEY = "_VERB_"

type Daemon struct {
	Path string
	Head T
}

type Client struct {
	Path string
}

func cmdRpcServe(fr *Frame, argv []T) T {
	pathT, headT := Arg2(argv)
	head := MkList(headT.List())

	http.HandleFunc(pathT.String(), func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				if re, ok := r.(error); ok {
					r = re.Error() // Convert error to string.
				}
				if re, ok := r.(Stringer); ok {
					r = re.String()
				}
				if rs, ok := r.(string); ok {
					r = rs + Sprintf("\n\tin cmdRpcServe\n\t\t%q", headT.String())
				}

				http.Error(w, Sprintf("%s", r), 555)
			}
		}()

		err := req.ParseForm()
		if err != nil {
			panic(err)
		}

		verb := Empty
		h := make(Hash)
		for k, v := range req.Form {
			if len(v) > 0 {
				if k == SPECIAL_VERB_KEY {
					verb = MkString(v[0])
				} else {
					h[k] = MkString(v[0])
				}
			} else {
				h[k] = Empty
			}
		}

		middle := MkList([]T{verb})
		tail := MkList([]T{MkHash(h)})
		command := ConcatLists([]T{head, middle, tail})

		z := fr.Apply(command)

		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(z.String()))
	})
	return Empty
}

func cmdRpcConnect(fr *Frame, argv []T) T {
	pathT := Arg1(argv)
	cp := &Client{
		Path: pathT.String(),
	}
	return MkValue(reflect.ValueOf(cp))
}

func cmdRpcClose(fr *Frame, argv []T) T {
	return Empty
}

func cmdRpcCall(fr *Frame, argv []T) T {
	clientT, verbT, hashT := Arg3(argv)

	cp := clientT.Raw().(*Client)
	verb := verbT.String()

	d := make(url.Values)
	d[SPECIAL_VERB_KEY] = []string{verb}
	for k, v := range hashT.Hash() {
		d[k] = []string{v.String()} // Always a singleton.
	}
	resp, err := http.PostForm(cp.Path, d)

	if err != nil {
		panic(err)
	}

	if resp.ContentLength <= 0 {
		panic(Sprintf("rpc call %q bad ContentLength: %d", verb, resp.ContentLength))
	}

	buf := make([]byte, resp.ContentLength)
	_, err = io.ReadFull(resp.Body, buf)
	if err != nil {
		panic(Sprintf("rpc %q call %q bad ReadFully: %q", cp.Path, verb, err.Error()))
	}

	if resp.StatusCode == 555 {
		panic(Sprintf("rpc %q call %q ERROR: %s", cp.Path, verb, string(buf)))
	}

	if resp.StatusCode != 200 {
		panic(Sprintf("rpc %q call %q bad response code: %d", cp.Path, verb, resp.StatusCode))
	}

	return MkString(string(buf))
}

var rpcEnsemble = []EnsembleItem{
	EnsembleItem{Name: "serve", Cmd: cmdRpcServe},
	EnsembleItem{Name: "connect", Cmd: cmdRpcConnect},
	EnsembleItem{Name: "close", Cmd: cmdRpcClose},
	EnsembleItem{Name: "call", Cmd: cmdRpcCall},
}

func init() {
	if Unsafes == nil {
		Unsafes = make(map[string]Command, 333)
	}
	Unsafes["rpc"] = MkEnsemble(rpcEnsemble)
}
