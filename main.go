package main

import (
	"encoding/json"
	"fmt"
	"genrefinder/auth"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// SpotifyPlaylist data from spotify playlist
type SpotifyPlaylist struct {
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

func main() {
	spotify := &auth.AppSecrets{}
	token := spotify.GetToken()
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
	req.Header.Add("Accept", contentTypeHdr)
	req.Header.Add("Content-Type", contentTypeHdr)
	req.Header.Add("Authorization", tokenHdr)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	result := SpotifyPlaylist{}
	json.Unmarshal(body, &result)
	fmt.Println(result.Href, result.Next, result.Total)
	for _, v := range result.Items {
		fmt.Println(v.Track.Href)
		for _, artist := range v.Track.Artists {
			fmt.Println(artist.Name)
		}
	}

}
