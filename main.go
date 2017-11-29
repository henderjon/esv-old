package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"time"
)

var (
	ref                string
	yearN, weekN, dayN int
	extendedSet, foundationSet,
	gospel, wisdom, nt, ot,
	text, help bool
	memoryFlags, readingFlags, refFlags *flag.FlagSet
)

type passage struct {
	BibleOrder, Year, Week int
	Set, Ref               string
}

// init sets up our subcommand flags
func init() {
	memoryFlags = flag.NewFlagSet("memoryFlags", flag.ContinueOnError)
	memoryFlags.BoolVar(&extendedSet, "extended", false, "return the current Fighter Verse from the extended set")
	memoryFlags.BoolVar(&foundationSet, "foundation", false, "return the current Fighter Verse from the foundation set")
	memoryFlags.IntVar(&yearN, "year", 0, "specify the year of the Fighter Verse to look up")
	memoryFlags.IntVar(&weekN, "week", 0, "specify the week of the Fighter Verse to look up")
	memoryFlags.BoolVar(&text, "text", false, "return the full text, not just the reference")
	memoryFlags.BoolVar(&help, "help", false, "show this message")

	readingFlags = flag.NewFlagSet("readingFlags", flag.ContinueOnError)
	readingFlags.BoolVar(&gospel, "gospel", true, "specify the gospel genre to read for the given day (default)")
	readingFlags.BoolVar(&wisdom, "wisdom", false, "specify the wisdom genre to read for the given day")
	readingFlags.BoolVar(&nt, "nt", false, "specify the nt genre to read for the given day")
	readingFlags.BoolVar(&ot, "ot", false, "specify the ot genre to read for the given day")
	readingFlags.IntVar(&dayN, "day", 0, "specify the day of the current month (1-25)")
	readingFlags.BoolVar(&text, "text", false, "return the full text, not just the reference")
	readingFlags.BoolVar(&help, "help", false, "show this message")

	refFlags = flag.NewFlagSet("refFlags", flag.ContinueOnError)
	refFlags.StringVar(&ref, "ref", "Isa 26:3", "the reference to lookup")
	refFlags.BoolVar(&help, "help", false, "show this message; subcommands are fv & reading")
}

func main() {
	if len(os.Args) < 2 {
		lookup()
	} else {
		switch os.Args[1] {
		case "fv":
			verse()
		case "reading":
			readingPlan()
		default:
			lookup()
		}
	}
}

func verse() {
	memoryFlags.Parse(os.Args[2:])
	if help {
		memoryFlags.PrintDefaults()
		return
	}

	now := time.Now()
	year, week := now.ISOWeek()

	// reset if a year was passed in
	if yearN >= 2015 {
		year = yearN
	}

	// reset if a week was passed in
	if weekN <= 52 && weekN > 0 {
		week = weekN
	}

	// normalize to a known year range
	for year > 2020 {
		year -= 5
	}

	var set map[int]map[int]passage

	switch true {
	case foundationSet:
		set = getFoundationsSet()
	case extendedSet:
		set = getExtendedSet()
	default:
		set = getBaseSet()
	}

	var out string

	if ref, ok := set[year][week]; ok {
		out = ref.Ref
	}

	if text {
		out = query(out)
	}

	render(out)

}

func lookup() {
	if len(os.Args) < 2 {
		refB := bufio.NewReader(os.Stdin)
		ref, _ = refB.ReadString('\n')
	} else {
		refFlags.Parse(os.Args[2:])
	}

	if help {
		refFlags.PrintDefaults()
		return
	}

	render(query(ref))
}

func readingPlan() {
	readingFlags.Parse(os.Args[2:])
	if help {
		readingFlags.PrintDefaults()
		return
	}

	now := time.Now()
	month, day := now.Month(), now.Day()

	if dayN <= 25 && dayN > 0 {
		day = dayN
	}

	plan := getReadingPlan()

	var out string

	ref, ok := plan[int(month)][int(day)]
	if !ok {
		log.Println("Looks like it's time to review.")
		return
	}

	switch true {
	case ot:
		out = ref.ot
	case nt:
		out = ref.nt
	case wisdom:
		out = ref.wisdom
	case gospel:
		fallthrough
	default:
		out = ref.gospel
	}

	if text {
		out = query(out)
	}

	render(out)
}
