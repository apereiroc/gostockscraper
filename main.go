package main

import (
	"flag"
	"gostockscraper/scraper"
)

func main() {
	var company string
	var file string
	flag.StringVar(&company, "c", "", "Company symbol to be sent to the scraper")
	flag.StringVar(&file, "f", "", "File with company symbols to be sent to the scraper")
	flag.Parse()

	scraper := scraper.New()

	// TODO
	// Improve this provisional splitting
	if len(company) > 0 {
		scraper.Scrap(company, false)
	}
	if len(file) > 0 {
		scraper.Scrap(file, true)
	}
}
