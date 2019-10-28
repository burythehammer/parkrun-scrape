package main

import (
	"github.com/gocolly/colly"
	"log"
	"fmt"
	"time"
)

func main() {
	log.Printf("Not doing anything")
}

func ScrapeAllLatestResults() {
	// get list of parkrun ids
	parkruns := []string{""}

	for _, parkrun := range parkruns {
		_ = ScrapeParkrunLatestResults(parkrun)
	}
}

// TODO not strings
type AthleteResult struct {
	Position       string
	Name           string
	Time           string
	AgeCategory    string
	AgeGrading     string
	Gender         string
	GenderPosition string
	Club           string
	PbNote         string
	TotalRuns      string
	ParkrunClubs   string
}

type ParkrunResult struct {
	eventId   string
	eventNo   string
	eventDate string
	results   []AthleteResult
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

func ScrapeParkrunLatestResults(parkrunName string) *ParkrunResult {
	results := []AthleteResult{}

	c := newCollector()
	url := fmt.Sprintf("https://www.parkrun.org.uk/%s/results/latestresults/", parkrunName)

	// Iterate through table rows

	c.OnHTML("table", func(tableElement *colly.HTMLElement) {
		log.Printf("found the results table")

		tableElement.ForEach("tr", func(row int, rowElement *colly.HTMLElement) {

			result := AthleteResult{}

			if row == 0 {
				return
			}

			rowElement.ForEach("td", func(col int, columnElement *colly.HTMLElement) {
				switch col {
				case 0:
					result.Position = columnElement.Text
					break
				case 1:
					result.Name = columnElement.Text
					break
				case 2:
					result.Time = columnElement.Text
					break
				case 3:
					result.AgeCategory = columnElement.Text
					break
				case 4:
					result.AgeGrading = columnElement.Text
					break
				case 5:
					result.Gender = columnElement.Text
				case 6:
					result.GenderPosition = columnElement.Text
					break
				case 7:
					result.Club = columnElement.Text
					break
				case 8:
					result.PbNote = columnElement.Text
					break
				case 9:
					result.TotalRuns = columnElement.Text
					break
				case 10:
					result.ParkrunClubs = columnElement.Text
					break
				default:
					panic("something went wrong")
				}
			})

			results = append(results, result)

		})

	})

	c.OnHTML("tr", func(e *colly.HTMLElement) {

	})

	c.Visit(url)

	return &ParkrunResult{results: results}
}
