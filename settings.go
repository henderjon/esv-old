package main

import (
	"github.com/henderjon/jsonstore"
	// "os"
	"log"
	"os/user"
)

var settingsDir *jsonstore.Bucket

func init() {
	u, _ := user.Current()

	var err error
	settingsDir, err = jsonstore.Open(u.HomeDir + "/.esv")
	if err != nil {
		log.Fatal(err, u.HomeDir)
	}
}

type settings struct {
	Fighter_verse_start     string
	Fighter_verse_set       int
	Journal_last_read_month int
	Journal_last_read_day   int
}

func (s *settings) get() {
	settingsDir.Get("settings.json", s)
}

func (s *settings) set() {
	settingsDir.Put("settings.json", s)
}
