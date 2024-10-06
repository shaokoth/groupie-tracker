package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"groupie-tracker/api"
)

type ArtistDetails struct {
	ArtistName   string
	ArtistImage  string
	BandMembers  []string
	CreationDate int
	FirstAlbum   string
	Concerts     map[string][]string
}

type ClickedArtist struct {
	ArtistID string `json:"artistId"`
}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var clickedArtist ClickedArtist
	err := json.NewDecoder(r.Body).Decode(&clickedArtist)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	artistID, _ := strconv.Atoi(clickedArtist.ArtistID)
	artist := Artist(artistID)

	t, err := template.ParseFiles("templates/details.html")
	if err != nil {
		fmt.Println("err")
	}
	t.Execute(w, artist)
}

func Artist(id int) ArtistDetails {
	artistList := api.DecodeArtists()
	artist := artistList[id-1]

	concerts := api.RelationMap(api.DecodeRelations())

	artistD := ArtistDetails{
		ArtistName:   artist.ArtistName,
		ArtistImage:  artist.ArtistImage,
		BandMembers:  artist.BandMembers,
		CreationDate: artist.CreationDate,
		FirstAlbum:   artist.FirstAlbum,
		Concerts:     concerts[id],
	}

	return artistD
}
