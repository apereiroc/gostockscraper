package main

import (
	"fmt"
	"gostockscrape/argparser"
	"gostockscrape/scraper"
	"os"
)

func main() {
	parser := argparser.New(os.Args[1:])
	arg := parser.Get()
	isJson := parser.IsJSON()

	fmt.Printf("Parsed args: %s\n", parser.Get())
	fmt.Printf("Is JSON? %t\n", parser.IsJSON())

	scraper := scraper.NewScraper()

	scraper.Scrap(arg, isJson)
}
