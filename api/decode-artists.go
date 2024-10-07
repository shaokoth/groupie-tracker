package api

import (
	"encoding/json"
)

type Artist struct {
	ArtistID     int      `json:"id"`
	ArtistName   string   `json:"name"`
	ArtistImage  string   `json:"image"`
	BandMembers  []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func DecodeArtists() ([]Artist, error) {
	apiBody, err := FetchAPI("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}

	artists := []Artist{}

	err = json.Unmarshal(apiBody, &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}
