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

func getMarketOpen(doc *goquery.Document) string {
	return doc.Find("#quote-market-notice").Text()
}

func getCompanyDataStr(str, companySymbol string, doc *goquery.Document) string {
	findString := fmt.Sprintf("fin-streamer[data-field=%s][data-symbol=%s]", str, companySymbol)
	return doc.Find(findString).AttrOr("value", "")
}

func getCompanyDataFloat(str, companySymbol string, doc *goquery.Document) float32 {
	valueStr := getCompanyDataStr(str, companySymbol, doc)

	// Cast to float
	value, err := strconv.ParseFloat(valueStr, 32)
	if err != nil {
		log.Fatal("Error:", err)
	}

	return float32(value)
}

func getRegularMarketPrice(companySymbol string, doc *goquery.Document) float32 {
	return getCompanyDataFloat("regularMarketPrice", companySymbol, doc)
}

func getRegularMarketChange(companySymbol string, doc *goquery.Document) float32 {
	return getCompanyDataFloat("regularMarketChange", companySymbol, doc)
}
