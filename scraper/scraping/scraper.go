package scraping

import (
	"github.com/gocolly/colly"
	"time"
	"fmt"
	"strings"
)

type ParkrunEvent struct {
	eventName   string
	eventNumber int
}

type Scraper struct {
	collector *colly.Collector
}

func NewCollector() *colly.Collector {
	c := colly.NewCollector()

	c.Limit(&colly.LimitRule{
		// Set a delay between requests to these domains
		Delay: 2 * time.Second,
		// Add an additional random delay
		RandomDelay: 2 * time.Second,
	})

	return c
}

func (s Scraper) ScrapeLatestResults(parkrunName string) *ParkrunResult {
	url := fmt.Sprintf("https://www.parkrun.org.uk/%s/results/latestResults", parkrunName)
	results := s.scrapeAthleteResults(url)
	return &ParkrunResult{results: results}
}

func (s Scraper) ScrapeParkrunEvent(parkrun ParkrunEvent) *ParkrunResult {

	url := fmt.Sprintf("https://www.parkrun.org.uk/%s/results/weeklyresults?runSeqNumber=%d", parkrun.eventName, parkrun.eventNumber)

	results := s.scrapeAthleteResults(url)

	return &ParkrunResult{results: results}
}

func (s Scraper) scrapeAthleteResults(url string) []AthleteResult {
	results := []AthleteResult{}
	s.collector.OnHTML("table", func(tableElement *colly.HTMLElement) {
		tableElement.ForEach("tr", func(row int, rowElement *colly.HTMLElement) {
			if row == 0 {
				return
			}

			result := AthleteResult{}

			unknownResult := strings.Contains(rowElement.Text, "Unknown")

			rowElement.ForEach("td", func(col int, columnElement *colly.HTMLElement) {

				if unknownResult {
					result = extractUnknownRunnerInfo(result, col, columnElement.Text)
				} else {
					result = extractKnownRunnerInfo(result, col, columnElement.Text)
				}
			})

			results = append(results, result)
		})

	})

	s.collector.Visit(url)
	return results
}
