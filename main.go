package main

import (
	// esv "github.com/henderjon/esv/utils"
	dj "github.com/henderjon/esv/djrp"
	"github.com/henderjon/esv/fv"
	// "encoding/json"
	"os"
	// "log"
	"flag"
)

var (
	sub string
	genre string
	next, prev int
	help bool
	fvFlags, djFlags *flag.FlagSet
)

func init() {
	fvFlags = flag.NewFlagSet("fvFlags", flag.ContinueOnError)
	fvFlags.IntVar(&next, "next", 0, "(optional) the verse for this week +n to view")
	fvFlags.IntVar(&prev, "prev", 0, "(optional) the verse for this week -n to view")
	fvFlags.BoolVar(&help, "help", false, "(optional) show this message")

	djFlags = flag.NewFlagSet("djFlags", flag.ContinueOnError)
	djFlags.StringVar(&genre, "genre", "all", "(optional) [gospel|wisdom|nt|ot|all] which genre to read for the given day")
	djFlags.BoolVar(&help, "help", false, "(optional) show this message")
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
	fvFlags.Parse(os.Args[2:])
	if help {
		fvFlags.PrintDefaults()
		return
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
	if help {
		djFlags.PrintDefaults()
		return
	}
	s := &settings{}
	s.get()

	if s.Journal_last_read_month < 0 {
		s.Journal_last_read_month = 0
	}

	if s.Journal_last_read_day < 0 {
		s.Journal_last_read_day = 0
	}

	s.set() // record defaults

	g := make(map[int]int, 0)
	g[dj.Gospel] = dj.Gospel
	g[dj.Wisdom] = dj.Wisdom
	g[dj.Nt]     = dj.Nt
	g[dj.Ot]     = dj.Ot

	for _, v := range g {
		switch genre {
		case "gospel" :
			dj.Render(dj.Gospel, s.Journal_last_read_month, s.Journal_last_read_day)
			return
		case "wisdom" :
			dj.Render(dj.Wisdom, s.Journal_last_read_month, s.Journal_last_read_day)
			return
		case "nt" :
			dj.Render(dj.Nt, s.Journal_last_read_month, s.Journal_last_read_day)
			return
		case "ot" :
			dj.Render(dj.Ot, s.Journal_last_read_month, s.Journal_last_read_day)
			return
		default:
			dj.Render(v, s.Journal_last_read_month, s.Journal_last_read_day)
			return
		}
	}


}

