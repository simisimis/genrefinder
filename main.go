package main

import (
	"fmt"

	"github.com/simisimis/genrefinder/apicalls"
	"github.com/simisimis/genrefinder/auth"
)

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
	genreMap, err := apicalls.GetGenreMap(token, artistList)
	for name, genres := range genreMap {
		fmt.Printf("Artist: %s likes following genres: %+q\n", name, genres)
	}
}
