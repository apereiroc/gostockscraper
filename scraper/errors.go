package scraper

import "log"

func handleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
