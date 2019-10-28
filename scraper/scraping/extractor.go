package scraping

import "strconv"

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
