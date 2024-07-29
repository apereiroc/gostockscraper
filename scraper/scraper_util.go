package scraper

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// Everything is based on yahoo finances
func getUrl(company string) string {
	return "https://finance.yahoo.com/quote/" + company
}

func getCompanyName(doc *goquery.Document) (string, error) {
	name := doc.Find("#quote-header-info").Find("h1").Text()

	if len(name) == 0 {
		return name, errors.New("Company name is empty")
	}
	return name, nil
}

func getMarketOpen(doc *goquery.Document) (string, error) {
	result := doc.Find("#quote-market-notice").Text()

	if len(result) == 0 {
		return result, errors.New("Market open is empty")
	}
	return result, nil
}

func getCompanyDataStr(str, companySymbol string, doc *goquery.Document) (string, error) {
	findString := fmt.Sprintf("fin-streamer[data-field=%s][data-symbol=%s]", str, companySymbol)
	result := doc.Find(findString).AttrOr("value", "")
	if len(result) == 0 {
		return result, errors.New("Company data  is empty")
	}
	return result, nil
}

func getCompanyDataFloat(str, companySymbol string, doc *goquery.Document) (float32, error) {
	valueStr, err := getCompanyDataStr(str, companySymbol, doc)

	handleErr(err)

	// Cast to float
	value, err := strconv.ParseFloat(valueStr, 32)

	handleErr(err)

	return float32(value), nil
}

func getRegularMarketPrice(companySymbol string, doc *goquery.Document) (float32, error) {
	return getCompanyDataFloat("regularMarketPrice", companySymbol, doc)
}

func getRegularMarketChange(companySymbol string, doc *goquery.Document) (float32, error) {
	return getCompanyDataFloat("regularMarketChange", companySymbol, doc)
}
