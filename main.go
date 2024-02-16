package main

import (
	"gostockscraper/argparser"
	"gostockscraper/scraper"
	"log"
	"os"
)

func main() {
	parser := argparser.New(os.Args[1:])
	arg := parser.Get()
	isJson := parser.IsJSON()

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	logger.Println("Arg passed:", arg)
	logger.Println("Requested JSON:", isJson)

	scraper := scraper.New()

	scraper.Scrap(arg, isJson)
}
