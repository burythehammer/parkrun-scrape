package main

import (
	"github.com/gocolly/colly"
	"log"
	"fmt"
	"time"
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

type ParkrunResult struct {
	eventId string
	eventNo int
	eventDate string // TODO datetime
	athleteName string // TODO proper struct with more details
	time string // TODO got to be a better format than this
	ageGrading float32
	position int
	pb bool
}

func newCollector() *colly.Collector {
	c := colly.NewCollector()

	c.Limit(&colly.LimitRule{
		// Set a delay between requests to these domains
		Delay: 2 * time.Second,
		// Add an additional random delay
		RandomDelay: 2 * time.Second,
	})

	return c
}

func ScrapeParkrunLatestResults(parkrunName string) []ParkrunResult {

	results := []ParkrunResult{}


	c := newCollector()
	url := fmt.Sprintf("https://www.parkrun.org.uk/%s/results/latestresults/", parkrunName)

	// Iterate through table rows
	c.OnHTML("tr", func(e *colly.HTMLElement) {

		result := ParkrunResult{}

		e.ForEach("td", func(i int, elem *colly.HTMLElement) {
			if i == 1 {
				result.athleteName = elem.Text
			}
		})

		results = append(results, result)
	})

	c.Visit(url)
	return results
}
