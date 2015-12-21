package djrp

import (
	"bytes"
	"github.com/henderjon/esv/esvapi"
	"log"
	"os"
	"text/template"
)

const (
	Jan = iota
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
	Gospel
	Wisdom
	Ot
	Nt
)

// renders the given genre's reading for the given month/day
func Render(g, m, d int) {
	t, _ := template.New("all").Parse("\n\n{{.}}\n\n")

	m -= 1 // zero index
	d -= 1 // zero index

	var p bytes.Buffer

	switch {
	case g == Gospel:
		p = esvapi.Query(Plan[m][d].gospel)
	case g == Wisdom:
		p = esvapi.Query(Plan[m][d].wisdom)
	case g == Ot:
		p = esvapi.Query(Plan[m][d].ot)
	case g == Nt:
		p = esvapi.Query(Plan[m][d].nt)
	}

	err := t.Execute(os.Stdout, (&p).String())

	if err != nil {
		log.Fatal(err)
	}
}

// renders the given genre's reference for the given month/day
func Ref(g, m, d int) {
	t := template.New("all")

	m -= 1 // zero index
	d -= 1 // zero index

	var ref string

	switch {
	case g == Gospel:
		ref = Plan[m][d].gospel
		t, _ = t.Parse("\nGospel: {{.}}\n\n")
	case g == Wisdom:
		ref = Plan[m][d].nt
		t, _ = t.Parse("\nWisdom: {{.}}\n\n")
	case g == Ot:
		ref = Plan[m][d].ot
		t, _ = t.Parse("\nOld Testament: {{.}}\n\n")
	case g == Nt:
		ref = Plan[m][d].wisdom
		t, _ = t.Parse("\nNew Testament: {{.}}\n\n")
	}

	err := t.Execute(os.Stdout, ref)

	if err != nil {
		log.Fatal(err)
	}
}
