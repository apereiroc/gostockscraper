package scraper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Scraper struct{}

func New() *Scraper {
	return &Scraper{}
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
	url := getUrl(company)

	log.Println("Requested to get info from URL:", url)

	res, err := http.Get(url)
	handleErr(err)
	defer res.Body.Close()

	if res.StatusCode != 200 {
		if res.StatusCode == 404 {
			log.Panicf("url not found. is %s an existing company code?", company)
		}
		log.Panicf("status code error: %d %s\n", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	handleErr(err)

	// Get company title
	companyTitle, err := getCompanyTitle(doc)
	handleErr(err)
	fmt.Println("Company:", companyTitle)

	// Get market status
	marketOpen, err := isMarketOpen(doc)
	handleErr(err)
	fmt.Println("Is market open:", marketOpen)

	// Get company data of interest
	currentValue, err := getRegularMarketPrice(doc)
	handleErr(err)

	currentChange, err := getRegularMarketChangeAbsolute(doc)
	handleErr(err)

	currentChangePercent, err := getRegularMarketChangePercent(doc)
	handleErr(err)

	// Print results
	fmt.Printf("Current value: %s  Change: %s (%s %%)\n", currentValue, currentChange, currentChangePercent)
}
