package scraper

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type Scraper struct {
	logger *log.Logger
}

func New(logger *log.Logger) *Scraper {
	return &Scraper{logger: logger}
}

func (scraper *Scraper) Scrap(arg string, isJson bool) {
	if isJson {
		scraper.scrapFile(arg)
	} else {
		scraper.scrapSingleCompany(arg)
	}
}

func (scraper *Scraper) scrapFile(file string) {
}

func (sc *Scraper) scrapSingleCompany(company string) {
	//
	url := "https://finance.yahoo.com/quote/" + company

	sc.logger.Println("Requested to get info from URL:", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	sc.logger.Println("Get result:", res)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	sc.logger.Println("Goquery result:", doc)
}
