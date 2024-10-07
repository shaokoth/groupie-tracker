package handlers

import (
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

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	artistID, err := strconv.Atoi(id)
	if err != nil || artistID < 1 {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error")
	}

	artist := Artist(artistID)

	t, err := template.ParseFiles("templates/details.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error")
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
