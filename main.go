package main

import (
	"github.com/simisimis/genrefinder/apicalls"
	"github.com/simisimis/genrefinder/auth"
)

func main() {
	spotify := &auth.AppSecrets{}
	token := spotify.GetToken()
	songs := &apicalls.PlaylistData{}
	songs.GetArtists(token)
}
