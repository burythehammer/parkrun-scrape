package scraping

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
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

	result := ParkrunResult{
		eventId: parkrunName,
		results: results,
	}

	return &result
}

func (s Scraper) ScrapeParkrunEvent(parkrun ParkrunEvent) *ParkrunResult {

	url := fmt.Sprintf("https://www.parkrun.org.uk/%s/results/weeklyresults?runSeqNumber=%d", parkrun.eventName, parkrun.eventNumber)

	results := s.scrapeAthleteResults(url)

	result := ParkrunResult{
		eventId:     parkrun.eventName,
		eventNumber: parkrun.eventNumber,
		results:     results,
	}

	return &result
}

func (s Scraper) scrapeAthleteResults(url string) []AthleteResult {
	results := []AthleteResult{}
	s.collector.OnHTML("tr", func(rowElement *colly.HTMLElement) {

		if rowElement.Attr("class") != "Results-table-row" {
			return
		}

		var result AthleteResult

		if strings.Contains(rowElement.Text, "Unknown") {
			result = AthleteResult{
				Position: stringToInt(rowElement.Attr("data-position")),
				Name:     rowElement.Attr("data-name"),
			}
		} else {
			result = AthleteResult{
				Position:    stringToInt(rowElement.Attr("data-position")),
				Name:        rowElement.Attr("data-name"),
				AgeGroup:    rowElement.Attr("data-agegroup"),
				AgeGrading:  rowElement.Attr("data-agegrade"),
				Gender:      rowElement.Attr("data-gender"),
				Club:        rowElement.Attr("data-club"),
				Achievement: rowElement.Attr("data-achievement"),
				Runs:        stringToInt(rowElement.Attr("data-runs")),
				Time:        extractTimeElement(rowElement),
			}
		}

		results = append(results, result)
	})

	s.collector.Visit(url)
	return results
}

func extractTimeElement(row *colly.HTMLElement) string {

	var parkrunTime string

	row.ForEach("td.Results-table-td--time", func(i int, timeElement *colly.HTMLElement) {

		timeElement.ForEach("div.compact", func(i int, timeElement *colly.HTMLElement) {
			parkrunTime = timeElement.Text
		})
	})

	return parkrunTime
}
