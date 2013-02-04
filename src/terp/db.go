package terp

import (
	"bytes"
	. "fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Record struct {
	Site   string
	Field  string
	Volume string
	Page   string
	Suffix string
	Values []T
}

func (r *Record) String() string {
	// Stringing the "logical" way, with Field after Page,
	// instead of the "physical sorting" way, with Field after Site.
	list := MkList([]T{
		MkString(r.Site),
		MkString(r.Volume),
		MkString(r.Page),
		MkString(r.Field),
		MkString(r.Suffix),
		MkList(r.Values),
	})
	return list.String() // Use Tcl list format.
}

var Records []*Record
var HackGlobalDataDirectory string // hack: remember this.

var ColumnSplit_rx = regexp.MustCompile("^([A-Za-z0-9_]+)([:](.*))$")

var InternalFileName_rx = regexp.MustCompile("^[a-z][.]([-A-Za-z0-9_.]+)$")

func cmdSaveRecords(fr *Frame, argv []T) T {
	dataDir, records := Arg2(argv)
	recTs := records.List()
	recs := make([]*Record, len(recTs))
	for i, r := range recTs {
		recs[i] = r.Raw().(*Record) // Get the *Record from the Value.
	}
	SaveRecords(dataDir.String(), recs)
	return Empty
}

func SaveRecords(dataDir string, newRecs []*Record) {
	// Gather the records by site:vol:page
	h := make(map[string][]*Record)
	for _, r := range newRecs {
		key := Sprintf("%s:%s:%s", r.Site, r.Volume, r.Page)
		if _, ok := h[key]; !ok {
			h[key] = make([]*Record, 0, 4)
		}
		h[key] = append(h[key], r)
	}

	// Use the current Unix time as the timestamp.
	tstamp := strconv.FormatInt(time.Now().Unix(), 10)

	for _, recs := range h {
		volumeDir := filepath.Join(dataDir, "s."+recs[0].Site, "v."+recs[0].Volume)
		fname := filepath.Join(volumeDir, "p."+recs[0].Page, "f.@wiki", "r."+tstamp)
		RewriteFileWithRecords(fname, recs[0].Site, recs[0].Volume, recs[0].Page, recs)
	}

	// The really stupid way to do this.  Reread all files.
	Records = ScanSites(dataDir)
}

func RewriteFileWithRecords(fname string, site, volume, page string, newRecs []*Record) {
	log.Printf("RewriteFileWithRecords < %q %q %q %q %v", fname, site, volume, page, newRecs)
	oldFname := CurrentRevFilename("", site, volume, page)

	oldRecs, textlines := []*Record{}, []string{}
	_, fileErr := os.Stat(oldFname)
	if fileErr == nil {
		// File exists, so read it.
		log.Printf("... ParseFileToRecords OK ... %q %q %q %q", fname, site, volume, page)
		oldRecs, textlines = ParseFileToRecords(oldFname, site, volume, page, []*Record{})
	} else {
		log.Printf("... ParseFileToRecords ERROR %q ... %q %q %q %q", fileErr.Error(), fname, site, volume, page)
	}

	// Collect old records, by field:suffix as key.
	h := make(map[string]*Record)
	for _, oldR := range oldRecs {
		log.Printf("... OldRecord ... %v", *oldR)
		if oldR.Site != site || oldR.Volume != volume || oldR.Page != page {
			panic("Wrong oldR in RewriteFileWithRecords")
		}
		h[oldR.Field+":"+oldR.Suffix] = oldR
	}

	// Create or Replace new records, by field:suffix as key.
	for _, newR := range newRecs {
		log.Printf("... NewRecord ... %v", *newR)
		if newR.Site != site || newR.Volume != volume || newR.Page != page {
			panic("Wrong newR in RewriteFileWithRecords")
		}
		h[newR.Field+":"+newR.Suffix] = newR
	}

	// sort the keys.
	keys := make([]string, 0, len(h))
	for k, _ := range h {
		log.Printf("... Key ... %q", k)
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// First write the non-/set textlines.
	buf := bytes.NewBuffer(nil)
	for _, line := range textlines {
		buf.WriteString(line)
		buf.WriteByte('\n')
	}

	// Then append /set lines, sorted.
	for _, fieldTag := range keys {
		list := []T{MkString("/set"), MkString(fieldTag)}
		list = append(list, h[fieldTag].Values...)
		str := MkList(list).String()
		if strings.Index(str, "\n") >= 0 {
			// TODO: warn: panic("cannot have newline in db")
			continue
		}
		buf.WriteString(str)
		buf.WriteByte('\n')
	}
	log.Printf("... buf ... <<<%s>>>", buf.String())

	// TODO: create new rev.  Here we just overwrite it.
	e := ioutil.WriteFile(fname, buf.Bytes(), 0600)
	log.Printf("... ioutil.WriteFile ... < %q > %v", fname, e)
	if e != nil {
		panic(e)
	}
}

func ParseFileToRecords(fname string, site, volume, page string, z []*Record) ([]*Record, []string) {
	all, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	textlines := make([]string, 0, 10)
	lines := strings.Split(string(all), "\n")

	for _, line := range lines {
		if len(line) > 4 && line[:4] == "/set" {
			words, parseErr := ParseListOrRecover(line) // TODO: Ignore errors.
			if parseErr != nil {
				// TODO: warn.
				continue // Ignore bad lines.
			}
			if len(words) < 2 {
				continue // Ignore short lines.
			}
			m := ColumnSplit_rx.FindStringSubmatch(words[1].String())
			if len(m) > 0 {
				r := &Record{
					Site:   site,
					Field:  m[1],
					Volume: volume,
					Page:   page,
					Suffix: m[3],
					Values: words[2:],
				}
				z = append(z, r)
			}
		} else {
			textlines = append(textlines, line)
		}
	}
	log.Printf("ParseFile %s -> %d records, %d other lines", fname, len(z), len(textlines))
	return z, textlines
}

func ScanSites(dataDir string) []*Record {
	// Hack: remember dataDir in global var.
	HackGlobalDataDirectory = dataDir

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
			page := m[1]
			z = ScanCurrentRev(filepath.Join(volumeDir, p.Name()), site, volume, page, z)
		}
	}
	return z
}

