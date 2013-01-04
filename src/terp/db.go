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

