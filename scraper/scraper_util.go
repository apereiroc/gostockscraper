package scraper

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Everything is based on yahoo finances
func getUrl(company string) string {
	return "https://finance.yahoo.com/quote/" + company
}

func getStringFromDocument(stringToBeFound, errorString string, doc *goquery.Document) (string, error) {
	result := doc.Find(stringToBeFound).Text()

	if len(result) == 0 {
		return result, errors.New(errorString)
	}

	return result, nil
}

func getCompanyTitle(doc *goquery.Document) (string, error) {
	// name := doc.Find("#quote-header-info").Find("h1").Text()
	//
	// if len(name) == 0 {
	// 	return name, errors.New("company name is empty")
	// }
	// return name, nil
	identifier := "h1.yf-3a2v0c"
	errorString := "company title not found"
	return getStringFromDocument(identifier, errorString, doc)
}

func getMarketOpen(doc *goquery.Document) (string, error) {
	// result := doc.Find("#quote-market-notice").Text()
	//
	// if len(result) == 0 {
	// 	return result, errors.New("market open is empty")
	// }
	// return result, nil
	identifier := "span.yf-1dnpe7s"
	errorString := "market open not found"
	return getStringFromDocument(identifier, errorString, doc)
}

func isMarketOpen(doc *goquery.Document) (bool, error) {
	marketString, err := getMarketOpen(doc)
	if err != nil {
		return false, err
	}
	return strings.Contains(marketString, "Open"), nil
}

// Generic function to get the company data
func getCompanyDataStr(findString string, doc *goquery.Document) (string, error) {
	selection := doc.Find(findString)
	if selection.Length() == 0 {
		return "", fmt.Errorf("element not found")
	}
	dataValue, exists := selection.Attr("data-value")
	if !exists {
		return "", fmt.Errorf("data-value attribute not found")
	}
	return dataValue, nil
}

// Function to get the regular market price
func getRegularMarketPrice(doc *goquery.Document) (string, error) {
	findString := "fin-streamer.livePrice.yf-mgkamr[data-field='regularMarketPrice']"
	return getCompanyDataStr(findString, doc)
}

// Function to get the regular market change
func getRegularMarketChangeAbsolute(doc *goquery.Document) (string, error) {
	findString := "fin-streamer.priceChange.yf-mgkamr[data-field='regularMarketChange']"
	return getCompanyDataStr(findString, doc)
}

// Function to get the regular market change in percent
func getRegularMarketChangePercent(doc *goquery.Document) (string, error) {
	findString := "fin-streamer.priceChange.yf-mgkamr[data-field='regularMarketChangePercent']"
	return getCompanyDataStr(findString, doc)
}

func parseDataToFloat(valueString string) (float64, error) {
	// Cast to float
	value, err := strconv.ParseFloat(valueString, 64)

	return value, err
}
