package lookup

import (
	"github.com/henderjon/esv/esvapi"
	"os"
	"log"
	"text/template"
)

// renders the given genre's reading for the given month/day
func Render(ref string) {
	t, _ := template.New("all").Parse("\n\n{{.}}\n\n")

	p := esvapi.Query(ref)

	err := t.Execute(os.Stdout, (&p).String())

	if err != nil {
		log.Fatal(err)
	}
}


