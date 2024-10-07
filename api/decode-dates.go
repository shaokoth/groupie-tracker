package api

import (
	"encoding/json"
)

type Dates struct {
	Index []Date
}

type Date struct {
	ArtistID int      `json:"id"`
	Date     []string `json:"dates"`
}

func DecodeDates() (Dates, error) {
	apiBody, err := FetchAPI("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return Dates{}, err
	}

	dates := Dates{}
	err = json.Unmarshal(apiBody, &dates)
	if err != nil {
		return Dates{}, err
	}

	return dates, nil
}

func DateMap(dates Dates) map[int][]string {
	dMap := make(map[int][]string)

	for _, Index := range dates.Index {
		dMap[Index.ArtistID] = Index.Date
	}

	return dMap
}
