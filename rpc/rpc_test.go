package rpc

import (
	_ "github.com/yak-labs/chirp-lang/goapi/default"

	"github.com/yak-labs/chirp-lang"
	"testing"
)

var rpcTests = `
    # This RPC will return 3 words: "extra_word", the verb, and the input hash:
	rpc serve /just_list "list extra_word"

    # Serve on port 31234, and wait 0.3 sec for it to start.
	go {/net/http/ListenAndServe localhost:31234 ""}
	/time/Sleep [expr {300 * [/time/Millisecond]}]

    # Call the RPC, storing resulting 3 words in w,v,h.
	set r [rpc connect http://localhost:31234/just_list]
	set w,v,h [rpc call $r SomeVerb [hash abc 123 xyz 789]]
	rpc close $r

    # Check $w and $v.
	must "extra_word" $w
	must "SomeVerb" $v

    # Check the hash.
	set h2 [hash $h]
	must "abc xyz" [hkeys $h2]
	must "123" [hget $h2 abc]
	must "789" [hget $h2 xyz]
`

func TestRpc(a *testing.T) {
	// chirp.Debug['a'] = true
	// chirp.Debug['e'] = true
	// chirp.Debug['w'] = true
	fr := chirp.NewInterpreter()
	fr.Eval(chirp.MkString(rpcTests))
}
