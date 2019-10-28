package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TODO will fail every week unless randomly chosen person keeps running same event. Must think of how to stub / customise

func TestScrapeParkrunLatestResults(t *testing.T) {
	results := ScrapeParkrunLatestResults("finsbury")
	assert.NotNil(t, results)

	exampleResult := AthleteResult{
		Position:       "41",
		Name:           "Paul SINTON-HEWITT",
		Time:           "20:51",
		AgeCategory:    "VM55-59",
		AgeGrading:     "76.26 %",
		Gender:         "M",
		GenderPosition: "36",
		Club:           "Ranelagh Harriers",
		PbNote:         "New PB!",
		TotalRuns:      "449",
		ParkrunClubs:   "",
	}

	assert.Contains(t, results.results, exampleResult)
}
