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
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	handleErr(err)

	// Get company name
	companyName, err := getCompanyName(doc)
	handleErr(err)

	// Get company data of interest
	currentValue, err := getRegularMarketPrice(company, doc)
	handleErr(err)

	currentChange, err := getRegularMarketChange(company, doc)
	handleErr(err)

	fmt.Println("Name:", companyName)
	fmt.Printf("Current value: %f (%f)\n", currentValue, currentChange)
}
