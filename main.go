package main

import (
	"genrefinder/apicalls"
	"genrefinder/auth"
)

func main() {
	spotify := &auth.AppSecrets{}
	token := spotify.GetToken()
	songs := &apicalls.PlaylistData{}
	songs.GetArtists(token)
}
