// Package spotify is for storing all the different calls to spotify apis
package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Artist information about artist
type Artist struct {
	ID       string              `json:"id"`
	Genres   []string            `json:"genres"`
	Href     string              `json:"href"`
	Name     string              `json:"name"`
	Playlist map[string]struct{} `json:"playlist"`
}

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

// GetArtists2 retrieves list of artist from given playlist
func GetArtists(token, playlist string, artists map[string]Artist) (map[string]Artist, error) {
	offset := "0"
	limit := "50"
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks?offset=%s&limit=%s", playlist, offset, limit)
	tokenHdr := "Bearer " + token
	queryMore := true

	for queryMore {
		songs, err := getSongs(tokenHdr, url)
		if err != nil {
			return nil, err
		}

		for _, v := range songs.Items {
			for _, artist := range v.Track.Artists {
				var currentArtist Artist
				currentArtist, found := artists[artist.ID]
				if !found {
					currentArtist = Artist{
						ID:       artist.ID,
						Name:     artist.Name,
						Href:     artist.Href,
						Playlist: make(map[string]struct{}),
						Genres:   []string{},
					}
				}

				currentArtist.Playlist[playlist] = struct{}{}
				artists[artist.ID] = currentArtist
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
