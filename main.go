package main

import (
	"fmt"

	"github.com/simisimis/genrefinder/apicalls"
	"github.com/simisimis/genrefinder/auth"
)

func main() {
	token, err := auth.GetToken()
	var artistList []string
	artistList, err = apicalls.GetArtists(token)

	if err != nil {
		panic(err)
	}
	genreMap, err := apicalls.GetGenreMap(token, artistList)
	for name, genres := range genreMap {
		fmt.Printf("Artist: %s likes following genres: %s\n", name, genres)
	}
}
