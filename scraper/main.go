package main

import (
	"github.com/gocolly/colly"
	"log"
	"fmt"
	"time"
	"strconv"
	"strings"
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
	Position       int
	Name           string
	Time           string
	AgeCategory    string
	AgeGrading     string
	Gender         string
	GenderPosition int
	Club           string
	PbNote         string
	TotalRuns      int
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

	c.Visit(url)
	return &ParkrunResult{results: results}
}
func extractUnknownRunnerInfo(result AthleteResult, col int, tableElementText string) AthleteResult {

	switch col {
	case 0:
		result.Position = stringToInt(tableElementText)
	case 1:
		result.Name = "Unknown"
	}

	return result
}

func extractKnownRunnerInfo(result AthleteResult, col int, tableElementText string) AthleteResult {
	switch col {
	case 0:
		result.Position = stringToInt(tableElementText)
		break
	case 1:
		result.Name = tableElementText
		break
	case 2:
		result.Time = tableElementText
		break
	case 3:
		result.AgeCategory = tableElementText
		break
	case 4:
		result.AgeGrading = tableElementText
		break
	case 5:
		result.Gender = tableElementText
	case 6:
		result.GenderPosition = stringToInt(tableElementText)
		break
	case 7:
		result.Club = tableElementText
		break
	case 8:
		result.PbNote = tableElementText
		break
	case 9:
		result.TotalRuns = stringToInt(tableElementText)
		break
	case 10:
		result.ParkrunClubs = tableElementText
		break
	default:
		panic("something went wrong")
	}

	return result
}

func stringToInt(s string) int {

	position, err := strconv.ParseInt(s, 10, 0)

	if err != nil {
		panic(err)
	}

	return int(position)
}
