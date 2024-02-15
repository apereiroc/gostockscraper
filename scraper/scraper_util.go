package scraper

import (
	"fmt"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func getCompanyName(doc *goquery.Document) string {
	return doc.Find("#quote-header-info").Find("h1").Text()
}

func getCompanyDataStr(str string, doc *goquery.Document) string {
	findString := fmt.Sprintf("fin-streamer[data-field=%s]", str)
	return doc.Find(findString).AttrOr("value", "")
}

func getCompanyDataFloat(str string, doc *goquery.Document) float32 {
	valueStr := getCompanyDataStr(str, doc)

	// Cast to float
	value, err := strconv.ParseFloat(valueStr, 32)
	if err != nil {
		log.Fatal("Error:", err)
	}

	return float32(value)
}

func getRegularMarketPrice(doc *goquery.Document) float32 {
	return getCompanyDataFloat("regularMarketPrice", doc)
}

func getRegularMarketChange(doc *goquery.Document) float32 {
	return getCompanyDataFloat("regularMarketChange", doc)
}
