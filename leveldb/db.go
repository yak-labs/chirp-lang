package leveldb

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	. "github.com/yak-labs/chirp-lang"
)

func cmdOpenFile(fr *Frame, argv []T) T {
	nameT := Arg1(argv)

	db, err := leveldb.OpenFile(nameT.String(), nil)
	if err != nil {
		log.Panicf("cannot leveldb.OpenFile: %q: %v", nameT.String(), err)
	}
	return MkT(db)
}

func cmdScanRange(fr *Frame, argv []T) T {
	dbT, startT, limitT := Arg3(argv)

	iter := dbT.Raw().(*leveldb.DB).NewIterator(&util.Range{Start: []byte(startT.String()), Limit: []byte(limitT.String())}, nil)
	return MkT(iter)
}

func cmdScanPrefix(fr *Frame, argv []T) T {
	dbT, prefixT := Arg2(argv)

	iter := dbT.Raw().(*leveldb.DB).NewIterator(util.BytesPrefix([]byte(prefixT.String())), nil)
	return MkT(iter)
}

func cmdScanAll(fr *Frame, argv []T) T {
	dbT := Arg1(argv)

	iter := dbT.Raw().(*leveldb.DB).NewIterator(nil, nil)
	return MkT(iter)
}

func cmdBatch(fr *Frame, argv []T) T {
	Arg0(argv)

	batch := new(leveldb.Batch)
	return MkT(batch)
}

func init() {
	if Unsafes == nil {
		Unsafes = make(map[string]Command, 333)
	}

	Unsafes["leveldb.OpenFile"] = cmdOpenFile
	Unsafes["leveldb.ScanRange"] = cmdScanRange
	Unsafes["leveldb.ScanPrefix"] = cmdScanPrefix
	Unsafes["leveldb.ScanAll"] = cmdScanAll
	Unsafes["leveldb.Batch"] = cmdBatch
}
