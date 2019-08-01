package main

import (
	"fmt"

	"github.com/simisimis/genrefinder/apicalls"
	"github.com/simisimis/genrefinder/auth"
)

func main() {
	token, err := auth.GetToken()
	var artistList []string
	if err == nil {
		artistList, err = apicalls.GetArtists(token)
	}
	for _, artist := range artistList {
		fmt.Println(artist)
	}
}
