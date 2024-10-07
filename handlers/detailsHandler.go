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
		errTxt := "Paths cross empty void\nSeeking what once existed\nSilence answers all."
		ErrorHandler(w, r, http.StatusNotFound, errTxt, "404 Not Found", "Artist")
		return
	}

	artist, err := Artist(artistID)
	if artist.ArtistName == "Not Found" {
		errTxt := "Paths cross empty void\nSeeking what once existed\nSilence answers all."
		ErrorHandler(w, r, http.StatusNotFound, errTxt, "404 Not Found", "Artist")
		return
	}

	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}

	t, err := template.ParseFiles("templates/details.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}
	t.Execute(w, artist)
}

func Artist(id int) (ArtistDetails, error) {
	artistList, err := api.DecodeArtists()
	if err != nil {
		return ArtistDetails{}, err
	}

	if len(artistList) < id {
		return ArtistDetails{ArtistName: "Not Found"}, nil
	}

	artist := artistList[id-1]

	concertAPI, err := api.DecodeRelations()
	if err != nil {
		return ArtistDetails{}, err
	}

	concerts := api.RelationMap(concertAPI)

	artistD := ArtistDetails{
		ArtistName:   artist.ArtistName,
		ArtistImage:  artist.ArtistImage,
		BandMembers:  artist.BandMembers,
		CreationDate: artist.CreationDate,
		FirstAlbum:   artist.FirstAlbum,
		Concerts:     concerts[id],
	}

	return artistD, nil
}
