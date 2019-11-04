package scraping

// TODO not strings

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

type ParkrunResult struct {
	eventId     string
	eventNumber int
	eventDate   string
	results     []AthleteResult
}
