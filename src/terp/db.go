package terp

import (
	. "fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Record struct {
	Site	string
	Field	string
	Volume	string
	Page	string
	Suffix	string
	Values	[]T
}

var Records []*Record

var ColumnSplit_rx = regexp.MustCompile("^([A-Za-z0-9_]+)([:](.*))$")

var InternalFileName_rx = regexp.MustCompile("^[a-z][.]([-A-Za-z0-9_.]+)$")

func ParseFileToRecords(fname string, site, volume, page string, z []*Record) []*Record {
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
				r := &Record{
					Site: site,
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

func ScanSites(dataDir string) []*Record {
	log.Printf("ScanSites %s", dataDir)
	var z []*Record = make([]*Record, 0, 4)

	sites, err := ioutil.ReadDir(dataDir) 
	if err != nil {
		panic(err)
	}

	for _, f := range sites {
		m := InternalFileName_rx.FindStringSubmatch(f.Name())
		if f.IsDir() && len(m) > 0 {
			site := m[1]
			z = ScanVolumes(filepath.Join(dataDir, f.Name()), site, z)
		}
	}

	return z
}

func ScanVolumes(siteDir string, site string, z []*Record) []*Record {
	log.Printf("ScanVolumes %s %s", siteDir, site)
	volumes, err := ioutil.ReadDir(siteDir)
	if err != nil {
		panic(err)
	}

	for _, v := range volumes {
		m := InternalFileName_rx.FindStringSubmatch(v.Name())
		if v.IsDir() && len(m) > 0 {
			volume := m[1]
			z = ScanPages(filepath.Join(siteDir, v.Name()), site, volume, z)
		}
	}
	return z
}

func ScanPages(volumeDir string, site, volume string, z []*Record) []*Record {
	log.Printf("ScanPages %s %s %s", volumeDir, site, volume)
	pages, err := ioutil.ReadDir(volumeDir)
	if err != nil {
		panic(err)
	}

	for _, p := range pages {
		m := InternalFileName_rx.FindStringSubmatch(p.Name())
		if p.IsDir() && len(m) > 0 {
			fname := filepath.Join(volumeDir, p.Name(), "f.@wiki", "r.0")
			page := m[1]

			// Test whether a db file exists.
			fd, fdErr := os.Open(fname)
			if fdErr != nil {
				continue  // Skip if cannot open.
			}
			fd.Close()  // Close the test.

			z = ParseFileToRecords(fname, site, volume, page, z)
		}
	}
	return z
}

func SelectLike(db []*Record, site, field, volume, page, suffix, value string) []*Record {
	var z []*Record = make([]*Record, 0, 4)

	for _, r := range db {
		if !MatchTailStar(site, r.Site) {
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

// eget [site] table id field tag -> values

func EntityGet(db []*Record, site, table, id, field, tag string) []T {
	volume := "Entity"
	page := Sprintf("Table%s_%s", table, id) 

	recs := SelectLike(db, site, field, volume, page, tag, "*")
	vec := make([]T, len(recs))
	for i, r := range recs {
		vec[i] = MkList(r.Values)
	}
	return vec
}

func cmdEntityGet(fr *Frame, argv []T) T {
	site, table, id, field, tag := Arg5(argv)
	return MkList(EntityGet(Records, site.String(), table.String(), id.String(), field.String(), tag.String()))
}

// eput [site] table id field tag values

func EntityPut(db []*Record, site, table, id, field, tag string, values []T) *Record {
	volume := "Entity"
	page := Sprintf("Table%s_%s", table, id) 

	// TODO: (Re)write page, update Records, (future) update indices.
	// HACK: For now, assume it's a new record, not an update.
	r := &Record{
		Site: site,
		Field: field,
		Volume: volume,
		Page: page,
		Suffix: tag,
		Values: values,
	}
	Records = append(Records, r)
	return r
}

func cmdEntityPut(fr *Frame, argv []T) T {
	site, table, id, field, tag, values := Arg6(argv)
	return MkT(EntityPut(Records, site.String(), table.String(), id.String(), field.String(), tag.String(), values.List()))
}

// elike [site] table fieldPattern tagPattern valuePattern -> list(id)

func EntityLike(db []*Record, site, table, field, tag, value string) []string {
	volume := "Entity"
	page := Sprintf("Table%s_*", table) 

	recs := SelectLike(db, site, field, volume, page, tag, value)

	// Collect set of unique ids, using a map ents.
	ents := make(map[string]bool, 10)
	for _, r := range recs {
		i := strings.Index(r.Page, "_")
		ents[r.Page[i+1:]] = true
	}

	// Convert to a list of ids.
	vec := make([]string, 0, len(ents))
	for id, _ := range ents {
		vec = append(vec, id)
	}
	sort.Strings(vec)
	return vec
}

func cmdEntityLike(fr *Frame, argv []T) T {
	site, table, field, tag, value := Arg5(argv)
	return MkStringList(EntityLike(Records, site.String(), table.String(), field.String(), tag.String(), value.String()))
}

// etriples [site] table id fieldPattern tagPattern valuePattern -> list(field tag value)

func EntityTriples(db []*Record, site, table, id, field, tag, value string) []T {
	volume := "Entity"
	page := Sprintf("Table%s_%s", table, id) 

	recs := SelectLike(db, site, field, volume, page, tag, value)

	// Convert to a list of ids.
	vec := make([]T, 0, len(recs))
	for _, r := range recs {
		triple := MkList([]T{
			MkString(r.Field),
			MkString(r.Suffix),
			MkList(r.Values),
		})
		vec = append(vec, triple)
	}
	SortListByString(vec)
	return vec
}

func cmdEntityTriples(fr *Frame, argv []T) T {
	site, table, id, field, tag, value := Arg6(argv)
	return MkList(EntityTriples(Records, site.String(), table.String(), id.String(), field.String(), tag.String(), value.String()))
}

//

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

func cmdRebuild(fr *Frame, argv []T) T {
	dataDir := Arg1(argv)
	
	Records = ScanSites(dataDir.String())
	return Empty
}

func cmdDbSelectLike(fr *Frame, argv []T) T {
	site, field, volume, page, suffix, value := Arg6(argv)

	return MkT(SelectLike(
		Records,
		site.String(),
		field.String(),
		volume.String(),
		page.String(),
		suffix.String(),
		value.String()))
}

func init() {
	if Unsafes == nil {
	    Unsafes = make(map[string]Command, 333)
	}

	Unsafes["db-select-like"] = cmdDbSelectLike
	Unsafes["db-rebuild"] = cmdRebuild
	Unsafes["entity-get"] = cmdEntityGet
	Unsafes["entity-put"] = cmdEntityPut
	Unsafes["entity-like"] = cmdEntityLike
	Unsafes["entity-triples"] = cmdEntityTriples
}
