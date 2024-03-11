package main

import (
	"flag"
	"fmt"
	"gostockscraper/scraper"
)

func main() {
	var company string
	var file string
	flag.StringVar(&company, "c", "", "Company symbol to be sent to the scraper")
	flag.StringVar(&file, "f", "", "File with company symbols to be sent to the scraper")
	flag.Parse()

	rest := flag.Args()

	companies := []string{company}
	companies = append(companies, company)
	companies = append(companies, rest...)

	fmt.Printf("Company: %s\n", company)
	fmt.Printf("File: %s\n", file)
	fmt.Printf("Tail: %s\n", flag.Args())

	scraper := scraper.New()

	// TODO
	// Improve this provisional splitting
	if len(company) > 0 {
		for _, c := range companies {
			scraper.Scrap(c, false)
		}
	}
	if len(file) > 0 {
		scraper.Scrap(file, true)
	}
}
