package main

import (
	_ "github.com/yak-labs/chirp-lang/goapi/default"
	_ "github.com/yak-labs/chirp-lang/http"
	_ "github.com/yak-labs/chirp-lang/img"
	_ "github.com/yak-labs/chirp-lang/posix"
	_ "github.com/yak-labs/chirp-lang/rpc"
)

import (
	"github.com/yak-labs/chirp-lang/cli"
)

func main() {
	cli.Main()
}
