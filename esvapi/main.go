package esvapi

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
)

var (
	api  = "http://www.esvapi.org/v2/rest/passageQuery"
	opts = map[string]string{
		"key":                        "IP",
		"output-format":              "plain-text",
		"include-short-copyright":    "0",
		"include-copyright":          "0",
		"include-audio-link":         "0",
		"include-word-ids":           "0",
		"include-verse-ids":          "0",
		"include-headings":           "0",
		"include-subheadings":        "0",
		"include-footnote-links":     "0",
		"include-footnotes":          "0",
		"include-verse-numbers":      "1",
		"include-passage-references": "1",
	}
)

// get a passage of scripture by reference from the ESV Web API
func Query(ref string) bytes.Buffer {

	vals := &url.Values{}
	vals.Set("passage", ref)
	for k, v := range opts {
		vals.Set(k, v)
	}

	url, err := url.Parse(api)
	if err != nil {
		log.Fatal(err)
	}

	url.RawQuery = vals.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)

	// req.Header.Set("", "")

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.Buffer{}

	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	res.Body.Close()
	return buf
}
