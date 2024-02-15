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
	url := "https://finance.yahoo.com/quote/" + company

	log.Println("Requested to get info from URL:", url)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Get company name
	companyName := getCompanyName(doc)

	// Get company data of interest
	currentValue := getRegularMarketPrice(doc)
	currentChange := getRegularMarketChange(doc)

	fmt.Println("Name:", companyName)
	fmt.Printf("Current value: %f (%f)\n", currentValue, currentChange)
}
