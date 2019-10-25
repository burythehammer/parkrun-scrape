package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestScrapeParkrunLatestResults(t *testing.T) {
	results := ScrapeParkrunLatestResults("victoriadock")
	assert.NotNil(t, results)
	assert.Contains(t, results, "Holly COOK")
}
