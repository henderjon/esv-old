package fv

import (
	"github.com/henderjon/esv/esvapi"
	"log"
	"math"
	"os"
	"text/template"
	"time"
)

const (
	start = "Jan 10, 2011" // 5 year cycle started here
	Base  = iota
	Extended
	Foundation
)

var sets map[int][]entry

func init() {
	sets = make(map[int][]entry, 0)
	sets[Base] = BaseSet
	sets[Extended] = ExtendedSet
	sets[Foundation] = FoundationSet
}

// get the current verse, given a start date
func Current() int {
	// Mon, 02 Jan 2006 15:04:05 -0700
	format := "Jan 2, 2006"
	t, err := time.Parse(format, start)
	if err != nil {
		log.Fatal(err)
	}
	return int(math.Floor(time.Since(t).Hours() / 24 / 7))
}

// output the verse to STDOUT
func Render(set, vn int) {
	t, _ := template.New("all").Parse("\n\n{{.}}\n\n")

	// vn -= 1 // zero index
	for vn > 259 {
		vn -= 260 // start the cycle over
	}

	v := esvapi.Query(sets[set][vn].Ref)

	err := t.Execute(os.Stdout, (&v).String())

	if err != nil {
		log.Fatal(err)
	}
}

func Ref(set, vn int) {
	t, _ := template.New("all").Parse(`

Reference: {{.Ref}}
Set: {{.Set}}
Bible Order: {{.BibleOrder}}
Set Order: {{.SetOrder}}

`)

	// vn -= 1 // zero index
	for vn > 259 {
		vn -= 260 // start the cycle over
	}

	err := t.Execute(os.Stdout, sets[set][vn])

	if err != nil {
		log.Fatal(err)
	}

}
