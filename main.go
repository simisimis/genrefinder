package main

import (
	"flag"
	"fmt"

	"github.com/simisimis/genrefinder/auth"
	"github.com/simisimis/genrefinder/elastic"
	"github.com/simisimis/genrefinder/spotify"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("pkg", "main")

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
		log.Fatal(err)
	}
	// retrieve token
	token, err := auth.GetToken()

	allPlaylists, err := spotify.GetPlaylists(token, username)
	if err != nil {
		log.Fatal(err)
	}
	/*Scenario1: user selects playlist from the prompt
	plistKeys := make([]string, 0, len(allPlaylists))
	for key := range allPlaylists {
		plistKeys = append(plistKeys, key)
	}
	prompt := promptui.Select{
		Label: "Select playlist:",
		Items: plistKeys,
		Size:  8,
	}

	_, plistSelect, err := prompt.Run()
	if err != nil {
		panic(err)

	artistList, err = spotify.GetArtists(token, plistSelect, allPlaylists[plistSelect], artistList)
	}*/

	//Scenario2: program scans all of the input username playlists
	//*/
	// Create artistList object to store retrieved data
	artistList := make(map[string]spotify.Artist)

	// Loop through all user playlists and retrieve artist data
	for playlistKey, playlistName := range allPlaylists {
		artistList, err = spotify.GetArtists(token, playlistKey, playlistName, artistList)
	}

	// Retrieve genres for every artist
	resultGenreData, err := spotify.GetGenreMap(token, artistList)
	///*

	elastic.PostBulkData(resultGenreData)
}
