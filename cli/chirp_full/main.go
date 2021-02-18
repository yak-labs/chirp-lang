package main

// Importing "goapi/full" adds 40ms to startup time (on a slow machine) and 7M to binary size,
// but also a minute of link time.
import (
	_ "github.com/yak-labs/chirp-lang/goapi/full"
	_ "github.com/yak-labs/chirp-lang/http"
	_ "github.com/yak-labs/chirp-lang/img"
	_ "github.com/yak-labs/chirp-lang/leveldb"
	_ "github.com/yak-labs/chirp-lang/posix"
	_ "github.com/yak-labs/chirp-lang/rpc"
)

import (
	"github.com/yak-labs/chirp-lang/cli"
)

func main() {
	cli.Main()
}