func CurrentRevFilename(pageDir string, site, volume, page string) string {
	if pageDir == "" {
		// When called from ScanPages, we know the pageDir.
		// But from elsewhere, we may not, so reconstruct it.
		pageDir = filepath.Join(HackGlobalDataDirectory, "s." + site, "v." + volume, "p." + page)
	}
	log.Printf("ScanCurrentRev %s %s %s %s", pageDir, site, volume, page)
	fileDir := filepath.Join(pageDir, "f.@wiki")
	revfiles, err := ioutil.ReadDir(fileDir)
	if err != nil {
		panic(err)
	}

	n := len(revfiles)
	if n == 0 {
		// No revisions.  Return early.
		return ""  // TODO -- strick guesses "" is correct, or should we panic?
	}

	// Get the Unix timestamps into a slice of strings from the file names.
	var revs []string = make([]string, n)
	for i, r := range revfiles {
		m := InternalFileName_rx.FindStringSubmatch(r.Name())
		if !r.IsDir() && len(m) > 0 {
			revs[i] = m[1]
		}
	}

	sort.Strings(revs)

	// The current revision
	rev := revs[n-1]
	// The full file path to the current revision.
	fname := filepath.Join(fileDir, "r."+rev)
	return fname
}
func ScanCurrentRev(pageDir string, site, volume, page string, z []*Record) []*Record {
	// pageDir may be empty, if yet unknown.
	fname := CurrentRevFilename(pageDir, site, volume, page)

	// Test whether a db file exists.
	fd, fdErr := os.Open(fname)
	if fdErr != nil {
		return z // Skip if cannot open.
	}
	fd.Close() // Close the test.

	z, _ = ParseFileToRecords(fname, site, volume, page, z)
	return z
}

func SelectLike(db []*Record, site, field, volume, page, suffix, value string) []*Record {
	var z []*Record = make([]*Record, 0, 4)

	for _, r := range db {
		if !StringMatch(site, r.Site) {
			continue
		}
		if !StringMatch(field, r.Field) {
			continue
		}
		if !StringMatch(volume, r.Volume) {
			continue
		}
		if !StringMatch(page, r.Page) {
			continue
		}
		if !StringMatch(suffix, r.Suffix) {
			continue
		}

		for _, v := range r.Values {
			if StringMatch(value, v.String()) {
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
		Site:   site,
		Field:  field,
		Volume: volume,
		Page:   page,
		Suffix: tag,
		Values: values,
	}
	// Old Hack:
	Records = append(Records, r)
	// New Hack:
	SaveRecords(HackGlobalDataDirectory, []*Record{r})
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
	Unsafes["db-save-records"] = cmdSaveRecords
	Unsafes["entity-get"] = cmdEntityGet
	Unsafes["entity-put"] = cmdEntityPut
	Unsafes["entity-like"] = cmdEntityLike
	Unsafes["entity-triples"] = cmdEntityTriples
}
