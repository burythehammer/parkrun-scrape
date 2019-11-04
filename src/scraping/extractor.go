package scraping

import "strconv"

func stringToInt(s string) int {

	position, err := strconv.ParseInt(s, 10, 0)

	if err != nil {
		panic(err)
	}

	return int(position)
}
