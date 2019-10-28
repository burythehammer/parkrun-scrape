package scraping


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
