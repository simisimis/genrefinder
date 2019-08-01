package main

import (
	"github.com/simisimis/genrefinder/apicalls"
	"github.com/simisimis/genrefinder/auth"
)

func main() {
	token, err := auth.GetToken()
	if err == nil {
		songs := &apicalls.PlaylistData{}
		songs.GetArtists(token)
	}
}
