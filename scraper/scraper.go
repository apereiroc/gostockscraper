package scraper

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// sc.logger.Println("Get result:", res)
	// sc.logger.Println("Result header:", res.Header)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// sc.logger.Println("Goquery result:", doc)

	// Find the review items
	doc.Find("#quote-header-info").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the name.
		name := s.Find("h1").Text()
		// name := s.Find("#qsp-price").Text()
		fmt.Println(name)
	})
	doc.Find("#mrt-node-Lead-5-QuoteHeader").Each(func(i int, s *goquery.Selection) {
		// Define function to get the desired value in string
		getValueStr := func(str string, s *goquery.Selection) string {
			findString := fmt.Sprintf("fin-streamer[data-field=%s]", str)
			return s.Find(findString).AttrOr("value", "")
		}

		// Define data fields of interest
		currentValueField := "regularMarketPrice"
		currentChangeField := "regularMarketChange"

		// Find the values in the doc
		currentValueStrResult := getValueStr(currentValueField, s)
		currentChangeStrResult := getValueStr(currentChangeField, s)

		// Cast to float
		currentValue, err := strconv.ParseFloat(currentValueStrResult, 32)
		if err != nil {
			log.Fatal("Error:", err)
		}

		currentChange, err := strconv.ParseFloat(currentChangeStrResult, 32)
		if err != nil {
			log.Fatal("Error:", err)
		}

		fmt.Printf("current value: %f (%f)\n", currentValue, currentChange)
	})
}
