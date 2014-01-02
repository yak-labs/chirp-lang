package main

import (
	_ "github.com/yak-labs/chirp-lang/goapi/tiny"
	_ "github.com/yak-labs/chirp-lang/posix"
)

import (
	"github.com/yak-labs/chirp-lang/cli"
)

func main() {
	cli.Main()
}
