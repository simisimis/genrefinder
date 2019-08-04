package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/simisimis/genrefinder/apicalls"
	"github.com/simisimis/genrefinder/auth"
)

type genreDesc struct {
	Repeats int      `json:"repeats"`
	Artists []string `json:"artists"`
}

func main() {
	// retrieve token
	token, err := auth.GetToken()
	var artistList []string

	// retrieve artists from playlist songs
	artistList, err = apicalls.GetArtists(token)

	if err != nil {
		panic(err)
	}
	// retrieve genres per artist
	resultGenreData, err := apicalls.GetGenreMap(token, artistList)
	if err != nil {
		panic(err)
	}

	genresMap := map[string]*genreDesc{}

	for artist, genres := range resultGenreData {
		for _, genre := range genres {
			if _, ok := genresMap[genre]; ok {
				genresMap[genre].Repeats++
				genresMap[genre].Artists = append(genresMap[genre].Artists, artist)
			} else {
				genresMap[genre] = &genreDesc{1, []string{artist}}
			}

		}
	}

	genresJSON, err := json.Marshal(genresMap)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("genres_out.json", genresJSON, 0644)
	if err != nil {
		panic(err)
	}
}
