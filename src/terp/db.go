package terp

import (
	"io/ioutil"
	"regexp"
	"strings"
)

type Record struct {
	Bundle	string
	Field	string
	Volume	string
	Page	string
	Suffix	string
	Values	[]T
}

var ColumnSplit_rx = regexp.MustCompile("^([A-Za-z0-9_]+)([:](.*))$")

func ParseFileToRecords(fname string, bundle, volume, page string) []Record {
	all, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(all), "\n")

	z := make([]Record, 0, 4)
	for _, line := range lines {
		if len(line) > 4 && line[:4] == "/set" {
			words := ParseList(line) // TODO: Ignore errors.
			if len(words) < 2 {
				continue // Ignore short lines.
			}
			m := ColumnSplit_rx.FindStringSubmatch(words[1].String())
			if len(m) > 0 {
				r := Record{
					Bundle: bundle,
					Field: m[1],
					Volume: volume,
					Page: page,
					Suffix: m[3],
					Values: words[2:],
				}
				z = append(z, r)
			}
		}
	}
	return z
}

func ScanForDatabaseEntityFiles(dataDir string) []Record {
	return nil
}


