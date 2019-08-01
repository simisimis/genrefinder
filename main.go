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
	for _, artist := range artistList {
		fmt.Println(artist)
	}
}
