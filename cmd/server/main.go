package main

import (
	"flag"
	"log"
	"os"

	"github.com/sano-home/thermohygrometer/api"
)

var (
	host   = flag.String("host", "0.0.0.0", "The listen host")
	port   = flag.String("port", "8000", "The listen port")
	dbPath = flag.String("db", "", "The path to the SQLite3 db file")
)

func main() {
	flag.Parse()

	s, err := api.NewServer(*dbPath)
	if err != nil {
		log.Printf("api.NewServer failed: %v", err)
		os.Exit(1)
	}
	if err := s.Run(*host, *port); err != nil {
		log.Printf("s.Run failed: %v", err)
		os.Exit(1)
	}
}
