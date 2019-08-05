// Package apicalls is for storing all the different calls to spotify apis
package apicalls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// playlistData data from spotify playlist
type playlistData struct {
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

// GetArtists retrieves list of artist from given playlist
func GetArtists(token, playlist string) ([]string, error) {
	offset := "0"
	limit := "50"
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks?offset=%s&limit=%s", playlist, offset, limit)
	tokenHdr := "Bearer " + token
	queryMore := true
	var artists []string
	for queryMore {
		songs, err := getSongs(tokenHdr, url)
		if err != nil {
			return nil, err
		}

		for _, v := range songs.Items {
			for _, artist := range v.Track.Artists {
				artists = append(artists, artist.Href)
			}
		}
		url = songs.Next
		if songs.Next == "" {
			queryMore = false
		}
	}
	return artists, nil
}

// getSongs retrieves songs from a playlist
func getSongs(tokenHdr, url string) (playlistData, error) {

	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	contentTypeHdr := fmt.Sprint("application/json")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return playlistData{}, err
	}
	req.Header.Add("Accept", contentTypeHdr)
	req.Header.Add("Content-Type", contentTypeHdr)
	req.Header.Add("Authorization", tokenHdr)

	resp, err := client.Do(req)
	if err != nil {
		return playlistData{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return playlistData{}, err
	}
	songs := playlistData{}

	if err := json.Unmarshal(body, &songs); err != nil {
		return playlistData{}, err
	}
	return songs, nil

}
