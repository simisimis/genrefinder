// Package spotify is for storing all the different calls to spotify apis
package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// genresData is a struct to receive json data about artist genres
type genresData struct {
	Name   string   `json:"name"`
	Genres []string `json:"genres"`
}

// GetGenreMap retrieves an array of genres per artist
func GetGenreMap(token string, artists map[string]Artist) (map[string]Artist, error) {
	tokenHdr := "Bearer " + token

	for id, artist := range artists {
		genreList, err := getGenres(tokenHdr, artist.Href)
		if err != nil {
			return artists, err
		}
		fmt.Println(genreList)
		artist.Genres = genreList
		artists[id] = artist
	}
	return artists, nil
}

// getGenres retrieves a list of genres for given artist
func getGenres(tokenHdr, url string) ([]string, error) {

	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	contentTypeHdr := fmt.Sprint("application/json")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []string{}, err
	}
	req.Header.Add("Accept", contentTypeHdr)
	req.Header.Add("Content-Type", contentTypeHdr)
	req.Header.Add("Authorization", tokenHdr)

	resp, err := client.Do(req)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	genres := genresData{}

	if err := json.Unmarshal(body, &genres); err != nil {
		return []string{}, err
	}
	return genres.Genres, nil

}
