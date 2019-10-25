package main

import (
	"github.com/gocolly/colly"
	"log"
	"fmt"
	"strings"
)

func main(){
	log.Printf("Not doing anything")
}

func ScrapeAllLatestResults() {
	// get list of parkrun ids
	parkruns := []string{""}

	for _, parkrun := range parkruns {
		_ = ScrapeParkrunLatestResults(parkrun)
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
		}
	})

	url := fmt.Sprintf("https://www.parkrun.org.uk/%s/results/latestresults/", parkrunName)
	c.Visit(url)

	return parkrunners
}
