package terp

import (
	"io/ioutil"
	"log"
	"path/filepath"
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

func (fr *Frame) initDbCmds() {
	Builtins["db-select-like"] = cmdDbSelectLike
	Builtins["db-scan"] = cmdDbScan
}

var ColumnSplit_rx = regexp.MustCompile("^([A-Za-z0-9_]+)([:](.*))$")

var InternalFileName_rx = regexp.MustCompile("^[a-z][.]([-A-Za-z0-9_.]+)$")

func ParseFileToRecords(fname string, bundle, volume, page string, z []Record) []Record {
	log.Printf("ParseFile %s", fname)
	all, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(all), "\n")

	for _, line := range lines {
		log.Printf("LINE: %q", line)
		if len(line) > 4 && line[:4] == "/set" {
			log.Printf("YES")
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
			log.Printf("z: %v", z)
		}
	}
	return z
}

func ScanBundles(dataDir string) []Record {
	log.Printf("ScanBundles %s", dataDir)
	var z []Record = make([]Record, 0, 4)

	bundles, err := ioutil.ReadDir(dataDir) 
	if err != nil {
		panic(err)
	}

	for _, b := range bundles {
		m := InternalFileName_rx.FindStringSubmatch(b.Name())
		if b.IsDir() && len(m) > 0 {
			bundle := m[1]
			z = ScanVolumes(filepath.Join(dataDir, b.Name()), bundle, z)
		}
	}

	return z
}

func ScanVolumes(bundleDir string, bundle string, z []Record) []Record {
	log.Printf("ScanVolumes %s %s", bundleDir, bundle)
	volumes, err := ioutil.ReadDir(bundleDir)
	if err != nil {
		panic(err)
	}

	for _, v := range volumes {
		m := InternalFileName_rx.FindStringSubmatch(v.Name())
		if v.IsDir() && len(m) > 0 {
			volume := m[1]
			z = ScanPages(filepath.Join(bundleDir, v.Name()), bundle, volume, z)
		}
	}
	return z
}

func ScanPages(volumeDir string, bundle, volume string, z []Record) []Record {
	log.Printf("ScanPages %s %s %s", volumeDir, bundle, volume)
	pages, err := ioutil.ReadDir(volumeDir)
	if err != nil {
		panic(err)
	}

	for _, p := range pages {
		m := InternalFileName_rx.FindStringSubmatch(p.Name())
		if p.IsDir() && len(m) > 0 {
			fname := filepath.Join(volumeDir, p.Name(), "r.0")
			page := m[1]

			z = ParseFileToRecords(fname, bundle, volume, page, z)
		}
	}
	return z
}

func SelectLike(db []Record, bundle, field, volume, page, suffix, value string) []Record {
	var z []Record = make([]Record, 0, 4)

	for _, r := range db {
		if !MatchTailStar(bundle, r.Bundle) {
			continue
		}
		if !MatchTailStar(field, r.Field) {
			continue
		}
		if !MatchTailStar(volume, r.Volume) {
			continue
		}
		if !MatchTailStar(page, r.Page) {
			continue
		}
		if !MatchTailStar(suffix, r.Suffix) {
			continue
		}

		for _, v := range r.Values {
			if MatchTailStar(value, v.String()) {
				z = append(z, r)
				break
			}
		}
	}

	return z
}

func MatchTailStar(pattern, str string) bool {
	println("pattern ", pattern)
	println("str ", str)
	if len(pattern) >= 1 && pattern[len(pattern)-1] == '*' {
		println("=1= ", str)
		if len(str) >= len(pattern)-1 {
		println("=2= ", str)
			return pattern[:len(pattern)-1] == str[:len(pattern)-1]
		}
	}

	println("=3= ", str)
	return pattern == str
}

func cmdDbScan(fr *Frame, argv []T) T {
	dataDir := Arg1(argv)
	
	return MkT(ScanBundles(dataDir.String()))
}

func cmdDbSelectLike(fr *Frame, argv []T) T {
	database, bundle, field, volume, page, suffix, value := Arg7(argv)

	//var db []Record = make([]Record, 0, 4)
	//for _, r := range database.List() {
	//	db = append(db, r.Raw().(Record))
	//}

	db := database.Raw().([]Record)
	
	return MkT(SelectLike(
		db,
		bundle.String(),
		field.String(),
		volume.String(),
		page.String(),
		suffix.String(),
		value.String()))
}
