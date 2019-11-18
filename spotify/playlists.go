// Package spotify is for storing all the different calls to spotify apis
package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type userPlaylists struct {
	Href  string `json:"href"`
	Items []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"items"`
	Limit  int    `json:"limit"`
	Next   string `json:"next"`
	Offset int    `json:"offset"`
	Total  int    `json:"total"`
}

// GetPlaylists returns a map of user playlists
func GetPlaylists(token, user string) (map[string]string, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/users/%s/playlists", user)
	tokenHdr := "Bearer " + token
	queryMore := true
	playlistMAP := make(map[string]string)
	for queryMore {
		playlists, err := queryUser(tokenHdr, url)
		if err != nil {
			return nil, err
		}

		for _, playLST := range playlists.Items {
			playlistMAP[playLST.Name] = playLST.ID
		}
		url = playlists.Next
		if playlists.Next == "" {
			queryMore = false
		}
	}
	return playlistMAP, nil
}
func queryUser(tokenHdr, url string) (userPlaylists, error) {

	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	contentTypeHdr := fmt.Sprint("application/json")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return userPlaylists{}, err
	}
	req.Header.Add("Accept", contentTypeHdr)
	req.Header.Add("Content-Type", contentTypeHdr)
	req.Header.Add("Authorization", tokenHdr)

	resp, err := client.Do(req)
	if err != nil {
		return userPlaylists{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return userPlaylists{}, err
	}
	playlists := userPlaylists{}

	if err := json.Unmarshal(body, &playlists); err != nil {
		return userPlaylists{}, err
	}
	return playlists, nil

}
