package api

import (
	"encoding/json"
	"fmt"
	"os"
)

type Dates struct {
	Index []Date
}

type Date struct {
	ArtistID int      `json:"id"`
	Date     []string `json:"dates"`
}

func DecodeDates() Dates {
	apiBody := FetchAPI("https://groupietrackers.herokuapp.com/api/dates")
	Dates := Dates{}

	err := json.Unmarshal(apiBody, &Dates)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return Dates
}
