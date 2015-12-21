package main

import (
	// esv "github.com/henderjon/esv/utils"
	dj "github.com/henderjon/esv/djrp"
	"github.com/henderjon/esv/fv"
	// "encoding/json"
	"os"
	"log"
	"flag"
)

var (
	sub string
	genre string
	next, prev int
	fvFlags, djFlags *flag.FlagSet
)

func init() {
	fvFlags = flag.NewFlagSet("fvFlags", flag.ContinueOnError)
	fvFlags.IntVar(&next, "next", 0, "(optional) the verse for this week +n to view")
	fvFlags.IntVar(&prev, "prev", 0, "(optional) the verse for this week -n to view")

	djFlags = flag.NewFlagSet("djFlags", flag.ContinueOnError)
	djFlags.StringVar(&genre, "genre", "all", "(optional) [gospel|wisdom|nt|ot|all] which genre to read for the given day")
}

func main() {
	switch os.Args[1] {
	case "fv" :
		verse()
	case "read" :
		reading()
	}
}

func verse() {
	err := fvFlags.Parse(os.Args[2:])
	if err != nil {
		log.Fatal(err)
	}
	s := &settings{}
	s.get()

	if s.Fighter_verse_start == "" {
		s.Fighter_verse_start = "Jan 10, 2011"
	}

	if s.Fighter_verse_set < 1 {
		s.Fighter_verse_set = fv.Base
	}

	c := fv.Current()

	c += next
	c -= prev

	fv.Render(s.Fighter_verse_set, c)
	s.set()
}

func reading() {
	djFlags.Parse(os.Args[2:])
	s := &settings{}
	s.get()

	if s.Journal_last_read_month < 1 {
		s.Journal_last_read_month = 1
	}

	if s.Journal_last_read_day < 1 {
		s.Journal_last_read_day = 1
	}

	g := make(map[int]int, 0)
	g[dj.Gospel] = dj.Gospel
	g[dj.Wisdom] = dj.Wisdom
	g[dj.Nt]     = dj.Nt
	g[dj.Ot]     = dj.Ot

	for _, v := range g {
		switch genre {
		case "gospel" :
			dj.Ref(dj.Gospel, s.Journal_last_read_month, s.Journal_last_read_day)
		case "wisdom" :
			dj.Ref(dj.Wisdom, s.Journal_last_read_month, s.Journal_last_read_day)
		case "nt" :
			dj.Ref(dj.Nt, s.Journal_last_read_month, s.Journal_last_read_day)
		case "ot" :
			dj.Ref(dj.Ot, s.Journal_last_read_month, s.Journal_last_read_day)
		default:
			dj.Ref(v, s.Journal_last_read_month, s.Journal_last_read_day)
		}
	}

	s.set()
}

