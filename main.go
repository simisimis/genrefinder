package main

import (
	"flag"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/simisimis/genrefinder/auth"
	"github.com/simisimis/genrefinder/spotify"
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

	printPlaylists, err := spotify.GetPlaylists(token, username)
	if err != nil {
		panic(err)
	}

	plistKeys := make([]string, 0, len(printPlaylists))
	for name := range printPlaylists {
		plistKeys = append(plistKeys, name)
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
	artistList := make(map[string]spotify.Artist)
	artistList, err = spotify.GetArtists(token, printPlaylists[plistSelect], artistList)

	resultGenreData, err := spotify.GetGenreMap(token, artistList)
	for _, artist := range resultGenreData {
		fmt.Printf("artistID: %s,\n genres: %s,\n href: %s,\n name: %s,\n playlists: %s\n", artist.ID, artist.Genres, artist.Href, artist.Name, artist.Playlist)
	}
}
