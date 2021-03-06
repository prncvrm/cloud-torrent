package main

import (
	"log"

	"github.com/prncvrm/cloud-torrent/server"
	"github.com/jpillora/opts"
)

var VERSION = "0.0.1-src" //set with ldflags

func main() {
	log.Println("here")
	s := server.Server{
		Title:      "Cloud Torrent",
		Port:       3000,
		ConfigPath: "cloud-torrent.json",
	}

	o := opts.New(&s)
	o.Version(VERSION)
	o.PkgRepo()
	o.LineWidth = 96
	o.Parse()

	if err := s.Run(VERSION); err != nil {
		log.Fatal(err)
	}
}
