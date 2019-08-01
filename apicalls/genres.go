// Package apicalls is for storing all the different calls to spotify apis
package apicalls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GenresData genres from spotify playlist
type genresData struct {
	Name   string   `json:"name"`
	Genres []string `json:"genres"`
}

// GetGenreMap retrieves list of genres of a given artist list
func GetGenreMap(token string, artistURLS []string) (map[string][]string, error) {
	tokenHdr := "Bearer " + token
	artistGenres := make(map[string][]string)
	for _, apiURL := range artistURLS {
		tempGenres, err := getGenres(tokenHdr, apiURL)
		if err != nil {
			return nil, err
		}
		artistGenres[tempGenres.Name] = tempGenres.Genres
	}
	return artistGenres, nil
}

// getGenres retrieves genres for a given artist
func getGenres(tokenHdr, url string) (genresData, error) {

	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	contentTypeHdr := fmt.Sprint("application/json")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return genresData{}, err
	}
	req.Header.Add("Accept", contentTypeHdr)
	req.Header.Add("Content-Type", contentTypeHdr)
	req.Header.Add("Authorization", tokenHdr)

	resp, err := client.Do(req)
	if err != nil {
		return genresData{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return genresData{}, err
	}
	genres := genresData{}

	if err := json.Unmarshal(body, &genres); err != nil {
		return genresData{}, err
	}
	return genres, nil

}
