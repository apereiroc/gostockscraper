package scraper

type Scraper struct{}

func NewScraper() *Scraper {
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
}
