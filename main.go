package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/manifoldco/promptui"
	"github.com/simisimis/genrefinder/apicalls"
	"github.com/simisimis/genrefinder/auth"
)

type genreDesc struct {
	Repeats int      `json:"repeats"`
	Artists []string `json:"artists"`
}

func main() {
	var username string
	flag.StringVar(&username, "u", "", "Spotify username")
	flag.Parse()
	if username == "" {
		err := fmt.Errorf("Program expects spotify username as a flag")
		panic(err)
	}
	// retrieve token
	token, err := auth.GetToken()

	printPlaylists, err := apicalls.GetPlaylists(token, username)
	if err != nil {
		panic(err)
	}

	plistKeys := make([]string, 0, len(printPlaylists))
	for plistName := range printPlaylists {
		plistKeys = append(plistKeys, plistName)
	}
	prompt := promptui.Select{
		Label: "Select playlist:",
		Items: plistKeys,
		Size:  8,
	}

	_, plistSelect, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	// retrieve artists from playlist songs
	var artistList []string
	artistList, err = apicalls.GetArtists(token, printPlaylists[plistSelect])

	if err != nil {
		panic(err)
	}
	// retrieve genres per artist
	resultGenreData, err := apicalls.GetGenreMap(token, artistList)
	if err != nil {
		panic(err)
	}

	genresJSON, err := prepareJSON(resultGenreData)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("genres_out.json", genresJSON, 0644)
	if err != nil {
		panic(err)
	}
}

func prepareJSON(artistINFO map[string][]string) ([]byte, error) {
	genresMap := map[string]*genreDesc{}

	for artist, genres := range artistINFO {
		for _, genre := range genres {
			if _, ok := genresMap[genre]; ok {
				genresMap[genre].Repeats++
				genresMap[genre].Artists = append(genresMap[genre].Artists, artist)
			} else {
				genresMap[genre] = &genreDesc{1, []string{artist}}
			}

		}
	}
	return json.Marshal(genresMap)
}
