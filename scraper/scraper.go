package scraper

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type Scraper struct{}

func New() *Scraper {
	return &Scraper{}
}

func (scraper Scraper) Scrap(arg string, isJson bool) {
	if isJson {
		scraper.scrapFile(arg)
	} else {
		scraper.scrapSingleCompany(arg)
	}
}

func (scraper Scraper) scrapFile(file string) {
}

func (scraper Scraper) scrapSingleCompany(company string) {
	//
	url := "https://finance.yahoo.com/quote/" + company

	log.Println("Requested to get info from URL: ", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Result:", res)
	fmt.Println("Doc:", doc)
}
