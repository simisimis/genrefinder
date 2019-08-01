// Package apicalls is for storing all the different calls to spotify apis
package apicalls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// PlaylistData data from spotify playlist
type PlaylistData struct {
	Href  string `json:"href"`
	Items []struct {
		Track struct {
			Artists []struct {
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"artists"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"track"`
	} `json:"items"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
}

// GetArtists retrieves a list of artists from a playlist
func (songs *PlaylistData) GetArtists(token string) {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	tokenHdr := "Bearer " + token
	playlist := "6zr6LLfSZCVr6lsReGXpL2"
	offset := "0"
	limit := "10"
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks?offset=%s&limit=%s", playlist, offset, limit)
	contentTypeHdr := fmt.Sprint("application/json")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Accept", contentTypeHdr)
	req.Header.Add("Content-Type", contentTypeHdr)
	req.Header.Add("Authorization", tokenHdr)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(body, &songs)
	fmt.Println(songs.Href, songs.Next, songs.Total)
	for _, v := range songs.Items {
		fmt.Println(v.Track.Href)
		for _, artist := range v.Track.Artists {
			fmt.Println(artist.Name)
		}
	}
}
