package scraping

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testResult *ParkrunResult

var parkrunTestEvent = ParkrunEvent{
	eventName:   "finsbury",
	eventNumber: 488,
}

func getTestResult() *ParkrunResult {
	if testResult == nil {
		scraper := Scraper{collector: NewCollector()}
		testResult = scraper.ScrapeParkrunEvent(parkrunTestEvent)
	}
	return testResult
}

func Test_Scraper_ContainsEventInfo(t *testing.T) {
	results := getTestResult()
	require.NotNil(t, results)

	assert.Equal(t, results.eventId, parkrunTestEvent.eventName)
	assert.Equal(t, results.eventNumber, parkrunTestEvent.eventNumber)
}

func Test_Scraper_ContainsResult(t *testing.T) {
	results := getTestResult()
	require.NotNil(t, results)
	require.NotEmpty(t, results)

	exampleResult := AthleteResult{
		Position:    41,
		Name:        "Paul SINTON-HEWITT",
		Time:        "20:51",
		AgeGroup:    "VM55-59",
		AgeGrading:  "76.26",
		Gender:      "Male",
		Club:        "Ranelagh Harriers",
		Achievement: "New PB!",
		Runs:        449,
	}

	assert.Contains(t, results.results, exampleResult)

}

func Test_Scraper_DoesNotContainEmptyResult(t *testing.T) {
	results := getTestResult()
	require.NotNil(t, results)

	for _, result := range results.results {
		assert.NotEmpty(t, result.Name)
	}

}

func Test_Scraper_ContainsUnknownResult(t *testing.T) {
	results := getTestResult()
	require.NotNil(t, results)

	unknownResult := AthleteResult{
		Position: 51,
		Name:     "Unknown",
	}

	assert.Contains(t, results.results, unknownResult)
}
