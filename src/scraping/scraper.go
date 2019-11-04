package scraping

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

//AthleteResult contains the scraped data associated with a single athlete at a single parkrun event.
type AthleteResult struct {
	Position    int
	Name        string
	AgeGroup    string
	AgeGrading  string
	Gender      string
	Club        string
	Achievement string
	Runs        int
	Time        string
}

//ParkrunResult contains the scraped data associated with all athletes at a single parkrun event.
type ParkrunResult struct {
	eventId     string
	eventNumber int
	eventDate   string
	results     []AthleteResult
}

//ParkrunEvent is the input data required to scrape a parkrun event.
type ParkrunEvent struct {
	eventName   string
	eventNumber int
}

//Scraper is an object used to scrape various kinds data from the parkrun website.
type Scraper struct {
	collector *colly.Collector
}

//NewCollector generates a colly collector with default parameters.
func NewCollector() *colly.Collector {
	c := colly.NewCollector()

	err := c.Limit(&colly.LimitRule{
		DomainRegexp: ".*",
		Delay:        2 * time.Second,
		RandomDelay:  2 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	return c
}

func (s Scraper) ScrapeLatestResults(parkrunName string) *ParkrunResult {
	url := fmt.Sprintf("https://www.parkrun.org.uk/%s/results/latestResults", parkrunName)
	results, err := s.scrapeAthleteResults(url)

	if err != nil {
		panic(err)
	}

	result := ParkrunResult{
		eventId: parkrunName,
		results: results,
	}

	return &result
}

func (s Scraper) ScrapeParkrunEvent(parkrun ParkrunEvent) *ParkrunResult {

	url := fmt.Sprintf("https://www.parkrun.org.uk/%s/results/weeklyresults?runSeqNumber=%d", parkrun.eventName, parkrun.eventNumber)

	results, err := s.scrapeAthleteResults(url)

	if err != nil {
		panic(err)
	}

	result := ParkrunResult{
		eventId:     parkrun.eventName,
		eventNumber: parkrun.eventNumber,
		results:     results,
	}

	return &result
}

func (s Scraper) scrapeAthleteResults(url string) ([]AthleteResult, error) {
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

	err := s.collector.Visit(url)
	return results, err
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

func stringToInt(s string) int {

	position, err := strconv.ParseInt(s, 10, 0)

	if err != nil {
		panic(err)
	}

	return int(position)
}
