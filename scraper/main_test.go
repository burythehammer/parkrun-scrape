package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TODO will fail every week unless randomly chosen person keeps running same event. Must think of how to stub / customise

var testResult *ParkrunResult

func getTestResult() *ParkrunResult {
	if testResult == nil {
		testResult = ScrapeParkrunLatestResults("finsbury")
	}
	return testResult
}

func Test_Scraper_ContainsResult(t *testing.T) {
	results := getTestResult()
	require.NotNil(t, results)

	exampleResult := AthleteResult{
		Position:       41,
		Name:           "Paul SINTON-HEWITT",
		Time:           "20:51",
		AgeCategory:    "VM55-59",
		AgeGrading:     "76.26 %",
		Gender:         "M",
		GenderPosition: 36,
		Club:           "Ranelagh Harriers",
		PbNote:         "New PB!",
		TotalRuns:      449,
		ParkrunClubs:   "",
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
