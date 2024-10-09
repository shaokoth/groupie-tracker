package handlers

import (
	"net/http"
	"strconv"

	"groupie-tracker/api"
)

type ArtistDetails struct {
	ArtistID     int
	ArtistName   string
	ArtistImage  string
	BandMembers  []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	Dates        []string
	Concerts     map[string][]string
}

<<<<<<< HEAD:handlers/Filter.go
func GetID(r *http.Request) (int, error) {
=======
func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed", "Error", "Reload")
	}
>>>>>>> 1fc0de3f1a7dea6bd1038a14477f9970727df077:handlers/detailsHandler.go
	id := r.URL.Query().Get("id")

	artistID, err := strconv.Atoi(id)
	if err != nil || artistID < 1 {
		return artistID, err
	}

	return artistID, nil
}

func Artist(r *http.Request) (ArtistDetails, error) {
	id, err := GetID(r)
	if err != nil {
		return ArtistDetails{}, err
	}

	artistList, err := api.DecodeArtists()
	if err != nil {
		return ArtistDetails{}, err
	}

	if len(artistList) < id {
		return ArtistDetails{ArtistName: "Not Found"}, nil
	}

	artist := artistList[id-1]

	concertAPI, errConst := api.DecodeRelations()
	dateAPI, errDt := api.DecodeDates()
	locationAPI, errLoc := api.DecodeLocations()
	if errConst != nil || errDt != nil || errLoc != nil {
		return ArtistDetails{}, err
	}

	concerts := api.RelationMap(concertAPI)
	dates := api.DateMap(dateAPI)
	locations := api.LocationMap(locationAPI)

	artistD := ArtistDetails{
		ArtistID:     id,
		ArtistName:   artist.ArtistName,
		ArtistImage:  artist.ArtistImage,
		BandMembers:  artist.BandMembers,
		CreationDate: artist.CreationDate,
		FirstAlbum:   artist.FirstAlbum,
		Dates:        dates[id],
		Locations:    locations[id],
		Concerts:     concerts[id],
	}

	return artistD, nil
}
