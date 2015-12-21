package lookup

import (
	"github.com/henderjon/esv/esvapi"
	"log"
	"os"
	"text/template"
)

// renders the given passage reference
func Render(ref string) {
	t, _ := template.New("all").Parse("\n\n{{.}}\n\n")

	p := esvapi.Query(ref)

	err := t.Execute(os.Stdout, (&p).String())

	if err != nil {
		log.Fatal(err)
	}
}
