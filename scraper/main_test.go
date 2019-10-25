package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestScrapeParkrunLatestResults(t *testing.T) {
	results := ScrapeParkrunLatestResults("victoriadock")
	assert.NotNil(t, results)

	exampleResult := ParkrunResult{
		eventId:     "",
		eventNo:     0,
		eventDate:   "",
		athleteName: "Holly COOK",
		time:        "",
		ageGrading:  0,
		position:    0,
		pb:          false,
	}

	assert.Contains(t, results, exampleResult)
}
