package main

import (
	"github.com/gocolly/colly"
	"log"
	"fmt"
	"strings"
)

func main() {
	ScrapeAllLatestResults()
}

func ScrapeAllLatestResults() {
	// get list of parkrun ids

	parkruns := []string{"victoriadock"}

	for _, parkrun := range parkruns {
		ScrapeParkrunLatestResults(parkrun)
	}

}

func ScrapeParkrunLatestResults(parkrunName string) []string {
	parkrunners := []string{}

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		val, exists := e.DOM.Attr("href")

		if exists && strings.Contains(val, "athletehistory"){
			parkrunners = append(parkrunners, e.Text)
			log.Printf("Found parkrunner %s", e.Text)
		}
	})

	url := fmt.Sprintf("https://www.parkrun.org.uk/%s/results/latestresults/", parkrunName)
	c.Visit(url)

	return parkrunners
}
